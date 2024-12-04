package payment

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"net"
	"net/http"
	"os"
	"time"

	"github.com/go-park-mail-ru/2024_2_GOATS/payment_service/config"
	"github.com/go-park-mail-ru/2024_2_GOATS/payment_service/internal/db"
	"github.com/go-park-mail-ru/2024_2_GOATS/payment_service/internal/interceptors"
	"github.com/go-park-mail-ru/2024_2_GOATS/payment_service/internal/payment/delivery"
	"github.com/go-park-mail-ru/2024_2_GOATS/payment_service/internal/payment/repository"
	"github.com/go-park-mail-ru/2024_2_GOATS/payment_service/internal/payment/service"
	payment "github.com/go-park-mail-ru/2024_2_GOATS/payment_service/pkg/payment_v1"
	"github.com/grpc-ecosystem/go-grpc-prometheus"
	_ "github.com/lib/pq"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/rs/zerolog"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type PaymentApp struct {
	logger   *zerolog.Logger
	database *sql.DB
	config   *config.Config
	srv      *grpc.Server
}

func New(isTest bool) (*PaymentApp, error) {
	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	logger := zerolog.New(os.Stdout).With().Timestamp().Logger()

	cfg, err := config.New(isTest)
	if err != nil {
		return nil, fmt.Errorf("error initialize payment_app cfg: %w", err)
	}

	ctx := config.WrapContext(context.Background(), cfg)
	ctx, cancel := context.WithTimeout(ctx, time.Second*30)
	defer cancel()

	db, err := db.SetupDatabase(ctx, cancel)
	if err != nil {
		return nil, fmt.Errorf("error initialize payment_app database: %w", err)
	}

	srv := grpc.NewServer(
		grpc.UnaryInterceptor(interceptors.ChainUnaryInterceptors(
			interceptors.WithLogger,
			interceptors.AccessLogInterceptor,
		)),
	)

	reflection.Register(srv)
	usrRepo := repository.NewPaymentRepository(db)
	usrServ := service.NewPaymentService(usrRepo)
	payment.RegisterPaymentRPCServer(srv, delivery.NewPaymentHandler(ctx, usrServ))
	grpc_prometheus.Register(srv)

	go func() {
		http.Handle("/metrics", promhttp.Handler())
		http.ListenAndServe(":9082", nil)
	}()

	return &PaymentApp{
		logger:   &logger,
		database: db,
		config:   cfg,
		srv:      srv,
	}, nil
}

func (ua *PaymentApp) Run() {
	lis, err := net.Listen("tcp", ua.config.Listener.Port)
	if err != nil {
		ua.logger.Fatal().Msgf("failed to setup payment_app listener: %v", err)
	}

	defer func() {
		if err := ua.GracefulShutdown(); err != nil {
			ua.logger.Fatal().Msgf("failed to graceful shutdown payment_app %v", err)
		}
	}()

	ua.logger.Info().Msgf("starting payment_app server at %s", ua.config.Listener.Port)

	if err := ua.srv.Serve(lis); err != nil {
		if errors.Is(err, grpc.ErrServerStopped) {
			ua.logger.Info().Msg("server closed under request")
		} else {
			ua.logger.Info().Msgf("server stopped: %v", err)
		}
	}
}

func (ua *PaymentApp) GracefulShutdown() error {
	ua.logger.Info().Msg("Starting graceful shutdown payment_app")
	if err := ua.database.Close(); err != nil {
		ua.logger.Error().Err(err).Msg("failed to close payment_app Postgres")
		return err
	}

	ua.srv.GracefulStop()
	ua.logger.Info().Msg("payment_app is shutdown")

	return nil
}
