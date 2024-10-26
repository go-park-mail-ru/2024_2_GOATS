package actor

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/go-park-mail-ru/2024_2_GOATS/internal/app/models"
	"github.com/rs/zerolog/log"
)

func FindById(ctx context.Context, actorId int, db *sql.DB) (*models.ActorInfo, error) {
	logger := log.Ctx(ctx)
	actorInfo := &models.ActorInfo{}

	actorSqlStatement := `
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

	row := db.QueryRowContext(ctx, actorSqlStatement, actorId)

	logger.Info().Msg("postgres: successfully select actor info")

	err := row.Scan(
		&actorInfo.Id,
		&actorInfo.Name,
		&actorInfo.Surname,
		&actorInfo.Biography,
		&actorInfo.Birthdate,
		&actorInfo.BigPhotoUrl,
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
