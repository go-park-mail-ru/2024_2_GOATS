package validation

import (
	"github.com/go-park-mail-ru/2024_2_GOATS/auth_service/internal/errs"
)

// ValidateCookie validates cookie presence
func ValidateCookie(cookie string) error {
	if cookie == "" {
		return errs.ErrInvalidCookie
	}

	return nil
}

// ValidateUserID validates userID presense
func ValidateUserID(usrID uint64) error {
	if usrID == 0 {
		return errs.ErrInvalidUserID
	}

	return nil
}
