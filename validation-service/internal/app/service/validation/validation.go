package validation

import (
	api "github.com/go-park-mail-ru/2024_2_GOATS/validation-service/internal/app/api/validation"
)

var _ api.ValidationService = (*serv)(nil)

type serv struct{}

func NewService() *serv {
	return &serv{}
}
