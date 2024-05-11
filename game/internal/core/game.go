package core

import (
	"context"
	"errors"

	"github.com/Prokopevs/ccc/game/internal/model"
	"github.com/Prokopevs/ccc/schema"
	"github.com/Prokopevs/ccc/game/internal/pg"
)

func (s *ServiceImpl) GetGame(ctx context.Context, id int) (*model.Game, error) {
	exist, err := s.usersClient.IsUserWithIdExists(ctx, &schema.IsUserWithIdExistsRequest{
		Id: int64(id),
	})
	if err != nil {
		return &model.Game{}, err
	}

	if !exist.Exists {
		return &model.Game{}, ErrNoSuchUser
	}

	game, err := s.db.GetGame(ctx, id)
	if err != nil {
		return &model.Game{}, err
	}
	
	return game, err
}

func (s *ServiceImpl) UpdateScore(ctx context.Context, score *model.Score) (Code, error) {
	exist, err := s.usersClient.IsUserWithIdExists(ctx, &schema.IsUserWithIdExistsRequest{
		Id: int64(score.Id),
	})
	if err != nil {
		return CodeInternal, err
	}

	if !exist.Exists {
		return CodeBadRequest, ErrNoSuchUser
	}

	err = s.db.UpdateScore(ctx, score)
	if err != nil {
		return CodeDBFail, err
	}

	return CodeOK, err 
}

func (s *ServiceImpl) UpdateMultiplicator(ctx context.Context, MultipUpdate *model.MultipUpdate) (Code, error) {
	exist, err := s.usersClient.IsUserWithIdExists(ctx, &schema.IsUserWithIdExistsRequest{
		Id: int64(MultipUpdate.Id),
	})
	if err != nil {
		return CodeInternal, err
	}

	if !exist.Exists {
		return CodeBadRequest, ErrNoSuchUser
	}

	corect := s.CheckCorectMultiplicatorType(MultipUpdate.NameType) 
	if !corect {
		return CodeNoMultiplicator, ErrNoSuchMultiplicator
	}

	err = s.db.UpdateMultiplicator(ctx, MultipUpdate)
	if err != nil {
		if errors.Is(err, ErrNoSuchMultiplicator) || errors.Is(err, pg.ErrMaxLevel) || errors.Is(err, pg.ErrNoEnoughScore) {
			return CodeBadRequest, err
		}

		return CodeDBFail, err
	}

	return CodeOK, err
}

func (s *ServiceImpl) GetPrices() map[int]int {
	prices := s.db.GetPrices()
	
	return prices
}