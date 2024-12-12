package repository

import (
	"github.com/go-park-mail-ru/2024_2_GOATS/auth_service/internal/auth/service"
	"github.com/go-redis/redis/v8"
)

var _ service.AuthRepositoryInterface = (*AuthRepository)(nil)

// AuthRepository is a auth_service repo layer struct
type AuthRepository struct {
	Redis *redis.Client
}

// NewAuthRepository returns an instance of AuthRepositoryInterface
func NewAuthRepository(rdb *redis.Client) service.AuthRepositoryInterface {
	return &AuthRepository{
		Redis: rdb,
	}
}
