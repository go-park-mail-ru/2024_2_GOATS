package errors

import (
	"encoding/json"
	"errors"
)

const (
	DuplicateErrKey        = "23505"
	ErrBrokenCookieCode    = "broken_cookie"
	ErrNoCookieCode        = "no_cookie_provided"
	ErrCookieMissmatchCode = "no_cookie_matches"
	ErrConvertionCode      = "convertion_error"
	ErrServerCode          = "something_went_wrong"
	ErrGenerateTokenCode   = "auth_token_generation_error"
	ErrCreateUserCode      = "create_user_error"
	ErrUnauthorizedCode    = "user_unauthorized"
	ErrRedisClearCode      = "failed_delete_from_redis"
	ErrRedisWriteCode      = "failed_write_into_redis"
	ErrFileUploadCode      = "file_upload_err"
	ErrUpdatePasswordCode  = "update_password_error"
	ErrUpdateProfileCode   = "update_profile_error"
	ErrInvalidEmailCode    = "invalid_email"
	ErrInvalidPasswordCode = "invalid_password"
	ErrUserNotFoundCode    = "user_not_found"
	ErrInvalidUsernameCode = "invalid_username"
)

var (
	ErrInvalidEmailText          = CustomError{Err: errors.New("email is incorrect")}
	ErrInvalidPasswordText       = CustomError{Err: errors.New("password is too short. The minimal len is 8")}
	ErrInvalidUsernameText       = CustomError{Err: errors.New("username is too short. The minimal len is 6")}
	ErrInvalidPasswordsMatchText = CustomError{Err: errors.New("password doesnt match with passwordConfirmation")}
	ErrInvalidOldPasswordText    = CustomError{Err: errors.New("invalid old password")}
	ErrUserNotFoundText          = CustomError{Err: errors.New("cannot find user by given params")}
	ErrBrokenCookieText          = CustomError{Err: errors.New("broken cookie was given")}
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
