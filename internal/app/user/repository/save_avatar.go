package repository

import (
	"context"
	"fmt"
	"os"

	"github.com/go-park-mail-ru/2024_2_GOATS/config"
	errVals "github.com/go-park-mail-ru/2024_2_GOATS/internal/app/errors"
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

	return relativePath, outFile, nil
}
