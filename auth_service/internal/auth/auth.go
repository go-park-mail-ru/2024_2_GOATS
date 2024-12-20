package auth

import (
	"context"
	"errors"
	"fmt"
	"net"
	"net/http"
	"os"

	"github.com/go-park-mail-ru/2024_2_GOATS/auth_service/config"
	"github.com/go-park-mail-ru/2024_2_GOATS/auth_service/internal/auth/delivery"
	"github.com/go-park-mail-ru/2024_2_GOATS/auth_service/internal/auth/repository"
	"github.com/go-park-mail-ru/2024_2_GOATS/auth_service/internal/auth/service"
	"github.com/go-park-mail-ru/2024_2_GOATS/auth_service/internal/interceptors"
	auth "github.com/go-park-mail-ru/2024_2_GOATS/auth_service/pkg/auth_v1"
	"github.com/go-redis/redis/v8"
	grpc_prometheus "github.com/grpc-ecosystem/go-grpc-prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/rs/zerolog"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

// AppAuth is a root struct of auth_service
type AppAuth struct {
	rdb    *redis.Client
	logger *zerolog.Logger
	srv    *grpc.Server
	cfg    *config.Config
}

// New returns an instance of AppAuth
func New(isTest bool) (*AppAuth, error) {
	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	logger := zerolog.New(os.Stdout).With().Timestamp().Logger()

	cfg, err := config.New(isTest)
	if err != nil {
		return nil, fmt.Errorf("error initialize app cfg: %w", err)
	}

	srv := grpc.NewServer(
		grpc.UnaryInterceptor(interceptors.ChainUnaryInterceptors(
			interceptors.WithLogger,
			interceptors.AccessLogInterceptor,
		)),
	)

	addr := fmt.Sprintf("%s:%d", cfg.Databases.Redis.Host, cfg.Databases.Redis.Port)
	rdb := redis.NewClient(&redis.Options{Addr: addr})
	ctx := config.WrapRedisContext(context.Background(), &cfg.Databases.Redis)

	reflection.Register(srv)
	sessRepo := repository.NewAuthRepository(rdb)
	sessServ := service.NewAuthService(sessRepo)
	auth.RegisterSessionRPCServer(srv, delivery.NewAuthManager(ctx, sessServ))

	grpc_prometheus.Register(srv)

	go func() {
		http.Handle("/metrics", promhttp.Handler())
		if err := http.ListenAndServe(":9081", nil); err != nil {
			logger.Error().Err(err).Msg("Metrics stopped")
		}
	}()

	return &AppAuth{
		rdb:    rdb,
		srv:    srv,
		logger: &logger,
		cfg:    cfg,
	}, nil
}

// Run starts grpc server
func (a *AppAuth) Run() {
	lis, err := net.Listen("tcp", a.cfg.Listener.Port)
	if err != nil {
		a.logger.Fatal().Msgf("failed to setup listener: %v", err)
	}

	a.logger.Info().Msgf("starting server at %s", a.cfg.Listener.Port)

	defer func() {
		if err := a.GracefulShutdown(); err != nil {
			a.logger.Fatal().Msgf("failed to graceful shutdown: %v", err)
		}
	}()

	if err := a.srv.Serve(lis); err != nil {
		if errors.Is(err, grpc.ErrServerStopped) {
			a.logger.Info().Msg("server closed under request")
		} else {
			a.logger.Info().Msgf("server stopped: %v", err)
		}
	}
}

// GracefulShutdown gracefully shutdowns AppAuth
func (a *AppAuth) GracefulShutdown() error {
	a.logger.Info().Msg("Starting graceful shutdown")

	if err := a.rdb.Close(); err != nil {
		a.logger.Error().Err(err).Msg("Cannot graceful shutdown Redis")
		return err
	}

	a.srv.GracefulStop()
	a.logger.Info().Msg("Auth grpc shut down")

	return nil
}
