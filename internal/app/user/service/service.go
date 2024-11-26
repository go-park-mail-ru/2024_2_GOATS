package service

import (
	"github.com/go-park-mail-ru/2024_2_GOATS/internal/app/client"
	"github.com/go-park-mail-ru/2024_2_GOATS/internal/app/user/delivery"
)

var _ delivery.UserServiceInterface = (*UserService)(nil)

type UserService struct {
	userClient client.UserClientInterface
}

func NewUserService(usrClient client.UserClientInterface) delivery.UserServiceInterface {
	return &UserService{
		userClient: usrClient,
	}
}
