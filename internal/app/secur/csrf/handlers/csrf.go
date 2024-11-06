package handlers

import (
	token_gen "github.com/go-park-mail-ru/2024_2_GOATS/internal/app/secur/csrf/token"
	"github.com/gorilla/sessions"
	"github.com/rs/zerolog/log"
	"net/http"
)

var store = sessions.NewCookieStore([]byte("secret-key")) // Этот ключ должен быть одинаковым

// GenerateCSRFTokenHandler создает CSRF-токен и сохраняет его в сессии
func GenerateCSRFTokenHandler(w http.ResponseWriter, r *http.Request) {
	token, err := token_gen.GenerateToken()

	if err != nil {
		http.Error(w, "Failed to generate CSRF token", http.StatusInternalServerError)
		log.Info().Msg("Error generating CSRF token")
		return
	}

	session, err := store.Get(r, "session-name")
	if err != nil {
		http.Error(w, "Failed to get session", http.StatusInternalServerError)
		return
	}

	session.Values["csrf_token"] = token
	err = session.Save(r, w)
	if err != nil {
		http.Error(w, "Failed to save session", http.StatusInternalServerError)
		return
	}

	w.Header().Set("X-CSRF-Token", token)

	// Возвращаем успешный ответ
	w.WriteHeader(http.StatusOK)
}
