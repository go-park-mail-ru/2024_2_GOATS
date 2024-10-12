package repository

import (
	"context"
	"fmt"
	"net/http"

	"github.com/go-park-mail-ru/2024_2_GOATS/config"
	authModels "github.com/go-park-mail-ru/2024_2_GOATS/internal/app/models/auth"
)

func (r *Repo) SetCookie(ctx context.Context, token *authModels.Token) (*authModels.CookieData, error, int) {
	cookieCfg := config.FromContext(ctx).Databases.Redis.Cookie

	err := r.Redis.Set(ctx, token.TokenID, fmt.Sprint(token.UserID), cookieCfg.MaxAge)
	if err.Err() != nil {
		return nil, fmt.Errorf("cannot set cookie into redis: %w", err), http.StatusInternalServerError
	}

	return &authModels.CookieData{
		Name:   cookieCfg.Name,
		Value:  token.TokenID,
		Expiry: token.Expiry,
		UserID: token.UserID,
	}, nil, http.StatusOK
}
