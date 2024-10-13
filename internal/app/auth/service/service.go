package service

import (
	"context"

	api "github.com/go-park-mail-ru/2024_2_GOATS/internal/app/auth/delivery"
	errVals "github.com/go-park-mail-ru/2024_2_GOATS/internal/app/errors"
	"github.com/go-park-mail-ru/2024_2_GOATS/internal/app/models"
	authModels "github.com/go-park-mail-ru/2024_2_GOATS/internal/app/models/auth"
)

var _ api.AuthServiceInterface = (*AuthService)(nil)

//go:generate mockgen -source=service.go -destination=mocks/mock.go
type AuthRepositoryInterface interface {
	UserByEmail(ctx context.Context, loginData *authModels.LoginData) (*models.User, *errVals.ErrorObj, int)
	CreateUser(ctx context.Context, registerData *authModels.RegisterData) (*models.User, *errVals.ErrorObj, int)
	UserById(ctx context.Context, userId string) (*models.User, *errVals.ErrorObj, int)
	DestroySession(ctx context.Context, cookie string) (*errVals.ErrorObj, int)
	SetCookie(ctx context.Context, token *authModels.Token) (*authModels.CookieData, *errVals.ErrorObj, int)
	GetFromCookie(ctx context.Context, cookie string) (string, *errVals.ErrorObj, int)
}

type AuthService struct {
	authRepository AuthRepositoryInterface
}

func NewService(authRepo AuthRepositoryInterface) api.AuthServiceInterface {
	return &AuthService{
		authRepository: authRepo,
	}
}
