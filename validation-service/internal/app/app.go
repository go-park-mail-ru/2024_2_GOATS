package app

import (
	"context"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	validationAPI "github.com/go-park-mail-ru/2024_2_GOATS/validation-service/internal/app/api/validation"
	validationService "github.com/go-park-mail-ru/2024_2_GOATS/validation-service/internal/app/service/validation"
	desc "github.com/go-park-mail-ru/2024_2_GOATS/validation-service/internal/pb/validation"
)

type App struct {
	serverPort        int
	validationService validationAPI.ValidationService
}

func New(ctx context.Context, serverPort int) (*App, error) {
	validationService := validationService.NewService()

	return &App{
		serverPort:        serverPort,
		validationService: validationService,
	}, nil
}

func (a *App) Run() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", a.serverPort))
	if err != nil {
		log.Fatal("Failed to listen: %v", err)
	}

	validationSrv := validationService.NewService()
	s := grpc.NewServer()
	reflection.Register(s)

	desc.RegisterValidationServer(s, validationAPI.NewImplementation(validationSrv))

	log.Printf("Server listening at %v", lis.Addr())

	if err = s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
