package api

import (
	errVals "github.com/go-park-mail-ru/2024_2_GOATS/internal/app/errors"
	"github.com/go-park-mail-ru/2024_2_GOATS/internal/app/models"
)

type RegisterRequest struct {
	Email                string `json:"email"`
	Username             string `json:"username"`
	Password             string `json:"password"`
	PasswordConfirmation string `json:"passwordConfirmation"`
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Cookie   string `json:"-"`
}
type SessionResponse struct {
	Success    bool `json:"success"`
	UserData   User `json:"user_data"`
	StatusCode int  `json:"-"`
}

type AuthResponse struct {
	Success    bool               `json:"success"`
	NewCookie  *models.CookieData `json:"-"`
	StatusCode int                `json:"-"`
}

type CollectionsResponse struct {
	Success     bool                `json:"success"`
	Collections []models.Collection `json:"collections"`
	StatusCode  int                 `json:"-"`
}

type ErrorResponse struct {
	Success    bool               `json:"success"`
	StatusCode int                `json:"-"`
	Errors     []errVals.ErrorObj `json:"errors"`
}

type User struct {
	Id       int    `json:"id"`
	Email    string `json:"email"`
	Username string `json:"username"`
}
