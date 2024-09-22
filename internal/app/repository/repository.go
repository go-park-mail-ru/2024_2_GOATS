package repository

import (
	"github.com/go-park-mail-ru/2024_2_GOATS/internal/app/service"
)

var _ service.RepositoryInterface = (*Repo)(nil)
type Repo struct {
}

func NewRepository() *Repo {
	return &Repo{}
}
