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

type CookieStore struct {
	RedisDB *redis.Client
}

func NewCookieStore(ctx context.Context) (*CookieStore, error) {
	// addr := fmt.Sprintf("%s:6379", "localhost") local
	host := config.FromContext(ctx).Redis.Host
	port := config.FromContext(ctx).Redis.Port
	addr := fmt.Sprintf("%s:%d", host, port)
	rdb := redis.NewClient(&redis.Options{
		Addr: addr,
	})
	return &CookieStore{RedisDB: rdb}, nil
}

func (cs *CookieStore) SetCookie(ctx context.Context, token *authModels.Token) (string, error) {
	err := cs.RedisDB.Set(ctx, token.TokenID, fmt.Sprint(token.UserID), config.FromContext(ctx).Redis.Cookie.MaxAge).Err()
	if err != nil {
		return "", err
	}

	cookie := &http.Cookie{
		Name:     config.FromContext(ctx).Redis.Cookie.Name,
		Value:    token.TokenID,
		Expires:  token.Expiry,
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteStrictMode,
	}

	return cookie.String(), nil
}

func (cs *CookieStore) GetFromCookie(ctx context.Context, cookie string) (string, error) {
	var userID string
	err := cs.RedisDB.Get(ctx, cookie).Scan(&userID)
	if err != nil {
		return "", err
	}

	return userID, nil
}

func (cs *CookieStore) DeleteCookie(ctx context.Context, userId int) error {
	_, err := cs.RedisDB.Del(ctx, fmt.Sprint(userId)).Result()
	return err
}

func GenerateToken(ctx context.Context, userID int) (*authModels.Token, error) {
	tokenID := generateRandomString(32)

	expiry := time.Now().Add(config.FromContext(ctx).Redis.Cookie.MaxAge)

	return &authModels.Token{
		UserID:  userID,
		TokenID: tokenID,
		Expiry:  expiry,
	}, nil
}

func generateRandomString(length int) string {
	bytes := make([]byte, length)
	_, err := rand.Read(bytes)
	if err != nil {
		panic(err)
	}
	return base64.URLEncoding.EncodeToString(bytes)
}
