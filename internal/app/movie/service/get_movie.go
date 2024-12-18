package service

import (
	"context"
	"fmt"

	"github.com/go-park-mail-ru/2024_2_GOATS/config"

	errVals "github.com/go-park-mail-ru/2024_2_GOATS/internal/app/errors"
	"github.com/go-park-mail-ru/2024_2_GOATS/internal/app/models"
)

// GetMovie gets movie by id
func (s *MovieService) GetMovie(ctx context.Context, mvID int) (*models.MovieInfo, *errVals.ServiceError) {
	mv, err := s.movieClient.GetMovie(ctx, mvID)

	if err != nil {
		return nil, &errVals.ServiceError{
			Code:  "GetMovie",
			Error: fmt.Errorf("error GetMovie: %w", err),
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
				Code:  "GetMovie",
				Error: fmt.Errorf("error GetMovie: %w", err),
			}
		}

		mv.IsFavorite = isFav

		if mv.WithSubscription {
			usrData, err := s.userClient.FindByID(ctx, uint64(usrID))
			if err != nil {
				return nil, &errVals.ServiceError{
					Code:  "GetUserData",
					Error: fmt.Errorf("error GetUserData: %w", err),
				}
			}

			if !usrData.SubscriptionStatus {
				mv.VideoURL = ""
			}
		}
	}

	if mv.WithSubscription && usrID == 0 {
		mv.VideoURL = ""
	}

	return mv, nil
}
