package service

import (
	"context"
	errVals "github.com/go-park-mail-ru/2024_2_GOATS/internal/app/errors"
	"github.com/go-park-mail-ru/2024_2_GOATS/internal/app/models"
	"github.com/go-park-mail-ru/2024_2_GOATS/internal/app/service/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

type getGetCollection struct {
	result []models.Collection
	err    *errVals.ErrorObj
	want2  int
}

func TestService_GetCollection(t *testing.T) {
	t.Parallel()

	type fields struct {
		repository RepositoryInterface
	}
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name             string
		fields           fields
		args             args
		want             *models.CollectionsResponse
		want1            *models.ErrorResponse
		getGetCollection *getGetCollection
	}{
		{
			name: "ok",
			getGetCollection: &getGetCollection{result: []models.Collection{
				{
					Id:     1,
					Title:  "string",
					Movies: []*models.Movie{},
				},
			}},
			want: &models.CollectionsResponse{
				Success: true,
				Collections: []models.Collection{
					{
						Id:     1,
						Title:  "string",
						Movies: []*models.Movie{},
					},
				},
			},
			want1: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			repository := &mocks.RepositoryInterfaceMock{}

			repository.EXPECT().GetCollection(mock.Anything).Return(tt.getGetCollection.result, tt.getGetCollection.err, tt.getGetCollection.want2)

			s := NewService(repository)

			got, got1 := s.GetCollection(tt.args.ctx)
			if got != nil {
				assert.Equalf(t, tt.want, got, "GetDictionaryValuesInfoByIDs()")
				assert.Equalf(t, tt.want1, got1, "GetDictionaryValuesInfoByIDs()")
			} else {
				assert.Equalf(t, tt.want1, got1, "GetDictionaryValuesInfoByIDs()")
			}
		})
	}
}
