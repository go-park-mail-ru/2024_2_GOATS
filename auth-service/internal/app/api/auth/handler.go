package handler

import (
	"context"
	"github.com/go-park-mail-ru/2024_2_GOATS/internal/app/model"
	desc "github.com/go-park-mail-ru/2024_2_GOATS/internal/pb/auth"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type IAuth interface {
	Register(ctx context.Context, email string, password string, passwordConfirm string, nickname string, sex string, birthdate *timestamppb.Timestamp) (bool, string, []*model.ErrorMessage)
	Login(ctx context.Context, email string, password string) (bool, string, []*model.ErrorMessage)
	Logout(ctx context.Context, email string) (success bool, Errors []*model.ErrorMessage)
	Close() error
}

type Implementation struct {
	desc.UnimplementedAuthServer
	auth IAuth
}

func NewAuth(rf IAuth) *Implementation {
	return &Implementation{auth: rf}
}
