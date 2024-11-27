package actor

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/go-park-mail-ru/2024_2_GOATS/movie_service/internal/movie/models"
	"github.com/go-park-mail-ru/2024_2_GOATS/movie_service/internal/movie/repository/staff"
	"github.com/rs/zerolog/log"
)

func FindById(ctx context.Context, staffId int, post string, db *sql.DB) (*models.ActorInfo, error) {
	logger := log.Ctx(ctx)
	actorInfo := &models.ActorInfo{}

	row := staff.FindById(ctx, staffId, post, db)

	err := row.Scan(
		&actorInfo.ID,
		&actorInfo.Name,
		&actorInfo.Surname,
		&actorInfo.Biography,
		&actorInfo.Birthdate,
		&actorInfo.BigPhotoURL,
		&actorInfo.Country,
	)

	if err != nil {
		errMsg := fmt.Errorf("postgres: error while selecting actor info: %w", err)
		logger.Err(errMsg)

		return nil, errMsg
	}

	logger.Info().Msg("postgres: successfully select actor info")

	return actorInfo, nil
}
