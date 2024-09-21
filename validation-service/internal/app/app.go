package app

import (
	"context"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"github.com/go-park-mail-ru/2024_2_GOATS/validation-service/config"
	validationAPI "github.com/go-park-mail-ru/2024_2_GOATS/validation-service/internal/app/api/validation"
	validationService "github.com/go-park-mail-ru/2024_2_GOATS/validation-service/internal/app/service/validation"
	desc "github.com/go-park-mail-ru/2024_2_GOATS/validation-service/internal/pb/validation"
)

type App struct {
	validationService validationAPI.ValidationService
}

func New() (*App, context.Context, error) {
	ctx, err := config.CreateConfigContext()
	if err != nil {
		log.Fatal("Failed to read config: %v", err)
	}

	return &App{
		validationService: validationService.NewService(ctx),
	}, ctx, nil
}

func (a *App) Run(ctx context.Context) {
	cfg := config.GetConfigFromContext(ctx)
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", cfg.Listener.Port))
	if err != nil {
		log.Fatal("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	reflection.Register(s)
	desc.RegisterValidationServer(s, validationAPI.NewImplementation(ctx, a.validationService))

	log.Printf("Server listening at %v:%d", cfg.Listener.Address, cfg.Listener.Port)

	if err = s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
