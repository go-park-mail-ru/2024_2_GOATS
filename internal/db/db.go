package db

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/go-park-mail-ru/2024_2_GOATS/config"
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
			log.Printf("Failed to connect to database. Retrying...")
			time.Sleep(5 * time.Second)
		}
	}
}

func ConnectDB(cfg *config.Config) (*sql.DB, error) {
	connString := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		cfg.Postgres.Host,
		cfg.Postgres.Port,
		cfg.Postgres.User,
		cfg.Postgres.Password,
		cfg.Postgres.Name,
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
	sqlFile, err := os.ReadFile(viper.GetString("SQL_PATH"))
	if err != nil {
		return nil, fmt.Errorf("error read sql script: %w", err)
	}

	_, err = DB.Exec(string(sqlFile))
	if err != nil {
		return nil, fmt.Errorf("error while migrating DB: %w", err)
	}

	return DB, nil
}
