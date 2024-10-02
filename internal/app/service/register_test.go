package service

import (
	"context"
	errVals "github.com/go-park-mail-ru/2024_2_GOATS/internal/app/errors"
	"github.com/go-park-mail-ru/2024_2_GOATS/internal/app/models"
	authModels "github.com/go-park-mail-ru/2024_2_GOATS/internal/app/models/auth"
	"github.com/go-park-mail-ru/2024_2_GOATS/internal/app/service/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

type getRegister struct {
	result *authModels.Token
	err    *errVals.ErrorObj
	want2  int
}

func TestService_Register(t *testing.T) {
	t.Parallel()

	type fields struct {
		repository RepositoryInterface
	}
	type args struct {
		ctx          context.Context
		registerData *authModels.RegisterData
	}
	tests := []struct {
		name        string
		fields      fields
		args        args
		want        *authModels.AuthResponse
		want1       *models.ErrorResponse
		getRegister *getRegister
	}{
		{
			name: "Register",
			args: args{
				registerData: &authModels.RegisterData{
					Email:                "string",
					Username:             "string",
					Password:             "string",
					PasswordConfirmation: "string",
				},
			},
			getRegister: &getRegister{result: &authModels.Token{
				UserID:  1,
				TokenID: "string",
			}, want2: 1},
			want: nil,
			want1: &models.ErrorResponse{
				Success: false,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			repository := &mocks.RepositoryInterfaceMock{}

			repository.EXPECT().Login(mock.Anything, mock.Anything).Return(tt.getRegister.result, tt.getRegister.err, tt.getRegister.want2)

			s := NewService(repository)

			got, got1 := s.Register(tt.args.ctx, tt.args.registerData)

			assert.Equalf(t, tt.want, got, "Register()")
			assert.Equalf(t, tt.want1.Success, got1.Success, "Register()")
		})
	}
}
