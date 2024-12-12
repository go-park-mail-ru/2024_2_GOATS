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
	grpc_prometheus "github.com/grpc-ecosystem/go-grpc-prometheus"

	// postgres driver
	_ "github.com/lib/pq"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/rs/zerolog"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

// AppPayment is a root struct of payment_service
type AppPayment struct {
	logger   *zerolog.Logger
	database *sql.DB
	config   *config.Config
	srv      *grpc.Server
}

// New returns an instance of AppPayment
func New(isTest bool) (*AppPayment, error) {
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
	payment.RegisterPaymentRPCServer(srv, delivery.NewPaymentHandler(usrServ))
	grpc_prometheus.Register(srv)

	go func() {
		http.Handle("/metrics", promhttp.Handler())
		if err := http.ListenAndServe(":9084", nil); err != nil {
			logger.Error().Err(err).Msg("Metrics stopped")
		}
	}()

	return &AppPayment{
		logger:   &logger,
		database: db,
		config:   cfg,
		srv:      srv,
	}, nil
}

// Run starts grpc server
func (ua *AppPayment) Run() {
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

// GracefulShutdown gracefully shutdowns AppPayment
func (ua *AppPayment) GracefulShutdown() error {
	ua.logger.Info().Msg("Starting graceful shutdown payment_app")
	if err := ua.database.Close(); err != nil {
		ua.logger.Error().Err(err).Msg("failed to close payment_app Postgres")
		return err
	}

	ua.srv.GracefulStop()
	ua.logger.Info().Msg("payment_app is shutdown")

	return nil
}
