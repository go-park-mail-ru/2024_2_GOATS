package middleware

import (
	"fmt"
	"net/http"

	"github.com/go-park-mail-ru/2024_2_GOATS/internal/app/api"
	"github.com/go-park-mail-ru/2024_2_GOATS/internal/app/api/converter"
	"github.com/go-park-mail-ru/2024_2_GOATS/internal/app/auth/delivery"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

type SessionMiddleware struct {
	authServ delivery.AuthServiceInterface
}

func NewSessionMiddleware(authServ delivery.AuthServiceInterface, lg *zerolog.Logger) *SessionMiddleware {
	return &SessionMiddleware{
		authServ: authServ,
	}
}

func (mw *SessionMiddleware) AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		lg := log.Ctx(r.Context())
		ck, err := r.Cookie("session_id")

		if err == http.ErrNoCookie {
			lg.Error().Msg("no cookie")
			w.WriteHeader(http.StatusForbidden)

			return
		} else if err != nil {
			lg.Error().Msg(fmt.Sprintf("problems with getting cookie: %v", err))
			w.WriteHeader(http.StatusInternalServerError)
		}

		sessionSrvResp, errSrvResp := mw.authServ.Session(r.Context(), ck.Value)

		_, errResp := converter.ToApiSessionResponse(sessionSrvResp), converter.ToApiErrorResponse(errSrvResp)
		if errResp != nil {
			api.Response(r.Context(), w, errResp.StatusCode, errResp)
			return
		}

		next.ServeHTTP(w, r)
	})
}
