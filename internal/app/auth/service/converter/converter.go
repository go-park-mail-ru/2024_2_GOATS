package converter

import (
	"github.com/go-park-mail-ru/2024_2_GOATS/internal/app/models"
	"github.com/go-park-mail-ru/2024_2_GOATS/internal/app/user/repository/dto"
)

func ToDBRegisterFromRegister(rd *models.RegisterData) *dto.DBRegisterData {
	if rd == nil {
		return nil
	}

	return &dto.DBRegisterData{
		Email:                rd.Email,
		Username:             rd.Username,
		Password:             rd.Password,
		PasswordConfirmation: rd.PasswordConfirmation,
	}
}
