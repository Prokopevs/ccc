package core

import (
	"context"

	"github.com/Prokopevs/ccc/game/internal/model"
	"github.com/Prokopevs/ccc/schema"
)

type DB interface {
	GetGame(context.Context, int) (*model.Game, error)
	UpdateScore(context.Context, *model.Score) (error)
	UpdateMultiplicator(context.Context, *model.MultipUpdate) (error)
}

type ServiceImpl struct {
	usersClient schema.UsersClient
	db DB
}

func NewService(usersClient schema.UsersClient, db DB) *ServiceImpl {
	return &ServiceImpl{
		usersClient: usersClient,
		db: db,
	}
}