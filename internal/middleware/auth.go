package middleware

import (
	"net/http"

	"github.com/go-park-mail-ru/2024_2_GOATS/internal/app/api"
	"github.com/go-park-mail-ru/2024_2_GOATS/internal/app/logger"
	"github.com/go-park-mail-ru/2024_2_GOATS/internal/app/api/converter"
	"github.com/go-park-mail-ru/2024_2_GOATS/internal/app/auth/delivery"
)

type SessionMiddleware struct {
	authServ delivery.AuthServiceInterface
	lg       *logger.BaseLogger
}

func NewSessionMiddleware(authServ delivery.AuthServiceInterface, lg *logger.BaseLogger) *SessionMiddleware {
	return &SessionMiddleware{
		authServ: authServ,
		lg:       lg,
	}
}

func (mw *SessionMiddleware) AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ck, err := r.Cookie("session_id")

		if err == http.ErrNoCookie {
			mw.lg.LogError("no cookie", err, w.Header().Get("request-id"))
			w.WriteHeader(http.StatusForbidden)

			return
		} else if err != nil {
			mw.lg.LogError("problems with getting cookie", err, w.Header().Get("request-id"))
			w.WriteHeader(http.StatusInternalServerError)
		}

		ctx := api.BaseContext(w, r, mw.lg)
		sessionSrvResp, errSrvResp := mw.authServ.Session(ctx, ck.Value)

		_, errResp := converter.ToApiSessionResponse(sessionSrvResp), converter.ToApiErrorResponse(errSrvResp)
		if errResp != nil {
			api.Response(w, errResp.StatusCode, errResp)
			return
		}

		next.ServeHTTP(w, r)
	})
}
