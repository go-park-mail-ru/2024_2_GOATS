package repository

import (
	"database/sql"

	"github.com/go-park-mail-ru/2024_2_GOATS/user_service/internal/user/service"
)

// UserRepo is a user_repository struct
type UserRepo struct {
	Database *sql.DB
}

// NewUserRepository returns an instance of UserRepoInterface
func NewUserRepository(db *sql.DB) service.UserRepoInterface {
	return &UserRepo{
		Database: db,
	}
}
