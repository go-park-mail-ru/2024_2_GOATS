package user

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/go-park-mail-ru/2024_2_GOATS/internal/app/auth/repository/password"
	"github.com/go-park-mail-ru/2024_2_GOATS/internal/app/models"
	authModels "github.com/go-park-mail-ru/2024_2_GOATS/internal/app/models/auth"
)

func Create(ctx context.Context, registerData authModels.RegisterData, db *sql.DB) (*models.User, error) {
	hashPass, err := password.HashAndSalt(registerData.Password)
	if err != nil {
		return nil, fmt.Errorf("error hashing password: %w", err)
	}

	sqlStatement := `
		INSERT INTO users (email, username, password_hash)
		VALUES ($1, $2, $3)
		RETURNING id`

	user := models.User{}
	err = db.QueryRowContext(ctx, sqlStatement, registerData.Email, registerData.Username, hashPass).Scan(&user.Id)
	if err != nil {
		return nil, fmt.Errorf("error while creating user: %w", err)
	}

	return &user, nil
}

func FindByEmail(ctx context.Context, email string, db *sql.DB) (*models.User, error) {
	var user models.User
	err := db.QueryRowContext(
		ctx,
		"SELECT id, email, username, password_hash FROM USERS WHERE email = $1", email,
	).Scan(&user.Id, &user.Email, &user.Username, &user.Password)

	if err != nil {
		return nil, fmt.Errorf("error while scanning user by email: %w", err)
	}

	return &user, nil
}

func FindById(ctx context.Context, userId string, db *sql.DB) (*models.User, error) {
	var user models.User
	err := db.QueryRowContext(
		ctx,
		"SELECT id, email, username, password_hash FROM USERS WHERE id = $1", userId,
	).Scan(&user.Id, &user.Email, &user.Username, &user.Password)

	if err != nil {
		return nil, fmt.Errorf("error while scanning user by id: %w", err)
	}

	return &user, nil
}
