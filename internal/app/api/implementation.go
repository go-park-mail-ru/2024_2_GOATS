package api

import (
	"context"

	"github.com/go-park-mail-ru/2024_2_GOATS/internal/app/models"
	authModels "github.com/go-park-mail-ru/2024_2_GOATS/internal/app/models/auth"
)

type ServiceInterface interface {
	Login(ctx context.Context, loginData *authModels.LoginData) (*authModels.AuthResponse, *models.ErrorResponse)
	Register(ctx context.Context, registerData *authModels.RegisterData) (*authModels.AuthResponse, *models.ErrorResponse)
	GetCollection(ctx context.Context) (*models.CollectionsResponse, *models.ErrorResponse)
	Session(ctx context.Context, cookie string) (*authModels.SessionResponse, *models.ErrorResponse)
}

type Implementation struct {
	ctx     context.Context
	service ServiceInterface
}

func NewImplementation(ctx context.Context, srv ServiceInterface) *Implementation {
	return &Implementation{
		ctx:     ctx,
		service: srv,
	}
}
