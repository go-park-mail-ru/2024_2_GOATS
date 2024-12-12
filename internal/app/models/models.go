package models

import (
	"database/sql"
	"fmt"
	"mime/multipart"
	"strings"
	"time"
)

// LoginData represents user login data
type LoginData struct {
	Email    string
	Password string
	Cookie   string
}

// RegisterData represents user register data
type RegisterData struct {
	Email                string
	Username             string
	Password             string
	PasswordConfirmation string
}

// SessionRespData represents user session data
type SessionRespData struct {
	UserData User
}

// AuthRespData represents cookie creation response
type AuthRespData struct {
	NewCookie *CookieData
}

// CollectionsRespData represents collections response
type CollectionsRespData struct {
	Collections []Collection
}

// User represents user full info
type User struct {
	ID                         int
	Email                      string
	Username                   string
	Password                   string
	AvatarURL                  string
	AvatarName                 string
	AvatarFile                 multipart.File
	SubscriptionStatus         bool
	SubscriptionExpirationDate string
}

// DBEpisode represents db episode data
type DBEpisode struct {
	ID            sql.NullInt64
	Title         sql.NullString
	Description   sql.NullString
	EpisodeNumber sql.NullInt64
	ReleaseDate   sql.NullString
	Rating        sql.NullFloat64
	PreviewURL    sql.NullString
	VideoURL      sql.NullString
}

// DirectorInfo represents director full info
type DirectorInfo struct {
	Person
	ID int
}

// CookieData represents cookie full info
type CookieData struct {
	Name  string
	Token *Token
}

// Token represents cookie token full info
type Token struct {
	UserID  int
	TokenID string
	Expiry  time.Time
}

// PasswordData represents update_password data
type PasswordData struct {
	UserID               int
	OldPassword          string
	Password             string
	PasswordConfirmation string
}

// Person represents person's name and surname
type Person struct {
	Name    string
	Surname string
}

// Favorite represents favorite's relations
type Favorite struct {
	UserID  int
	MovieID int
}

// SubscriptionData represents base subscription params
type SubscriptionData struct {
	UserID int
	Amount uint64
}

// PaymentCallbackData contains YooMoney payment_callback data
type PaymentCallbackData struct {
	NotificationType string
	OperationID      string
	Amount           int64
	Currency         string
	Sender           string
	Label            string
	Unaccepted       bool
}

// CreatePaymentData represents data for payment creation
type CreatePaymentData struct {
	SubscriptionID int
	Amount         uint64
}

// FullName returns the person's fullname
func (p Person) FullName() string {
	return strings.TrimSpace(fmt.Sprintf("%s %s", p.Name, p.Surname))
}
