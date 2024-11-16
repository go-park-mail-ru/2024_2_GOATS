package dto

import "time"

type SrvCreateCookie struct {
	UserID uint64
}

type SrvDeleteCookie struct {
	Token string
}

type SrvCheckCookie struct {
	Token string
}

type SrvSuccessResp struct {
	Success bool
}

type Cookie struct {
	Name    string
	UserID  uint64
	TokenID string
	Expiry  time.Time
}

type Token struct {
	UserID  uint64
	TokenID string
	Expiry  time.Time
}
