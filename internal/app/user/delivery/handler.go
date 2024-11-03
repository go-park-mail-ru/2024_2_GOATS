package delivery

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-park-mail-ru/2024_2_GOATS/config"
	"github.com/go-park-mail-ru/2024_2_GOATS/internal/app/api"
	"github.com/go-park-mail-ru/2024_2_GOATS/internal/app/api/handlers"
	"github.com/go-park-mail-ru/2024_2_GOATS/internal/app/api/converter"
	errVals "github.com/go-park-mail-ru/2024_2_GOATS/internal/app/errors"
	"github.com/go-park-mail-ru/2024_2_GOATS/internal/app/logger"
	"github.com/go-park-mail-ru/2024_2_GOATS/internal/app/validation"
	"github.com/gorilla/mux"
)

var _ handlers.UserHandlerInterface = (*UserHandler)(nil)

const (
	rParseErr      = "user_request_parse_error"
	vlErr          = "user_validation_error"
	uploadFileSize = 5 * 1024 * 1024
)

type UserHandler struct {
	userService UserServiceInterface
	lg          *logger.BaseLogger
	locS        *config.LocalStorage
}

func NewUserHandler(ctx context.Context, srv UserServiceInterface) handlers.UserHandlerInterface {
	locS := config.FromContext(ctx).Databases.LocalStorage

	return &UserHandler{
		userService: srv,
		lg:          config.FromContext(ctx).Logger,
		locS:        &locS,
	}
}

func (u *UserHandler) UpdatePassword(w http.ResponseWriter, r *http.Request) {
	passwordReq := &api.UpdatePasswordRequest{}
	api.DecodeBody(w, r, passwordReq)

	vars := mux.Vars(r)
	usrId, err := getUserId(vars)
	if err != nil {
		errMsg := fmt.Errorf("updateProfile action: Path params err - %w", err)
		api.RequestError(w, u.lg, rParseErr, http.StatusBadRequest, errMsg)

		return
	}

	passwordReq.UserId = usrId

	if err := validation.ValidatePassword(passwordReq.Password, passwordReq.PasswordConfirmation); err != nil {
		errMsg := fmt.Errorf("updatePassword action: Password err - %w", err.Err)
		api.RequestError(w, u.lg, vlErr, http.StatusBadRequest, errMsg)

		return
	}

	ctx := api.BaseContext(w, r, u.lg)
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
		api.RequestError(w, u.lg, rParseErr, http.StatusBadRequest, errMsg)

		return
	}

	err = r.ParseMultipartForm(uploadFileSize)
	if err != nil {
		errMsg := fmt.Errorf("updateProfile action: Error parsing multipartForm - %w", err)
		api.RequestError(w, u.lg, rParseErr, http.StatusBadRequest, errMsg)

		return
	}

	profileReq, err := u.parseProfileRequest(w, r, usrId)
	if err != nil {
		errMsg := fmt.Errorf("cannot read file from request: %w", err)
		u.lg.LogError(errMsg.Error(), errMsg, api.GetRequestId(w))
		api.Response(w, http.StatusBadRequest, api.PreparedDefaultError("parse_request_error", errMsg))

		return
	}

	var errs []errVals.ErrorObj
	if profileReq.Username != "" {
		if valErr := validation.ValidateUsername(profileReq.Username); valErr != nil {
			errMsg := fmt.Errorf("updateProfile action: Username err - %w", valErr.Err)
			errs = append(errs, errVals.ErrorObj{Code: vlErr, Error: errVals.CustomError{Err: errMsg}})
		}
	}

	if profileReq.Email != "" {
		if valErr := validation.ValidateEmail(profileReq.Email); valErr != nil {
			errMsg := fmt.Errorf("updateProfile action: Email err - %w", valErr.Err)
			errs = append(errs, errVals.ErrorObj{Code: vlErr, Error: errVals.CustomError{Err: errMsg}})
		}
	}

	if len(errs) > 0 {
		errResp := &api.ErrorResponse{
			Errors:     errs,
			StatusCode: http.StatusBadRequest,
		}

		api.Response(w, errResp.StatusCode, errResp)
		return
	}

	ctx := api.BaseContext(w, r, u.lg)
	ctx = config.WrapLocalStorageContext(ctx, u.locS)
	profileServData := converter.ToServUserData(profileReq)
	usrSrvResp, errSrvResp := u.userService.UpdateProfile(ctx, profileServData)
	usrResp, errResp := converter.ToApiUpdateUserResponse(usrSrvResp), converter.ToApiErrorResponse(errSrvResp)

	if errResp != nil {
		api.Response(w, errResp.StatusCode, errResp)
		return
	}

	api.Response(w, usrResp.StatusCode, usrResp)
}

func (u *UserHandler) parseProfileRequest(w http.ResponseWriter, r *http.Request, usrId int) (*api.UpdateProfileRequest, error) {
	formData := r.MultipartForm.Value
	file, handler, err := r.FormFile("avatar")

	if err != nil {
		if errors.Is(err, http.ErrMissingFile) {
			u.lg.Log("file was not given", api.GetRequestId(w))
		} else {
			errMsg := fmt.Errorf("cannot read file from request: %w", err)
			u.lg.LogError(errMsg.Error(), errMsg, api.GetRequestId(w))

			return nil, err
		}
	}

	defer func() {
		if file != nil {
			if err := file.Close(); err != nil {
				u.lg.LogError("file_close_error", fmt.Errorf("cannot close file: %w", err), api.GetRequestId(w))
			}
		}
	}()

	var filename string
	if handler != nil {
		filename = handler.Filename
	}

	profileReq := &api.UpdateProfileRequest{
		UserId:     usrId,
		Email:      getFormValue(formData, "email"),
		Username:   getFormValue(formData, "username"),
		Avatar:     file,
		AvatarName: filename,
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
