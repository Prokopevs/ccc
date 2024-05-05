package core

import (
	"context"

	"github.com/Prokopevs/ccc/game/internal/model"
)

type DB interface {
	GetGame(context.Context, int) (*model.Game, error)
	UpdateScore(context.Context, *model.Score) (error)
	UpdateMultiplicator(context.Context, *model.MultipUpdate) (error)
}

type ServiceImpl struct {
	db DB
}

func NewService(db DB) *ServiceImpl {
	return &ServiceImpl{
		db: db,
	}
}