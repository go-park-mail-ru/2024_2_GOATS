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
		errMsg := fmt.Errorf("UpdateProfile action: Path params err - %w", err)
		u.logger.Error().Msg(errMsg.Error())
		api.Response(w, http.StatusBadRequest, api.PreparedDefaultError("parse_request_error", errMsg))

		return
	}

	passwordReq.UserId = usrId

	if err := validation.ValidatePassword(passwordReq.Password, passwordReq.PasswordConfirmation); err != nil {
		errMsg := fmt.Errorf("UpdatePassword action: Password err - %w", err.Err)
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
	profileReq := &api.UpdateProfileRequest{}
	api.DecodeBody(w, r, profileReq)

	vars := mux.Vars(r)
	usrId, err := strconv.Atoi(vars["id"])
	if err != nil {
		errMsg := fmt.Errorf("UpdateProfile action: Path params err - %w", err)
		u.logger.Error().Msg(errMsg.Error())
		api.Response(w, http.StatusBadRequest, api.PreparedDefaultError("parse_request_error", errMsg))

		return
	}

	profileReq.UserId = usrId

	if valErr := validation.ValidateEmail(profileReq.Email); valErr != nil {
		errMsg := fmt.Errorf("UpdateProfile action: Email err - %w", valErr.Err)
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
