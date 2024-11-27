package errs

import "errors"

var (
	ErrBadRequest    = errors.New("bad_request")
	ErrInvalidCookie = errors.New("empty_cookie")
	ErrInvalidUserID = errors.New("invalid_user_id")
)
