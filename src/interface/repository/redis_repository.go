package repository

import (
	"api_client/usecase/repository"

	"github.com/go-redis/redis/v8"
)

type redisRepository struct {
	redis *redis.Client
}

func NewRedisRepository(r *redis.Client) repository.RedisRepository {
	return &redisRepository{r}
}

func (rr *redisRepository) GetCookie(accountID string) error {
	// get cookie store in redis
	return nil
}
