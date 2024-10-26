package delivery

import (
	"context"

	"github.com/go-park-mail-ru/2024_2_GOATS/internal/app/models"
)

//go:generate mockgen -source=delivery.go -destination=mocks/mock.go
type UserServiceInterface interface {
	UpdateProfile(ctx context.Context, profileData *models.User) (*models.UpdateUserRespData, *models.ErrorRespData)
	UpdatePassword(ctx context.Context, passwordData *models.PasswordData) (*models.UpdateUserRespData, *models.ErrorRespData)
}
