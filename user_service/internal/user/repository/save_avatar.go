package repository

import (
	"context"
	"fmt"
	"os"

	"github.com/go-park-mail-ru/2024_2_GOATS/user_service/config"
	"github.com/rs/zerolog/log"
)

// SaveUserAvatar saves user avatar on the machine
func (u *UserRepo) SaveUserAvatar(ctx context.Context, avatarName string, file []byte) (string, error) {
	logger := log.Ctx(ctx)
	lclStrg := config.FromLocalStorageContext(ctx)
	fullPath := lclStrg.UserAvatarsFullURL + avatarName
	relativePath := lclStrg.UserAvatarsRelativeURL + avatarName

	err := os.WriteFile(fullPath, file, 0644)
	if err != nil {
		errMsg := fmt.Errorf("failed_to_save_file: %w", err)
		logger.Error().Err(errMsg).Msg("cannot save file")
		return "", errMsg
	}

	log.Printf("File %s saved successfully", fullPath)
	return relativePath, nil
}
