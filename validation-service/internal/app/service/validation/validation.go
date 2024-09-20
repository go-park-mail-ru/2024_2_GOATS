package validation

import (
	"context"

	api "github.com/go-park-mail-ru/2024_2_GOATS/validation-service/internal/app/api/validation"
)

var _ api.ValidationService = (*serv)(nil)

type serv struct{
	ctx context.Context
}

func NewService(ctx context.Context) *serv {
	return &serv{
		ctx: ctx,
	}
}
