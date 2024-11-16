package repository

import (
	"context"
	"fmt"
	"os"

	"github.com/go-park-mail-ru/2024_2_GOATS/config"
	errVals "github.com/go-park-mail-ru/2024_2_GOATS/internal/app/errors"
	"github.com/rs/zerolog/log"
)

func (u *UserRepo) SaveUserAvatar(ctx context.Context, avatarName string) (string, *os.File, *errVals.RepoError) {
	lclStrg := config.FromLocalStorageContext(ctx)
	fullPath := lclStrg.UserAvatarsFullURL + avatarName
	relativePath := lclStrg.UserAvatarsRelativeURL + avatarName

	outFile, fileErr := os.Create(fullPath)
	if fileErr != nil {
		return "", nil, errVals.NewRepoError(
			errVals.ErrFileUploadCode,
			errVals.NewCustomError(fmt.Sprintf("cannot find or create nginx static folder: %v", fileErr)),
		)
	}

	defer func() {
		if err := outFile.Close(); err != nil {
			log.Ctx(ctx).Err(fmt.Errorf("failed to close outFile: %w", err))
		}
	}()

	return relativePath, outFile, nil
}
