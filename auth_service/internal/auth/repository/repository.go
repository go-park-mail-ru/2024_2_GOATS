package repository

import (
	"github.com/go-park-mail-ru/2024_2_GOATS/auth_service/internal/auth/service"
	"github.com/go-redis/redis/v8"
)

var _ service.AuthRepositoryInterface = (*AuthRepository)(nil)

type AuthRepository struct {
	Redis *redis.Client
}

func NewAuthRepository(rdb *redis.Client) service.AuthRepositoryInterface {
	return &AuthRepository{
		Redis: rdb,
	}
}
