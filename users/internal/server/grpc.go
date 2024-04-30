package server

import (
	"context"

	"github.com/Prokopevs/ccc/schema"
	"github.com/Prokopevs/ccc/users/internal/core"
)

type Service interface {
	AddUser(ctx context.Context, user *core.User) error
	GetUser(ctx context.Context, id int) (*core.User, bool, error)
	IsUserWithIdExists(ctx context.Context, id int) (bool, error)
}

type GRPC struct {
	schema.UnimplementedUsersServer

	service Service
}

func NewGRPC(usersService Service) *GRPC {
	return &GRPC{
		service: usersService,
	}
}
