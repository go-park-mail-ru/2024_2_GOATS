package delivery

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-park-mail-ru/2024_2_GOATS/config"
	"github.com/go-park-mail-ru/2024_2_GOATS/internal/app/api"
	"github.com/go-park-mail-ru/2024_2_GOATS/internal/app/api/converter"
	"github.com/go-park-mail-ru/2024_2_GOATS/internal/app/api/handlers"
	"github.com/go-park-mail-ru/2024_2_GOATS/internal/app/validation"
	"github.com/gorilla/mux"
	"github.com/rs/zerolog"
)

var _ handlers.UserImplementationInterface = (*UserHandler)(nil)

const (
	rParseErr      = "user_request_parse_error"
	vlErr          = "user_validation_error"
	uploadFileSize = 5 * 1024 * 1024
)

type UserHandler struct {
	userService UserServiceInterface
	logger      *zerolog.Logger
}

func NewUserHandler(ctx context.Context, srv UserServiceInterface) *UserHandler {
	return &UserHandler{
		userService: srv,
		logger:      &config.FromContext(ctx).Logger,
	}
}

func (u *UserHandler) UpdatePassword(w http.ResponseWriter, r *http.Request) {
	passwordReq := &api.UpdatePasswordRequest{}
	api.DecodeBody(w, r, passwordReq)

	vars := mux.Vars(r)
	usrId, err := getUserId(vars)
	if err != nil {
		errMsg := fmt.Errorf("updateProfile action: Path params err - %w", err)
		api.RequestError(w, u.logger, rParseErr, errMsg)

		return
	}

	passwordReq.UserId = usrId

	if err := validation.ValidatePassword(passwordReq.Password, passwordReq.PasswordConfirmation); err != nil {
		errMsg := fmt.Errorf("updatePassword action: Password err - %w", err.Err)
		api.RequestError(w, u.logger, vlErr, errMsg)

		return
	}

	ctx := u.logger.WithContext(r.Context())
	passwordServData := converter.ToServPasswordData(passwordReq)
	usrSrvResp, errSrvResp := u.userService.UpdatePassword(ctx, passwordServData)
	usrResp, errResp := converter.ToApiUpdateUserResponse(usrSrvResp), converter.ToApiErrorResponse(errSrvResp)

	if errResp != nil {
		api.Response(w, errResp.StatusCode, errResp)
		return
	}

	api.Response(w, usrResp.StatusCode, usrResp)
}

func (u *UserHandler) UpdateProfile(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	usrId, err := getUserId(vars)
	if err != nil {
		errMsg := fmt.Errorf("updateProfile action: Path params err - %w", err)
		api.RequestError(w, u.logger, rParseErr, errMsg)

		return
	}

	err = r.ParseMultipartForm(uploadFileSize)
	if err != nil {
		errMsg := fmt.Errorf("updateProfile action: Error parsing multipartForm - %w", err)
		api.RequestError(w, u.logger, rParseErr, errMsg)

		return
	}

	profileReq, err := u.parseProfileRequest(r, usrId)
	if err != nil {
		errMsg := fmt.Errorf("cannot read file from request: %w", err)
		u.logger.Err(errMsg)
		api.Response(w, http.StatusBadRequest, api.PreparedDefaultError("parse_request_error", errMsg))

		return
	}

	if valErr := validation.ValidateEmail(profileReq.Email); valErr != nil {
		errMsg := fmt.Errorf("updateProfile action: Email err - %w", valErr.Err)
		api.RequestError(w, u.logger, vlErr, errMsg)

		return
	}

	ctx := u.logger.WithContext(r.Context())
	profileServData := converter.ToServUserData(profileReq)
	usrSrvResp, errSrvResp := u.userService.UpdateProfile(ctx, profileServData)
	usrResp, errResp := converter.ToApiUpdateUserResponse(usrSrvResp), converter.ToApiErrorResponse(errSrvResp)

	if errResp != nil {
		api.Response(w, errResp.StatusCode, errResp)
		return
	}

	api.Response(w, usrResp.StatusCode, usrResp)
}

func (u *UserHandler) parseProfileRequest(r *http.Request, usrId int) (*api.UpdateProfileRequest, error) {
	formData := r.MultipartForm.Value
	file, handler, err := r.FormFile("avatar")

	if errors.Is(err, http.ErrMissingFile) {
		u.logger.Info().Msg("file was not given")
	} else {
		errMsg := fmt.Errorf("cannot read file from request: %w", err)
		u.logger.Err(errMsg)

		return nil, err
	}

	defer func() {
		if file != nil {
			if err := file.Close(); err != nil {
				u.logger.Err(fmt.Errorf("cannot close file: %w", err))
			}
		}
	}()

	profileReq := &api.UpdateProfileRequest{
		UserId:     usrId,
		Email:      getFormValue(formData, "email"),
		Username:   getFormValue(formData, "username"),
		Avatar:     file,
		AvatarName: handler.Filename,
	}

	return profileReq, nil
}

func getFormValue(formData map[string][]string, key string) string {
	if val, ok := formData[key]; ok && len(val) > 0 {
		return val[0]
	}
	return ""
}

func getUserId(vars map[string]string) (int, error) {
	usrId, err := strconv.Atoi(vars["id"])
	if err != nil {
		return 0, err
	}
	return usrId, nil
}
