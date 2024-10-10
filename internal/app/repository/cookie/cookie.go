package cookie

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"time"

	"github.com/go-park-mail-ru/2024_2_GOATS/config"
	authModels "github.com/go-park-mail-ru/2024_2_GOATS/internal/app/models/auth"
	"github.com/go-redis/redis/v8"
)

type Store struct {
	RedisDB  *redis.Client
	RedisCfg config.Redis
	Ctx      context.Context
}

func NewCookieStore(ctx context.Context, rdb *redis.Client) *Store {
	cfg := config.FromContext(ctx)

	return &Store{
		RedisDB:  rdb,
		RedisCfg: cfg.Databases.Redis,
		Ctx:      ctx,
	}
}

func (cs *Store) SetCookie(token *authModels.Token) (*authModels.CookieData, error) {
	err := cs.RedisDB.Set(cs.Ctx, token.TokenID, fmt.Sprint(token.UserID), cs.RedisCfg.Cookie.MaxAge).Err()
	if err != nil {
		return nil, fmt.Errorf("cannot set cookie into redis: %w", err)
	}

	return &authModels.CookieData{
		Name:   cs.RedisCfg.Cookie.Name,
		Value:  token.TokenID,
		Expiry: token.Expiry,
		UserID: token.UserID,
	}, nil
}

func (cs *Store) GetFromCookie(cookie string) (string, error) {
	var userID string
	err := cs.RedisDB.Get(cs.Ctx, cookie).Scan(&userID)
	if err != nil {
		return "", fmt.Errorf("cannot get cookie from redis: %w", err)
	}

	return userID, nil
}

func (cs *Store) DeleteCookie(token string) (*authModels.CookieData, error) {
	_, err := cs.RedisDB.Del(cs.Ctx, token).Result()

	if err != nil {
		return nil, fmt.Errorf("failed to delete old cookie: %w", err)
	}

	return &authModels.CookieData{
		Name:   cs.RedisCfg.Cookie.Name,
		Value:  "",
		Expiry: time.Unix(0, 0),
	}, nil
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
