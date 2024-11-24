package favoritedb

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/go-park-mail-ru/2024_2_GOATS/internal/app/user/repository/dto"
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
		SELECT movies.id, movies.title, movies.card_url, movies.album_url, movies.rating, movies.release_date, movies.movie_type, countries.title FROM favorites
		JOIN movies on movies.id = favorites.movie_id
		JOIN countries on countries.id = movies.country_id
		WHERE user_id = $1
	`

	favCheckSQL = `
		SELECT id FROM favorites
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

func FindByUserID(ctx context.Context, userID int, db *sql.DB) (*sql.Rows, error) {
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

func CheckFavoriteExists(ctx context.Context, favData *dto.RepoFavorite, db *sql.DB) (bool, error) {
	logger := log.Ctx(ctx)

	rows, err := db.QueryContext(ctx, favCheckSQL, favData.UserID, favData.MovieID)

	defer func() {
		if err := rows.Close(); err != nil {
			errMsg := fmt.Errorf("cannot close rows while checking favorite existence: %w", err)
			log.Error().Err(errMsg).Msg("pg_error")
		}
	}()

	if err != nil {
		errMsg := fmt.Errorf("postgres: error while checking favorite existence - %w", err)
		logger.Error().Err(errMsg).Msg("pg_error")

		return false, errMsg
	}

	present := rows.Next()
	logger.Info().Msgf("postgres: check favorite completed with status: %v", present)

	return present, nil
}