package service

import (
	"context"

	"github.com/go-park-mail-ru/2024_2_GOATS/internal/app/auth/delivery"
	"github.com/go-park-mail-ru/2024_2_GOATS/internal/app/client"
	errVals "github.com/go-park-mail-ru/2024_2_GOATS/internal/app/errors"
	"github.com/go-park-mail-ru/2024_2_GOATS/internal/app/models"
)

var _ delivery.AuthServiceInterface = (*AuthService)(nil)

//go:generate mockgen -source=service.go -destination=mocks/mock.go
type AuthRepositoryInterface interface {
	DestroySession(ctx context.Context, cookie string) *errVals.RepoError
	SetCookie(ctx context.Context, token *models.Token) (*models.CookieData, *errVals.RepoError)
	GetFromCookie(ctx context.Context, cookie string) (string, *errVals.RepoError)
}

type AuthService struct {
	authClient client.AuthClientInterface
	userClient client.UserClientInterface
}

func NewAuthService(
	authClient client.AuthClientInterface,
	usrClient client.UserClientInterface,
) delivery.AuthServiceInterface {
	return &AuthService{
		authClient: authClient,
		userClient: usrClient,
	}
}
