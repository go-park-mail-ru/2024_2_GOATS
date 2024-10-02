package errors

import (
	"encoding/json"
	"errors"
)

var (
	ErrInvalidEmailCode          = "invalid_email"
	ErrInvalidEmailText          = CustomError{Err: errors.New("email is incorrect")}
	ErrInvalidPasswordCode       = "invalid_password"
	ErrInvalidPasswordText       = CustomError{Err: errors.New("password is too short. The minimal len is 8")}
	ErrInvalidPasswordsMatchText = CustomError{Err: errors.New("password doesnt match with passwordConfirmation")}
	ErrUserNotFoundCode          = "user_not_found"
	ErrUserNotFoundText          = CustomError{Err: errors.New("cannot find user with translated email")}
	ErrServerCode                = "something_went_wrong"
	ErrGenerateTokenCode         = "auth_token_generation_error"
	ErrCreateUserCode            = "create_user_error"
	ErrUnauthorizedCode          = "user_unauthorized"
)

type CustomError struct {
	Err error
}

type ErrorObj struct {
	Code  string
	Error CustomError
}

func (ce *CustomError) MarshalJSON() ([]byte, error) {
	return json.Marshal(ce.Err.Error())
}

func NewErrorObj(code string, text CustomError) *ErrorObj {
	return &ErrorObj{
		Code:  code,
		Error: text,
	}
}
