package delivery

import (
	"context"

	srvDTO "github.com/go-park-mail-ru/2024_2_GOATS/user_service/internal/user/service/dto"
)

// UserServiceInterface contains user_service methods
//
//go:generate mockgen -source=delivery.go -destination=mocks/mock.go
type UserServiceInterface interface {
	Create(ctx context.Context, createData *srvDTO.CreateUserData) (uint64, error)
	SetFavorite(ctx context.Context, favData *srvDTO.Favorite) error
	GetFavorites(ctx context.Context, usrID uint64) ([]uint64, error)
	ResetFavorite(ctx context.Context, favData *srvDTO.Favorite) error
	UpdatePassword(ctx context.Context, passwordData *srvDTO.PasswordData) error
	UpdateProfile(ctx context.Context, usrData *srvDTO.User) error
	FindByID(ctx context.Context, usrID uint64) (*srvDTO.User, error)
	FindByEmail(ctx context.Context, email string) (*srvDTO.User, error)
	CheckFavorite(ctx context.Context, favData *srvDTO.Favorite) (bool, error)
	CreateSubscription(ctx context.Context, createData *srvDTO.CreateSubscriptionData) (uint64, error)
	UpdateSubscribtionStatus(ctx context.Context, subID uint64) error
}
