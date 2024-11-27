package db

import (
	"context"
	"database/sql"
	"fmt"
	"os"
	"time"

	"github.com/go-park-mail-ru/2024_2_GOATS/config"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
)

func SetupDatabase(ctx context.Context, cancel context.CancelFunc) (*sql.DB, error) {
	ctxVals := config.FromContext(ctx)
	defer cancel()

	for {
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		default:
			DB, err := ConnectDB(ctxVals)
			if err == nil {
				return DB, nil
			}

			log.Error().Err(fmt.Errorf("failed to connect to database. Error: %v Retrying", err)).Msg("setup_db_error")
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
		errMsg := fmt.Errorf("error while opening DB: %w", err)
		log.Error().Err(errMsg).Msg("connect_db_error")

		return nil, errMsg
	}

	log.Info().Msg("Database connection opened successfully")
	time.Sleep(5 * time.Second)

	err = DB.Ping()
	if err != nil {
		errMsg := fmt.Errorf("error while pinging DB: %w", err)
		log.Error().Err(errMsg).Msg("ping_db_error")

		return nil, errMsg
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
		errMsg := fmt.Errorf("migration: error read sql script - %w", err)
		log.Error().Err(errMsg).Msg("migrate_db_error")

		return errMsg
	}

	_, err = db.Exec(string(sqlFile))
	if err != nil {
		errMsg := fmt.Errorf("migration: error while exec sqlFile: %w", err)
		log.Error().Err(errMsg).Msg("migrate_db_error")

		return errMsg
	}

	log.Info().Msg("database successfully migrated")
	return nil
}

func seed(db *sql.DB) error {
	seedsFile, err := os.ReadFile(viper.GetString("SEEDS_PATH"))

	if err != nil {
		errMsg := fmt.Errorf("seed: error read sql script - %w", err)
		log.Error().Err(errMsg).Msg("seed_db_error")

		return errMsg
	}

	_, err = db.Exec(string(seedsFile))
	if err != nil {
		errMsg := fmt.Errorf("seed: error while exec seedsFile - %w", err)
		log.Error().Err(errMsg).Msg("seed_db_error")

		return errMsg
	}

	log.Info().Msg("database successfully seeded")
	return nil
}
