package app

import (
	"context"
	"errors"
	"fmt"
	ws "github.com/go-park-mail-ru/2024_2_GOATS/internal/app/room/ws"
	"github.com/go-redis/redis/v8"
	"github.com/prometheus/client_golang/prometheus/promhttp"

	"log"
	"net/http"
	"os"
	"time"

	"github.com/rs/zerolog"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/gorilla/mux"

	auth "github.com/go-park-mail-ru/2024_2_GOATS/auth_service/pkg/auth_v1"
	"github.com/go-park-mail-ru/2024_2_GOATS/config"
	authApi "github.com/go-park-mail-ru/2024_2_GOATS/internal/app/auth/delivery"
	authServ "github.com/go-park-mail-ru/2024_2_GOATS/internal/app/auth/service"
	"github.com/go-park-mail-ru/2024_2_GOATS/internal/app/client"

	movieApi "github.com/go-park-mail-ru/2024_2_GOATS/internal/app/movie/delivery"
	movieServ "github.com/go-park-mail-ru/2024_2_GOATS/internal/app/movie/service"
	roomRepo "github.com/go-park-mail-ru/2024_2_GOATS/internal/app/room/repository"
	roomApi "github.com/go-park-mail-ru/2024_2_GOATS/internal/app/room/room_handler"
	roomServ "github.com/go-park-mail-ru/2024_2_GOATS/internal/app/room/service"
	payApi "github.com/go-park-mail-ru/2024_2_GOATS/internal/app/payment/delivery"
	payServ "github.com/go-park-mail-ru/2024_2_GOATS/internal/app/payment/service"
	"github.com/go-park-mail-ru/2024_2_GOATS/internal/app/router"
	subApi "github.com/go-park-mail-ru/2024_2_GOATS/internal/app/subscription/delivery"
	subServ "github.com/go-park-mail-ru/2024_2_GOATS/internal/app/subscription/service"
	userApi "github.com/go-park-mail-ru/2024_2_GOATS/internal/app/user/delivery"
	userServ "github.com/go-park-mail-ru/2024_2_GOATS/internal/app/user/service"
	"github.com/go-park-mail-ru/2024_2_GOATS/internal/middleware"
	movie "github.com/go-park-mail-ru/2024_2_GOATS/movie_service/pkg/movie_v1"
	payment "github.com/go-park-mail-ru/2024_2_GOATS/payment_service/pkg/payment_v1"
	user "github.com/go-park-mail-ru/2024_2_GOATS/user_service/pkg/user_v1"
)

// App root facade struct
type App struct {
	Config            *config.Config
	Logger            *zerolog.Logger
	Server            *http.Server
	AcceptConnections bool
}

// New returns an instance of App
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

// Run Starts http server
func (a *App) Run() {
	ctx := config.WrapContext(context.Background(), a.Config)
	aGrpcConn, err := grpc.NewClient(
		"auth_app:8081",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatalf("cant connect to grpc")
	}

	defer func() {
		if clErr := aGrpcConn.Close(); clErr != nil {
			log.Fatal("cannot close authGrpcConnection")
		}
	}()

	uGrpcConn, err := grpc.NewClient(
		"user_app:8082",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatalf("cant connect to grpc")
	}

	defer func() {
		if clErr := uGrpcConn.Close(); clErr != nil {
			log.Fatal("cannot close userGrpcConnection")
		}
	}()

	mGrpcConn, err := grpc.NewClient(
		"movie_app:8083",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatalf("cant connect to grpc")
	}

	defer func() {
		if clErr := mGrpcConn.Close(); clErr != nil {
			log.Fatal("cannot close movieGrpcConnection")
		}
	}()

	pGrpcConn, err := grpc.NewClient(
		"payment_app:8084",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatalf("cant connect to grpc")
	}

	defer func() {
		if clErr := pGrpcConn.Close(); clErr != nil {
			log.Fatal("cannot close paymentGrpcConnection")
		}
	}()

	sessManager := client.NewAuthClient(auth.NewSessionRPCClient(aGrpcConn))
	usrManager := client.NewUserClient(user.NewUserRPCClient(uGrpcConn))
	mvManager := client.NewMovieClient(movie.NewMovieServiceClient(mGrpcConn))
	payManager := client.NewPaymentClient(payment.NewPaymentRPCClient(pGrpcConn))

	srvUser := userServ.NewUserService(usrManager, mvManager)
	delUser := userApi.NewUserHandler(srvUser)

	srvAuth := authServ.NewAuthService(sessManager, usrManager)
	delAuth := authApi.NewAuthHandler(srvAuth, srvUser)

	srvMov := movieServ.NewMovieService(mvManager, usrManager)
	delMov := movieApi.NewMovieHandler(srvMov)

	srvPay := payServ.NewPaymentService(payManager, usrManager)
	delPay := payApi.NewPaymentHandler(srvPay)

	srvSub := subServ.NewSubscriptionService(payManager, usrManager)
	delSub := subApi.NewSubscriptionHandler(srvSub)

	addr := fmt.Sprintf("%s:%d", a.Config.Databases.Redis.Host, a.Config.Databases.Redis.Port)
	rdb := redis.NewClient(&redis.Options{Addr: addr})

	repoRoom := roomRepo.NewRepository(rdb)
	roomHub := ws.NewRoomHub()
	timer := roomServ.NewTimerManager(roomHub)
	srvRoom := roomServ.NewService(repoRoom, mvManager, usrManager, roomHub, timer)
	delRoom := roomApi.NewRoomHandler(srvRoom, roomHub)

	log.Println("XZXZXZ")
	go roomHub.Run()
	log.Println("XZ2XZ2XZ2")

	mx := mux.NewRouter()
	authMW := middleware.NewSessionMiddleware(srvAuth)
	router.UseCommonMiddlewares(mx, authMW)
	router.SetupCsrf(mx)
	router.SetupAuth(delAuth, mx)
	router.SetupUser(delUser, mx)
	router.SetupMovie(delMov, mx)
	router.SetupPayment(delPay, mx)
	router.SetupSubscription(delSub, mx)

	mx.HandleFunc("/metrics", func(w http.ResponseWriter, r *http.Request) {
		promhttp.Handler().ServeHTTP(w, r)
	}).Methods(http.MethodGet)

	ctxValues := config.FromContext(ctx)

	router.SetupRoom(roomHub, delRoom, mx)

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

// GracefulShutdown gracefully shutdowns grpc server
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
