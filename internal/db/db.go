package db

import (
	"context"
	"database/sql"
	"fmt"
	"os"
	"time"

	"github.com/go-park-mail-ru/2024_2_GOATS/config"
	"github.com/labstack/gommon/log"
	"github.com/spf13/viper"
)

func SetupDatabase(ctx context.Context) (*sql.DB, error) {
	ctxVals := config.FromContext(ctx)
	ctxTimeout, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()

	for {
		select {
		case <-ctxTimeout.Done():
			return nil, ctxTimeout.Err()
		default:
			DB, err := ConnectDB(ctxVals)
			if err == nil {
				return DB, nil
			}

			log.Errorf("Failed to connect to database. Error: %v. Retrying...", err)
			time.Sleep(5 * time.Second)
		}
	}
}

func ConnectDB(cfg *config.Config) (*sql.DB, error) {
	connString := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		cfg.Databases.Postgres.Host,
		cfg.Databases.Postgres.Port,
		cfg.Databases.Postgres.User,
		cfg.Databases.Postgres.Password,
		cfg.Databases.Postgres.Name,
	)

	DB, err := sql.Open("postgres", connString)
	if err != nil {
		return nil, fmt.Errorf("error while opening DB: %w", err)
	}

	log.Printf("Database connection opened successfully")
	time.Sleep(5 * time.Second)

	err = DB.Ping()
	if err != nil {
		return nil, fmt.Errorf("error while pinging DB: %w", err)
	}

	log.Printf("Database pinged successfully")

	if err = migrate(DB); err != nil {
		return nil, fmt.Errorf("error while migrating DB: %w", err)
	}

	if err = seed(DB); err != nil {
		return nil, fmt.Errorf("error while seeding DB: %w", err)
	}

	return DB, nil
}

func migrate(db *sql.DB) error {
	sqlFile, err := os.ReadFile(viper.GetString("SCHEMA_PATH"))
	if err != nil {
		return fmt.Errorf("error read sql script: %w", err)
	}

	_, err = db.Exec(string(sqlFile))
	if err != nil {
		return fmt.Errorf("error while exec sqlFile: %w", err)
	}

	return nil
}

func seed(db *sql.DB) error {
	seedsFile, err := os.ReadFile(viper.GetString("SEEDS_PATH"))

	if err != nil {
		return fmt.Errorf("error read sql script: %w", err)
	}

	_, err = db.Exec(string(seedsFile))
	if err != nil {
		return fmt.Errorf("error while exec seedsFile: %w", err)
	}

	return nil
}
