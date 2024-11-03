package validation

import (
	"regexp"

	"github.com/go-park-mail-ru/2024_2_GOATS/internal/app/errors"
	errVals "github.com/go-park-mail-ru/2024_2_GOATS/internal/app/errors"
)

const PasswordLength = 6
const UsernameLength = 8

var emailRegex = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")

func ValidatePassword(pass, passConf string) *errors.CustomError {
	if pass != passConf {
		return &errVals.ErrInvalidPasswordsMatchText
	}

	if len(pass) < PasswordLength {
		return &errVals.ErrInvalidPasswordText
	}

	return nil
}

func ValidateEmail(email string) *errors.CustomError {
	if !emailRegex.MatchString(email) {
		return &errVals.ErrInvalidEmailText
	}

	return nil
}

func ValidateUsername(username string) *errors.CustomError {
	if len(username) < UsernameLength {
		return &errVals.ErrInvalidUsernameText
	}

	return nil
}

func ValidateCookie(cookie string) *errors.CustomError {
	if len(cookie) == 0 {
		return &errVals.ErrBrokenCookieText
	}

	return nil
}
