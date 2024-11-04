package middleware

import (
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/go-park-mail-ru/2024_2_GOATS/internal/app/api"
	"github.com/go-park-mail-ru/2024_2_GOATS/internal/app/api/converter"
	"github.com/go-park-mail-ru/2024_2_GOATS/internal/app/auth/delivery"
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
		if !strings.HasPrefix(r.URL.Path, "/api/users") {
			next.ServeHTTP(w, r)
			return
		}

		lg := log.Ctx(r.Context())
		ck, err := r.Cookie("session_id")

		if err == http.ErrNoCookie {
			lg.Error().Err(fmt.Errorf("sessionMiddleware: no cookie %w", err)).Msg("no_cookie_err")
			w.WriteHeader(http.StatusForbidden)

			return
		} else if err != nil {
			lg.Error().Err(fmt.Errorf("problems with getting cookie: %w", err)).Msg("check_cookie_err")
			w.WriteHeader(http.StatusInternalServerError)
		}

		sessionSrvResp, errSrvResp := mw.authServ.Session(r.Context(), ck.Value)

		_, errResp := converter.ToApiSessionResponse(sessionSrvResp), converter.ToApiErrorResponse(errSrvResp)
		if errResp != nil {
			errMsg := errors.New("failed to authorize")
			lg.Error().Err(errMsg).Interface("sessionResp", errResp).Msg("request_failed")
			api.Response(r.Context(), w, errResp.StatusCode, errResp)

			return
		}

		lg.Info().Interface("sessionResp", sessionSrvResp).Msg("authMiddleware success")
		next.ServeHTTP(w, r)
	})
}
