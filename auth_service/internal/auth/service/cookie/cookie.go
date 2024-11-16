package cookie

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"time"

	"github.com/go-park-mail-ru/2024_2_GOATS/auth_service/internal/auth/service/dto"
	"github.com/go-park-mail-ru/2024_2_GOATS/auth_service/config"
	"github.com/rs/zerolog/log"
)

func GenerateToken(ctx context.Context, userID uint64) (*dto.Token, error) {
	tokenID, err := generateRandomString(32)
	if err != nil {
		errMsg := fmt.Errorf("cookie: failed to generate cookie token - %w", err)
		log.Ctx(ctx).Error().Err(errMsg).Msg("token_generation_error")

		return nil, errMsg
	}

	expiry := time.Now().Add(config.FromRedisContext(ctx).Cookie.MaxAge)

	return &dto.Token{
		UserID:  userID,
		TokenID: tokenID,
		Expiry:  expiry,
	}, nil
}

func generateRandomString(length int) (string, error) {
	bytes := make([]byte, length)
	_, err := rand.Read(bytes)
	if err != nil {
		errMsg := fmt.Errorf("cookie: failed to generate random string - %w", err)
		log.Error().Err(errMsg).Msg("gen_random_string_error")

		return "", errMsg
	}

	return base64.URLEncoding.EncodeToString(bytes), nil
}
