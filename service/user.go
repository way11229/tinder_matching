package service

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/way11229/tinder_matching/domain"
)

type UserService struct {
	usersMaleMemDB   domain.UsersMaleMemDB
	usersFemaleMemDB domain.UsersFemaleMemDB
}

func NewUserService(
	usersMaleMemDB domain.UsersMaleMemDB,
	usersFemaleMemDB domain.UsersFemaleMemDB,
) domain.UserService {
	return &UserService{
		usersMaleMemDB:   usersMaleMemDB,
		usersFemaleMemDB: usersFemaleMemDB,
	}
}

func (u *UserService) CreateUserAndListMatches(ctx context.Context, input *domain.CreateUser) (*domain.CreateUserResp, error) {
	if err := u.validateCreateUserData(input); err != nil {
		return nil, err
	}

	createUser := &domain.UsersMemDbCreate{
		Name:                input.Name,
		Height:              input.Height,
		RemainNumberOfDates: input.RemainNumberOfDates,
	}
	var newUserId uuid.UUID

	switch input.Gender {
	case domain.USER_GENDER_MALE:
		resp, err := u.usersMaleMemDB.Create(ctx, createUser)
		if err != nil {
			return nil, err
		}

		newUserId = resp.UUID
	case domain.USER_GENDER_FEMALE:
		resp, err := u.usersFemaleMemDB.Create(ctx, createUser)
		if err != nil {
			return nil, err
		}

		newUserId = resp.UUID
	default:
		return nil, domain.ErrorUserGenderInvalid
	}

	matches, err := u.listMatchesByUserId(ctx, &listMatchesByUserIdType{
		UserId:              newUserId,
		Limit:               nil,
		ReduceNumberOfDates: true,
	})
	if err != nil {
		return &domain.CreateUserResp{
			UserId:  newUserId,
			Matches: []*domain.User{},
		}, nil
	}

	return &domain.CreateUserResp{
		UserId:  newUserId,
		Matches: matches,
	}, nil
}

func (u *UserService) DeleteUserById(ctx context.Context, id uuid.UUID) error {
	user, err := u.getUserById(ctx, id)
	if err != nil {
		return err
	}

	switch user.Gender {
	case domain.USER_GENDER_MALE:
		if err := u.usersMaleMemDB.DeleteById(ctx, user.Id); err != nil {
			return err
		}
	case domain.USER_GENDER_FEMALE:
		if err := u.usersFemaleMemDB.DeleteById(ctx, user.Id); err != nil {
			return err
		}
	}

	return nil
}

func (u *UserService) ListMatchesByUserId(ctx context.Context, input *domain.UserListMatchesByUserId) ([]*domain.User, error) {
	limit := input.Limit
	if limit <= 0 {
		limit = domain.SEARCH_LIMIT_DEFAULT
	}

	matches, err := u.listMatchesByUserId(ctx, &listMatchesByUserIdType{
		UserId:              input.UserId,
		Limit:               &limit,
		ReduceNumberOfDates: false,
	})
	if err != nil {
		return nil, err
	}

	return matches, nil
}

/********************
 ********************
 ** private method **
 ********************
 ********************/

func (u *UserService) validateCreateUserData(input *domain.CreateUser) error {
	if !u.validateUserName(input.Name) {
		return domain.ErrorUserNameInvalid
	}

	if !u.validateUserHeight(input.Height) {
		return domain.ErrorUserHeightInvalid
	}

	if !u.validateUserNumberOfWantedDates(input.RemainNumberOfDates) {
		return domain.ErrorUserNumberOfWantedDatesInvalid
	}

	return nil
}

func (u *UserService) validateUserName(name string) bool {
	runeLen := len([]rune(name))
	return runeLen > 0 && runeLen <= domain.USER_NAME_LEN_MAX
}

func (u *UserService) validateUserHeight(height uint32) bool {
	return height > 0 && height <= domain.USER_HEIGHT_MAX
}

func (u *UserService) validateUserNumberOfWantedDates(number uint32) bool {
	return number > 0 && number <= domain.USER_NUMBER_OF_WANTED_DATES_MAX
}

type listMatchesByUserIdType struct {
	UserId              uuid.UUID
	Limit               *int
	ReduceNumberOfDates bool
}

func (u *UserService) listMatchesByUserId(ctx context.Context, input *listMatchesByUserIdType) ([]*domain.User, error) {
	user, err := u.getUserById(ctx, input.UserId)
	if err != nil {
		return nil, err
	}

	search := &domain.UsersMemDbHeightSearch{
		Limit: input.Limit,
		Bound: user.Height,
	}
	matches := []*domain.User{}

	switch user.Gender {
	case domain.USER_GENDER_MALE:
		resp, err := u.usersFemaleMemDB.ListByHeightUpperBoundWithoutEqual(ctx, search)
		if err != nil {
			return nil, err
		}

		matches = resp

		if input.ReduceNumberOfDates {
			// error don't affect the response
			u.reduceUserNumberOfDatesAndDeleteUserWhenBecomeToZero(ctx, &reduceUserNumberOfDatesAndDeleteUserWhenBecomeToZeroType{
				User:       user,
				Matches:    matches,
				UsersMemDb: u.usersMaleMemDB,
			})
		}
	case domain.USER_GENDER_FEMALE:
		resp, err := u.usersMaleMemDB.ListByHeightLowerBoundWithoutEqual(ctx, search)
		if err != nil {
			return nil, err
		}

		matches = resp

		if input.ReduceNumberOfDates {
			// error don't affect the response
			u.reduceUserNumberOfDatesAndDeleteUserWhenBecomeToZero(ctx, &reduceUserNumberOfDatesAndDeleteUserWhenBecomeToZeroType{
				User:       user,
				Matches:    matches,
				UsersMemDb: u.usersFemaleMemDB,
			})
		}
	}

	return matches, nil
}

func (u *UserService) getUserById(ctx context.Context, id uuid.UUID) (*domain.User, error) {
	user, err := u.usersMaleMemDB.GetById(ctx, id)
	if err != nil && !errors.Is(err, domain.ErrorRecordNotFound) {
		return nil, err
	}

	if user != nil {
		return user, nil
	}

	user, err = u.usersFemaleMemDB.GetById(ctx, id)
	if err != nil {
		return nil, err
	}

	return user, nil
}

type reduceUserNumberOfDatesAndDeleteUserWhenBecomeToZeroType struct {
	User    *domain.User
	Matches []*domain.User

	UsersMemDb domain.UsersMemDB
}

func (u *UserService) reduceUserNumberOfDatesAndDeleteUserWhenBecomeToZero(ctx context.Context, input *reduceUserNumberOfDatesAndDeleteUserWhenBecomeToZeroType) error {
	trxParams := &domain.ReduceNumberOfDatesOfUserAndMatchesTrx{
		UpdateMatches:    []*domain.UsersMemDbUpdate{},
		DeleteMatchesIds: []uuid.UUID{},
	}

	for _, e := range input.Matches {
		e.RemainNumberOfDates -= 1
		input.User.RemainNumberOfDates -= 1

		if e.RemainNumberOfDates <= 0 {
			e.RemainNumberOfDates = 0
			trxParams.DeleteMatchesIds = append(trxParams.DeleteMatchesIds, e.Id)
			continue
		}

		trxParams.UpdateMatches = append(trxParams.UpdateMatches, &domain.UsersMemDbUpdate{
			Id:                  e.Id,
			Name:                e.Name,
			Height:              e.Height,
			RemainNumberOfDates: e.RemainNumberOfDates,
		})
	}

	if input.User.RemainNumberOfDates > 0 {
		trxParams.UserUpdate = &domain.UsersMemDbUpdate{
			Id:                  input.User.Id,
			Name:                input.User.Name,
			Height:              input.User.Height,
			RemainNumberOfDates: input.User.RemainNumberOfDates,
		}
	} else {
		trxParams.UserDeleteId = uuid.NullUUID{
			Valid: true,
			UUID:  input.User.Id,
		}
	}

	if err := input.UsersMemDb.ReduceNumberOfDatesOfUserAndMatchesTrx(ctx, trxParams); err != nil {
		return err
	}

	return nil
}
