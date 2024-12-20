package favoritedb

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/go-park-mail-ru/2024_2_GOATS/user_service/internal/user/repository/dto"
	metricsutils "github.com/go-park-mail-ru/2024_2_GOATS/user_service/internal/user/repository/metrics_utils"
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
		SELECT movie_id FROM favorites
		WHERE user_id = $1 and movie_id = $2
	`
)

// Create creates user favorite
func Create(ctx context.Context, favReq *dto.RepoFavorite, db *sql.DB) error {
	start := time.Now()
	logger := log.Ctx(ctx)

	stmt, err := db.Prepare(favCreateSQL)
	if err != nil {
		return fmt.Errorf("prepareStatement#createFavorite: %w", err)
	}

	defer func() {
		if clErr := stmt.Close(); clErr != nil {
			logger.Error().Err(clErr).Msg("failed_to_close_statement")
		}
	}()

	err = stmt.QueryRowContext(
		ctx,
		favReq.UserID, favReq.MovieID,
	).Err()

	if err != nil {
		metricsutils.SaveErrorMetric("create_favorite", "favorites")
		errMsg := fmt.Errorf("postgres: error while creating favorite - %w", err)
		logger.Error().Err(errMsg).Msg("pg_error")

		return errMsg
	}

	metricsutils.SaveSuccessMetric(start, "create_favorite", "favorites")
	logger.Info().Msg("postgres: favorite created successfully")

	return nil
}

// Destroy destroys user favorite
func Destroy(ctx context.Context, favReq *dto.RepoFavorite, db *sql.DB) error {
	start := time.Now()
	logger := log.Ctx(ctx)

	stmt, err := db.Prepare(favDestroySQL)
	if err != nil {
		return fmt.Errorf("prepareStatement#destroyFavorite: %w", err)
	}

	defer func() {
		if clErr := stmt.Close(); clErr != nil {
			logger.Error().Err(clErr).Msg("failed_to_close_statement")
		}
	}()

	err = stmt.QueryRowContext(
		ctx,
		favReq.UserID, favReq.MovieID,
	).Err()

	if err != nil {
		metricsutils.SaveErrorMetric("destroy_favorite", "favorites")
		errMsg := fmt.Errorf("postgres: error while destroying favorite - %w", err)
		logger.Error().Err(errMsg).Msg("pg_error")

		return errMsg
	}

	metricsutils.SaveSuccessMetric(start, "destroy_favorite", "favorites")
	logger.Info().Msg("postgres: favorite destroyed successfully")

	return nil
}

// FindByUserID find user's favorites
func FindByUserID(ctx context.Context, userID uint64, db *sql.DB) (*sql.Rows, error) {
	start := time.Now()
	logger := log.Ctx(ctx)

	stmt, err := db.Prepare(favGetSQL)
	if err != nil {
		return nil, fmt.Errorf("prepareStatement#favoritesByUserID: %w", err)
	}

	defer func() {
		if clErr := stmt.Close(); clErr != nil {
			logger.Error().Err(clErr).Msg("failed_to_close_statement")
		}
	}()

	rows, err := stmt.QueryContext(ctx, userID)
	if err != nil {
		metricsutils.SaveErrorMetric("find_user_favorites", "favorites")
		errMsg := fmt.Errorf("postgres: error while scanning favorites by user_id - %w", err)
		logger.Error().Err(errMsg).Msg("pg_error")

		return nil, errMsg
	}

	metricsutils.SaveSuccessMetric(start, "find_user_favorites", "favorites")
	logger.Info().Msgf("postgres: favorites with id %d found", userID)

	return rows, nil
}

// Check checks user's favorites
func Check(ctx context.Context, favData *dto.RepoFavorite, db *sql.DB) (bool, error) {
	start := time.Now()
	logger := log.Ctx(ctx)

	stmt, err := db.Prepare(favCheckSQL)
	if err != nil {
		return false, fmt.Errorf("prepareStatement#checkFavorite: %w", err)
	}

	defer func() {
		if clErr := stmt.Close(); clErr != nil {
			logger.Error().Err(clErr).Msg("failed_to_close_statement")
		}
	}()

	rows, err := stmt.QueryContext(ctx, favData.UserID, favData.MovieID)
	if err != nil {
		metricsutils.SaveErrorMetric("check_favorite_existence", "favorites")
		errMsg := fmt.Errorf("postgres: failed to check favorite existence: %w", err)
		logger.Error().Err(errMsg).Msg("database query error")
		return false, errMsg
	}
	defer func() {
		if clErr := rows.Close(); clErr != nil {
			logger.Error().Err(clErr).Msg("cannot close rows")
		}
	}()

	if rows.Next() {
		logger.Info().Msgf("postgres: favorite pair for user %d and movie_service %d found", favData.UserID, favData.MovieID)
		return true, nil
	}

	metricsutils.SaveSuccessMetric(start, "check_favorite_existence", "favorites")
	logger.Info().Msgf("postgres: favorite pair for user %d and movie_service %d not found", favData.UserID, favData.MovieID)
	return false, nil
}
