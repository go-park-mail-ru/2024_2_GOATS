package db

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/go-park-mail-ru/2024_2_GOATS/movie_service/config"
	"github.com/golang-migrate/migrate"
	"github.com/golang-migrate/migrate/database/postgres"

	// migration driver
	_ "github.com/golang-migrate/migrate/source/file"
	"github.com/rs/zerolog/log"
)

// SetupDatabase connects to Postgres and returns instance of sql.DB
func SetupDatabase(ctx context.Context, cancel context.CancelFunc) (*sql.DB, error) {
	ctxVals := config.FromContext(ctx)
	defer cancel()

	for {
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		default:
			DB, err := connectDB(ctxVals)
			if err == nil {
				return DB, nil
			}

			log.Error().Err(fmt.Errorf("failed to connect to database. Error: %v Retrying", err)).Msg("setup_db_error")
			time.Sleep(5 * time.Second)
		}
	}
}

func connectDB(cfg *config.Config) (*sql.DB, error) {
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

	DB.SetMaxOpenConns(cfg.Databases.Postgres.MaxOpenConns)
	DB.SetMaxIdleConns(cfg.Databases.Postgres.MaxIdleConns)
	DB.SetConnMaxLifetime(time.Duration(cfg.Databases.Postgres.ConnMaxLifetime) * time.Minute)
	DB.SetConnMaxIdleTime(time.Duration(cfg.Databases.Postgres.ConnMaxIdleTime) * time.Minute)

	log.Info().Msg("Database connection opened successfully")
	time.Sleep(5 * time.Second)

	err = DB.Ping()
	if err != nil {
		errMsg := fmt.Errorf("error while pinging DB: %w", err)
		log.Error().Err(errMsg).Msg("ping_db_error")

		return nil, errMsg
	}

	log.Info().Msg("Database pinged successfully")

	if err = migDB(DB); err != nil {
		return nil, fmt.Errorf("error while migrating DB: %w", err)
	}

	return DB, nil
}

func migDB(db *sql.DB) error {
	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		errMsg := fmt.Errorf("migDB: error get sql driver - %w", err)
		log.Error().Msg(errMsg.Error())

		return errMsg
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file://internal/db/migrations/",
		"postgres",
		driver,
	)

	if err != nil {
		errMsg := fmt.Errorf("migDB: cannot create migrator - %w", err)
		log.Error().Msg(errMsg.Error())

		return errMsg
	}

	if err = m.Up(); err != nil {
		log.Error().Err(err).Msg("failed_to_migrate_db")
	}

	return nil
}
