package delivery

import (
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
	destroyFavOp   = "destroy_favorite"
	setFavOp       = "set_favorite"
)

// UserHandler user http handler struct
type UserHandler struct {
	userService UserServiceInterface
}

// NewUserHandler returns an instance of UserHandlerInterface
func NewUserHandler(srv UserServiceInterface) handlers.UserHandlerInterface {
	return &UserHandler{
		userService: srv,
	}
}

// UpdatePassword updates_user_password http handler method
func (u *UserHandler) UpdatePassword(w http.ResponseWriter, r *http.Request) {
	logger := log.Ctx(r.Context())
	passwordReq := &api.UpdatePasswordRequest{}
	api.DecodeBody(w, r, passwordReq)

	vars := mux.Vars(r)
	usrID, err := getUserID(vars)
	if err != nil {
		errMsg := fmt.Errorf("updateProfile action: Path params err - %w", err)
		api.RequestError(r.Context(), w, rParseErr, http.StatusBadRequest, errMsg)

		return
	}

	passwordReq.UserID = usrID

	if err := validation.ValidatePassword(passwordReq.Password, passwordReq.PasswordConfirmation); err != nil {
		errMsg := fmt.Errorf("updatePassword action: Password err - %w", err.Err)
		api.RequestError(r.Context(), w, vlErr, http.StatusBadRequest, errMsg)

		return
	}

	passwordServData := converter.ToServPasswordData(passwordReq)
	errSrvResp := u.userService.UpdatePassword(r.Context(), passwordServData)
	errResp := errVals.ToDeliveryErrorFromService(errSrvResp)

	if errResp != nil {
		errMsg := errors.New("failed to update password")
		logger.Error().Err(errMsg).Interface("updatePasswdResp", errResp).Msg("request_failed")
		api.Response(r.Context(), w, errResp.HTTPStatus, errResp)

		return
	}

	logger.Info().Interface("updatePasswdResp", true).Msg("updatePasswd success")

	api.Response(r.Context(), w, http.StatusOK, nil)
}

// UpdateProfile updates_user_profile http handler method
func (u *UserHandler) UpdateProfile(w http.ResponseWriter, r *http.Request) {
	logger := log.Ctx(r.Context())
	vars := mux.Vars(r)
	usrID, err := getUserID(vars)
	if err != nil {
		errMsg := fmt.Errorf("updateProfile action: Path params err - %w", err)
		api.RequestError(r.Context(), w, rParseErr, http.StatusBadRequest, errMsg)

		return
	}

	err = r.ParseMultipartForm(uploadFileSize)
	if err != nil {
		errMsg := fmt.Errorf("updateProfile action: Error parsing multipartForm - %w", err)
		api.RequestError(r.Context(), w, rParseErr, http.StatusBadRequest, errMsg)

		return
	}

	profileReq, err := u.parseProfileRequest(r, usrID)
	if err != nil {
		errMsg := fmt.Errorf("cannot read file from request: %w", err)
		logger.Error().Err(errMsg).Msg("read_file_error")
		api.Response(r.Context(), w, http.StatusBadRequest, api.PreparedDefaultError("parse_request_error", errMsg))

		return
	}

	var errs = make([]errVals.ErrorItem, 0)
	if profileReq.Username != "" {
		if valErr := validation.ValidateUsername(profileReq.Username); valErr != nil {
			errMsg := fmt.Errorf("updateProfile action: Username err - %w", valErr.Err)
			errs = append(errs, errVals.ErrorItem{Code: vlErr, Error: errVals.NewCustomError(errMsg.Error())})
		}
	}

	if profileReq.Email != "" {
		if valErr := validation.ValidateEmail(profileReq.Email); valErr != nil {
			errMsg := fmt.Errorf("updateProfile action: Email err - %w", valErr.Err)
			errs = append(errs, errVals.ErrorItem{Code: vlErr, Error: errVals.NewCustomError(errMsg.Error())})
		}
	}

	if len(errs) > 0 {
		errResp := errVals.NewDeliveryError(http.StatusBadRequest, errs)

		api.Response(r.Context(), w, errResp.HTTPStatus, errResp)
		return
	}

	profileServData := converter.ToServUserData(profileReq)
	errSrvResp := u.userService.UpdateProfile(r.Context(), profileServData)
	errResp := errVals.ToDeliveryErrorFromService(errSrvResp)

	if errResp != nil {
		errMsg := errors.New("failed to update profile")
		logger.Error().Err(errMsg).Interface("updateProfileResp", errResp).Msg("request_failed")
		api.Response(r.Context(), w, errResp.HTTPStatus, errResp)
		return
	}

	logger.Info().Bool("updateProfileResp", true).Msg("updateProfile success")

	api.Response(r.Context(), w, http.StatusOK, nil)
}

func (u *UserHandler) parseProfileRequest(r *http.Request, usrID int) (*api.UpdateProfileRequest, error) {
	logger := log.Ctx(r.Context())
	formData := r.MultipartForm.Value
	file, handler, err := r.FormFile("avatar")

	defer func() {
		if file != nil {
			if clErr := file.Close(); clErr != nil {
				logger.Error().Err(fmt.Errorf("cannot close file: %v", clErr)).Msg("close_file_error")
			}
		}
	}()

	if err != nil {
		if errors.Is(err, http.ErrMissingFile) {
			logger.Info().Msg("file was not given")
		} else {
			errMsg := fmt.Errorf("cannot read file from request: %w", err)
			logger.Error().Err(errMsg).Msg("read_file_error")

			return nil, err
		}
	}

	var filename string
	if handler != nil {
		filename = handler.Filename
	}

	profileReq := &api.UpdateProfileRequest{
		UserID:     usrID,
		Email:      getFormValue(formData, "email"),
		Username:   getFormValue(formData, "username"),
		AvatarFile: file,
		AvatarName: filename,
	}

	return profileReq, nil
}

// SetFavorite set_user_favorite http handler method
func (u *UserHandler) SetFavorite(w http.ResponseWriter, r *http.Request) {
	u.toggleFavorite(w, r)
}

// ResetFavorite reset_user_favorite http handler method
func (u *UserHandler) ResetFavorite(w http.ResponseWriter, r *http.Request) {
	u.toggleFavorite(w, r)
}

func (u *UserHandler) toggleFavorite(w http.ResponseWriter, r *http.Request) {
	var err *errVals.ServiceError
	var op string

	logger := log.Ctx(r.Context())
	favReq := &api.FavReq{}
	api.DecodeBody(w, r, favReq)
	favReq.UserID = config.CurrentUserID(r.Context())

	favSrvData := converter.ToServFavData(favReq)

	if r.Method == http.MethodDelete {
		op = destroyFavOp
		err = u.userService.ResetFavorite(r.Context(), favSrvData)
	} else {
		op = setFavOp
		err = u.userService.AddFavorite(r.Context(), favSrvData)
	}

	if err != nil {
		errResp := errVals.ToDeliveryErrorFromService(err)
		errMsg := fmt.Errorf("failed to %s", op)
		logger.Error().Err(errMsg).Interface("favResp", errResp).Msg("request_failed")
		api.Response(r.Context(), w, errResp.HTTPStatus, errResp)

		return
	}

	logger.Info().Msgf("%s success", op)

	api.Response(r.Context(), w, http.StatusOK, nil)
}

// GetFavorites get_user_favorites http handler method
func (u *UserHandler) GetFavorites(w http.ResponseWriter, r *http.Request) {
	logger := log.Ctx(r.Context())
	vars := mux.Vars(r)
	usrID, err := getUserID(vars)
	if err != nil {
		errMsg := fmt.Errorf("updateProfile action: Path params err - %w", err)
		api.RequestError(r.Context(), w, rParseErr, http.StatusBadRequest, errMsg)

		return
	}

	srvResp, srvRespErr := u.userService.GetFavorites(r.Context(), usrID)
	resp, respErr := converter.ToAPIMovieShortInfos(srvResp), errVals.ToDeliveryErrorFromService(srvRespErr)
	if respErr != nil {
		errMsg := errors.New("failed to get user favorites")
		logger.Error().Err(errMsg).Interface("favResp", respErr).Msg("request_failed")
		api.Response(r.Context(), w, respErr.HTTPStatus, respErr)

		return
	}

	logger.Info().Interface("GetfavResp", resp).Msg("Favorites success")

	api.Response(r.Context(), w, http.StatusOK, resp)
}

func getFormValue(formData map[string][]string, key string) string {
	return formData[key][0]
}

func getUserID(vars map[string]string) (int, error) {
	usrID, err := strconv.Atoi(vars["id"])
	if err != nil {
		return 0, err
	}
	return usrID, nil
}
