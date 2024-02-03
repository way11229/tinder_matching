package memdb

import (
	"context"
	"log"

	"github.com/google/uuid"
	"github.com/hashicorp/go-memdb"
	"github.com/way11229/tinder_matching/domain"
)

const (
	USERS_MALE_MEM_DB_TABLE_NAME = "users_male"
)

type UsersMaleMemDB struct {
	memdb *MemDB
}

func getUsersMaleTableSchema() *memdb.TableSchema {
	return &memdb.TableSchema{
		Name: USERS_MALE_MEM_DB_TABLE_NAME,
		Indexes: map[string]*memdb.IndexSchema{
			"id": {
				Name:         "id",
				AllowMissing: false,
				Unique:       true,
				Indexer: &memdb.UUIDFieldIndex{
					Field: "Id",
				},
			},
			"name": {
				Name:         "name",
				AllowMissing: false,
				Unique:       false,
				Indexer: &memdb.StringFieldIndex{
					Field: "Name",
				},
			},
			"height": {
				Name:         "height",
				AllowMissing: false,
				Unique:       false,
				Indexer: &memdb.UintFieldIndex{
					Field: "Height",
				},
			},
			"remain_number_of_dates": {
				Name:         "remain_number_of_dates",
				AllowMissing: false,
				Unique:       false,
				Indexer: &memdb.UintFieldIndex{
					Field: "RemainNumberOfDates",
				},
			},
		},
	}
}

func NewUsersMaleMemDB(db *MemDB) domain.UsersMaleMemDB {
	return &UsersMaleMemDB{
		memdb: db,
	}
}

func (u *UsersMaleMemDB) Create(ctx context.Context, input *domain.UsersMemDbCreate) (uuid.NullUUID, error) {
	userId, err := uuid.NewUUID()
	if err != nil {
		log.Printf("UsersMaleMemDB, create, get user id error: %v", err)
		return uuid.NullUUID{}, domain.ErrorInternalServerError
	}

	user := &User{
		Id:                  userId,
		Name:                input.Name,
		Height:              input.Height,
		RemainNumberOfDates: input.RemainNumberOfDates,
	}

	err = u.memdb.ExecTrx(true, func(t *memdb.Txn) error {
		if err := t.Insert(USERS_MALE_MEM_DB_TABLE_NAME, user); err != nil {
			log.Printf("UsersMaleMemDB, create, insert error: %v, input: %v", err, user)
			return domain.ErrorInternalServerError
		}

		return nil
	})
	if err != nil {
		return uuid.NullUUID{}, err
	}

	return uuid.NullUUID{
		Valid: true,
		UUID:  userId,
	}, nil
}

func (u *UsersMaleMemDB) UpdateById(ctx context.Context, input *domain.UsersMemDbUpdateById) error {
	user := &User{
		Id:                  input.Id,
		Name:                input.Name,
		Height:              input.Height,
		RemainNumberOfDates: input.RemainNumberOfDates,
	}

	err := u.memdb.ExecTrx(true, func(t *memdb.Txn) error {
		if err := t.Insert(USERS_MALE_MEM_DB_TABLE_NAME, user); err != nil {
			log.Printf("UsersMaleMemDB, updateById, insert error: %v, input: %v", err, user)
			return domain.ErrorInternalServerError
		}

		return nil
	})
	if err != nil {
		return err
	}

	return nil
}

func (u *UsersMaleMemDB) DeleteById(ctx context.Context, id uuid.UUID) error {
	user := &User{Id: id}
	err := u.memdb.ExecTrx(true, func(t *memdb.Txn) error {
		if err := t.Delete(USERS_MALE_MEM_DB_TABLE_NAME, user); err != nil {
			log.Printf("UsersMaleMemDB, deleteById, delete error: %v, input: %v", err, user)
			return domain.ErrorInternalServerError
		}
		return nil
	})

	if err != nil {
		return err
	}

	return nil
}

func (u *UsersMaleMemDB) GetById(ctx context.Context, id uuid.UUID) (*domain.User, error) {
	user := &User{}

	err := u.memdb.ExecTrx(false, func(t *memdb.Txn) error {
		raw, err := t.First(USERS_MALE_MEM_DB_TABLE_NAME, "id", id)
		if err != nil {
			log.Printf("UsersMaleMemDB, getById, first error: %v, input: %v", err, id)
			return domain.ErrorInternalServerError
		}

		if raw == nil {
			return domain.ErrorRecordNotFound
		}

		assignUser, ok := raw.(*User)
		if !ok {
			log.Printf("UsersMaleMemDB, getById, assign type error, raw data: %v", raw)
			return domain.ErrorInternalServerError
		}

		user = assignUser

		return nil
	})
	if err != nil {
		return nil, err
	}

	return u.remodelUser([]*User{user})[0], nil
}

func (u *UsersMaleMemDB) ListByHeightUpperBoundWithoutEqual(ctx context.Context, search *domain.UsersMemDbHeightSearch) ([]*domain.User, error) {
	users := []*User{}

	err := u.memdb.ExecTrx(false, func(t *memdb.Txn) error {
		it, err := t.ReverseLowerBound(USERS_MALE_MEM_DB_TABLE_NAME, "height", search.Bound)
		if err != nil {
			log.Printf("UsersMaleMemDB, ListByHeightUpperBoundWithoutEqual, reverseLowerBound, error: %v, input: %d", err, search.Bound)
			return domain.ErrorInternalServerError
		}

		for raw := it.Next(); raw != nil; raw = it.Next() {
			assignUser, ok := raw.(*User)
			if !ok {
				log.Printf("UsersMaleMemDB, ListByHeightUpperBoundWithoutEqual, assign type error, raw data: %v", raw)
				return domain.ErrorInternalServerError
			}

			if assignUser.Height == search.Bound {
				continue
			}

			users = append(users, assignUser)
			if search.Limit != nil && len(users) == *search.Limit {
				break
			}
		}

		return nil
	})
	if err != nil {
		return nil, err
	}

	return u.remodelUser(users), nil
}

func (u *UsersMaleMemDB) ListByHeightLowerBoundWithoutEqual(ctx context.Context, search *domain.UsersMemDbHeightSearch) ([]*domain.User, error) {
	users := []*User{}

	err := u.memdb.ExecTrx(false, func(t *memdb.Txn) error {
		it, err := t.LowerBound(USERS_MALE_MEM_DB_TABLE_NAME, "height", search.Bound)
		if err != nil {
			log.Printf("UsersMaleMemDB, ListByHeightLowerBoundWithoutEqual, lowerBound, error: %v, input: %d", err, search.Bound)
			return domain.ErrorInternalServerError
		}

		for raw := it.Next(); raw != nil; raw = it.Next() {
			assignUser, ok := raw.(*User)
			if !ok {
				log.Printf("UsersMaleMemDB, ListByHeightLowerBoundWithoutEqual, assign type error, raw data: %v", raw)
				return domain.ErrorInternalServerError
			}

			if assignUser.Height == search.Bound {
				continue
			}

			users = append(users, assignUser)
			if search.Limit != nil && len(users) == *search.Limit {
				break
			}
		}

		return nil
	})
	if err != nil {
		return nil, err
	}

	return u.remodelUser(users), nil
}

func (u *UsersMaleMemDB) remodelUser(data []*User) []*domain.User {
	rtn := []*domain.User{}
	for _, e := range data {
		rtn = append(rtn, &domain.User{
			Id:                  e.Id,
			Name:                e.Name,
			Height:              e.Height,
			Gender:              domain.USER_GENDER_MALE,
			RemainNumberOfDates: e.RemainNumberOfDates,
		})
	}

	return rtn
}