package app

import (
	"context"
	"fmt"
	"log"
	"net/http"

	_ "github.com/lib/pq"

	"github.com/go-park-mail-ru/2024_2_GOATS/config"
	"github.com/go-park-mail-ru/2024_2_GOATS/internal/app/api"
	"github.com/go-park-mail-ru/2024_2_GOATS/internal/app/repository"
	"github.com/go-park-mail-ru/2024_2_GOATS/internal/app/router"
	"github.com/go-park-mail-ru/2024_2_GOATS/internal/app/service"
	"github.com/go-park-mail-ru/2024_2_GOATS/internal/db"
)

type App struct{}

func New() (*App, context.Context, error) {
	ctx, err := config.CreateConfigContext()

	if err != nil {
		return nil, nil, err
	}

	return &App{}, ctx, nil
}

func (a *App) Run(ctx context.Context) {
	repoLayer := repository.NewRepository()
	srvLayer := service.NewService(repoLayer)
	apiLayer := api.NewImplementation(ctx, srvLayer)
	router := router.SetupRouter(ctx, apiLayer)
	db.SetupDatabase(ctx)

	ctxValues := config.GetConfigFromContext(ctx)

	srv := &http.Server{
		Addr:         fmt.Sprintf(":%d", ctxValues.Listener.Port),
		Handler:      router,
		ReadTimeout:  ctxValues.HttpServer.Timeout,
		WriteTimeout: ctxValues.HttpServer.Timeout,
		IdleTimeout:  ctxValues.HttpServer.IdleTimeout,
	}

	log.Printf("Server is listening: %s:%d", ctxValues.Listener.Address, ctxValues.Listener.Port)

	if err := srv.ListenAndServe(); err != nil {
		log.Fatalf(err.Error())
	}
}
