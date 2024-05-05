package core

import (
	"context"

	"github.com/Prokopevs/ccc/game/internal/model"
)

func (s *ServiceImpl) GetGame(ctx context.Context, id int) (*model.Game, error) {
	game, err := s.db.GetGame(ctx, id)
	if err != nil {
		return &model.Game{}, err
	}

	return game, err
}

func (s *ServiceImpl) UpdateScore(ctx context.Context, score *model.Score) (Code, error) {
	err := s.db.UpdateScore(ctx, score)
	if err != nil {
		return CodeDBFail, err
	}

	return CodeOK, err 
}

func (s *ServiceImpl) UpdateMultiplicator(ctx context.Context, MultipUpdate *model.MultipUpdate) (Code, error) {
	exist := s.CheckCorectMultiplicatorType(MultipUpdate.NameType) 
	if !exist {
		return CodeNoMultiplicator, ErrNoSuchMultiplicator
	}

	err := s.db.UpdateMultiplicator(ctx, MultipUpdate)
	if err != nil {
		return CodeDBFail, err
	}

	return CodeOK, err
}