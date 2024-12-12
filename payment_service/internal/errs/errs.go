package errs

import "errors"

var (
	// ErrBadRequest represents bad_request error
	ErrBadRequest = errors.New("bad_request")
)
