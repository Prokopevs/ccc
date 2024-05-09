package redis

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
)

var cacheExpiration = 1 * time.Minute

func (r *rbd) IsSignatureExist(ctx context.Context, key string, signature string) (bool, error) {
	val, err := r.rbd.Get(ctx, key).Result()
	if err == redis.Nil {
		err = r.SetValue(ctx, key, signature)
		if err != nil {
			return false, err
		}

		return false, nil
	} 
	if err != nil {
		return false, err
	}

	if val == signature {
		return true, nil
	}

	err = r.SetValue(ctx, key, signature)
	if err != nil {
		return false, err
	}
	return false, nil
}

func (r *rbd) SetValue(ctx context.Context, key string, signature string) error {
	err := r.rbd.Set(ctx, key, signature, cacheExpiration).Err()
	if err != nil {
		return err
	}
	return nil
}
