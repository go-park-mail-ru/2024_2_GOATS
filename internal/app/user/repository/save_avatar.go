package repository

import (
	"context"
	"fmt"
	"io"
	"os"

	errVals "github.com/go-park-mail-ru/2024_2_GOATS/internal/app/errors"
	"github.com/go-park-mail-ru/2024_2_GOATS/internal/app/models"
	"github.com/rs/zerolog/log"
)

func (u *UserRepo) SaveAvatar(ctx context.Context, usrData *models.User) (string, *errVals.ErrorObj) {
	path := fmt.Sprintf("images/user_avatars/%s", usrData.AvatarName)
	outFile, osErr := os.Create(path)
	if osErr != nil {
		return "", &errVals.ErrorObj{
			Code:  "file_upload_err",
			Error: errVals.CustomError{Err: fmt.Errorf("cannot find or create nginx static folder: %w", osErr)},
		}
	}

	defer func() {
		if err := outFile.Close(); err != nil {
			log.Ctx(ctx).Err(fmt.Errorf("failed to close outFile: %w", err))
		}
	}()

	_, osErr = io.Copy(outFile, usrData.Avatar)
	if osErr != nil {
		return "", &errVals.ErrorObj{
			Code:  "file_upload_err",
			Error: errVals.CustomError{Err: fmt.Errorf("cannot save file into nginx static folder: %w", osErr)},
		}
	}

	return path, nil
}
