package delivery

import (
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/go-park-mail-ru/2024_2_GOATS/internal/app/api"
	"github.com/go-park-mail-ru/2024_2_GOATS/internal/app/api/converter"
	"github.com/go-park-mail-ru/2024_2_GOATS/internal/app/api/handlers"
	errVals "github.com/go-park-mail-ru/2024_2_GOATS/internal/app/errors"
	"github.com/go-park-mail-ru/2024_2_GOATS/internal/app/models"
	userDel "github.com/go-park-mail-ru/2024_2_GOATS/internal/app/user/delivery"
	"github.com/go-park-mail-ru/2024_2_GOATS/internal/app/validation"
	"github.com/rs/zerolog/log"
)

var _ handlers.AuthHandlerInterface = (*AuthHandler)(nil)

const (
	rParseErr = "auth_request_parse_error"
	vlErr     = "auth_validation_error"
)

// AuthHandler is a auth delivery handler struct
type AuthHandler struct {
	authService AuthServiceInterface
	userService userDel.UserServiceInterface
}

// NewAuthHandler returns an instance of AuthHandlerInterface
func NewAuthHandler(authSrv AuthServiceInterface, usrSrv userDel.UserServiceInterface) handlers.AuthHandlerInterface {
	return &AuthHandler{
		authService: authSrv,
		userService: usrSrv,
	}
}

// Logout logout http handler method
func (a *AuthHandler) Logout(w http.ResponseWriter, r *http.Request) {
	logger := log.Ctx(r.Context())
	ck, err := r.Cookie("session_id")
	if errors.Is(err, http.ErrNoCookie) {
		errMsg := fmt.Errorf("logout action: No cookie err - %w", err)
		api.RequestError(r.Context(), w, rParseErr, http.StatusBadRequest, errMsg)

		return
	}

	validErr := validation.ValidateCookie(ck.Value)
	if validErr != nil {
		errMsg := fmt.Errorf("logout action: Invalid cookie err - %w", validErr)
		api.RequestError(r.Context(), w, vlErr, http.StatusBadRequest, errMsg)

		return
	}

	errSrvResp := a.authService.Logout(r.Context(), ck.Value)
	errResp := errVals.ToDeliveryErrorFromService(errSrvResp)
	if errResp != nil {
		errMsg := errors.New("failed to logout")
		logger.Error().Err(errMsg).Interface("logoutResp", errResp).Msg("request_failed")
		api.Response(r.Context(), w, errResp.HTTPStatus, errResp)

		return
	}

	http.SetCookie(w, preparedExpiredCookie())
	logger.Info().Msg("Logout success")

	api.Response(r.Context(), w, http.StatusOK, nil)
}

// Login login http handler method
func (a *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	logger := log.Ctx(r.Context())
	loginRequest := &api.LoginRequest{}

	if !api.DecodeBody(w, r, loginRequest) {
		return
	}

	loginServData := converter.ToServLoginData(loginRequest)
	authSrvResp, errSrvResp := a.authService.Login(r.Context(), loginServData)

	authResp, errResp := converter.ToAPIAuthResponse(authSrvResp), errVals.ToDeliveryErrorFromService(errSrvResp)
	if errResp != nil {
		errMsg := errors.New("failed to login")
		logger.Error().Err(errMsg).Interface("loginResp", errResp).Msg("request_failed")
		api.Response(r.Context(), w, errResp.HTTPStatus, errResp)

		return
	}

	oldCookie, err := r.Cookie("session_id")
	if errors.Is(err, http.ErrNoCookie) {
		logger.Info().Msg("user dont have old cookie")
	}

	if oldCookie != nil && authResp.NewCookie.Token.TokenID != oldCookie.Value {
		logger.Info().Msg("successfully expire cookie")
		http.SetCookie(w, preparedExpiredCookie())
	}

	http.SetCookie(w, preparedCookie(authResp.NewCookie))
	logger.Info().Msg("successfully set new cookie")
	logger.Info().Interface("loginResp", authResp).Msg("login success")

	api.Response(r.Context(), w, http.StatusOK, authResp)
}

// Register register_user http handler method
func (a *AuthHandler) Register(w http.ResponseWriter, r *http.Request) {
	logger := log.Ctx(r.Context())

	registerReq := &api.RegisterRequest{}
	if !api.DecodeBody(w, r, registerReq) {
		return
	}

	errs := make([]errVals.ErrorItem, 0)

	if err := validation.ValidatePassword(registerReq.Password, registerReq.PasswordConfirmation); err != nil {
		logger.Error().Msgf("%s:%s", vlErr, err)
		addError(errVals.ErrInvalidPasswordCode, err, &errs)
	}

	if err := validation.ValidateEmail(registerReq.Email); err != nil {
		logger.Error().Msgf("%s:%s", vlErr, err)
		addError(errVals.ErrInvalidEmailCode, err, &errs)
	}

	if len(errs) > 0 {
		errResp := errVals.NewDeliveryError(http.StatusBadRequest, errs)
		api.Response(r.Context(), w, errResp.HTTPStatus, errResp)
		return
	}

	registerServData := converter.ToServRegisterData(registerReq)
	authSrvResp, errSrvResp := a.authService.Register(r.Context(), registerServData)
	authResp, errResp := converter.ToAPIAuthResponse(authSrvResp), errVals.ToDeliveryErrorFromService(errSrvResp)

	if errResp != nil {
		errMsg := errors.New("failed to register")
		logger.Error().Err(errMsg).Interface("registerResp", errResp).Msg("request_failed")
		api.Response(r.Context(), w, errResp.HTTPStatus, errResp)

		return
	}

	http.SetCookie(w, preparedCookie(authResp.NewCookie))
	logger.Info().Interface("registerResp", authResp).Msg("Register success")

	api.Response(r.Context(), w, http.StatusOK, authResp)
}

// Session check_session http handler method
func (a *AuthHandler) Session(w http.ResponseWriter, r *http.Request) {
	logger := log.Ctx(r.Context())
	ck, err := r.Cookie("session_id")
	if errors.Is(err, http.ErrNoCookie) {
		errMsg := fmt.Errorf("session action: No cookie err - %w", err)
		api.RequestError(r.Context(), w, rParseErr, http.StatusForbidden, errMsg)

		return
	}

	sessionSrvResp, errSrvResp := a.authService.Session(r.Context(), ck.Value)

	sessionResp, errResp := converter.ToAPISessionResponse(sessionSrvResp), errVals.ToDeliveryErrorFromService(errSrvResp)
	if errResp != nil {
		errMsg := errors.New("failed to get session")
		logger.Error().Err(errMsg).Interface("sessionResp", errResp).Msg("request_failed")
		api.Response(r.Context(), w, errResp.HTTPStatus, errResp)

		return
	}

	logger.Info().Interface("sessionResp", sessionResp).Msg("getSession success")

	api.Response(r.Context(), w, http.StatusOK, sessionResp)
}

func preparedCookie(ck *models.CookieData) *http.Cookie {
	return &http.Cookie{
		Name:     ck.Name,
		Value:    ck.Token.TokenID,
		Expires:  ck.Token.Expiry,
		HttpOnly: true,
		SameSite: http.SameSiteLaxMode,
		Path:     "/api",
		Secure:   false,
	}
}

func preparedExpiredCookie() *http.Cookie {
	return &http.Cookie{
		Name:     "session_id",
		Value:    "",
		Expires:  time.Unix(0, 0),
		HttpOnly: true,
		SameSite: http.SameSiteLaxMode,
		Path:     "/api",
		Secure:   false,
	}
}

func addError(code string, err error, errors *[]errVals.ErrorItem) {
	errStruct := errVals.ErrorItem{
		Code:  code,
		Error: err.Error(),
	}

	*errors = append(*errors, errStruct)
}
