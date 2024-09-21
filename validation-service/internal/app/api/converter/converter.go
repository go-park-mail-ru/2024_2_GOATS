package converter

import (
	"github.com/go-park-mail-ru/2024_2_GOATS/validation-service/internal/app/model"
	desc "github.com/go-park-mail-ru/2024_2_GOATS/validation-service/internal/pb/validation"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func ToUserRegisterDataFromDesc(data *desc.ValidateRegistrationRequest) *model.UserRegisterData {
	birthday := data.GetBitrhdate().AsTime()
	ts := timestamppb.New(birthday)

	return &model.UserRegisterData{
		Email:           data.GetEmail(),
		Password:        data.GetPassword(),
		PasswordConfirm: data.GetPasswordConfirm(),
		Sex:             data.GetSex(),
		Birthday:        int(ts.Seconds),
	}
}

func ToErrorsFromServ(data *model.ErrorResponse) *desc.ErrorMessage {
	return &desc.ErrorMessage{
		Code:  data.Code,
		Error: data.ErrorObj.Error(),
	}
}
