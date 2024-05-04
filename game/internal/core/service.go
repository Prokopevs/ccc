package core 

import (
	"context"
)

type DB interface {
	UpdateScore(context.Context, int, int) (error)
	UpdateMultiplicator(context.Context, int, string) (error)
}

type ServiceImpl struct {
	db DB
}

func NewService(db DB) *ServiceImpl {
	return &ServiceImpl{
		db: db,
	}
}