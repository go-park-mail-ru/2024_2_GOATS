package delivery

import (
	"context"

	errVals "github.com/go-park-mail-ru/2024_2_GOATS/internal/app/errors"
	"github.com/go-park-mail-ru/2024_2_GOATS/internal/app/models"
)

//go:generate mockgen -source=delivery.go -destination=mocks/mock.go
type AuthServiceInterface interface {
	Login(ctx context.Context, loginData *models.LoginData) (*models.AuthRespData, *errVals.ServiceError)
	Register(ctx context.Context, registerData *models.RegisterData) (*models.AuthRespData, *errVals.ServiceError)
	Session(ctx context.Context, cookie string) (*models.SessionRespData, *errVals.ServiceError)
	Logout(ctx context.Context, cookie string) *errVals.ServiceError
}
