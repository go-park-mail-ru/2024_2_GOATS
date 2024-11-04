package handlers

import (
	token_gen "github.com/go-park-mail-ru/2024_2_GOATS/internal/app/secur/csrf/token"
	"net/http"
)

// Создаем CSRF-токен и отправляет его клиенту
func GenerateCSRFTokenHandler(w http.ResponseWriter, r *http.Request) {
	token, err := token_gen.GenerateToken()
	if err != nil {
		http.Error(w, "Failed to generate CSRF token", http.StatusInternalServerError)
		return
	}

	token_gen.SetCSRFTokenCookie(w, token)
	w.Write([]byte("CSRF token generated"))
}
