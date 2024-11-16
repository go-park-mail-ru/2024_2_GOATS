package service

import (
	"context"

	"github.com/go-park-mail-ru/2024_2_GOATS/auth_service/internal/auth/delivery"
	repoDTO "github.com/go-park-mail-ru/2024_2_GOATS/auth_service/internal/auth/repository/dto"
	"github.com/go-park-mail-ru/2024_2_GOATS/auth_service/internal/auth/service/dto"
	"github.com/go-park-mail-ru/2024_2_GOATS/auth_service/internal/errors"
)

type AuthRepositoryInterface interface {
	SetCookie(ctx context.Context, token *repoDTO.TokenData) (*dto.Cookie, *errors.ErrorObj)
	DestroySession(ctx context.Context, cookie string) *errors.ErrorObj
	GetSessionData(ctx context.Context, cookie string) (string, *errors.ErrorObj)
}

type AuthService struct {
	authRepository AuthRepositoryInterface
}

func NewAuthService(authRepository AuthRepositoryInterface) delivery.AuthServiceInterface {
	return &AuthService{
		authRepository: authRepository,
	}
}
