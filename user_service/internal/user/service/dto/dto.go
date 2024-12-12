package dto

// User represents full user info
type User struct {
	ID                         uint64
	Email                      string
	Username                   string
	Password                   string
	AvatarURL                  string
	AvatarName                 string
	AvatarFile                 []byte
	SubscriptionStatus         bool
	SubscriptionExpirationDate string
}

// Favorite contains user and movie relation
type Favorite struct {
	UserID  uint64
	MovieID uint64
}

// PasswordData for update user password
type PasswordData struct {
	UserID               uint64
	OldPassword          string
	Password             string
	PasswordConfirmation string
}

// CreateUserData for create_user action
type CreateUserData struct {
	Email                string
	Username             string
	Password             string
	PasswordConfirmation string
}

// CreateSubscriptionData for create_subscription action
type CreateSubscriptionData struct {
	UserID uint64
	Amount uint64
}
