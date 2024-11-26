package movie

import (
	"context"
	"crypto/tls"
	"database/sql"
	"errors"
	"fmt"
	"github.com/elastic/go-elasticsearch/v7"
	"github.com/go-park-mail-ru/2024_2_GOATS/movie_service/config"
	"github.com/go-park-mail-ru/2024_2_GOATS/movie_service/internal/db"
	"github.com/go-park-mail-ru/2024_2_GOATS/movie_service/internal/interceptors"
	"github.com/go-park-mail-ru/2024_2_GOATS/movie_service/internal/movie/delivery"
	"github.com/go-park-mail-ru/2024_2_GOATS/movie_service/internal/movie/repository"
	"github.com/go-park-mail-ru/2024_2_GOATS/movie_service/internal/movie/service"
	"github.com/go-park-mail-ru/2024_2_GOATS/movie_service/pkg/clients"
	movie "github.com/go-park-mail-ru/2024_2_GOATS/movie_service/pkg/movie_v1"
	user "github.com/go-park-mail-ru/2024_2_GOATS/movie_service/pkg/user_v1"
	"github.com/rs/zerolog"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/reflection"
	"net"
	"net/http"
	"os"
	"time"
)

type MovieApp struct {
	database *sql.DB
	logger   *zerolog.Logger
	srv      *grpc.Server
	cfg      *config.Config
}

func New(isTest bool) (*MovieApp, error) {
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

	ctx := config.WrapContext(context.Background(), cfg)
	ctx, cancel := context.WithTimeout(ctx, time.Second*30)
	defer cancel()
	db, err := db.SetupDatabase(ctx, cancel)
	if err != nil {
		return nil, fmt.Errorf("error initialize user_app database: %w", err)
	}

	cfgEl := elasticsearch.Config{
		Addresses: []string{"http://elasticsearch:9200"},
		Transport: &http.Transport{
			MaxIdleConnsPerHost:   10,
			ResponseHeaderTimeout: time.Second,
			DialContext:           (&net.Dialer{Timeout: time.Second}).DialContext,
			TLSClientConfig:       &tls.Config{MinVersion: tls.VersionTLS12}}}

	esClient, err := elasticsearch.NewClient(cfgEl)

	reflection.Register(srv)
	sessRepo := repository.NewMovieRepository(db, esClient)

	uGrpcConn, err := grpc.NewClient(
		"user_app:8082",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)

	usrClient := client.NewUserClient(user.NewUserRPCClient(uGrpcConn))

	movieService := service.NewMovieService(sessRepo, usrClient)

	movie.RegisterMovieServiceServer(srv, delivery.NewMovieHandler(ctx, movieService))

	return &MovieApp{
		srv:    srv,
		logger: &logger,
		cfg:    cfg,
	}, nil
}

func (a *MovieApp) Run() {
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

func (ua *MovieApp) GracefulShutdown() error {
	ua.logger.Info().Msg("Starting graceful shutdown user_app")
	if err := ua.database.Close(); err != nil {
		ua.logger.Error().Err(err).Msg("failed to close user_app Postgres")
		return err
	}

	ua.srv.GracefulStop()
	ua.logger.Info().Msg("user_app is shutdown")

	return nil
}
