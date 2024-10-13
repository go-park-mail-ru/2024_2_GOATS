package delivery

import (
	"context"
	"net/http"

	"github.com/go-park-mail-ru/2024_2_GOATS/internal/app/auth/delivery/validation"
	errVals "github.com/go-park-mail-ru/2024_2_GOATS/internal/app/errors"
	"github.com/go-park-mail-ru/2024_2_GOATS/internal/app/models"
	authModels "github.com/go-park-mail-ru/2024_2_GOATS/internal/app/models/auth"
)

func (i *Implementation) Register(ctx context.Context, registerData *authModels.RegisterData) (*authModels.AuthResponse, *models.ErrorResponse) {
	success := true
	errors := make([]errVals.ErrorObj, 0)

	if err := validation.ValidatePassword(registerData.Password, registerData.PasswordConfirmation); err != nil {
		success = addError(errVals.ErrInvalidPasswordCode, *err, &errors)
	}

	if err := validation.ValidateEmail(registerData.Email); err != nil {
		success = addError(errVals.ErrInvalidEmailCode, *err, &errors)
	}

	if len(errors) > 0 {
		return nil, &models.ErrorResponse{
			Success:    success,
			Errors:     errors,
			StatusCode: http.StatusBadRequest,
		}
	}

	resp, errData := i.authService.Register(ctx, registerData)
	if errData != nil {
		return nil, errData
	}

	return resp, nil
}

func (i *Implementation) Login(ctx context.Context, loginData *authModels.LoginData) (*authModels.AuthResponse, *models.ErrorResponse) {
	resp, errData := i.authService.Login(ctx, loginData)
	if errData != nil {
		return nil, errData
	}

	return resp, nil
}

func (i *Implementation) Session(ctx context.Context, cookie string) (*authModels.SessionResponse, *models.ErrorResponse) {
	resp, errData := i.authService.Session(ctx, cookie)
	if errData != nil {
		return nil, errData
	}

	return resp, nil
}

func (i *Implementation) Logout(ctx context.Context, cookie string) (*authModels.AuthResponse, *models.ErrorResponse) {
	validErr := validation.ValidateCookie(cookie)
	if validErr != nil {
		return nil, &models.ErrorResponse{
			Success:    false,
			StatusCode: http.StatusBadRequest,
			Errors:     []errVals.ErrorObj{{Code: errVals.ErrBrokenCookieCode, Error: *validErr}},
		}
	}

	resp, errData := i.authService.Logout(ctx, cookie)
	if errData != nil {
		return nil, errData
	}

	return resp, nil
}

func addError(code string, err errVals.CustomError, errors *[]errVals.ErrorObj) bool {
	errStruct := errVals.ErrorObj{
		Code:  code,
		Error: err,
	}

	*errors = append(*errors, errStruct)

	return false
}
