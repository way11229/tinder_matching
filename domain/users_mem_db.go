package domain

import (
	"context"

	"github.com/google/uuid"
)

type UsersMemDbCreate struct {
	Name                string
	Height              uint32
	RemainNumberOfDates uint32
}

type UsersMemDbUpdate struct {
	Id                  uuid.UUID
	Name                string
	Height              uint32
	RemainNumberOfDates uint32
}

type UsersMemDbHeightSearch struct {
	Limit *int
	Bound uint32
}

type UsersMemDB interface {
	Create(ctx context.Context, input *UsersMemDbCreate) (uuid.NullUUID, error)
	Update(ctx context.Context, input *UsersMemDbUpdate) error
	UpdateBatch(ctx context.Context, input []*UsersMemDbUpdate) error
	DeleteById(ctx context.Context, id uuid.UUID) error
	DeleteByIds(ctx context.Context, ids []uuid.UUID) error
	GetById(ctx context.Context, id uuid.UUID) (*User, error)
	ListByHeightUpperBoundWithoutEqual(ctx context.Context, search *UsersMemDbHeightSearch) ([]*User, error)
	ListByHeightLowerBoundWithoutEqual(ctx context.Context, search *UsersMemDbHeightSearch) ([]*User, error)
}

type UsersMaleMemDB interface {
	UsersMemDB
}

type UsersFemaleMemDB interface {
	UsersMemDB
}
