package password

import (
	"context"
	"fmt"

	"github.com/rs/zerolog/log"
	"golang.org/x/crypto/bcrypt"
)

func HashAndSalt(ctx context.Context, password string) (string, error) {
	lg := log.Ctx(ctx)
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		errMsg := fmt.Errorf("failed to hash password: %w", err)
		lg.Error().Msg(errMsg.Error())

		return "", errMsg
	}

	return string(hashedPassword), nil
}
