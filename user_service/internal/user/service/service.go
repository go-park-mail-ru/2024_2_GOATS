package service

import "github.com/go-park-mail-ru/2024_2_GOATS/user_service/internal/user/delivery"

type UserRepoInterface interface {
}

type UserService struct {
	userRepo UserRepoInterface
}

func NewUserService(userRepo UserRepoInterface) delivery.UserServiceInterface {
	return &UserService{
		userRepo: userRepo,
	}
}
