package validation

import (
	"errors"
	"fmt"
	"regexp"
	"strings"

	user "github.com/go-park-mail-ru/2024_2_GOATS/user_service/pkg/user_v1"
)

const (
	favoriteValidationKey     = "incorrect_favorite_params"
	emailValidationKey        = "incorrect_email"
	usernameValidationKey     = "incorrect_username"
	passwordValidationKey     = "incorrect_password"
	passwordConfValidationKey = "incorrect_password_confirmation"
	wrongUsrIDValidationKey   = "incorrect_user_id"
	passwordLength            = 8
	usernameLength            = 6
)

var emailRegex = regexp.MustCompile(`^[$\/@ "'.!#\$%&'*+\-=?^_{|}~a-zA-z0-9]+@[a-z]+\.[a-z]{2,10}$`)

// ValidateFavoriteRequest validates favorite req params
func ValidateFavoriteRequest(favReq *user.HandleFavorite) error {
	if favReq.MovieID == 0 || favReq.UserID == 0 {
		return errors.New(favoriteValidationKey)
	}

	return nil
}

// ValidateCreateUserRequest validates create_user req params
func ValidateCreateUserRequest(req *user.CreateUserRequest) error {
	var validationErrors []string

	if req.Email == "" {
		validationErrors = append(validationErrors, emailValidationKey)
	} else if !emailRegex.MatchString(req.Email) {
		validationErrors = append(validationErrors, "email has invalid format")
	}

	if req.Username == "" {
		validationErrors = append(validationErrors, usernameValidationKey)
	} else if len(req.Username) < usernameLength {
		validationErrors = append(validationErrors, fmt.Sprintf("username must be at least %d characters long", usernameLength))
	}

	passErrs := validatePassword(req.Password, req.PasswordConfirmation)
	if passErrs != nil {
		validationErrors = append(validationErrors, passErrs...)
	}

	if len(validationErrors) == 0 {
		return nil
	}

	return fmt.Errorf("validation errors: %s", strings.Join(validationErrors, "; "))
}

func validatePassword(passwd string, passwdConf string) []string {
	var validationErrors []string
	if passwd == "" {
		validationErrors = append(validationErrors, passwordValidationKey)
	} else if len(passwd) < passwordLength {
		validationErrors = append(validationErrors, fmt.Sprintf("password must be at least %d characters long", passwordLength))
	}

	if passwd != passwdConf {
		validationErrors = append(validationErrors, passwordConfValidationKey)
	}

	if len(validationErrors) == 0 {
		return nil
	}

	return validationErrors
}

// ValidateUpdatePasswordRequesr validates update_password req params
func ValidateUpdatePasswordRequesr(req *user.UpdatePasswordRequest) error {
	var validationErrors []string

	if req.UserID == 0 {
		validationErrors = append(validationErrors, wrongUsrIDValidationKey)
	}

	passwdErrs := validatePassword(req.Password, req.PasswordConfirmation)
	if passwdErrs != nil {
		validationErrors = append(validationErrors, passwdErrs...)
	}

	if len(validationErrors) == 0 {
		return nil
	}

	return fmt.Errorf("validation errors: %s", strings.Join(validationErrors, "; "))
}
