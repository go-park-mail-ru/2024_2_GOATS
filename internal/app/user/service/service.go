package service

import (
	"context"
	"os"

	"github.com/go-park-mail-ru/2024_2_GOATS/internal/app/client"
	errVals "github.com/go-park-mail-ru/2024_2_GOATS/internal/app/errors"
	"github.com/go-park-mail-ru/2024_2_GOATS/internal/app/models"
	"github.com/go-park-mail-ru/2024_2_GOATS/internal/app/user/delivery"
	"github.com/go-park-mail-ru/2024_2_GOATS/internal/app/user/repository/dto"
)

var _ delivery.UserServiceInterface = (*UserService)(nil)

//go:generate mockgen -source=service.go -destination=mocks/mock.go
type UserRepositoryInterface interface {
	UserByEmail(ctx context.Context, email string) (*models.User, *errVals.RepoError)
	CreateUser(ctx context.Context, registerData *dto.RepoRegisterData) (*models.User, *errVals.RepoError)
	UserByID(ctx context.Context, userID int) (*models.User, *errVals.RepoError)
	UpdateProfileData(ctx context.Context, usrData *dto.RepoUser) *errVals.RepoError
	UpdatePassword(ctx context.Context, usrID int, pass string) *errVals.RepoError
	SaveUserAvatar(ctx context.Context, avatarName string) (string, *os.File, *errVals.RepoError)
	CreateFavorite(ctx context.Context, favData *dto.RepoFavorite) *errVals.RepoError
	DestroyFavorite(ctx context.Context, favData *dto.RepoFavorite) *errVals.RepoError
	GetFavorites(ctx context.Context, usrID int) ([]models.MovieShortInfo, *errVals.RepoError)
}

type UserService struct {
	userClient client.UserClientInterface
}

func NewUserService(usrClient client.UserClientInterface) delivery.UserServiceInterface {
	return &UserService{
		userClient: usrClient,
	}
}
