package service

import (
	"github.com/go-park-mail-ru/2024_2_GOATS/internal/app/client"
	"github.com/go-park-mail-ru/2024_2_GOATS/internal/app/movie/delivery"
)

//go:generate mockgen -source=service.go -destination=mocks/mock.go

// MovieService movie service layer struct
type MovieService struct {
	movieClient client.MovieClientInterface
	userClient  client.UserClientInterface
}

// NewMovieService returns an instance if MovieServiceInterface
func NewMovieService(mvClient client.MovieClientInterface, usrClient client.UserClientInterface) delivery.MovieServiceInterface {
	return &MovieService{
		movieClient: mvClient,
		userClient:  usrClient,
	}
}
