package app

import (
	"context"
	"database/sql"
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

type App struct {
	Database *sql.DB
	Context  context.Context
}

func New() (*App, error) {
	cfg, err := config.New()
	if err != nil {
		return nil, fmt.Errorf("error initialize app cfg: %w", err)
	}

	ctx, err := config.WrapContext(cfg)
	if err != nil {
		return nil, fmt.Errorf("error wrap app context: %w", err)
	}

	db, err := db.SetupDatabase(ctx)
	if err != nil {
		return nil, fmt.Errorf("error initialize database: %w", err)
	}

	return &App{
		Database: db,
		Context:  ctx,
	}, nil
}

func (a *App) Run() {
	repoLayer := repository.NewRepository(a.Database)
	srvLayer := service.NewService(repoLayer)
	apiLayer := api.NewImplementation(a.Context, srvLayer)
	mux := router.Setup(a.Context, apiLayer)

	ctxValues := config.FromContext(a.Context)

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
