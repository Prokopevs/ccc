package redis

import (
	"github.com/redis/go-redis/v9"
)

type rbd struct {
	rbd *redis.Client
}

func Connect(connStr string) (*rbd, error) {
	opts, err := redis.ParseURL(connStr)
	if err != nil {
		return nil, err
	}

	return &rbd{
		rbd: redis.NewClient(opts),
	}, nil
}