package pg

import (
	"context"
	"database/sql"

	"github.com/Prokopevs/ccc/game/internal/model"
)

func (d *db) GetGame(ctx context.Context, id int) (*model.Game, error) {
	const q = "select * from game where ownerId=$1"

	game := model.Game{}
	err := d.db.SelectContext(ctx, &game, q, id)

	return &game, err
}

func (d *db) UpdateScore(ctx context.Context, score *model.Score) error {
	const UpdateScoreQ = "UPDATE game SET score = score + $1 WHERE ownerId = $2;"

	_, err := d.db.ExecContext(ctx, UpdateScoreQ, score.Id, score.Score)

	return err
}

func (d *db) UpdateMultiplicator(ctx context.Context, MultipUpdate *model.MultipUpdate) error {
	const (
		MultiplicatorQ = "select $1 from game where ownerId=$2"
		getScoreQ      = "select score from game where ownerId=$1"
		updateScoreQ  = "UPDATE game SET score = score - $1 WHERE ownerId = $2;"
		updateMultiplicatorQ  = "UPDATE game SET $1 = $1 + 1 WHERE ownerId = $2;"
	)

	prices := d.getPrices()
	tx, err := d.db.BeginTxx(ctx, &sql.TxOptions{})
	if err != nil {
		return err
	}

	var m int
	err = tx.GetContext(ctx, m, MultiplicatorQ, MultipUpdate.NameType, MultipUpdate.Id) // get multiplicator level
	if err != nil {
		tx.Rollback()
		return err
	}

	var s int
	err = tx.GetContext(ctx, s, getScoreQ, MultipUpdate.Id) // get score
	if err != nil {
		tx.Rollback()
		return err
	}

	if s < prices[m+1] {
		err = ErrNoEnoughScore
		tx.Rollback()
		return err
	}

	_, err = tx.ExecContext(ctx, updateScoreQ, prices[m+1], MultipUpdate.Id)
	if err != nil {
		tx.Rollback()
		return err
	}

	_, err = tx.ExecContext(ctx, updateMultiplicatorQ, MultipUpdate.NameType, MultipUpdate.Id)
	if err != nil {
		tx.Rollback()
		return err
	}

	err = tx.Commit()

	return err
}
