package repository

import (
	"context"
	"time"

	"go-bank/domain"
	"go-bank/internal/config"

	"github.com/redis/go-redis/v9"
)

type redisCacheRepository struct {
	rdb *redis.Client
}

func NewRedisClient(cnf *config.Config) domain.CacheRepository {
	return &redisCacheRepository{
		rdb: redis.NewClient(&redis.Options{
			Addr:     cnf.Redis.Addr,
			Password: cnf.Redis.Pass,
			DB:       0,
		}),
	}
}

func (r *redisCacheRepository) Set(key string, entry []byte) error {
	return r.rdb.Set(context.Background(), key, entry, 15*time.Minute).Err()
}

func (r *redisCacheRepository) Get(key string) ([]byte, error) {
	return r.rdb.Get(context.Background(), key).Bytes()
}
