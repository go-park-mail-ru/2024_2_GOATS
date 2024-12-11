package validation

import (
	"regexp"

	errVals "github.com/go-park-mail-ru/2024_2_GOATS/internal/app/errors"
)

const (
	passwordLength = 8
	usernameLength = 6
)

var emailRegex = regexp.MustCompile(`^[$\/@ "'.!#\$%&'*+\-=?^_{|}~a-zA-z0-9]+@[a-z]+\.[a-z]{2,10}$`)

// ValidatePassword validates password len and matching with password_confirmation
func ValidatePassword(pass, passConf string) *errVals.CustomError {
	if pass != passConf {
		return &errVals.ErrInvalidPasswordsMatch
	}

	if len(pass) < passwordLength {
		return &errVals.ErrInvalidPassword
	}

	return nil
}

// ValidateEmail validates email matches regexp
func ValidateEmail(email string) *errVals.CustomError {
	if !emailRegex.MatchString(email) {
		return &errVals.ErrInvalidEmail
	}

	return nil
}

// ValidateUsername validates username length
func ValidateUsername(username string) *errVals.CustomError {
	if len(username) < usernameLength {
		return &errVals.ErrInvalidUsername
	}

	return nil
}

// ValidateCookie validates cookie presence
func ValidateCookie(cookie string) *errVals.CustomError {
	if len(cookie) == 0 {
		return &errVals.ErrBrokenCookie
	}

	return nil
}
