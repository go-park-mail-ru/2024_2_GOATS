package errs

import "errors"

var (
	// ErrBadRequest represents bad_request error
	ErrBadRequest = errors.New("bad_request")
	// ErrInvalidCookie represents unauthorized request error
	ErrInvalidCookie = errors.New("empty_cookie")
	// ErrInvalidUserID represents invalid user_id error
	ErrInvalidUserID = errors.New("invalid_user_id")
)
