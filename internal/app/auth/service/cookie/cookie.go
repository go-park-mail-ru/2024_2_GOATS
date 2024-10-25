package cookie

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"time"

	"github.com/go-park-mail-ru/2024_2_GOATS/config"
	"github.com/go-park-mail-ru/2024_2_GOATS/internal/app/models"
	"github.com/rs/zerolog/log"
)

func GenerateToken(ctx context.Context, userID int) (*models.Token, error) {
	tokenID, err := generateRandomString(32)
	if err != nil {
		errMsg := fmt.Errorf("cookie: failed to generate cookie token - %w", err)
		log.Ctx(ctx).Error().Msg(errMsg.Error())

		return nil, errMsg
	}

	expiry := time.Now().Add(config.FromRedisContext(ctx).Cookie.MaxAge)

	return &models.Token{
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
		log.Error().Msg(errMsg.Error())

		return "", errMsg
	}

	return base64.URLEncoding.EncodeToString(bytes), nil
}
