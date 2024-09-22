package service

import (
	"context"

	"github.com/go-park-mail-ru/2024_2_GOATS/internal/app/api"
)

var _ api.ServiceInterface = (*Service)(nil)

type RepositoryInterface interface {
	Login(ctx context.Context)
	Register(ctx context.Context)
	GetCollection(ctx context.Context)
}

type Service struct {
	repository RepositoryInterface
}

func NewService(repo RepositoryInterface) *Service {
	return &Service{
		repository: repo,
	}
}
