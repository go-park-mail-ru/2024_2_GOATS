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

type getLogin struct {
	result *authModels.Token
	err    *errVals.ErrorObj
	want2  int
}

func TestService_Login1(t *testing.T) {
	t.Parallel()

	type fields struct {
		repository RepositoryInterface
	}
	type args struct {
		ctx       context.Context
		loginData *authModels.LoginData
	}
	tests := []struct {
		name     string
		fields   fields
		args     args
		want     *authModels.AuthResponse
		want1    *models.ErrorResponse
		getLogin *getLogin
	}{
		{
			name:     "ok",
			getLogin: &getLogin{result: &authModels.Token{}},
			want: &authModels.AuthResponse{
				Success: true,
				Token: &authModels.Token{
					UserID:  1,
					TokenID: "string",
				},
			},
			want1: nil,
		},

		{
			name:     "error",
			getLogin: &getLogin{result: &authModels.Token{}, err: &errVals.ErrorObj{Code: "1"}},
			want: &authModels.AuthResponse{
				Success: true,
				Token: &authModels.Token{
					UserID:  1,
					TokenID: "string",
				},
			},
			want1: &models.ErrorResponse{Success: false, StatusCode: 0, Errors: []errVals.ErrorObj{
				{
					Code: "1",
				},
			}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			repository := &mocks.RepositoryInterfaceMock{}

			repository.EXPECT().Login(mock.Anything, mock.Anything).Return(tt.getLogin.result, tt.getLogin.err, tt.getLogin.want2)

			s := NewService(repository)

			got, got2 := s.Login(tt.args.ctx, tt.args.loginData)
			if got2 != nil {
				assert.Equal(t, tt.want1, got2)
			}
			if got != nil {
				assert.Equalf(t, tt.want.Success, got.Success, "Login()")
			}
		})
	}
}
