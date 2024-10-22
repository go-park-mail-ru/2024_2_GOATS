package delivery

import (
	"context"

	"github.com/go-park-mail-ru/2024_2_GOATS/internal/app/models"
)

//go:generate mockgen -source=delivery.go -destination=mocks/mock.go
type AuthServiceInterface interface {
	Login(ctx context.Context, loginData *models.LoginData) (*models.AuthRespData, *models.ErrorRespData)
	Register(ctx context.Context, registerData *models.RegisterData) (*models.AuthRespData, *models.ErrorRespData)
	Session(ctx context.Context, cookie string) (*models.SessionRespData, *models.ErrorRespData)
	Logout(ctx context.Context, cookie string) (*models.AuthRespData, *models.ErrorRespData)
}
