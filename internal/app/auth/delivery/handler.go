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
	"github.com/go-park-mail-ru/2024_2_GOATS/internal/app/logger"
	"github.com/go-park-mail-ru/2024_2_GOATS/internal/app/models"
	userDel "github.com/go-park-mail-ru/2024_2_GOATS/internal/app/user/delivery"
	"github.com/go-park-mail-ru/2024_2_GOATS/internal/app/validation"
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
	lg          *logger.BaseLogger
}

func NewAuthHandler(ctx context.Context, authSrv AuthServiceInterface, usrSrv userDel.UserServiceInterface) handlers.AuthHandlerInterface {
	redisCfg := config.FromContext(ctx).Databases.Redis
	logger := config.FromContext(ctx).Logger

	return &AuthHandler{
		authService: authSrv,
		userService: usrSrv,
		redisCfg:    &redisCfg,
		lg:          logger,
	}
}

func (a *AuthHandler) Logout(w http.ResponseWriter, r *http.Request) {
	ctx := api.BaseContext(w, r, a.lg)

	ck, err := r.Cookie("session_id")
	if errors.Is(err, http.ErrNoCookie) {
		errMsg := fmt.Errorf("logout action: No cookie err - %w", err)
		api.RequestError(w, a.lg, rParseErr, http.StatusBadRequest, errMsg)

		return
	}

	validErr := validation.ValidateCookie(ck.Value)
	if validErr != nil {
		errMsg := fmt.Errorf("logout action: Invalid cookie err - %w", validErr.Err)
		api.RequestError(w, a.lg, vlErr, http.StatusBadRequest, errMsg)

		return
	}

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
	ctx := api.BaseContext(w, r, a.lg)
	ctx = config.WrapRedisContext(ctx, a.redisCfg)

	authSrvResp, errSrvResp := a.authService.Login(ctx, loginServData)

	authResp, errResp := converter.ToApiAuthResponse(authSrvResp), converter.ToApiErrorResponse(errSrvResp)
	if errResp != nil {
		api.Response(w, errResp.StatusCode, errResp)
		return
	}

	oldCookie, err := r.Cookie("session_id")
	if errors.Is(err, http.ErrNoCookie) {
		a.lg.Log("user dont have old cookie", api.GetRequestId(w))
	}

	if oldCookie != nil && authResp.NewCookie.Token.TokenID != oldCookie.Value {
		a.lg.Log("successfully expire cookie", api.GetRequestId(w))
		http.SetCookie(w, preparedExpiredCookie())
	}

	http.SetCookie(w, preparedCookie(authResp.NewCookie))
	a.lg.Log("successfully set new cookie", api.GetRequestId(w))

	api.Response(w, authResp.StatusCode, authResp)
}

func (a *AuthHandler) Register(w http.ResponseWriter, r *http.Request) {
	registerReq := &api.RegisterRequest{}
	api.DecodeBody(w, r, registerReq)

	errs := make([]errVals.ErrorObj, 0)

	if err := validation.ValidatePassword(registerReq.Password, registerReq.PasswordConfirmation); err != nil {
		a.lg.LogError(err.Err.Error(), err.Err, api.GetRequestId(w))
		addError(errVals.ErrInvalidPasswordCode, *err, &errs)
	}

	if err := validation.ValidateEmail(registerReq.Email); err != nil {
		a.lg.LogError(err.Err.Error(), err.Err, api.GetRequestId(w))
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

	ctx := api.BaseContext(w, r, a.lg)
	ctx = config.WrapRedisContext(ctx, a.redisCfg)

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
		api.RequestError(w, a.lg, rParseErr, http.StatusForbidden, errMsg)

		return
	}

	ctx := api.BaseContext(w, r, a.lg)
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
