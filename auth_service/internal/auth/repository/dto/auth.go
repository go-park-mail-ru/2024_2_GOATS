package dto

import "time"

type CookieData struct {
	Name  string
	Token *TokenData
}

type TokenData struct {
	UserID  uint64
	TokenID string
	Expiry  time.Time
}
