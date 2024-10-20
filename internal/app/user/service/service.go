package service

import (
	"context"

	errVals "github.com/go-park-mail-ru/2024_2_GOATS/internal/app/errors"
	"github.com/go-park-mail-ru/2024_2_GOATS/internal/app/models"
	"github.com/go-park-mail-ru/2024_2_GOATS/internal/app/user/delivery"
)

var _ delivery.UserServiceInterface = (*UserService)(nil)

//go:generate mockgen -source=service.go -destination=mocks/mock.go
type UserRepositoryInterface interface {
	UserByEmail(ctx context.Context, email string) (*models.User, *errVals.ErrorObj, int)
	CreateUser(ctx context.Context, registerData *models.RegisterData) (*models.User, *errVals.ErrorObj, int)
	UserById(ctx context.Context, userId int) (*models.User, *errVals.ErrorObj, int)
	UpdateProfileData(ctx context.Context, profileData *models.User) (*errVals.ErrorObj, int)
	UpdatePassword(ctx context.Context, usrId int, pass string) (*errVals.ErrorObj, int)
}

type UserService struct {
	userRepo UserRepositoryInterface
}

func NewUserService(userRepo UserRepositoryInterface) delivery.UserServiceInterface {
	return &UserService{
		userRepo: userRepo,
	}
}
