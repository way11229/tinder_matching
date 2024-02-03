package domain

import (
	"context"

	"github.com/google/uuid"
)

type UsersMemDbCreate struct {
	Name          string
	Height        uint32
	Gender        UserGender
	NumberOfDates uint32
	DateStatus    bool
}

type UsersMemDbUpdateById struct {
	Id            uuid.UUID
	Name          string
	Height        uint32
	Gender        UserGender
	NumberOfDates uint32
	DateStatus    bool
}

type UsersMemDbSearch struct {
	Gender           *UserGender
	DateStatus       *bool
	HeightUpperBound *uint
	HeightLowerBound *uint
}

type UsersMemDB interface {
	Create(ctx context.Context, input *UsersMemDbCreate) (uuid.NullUUID, error)
	GetById(ctx context.Context, id uuid.UUID) (*User, error)
	UpdateById(ctx context.Context, input *UsersMemDbUpdateById) error
	ListBySearch(ctx context.Context, search *UsersMemDbSearch) ([]*User, error)
}
