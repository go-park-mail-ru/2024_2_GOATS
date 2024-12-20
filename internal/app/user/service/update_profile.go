package service

import (
	"context"
	"crypto/tls"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"strings"

	errVals "github.com/go-park-mail-ru/2024_2_GOATS/internal/app/errors"
	"github.com/go-park-mail-ru/2024_2_GOATS/internal/app/models"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
)

// UpdateProfile updates user profile by calling userClient UpdateProfile
func (u *UserService) UpdateProfile(ctx context.Context, handler *multipart.FileHeader, usrData *models.User) *errVals.ServiceError {
	logger := log.Ctx(ctx)
	err := u.userClient.UpdateProfile(ctx, usrData)
	if err != nil {
		if strings.Contains(err.Error(), errVals.DuplicateErrCode) {
			logger.Error().Err(err).Msg("duplicate entry")
			return errVals.NewServiceError(errVals.DuplicateErrCode, fmt.Errorf("failed to register: %w", err))
		}

		logger.Error().Err(err).Msg("something went wrong")
		return errVals.NewServiceError("internal_error", fmt.Errorf("failed to register: %w", err))
	}

	if handler == nil {
		logger.Info().Msg("no avatar provided")
		return nil
	}

	useSSL := true

	bucketName := viper.GetString("VK_CLOUD_BUCKET_NAME")
	transport := &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true,
		},
	}

	minioClient, err := minio.New(viper.GetString("VK_CLOUD_ENDPOINT"), &minio.Options{
		Transport: transport,
		Creds:     credentials.NewStaticV4(viper.GetString("VK_CLOUD_ACCESS_KEY"), viper.GetString("VK_CLOUD_SECRET"), ""),
		Secure:    useSSL,
	})

	if err != nil {
		logger.Error().Err(err).Msg("failed to initialize minioClient")
		return errVals.NewServiceError("cannot_upload_file", err)
	}

	err = uploadFile(ctx, minioClient, bucketName, usrData.AvatarName, usrData.AvatarFile, handler.Size)
	if err != nil {
		logger.Error().Err(err).Msg("failed to upload file")
		return errVals.NewServiceError("cannot_upload_file", err)
	}

	return nil
}

func uploadFile(ctx context.Context, minioClient *minio.Client, bucketName string, objectName string, file multipart.File, filesize int64) error {
	logger := log.Ctx(ctx)
	err := minioClient.MakeBucket(ctx, bucketName, minio.MakeBucketOptions{})
	if err != nil {
		exists, errBucketExists := minioClient.BucketExists(ctx, bucketName)
		if !(errBucketExists == nil && exists) {
			logger.Error().Err(err).Msg("cannot connect to bucket")
			return err
		}
	}

	_, err = file.Seek(0, io.SeekStart)
	if err != nil {
		logger.Error().Err(err).Msg("cannot rewind file")
		return err
	}

	_, err = minioClient.PutObject(ctx, bucketName, objectName, file, filesize, minio.PutObjectOptions{
		UserMetadata: map[string]string{
			"x-amz-acl": viper.GetString("VK_CLOUD_ACCESS_MODE"),
		},
	})

	if err != nil {
		logger.Error().Err(err).Msg("cannot put avatar into bucket")
		return err
	}

	return nil
}
