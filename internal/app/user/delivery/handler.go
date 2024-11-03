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
	errVals "github.com/go-park-mail-ru/2024_2_GOATS/internal/app/errors"
	"github.com/go-park-mail-ru/2024_2_GOATS/internal/app/validation"
	"github.com/gorilla/mux"
	"github.com/rs/zerolog/log"
)

var _ handlers.UserHandlerInterface = (*UserHandler)(nil)

const (
	rParseErr      = "user_request_parse_error"
	vlErr          = "user_validation_error"
	uploadFileSize = 5 * 1024 * 1024
)

type UserHandler struct {
	userService UserServiceInterface
	locS        *config.LocalStorage
}

func NewUserHandler(ctx context.Context, srv UserServiceInterface) handlers.UserHandlerInterface {
	locS := config.FromContext(ctx).Databases.LocalStorage

	return &UserHandler{
		userService: srv,
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
		api.RequestError(r.Context(), w, rParseErr, http.StatusBadRequest, errMsg)

		return
	}

	passwordReq.UserId = usrId

	if err := validation.ValidatePassword(passwordReq.Password, passwordReq.PasswordConfirmation); err != nil {
		errMsg := fmt.Errorf("updatePassword action: Password err - %w", err.Err)
		api.RequestError(r.Context(), w, vlErr, http.StatusBadRequest, errMsg)

		return
	}

	passwordServData := converter.ToServPasswordData(passwordReq)
	usrSrvResp, errSrvResp := u.userService.UpdatePassword(r.Context(), passwordServData)
	usrResp, errResp := converter.ToApiUpdateUserResponse(usrSrvResp), converter.ToApiErrorResponse(errSrvResp)

	if errResp != nil {
		api.Response(r.Context(), w, errResp.StatusCode, errResp)
		return
	}

	api.Response(r.Context(), w, usrResp.StatusCode, usrResp)
}

func (u *UserHandler) UpdateProfile(w http.ResponseWriter, r *http.Request) {
	lg := log.Ctx(r.Context())
	ctx := config.WrapLocalStorageContext(r.Context(), u.locS)
	vars := mux.Vars(r)
	usrId, err := getUserId(vars)
	if err != nil {
		errMsg := fmt.Errorf("updateProfile action: Path params err - %w", err)
		api.RequestError(ctx, w, rParseErr, http.StatusBadRequest, errMsg)

		return
	}

	err = r.ParseMultipartForm(uploadFileSize)
	if err != nil {
		errMsg := fmt.Errorf("updateProfile action: Error parsing multipartForm - %w", err)
		api.RequestError(ctx, w, rParseErr, http.StatusBadRequest, errMsg)

		return
	}

	profileReq, err := u.parseProfileRequest(r, usrId)
	if err != nil {
		errMsg := fmt.Errorf("cannot read file from request: %w", err)
		lg.Error().Msg(errMsg.Error())
		api.Response(ctx, w, http.StatusBadRequest, api.PreparedDefaultError("parse_request_error", errMsg))

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

		api.Response(ctx, w, errResp.StatusCode, errResp)
		return
	}

	profileServData := converter.ToServUserData(profileReq)
	usrSrvResp, errSrvResp := u.userService.UpdateProfile(ctx, profileServData)
	usrResp, errResp := converter.ToApiUpdateUserResponse(usrSrvResp), converter.ToApiErrorResponse(errSrvResp)

	if errResp != nil {
		api.Response(ctx, w, errResp.StatusCode, errResp)
		return
	}

	api.Response(ctx, w, usrResp.StatusCode, usrResp)
}

func (u *UserHandler) parseProfileRequest(r *http.Request, usrId int) (*api.UpdateProfileRequest, error) {
	lg := log.Ctx(r.Context())
	formData := r.MultipartForm.Value
	file, handler, err := r.FormFile("avatar")

	if err != nil {
		if errors.Is(err, http.ErrMissingFile) {
			lg.Info().Msg("file was not given")
		} else {
			errMsg := fmt.Errorf("cannot read file from request: %w", err)
			lg.Error().Msg(errMsg.Error())

			return nil, err
		}
	}

	defer func() {
		if file != nil {
			if err := file.Close(); err != nil {
				lg.Error().Msg(fmt.Sprintf("cannot close file: %v", err))
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
