package domain

import (
	"context"

	"github.com/google/uuid"
)

type UsersMemDbCreate struct {
	Name                string
	Height              uint32
	Gender              UserGender
	RemainNumberOfDates uint32
}

type UsersMemDbUpdateById struct {
	Id                  uuid.UUID
	Name                string
	Height              uint32
	Gender              UserGender
	RemainNumberOfDates uint32
}

type UsersMemDbHeightSearch struct {
	Limit *int
	Bound uint32
}

type UsersMemDB interface {
	Create(ctx context.Context, input *UsersMemDbCreate) (uuid.NullUUID, error)
	UpdateById(ctx context.Context, input *UsersMemDbUpdateById) error
	DeleteById(ctx context.Context, id uuid.UUID) error
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
