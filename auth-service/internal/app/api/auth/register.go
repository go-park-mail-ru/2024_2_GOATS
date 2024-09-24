package handler

import (
	"context"
	desc "github.com/go-park-mail-ru/2024_2_GOATS/internal/pb/auth"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"log"
)

func (i *Implementation) Register(ctx context.Context, req *desc.RegisterRequest) (*desc.AuthResponse, error) {

	success, token, err := i.auth.Register(ctx, req.Email, req.Password, req.PasswordConfirm, req.Sex, req.Nickname, req.Birthdate)

	var errSlice []*desc.ErrorMessage

	for _, v := range err {
		errSlice = append(errSlice, &desc.ErrorMessage{
			Code:  v.Code,
			Error: v.Error,
		})
	}

	if success {
		// Устанавливаем куки с токеном

		// Добавляем токен в метаданные контекста
		md := metadata.Pairs("auth-token", token)
		grpc.SetHeader(ctx, md)

		log.Printf("Set metadata in gRPC: %v", md)
	}

	return &desc.AuthResponse{
		Success: success,
		Errors:  errSlice,
	}, nil
}
