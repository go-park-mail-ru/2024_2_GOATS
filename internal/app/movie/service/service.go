package service

import (
	"github.com/go-park-mail-ru/2024_2_GOATS/internal/app/client"
	"github.com/go-park-mail-ru/2024_2_GOATS/internal/app/movie/delivery"
)

//go:generate mockgen -source=service.go -destination=mocks/mock.go

type MovieService struct {
	movieClient client.MovieClientInterface
	userClient  client.UserClientInterface
}

func NewMovieService(mvClient client.MovieClientInterface, usrClient client.UserClientInterface) delivery.MovieServiceInterface {
	return &MovieService{
		movieClient: mvClient,
		userClient:  usrClient,
	}
}
