package user

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/go-park-mail-ru/2024_2_GOATS/config"
	"github.com/go-park-mail-ru/2024_2_GOATS/internal/app/models"
)

func Create(ctx context.Context, registerData models.RegisterData, db *sql.DB) (*models.User, error) {
	logger, requestId := config.FromBaseContext(ctx)
	sqlStatement := `
		INSERT INTO users (email, username, password_hash)
		VALUES ($1, $2, $3)
		RETURNING id, email`

	usr := models.User{}
	err := db.QueryRowContext(
		ctx,
		sqlStatement,
		registerData.Email, registerData.Username, registerData.Password,
	).Scan(&usr.Id, &usr.Email)

	if err != nil {
		errMsg := fmt.Errorf("postgres: error while creating user - %w", err)
		logger.LogError(errMsg.Error(), errMsg, requestId)

		return nil, errMsg
	}

	logger.Log("postgres: successfully create user", requestId)

	return &usr, nil
}

func FindByEmail(ctx context.Context, email string, db *sql.DB) (*models.User, error) {
	var usr models.User
	logger, requestId := config.FromBaseContext(ctx)

	err := db.QueryRowContext(
		ctx,
		"SELECT id, email, username, password_hash FROM USERS WHERE email = $1", email,
	).Scan(&usr.Id, &usr.Email, &usr.Username, &usr.Password)

	if err != nil {
		errMsg := fmt.Errorf("postgres: error while scanning user by email - %w", err)
		logger.LogError(errMsg.Error(), errMsg, requestId)

		return nil, errMsg
	}

	logger.Log("postgres: user found by email", requestId)

	return &usr, nil
}

func FindById(ctx context.Context, userId int, db *sql.DB) (*models.User, error) {
	var usr models.User
	logger, requestId := config.FromBaseContext(ctx)

	err := db.QueryRowContext(
		ctx,
		"SELECT id, email, username, password_hash, avatar_url FROM USERS WHERE id = $1", userId,
	).Scan(&usr.Id, &usr.Email, &usr.Username, &usr.Password, &usr.AvatarUrl)

	if err != nil {
		errMsg := fmt.Errorf("postgres: error while scanning user by id - %w", err)
		logger.LogError(errMsg.Error(), errMsg, requestId)

		return nil, errMsg
	}

	logger.Log(fmt.Sprintf("postgres: user with id %d found", usr.Id), requestId)

	return &usr, nil
}

func UpdatePassword(ctx context.Context, userId int, pass string, db *sql.DB) error {
	logger, requestId := config.FromBaseContext(ctx)

	sqlStatement := "UPDATE users SET password_hash = $1, updated_at = $2 WHERE id = $3"

	_, err := db.ExecContext(ctx, sqlStatement, pass, time.Now(), userId)

	if err != nil {
		errMsg := fmt.Errorf("postgres: error while updating user password - %w", err)
		logger.LogError(errMsg.Error(), errMsg, requestId)

		return errMsg
	}

	logger.Log(fmt.Sprintf("postgres: successfully update password for user with id - %d", userId), requestId)

	return nil
}

func UpdateProfile(ctx context.Context, usrData *models.User, db *sql.DB) error {
	logger, requestId := config.FromBaseContext(ctx)

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
	if usrData.AvatarUrl != "" {
		sets = append(sets, fmt.Sprintf("avatar_url = $%d", argCount))
		args = append(args, usrData.AvatarUrl)
		argCount++
	}

	if len(sets) == 0 {
		errMsg := fmt.Errorf("no data to update")
		logger.LogError(errMsg.Error(), errMsg, requestId)

		return errMsg
	}

	sets = append(sets, fmt.Sprintf("updated_at = $%d", argCount))
	args = append(args, time.Now())
	argCount++

	sqlStatement += strings.Join(sets, ", ") + fmt.Sprintf(" WHERE id = $%d", argCount)
	args = append(args, usrData.Id)

	_, err := db.ExecContext(ctx, sqlStatement, args...)
	if err != nil {
		errMsg := fmt.Errorf("postgres: error while updating user profile - %w", err)
		logger.LogError(errMsg.Error(), errMsg, requestId)

		return errMsg
	}

	logger.Log(fmt.Sprintf("postgres: successfully updated profile for user with id - %d", usrData.Id), requestId)

	return nil
}
