package service

import (
	"context"

	"github.com/go-park-mail-ru/2024_2_GOATS/user_service/internal/user/delivery"
	repoDTO "github.com/go-park-mail-ru/2024_2_GOATS/user_service/internal/user/repository/dto"
	"github.com/go-park-mail-ru/2024_2_GOATS/user_service/internal/user/service/dto"
)

//go:generate mockgen -source=service.go -destination=mocks/mock.go
type UserRepoInterface interface {
	CreateUser(ctx context.Context, registerData *repoDTO.RepoCreateData) (*dto.User, error)
	SetFavorite(ctx context.Context, favData *repoDTO.RepoFavorite) error
	ResetFavorite(ctx context.Context, favData *repoDTO.RepoFavorite) error
	GetFavorites(ctx context.Context, usrID uint64) ([]uint64, error)
	SaveUserAvatar(ctx context.Context, avatarName string, file []byte) (string, error)
	UpdatePassword(ctx context.Context, usrID uint64, pass string) error
	UpdateProfileData(ctx context.Context, profileData *repoDTO.RepoUser) error
	UserByEmail(ctx context.Context, email string) (*dto.User, error)
	UserByID(ctx context.Context, userID uint64) (*dto.User, error)
	CheckFavorite(ctx context.Context, favData *repoDTO.RepoFavorite) (bool, error)
}

type UserService struct {
	userRepo UserRepoInterface
}

func NewUserService(userRepo UserRepoInterface) delivery.UserServiceInterface {
	return &UserService{
		userRepo: userRepo,
	}
}
