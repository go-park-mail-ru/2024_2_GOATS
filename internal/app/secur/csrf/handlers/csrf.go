package handlers

import (
	errVals "github.com/go-park-mail-ru/2024_2_GOATS/internal/app/errors"
	token_gen "github.com/go-park-mail-ru/2024_2_GOATS/internal/app/secur/csrf/token"
	"github.com/rs/zerolog/log"
	"net/http"
)

// GenerateCSRFTokenHandler создает CSRF-токен и отправляет его клиенту
func GenerateCSRFTokenHandler(w http.ResponseWriter, r *http.Request) {
	token, err := token_gen.GenerateToken()
	if err != nil {
		http.Error(w, "Failed to generate CSRF token", http.StatusInternalServerError)
		log.Info().Msg(errVals.ErrGenCSRF)

		return
	}

	token_gen.SetCSRFTokenCookie(w, token)

	w.WriteHeader(http.StatusOK)
}
