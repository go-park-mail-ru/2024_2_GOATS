package validation

import (
	"context"

	"github.com/go-park-mail-ru/2024_2_GOATS/validation-service/internal/app/model"
	desc "github.com/go-park-mail-ru/2024_2_GOATS/validation-service/internal/pb/validation"
)

type ValidationService interface {
	ValidateRegistration(ctx context.Context, userData *model.UserRegisterData) *model.ValidationResponse
}

type Implementation struct {
	desc.UnimplementedValidationServer
	ctx               context.Context
	validationService ValidationService
}

func NewImplementation(ctx context.Context, validationService ValidationService) *Implementation {
	return &Implementation{
		ctx:               ctx,
		validationService: validationService,
	}
}
