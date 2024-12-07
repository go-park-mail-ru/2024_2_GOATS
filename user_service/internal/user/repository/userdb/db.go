package userdb

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/go-park-mail-ru/2024_2_GOATS/user_service/internal/user/repository/dto"
	metricsutils "github.com/go-park-mail-ru/2024_2_GOATS/user_service/internal/user/repository/metrics_utils"
	"github.com/rs/zerolog/log"
)

const (
	usrCreateSQL = `
		INSERT INTO users (email, username, password_hash)
		VALUES ($1, $2, $3)
		RETURNING id, email
	`

	usrFindByID = `
		SELECT users.id, users.email, users.username, users.password_hash, users.avatar_url
		FROM users
		WHERE users.id = $1
	`

	usrFindByEmail       = "SELECT id, email, username, password_hash FROM users WHERE email = $1"
	usrUpdatePasswordSQL = "UPDATE users SET password_hash = $1, updated_at = $2 WHERE id = $3"
)

func Create(ctx context.Context, registerData dto.RepoCreateData, db *sql.DB) (*dto.RepoUser, error) {
	start := time.Now()
	logger := log.Ctx(ctx)

	usr := dto.RepoUser{}
	stmt, err := db.Prepare(usrCreateSQL)
	if err != nil {
		return nil, fmt.Errorf("prepareStatement#createUser: %w", err)
	}

	defer func() {
		if err := stmt.Close(); err != nil {
			logger.Error().Err(err).Msg("failed_to_close_statement")
		}
	}()

	err = stmt.QueryRowContext(
		ctx,
		registerData.Email, registerData.Username, registerData.Password,
	).Scan(&usr.ID, &usr.Email)

	if err != nil {
		metricsutils.SaveErrorMetric(start, "create_user", "users")
		errMsg := fmt.Errorf("postgres: error while creating user - %w", err)
		logger.Error().Err(errMsg).Msg("pg_error")

		return nil, errMsg
	}

	metricsutils.SaveSuccessMetric(start, "create_user", "users")
	logger.Info().Msg("postgres: user created successfully")

	return &usr, nil
}

func FindByEmail(ctx context.Context, email string, db *sql.DB) (*dto.RepoUser, error) {
	var usr dto.RepoUser
	start := time.Now()
	logger := log.Ctx(ctx)

	err := db.QueryRowContext(ctx, usrFindByEmail, email).Scan(&usr.ID, &usr.Email, &usr.Username, &usr.Password)

	if err != nil {
		metricsutils.SaveErrorMetric(start, "find_user_by_email", "users")
		errMsg := fmt.Errorf("postgres: error while scanning user by email - %w", err)
		logger.Error().Err(errMsg).Msg("pg_error")

		return nil, errMsg
	}

	metricsutils.SaveSuccessMetric(start, "find_user_by_email", "users")
	logger.Info().Msg("postgres: user found by email")

	return &usr, nil
}

func FindByID(ctx context.Context, userID uint64, db *sql.DB) (*dto.RepoUser, error) {
	var usr dto.RepoUser
	start := time.Now()
	logger := log.Ctx(ctx)

	err := db.QueryRowContext(ctx, usrFindByID, userID).Scan(
		&usr.ID,
		&usr.Email,
		&usr.Username,
		&usr.Password,
		&usr.AvatarURL,
	)

	if err != nil {
		metricsutils.SaveErrorMetric(start, "find_user_by_id", "users")
		errMsg := fmt.Errorf("postgres: error while scanning user by id - %w", err)
		logger.Error().Err(errMsg).Msg("pg_error")

		return nil, errMsg
	}

	metricsutils.SaveSuccessMetric(start, "find_user_by_id", "users")
	logger.Info().Msgf("postgres: user with id %d found", usr.ID)

	return &usr, nil
}

func UpdatePassword(ctx context.Context, userID uint64, pass string, db *sql.DB) error {
	start := time.Now()
	logger := log.Ctx(ctx)

	stmt, err := db.Prepare(usrUpdatePasswordSQL)
	if err != nil {
		return fmt.Errorf("prepareStatement#updatePassword: %w", err)
	}

	defer func() {
		if err := stmt.Close(); err != nil {
			logger.Error().Err(err).Msg("failed_to_close_statement")
		}
	}()

	_, err = stmt.ExecContext(ctx, pass, time.Now(), userID)

	if err != nil {
		metricsutils.SaveErrorMetric(start, "update_password", "users")
		errMsg := fmt.Errorf("postgres: error while updating user password - %w", err)
		logger.Error().Err(errMsg).Msg("pg_error")

		return errMsg
	}

	metricsutils.SaveSuccessMetric(start, "update_password", "users")
	logger.Info().Msgf("postgres: successfully update password for user with id - %d", userID)

	return nil
}

func UpdateProfile(ctx context.Context, usrData *dto.RepoUser, db *sql.DB) error {
	start := time.Now()
	logger := log.Ctx(ctx)

	sqlStatement := "UPDATE users SET "
	var sets []string
	var args []interface{}
	argCount := 1

	if usrData.Email != "" {
		sets = append(sets, fmt.Sprintf("email = $%d", argCount))
		args = append(args, usrData.Email)
		argCount++
	}
	if usrData.Username != "" {
		sets = append(sets, fmt.Sprintf("username = $%d", argCount))
		args = append(args, usrData.Username)
		argCount++
	}
	if usrData.AvatarURL != "" {
		sets = append(sets, fmt.Sprintf("avatar_url = $%d", argCount))
		args = append(args, usrData.AvatarURL)
		argCount++
	}

	if len(sets) == 0 {
		logger.Info().Msg("empty_update_data")
		return nil
	}

	sets = append(sets, fmt.Sprintf("updated_at = $%d", argCount))
	args = append(args, time.Now())
	argCount++

	sqlStatement += strings.Join(sets, ", ") + fmt.Sprintf(" WHERE id = $%d", argCount)
	args = append(args, usrData.ID)

	stmt, err := db.Prepare(sqlStatement)
	if err != nil {
		return fmt.Errorf("prepareStatement#updateUser: %w", err)
	}

	defer func() {
		if err := stmt.Close(); err != nil {
			logger.Error().Err(err).Msg("failed_to_close_statement")
		}
	}()

	_, err = stmt.ExecContext(ctx, args...)
	if err != nil {
		metricsutils.SaveErrorMetric(start, "update_profile", "users")
		errMsg := fmt.Errorf("postgres: error while updating user profile - %w", err)
		logger.Error().Err(errMsg).Msg("pg_error")

		return errMsg
	}

	metricsutils.SaveSuccessMetric(start, "update_profile", "users")
	logger.Info().Msgf("postgres: successfully updated profile for user with id - %d", usrData.ID)

	return nil
}
