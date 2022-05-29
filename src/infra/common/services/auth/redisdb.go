package auth_service

import (
	auth_service "salvegame197/golangApi/src/domain/common/services/auth"

	"github.com/go-redis/redis"
)

type RedisService struct {
	Auth   auth_service.AuthInterface
	Client *redis.Client
}

func NewRedisService(client *redis.Client) *RedisService {
	return &RedisService{
		Auth:   NewAuth(client),
		Client: client,
	}
}
