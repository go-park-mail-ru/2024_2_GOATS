package service

import (
	"context"

	"github.com/go-park-mail-ru/2024_2_GOATS/auth_service/internal/auth/delivery"
	repoDTO "github.com/go-park-mail-ru/2024_2_GOATS/auth_service/internal/auth/repository/dto"
	"github.com/go-park-mail-ru/2024_2_GOATS/auth_service/internal/auth/service/dto"
)

// AuthRepositoryInterface defines methods for auth_service repo layer
//
//go:generate mockgen -source=service.go -destination=mocks/mock.go
type AuthRepositoryInterface interface {
	SetCookie(ctx context.Context, token *repoDTO.TokenData) (*dto.Cookie, error)
	DestroySession(ctx context.Context, cookie string) error
	GetSessionData(ctx context.Context, cookie string) (string, error)
}

// AuthService auth_service service layer struct
type AuthService struct {
	authRepository AuthRepositoryInterface
}

// NewAuthService returns an instance of AuthServiceInterface
func NewAuthService(authRepository AuthRepositoryInterface) delivery.AuthServiceInterface {
	return &AuthService{
		authRepository: authRepository,
	}
}
