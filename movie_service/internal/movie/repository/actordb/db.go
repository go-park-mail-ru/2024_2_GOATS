package actordb

import (
	"context"
	"database/sql"
	"fmt"

	"time"

	"github.com/go-park-mail-ru/2024_2_GOATS/movie_service/internal/movie/repository/dto"
	"github.com/rs/zerolog/log"

	"github.com/go-park-mail-ru/2024_2_GOATS/movie_service/internal/movie/repository/metrics_utils"
)

const (
	actorFindByIDSQL = `
		SELECT
			actors.id,
			actors.first_name,
			actors.second_name,
			actors.biography,
			actors.birthdate,
			actors.big_photo_url,
			countries.title
		FROM actors
		JOIN countries on countries.id = actors.country_id
		WHERE actors.id = $1
	`
)

func FindByID(ctx context.Context, actorID int, db *sql.DB) (*dto.RepoActor, error) {
	start := time.Now()
	logger := log.Ctx(ctx)
	actorInfo := &dto.RepoActor{}

	row := db.QueryRowContext(ctx, actorFindByIDSQL, actorID)
	logger.Info().Msg("postgres: successfully select actor info")

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
		metricsutils.SaveErrorMetric(start, "get_actor_by_id", "actors")
		errMsg := fmt.Errorf("postgres: error while selecting actor info: %w", err)
		logger.Error().Err(errMsg).Msg("pg_error")

		return nil, errMsg
	}

	metricsutils.SaveSuccessMetric(start, "get_actor_by_id", "actors")
	logger.Info().Msg("postgres: successfully select actor info")

	return actorInfo, nil
}
