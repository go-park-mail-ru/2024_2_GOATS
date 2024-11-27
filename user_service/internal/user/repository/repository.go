package repository

import (
	"database/sql"

	"github.com/go-park-mail-ru/2024_2_GOATS/user_service/internal/user/service"
)

type UserRepo struct {
	Database *sql.DB
}

func NewUserRepository(db *sql.DB) service.UserRepoInterface {
	return &UserRepo{
		Database: db,
	}
}
