package auth

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/go-park-mail-ru/2024_2_GOATS/internal/app/api"
	authModels "github.com/go-park-mail-ru/2024_2_GOATS/internal/app/models/auth"
	"github.com/go-park-mail-ru/2024_2_GOATS/internal/app/models/cookie"
)

type AuthHandler struct {
	Context  context.Context
	ApiLayer *api.Implementation
}

func NewHandler(ctx context.Context, api *api.Implementation) *AuthHandler {
	return &AuthHandler{
		Context: ctx,
		ApiLayer: api,
	}
}

func (a *AuthHandler) Login(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		loginData := &authModels.LoginData{}
		err := json.NewDecoder(r.Body).Decode(loginData)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		loginResp, errResp := a.ApiLayer.Login(a.Context, loginData)
		if errResp != nil {
			w.WriteHeader(errResp.StatusCode)
			json.NewEncoder(w).Encode(errResp)
			return
		}

		cookieProcessor(a.Context, w, loginResp.Token)

		json.NewEncoder(w).Encode(loginResp)
	})
}

func (a *AuthHandler) Register(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		registerData := &authModels.RegisterData{}
		err := json.NewDecoder(r.Body).Decode(&registerData)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		regResp, errResp := a.ApiLayer.Register(a.Context, registerData)
		if errResp != nil {
			w.WriteHeader(errResp.StatusCode)
			json.NewEncoder(w).Encode(errResp)
			return
		}

		cookieProcessor(a.Context, w, regResp.Token)

		json.NewEncoder(w).Encode(regResp)
	})
}

func (a *AuthHandler) Session(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		cookie, err := r.Cookie("session_id")
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		sessionResp, errResp := a.ApiLayer.Session(a.Context, cookie.Value)
		if errResp != nil {
			w.WriteHeader(errResp.StatusCode)
			json.NewEncoder(w).Encode(errResp)
			return
		}

		json.NewEncoder(w).Encode(sessionResp)
	})
}

func cookieProcessor(ctx context.Context, w http.ResponseWriter, token *authModels.Token) {
	cookieStore, err := cookie.NewCookieStore(ctx)
	if err != nil {
		http.Error(w, "Failed to connect to Redis", http.StatusInternalServerError)
		return
	}

	defer cookieStore.RedisDB.Close()

	cookieStore.DeleteCookie(ctx, token.UserID)

	defer cookieStore.RedisDB.Close()

	sessionCookie, err := cookieStore.SetCookie(ctx, token)
	if err != nil {
		http.Error(w, "Failed to set cookie", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Set-Cookie", sessionCookie)
}
