package pg

import (
	"context"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type db struct {
	db *sqlx.DB
}

func Connect(ctx context.Context, addr string) (*db, error) {
	d, err := sqlx.ConnectContext(ctx, "postgres", addr)
	if err != nil {
		return nil, err
	}

	return &db{
		db: d,
	}, nil
}


