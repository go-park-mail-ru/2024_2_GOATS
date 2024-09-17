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
	validationService ValidationService
}

func NewImplementation(validationService ValidationService) *Implementation {
	return &Implementation{
		validationService: validationService,
	}
}
