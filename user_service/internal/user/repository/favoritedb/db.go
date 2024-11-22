package favoritedb

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/go-park-mail-ru/2024_2_GOATS/user_service/internal/user/repository/dto"
	"github.com/rs/zerolog/log"
)

const (
	favCreateSQL = `
		INSERT INTO favorites (user_id, movie_id)
		VALUES ($1, $2)
	`

	favDestroySQL = `
		DELETE FROM favorites
		WHERE user_id = $1 and movie_id = $2
	`

	favGetSQL = `
		SELECT movie_id FROM favorites
		WHERE user_id = $1
	`

	favCheckSQL = `
		SELECT count(movie_id) FROM favorites
		WHERE user_id = $1 and movie_id = $2
	`
)

func Create(ctx context.Context, favReq *dto.RepoFavorite, db *sql.DB) error {
	logger := log.Ctx(ctx)

	err := db.QueryRowContext(
		ctx,
		favCreateSQL,
		favReq.UserID, favReq.MovieID,
	).Err()

	if err != nil {
		errMsg := fmt.Errorf("postgres: error while creating favorite - %w", err)
		logger.Error().Err(errMsg).Msg("pg_error")

		return errMsg
	}

	logger.Info().Msg("postgres: favorite created successfully")

	return nil
}

func Destroy(ctx context.Context, favReq *dto.RepoFavorite, db *sql.DB) error {
	logger := log.Ctx(ctx)

	err := db.QueryRowContext(
		ctx,
		favDestroySQL,
		favReq.UserID, favReq.MovieID,
	).Err()

	if err != nil {
		errMsg := fmt.Errorf("postgres: error while destroying favorite - %w", err)
		logger.Error().Err(errMsg).Msg("pg_error")

		return errMsg
	}

	logger.Info().Msg("postgres: favorite destroyed successfully")

	return nil
}

func FindByUserID(ctx context.Context, userID uint64, db *sql.DB) (*sql.Rows, error) {
	logger := log.Ctx(ctx)

	rows, err := db.QueryContext(ctx, favGetSQL, userID)
	if err != nil {
		errMsg := fmt.Errorf("postgres: error while scanning favorites by user_id - %w", err)
		logger.Error().Err(errMsg).Msg("pg_error")

		return nil, errMsg
	}

	logger.Info().Msgf("postgres: favorites with id %d found", userID)

	return rows, nil
}

func Check(ctx context.Context, favData *dto.RepoFavorite, db *sql.DB) (bool, error) {
	logger := log.Ctx(ctx)

	rows, err := db.QueryContext(ctx, favCheckSQL, favData)
	if err != nil {
		errMsg := fmt.Errorf("postgres: failed to check favorite existence: %w", err)
		logger.Error().Err(errMsg).Msg("database query error")
		return false, errMsg
	}
	defer rows.Close()

	if rows.Next() {
		logger.Info().Msgf("postgres: favorite pair for user %d and movie found", favData.UserID, favData.MovieID)
		return true, nil
	}

	logger.Info().Msgf("postgres: favorite pair for user %d and movie not found", favData.UserID, favData.MovieID)
	return false, nil
}
