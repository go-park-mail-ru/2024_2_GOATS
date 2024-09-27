package app

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

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
	Server   *http.Server
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

	repoLayer := repository.NewRepository(db)
	srvLayer := service.NewService(repoLayer)
	apiLayer := api.NewImplementation(ctx, srvLayer)
	mux := router.Setup(ctx, apiLayer)

	srv := &http.Server{
		Addr:         fmt.Sprintf(":%d", cfg.Listener.Port),
		Handler:      mux,
		ReadTimeout:  cfg.Listener.Timeout,
		WriteTimeout: cfg.Listener.Timeout,
		IdleTimeout:  cfg.Listener.IdleTimeout,
	}

	return &App{
		Database: db,
		Context:  ctx,
		Server:   srv,
	}, nil
}

func (a *App) Run() {
	ctxValues := config.FromContext(a.Context)
	log.Printf("Server is listening: %s:%d", ctxValues.Listener.Address, ctxValues.Listener.Port)

	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		for {
			sig := <-c
			log.Println("Received shutdown signal:", sig)

			if err := a.GracefulShutdown(); err != nil {
				log.Fatalf("Failed to shut down gracefully: %v", err)
			}

			return
		}
	}()

	if err := a.Server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf(err.Error())
	}
}

func (a *App) GracefulShutdown() error {
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

	os.Exit(1)

	return nil
}