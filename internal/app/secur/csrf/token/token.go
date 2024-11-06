package token

import (
	"crypto/rand"
	"encoding/base64"
)

// GenerateToken генерирует случайный CSRF-токен
func GenerateToken() (string, error) {
	token := make([]byte, 32)
	_, err := rand.Read(token)
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(token), nil
}
