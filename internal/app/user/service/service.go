package service

import (
	"context"
	"os"

	errVals "github.com/go-park-mail-ru/2024_2_GOATS/internal/app/errors"
	"github.com/go-park-mail-ru/2024_2_GOATS/internal/app/models"
	"github.com/go-park-mail-ru/2024_2_GOATS/internal/app/user/delivery"
	"github.com/go-park-mail-ru/2024_2_GOATS/internal/app/user/repository/dto"
)

var _ delivery.UserServiceInterface = (*UserService)(nil)

//go:generate mockgen -source=service.go -destination=mocks/mock.go
type UserRepositoryInterface interface {
	UserByEmail(ctx context.Context, email string) (*models.User, *errVals.RepoError)
	CreateUser(ctx context.Context, registerData *dto.DBRegisterData) (*models.User, *errVals.RepoError)
	UserByID(ctx context.Context, userID int) (*models.User, *errVals.RepoError)
	UpdateProfileData(ctx context.Context, usrData *dto.DBUser) *errVals.RepoError
	UpdatePassword(ctx context.Context, usrID int, pass string) *errVals.RepoError
	SaveUserAvatar(ctx context.Context, avatarName string) (string, *os.File, *errVals.RepoError)
	CreateFavorite(ctx context.Context, favData *dto.DBFavorite) *errVals.RepoError
	DestroyFavorite(ctx context.Context, favData *dto.DBFavorite) *errVals.RepoError
	GetFavorites(ctx context.Context, usrID int) ([]models.MovieShortInfo, *errVals.RepoError)
}

type UserService struct {
	userRepo UserRepositoryInterface
}

func NewUserService(userRepo UserRepositoryInterface) delivery.UserServiceInterface {
	return &UserService{
		userRepo: userRepo,
	}
}
