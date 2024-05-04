package pg

import (
	"context"
	"database/sql"
)

type Game struct {
	OwnerId    int    `db:"ownerId"`
	Score      string `db:"score"`
	GasStorage int    `db:"gasStorage"`
	GasMining  string `db:"gasMining"`
	Protection int    `db:"protection"`
}

func (d *db) UpdateScore(ctx context.Context, id int, score int) error {
	const UpdateScoreQ = "UPDATE game SET score = score + $1 WHERE ownerId = $2;"

	_, err := d.db.ExecContext(ctx, UpdateScoreQ, id, score)

	return err
}

func (d *db) UpdateMultiplicator(ctx context.Context, id int, mType string) error {
	const (
		MultiplicatorQ = "select $1 from game where ownerId=$2"
		getScoreQ      = "select score from game where ownerId=$1"
		updateScoreQ  = "UPDATE game SET score = score - $1 WHERE ownerId = $2;"
		updateMultiplicatorQ  = "UPDATE game SET $1 = $1 + 1 WHERE ownerId = $2;"
	)

	prices := d.GetPrices()
	tx, err := d.db.BeginTxx(ctx, &sql.TxOptions{})
	if err != nil {
		return err
	}

	var m int
	err = tx.GetContext(ctx, m, MultiplicatorQ, mType, id) // get multiplicator level
	if err != nil {
		tx.Rollback()
		return err
	}

	var s int
	err = tx.GetContext(ctx, s, getScoreQ, id) // get score
	if err != nil {
		tx.Rollback()
		return err
	}

	if s < prices[m+1] {
		err = ErrNoEnoughScore
		tx.Rollback()
		return err
	}

	_, err = tx.ExecContext(ctx, updateScoreQ, prices[m+1], id)
	if err != nil {
		tx.Rollback()
		return err
	}

	_, err = tx.ExecContext(ctx, updateMultiplicatorQ, mType, id)
	if err != nil {
		tx.Rollback()
		return err
	}

	err = tx.Commit()

	return err
}
