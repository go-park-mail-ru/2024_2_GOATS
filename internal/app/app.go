package app

import (
	"context"
	"errors"
	"fmt"

	"log"
	"net/http"
	"os"
	"time"

	"github.com/rs/zerolog"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"

	auth "github.com/go-park-mail-ru/2024_2_GOATS/auth_service/pkg/auth_v1"
	"github.com/go-park-mail-ru/2024_2_GOATS/config"
	authApi "github.com/go-park-mail-ru/2024_2_GOATS/internal/app/auth/delivery"
	authServ "github.com/go-park-mail-ru/2024_2_GOATS/internal/app/auth/service"
	"github.com/go-park-mail-ru/2024_2_GOATS/internal/app/client"

	movieApi "github.com/go-park-mail-ru/2024_2_GOATS/internal/app/movie/delivery"
	movieServ "github.com/go-park-mail-ru/2024_2_GOATS/internal/app/movie/service"
	"github.com/go-park-mail-ru/2024_2_GOATS/internal/app/router"
	userApi "github.com/go-park-mail-ru/2024_2_GOATS/internal/app/user/delivery"
	userServ "github.com/go-park-mail-ru/2024_2_GOATS/internal/app/user/service"
	"github.com/go-park-mail-ru/2024_2_GOATS/internal/middleware"
	movie "github.com/go-park-mail-ru/2024_2_GOATS/movie_service/pkg/movie_v1"
	user "github.com/go-park-mail-ru/2024_2_GOATS/user_service/pkg/user_v1"
)

type App struct {
	Config            *config.Config
	Logger            *zerolog.Logger
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

	return &App{
		Config: cfg,
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

	mGrpcConn, err := grpc.NewClient(
		"movie_app:8083",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatalf("cant connect to grpc")
	}

	defer mGrpcConn.Close()

	sessManager := client.NewAuthClient(auth.NewSessionRPCClient(aGrpcConn))
	usrManager := client.NewUserClient(user.NewUserRPCClient(uGrpcConn))
	mvManager := client.NewMovieClient(movie.NewMovieServiceClient(mGrpcConn))

	srvUser := userServ.NewUserService(usrManager, mvManager)
	delUser := userApi.NewUserHandler(ctx, srvUser)

	srvAuth := authServ.NewAuthService(sessManager, usrManager)
	delAuth := authApi.NewAuthHandler(ctx, srvAuth, srvUser)

	srvMov := movieServ.NewMovieService(mvManager)
	delMov := movieApi.NewMovieHandler(srvMov)

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
	router.SetupMovie(delMov, mx)
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
