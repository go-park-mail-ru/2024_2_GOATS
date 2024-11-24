package service

import (
	"github.com/go-park-mail-ru/2024_2_GOATS/internal/app/auth/delivery"
	"github.com/go-park-mail-ru/2024_2_GOATS/internal/app/client"
)

var _ delivery.AuthServiceInterface = (*AuthService)(nil)

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
