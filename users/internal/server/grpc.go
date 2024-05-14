package server

import (
	"context"

	"github.com/Prokopevs/ccc/schema"
	"github.com/Prokopevs/ccc/users/internal/model"
)

type Service interface {
	AddUser(ctx context.Context, user *model.UserReq) error
	GetUser(ctx context.Context, id int) (*model.UserRes, bool, error)
	IsUserWithIdExists(ctx context.Context, id int) (bool, error)
	GetUserReferrals(ctx context.Context, id int) ([]*model.UserReferrals, bool, error)
	GetUsers(ctx context.Context) ([]*model.UserRes, error)
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
