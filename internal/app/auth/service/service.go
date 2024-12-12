package service

import (
	"github.com/go-park-mail-ru/2024_2_GOATS/internal/app/auth/delivery"
	"github.com/go-park-mail-ru/2024_2_GOATS/internal/app/client"
)

// AuthService facade auth service struct
type AuthService struct {
	authClient client.AuthClientInterface
	userClient client.UserClientInterface
}

// NewAuthService returns an instance of AuthServiceInterface
func NewAuthService(
	authClient client.AuthClientInterface,
	usrClient client.UserClientInterface,
) delivery.AuthServiceInterface {
	return &AuthService{
		authClient: authClient,
		userClient: usrClient,
	}
}
