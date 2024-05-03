package core 

import (
	"context"
	"github.com/Prokopevs/ccc/users/internal/pg"
)

type DB interface {
	AddUser(context.Context, *pg.UserReq) (error)
	GetUser(context.Context, int) (*pg.UserRes, error)
	IsUserWithIdExists(context.Context, int) (bool, error)
}

type ServiceImpl struct {
	db DB
}

func NewServiceImpl(db DB) *ServiceImpl {
	return &ServiceImpl{
		db: db,
	}
}