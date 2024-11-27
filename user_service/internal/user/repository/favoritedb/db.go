package favoritedb

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/go-park-mail-ru/2024_2_GOATS/user_service/internal/user/repository/dto"
	"github.com/go-park-mail-ru/2024_2_GOATS/user_service/internal/user/repository/metrics_utils"
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

func Create(ctx context.Context, favReq *dto.RepoFavorite, db *sql.DB) error {
	start := time.Now()
	logger := log.Ctx(ctx)

	err := db.QueryRowContext(
		ctx,
		favCreateSQL,
		favReq.UserID, favReq.MovieID,
	).Err()

	if err != nil {
		metricsutils.SaveErrorMetric(start, "create_favorite", "favorites")
		errMsg := fmt.Errorf("postgres: error while creating favorite - %w", err)
		logger.Error().Err(errMsg).Msg("pg_error")

		return errMsg
	}

	metricsutils.SaveSuccessMetric(start, "create_favorite", "favorites")
	logger.Info().Msg("postgres: favorite created successfully")

	return nil
}

func Destroy(ctx context.Context, favReq *dto.RepoFavorite, db *sql.DB) error {
	start := time.Now()
	logger := log.Ctx(ctx)

	err := db.QueryRowContext(
		ctx,
		favDestroySQL,
		favReq.UserID, favReq.MovieID,
	).Err()

	if err != nil {
		metricsutils.SaveErrorMetric(start, "destroy_favorite", "favorites")
		errMsg := fmt.Errorf("postgres: error while destroying favorite - %w", err)
		logger.Error().Err(errMsg).Msg("pg_error")

		return errMsg
	}

	metricsutils.SaveSuccessMetric(start, "destroy_favorite", "favorites")
	logger.Info().Msg("postgres: favorite destroyed successfully")

	return nil
}

func FindByUserID(ctx context.Context, userID uint64, db *sql.DB) (*sql.Rows, error) {
	start := time.Now()
	logger := log.Ctx(ctx)

	rows, err := db.QueryContext(ctx, favGetSQL, userID)
	if err != nil {
		metricsutils.SaveErrorMetric(start, "find_user_favorites", "favorites")
		errMsg := fmt.Errorf("postgres: error while scanning favorites by user_id - %w", err)
		logger.Error().Err(errMsg).Msg("pg_error")

		return nil, errMsg
	}

	metricsutils.SaveSuccessMetric(start, "find_user_favorites", "favorites")
	logger.Info().Msgf("postgres: favorites with id %d found", userID)

	return rows, nil
}

func Check(ctx context.Context, favData *dto.RepoFavorite, db *sql.DB) (bool, error) {
	start := time.Now()
	logger := log.Ctx(ctx)

	rows, err := db.QueryContext(ctx, favCheckSQL, favData.UserID, favData.MovieID)
	if err != nil {
		metricsutils.SaveErrorMetric(start, "check_favorite_existence", "favorites")
		errMsg := fmt.Errorf("postgres: failed to check favorite existence: %w", err)
		logger.Error().Err(errMsg).Msg("database query error")
		return false, errMsg
	}
	defer rows.Close()

	if rows.Next() {
		logger.Info().Msgf("postgres: favorite pair for user %d and movie_service %d found", favData.UserID, favData.MovieID)
		return true, nil
	}

	metricsutils.SaveSuccessMetric(start, "check_favorite_existence", "favorites")
	logger.Info().Msgf("postgres: favorite pair for user %d and movie_service %d not found", favData.UserID, favData.MovieID)
	return false, nil
}
