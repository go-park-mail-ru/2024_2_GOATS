package middleware

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/go-park-mail-ru/2024_2_GOATS/config"
	"github.com/go-park-mail-ru/2024_2_GOATS/internal/app/api"
	"github.com/go-park-mail-ru/2024_2_GOATS/internal/app/api/converter"
	"github.com/go-park-mail-ru/2024_2_GOATS/internal/app/auth/delivery"
	errVals "github.com/go-park-mail-ru/2024_2_GOATS/internal/app/errors"
	"github.com/rs/zerolog/log"
)

type SessionMiddleware struct {
	authServ delivery.AuthServiceInterface
}

func NewSessionMiddleware(authServ delivery.AuthServiceInterface) *SessionMiddleware {
	return &SessionMiddleware{
		authServ: authServ,
	}
}

func (mw *SessionMiddleware) AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if strings.HasPrefix(r.URL.Path, "/api/auth") || strings.HasPrefix(r.URL.Path, "/api/csrf-token") {
			next.ServeHTTP(w, r)
			return
		}

		logger := log.Ctx(r.Context())
		ck, err := r.Cookie("session_id")

		if strings.HasPrefix(r.URL.Path, "/api/movie") && err == http.ErrNoCookie {
			ctx := context.WithValue(r.Context(), config.CurrentUserKey{}, 0)
			next.ServeHTTP(w, r.WithContext(ctx))
		}

		if err == http.ErrNoCookie {
			logger.Error().Err(fmt.Errorf("sessionMiddleware: no cookie %w", err)).Msg("no_cookie_err")
			w.WriteHeader(http.StatusForbidden)

			return
		} else if err != nil {
			logger.Error().Err(fmt.Errorf("problems with getting cookie: %w", err)).Msg("check_cookie_err")
			w.WriteHeader(http.StatusInternalServerError)
		}

		sessionSrvResp, errSrvResp := mw.authServ.Session(r.Context(), ck.Value)

		_, errResp := converter.ToApiSessionResponse(sessionSrvResp), errVals.ToDeliveryErrorFromService(errSrvResp)
		if errResp != nil {
			errMsg := errors.New("failed to authorize")
			logger.Error().Err(errMsg).Interface("sessionResp", errResp).Msg("request_failed")
			api.Response(r.Context(), w, errResp.HTTPStatus, errResp)

			return
		}

		ctx := context.WithValue(r.Context(), config.CurrentUserKey{}, sessionSrvResp.UserData.ID)

		logger.Info().Interface("sessionResp", sessionSrvResp).Msg("authMiddleware success")
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
