package app

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"sync"
	"time"

	"github.com/rs/zerolog"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/go-redis/redis/v8"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"

	"github.com/go-park-mail-ru/2024_2_GOATS/config"
	authApi "github.com/go-park-mail-ru/2024_2_GOATS/internal/app/auth/delivery"
	authRepo "github.com/go-park-mail-ru/2024_2_GOATS/internal/app/auth/repository"
	authServ "github.com/go-park-mail-ru/2024_2_GOATS/internal/app/auth/service"
	"github.com/go-park-mail-ru/2024_2_GOATS/internal/app/client"
	review "github.com/go-park-mail-ru/2024_2_GOATS/review_service/pkg/review_v1"
	movieApi "github.com/go-park-mail-ru/2024_2_GOATS/internal/app/movie/delivery"
	movieRepo "github.com/go-park-mail-ru/2024_2_GOATS/internal/app/movie/repository"
	movieServ "github.com/go-park-mail-ru/2024_2_GOATS/internal/app/movie/service"
	reviewApi "github.com/go-park-mail-ru/2024_2_GOATS/internal/app/review/delivery"
	reviewServ "github.com/go-park-mail-ru/2024_2_GOATS/internal/app/review/service"
	"github.com/go-park-mail-ru/2024_2_GOATS/internal/app/router"
	userApi "github.com/go-park-mail-ru/2024_2_GOATS/internal/app/user/delivery"
	userRepo "github.com/go-park-mail-ru/2024_2_GOATS/internal/app/user/repository"
	userServ "github.com/go-park-mail-ru/2024_2_GOATS/internal/app/user/service"
	"github.com/go-park-mail-ru/2024_2_GOATS/internal/db"
	"github.com/go-park-mail-ru/2024_2_GOATS/internal/middleware"
)

type App struct {
	Database          *sql.DB
	Redis             *redis.Client
	Server            *http.Server
	Mux               *mux.Router
	Logger            *zerolog.Logger
	Cfg               *config.Config
	AcceptConnections bool
}

func New(isTest bool) (*App, error) {
	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	logger := zerolog.New(os.Stdout).With().Timestamp().Logger()

	cfg, err := config.New(isTest)
	if err != nil {
		return nil, fmt.Errorf("error initialize app cfg: %w", err)
	}

	ctx := config.WrapContext(context.Background(), cfg)
	ctxDBTimeout, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()

	database, err := db.SetupDatabase(ctxDBTimeout, cancel)
	if err != nil {
		return nil, fmt.Errorf("error initialize database: %w", err)
	}

	addr := fmt.Sprintf("%s:%d", cfg.Databases.Redis.Host, cfg.Databases.Redis.Port)
	rdb := redis.NewClient(&redis.Options{Addr: addr})

	return &App{
		Database: database,
		Redis:    rdb,
		Logger:   &logger,
		Cfg:      cfg,
	}, nil
}

func (a *App) Run() {
	ctx := config.WrapContext(context.Background(), a.Cfg)

	repoUser := userRepo.NewUserRepository(a.Database)
	srvUser := userServ.NewUserService(repoUser)
	delUser := userApi.NewUserHandler(ctx, srvUser)

	repoAuth := authRepo.NewAuthRepository(a.Database, a.Redis)
	srvAuth := authServ.NewAuthService(repoAuth, repoUser)
	delAuth := authApi.NewAuthHandler(ctx, srvAuth, srvUser)

	repoMov := movieRepo.NewMovieRepository(a.Database)
	srvMov := movieServ.NewMovieService(repoMov)
	delMov := movieApi.NewMovieHandler(srvMov)

	mx := mux.NewRouter()
	router.UseCommonMiddlewares(mx)
	router.SetupCsrf(mx)
	router.SetupAuth(delAuth, mx)
	router.SetupMovie(delMov, mx)

	authMW := middleware.NewSessionMiddleware(srvAuth)
	router.SetupUser(delUser, authMW, mx)

	reviewGrpcConn, err := grpc.NewClient(
		"review_app:8082",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatalf("cant connect to grpc")
	}

	defer reviewGrpcConn.Close()

	reviewClient := client.NewReviewClient(review.NewReviewClient(reviewGrpcConn))
	reviewSrv := reviewServ.NewReviewService(reviewClient)
	reviewDel := reviewApi.NewReviewHandler(reviewSrv)

	router.SetupReview(reviewDel, authMW, mx)

	srv := &http.Server{
		Addr:         fmt.Sprintf(":%d", a.Cfg.Listener.Port),
		Handler:      mx,
		ReadTimeout:  a.Cfg.Listener.Timeout,
		WriteTimeout: a.Cfg.Listener.Timeout,
		IdleTimeout:  a.Cfg.Listener.IdleTimeout,
	}

	a.Server = srv

	mx.Use(a.AppReadyMiddleware)

	a.Logger.Info().Msgf("Server is listening: %s", a.Server.Addr)

	// Not ready yet
	defer func() {
		if err := a.GracefulShutdown(); err != nil {
			a.Logger.Fatal().Msgf("failed to graceful shutdown: %v", err)
		}
	}()

	a.AcceptConnections = true
	if err := a.Server.ListenAndServe(); err != nil {
		if errors.Is(err, http.ErrServerClosed) {
			a.Logger.Info().Msg("server closed under request")
		} else {
			a.Logger.Info().Msgf("server stopped: %v", err)
		}
	}
}

func (a *App) GracefulShutdown() error {
	a.AcceptConnections = false
	a.Logger.Info().Msg("Starting graceful shutdown")

	var wg sync.WaitGroup
	errChan := make(chan error, 2)

	shutdownFuncs := []func() error{
		a.Database.Close,
		a.Redis.Close,
	}

	wg.Add(len(shutdownFuncs))

	for _, shutdownFunc := range shutdownFuncs {
		go func(shutdownFunc func() error) {
			defer wg.Done()
			if err := shutdownFunc(); err != nil {
				errChan <- err
			}
		}(shutdownFunc)
	}

	wg.Wait()
	close(errChan)

	var errs []error
	for err := range errChan {
		if err != nil {
			errs = append(errs, err)
		}
	}

	if len(errs) > 0 {
		return fmt.Errorf("shutdown errors: %v", errs)
	}

	shutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := a.Server.Shutdown(shutdownCtx); err != nil {
		return fmt.Errorf("failed to shut down HTTP server: %w", err)
	}
	a.Logger.Info().Msg("HTTP shut down")

	select {
	case <-shutdownCtx.Done():
		a.Logger.Info().Msg("Graceful shutdown complete")
	default:
		a.Logger.Info().Msg("Waiting for all goroutines to finish...")
		time.Sleep(500 * time.Millisecond)
	}

	return nil
}
