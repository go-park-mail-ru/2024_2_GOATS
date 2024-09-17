package errors

var (
	ErrInvalidEmailCode     = "invalid_email"
	ErrInvalidEmailText     = "email is incorrect"
	ErrInvalidPasswordCode  = "invalid_password"
	ErrInvalidPasswordText  = "password is too short. The minimal len is 8"
	ErrInvalidSexCode       = "invalid_sex"
	ErrInvalidSexText       = "only male or female allowed"
	ErrInvalidBirthdateCode = "invalid_birthdate"
	ErrInvalidBirthdateText = "bithdate should be before current time"
)
