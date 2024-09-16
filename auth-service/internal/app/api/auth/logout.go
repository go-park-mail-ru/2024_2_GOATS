package handler

import (
	"context"
	desc "github.com/go-park-mail-ru/2024_2_GOATS/internal/pb/auth"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"log"
)

func (i *Implementation) Logout(ctx context.Context, req *desc.LogoutRequest) (*desc.AuthResponse, error) {

	success, err := i.auth.Logout(ctx, req.Email)

	var errSlice []*desc.ErrorMessage

	for _, v := range err {
		errSlice = append(errSlice, &desc.ErrorMessage{
			Code:  v.Code,
			Error: v.Error,
		})
	}

	if success {
		// Удаляем куку (очищаем токен)
		// Вот тут нужно проверить что да как
		md := metadata.Pairs("Set-Cookie", "auth-token=; Max-Age=0; Path=/; HttpOnly; SameSite=Lax")
		grpc.SetHeader(ctx, md)

		log.Printf("Removed auth-token in gRPC: %v", md)
	}

	return &desc.AuthResponse{
		Success: success,
		Errors:  errSlice,
	}, nil
}
