package service

import (
	"context"

	auth "github.com/go-park-mail-ru/2024_2_GOATS/auth_service/pkg/auth_v1"
	"github.com/go-park-mail-ru/2024_2_GOATS/internal/app/auth/delivery"
	errVals "github.com/go-park-mail-ru/2024_2_GOATS/internal/app/errors"
	"github.com/go-park-mail-ru/2024_2_GOATS/internal/app/models"
	usrServ "github.com/go-park-mail-ru/2024_2_GOATS/internal/app/user/service"
	user "github.com/go-park-mail-ru/2024_2_GOATS/user_service/pkg/user_v1"
)

var _ delivery.AuthServiceInterface = (*AuthService)(nil)

//go:generate mockgen -source=service.go -destination=mocks/mock.go
type AuthRepositoryInterface interface {
	DestroySession(ctx context.Context, cookie string) *errVals.RepoError
	SetCookie(ctx context.Context, token *models.Token) (*models.CookieData, *errVals.RepoError)
	GetFromCookie(ctx context.Context, cookie string) (string, *errVals.RepoError)
}

type AuthService struct {
	authMS         auth.SessionRPCClient
	userMS         user.UserRPCClient
	authRepository AuthRepositoryInterface
	userRepository usrServ.UserRepositoryInterface
}

func NewAuthService(authRepo AuthRepositoryInterface, usrRepo usrServ.UserRepositoryInterface, authMS auth.SessionRPCClient, usrMS user.UserRPCClient) delivery.AuthServiceInterface {
	return &AuthService{
		authMS:         authMS,
		userMS:         usrMS,
		authRepository: authRepo,
		userRepository: usrRepo,
	}
}
