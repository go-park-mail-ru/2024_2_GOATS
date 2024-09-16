package handler

import (
	"context"
	desc "github.com/go-park-mail-ru/2024_2_GOATS/internal/pb/auth"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"log"
)

func (i *Implementation) Login(ctx context.Context, req *desc.AuthRequest) (*desc.AuthResponse, error) {

	success, token, err := i.auth.Login(ctx, req.Email, req.Password)

	var errSlice []*desc.ErrorMessage

	for _, v := range err {
		errSlice = append(errSlice, &desc.ErrorMessage{
			Code:  v.Code,
			Error: v.Error,
		})
	}

	// Добавляем токен в метаданные контекста
	md := metadata.Pairs("auth-token", token)
	grpc.SetHeader(ctx, md)

	log.Printf("Set metadata in gRPC: %v", md)

	return &desc.AuthResponse{
		Success: success,
		Errors:  errSlice,
	}, nil
}
