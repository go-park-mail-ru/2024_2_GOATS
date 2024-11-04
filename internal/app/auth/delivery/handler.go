package delivery

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/go-park-mail-ru/2024_2_GOATS/config"
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

type AuthHandler struct {
	authService AuthServiceInterface
	userService userDel.UserServiceInterface
	redisCfg    *config.Redis
}

func NewAuthHandler(ctx context.Context, authSrv AuthServiceInterface, usrSrv userDel.UserServiceInterface) handlers.AuthHandlerInterface {
	redisCfg := config.FromContext(ctx).Databases.Redis

	return &AuthHandler{
		authService: authSrv,
		userService: usrSrv,
		redisCfg:    &redisCfg,
	}
}

func (a *AuthHandler) Logout(w http.ResponseWriter, r *http.Request) {
	lg := log.Ctx(r.Context())
	ck, err := r.Cookie("session_id")
	if errors.Is(err, http.ErrNoCookie) {
		errMsg := fmt.Errorf("logout action: No cookie err - %w", err)
		api.RequestError(r.Context(), w, rParseErr, http.StatusBadRequest, errMsg)

		return
	}

	validErr := validation.ValidateCookie(ck.Value)
	if validErr != nil {
		errMsg := fmt.Errorf("logout action: Invalid cookie err - %w", validErr.Err)
		api.RequestError(r.Context(), w, vlErr, http.StatusBadRequest, errMsg)

		return
	}

	logoutSrvResp, errSrvResp := a.authService.Logout(r.Context(), ck.Value)

	logoutResp, errResp := converter.ToApiAuthResponse(logoutSrvResp), converter.ToApiErrorResponse(errSrvResp)
	if errResp != nil {
		errMsg := errors.New("failed to logout")
		lg.Error().Err(errMsg).Interface("logoutResp", errResp).Msg("request_failed")
		api.Response(r.Context(), w, errResp.StatusCode, errResp)

		return
	}

	http.SetCookie(w, preparedExpiredCookie())
	lg.Info().Interface("logoutResp", logoutResp).Msg("Logout success")

	api.Response(r.Context(), w, logoutResp.StatusCode, logoutResp)
}

func (a *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	lg := log.Ctx(r.Context())
	loginRequest := &api.LoginRequest{}
	ctx := config.WrapRedisContext(r.Context(), a.redisCfg)

	api.DecodeBody(w, r, loginRequest)
	loginServData := converter.ToServLoginData(loginRequest)
	authSrvResp, errSrvResp := a.authService.Login(ctx, loginServData)

	authResp, errResp := converter.ToApiAuthResponse(authSrvResp), converter.ToApiErrorResponse(errSrvResp)
	if errResp != nil {
		errMsg := errors.New("failed to login")
		lg.Error().Err(errMsg).Interface("loginResp", errResp).Msg("request_failed")
		api.Response(ctx, w, errResp.StatusCode, errResp)

		return
	}

	oldCookie, err := r.Cookie("session_id")
	if errors.Is(err, http.ErrNoCookie) {
		lg.Info().Msg("user dont have old cookie")
	}

	if oldCookie != nil && authResp.NewCookie.Token.TokenID != oldCookie.Value {
		lg.Info().Msg("successfully expire cookie")
		http.SetCookie(w, preparedExpiredCookie())
	}

	http.SetCookie(w, preparedCookie(authResp.NewCookie))
	lg.Info().Msg("successfully set new cookie")
	lg.Info().Interface("loginResp", authResp).Msg("login success")

	api.Response(ctx, w, authResp.StatusCode, authResp)
}

func (a *AuthHandler) Register(w http.ResponseWriter, r *http.Request) {
	lg := log.Ctx(r.Context())
	ctx := config.WrapRedisContext(r.Context(), a.redisCfg)

	registerReq := &api.RegisterRequest{}
	api.DecodeBody(w, r, registerReq)

	errs := make([]errVals.ErrorObj, 0)

	if err := validation.ValidatePassword(registerReq.Password, registerReq.PasswordConfirmation); err != nil {
		lg.Error().Err(err.Err).Msg(vlErr)
		addError(errVals.ErrInvalidPasswordCode, *err, &errs)
	}

	if err := validation.ValidateEmail(registerReq.Email); err != nil {
		lg.Error().Err(err.Err).Msg(vlErr)
		addError(errVals.ErrInvalidEmailCode, *err, &errs)
	}

	if len(errs) > 0 {
		errResp := &api.ErrorResponse{
			Errors:     errs,
			StatusCode: http.StatusBadRequest,
		}

		api.Response(ctx, w, errResp.StatusCode, errResp)
		return
	}

	registerServData := converter.ToServRegisterData(registerReq)
	authSrvResp, errSrvResp := a.authService.Register(ctx, registerServData)
	authResp, errResp := converter.ToApiAuthResponse(authSrvResp), converter.ToApiErrorResponse(errSrvResp)

	if errResp != nil {
		errMsg := errors.New("failed to register")
		lg.Error().Err(errMsg).Interface("registerResp", errResp).Msg("request_failed")
		api.Response(ctx, w, errResp.StatusCode, errResp)

		return
	}

	http.SetCookie(w, preparedCookie(authResp.NewCookie))
	lg.Info().Interface("registerResp", authResp).Msg("Register success")

	api.Response(ctx, w, authResp.StatusCode, authResp)
}

func (a *AuthHandler) Session(w http.ResponseWriter, r *http.Request) {
	lg := log.Ctx(r.Context())
	ck, err := r.Cookie("session_id")
	if errors.Is(err, http.ErrNoCookie) {
		errMsg := fmt.Errorf("session action: No cookie err - %w", err)
		api.RequestError(r.Context(), w, rParseErr, http.StatusForbidden, errMsg)

		return
	}

	sessionSrvResp, errSrvResp := a.authService.Session(r.Context(), ck.Value)

	sessionResp, errResp := converter.ToApiSessionResponse(sessionSrvResp), converter.ToApiErrorResponse(errSrvResp)
	if errResp != nil {
		errMsg := errors.New("failed to get session")
		lg.Error().Err(errMsg).Interface("sessionResp", errResp).Msg("request_failed")
		api.Response(r.Context(), w, errResp.StatusCode, errResp)

		return
	}

	lg.Info().Interface("sessionResp", sessionResp).Msg("getSession success")

	api.Response(r.Context(), w, sessionResp.StatusCode, sessionResp)
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

func addError(code string, err errVals.CustomError, errors *[]errVals.ErrorObj) {
	errStruct := errVals.ErrorObj{
		Code:  code,
		Error: err,
	}

	*errors = append(*errors, errStruct)
}
