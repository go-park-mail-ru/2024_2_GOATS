package delivery

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/go-park-mail-ru/2024_2_GOATS/config"
	"github.com/go-park-mail-ru/2024_2_GOATS/internal/app/api"
	"github.com/go-park-mail-ru/2024_2_GOATS/internal/app/api/converter"
	"github.com/go-park-mail-ru/2024_2_GOATS/internal/app/api/handlers"
	"github.com/go-park-mail-ru/2024_2_GOATS/internal/app/auth/delivery/validation"
	errVals "github.com/go-park-mail-ru/2024_2_GOATS/internal/app/errors"
	"github.com/go-park-mail-ru/2024_2_GOATS/internal/app/models"
)

var _ handlers.AuthImplementationInterface = (*AuthHandler)(nil)

type AuthHandler struct {
	authService AuthServiceInterface
	cfg         *config.Config
}

func NewAuthHandler(authSrv AuthServiceInterface, cfg *config.Config) *AuthHandler {
	return &AuthHandler{
		authService: authSrv,
		cfg:         cfg,
	}
}

func (a *AuthHandler) Logout(w http.ResponseWriter, r *http.Request) {
	ck, err := r.Cookie("session_id")
	if errors.Is(err, http.ErrNoCookie) {
		api.Response(w, http.StatusForbidden,
			preparedDefaultError(
				errVals.ErrNoCookieCode,
				fmt.Errorf("Logout action: No cookie err - %w", err),
			),
		)

		return
	}

	validErr := validation.ValidateCookie(ck.Value)
	if validErr != nil {
		api.Response(w, http.StatusBadRequest,
			preparedDefaultError(
				"cookie_validation_error",
				fmt.Errorf("Logout action: Invalid cookie err - %w", validErr.Err),
			),
		)

		return
	}

	ctx := config.WrapContext(r.Context(), a.cfg)
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
	ctx := config.WrapContext(r.Context(), a.cfg)
	authSrvResp, errSrvResp := a.authService.Login(ctx, loginServData)

	authResp, errResp := converter.ToApiAuthResponse(authSrvResp), converter.ToApiErrorResponse(errSrvResp)
	if errResp != nil {
		api.Response(w, errResp.StatusCode, errResp)
		return
	}

	oldCookie, err := r.Cookie("session_id")
	if errors.Is(err, http.ErrNoCookie) {
		log.Printf("user dont have old cookie")
	}

	if oldCookie != nil && authResp.NewCookie.Token.TokenID != oldCookie.Value {
		http.SetCookie(w, preparedExpiredCookie())
	}

	http.SetCookie(w, preparedCookie(authResp.NewCookie))

	api.Response(w, authResp.StatusCode, authResp)
}

func (a *AuthHandler) Register(w http.ResponseWriter, r *http.Request) {
	registerReq := &api.RegisterRequest{}
	api.DecodeBody(w, r, registerReq)

	errs := make([]errVals.ErrorObj, 0)

	if err := validation.ValidatePassword(registerReq.Password, registerReq.PasswordConfirmation); err != nil {
		addError(errVals.ErrInvalidPasswordCode, *err, &errs)
	}

	if err := validation.ValidateEmail(registerReq.Email); err != nil {
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

	ctx := config.WrapContext(r.Context(), a.cfg)
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
		api.Response(w, http.StatusForbidden,
			preparedDefaultError(
				errVals.ErrNoCookieCode,
				fmt.Errorf("Session action: No cookie err - %w", err),
			),
		)

		return
	}

	ctx := config.WrapContext(r.Context(), a.cfg)
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

func preparedDefaultError(code string, err error) *api.ErrorResponse {
	return &api.ErrorResponse{
		StatusCode: http.StatusForbidden,
		Errors: []errVals.ErrorObj{{
			Code:  code,
			Error: errVals.CustomError{Err: err},
		}},
	}
}

func addError(code string, err errVals.CustomError, errors *[]errVals.ErrorObj) {
	errStruct := errVals.ErrorObj{
		Code:  code,
		Error: err,
	}

	*errors = append(*errors, errStruct)
}
