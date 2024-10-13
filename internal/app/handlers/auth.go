package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/go-park-mail-ru/2024_2_GOATS/config"
	"github.com/go-park-mail-ru/2024_2_GOATS/internal/app/models"
	authModels "github.com/go-park-mail-ru/2024_2_GOATS/internal/app/models/auth"
)

type AuthHandler struct {
	ApiLayer AuthImplementationInterface
	Config   *config.Config
}

func NewAuthHandler(api AuthImplementationInterface, cfg *config.Config) *AuthHandler {
	return &AuthHandler{
		ApiLayer: api,
		Config:   cfg,
	}
}

func (a *AuthHandler) Logout(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ck, err := r.Cookie("session_id")
		if errors.Is(err, http.ErrNoCookie) {
			Response(w, http.StatusForbidden, err)
			return
		}

		ctx := config.WrapContext(r.Context(), a.Config)
		resp, errData := a.ApiLayer.Logout(ctx, ck.Value)
		if errData != nil {
			Response(w, errData.StatusCode, errData)
			return
		}

		http.SetCookie(w, preparedExpiredCookie())

		Response(w, resp.StatusCode, resp)
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
		if errors.Is(err, http.ErrNoCookie) {
			Response(w, http.StatusForbidden, err)
			return
		}

		ctx := config.WrapContext(r.Context(), a.Config)
		sessionResp, errResp := a.ApiLayer.Session(ctx, ck.Value)
		if errResp != nil {
			Response(w, errResp.StatusCode, errResp)
			return
		}

		Response(w, sessionResp.StatusCode, sessionResp)
	})
}

func (a *AuthHandler) handleAuth(w http.ResponseWriter, r *http.Request, decodeData interface{}, operation string) {
	err := json.NewDecoder(r.Body).Decode(decodeData)
	if err != nil {
		Response(w, http.StatusBadRequest, fmt.Errorf("cannot parse request: %w", err))
		return
	}

	ctx := config.WrapContext(r.Context(), a.Config)
	var authResp *authModels.AuthResponse
	var errResp *models.ErrorResponse

	if operation == "login" {
		authResp, errResp = a.ApiLayer.Login(ctx, decodeData.(*authModels.LoginData))
	} else if operation == "register" {
		authResp, errResp = a.ApiLayer.Register(ctx, decodeData.(*authModels.RegisterData))
	}

	if errResp != nil {
		Response(w, errResp.StatusCode, errResp)
		return
	}

	oldCookie, err := r.Cookie("session_id")
	if errors.Is(err, http.ErrNoCookie) {
		log.Printf("user dont have old cookie")
	}

	if oldCookie != nil && authResp.NewCookie.Value != oldCookie.Value {
		http.SetCookie(w, preparedExpiredCookie())
	}

	http.SetCookie(w, preparedCookie(authResp.NewCookie))

	Response(w, authResp.StatusCode, authResp)
}

func preparedCookie(ck *authModels.CookieData) *http.Cookie {
	return &http.Cookie{
		Name:     ck.Name,
		Value:    ck.Value,
		Expires:  ck.Expiry,
		HttpOnly: true,
		SameSite: http.SameSiteLaxMode,
		Secure:   false,
	}
}

func preparedExpiredCookie() *http.Cookie {
	return &http.Cookie{
		Name:     "session_id",
		Value:    "",
		Expires:  time.Unix(0, 0),
		HttpOnly: true,
		SameSite: http.SameSiteLaxMode,
		Secure:   false,
	}
}
