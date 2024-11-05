package validation

import (
	"regexp"

	errVals "github.com/go-park-mail-ru/2024_2_GOATS/internal/app/errors"
)

const (
	passwordLength = 8
	usernameLength = 6
)

var emailRegex = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")

func ValidatePassword(pass, passConf string) *errVals.CustomError {
	if pass != passConf {
		return &errVals.ErrInvalidPasswordsMatchText
	}

	if len(pass) < passwordLength {
		return &errVals.ErrInvalidPasswordText
	}

	return nil
}

func ValidateEmail(email string) *errVals.CustomError {
	if !emailRegex.MatchString(email) {
		return &errVals.ErrInvalidEmailText
	}

	return nil
}

func ValidateUsername(username string) *errVals.CustomError {
	if len(username) < usernameLength {
		return &errVals.ErrInvalidUsernameText
	}

	return nil
}

func ValidateCookie(cookie string) *errVals.CustomError {
	if len(cookie) == 0 {
		return &errVals.ErrBrokenCookieText
	}

	return nil
}
