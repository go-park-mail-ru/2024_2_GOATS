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

type getSession struct {
	result *models.User
	err    *errVals.ErrorObj
	want2  int
}

func TestService_Session(t *testing.T) {
	t.Parallel()

	type fields struct {
		repository RepositoryInterface
	}
	type args struct {
		ctx    context.Context
		cookie string
	}
	tests := []struct {
		name       string
		fields     fields
		args       args
		want       *authModels.SessionResponse
		want1      *models.ErrorResponse
		getSession *getSession
	}{
		{
			name:   "ok",
			fields: fields{},
			getSession: &getSession{
				result: &models.User{},
				err:    &errVals.ErrorObj{},
				want2:  1,
			},
			want1: &models.ErrorResponse{
				Success: false,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			repository := &mocks.RepositoryInterfaceMock{}

			repository.EXPECT().Session(mock.Anything, mock.Anything).Return(tt.getSession.result, tt.getSession.err, tt.getSession.want2)

			s := NewService(repository)

			got, got1 := s.Session(tt.args.ctx, tt.args.cookie)
			assert.Equalf(t, tt.want, got, "Session(%v, %v)", tt.args.ctx, tt.args.cookie)
			assert.Equalf(t, tt.want1.Success, got1.Success, "Session(%v, %v)", tt.args.ctx, tt.args.cookie)
		})
	}
}
