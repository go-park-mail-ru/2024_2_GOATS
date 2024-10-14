package cookie

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"time"

	"github.com/go-park-mail-ru/2024_2_GOATS/config"
	"github.com/go-park-mail-ru/2024_2_GOATS/internal/app/models"
)

func GenerateToken(ctx context.Context, userID int) (*models.Token, error) {
	tokenID, err := generateRandomString(32)
	if err != nil {
		return nil, fmt.Errorf("failed to generate cookie token: %w", err)
	}

	expiry := time.Now().Add(config.FromContext(ctx).Databases.Redis.Cookie.MaxAge)

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
		return "", fmt.Errorf("failed to generate random string: %w", err)
	}

	return base64.URLEncoding.EncodeToString(bytes), nil
}
