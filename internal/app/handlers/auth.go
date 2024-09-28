package handlers

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/go-park-mail-ru/2024_2_GOATS/internal/app/api"
	"github.com/go-park-mail-ru/2024_2_GOATS/internal/app/models"
	authModels "github.com/go-park-mail-ru/2024_2_GOATS/internal/app/models/auth"
	"github.com/go-park-mail-ru/2024_2_GOATS/internal/app/models/cookie"
	"github.com/labstack/gommon/log"
)

type AuthHandler struct {
	ApiLayer *api.Implementation
}

func NewAuthHandler(api *api.Implementation) *AuthHandler {
	return &AuthHandler{
		ApiLayer: api,
	}
}

func (a *AuthHandler) Logout(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ck, err := r.Cookie("session_id")
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}

		cs, err := cookie.NewCookieStore(a.ApiLayer.Ctx)
		if err != nil {
			log.Errorf("failed to connect to Redis: %w", err)
			http.Error(w, "Redis Server Error", http.StatusInternalServerError)
			return
		}

		defer func() {
			if err := cs.RedisDB.Close(); err != nil {
				log.Fatal("Error closing redis connection %w", err)
			}
		}()

		err = cs.DeleteCookie(ck.Value)
		if err != nil {
			log.Errorf("cookie error: %w", err)
			http.Error(w, "Redis Server Error", http.StatusInternalServerError)
			return
		}

		err = json.NewEncoder(w).Encode(map[string]bool{"success": true})
		if err != nil {
			log.Errorf("error while encoding success logout response: %w", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})
}

func (a *AuthHandler) Login(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		loginData := &authModels.LoginData{}
		a.handleAuth(w, r, loginData, "login")
	})
}

func (a *AuthHandler) Register(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		registerData := &authModels.RegisterData{}
		a.handleAuth(w, r, registerData, "register")
	})
}

func (a *AuthHandler) Session(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ck, err := r.Cookie("session_id")
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}

		sessionResp, errResp := a.ApiLayer.Session(a.ApiLayer.Ctx, ck.Value)
		if errResp != nil {
			w.WriteHeader(errResp.StatusCode)
			err = json.NewEncoder(w).Encode(errResp)
			if err != nil {
				log.Errorf("error while encoding bad session response: %w", err)
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
			return
		}

		err = json.NewEncoder(w).Encode(sessionResp)
		if err != nil {
			log.Errorf("error while encoding success session response: %w", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})
}

func (a *AuthHandler) handleAuth(w http.ResponseWriter, r *http.Request, decodeData interface{}, operation string) {
	err := json.NewDecoder(r.Body).Decode(decodeData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var authResp *authModels.AuthResponse
	var errResp *models.ErrorResponse

	if operation == "login" {
		authResp, errResp = a.ApiLayer.Login(a.ApiLayer.Ctx, decodeData.(*authModels.LoginData))
	} else if operation == "register" {
		authResp, errResp = a.ApiLayer.Register(a.ApiLayer.Ctx, decodeData.(*authModels.RegisterData))
	}

	if errResp != nil {
		w.WriteHeader(errResp.StatusCode)
		err = json.NewEncoder(w).Encode(errResp)
		if err != nil {
			log.Errorf("error while encoding bad auth response: %w", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		return
	}

	cookieProcessor(a.ApiLayer.Ctx, w, authResp.Token)

	err = json.NewEncoder(w).Encode(authResp)
	if err != nil {
		log.Errorf("error while encoding success auth response: %w", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func cookieProcessor(ctx context.Context, w http.ResponseWriter, token *authModels.Token) {
	cs, err := cookie.NewCookieStore(ctx)
	if err != nil {
		log.Errorf("failed to connect to Redis: %w", err)
		http.Error(w, "Redis Server Error", http.StatusInternalServerError)
		return
	}

	defer func() {
		if err := cs.RedisDB.Close(); err != nil {
			log.Fatal("Error closing redis connection %w", err)
		}
	}()

	err = cs.DeleteCookie(token.TokenID)
	if err != nil {
		log.Errorf("cookie error: %w", err)
		http.Error(w, "Redis Server Error", http.StatusInternalServerError)
		return
	}

	sessionCookie, err := cs.SetCookie(token)
	if err != nil {
		log.Errorf("cookie error: %w", err)
		http.Error(w, "Failed to set cookie", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Set-Cookie", sessionCookie)
}
