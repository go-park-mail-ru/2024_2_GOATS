package validation

import (
	"log"
	"regexp"

	"github.com/go-park-mail-ru/2024_2_GOATS/internal/app/errors"
	errVals "github.com/go-park-mail-ru/2024_2_GOATS/internal/app/errors"
)

const PasswordLength = 8

var emailRegex = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")

func ValidatePassword(pass, passConf string) *errors.CustomError {
	if pass != passConf {
		log.Println(errVals.ErrInvalidPasswordsMatchText.Err.Error())
		return &errVals.ErrInvalidPasswordsMatchText
	}

	if len(pass) < PasswordLength {
		log.Println(errVals.ErrInvalidPasswordText.Err.Error())
		return &errVals.ErrInvalidPasswordText
	}

	return nil
}

func ValidateEmail(email string) *errors.CustomError {
	if !emailRegex.MatchString(email) {
		log.Println(errVals.ErrInvalidEmailText.Err.Error())
		return &errVals.ErrInvalidEmailText
	}

	return nil
}

func ValidateCookie(cookie string) *errors.CustomError {
	if len(cookie) == 0 {
		return &errVals.ErrBrokenCookieText
	}

	return nil
}
