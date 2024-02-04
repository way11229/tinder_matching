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
		UserId: newUserId,
		Limit:  nil,
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
		UserId: input.UserId,
		Limit:  &limit,
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
				Users:      matches,
				UsersMemDb: u.usersFemaleMemDB,
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
				Users:      matches,
				UsersMemDb: u.usersMaleMemDB,
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
	Users      []*domain.User
	UsersMemDb domain.UsersMemDB
}

func (u *UserService) reduceUserNumberOfDatesAndDeleteUserWhenBecomeToZero(ctx context.Context, input *reduceUserNumberOfDatesAndDeleteUserWhenBecomeToZeroType) error {
	updateUsers := []*domain.UsersMemDbUpdate{}
	deleteUserIds := []uuid.UUID{}

	for _, e := range input.Users {
		remain := e.RemainNumberOfDates - 1
		if remain <= 0 {
			deleteUserIds = append(deleteUserIds, e.Id)
			continue
		}

		updateUsers = append(updateUsers, &domain.UsersMemDbUpdate{
			Id:                  e.Id,
			Name:                e.Name,
			Height:              e.Height,
			RemainNumberOfDates: remain,
		})
	}

	if len(updateUsers) > 0 {
		if err := input.UsersMemDb.UpdateBatch(ctx, updateUsers); err != nil {
			return err
		}
	}

	if len(deleteUserIds) > 0 {
		if err := input.UsersMemDb.DeleteByIds(ctx, deleteUserIds); err != nil {
			return err
		}
	}

	return nil
}
