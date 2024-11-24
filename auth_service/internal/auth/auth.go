package auth

import (
	"context"
	"errors"
	"fmt"
	"net"
	"os"

	"github.com/go-park-mail-ru/2024_2_GOATS/auth_service/config"
	"github.com/go-park-mail-ru/2024_2_GOATS/auth_service/internal/auth/delivery"
	"github.com/go-park-mail-ru/2024_2_GOATS/auth_service/internal/auth/repository"
	"github.com/go-park-mail-ru/2024_2_GOATS/auth_service/internal/auth/service"
	"github.com/go-park-mail-ru/2024_2_GOATS/auth_service/internal/interceptors"
	auth "github.com/go-park-mail-ru/2024_2_GOATS/auth_service/pkg/auth_v1"
	"github.com/go-redis/redis/v8"
	"github.com/rs/zerolog"
	"google.golang.org/grpc"
)

type AuthApp struct {
	rdb    *redis.Client
	logger *zerolog.Logger
	srv    *grpc.Server
	cfg    *config.Config
}

func New(isTest bool) (*AuthApp, error) {
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

	sessRepo := repository.NewAuthRepository(rdb)
	sessServ := service.NewAuthService(sessRepo)
	auth.RegisterSessionRPCServer(srv, delivery.NewAuthManager(ctx, sessServ))

	return &AuthApp{
		rdb:    rdb,
		srv:    srv,
		logger: &logger,
		cfg:    cfg,
	}, nil
}

func (a *AuthApp) Run() {
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

func (a *AuthApp) GracefulShutdown() error {
	a.logger.Info().Msg("Starting graceful shutdown")

	if err := a.rdb.Close(); err != nil {
		a.logger.Error().Err(err).Msg("Cannot graceful shutdown Redis")
		return err
	}

	a.srv.GracefulStop()
	a.logger.Info().Msg("Auth grpc shut down")

	return nil
}
