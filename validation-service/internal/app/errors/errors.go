package errors

import "errors"

var (
	ErrInvalidEmailCode     = "invalid_email"
	ErrInvalidEmailText     = errors.New("email is incorrect")
	ErrInvalidPasswordCode  = "invalid_password"
	ErrInvalidPasswordText  = errors.New("password is too short. The minimal len is 8")
	ErrInvalidSexCode       = "invalid_sex"
	ErrInvalidSexText       = errors.New("only male or female allowed")
	ErrInvalidBirthdateCode = "invalid_birthdate"
	ErrInvalidBirthdateText = errors.New("bithdate should be before current time")
)
