package pg

import (
	"context"

	"github.com/jmoiron/sqlx"
	_ "github.com/jackc/pgx/stdlib"
)

type db struct {
	db *sqlx.DB
}

func Connect(ctx context.Context, addr string) (*db, error) {
	d, err := sqlx.ConnectContext(ctx, "pgx", addr)
	if err != nil {
		return nil, err
	}

	return &db{
		db: d,
	}, nil
}

func Newdb(dbInstance *sqlx.DB) *db {
	return &db{db: dbInstance}
}