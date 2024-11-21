package repository

import (
	"database/sql"

	"github.com/go-park-mail-ru/2024_2_GOATS/internal/app/auth/service"
	"github.com/go-redis/redis/v8"
)

var _ service.AuthRepositoryInterface = (*AuthRepo)(nil)

type AuthRepo struct {
	Database *sql.DB
	Redis    *redis.Client
}

func NewAuthRepository(db *sql.DB, rdb *redis.Client) service.AuthRepositoryInterface {
	return &AuthRepo{
		Database: db,
		Redis:    rdb,
	}
}
