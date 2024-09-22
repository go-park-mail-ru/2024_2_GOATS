package app

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net/http"

	_ "github.com/lib/pq"
	"github.com/spf13/viper"

	"github.com/go-park-mail-ru/2024_2_GOATS/config"
	"github.com/go-park-mail-ru/2024_2_GOATS/internal/app/api"
	"github.com/go-park-mail-ru/2024_2_GOATS/internal/app/repository"
	"github.com/go-park-mail-ru/2024_2_GOATS/internal/app/router"
	"github.com/go-park-mail-ru/2024_2_GOATS/internal/app/service"
	"github.com/go-park-mail-ru/2024_2_GOATS/internal/db"
)

type App struct {
	Database *sql.DB
}

func New() (*App, context.Context, error) {
	setupViper()
	ctx, err := config.WrapContext(config.NewConfig())
	if err != nil {
		return nil, nil, err
	}

	db, err := db.SetupDatabase(ctx)
	if err != nil {
		return nil, nil, err
	}

	return &App{Database: db}, ctx, nil
}

func (a *App) Run(ctx context.Context) {
	repoLayer := repository.NewRepository(a.Database)
	srvLayer := service.NewService(repoLayer)
	apiLayer := api.NewImplementation(ctx, srvLayer)
	mux := router.Setup(ctx, apiLayer)

	ctxValues := config.FromContext(ctx)

	srv := &http.Server{
		Addr:         fmt.Sprintf(":%d", ctxValues.Listener.Port),
		Handler:      mux,
		ReadTimeout:  ctxValues.Listener.Timeout,
		WriteTimeout: ctxValues.Listener.Timeout,
		IdleTimeout:  ctxValues.Listener.IdleTimeout,
	}

	log.Printf("Server is listening: %s:%d", ctxValues.Listener.Address, ctxValues.Listener.Port)

	if err := srv.ListenAndServe(); err != nil {
		log.Fatalf(err.Error())
	}
}

func setupViper() {
	viper.SetConfigName(".env")
	viper.SetConfigType("env")
	viper.AddConfigPath(".")

	err := viper.ReadInConfig()
	if err != nil {
		fmt.Printf("Failed to read .env file: %v\n", err)
		return
	}

	viper.SetConfigName("config")
	viper.SetConfigType("yml")
	viper.AddConfigPath(viper.GetString("VIPER_CFG_PATH"))

	err = viper.MergeInConfig()
	if err != nil {
		fmt.Printf("Failed to read config.yml file: %v\n", err)
		return
	}
}
