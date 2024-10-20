package user

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/go-park-mail-ru/2024_2_GOATS/internal/app/models"
	"github.com/go-park-mail-ru/2024_2_GOATS/internal/app/user/repository/password"
	"github.com/rs/zerolog/log"
)

func Create(ctx context.Context, registerData models.RegisterData, db *sql.DB) (*models.User, error) {
	logger := log.Ctx(ctx)
	hashPass, err := password.HashAndSalt(registerData.Password)
	if err != nil {
		return nil, fmt.Errorf("error hashing password: %w", err)
	}

	sqlStatement := `
		INSERT INTO users (email, username, password_hash)
		VALUES ($1, $2, $3)
		RETURNING id, email`

	usr := models.User{}
	err = db.QueryRowContext(
		ctx,
		sqlStatement,
		registerData.Email, registerData.Username, hashPass,
	).Scan(&usr.Id, &usr.Email)

	if err != nil {
		errMsg := fmt.Errorf("postgres: error while creating user - %w", err)
		logger.Error().Msg(errMsg.Error())

		return nil, errMsg
	}

	logger.Info().Msg(fmt.Sprintf("postgres: successfully create user with Email - %s", usr.Email))

	return &usr, nil
}

func FindByEmail(ctx context.Context, email string, db *sql.DB) (*models.User, error) {
	var usr models.User
	logger := log.Ctx(ctx)
	err := db.QueryRowContext(
		ctx,
		"SELECT id, email, username, password_hash FROM USERS WHERE email = $1", email,
	).Scan(&usr.Id, &usr.Email, &usr.Username, &usr.Password)

	if err != nil {
		errMsg := fmt.Errorf("postgres: error while scanning user by email - %w", err)
		logger.Error().Msg(errMsg.Error())

		return nil, errMsg
	}

	logger.Info().Msg(fmt.Sprintf("postgres: user with email %s found", email))

	return &usr, nil
}

func FindById(ctx context.Context, userId int, db *sql.DB) (*models.User, error) {
	var usr models.User
	logger := log.Ctx(ctx)
	err := db.QueryRowContext(
		ctx,
		"SELECT id, email, username, password_hash, sex, birthdate FROM USERS WHERE id = $1", userId,
	).Scan(&usr.Id, &usr.Email, &usr.Username, &usr.Password, &usr.Sex, &usr.Birthdate)

	if err != nil {
		errMsg := fmt.Errorf("postgres: error while scanning user by id - %w", err)
		logger.Error().Msg(errMsg.Error())

		return nil, errMsg
	}

	logger.Info().Msg(fmt.Sprintf("postgres: user wuth id %d found", usr.Id))

	return &usr, nil
}

func UpdatePassword(ctx context.Context, userId int, pass string, db *sql.DB) error {
	logger := log.Ctx(ctx)
	hashPass, err := password.HashAndSalt(pass)
	if err != nil {
		return fmt.Errorf("error hashing password: %w", err)
	}

	sqlStatement := "UPDATE users SET password_hash = $1 WHERE id = $2"

	_, err = db.ExecContext(ctx, sqlStatement, hashPass, userId)

	if err != nil {
		errMsg := fmt.Errorf("postgres: error while updating user password - %w", err)
		logger.Error().Msg(errMsg.Error())

		return errMsg
	}

	logger.Info().Msg(fmt.Sprintf("postgres: successfully update password for user with id - %d", userId))

	return nil
}

func UpdateProfile(ctx context.Context, usrData *models.User, db *sql.DB) error {
	logger := log.Ctx(ctx)

	sqlStatement := "UPDATE users SET email = $1, username = $2, sex = $3, avatar_url = $4 WHERE id = $5"

	_, err := db.ExecContext(ctx, sqlStatement, usrData.Email, usrData.Username, usrData.Sex, usrData.AvatarUrl, usrData.Id)

	if err != nil {
		errMsg := fmt.Errorf("postgres: error while updating user password - %w", err)
		logger.Error().Msg(errMsg.Error())

		return errMsg
	}

	logger.Info().Msg(fmt.Sprintf("postgres: successfully update profile for user with id - %d", usrData.Id))

	return nil
}
