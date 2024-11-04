package token

import (
	"crypto/rand"
	"encoding/base64"
	"net/http"
)

// Генерирует случайный CSRF-токен
func GenerateToken() (string, error) {
	token := make([]byte, 32)
	_, err := rand.Read(token)
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(token), nil
}

// Сохраняет CSRF-токен в cookie
func SetCSRFTokenCookie(w http.ResponseWriter, token string) {
	http.SetCookie(w, &http.Cookie{
		Name:     "csrf_token",
		Value:    token,
		Path:     "/",
		HttpOnly: true,
		Secure:   true, // Для прода
	})
}
