package delivery

import (
	"context"

	"github.com/go-park-mail-ru/2024_2_GOATS/auth_service/internal/auth/service/dto"
	"github.com/go-park-mail-ru/2024_2_GOATS/auth_service/internal/errors"
)

//go:generate mockgen -source=delivery.go -destination=mocks/mock.go
type AuthServiceInterface interface {
	CreateSession(ctx context.Context, srvCreateCookieReq *dto.SrvCreateCookie) (*dto.Cookie, *errors.SrvErrorObj)
	DestroySession(ctx context.Context, cookie string) (bool, *errors.SrvErrorObj)
	GetSessionData(ctx context.Context, cookie string) (uint64, *errors.SrvErrorObj)
}