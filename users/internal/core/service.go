package core

import (
	"context"

	"github.com/Prokopevs/ccc/users/internal/model"
)

type DB interface {
	AddUser(context.Context, *model.UserReq) (error)
	GetUser(context.Context, int) (*model.UserRes, error)
	IsUserWithIdExists(context.Context, int) (bool, error)
	GetUserReferrals(context.Context, int) ([]*model.UserReferrals, error)
}

type ServiceImpl struct {
	db DB
}

func NewServiceImpl(db DB) *ServiceImpl {
	return &ServiceImpl{
		db: db,
	}
}