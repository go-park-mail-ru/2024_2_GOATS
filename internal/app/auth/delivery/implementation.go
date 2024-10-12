package delivery

import (
	"context"

	"github.com/go-park-mail-ru/2024_2_GOATS/internal/app/models"
	"github.com/go-park-mail-ru/2024_2_GOATS/internal/app/models/auth"
	authModels "github.com/go-park-mail-ru/2024_2_GOATS/internal/app/models/auth"
)

type AuthServiceInterface interface {
	Login(ctx context.Context, loginData *authModels.LoginData) (*authModels.AuthResponse, *models.ErrorResponse)
	Register(ctx context.Context, registerData *authModels.RegisterData) (*authModels.AuthResponse, *models.ErrorResponse)
	Session(ctx context.Context, cookie string) (*authModels.SessionResponse, *models.ErrorResponse)
	Logout(ctx context.Context, cookie string) (*auth.AuthResponse, *models.ErrorResponse)
}

type Implementation struct {
	Ctx     context.Context
	authService AuthServiceInterface
}

func NewImplementation(ctx context.Context, srv AuthServiceInterface) *Implementation {
	return &Implementation{
		Ctx:     ctx,
		authService: srv,
	}
}
