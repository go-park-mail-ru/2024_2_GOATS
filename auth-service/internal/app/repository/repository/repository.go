package repository

import (
	"context"
	"errors"
	pg "github.com/go-park-mail-ru/2024_2_GOATS/internal/pkg/database/postgres"
	"log"
	"time"
)

// IRepository defines the interface for database interaction
type IRepository interface {
	SaveUser(ctx context.Context, email string, passwordHash string, nickname string, sex string, birthdate time.Time) error
	GetUserByEmail(ctx context.Context, email string) (string, error)
}

// PostgresRepository implements the IRepository interface for PostgreSQL
type PostgresRepository struct {
	db pg.Database // Используем интерфейс Database вместо *PGDatabase
}

// NewRepo creates a new instance of PostgresRepository
func NewRepo(db pg.Database) *PostgresRepository {
	return &PostgresRepository{db: db}
}

// SaveUser inserts a new user into the database
func (r *PostgresRepository) SaveUser(ctx context.Context, email string, passwordHash string, nickname string, sex string, birthdate time.Time) error {
	query := `
		INSERT INTO users (email, password_hash, nickname, sex, birthdate)
		VALUES ($1, $2, $3, $4, $5)
	`

	// Выполняем запрос через Exec
	_, err := r.db.Exec(ctx, query, email, passwordHash, nickname, sex, birthdate)
	if err != nil {
		log.Printf("Error saving user: %v", err)
		return errors.New("failed to save user")
	}

	return nil
}

// GetUserByEmail retrieves the password hash for a user by their email
func (r *PostgresRepository) GetUserByEmail(ctx context.Context, email string) (string, error) {
	var passwordHash string
	query := `SELECT password_hash FROM users WHERE email = $1`

	// Используем QueryRow для получения данных
	err := r.db.QueryRow(ctx, query, email).Scan(&passwordHash)
	if err != nil {
		if err.Error() == "no rows in result set" { // Случай, если пользователь не найден
			return "", errors.New("user not found")
		}
		log.Printf("Error retrieving user by email: %v", err)
		return "", errors.New("failed to retrieve user")
	}

	return passwordHash, nil
}
