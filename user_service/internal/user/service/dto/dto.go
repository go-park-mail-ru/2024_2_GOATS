package dto

type User struct {
	ID         uint64
	Email      string
	Username   string
	Password   string
	AvatarURL  string
	AvatarName string
	AvatarFile []byte
}

type Favorite struct {
	UserID  uint64
	MovieID uint64
}

type PasswordData struct {
	UserID               uint64
	OldPassword          string
	Password             string
	PasswordConfirmation string
}

type CreateUserData struct {
	Email                string
	Username             string
	Password             string
	PasswordConfirmation string
}
