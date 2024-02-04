package domain

import (
	"context"

	"github.com/google/uuid"
)

type CreateUser struct {
	Name                string
	Height              uint32
	Gender              UserGender
	RemainNumberOfDates uint32
}

type CreateUserResp struct {
	UserId  uuid.UUID
	Matches []*User
}

type UserListMatchesByUserId struct {
	UserId uuid.UUID
	Limit  int
}

type UserService interface {
	CreateUserAndListMatches(ctx context.Context, input *CreateUser) (*CreateUserResp, error)
	DeleteUserById(ctx context.Context, id uuid.UUID) error
	ListMatchesByUserId(ctx context.Context, input *UserListMatchesByUserId) ([]*User, error)
}
