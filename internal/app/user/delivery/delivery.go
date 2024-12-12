package delivery

import (
	"context"

	errVals "github.com/go-park-mail-ru/2024_2_GOATS/internal/app/errors"
	"github.com/go-park-mail-ru/2024_2_GOATS/internal/app/models"
)

// UserServiceInterface defines facade service layer methods
//
//go:generate mockgen -source=delivery.go -destination=mocks/mock.go
type UserServiceInterface interface {
	UpdateProfile(ctx context.Context, profileData *models.User) *errVals.ServiceError
	UpdatePassword(ctx context.Context, passwordData *models.PasswordData) *errVals.ServiceError
	AddFavorite(ctx context.Context, favData *models.Favorite) *errVals.ServiceError
	ResetFavorite(ctx context.Context, favData *models.Favorite) *errVals.ServiceError
	GetFavorites(ctx context.Context, usrID int) ([]models.MovieShortInfo, *errVals.ServiceError)
}
