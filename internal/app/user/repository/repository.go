package repository

import (
	"database/sql"

	"github.com/go-park-mail-ru/2024_2_GOATS/internal/app/user/service"
)

var _ service.UserRepositoryInterface = (*UserRepo)(nil)

type UserRepo struct {
	Database *sql.DB
}

func NewUserRepository(db *sql.DB) service.UserRepositoryInterface {
	return &UserRepo{
		Database: db,
	}
}
