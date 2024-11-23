package auth

import (
	"context"
	"errors"
	"fmt"
	"net"
	"os"
	"time"

	"github.com/go-park-mail-ru/2024_2_GOATS/review/config"
	"github.com/go-park-mail-ru/2024_2_GOATS/review/internal/db"
	"github.com/go-park-mail-ru/2024_2_GOATS/review/internal/interceptors"
	"github.com/go-park-mail-ru/2024_2_GOATS/review/internal/review/delivery"
	"github.com/go-park-mail-ru/2024_2_GOATS/review/internal/review/repository"
	"github.com/go-park-mail-ru/2024_2_GOATS/review/internal/review/service"
	review "github.com/go-park-mail-ru/2024_2_GOATS/review/pkg/review_v1"
	"github.com/go-redis/redis/v8"
	"github.com/rs/zerolog"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
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

	reflection.Register(srv)

	ctx := config.WrapContext(context.Background(), cfg)
	ctx, cancel := context.WithTimeout(ctx, time.Second*30)
	defer cancel()

	db, err := db.SetupDatabase(ctx, cancel)
	if err != nil {
		return nil, fmt.Errorf("error initialize user_app database: %w", err)
	}

	//ctx := config.WrapRedisContext(context.Background(), &cfg.Databases.Redis)

	sessRepo := repository.NewReviewRepository(db)
	sessServ := service.NewReviewService(sessRepo)
	review.RegisterReviewServer(srv, delivery.NewReviewHandler(ctx, sessServ))

	return &AuthApp{
		srv:    srv,
		logger: &logger,
		cfg:    cfg,
	}, nil
}

func (a *AuthApp) Run() {
	lis, err := net.Listen("tcp", ":8081")
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
