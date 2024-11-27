package service

import (
	"github.com/go-park-mail-ru/2024_2_GOATS/internal/app/client"
	"github.com/go-park-mail-ru/2024_2_GOATS/internal/app/user/delivery"
)

type UserService struct {
	userClient client.UserClientInterface
	mvClient client.MovieClientInterface
}

func NewUserService(usrClient client.UserClientInterface, mvClient client.MovieClientInterface) delivery.UserServiceInterface {
	return &UserService{
		userClient: usrClient,
		mvClient: mvClient,
	}
}
