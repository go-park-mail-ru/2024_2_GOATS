package validation

import (
	"context"

	"github.com/go-park-mail-ru/2024_2_GOATS/validation-service/internal/app/api/converter"
	desc "github.com/go-park-mail-ru/2024_2_GOATS/validation-service/internal/pb/validation"
)

func (i *Implementation) ValidateRegistration(ctx context.Context, req *desc.ValidateRegistrationRequest) (*desc.ValidationResponse, error) {
	validData := i.validationService.ValidateRegistration(i.ctx, converter.ToUserRegisterDataFromDesc(req))
	descErrors := make([]*desc.ErrorMessage, 0)
	for _, errData := range validData.Errors {
		descErrors = append(descErrors, converter.ToErrorsFromServ(&errData))
	}

	return &desc.ValidationResponse{
		Success: validData.Success,
		Errors:  descErrors,
	}, nil
}
