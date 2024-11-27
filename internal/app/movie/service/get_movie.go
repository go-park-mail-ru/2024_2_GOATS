package service

import (
	"context"
	"fmt"

	"github.com/go-park-mail-ru/2024_2_GOATS/config"

	errVals "github.com/go-park-mail-ru/2024_2_GOATS/internal/app/errors"
	"github.com/go-park-mail-ru/2024_2_GOATS/internal/app/models"
)

type RepoFavorite struct {
	UserID  uint64
	MovieID uint64
}

func (s *MovieService) GetMovie(ctx context.Context, mvID int) (*models.MovieInfo, *errVals.ServiceError) {
	mv, err := s.movieClient.GetMovie(ctx, mvID)

	if err != nil {
		return nil, &errVals.ServiceError{
			Code:  "GetCollection",
			Error: fmt.Errorf("error GetCollection: %w", err),
		}
	}

	usrID := config.CurrentUserID(ctx)
	if usrID != 0 {
		fav := &models.Favorite{
			UserID:  usrID,
			MovieID: mv.ID,
		}

		isFav, err := s.userClient.CheckFavorite(ctx, fav)
		if err != nil {
			return nil, &errVals.ServiceError{
				Code:  "GetCollection",
				Error: fmt.Errorf("error GetCollection: %w", err),
			}
		}

		mv.IsFavorite = isFav
	}

	return mv, nil
}
