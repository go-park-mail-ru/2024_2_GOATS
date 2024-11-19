package user

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"net"
	"os"
	"time"

	"github.com/go-park-mail-ru/2024_2_GOATS/user_service/config"
	"github.com/go-park-mail-ru/2024_2_GOATS/user_service/internal/db"
	"github.com/go-park-mail-ru/2024_2_GOATS/user_service/internal/user/repository"
	"github.com/go-park-mail-ru/2024_2_GOATS/user_service/internal/user/service"
	"github.com/go-park-mail-ru/2024_2_GOATS/user_service/internal/user/delivery"
	user "github.com/go-park-mail-ru/2024_2_GOATS/user_service/pkg/user_v1"
	_ "github.com/lib/pq"
	"github.com/rs/zerolog"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type UserApp struct {
	logger   *zerolog.Logger
	database *sql.DB
	config   *config.Config
	srv      *grpc.Server
}

func New(isTest bool) (*UserApp, error) {
	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	logger := zerolog.New(os.Stdout).With().Timestamp().Logger()

	cfg, err := config.New(isTest)
	if err != nil {
		return nil, fmt.Errorf("error initialize user_app cfg: %w", err)
	}

	ctx := config.WrapContext(context.Background(), cfg)
	ctx, cancel := context.WithTimeout(ctx, time.Second*30)
	defer cancel()

	db, err := db.SetupDatabase(ctx, cancel)
	if err != nil {
		return nil, fmt.Errorf("error initialize user_app database: %w", err)
	}

	srv := grpc.NewServer()
	reflection.Register(srv)

	usrRepo := repository.NewUserRepository(db)
	usrServ := service.NewUserService(usrRepo)
	user.RegisterUserRPCServer(srv, delivery.NewUserHandler(ctx, usrServ))

	return &UserApp{
		logger:   &logger,
		database: db,
		config:   cfg,
		srv:      srv,
	}, nil
}

func (ua *UserApp) Run() {
	lis, err := net.Listen("tcp", ua.config.Listener.Port)
	if err != nil {
		ua.logger.Fatal().Msgf("failed to setup user_app listener: %v", err)
	}

	defer func() {
		if err := ua.GracefulShutdown(); err != nil {
			ua.logger.Fatal().Msgf("failed to graceful shutdown user_app %v", err)
		}
	}()

	ua.logger.Info().Msgf("starting user_app server at %s", ua.config.Listener.Port)

	if err := ua.srv.Serve(lis); err != nil {
		if errors.Is(err, grpc.ErrServerStopped) {
			ua.logger.Info().Msg("server closed under request")
		} else {
			ua.logger.Info().Msgf("server stopped: %v", err)
		}
	}
}

func (ua *UserApp) GracefulShutdown() error {
	ua.logger.Info().Msg("Starting graceful shutdown user_app")
	if err := ua.database.Close(); err != nil {
		ua.logger.Error().Err(err).Msg("failed to close user_app Postgres")
		return err
	}

	ua.srv.GracefulStop()
	ua.logger.Info().Msg("user_app is shutdown")

	return nil
}
