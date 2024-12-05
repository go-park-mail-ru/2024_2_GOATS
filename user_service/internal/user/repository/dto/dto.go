package dto

import (
	"database/sql"
)

type RepoCreateData struct {
	Email                string
	Username             string
	Password             string
	PasswordConfirmation string
}

type RepoCreateSubscriptionData struct {
	UserID uint64
	Amount uint64
}

type RepoUser struct {
	ID        uint64
	Email     string
	Username  string
	Password  string
	AvatarURL string
}

type RepoFavorite struct {
	UserID  uint64
	MovieID uint64
}

type RepoSubscription struct {
	Status         sql.NullString
	ExpirationDate sql.NullTime
}
