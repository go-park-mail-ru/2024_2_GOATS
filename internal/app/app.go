package app

import (
	"context"
	"crypto/tls"
	"database/sql"
	"errors"
	"fmt"

	"github.com/elastic/go-elasticsearch/v7"

	"log"
	"net"
	"net/http"
	"os"
	"time"

	"github.com/rs/zerolog"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/go-redis/redis/v8"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"

	auth "github.com/go-park-mail-ru/2024_2_GOATS/auth_service/pkg/auth_v1"
	"github.com/go-park-mail-ru/2024_2_GOATS/config"
	authApi "github.com/go-park-mail-ru/2024_2_GOATS/internal/app/auth/delivery"
	authServ "github.com/go-park-mail-ru/2024_2_GOATS/internal/app/auth/service"
	"github.com/go-park-mail-ru/2024_2_GOATS/internal/app/client"

	// movieApi "github.com/go-park-mail-ru/2024_2_GOATS/internal/app/movie/delivery"
	// movieRepo "github.com/go-park-mail-ru/2024_2_GOATS/internal/app/movie/repository"
	// movieServ "github.com/go-park-mail-ru/2024_2_GOATS/internal/app/movie/service"
	"github.com/go-park-mail-ru/2024_2_GOATS/internal/app/router"
	userApi "github.com/go-park-mail-ru/2024_2_GOATS/internal/app/user/delivery"
	userServ "github.com/go-park-mail-ru/2024_2_GOATS/internal/app/user/service"
	"github.com/go-park-mail-ru/2024_2_GOATS/internal/middleware"
	user "github.com/go-park-mail-ru/2024_2_GOATS/user_service/pkg/user_v1"
)

type App struct {
	Database          *sql.DB
	Redis             *redis.Client
	Config            *config.Config
	Logger            *zerolog.Logger
	Es                *elasticsearch.Client
	Server            *http.Server
	AcceptConnections bool
}

func New(isTest bool) (*App, error) {
	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	logger := zerolog.New(os.Stdout).With().Timestamp().Logger()

	cfg, err := config.New(isTest)
	if err != nil {
		return nil, fmt.Errorf("error initialize app cfg: %w", err)
	}

	// ctx := config.WrapContext(context.Background(), cfg)
	// ctxDBTimeout, cancel := context.WithTimeout(ctx, 30*time.Second)
	// defer cancel()

	// database, err := db.SetupDatabase(ctxDBTimeout, cancel)
	// if err != nil {
	// 	return nil, fmt.Errorf("error initialize database: %w", err)
	// }

	addr := fmt.Sprintf("%s:%d", cfg.Databases.Redis.Host, cfg.Databases.Redis.Port)
	rdb := redis.NewClient(&redis.Options{Addr: addr})

	cfgEl := elasticsearch.Config{
		Addresses: []string{"http://elasticsearch:9200"},
		Transport: &http.Transport{
			MaxIdleConnsPerHost:   10,
			ResponseHeaderTimeout: time.Second,
			DialContext:           (&net.Dialer{Timeout: time.Second}).DialContext,
			TLSClientConfig:       &tls.Config{MinVersion: tls.VersionTLS12}}}

	esClient, err := elasticsearch.NewClient(cfgEl)

	if err != nil {
		return nil, fmt.Errorf("failed to create Elasticsearch client: %w", err)
	}

	return &App{
		// Database: database,
		Redis:  rdb,
		Config: cfg,
		Es:     esClient,
		Logger: &logger,
	}, nil
}

func (a *App) Run() {
	ctx := config.WrapContext(context.Background(), a.Config)
	aGrpcConn, err := grpc.NewClient(
		"auth_app:8081",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatalf("cant connect to grpc")
	}

	defer aGrpcConn.Close()

	uGrpcConn, err := grpc.NewClient(
		"user_app:8082",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatalf("cant connect to grpc")
	}

	defer uGrpcConn.Close()

	sessManager := client.NewAuthClient(auth.NewSessionRPCClient(aGrpcConn))
	usrManager := client.NewUserClient(user.NewUserRPCClient(uGrpcConn))

	srvUser := userServ.NewUserService(usrManager)
	delUser := userApi.NewUserHandler(ctx, srvUser)

	srvAuth := authServ.NewAuthService(sessManager, usrManager)
	delAuth := authApi.NewAuthHandler(ctx, srvAuth, srvUser)

	// repoMov := movieRepo.NewMovieRepository(a.Database, rdb, esClient)
	// srvMov := movieServ.NewMovieService(repoMov, repoUser)
	// delMov := movieApi.NewMovieHandler(srvMov)

	// repoRoom := roomRepo.NewRepository(database, rdb)
	// srvRoom := roomServ.NewService(repoRoom, srvMov)
	// roomHub := ws.NewRoomHub()
	// delRoom := roomApi.NewRoomHandler(srvRoom, roomHub)

	// go roomHub.Run() // Запуск обработчика Hub'a

	mx := mux.NewRouter()
	authMW := middleware.NewSessionMiddleware(srvAuth)
	router.UseCommonMiddlewares(mx, authMW)
	router.SetupCsrf(mx)
	router.SetupAuth(delAuth, mx)
	router.SetupUser(delUser, mx)
	// router.SetupMovie(delMov, mx)
	ctxValues := config.FromContext(ctx)

	srv := &http.Server{
		Addr:         fmt.Sprintf(":%d", ctxValues.Listener.Port),
		Handler:      mx,
		ReadTimeout:  ctxValues.Listener.Timeout,
		WriteTimeout: ctxValues.Listener.Timeout,
		IdleTimeout:  ctxValues.Listener.IdleTimeout,
	}

	a.Server = srv
	mx.Use(a.AppReadyMiddleware)

	a.Logger.Info().Msgf("Server is listening: %s", srv.Addr)

	// Not ready yet
	defer func() {
		if err := a.GracefulShutdown(); err != nil {
			a.Logger.Fatal().Msgf("failed to graceful shutdown: %v", err)
		}
	}()

	a.AcceptConnections = true
	if err := srv.ListenAndServe(); err != nil {
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

	// var wg sync.WaitGroup
	// errChan := make(chan error, 2)

	// shutdownFuncs := []func() error{
	// 	a.Database.Close,
	// 	a.Redis.Close,
	// }

	// wg.Add(len(shutdownFuncs))

	// for _, shutdownFunc := range shutdownFuncs {
	// 	go func(shutdownFunc func() error) {
	// 		defer wg.Done()
	// 		if err := shutdownFunc(); err != nil {
	// 			errChan <- err
	// 		}
	// 	}(shutdownFunc)
	// }

	// wg.Wait()
	// close(errChan)

	// var errs []error
	// for err := range errChan {
	// 	if err != nil {
	// 		errs = append(errs, err)
	// 	}
	// }

	// if len(errs) > 0 {
	// 	return fmt.Errorf("shutdown errors: %v", errs)
	// }

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
