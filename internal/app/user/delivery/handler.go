package delivery

import (
	"context"
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
	usrId, err := strconv.Atoi(vars["id"])
	if err != nil {
		errMsg := fmt.Errorf("updateProfile action: Path params err - %w", err)
		u.logger.Error().Msg(errMsg.Error())
		api.Response(w, http.StatusBadRequest, api.PreparedDefaultError("parse_request_error", errMsg))

		return
	}

	passwordReq.UserId = usrId

	if err := validation.ValidatePassword(passwordReq.Password, passwordReq.PasswordConfirmation); err != nil {
		errMsg := fmt.Errorf("updatePassword action: Password err - %w", err.Err)
		u.logger.Error().Msg(errMsg.Error())
		api.Response(w, http.StatusBadRequest, api.PreparedDefaultError("password_validation_error", errMsg))

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
	usrId, err := strconv.Atoi(vars["id"])
	if err != nil {
		errMsg := fmt.Errorf("updateProfile action: Path params err - %w", err)
		u.logger.Error().Msg(errMsg.Error())
		api.Response(w, http.StatusBadRequest, api.PreparedDefaultError("parse_request_error", errMsg))

		return
	}

	r.ParseMultipartForm(5 * 1024 * 1024)
	formData := r.MultipartForm.Value

	file, handler, err := r.FormFile("avatar")
	if err != nil {
		errMsg := fmt.Errorf("cannot read file from request: %w", err)
		u.logger.Err(errMsg)
		api.Response(w, http.StatusBadRequest, api.PreparedDefaultError("parse_request_error", errMsg))

		return
	}

	defer file.Close()

	profileReq := &api.UpdateProfileRequest{
    UserId: usrId,
	}

	if email, ok := formData["email"]; ok && len(email) > 0 {
		profileReq.Email = email[0]
	}

	if username, ok := formData["username"]; ok && len(username) > 0 {
		profileReq.Username = username[0]
	}

	if birthdate, ok := formData["birthdate"]; ok && len(birthdate) > 0 {
		profileReq.Birthdate = birthdate[0]
	}

	if sex, ok := formData["sex"]; ok && len(sex) > 0 {
		profileReq.Sex = sex[0]
	}

	if file != nil && handler.Filename != "" {
		profileReq.Avatar = file
		profileReq.AvatarName = handler.Filename
	}

	if valErr := validation.ValidateEmail(profileReq.Email); valErr != nil {
		errMsg := fmt.Errorf("updateProfile action: Email err - %w", valErr.Err)
		u.logger.Error().Msg(errMsg.Error())
		api.Response(w, http.StatusBadRequest, api.PreparedDefaultError("email_validation_error", errMsg))

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
