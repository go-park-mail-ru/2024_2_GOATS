package delivery

import (
	"context"

	"github.com/go-park-mail-ru/2024_2_GOATS/auth_service/internal/auth/service/dto"
)

// AuthServiceInterface defines methods for auth_service service layer
//
//go:generate mockgen -source=delivery.go -destination=mocks/mock.go
type AuthServiceInterface interface {
	CreateSession(ctx context.Context, srvCreateCookieReq *dto.SrvCreateCookie) (*dto.Cookie, error)
	DestroySession(ctx context.Context, cookie string) (bool, error)
	GetSessionData(ctx context.Context, cookie string) (uint64, error)
}
