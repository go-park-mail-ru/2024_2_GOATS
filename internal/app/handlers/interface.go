package handlers

import (
	"context"
	"net/url"

	"github.com/go-park-mail-ru/2024_2_GOATS/internal/app/models"
	authModels "github.com/go-park-mail-ru/2024_2_GOATS/internal/app/models/auth"
)

//go:generate mockgen -source=interface.go -destination=mocks/mock.go
type MovieImplementationInterface interface {
	GetCollection(ctx context.Context, query url.Values) (*models.CollectionsResponse, *models.ErrorResponse)
}

type AuthImplementationInterface interface {
	Register(ctx context.Context, registerData *authModels.RegisterData) (*authModels.AuthResponse, *models.ErrorResponse)
	Login(ctx context.Context, loginData *authModels.LoginData) (*authModels.AuthResponse, *models.ErrorResponse)
	Session(ctx context.Context, cookie string) (*authModels.SessionResponse, *models.ErrorResponse)
	Logout(ctx context.Context, cookie string) (*authModels.AuthResponse, *models.ErrorResponse)
}
