package repository

import (
	"database/sql"

	"github.com/go-park-mail-ru/2024_2_GOATS/internal/app/service"
)

var _ service.RepositoryInterface = (*Repo)(nil)

type Repo struct {
	Database *sql.DB
}

func NewRepository(db *sql.DB) *Repo {
	return &Repo{
		Database: db,
	}
}
