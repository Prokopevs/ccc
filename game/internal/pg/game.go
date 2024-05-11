package pg

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/Prokopevs/ccc/game/internal/model"
)

func (d *db) GetGame(ctx context.Context, id int) (*model.Game, error) {
	const q = "select * from game where ownerId=$1"

	game := &model.Game{}
	err := d.db.GetContext(ctx, game, q, id)
	return game, err
}

func (d *db) UpdateScore(ctx context.Context, score *model.Score) error {
	const UpdateScoreQ = "UPDATE game SET score = score + $1 WHERE ownerId = $2;"
	_, err := d.db.ExecContext(ctx, UpdateScoreQ, score.Score, score.Id)

	return err
}

func (d *db) UpdateMultiplicator(ctx context.Context, MultipUpdate *model.MultipUpdate) error {
	MultiplicatorQ := fmt.Sprintf("SELECT %s FROM game WHERE ownerId=$1", MultipUpdate.NameType)
	updateMultiplicatorQ := fmt.Sprintf("UPDATE game SET %s = %s + 1 WHERE ownerId = $1;", MultipUpdate.NameType, MultipUpdate.NameType)
	const (
		getScoreQ      = "select score from game where ownerId=$1"
		updateScoreQ  = "UPDATE game SET score = score - $1 WHERE ownerId = $2;"
	)

	prices := d.GetPrices()
	tx, err := d.db.BeginTxx(ctx, &sql.TxOptions{})
	if err != nil {
		return err
	}

	var m int
	err = tx.QueryRowContext(ctx, MultiplicatorQ, MultipUpdate.Id).Scan(&m)
	if err != nil {
		tx.Rollback()
		return err
	}

	var s int
	err = tx.QueryRowContext(ctx, getScoreQ, MultipUpdate.Id).Scan(&s)
	if err != nil {
		tx.Rollback()
		return err
	}

	if m > 6 {
		err = ErrMaxLevel
		tx.Rollback()
		return err
	}

	if s < prices[m] {
		err = ErrNoEnoughScore
		tx.Rollback()
		return err
	}

	_, err = tx.ExecContext(ctx, updateScoreQ, prices[m], MultipUpdate.Id)
	if err != nil {
		tx.Rollback()
		return err
	}

	_, err = tx.ExecContext(ctx, updateMultiplicatorQ, MultipUpdate.Id)
	if err != nil {
		tx.Rollback()
		return err
	}

	err = tx.Commit()

	return err
}
