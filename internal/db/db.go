package db

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/go-park-mail-ru/2024_2_GOATS/config"
)

func SetupDatabase(ctx context.Context) {
	ctxVals := config.GetConfigFromContext(ctx)

	connString := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		ctxVals.Postgres.Host,
		ctxVals.Postgres.Port,
		ctxVals.Postgres.User,
		ctxVals.Postgres.Password,
		ctxVals.Postgres.Name,
	)

	DB, err := sql.Open("postgres", connString)
	if err != nil {
		log.Fatal("Opening DB: ", err)
	}

	defer DB.Close()

	log.Printf("Database connection opened successfully")
	time.Sleep(5 * time.Second)

	err = DB.Ping()
	if err != nil {
		log.Fatal("Pinging DB: ", err)
	}

	log.Printf("Database pinged successfully")
}
