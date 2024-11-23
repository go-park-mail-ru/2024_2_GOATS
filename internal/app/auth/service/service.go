package service

import (
	"context"

	"github.com/go-park-mail-ru/2024_2_GOATS/internal/app/auth/delivery"
	errVals "github.com/go-park-mail-ru/2024_2_GOATS/internal/app/errors"
	"github.com/go-park-mail-ru/2024_2_GOATS/internal/app/models"
	usrServ "github.com/go-park-mail-ru/2024_2_GOATS/internal/app/user/service"
)

var _ delivery.AuthServiceInterface = (*AuthService)(nil)

//go:generate mockgen -source=service.go -destination=mocks/mock.go
type AuthRepositoryInterface interface {
	DestroySession(ctx context.Context, cookie string) *errVals.RepoError
	SetCookie(ctx context.Context, token *models.Token) (*models.CookieData, *errVals.RepoError)
	GetFromCookie(ctx context.Context, cookie string) (string, *errVals.RepoError)
	SetActiveSessionTime(ctx context.Context, usrID string, seconds int) error
}

type AuthService struct {
	authRepository AuthRepositoryInterface
	userRepository usrServ.UserRepositoryInterface
}

func NewAuthService(authRepo AuthRepositoryInterface, usrRepo usrServ.UserRepositoryInterface) delivery.AuthServiceInterface {
	return &AuthService{
		authRepository: authRepo,
		userRepository: usrRepo,
	}
}
