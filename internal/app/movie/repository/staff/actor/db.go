package actor

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/go-park-mail-ru/2024_2_GOATS/internal/app/models"
	"github.com/go-park-mail-ru/2024_2_GOATS/internal/app/movie/repository/staff"
	"github.com/rs/zerolog/log"
)

func FindById(ctx context.Context, staffId int, post string, db *sql.DB) (*models.StaffInfo, error) {
	logger := log.Ctx(ctx)
	actorInfo := &models.StaffInfo{}

	row := staff.FindById(ctx, staffId, post, db)

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
