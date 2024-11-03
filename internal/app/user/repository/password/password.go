package password

import (
	"fmt"

	"github.com/rs/zerolog/log"
	"golang.org/x/crypto/bcrypt"
)

func HashAndSalt(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		errMsg := fmt.Errorf("failed to hash password: %w", err)
		log.With().Timestamp().Err(errMsg)

		return "", errMsg
	}

	return string(hashedPassword), nil
}
