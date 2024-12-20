package repository

import (
	"context"

	"github.com/go-park-mail-ru/2024_2_GOATS/user_service/config"
)

// SaveUserAvatar creates path for storing in cloud
func (u *UserRepo) SaveUserAvatar(ctx context.Context, avatarName string) (string, error) {
	lclStrg := config.FromLocalStorageContext(ctx)
	relativePath := lclStrg.UserAvatarsRelativeURL + avatarName

	return relativePath, nil
}
