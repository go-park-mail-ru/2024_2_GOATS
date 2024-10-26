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
	"github.com/rs/zerolog"
)

var _ handlers.AuthImplementationInterface = (*AuthHandler)(nil)

type AuthHandler struct {
	authService AuthServiceInterface
	userService userDel.UserServiceInterface
	redisCfg    *config.Redis
	logger      *zerolog.Logger
}

func NewAuthHandler(ctx context.Context, authSrv AuthServiceInterface, usrSrv userDel.UserServiceInterface) *AuthHandler {
	redisCfg := config.FromContext(ctx).Databases.Redis
	logger := config.FromContext(ctx).Logger

	return &AuthHandler{
		authService: authSrv,
		userService: usrSrv,
		redisCfg:    &redisCfg,
		logger:      &logger,
	}
}

func (a *AuthHandler) Logout(w http.ResponseWriter, r *http.Request) {
	ck, err := r.Cookie("session_id")
	if errors.Is(err, http.ErrNoCookie) {
		errMsg := fmt.Errorf("logout action: No cookie err - %w", err)
		a.logger.Error().Msg(errMsg.Error())
		api.Response(w, http.StatusForbidden, api.PreparedDefaultError(errVals.ErrNoCookieCode, errMsg))

		return
	}

	validErr := validation.ValidateCookie(ck.Value)
	if validErr != nil {
		errMsg := fmt.Errorf("logout action: Invalid cookie err - %w", validErr.Err)
		a.logger.Error().Msg(errMsg.Error())
		api.Response(w, http.StatusBadRequest, api.PreparedDefaultError("cookie_validation_error", errMsg))

		return
	}

	ctx := a.logger.WithContext(r.Context())
	logoutSrvResp, errSrvResp := a.authService.Logout(ctx, ck.Value)

	logoutResp, errResp := converter.ToApiAuthResponse(logoutSrvResp), converter.ToApiErrorResponse(errSrvResp)
	if errResp != nil {
		api.Response(w, errResp.StatusCode, errResp)
		return
	}

	http.SetCookie(w, preparedExpiredCookie())

	api.Response(w, logoutResp.StatusCode, logoutResp)
}

func (a *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	loginRequest := &api.LoginRequest{}
	api.DecodeBody(w, r, loginRequest)

	loginServData := converter.ToServLoginData(loginRequest)
	ctx := config.WrapRedisContext(r.Context(), a.redisCfg)
	ctx = a.logger.WithContext(ctx)
	authSrvResp, errSrvResp := a.authService.Login(ctx, loginServData)

	authResp, errResp := converter.ToApiAuthResponse(authSrvResp), converter.ToApiErrorResponse(errSrvResp)
	if errResp != nil {
		api.Response(w, errResp.StatusCode, errResp)
		return
	}

	oldCookie, err := r.Cookie("session_id")
	if errors.Is(err, http.ErrNoCookie) {
		a.logger.Info().Msg("user dont have old cookie")
	}

	if oldCookie != nil && authResp.NewCookie.Token.TokenID != oldCookie.Value {
		a.logger.Info().Msg("successfully expire cookie")
		http.SetCookie(w, preparedExpiredCookie())
	}

	http.SetCookie(w, preparedCookie(authResp.NewCookie))
	a.logger.Info().Msg("successfully set new cookie")

	api.Response(w, authResp.StatusCode, authResp)
}

func (a *AuthHandler) Register(w http.ResponseWriter, r *http.Request) {
	registerReq := &api.RegisterRequest{}
	api.DecodeBody(w, r, registerReq)

	errs := make([]errVals.ErrorObj, 0)

	if err := validation.ValidatePassword(registerReq.Password, registerReq.PasswordConfirmation); err != nil {
		a.logger.Error().Msg(err.Err.Error())
		addError(errVals.ErrInvalidPasswordCode, *err, &errs)
	}

	if err := validation.ValidateEmail(registerReq.Email); err != nil {
		a.logger.Error().Msg(err.Err.Error())
		addError(errVals.ErrInvalidEmailCode, *err, &errs)
	}

	if len(errs) > 0 {
		errResp := &api.ErrorResponse{
			Errors:     errs,
			StatusCode: http.StatusBadRequest,
		}

		api.Response(w, errResp.StatusCode, errResp)
		return
	}

	ctx := config.WrapRedisContext(r.Context(), a.redisCfg)
	ctx = a.logger.WithContext(ctx)

	registerServData := converter.ToServRegisterData(registerReq)
	authSrvResp, errSrvResp := a.authService.Register(ctx, registerServData)
	authResp, errResp := converter.ToApiAuthResponse(authSrvResp), converter.ToApiErrorResponse(errSrvResp)

	if errResp != nil {
		api.Response(w, errResp.StatusCode, errResp)
		return
	}

	http.SetCookie(w, preparedCookie(authResp.NewCookie))

	api.Response(w, authResp.StatusCode, authResp)
}

func (a *AuthHandler) Session(w http.ResponseWriter, r *http.Request) {
	ck, err := r.Cookie("session_id")
	if errors.Is(err, http.ErrNoCookie) {
		errMsg := fmt.Errorf("session action: No cookie err - %w", err)
		a.logger.Error().Msg(errMsg.Error())
		api.Response(w, http.StatusForbidden, api.PreparedDefaultError(errVals.ErrNoCookieCode, errMsg))

		return
	}

	ctx := a.logger.WithContext(r.Context())
	sessionSrvResp, errSrvResp := a.authService.Session(ctx, ck.Value)

	sessionResp, errResp := converter.ToApiSessionResponse(sessionSrvResp), converter.ToApiErrorResponse(errSrvResp)
	if errResp != nil {
		api.Response(w, errResp.StatusCode, errResp)
		return
	}

	api.Response(w, sessionResp.StatusCode, sessionResp)
}

func preparedCookie(ck *models.CookieData) *http.Cookie {
	return &http.Cookie{
		Name:     ck.Name,
		Value:    ck.Token.TokenID,
		Expires:  ck.Token.Expiry,
		HttpOnly: true,
		SameSite: http.SameSiteLaxMode,
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
