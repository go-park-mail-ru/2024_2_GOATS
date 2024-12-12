package converter

import (
	"database/sql"

	"github.com/go-park-mail-ru/2024_2_GOATS/user_service/internal/user/repository/dto"
	"github.com/go-park-mail-ru/2024_2_GOATS/user_service/internal/user/repository/subscriptiondb"
	srvDTO "github.com/go-park-mail-ru/2024_2_GOATS/user_service/internal/user/service/dto"
)

// ToUserFromRepoUser converts from repo User to srv DTO User
func ToUserFromRepoUser(u *dto.RepoUser, s *dto.RepoSubscription) *srvDTO.User {
	if u == nil {
		return nil
	}

	return &srvDTO.User{
		ID:                         u.ID,
		Email:                      u.Email,
		Username:                   u.Username,
		Password:                   u.Password,
		AvatarURL:                  u.AvatarURL,
		SubscriptionStatus:         subscriptionStatus(s.Status),
		SubscriptionExpirationDate: toStringFromNullTime(s.ExpirationDate),
	}
}

// ToUserShortFromRepoUser converts from repo User to srv DTO User
func ToUserShortFromRepoUser(u *dto.RepoUser) *srvDTO.User {
	if u == nil {
		return nil
	}

	return &srvDTO.User{
		ID:        u.ID,
		Email:     u.Email,
		Username:  u.Username,
		Password:  u.Password,
		AvatarURL: u.AvatarURL,
	}
}

func toStringFromNullTime(t sql.NullTime) string {
	if t.Valid {
		return t.Time.Format("2006-01-02")
	}

	return ""
}

func subscriptionStatus(st sql.NullString) bool {
	if st.Valid {
		return st.String == subscriptiondb.ActiveStatus
	}

	return false
}
