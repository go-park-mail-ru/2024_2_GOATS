package cookie

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"net/http"
	"time"

	"github.com/go-park-mail-ru/2024_2_GOATS/config"
	authModels "github.com/go-park-mail-ru/2024_2_GOATS/internal/app/models/auth"
	"github.com/go-redis/redis/v8"
)

type Store struct {
	RedisDB  *redis.Client
	RedisCfg config.Redis
}

func NewCookieStore(ctx context.Context) (*Store, error) {
	cfg := config.FromContext(ctx)
	addr := fmt.Sprintf("%s:%d", cfg.Databases.Redis.Host, cfg.Databases.Redis.Port)
	rdb := redis.NewClient(&redis.Options{
		Addr: addr,
	})

	return &Store{
		RedisDB:  rdb,
		RedisCfg: cfg.Databases.Redis,
	}, nil
}

func (cs *Store) SetCookie(token *authModels.Token) (string, error) {
	err := cs.RedisDB.Set(context.Background(), token.TokenID, fmt.Sprint(token.UserID), cs.RedisCfg.Cookie.MaxAge).Err()
	if err != nil {
		return "", fmt.Errorf("cannot set cookie into redis: %w", err)
	}

	cookie := &http.Cookie{
		Name:     cs.RedisCfg.Cookie.Name,
		Value:    token.TokenID,
		Expires:  token.Expiry,
		HttpOnly: true,
		SameSite: http.SameSiteLaxMode,
		Secure:   false,
	}

	return cookie.String(), nil
}

func (cs *Store) GetFromCookie(cookie string) (string, error) {
	var userID string
	err := cs.RedisDB.Get(context.Background(), cookie).Scan(&userID)
	if err != nil {
		return "", fmt.Errorf("cannot get cookie from redis: %w", err)
	}

	return userID, nil
}

func (cs *Store) DeleteCookie(token string) error {
	_, err := cs.RedisDB.Del(context.Background(), token).Result()

	if err != nil {
		return fmt.Errorf("failed to delete old cookie: %w", err)
	}

	return nil
}

func GenerateToken(ctx context.Context, userID int) (*authModels.Token, error) {
	tokenID, err := generateRandomString(32)
	if err != nil {
		return nil, fmt.Errorf("failed to generate cookie token: %w", err)
	}

	expiry := time.Now().Add(config.FromContext(ctx).Databases.Redis.Cookie.MaxAge)

	return &authModels.Token{
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
