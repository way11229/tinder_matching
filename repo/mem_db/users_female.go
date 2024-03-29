package memdb

import (
	"context"
	"log"

	"github.com/google/uuid"
	"github.com/hashicorp/go-memdb"
	"github.com/way11229/tinder_matching/domain"
)

const (
	USERS_FEMALE_MEM_DB_TABLE_NAME = "users_female"
)

type UsersFemaleMemDB struct {
	memdb *MemDB
}

func getUsersFemaleTableSchema() *memdb.TableSchema {
	return &memdb.TableSchema{
		Name: USERS_FEMALE_MEM_DB_TABLE_NAME,
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

func NewUsersFemaleMemDB(db *MemDB) domain.UsersFemaleMemDB {
	return &UsersFemaleMemDB{
		memdb: db,
	}
}

func (u *UsersFemaleMemDB) Create(ctx context.Context, input *domain.UsersMemDbCreate) (uuid.NullUUID, error) {
	userId, err := uuid.NewUUID()
	if err != nil {
		log.Printf("UsersFemaleMemDB, create, get user id error: %v", err)
		return uuid.NullUUID{}, domain.ErrorInternalServerError
	}

	user := &User{
		Id:                  userId.String(),
		Name:                input.Name,
		Height:              input.Height,
		RemainNumberOfDates: input.RemainNumberOfDates,
	}

	err = u.memdb.ExecTrx(true, func(t *memdb.Txn) error {
		if err := t.Insert(USERS_FEMALE_MEM_DB_TABLE_NAME, user); err != nil {
			log.Printf("UsersFemaleMemDB, create, insert error: %v, input: %v", err, user)
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

func (u *UsersFemaleMemDB) UpdateBatch(ctx context.Context, input []*domain.UsersMemDbUpdate) error {
	users := []*User{}
	for _, e := range input {
		users = append(users, &User{
			Id:                  e.Id.String(),
			Name:                e.Name,
			Height:              e.Height,
			RemainNumberOfDates: e.RemainNumberOfDates,
		})
	}

	err := u.memdb.ExecTrx(true, func(t *memdb.Txn) error {
		for _, e := range users {
			if err := t.Insert(USERS_FEMALE_MEM_DB_TABLE_NAME, e); err != nil {
				log.Printf("UsersFemaleMemDB, updateBatch, insert error: %v, input: %v", err, e)
				return domain.ErrorInternalServerError
			}
		}

		return nil
	})
	if err != nil {
		return err
	}

	return nil
}

func (u *UsersFemaleMemDB) DeleteById(ctx context.Context, id uuid.UUID) error {
	user := &User{Id: id.String()}
	err := u.memdb.ExecTrx(true, func(t *memdb.Txn) error {
		if err := t.Delete(USERS_FEMALE_MEM_DB_TABLE_NAME, user); err != nil {
			log.Printf("UsersFemaleMemDB, deleteById, delete error: %v, input: %v", err, user)
			return domain.ErrorInternalServerError
		}
		return nil
	})

	if err != nil {
		return err
	}

	return nil
}

func (u *UsersFemaleMemDB) DeleteByIds(ctx context.Context, ids []uuid.UUID) error {
	users := []*User{}
	for _, e := range ids {
		users = append(users, &User{Id: e.String()})
	}

	err := u.memdb.ExecTrx(true, func(t *memdb.Txn) error {
		for _, e := range users {
			if err := t.Delete(USERS_FEMALE_MEM_DB_TABLE_NAME, e); err != nil {
				log.Printf("UsersFemaleMemDB, deleteByIds, delete error: %v, input: %v", err, e)
				return domain.ErrorInternalServerError
			}
		}
		return nil
	})

	if err != nil {
		return err
	}

	return nil
}

func (u *UsersFemaleMemDB) GetById(ctx context.Context, id uuid.UUID) (*domain.User, error) {
	user := &User{}

	err := u.memdb.ExecTrx(false, func(t *memdb.Txn) error {
		raw, err := t.First(USERS_FEMALE_MEM_DB_TABLE_NAME, "id", id.String())
		if err != nil {
			log.Printf("UsersFemaleMemDB, getById, first error: %v, input: %v", err, id)
			return domain.ErrorInternalServerError
		}

		if raw == nil {
			return domain.ErrorRecordNotFound
		}

		assignUser, ok := raw.(*User)
		if !ok {
			log.Printf("UsersFemaleMemDB, getById, assign type error, raw data: %v", raw)
			return domain.ErrorInternalServerError
		}

		user = assignUser

		return nil
	})
	if err != nil {
		return nil, err
	}

	return u.remodelUsers([]*User{user})[0], nil
}

func (u *UsersFemaleMemDB) ListByHeightUpperBoundWithoutEqual(ctx context.Context, search *domain.UsersMemDbHeightSearch) ([]*domain.User, error) {
	users := []*User{}

	err := u.memdb.ExecTrx(false, func(t *memdb.Txn) error {
		it, err := t.ReverseLowerBound(USERS_FEMALE_MEM_DB_TABLE_NAME, "height", search.Bound)
		if err != nil {
			log.Printf("UsersFemaleMemDB, ListByHeightUpperBoundWithoutEqual, reverseLowerBound, error: %v, input: %d", err, search.Bound)
			return domain.ErrorInternalServerError
		}

		for raw := it.Next(); raw != nil; raw = it.Next() {
			assignUser, ok := raw.(*User)
			if !ok {
				log.Printf("UsersFemaleMemDB, ListByHeightUpperBoundWithoutEqual, assign type error, raw data: %v", raw)
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

	return u.remodelUsers(users), nil
}

func (u *UsersFemaleMemDB) ListByHeightLowerBoundWithoutEqual(ctx context.Context, search *domain.UsersMemDbHeightSearch) ([]*domain.User, error) {
	users := []*User{}

	err := u.memdb.ExecTrx(false, func(t *memdb.Txn) error {
		it, err := t.LowerBound(USERS_FEMALE_MEM_DB_TABLE_NAME, "height", search.Bound)
		if err != nil {
			log.Printf("UsersFemaleMemDB, ListByHeightLowerBoundWithoutEqual, lowerBound, error: %v, input: %d", err, search.Bound)
			return domain.ErrorInternalServerError
		}

		for raw := it.Next(); raw != nil; raw = it.Next() {
			assignUser, ok := raw.(*User)
			if !ok {
				log.Printf("UsersFemaleMemDB, ListByHeightLowerBoundWithoutEqual, assign type error, raw data: %v", raw)
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

	return u.remodelUsers(users), nil
}

func (u *UsersFemaleMemDB) ReduceNumberOfDatesOfUserAndMatchesTrx(ctx context.Context, input *domain.ReduceNumberOfDatesOfUserAndMatchesTrx) error {
	err := u.memdb.ExecTrx(false, func(t *memdb.Txn) error {
		if input.UserUpdate != nil {
			if err := t.Insert(USERS_FEMALE_MEM_DB_TABLE_NAME, &User{
				Id:                  input.UserUpdate.Id.String(),
				Name:                input.UserUpdate.Name,
				Height:              input.UserUpdate.Height,
				RemainNumberOfDates: input.UserUpdate.RemainNumberOfDates,
			}); err != nil {
				log.Printf("UsersFemaleMemDB, reduceNumberOfDatesOfUserAndMatchesTrx, insert error: %v, input: %v", err, input.UserUpdate)
				return domain.ErrorInternalServerError
			}

			if input.UserDeleteId.Valid {
				if err := t.Delete(USERS_FEMALE_MEM_DB_TABLE_NAME, &User{
					Id: input.UserDeleteId.UUID.String(),
				}); err != nil {
					log.Printf("UsersFemaleMemDB, reduceNumberOfDatesOfUserAndMatchesTrx, delete error: %v, input: %v", err, input.UserDeleteId)
					return domain.ErrorInternalServerError
				}
			}

			if len(input.UpdateMatches) > 0 {
				for _, e := range input.UpdateMatches {
					if err := t.Insert(USERS_MALE_MEM_DB_TABLE_NAME, &User{
						Id:                  e.Id.String(),
						Name:                e.Name,
						Height:              e.Height,
						RemainNumberOfDates: e.RemainNumberOfDates,
					}); err != nil {
						log.Printf("UsersFemaleMemDB, reduceNumberOfDatesOfUserAndMatchesTrx, insert error: %v, input: %v", err, e)
						return domain.ErrorInternalServerError
					}
				}
			}

			if len(input.DeleteMatchesIds) > 0 {
				for _, e := range input.DeleteMatchesIds {
					if err := t.Delete(USERS_MALE_MEM_DB_TABLE_NAME, &User{
						Id: e.String(),
					}); err != nil {
						log.Printf("UsersFemaleMemDB, reduceNumberOfDatesOfUserAndMatchesTrx, delete error: %v, input: %v", err, input.UserDeleteId)
						return domain.ErrorInternalServerError
					}
				}
			}
		}

		return nil
	})
	if err != nil {
		return err
	}

	return nil
}

func (u *UsersFemaleMemDB) remodelUsers(data []*User) []*domain.User {
	rtn := []*domain.User{}
	for _, e := range data {
		id, _ := uuid.Parse(e.Id)
		rtn = append(rtn, &domain.User{
			Id:                  id,
			Name:                e.Name,
			Height:              e.Height,
			Gender:              domain.USER_GENDER_FEMALE,
			RemainNumberOfDates: e.RemainNumberOfDates,
		})
	}

	return rtn
}
