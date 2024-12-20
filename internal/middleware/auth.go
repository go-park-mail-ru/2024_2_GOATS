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

// SessionMiddleware struct
type SessionMiddleware struct {
	authServ delivery.AuthServiceInterface
}

// NewSessionMiddleware returns an instance of SessionMiddleware
func NewSessionMiddleware(authServ delivery.AuthServiceInterface) *SessionMiddleware {
	return &SessionMiddleware{
		authServ: authServ,
	}
}

// AuthMiddleware checks if user is authorized
func (mw *SessionMiddleware) AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !strings.HasPrefix(r.URL.Path, "/api/room/create") && !strings.HasPrefix(r.URL.Path, "/api/users") && !strings.HasPrefix(r.URL.Path, "/api/movies") && !strings.HasPrefix(r.URL.Path, "/api/subscription") {
			next.ServeHTTP(w, r)
			return
		}

		logger := log.Ctx(r.Context())
		ck, err := r.Cookie("session_id")

		if strings.HasPrefix(r.URL.Path, "/api/movies") && errors.Is(err, http.ErrNoCookie) {
			ctx := context.WithValue(r.Context(), config.CurrentUserKey{}, 0)
			next.ServeHTTP(w, r.WithContext(ctx))
			return
		}

		if errors.Is(err, http.ErrNoCookie) {
			logger.Error().Err(fmt.Errorf("sessionMiddleware: no cookie %w", err)).Msg("no_cookie_err")
			w.WriteHeader(http.StatusForbidden)
			return
		} else if err != nil {
			logger.Error().Err(fmt.Errorf("unexpected error with cookie: %w", err)).Msg("check_cookie_err")
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		sessionSrvResp, errSrvResp := mw.authServ.Session(r.Context(), ck.Value)
		_, errResp := converter.ToAPISessionResponse(sessionSrvResp), errVals.ToDeliveryErrorFromService(errSrvResp)
		if errResp != nil {
			errRespp := errors.New("authorization failed")
			logger.Error().Err(errRespp).Msg("authorization failed")
			api.Response(r.Context(), w, errResp.HTTPStatus, errResp)
			return
		}

		ctx := context.WithValue(r.Context(), config.CurrentUserKey{}, sessionSrvResp.UserData.ID)
		logger.Info().Interface("sessionResp", sessionSrvResp).Msg("authMiddleware success")
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
