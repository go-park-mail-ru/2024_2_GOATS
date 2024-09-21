package validation

import (
	"context"

	errVals "github.com/go-park-mail-ru/2024_2_GOATS/validation-service/internal/app/errors"
	"github.com/go-park-mail-ru/2024_2_GOATS/validation-service/internal/app/model"
	userHelper "github.com/go-park-mail-ru/2024_2_GOATS/validation-service/internal/app/service/validation/helpers"
)

func (s *serv) ValidateRegistration(ctx context.Context, userData *model.UserRegisterData) *model.ValidationResponse {
	success := true
	errors := make([]model.ErrorResponse, 0)

	if err := userHelper.ValidatePassword(userData.Password, userData.PasswordConfirm); err != nil {
		success = AddError(errVals.ErrInvalidPasswordCode, err, &errors)
	}

	if err := userHelper.ValidateEmail(userData.Email); err != nil {
		success = AddError(errVals.ErrInvalidEmailCode, err, &errors)
	}

	if err := userHelper.ValidateBirthdate(userData.Birthday); err != nil {
		success = AddError(errVals.ErrInvalidBirthdateCode, err, &errors)
	}

	if err := userHelper.ValidateSex(userData.Sex); err != nil {
		success = AddError(errVals.ErrInvalidSexCode, err, &errors)
	}

	return &model.ValidationResponse{
		Success: success,
		Errors:  errors,
	}
}

func AddError(code string, err error, errors *[]model.ErrorResponse) bool {
	errStruct := model.ErrorResponse{
		Code:     code,
		ErrorObj: err,
	}

	*errors = append(*errors, errStruct)

	return false
}
