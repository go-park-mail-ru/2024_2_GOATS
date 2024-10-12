package auth

import (
	"time"

	"github.com/go-park-mail-ru/2024_2_GOATS/internal/app/models"
)

type CookieData struct {
	Name   string
	Value  string
	UserID int
	Expiry time.Time
}

type RegisterData struct {
	Email                string `json:"email"`
	Username             string `json:"username"`
	Password             string `json:"password"`
	PasswordConfirmation string `json:"passwordConfirmation"`
}

type LoginData struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Cookie   string `json:"-"`
}
type SessionResponse struct {
	Success    bool        `json:"success"`
	UserData   models.User `json:"user_data"`
	StatusCode int         `json:"-"`
}

type AuthResponse struct {
	Success    bool        `json:"success"`
	NewCookie  *CookieData `json:"-"`
	StatusCode int         `json:"-"`
}

type Token struct {
	UserID  int
	TokenID string
	Expiry  time.Time
}
