package dto

import "time"

// CookieData repo layer struct
type CookieData struct {
	Name  string
	Token *TokenData
}

// TokenData repo layer struct
type TokenData struct {
	UserID  uint64
	TokenID string
	Expiry  time.Time
}
