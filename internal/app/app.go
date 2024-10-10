package app

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"

	"github.com/go-park-mail-ru/2024_2_GOATS/config"
	"github.com/go-park-mail-ru/2024_2_GOATS/internal/app/api"
	"github.com/go-park-mail-ru/2024_2_GOATS/internal/app/repository"
	"github.com/go-park-mail-ru/2024_2_GOATS/internal/app/router"
	"github.com/go-park-mail-ru/2024_2_GOATS/internal/app/service"
	"github.com/go-park-mail-ru/2024_2_GOATS/internal/db"
)

type App struct {
	Database          *sql.DB
	Redis             *redis.Client
	Context           context.Context
	Server            *http.Server
	Mux               *mux.Router
	AcceptConnections bool
}

func New() (*App, error) {
	cfg, err := config.New()
	if err != nil {
		return nil, fmt.Errorf("error initialize app cfg: %w", err)
	}

	ctx := config.WrapContext(context.Background(), cfg)
	ctxDBTimeout, cancel := context.WithTimeout(ctx, 30*time.Second)
	database, err := db.SetupDatabase(ctxDBTimeout, cancel)
	if err != nil {
		return nil, fmt.Errorf("error initialize database: %w", err)
	}

	addr := fmt.Sprintf("%s:%d", cfg.Databases.Redis.Host, cfg.Databases.Redis.Port)
	rdb := redis.NewClient(&redis.Options{Addr: addr})

	repoLayer := repository.NewRepository(database, rdb)
	srvLayer := service.NewService(repoLayer)
	apiLayer := api.NewImplementation(ctx, srvLayer)
	appMx := router.Setup(ctx, apiLayer)

	srv := &http.Server{
		Addr:         fmt.Sprintf(":%d", cfg.Listener.Port),
		Handler:      appMx,
		ReadTimeout:  cfg.Listener.Timeout,
		WriteTimeout: cfg.Listener.Timeout,
		IdleTimeout:  cfg.Listener.IdleTimeout,
	}

	return &App{
		Database: database,
		Redis:    rdb,
		Context:  ctx,
		Server:   srv,
		Mux:      appMx,
	}, nil
}

func (a *App) Run() {
	ctxValues := config.FromContext(a.Context)
	a.Mux.Use(a.AppReadyMiddleware)

	log.Printf("Server is listening: %s:%d", ctxValues.Listener.Address, ctxValues.Listener.Port)

	// Not ready yet
	defer a.GracefulShutdown()

	a.AcceptConnections = true
	if err := a.Server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("server stopped: %v", err)
	}
}

func (a *App) GracefulShutdown() error {
	a.AcceptConnections = false
	log.Println("Starting graceful shutdown")

	if err := a.Database.Close(); err != nil {
		return fmt.Errorf("failed to close database: %w", err)
	}
	log.Println("Postgres shut down")

	shutdownCtx, cancel := context.WithTimeout(a.Context, 10*time.Second)
	defer cancel()

	if err := a.Server.Shutdown(shutdownCtx); err != nil {
		return fmt.Errorf("failed to shut down HTTP server: %w", err)
	}
	log.Println("HTTP server shut down")

	select {
	case <-shutdownCtx.Done():
		log.Println("Graceful shutdown complete")
	default:
		log.Println("Waiting for all goroutines to finish...")
		time.Sleep(500 * time.Millisecond)
	}

	os.Exit(0)

	return nil
}
