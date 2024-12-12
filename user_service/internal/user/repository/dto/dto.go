package dto

import (
	"database/sql"
)

// RepoCreateData repo create_user struct
type RepoCreateData struct {
	Email                string
	Username             string
	Password             string
	PasswordConfirmation string
}

// RepoCreateSubscriptionData repo create_subscription struct
type RepoCreateSubscriptionData struct {
	UserID uint64
	Amount uint64
}

// RepoUser repo user_data struct
type RepoUser struct {
	ID        uint64
	Email     string
	Username  string
	Password  string
	AvatarURL string
}

// RepoFavorite repo favorite struct
type RepoFavorite struct {
	UserID  uint64
	MovieID uint64
}

// RepoSubscription repo subscription struct
type RepoSubscription struct {
	Status         sql.NullString
	ExpirationDate sql.NullTime
}
