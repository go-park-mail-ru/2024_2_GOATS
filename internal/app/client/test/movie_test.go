package client_test

import (
	"context"
	"database/sql"
	"errors"
	"testing"

	"github.com/go-park-mail-ru/2024_2_GOATS/internal/app/client"
	mockMovie "github.com/go-park-mail-ru/2024_2_GOATS/internal/app/client/mocks"
	"github.com/go-park-mail-ru/2024_2_GOATS/internal/app/models"
	movie "github.com/go-park-mail-ru/2024_2_GOATS/movie_service/pkg/movie_v1"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestMovieClient_GetFavorites(t *testing.T) {
	tests := []struct {
		name          string
		movieIDs      []uint64
		mockSetup     func(mock *mockMovie.MockMovieServiceClient)
		expectedResp  []models.MovieShortInfo
		expectedError error
	}{
		{
			name:     "Success",
			movieIDs: []uint64{1, 2},
			mockSetup: func(mock *mockMovie.MockMovieServiceClient) {
				mock.EXPECT().
					GetFavorites(gomock.Any(), &movie.GetFavoritesRequest{MovieIds: []uint64{1, 2}}).
					Return(&movie.GetFavoritesResponse{
						Movies: []*movie.MovieShortInfo{
							{
								Id:          1,
								Title:       "Movie 1",
								CardUrl:     "card_url_1",
								AlbumUrl:    "album_url_1",
								Rating:      8.5,
								ReleaseDate: "2023-01-01",
								MovieType:   "Action",
								Country:     "USA",
							},
							{
								Id:          2,
								Title:       "Movie 2",
								CardUrl:     "card_url_2",
								AlbumUrl:    "album_url_2",
								Rating:      7.5,
								ReleaseDate: "2023-02-01",
								MovieType:   "Comedy",
								Country:     "UK",
							},
						},
					}, nil)
			},
			expectedResp: []models.MovieShortInfo{
				{
					ID:          1,
					Title:       "Movie 1",
					CardURL:     "card_url_1",
					AlbumURL:    "album_url_1",
					Rating:      8.5,
					ReleaseDate: "2023-01-01",
					MovieType:   "Action",
					Country:     "USA",
				},
				{
					ID:          2,
					Title:       "Movie 2",
					CardURL:     "card_url_2",
					AlbumURL:    "album_url_2",
					Rating:      7.5,
					ReleaseDate: "2023-02-01",
					MovieType:   "Comedy",
					Country:     "UK",
				},
			},
			expectedError: nil,
		},
		{
			name:     "Error from MovieService",
			movieIDs: []uint64{1, 2},
			mockSetup: func(mock *mockMovie.MockMovieServiceClient) {
				mock.EXPECT().
					GetFavorites(gomock.Any(), &movie.GetFavoritesRequest{MovieIds: []uint64{1, 2}}).
					Return(nil, errors.New("gRPC error"))
			},
			expectedResp:  nil,
			expectedError: errors.New("gRPC error"),
		},
		{
			name:     "Empty Input",
			movieIDs: []uint64{},
			mockSetup: func(mock *mockMovie.MockMovieServiceClient) {
				mock.EXPECT().
					GetFavorites(gomock.Any(), &movie.GetFavoritesRequest{MovieIds: []uint64{}}).
					Return(&movie.GetFavoritesResponse{Movies: []*movie.MovieShortInfo{}}, nil)
			},
			expectedResp:  []models.MovieShortInfo{},
			expectedError: nil,
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockMovieService := mockMovie.NewMockMovieServiceClient(ctrl)
			test.mockSetup(mockMovieService)

			movieClient := client.NewMovieClient(mockMovieService)
			resp, err := movieClient.GetFavorites(context.Background(), test.movieIDs)

			assert.Equal(t, test.expectedResp, resp)
			if test.expectedError != nil {
				assert.Error(t, err)
				assert.Contains(t, err.Error(), test.expectedError.Error())
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestMovieClient_GetMovie(t *testing.T) {
	tests := []struct {
		name          string
		movieID       int
		mockSetup     func(mock *mockMovie.MockMovieServiceClient)
		expectedResp  *models.MovieInfo
		expectedError error
	}{
		{
			name:    "Success",
			movieID: 1,
			mockSetup: func(mock *mockMovie.MockMovieServiceClient) {
				mock.EXPECT().
					GetMovie(gomock.Any(), &movie.GetMovieRequest{MovieId: 1}).
					Return(&movie.GetMovieResponse{
						Movie: &movie.MovieInfo{
							Id:          1,
							Title:       "Test Movie",
							CardUrl:     "card1",
							AlbumUrl:    "album1",
							Rating:      9.5,
							MovieType:   "Action",
							Country:     "USA",
							ReleaseDate: "2023-01-01",
							DirectorInfo: &movie.DirectorInfo{
								Name:    "Test",
								Surname: "Tester",
							},
						},
					}, nil)
			},
			expectedResp: &models.MovieInfo{
				ID:          1,
				Title:       "Test Movie",
				CardURL:     "card1",
				AlbumURL:    "album1",
				Rating:      9.5,
				MovieType:   "Action",
				Country:     "USA",
				ReleaseDate: "2023-01-01",
				Director: &models.DirectorInfo{
					Person: models.Person{
						Name:    "Test",
						Surname: "Tester",
					},
				},
			},
			expectedError: nil,
		},
		{
			name:    "Error from service",
			movieID: 2,
			mockSetup: func(mock *mockMovie.MockMovieServiceClient) {
				mock.EXPECT().
					GetMovie(gomock.Any(), &movie.GetMovieRequest{MovieId: 2}).
					Return(nil, errors.New("gRPC error"))
			},
			expectedResp:  nil,
			expectedError: errors.New("gRPC error"),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockMovieService := mockMovie.NewMockMovieServiceClient(ctrl)
			test.mockSetup(mockMovieService)

			movieClient := client.NewMovieClient(mockMovieService)
			resp, err := movieClient.GetMovie(context.Background(), test.movieID)

			assert.Equal(t, test.expectedResp, resp)
			if test.expectedError != nil {
				assert.Error(t, err)
				assert.Contains(t, err.Error(), test.expectedError.Error())
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestMovieClient_GetActor(t *testing.T) {
	tests := []struct {
		name          string
		actorID       int
		mockSetup     func(mock *mockMovie.MockMovieServiceClient)
		expectedResp  *models.ActorInfo
		expectedError error
	}{
		{
			name:    "Success",
			actorID: 1,
			mockSetup: func(mock *mockMovie.MockMovieServiceClient) {
				mock.EXPECT().
					GetActor(gomock.Any(), &movie.GetActorRequest{ActorId: 1}).
					Return(&movie.GetActorResponse{
						Actor: &movie.ActorInfo{
							Id:            1,
							Name:          "John",
							Surname:       "Doe",
							Biography:     "Actor Biography",
							SmallPhotoUrl: "small_photo_url",
							BigPhotoUrl:   "big_photo_url",
							Country:       "USA",
							Birthdate:     "1990-01-01",
						},
					}, nil)
			},
			expectedResp: &models.ActorInfo{
				ID: 1,
				Person: models.Person{
					Name:    "John",
					Surname: "Doe",
				},
				Biography:     "Actor Biography",
				SmallPhotoURL: "small_photo_url",
				BigPhotoURL:   "big_photo_url",
				Country:       "USA",
				Birthdate: sql.NullString{
					String: "1990-01-01",
					Valid:  true,
				},
			},
			expectedError: nil,
		},
		{
			name:    "Error from service",
			actorID: 2,
			mockSetup: func(mock *mockMovie.MockMovieServiceClient) {
				mock.EXPECT().
					GetActor(gomock.Any(), &movie.GetActorRequest{ActorId: 2}).
					Return(nil, errors.New("gRPC error"))
			},
			expectedResp:  nil,
			expectedError: errors.New("gRPC error"),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockMovieService := mockMovie.NewMockMovieServiceClient(ctrl)
			test.mockSetup(mockMovieService)

			movieClient := client.NewMovieClient(mockMovieService)
			resp, err := movieClient.GetActor(context.Background(), test.actorID)

			assert.Equal(t, test.expectedResp, resp)
			if test.expectedError != nil {
				assert.Error(t, err)
				assert.Contains(t, err.Error(), test.expectedError.Error())
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestMovieClient_GetCollection(t *testing.T) {
	tests := []struct {
		name          string
		filter        string
		mockSetup     func(mock *mockMovie.MockMovieServiceClient)
		expectedResp  []models.Collection
		expectedError error
	}{
		{
			name:   "Success",
			filter: "Top",
			mockSetup: func(mock *mockMovie.MockMovieServiceClient) {
				mock.EXPECT().
					GetCollections(gomock.Any(), &movie.GetCollectionsRequest{Filter: "Top"}).
					Return(&movie.GetCollectionsResponse{
						Collections: []*movie.Collection{
							{
								Id:    1,
								Title: "Top Collection",
								Movies: []*movie.MovieShortInfo{
									{Id: 1, Title: "Movie 1", CardUrl: "card_url_1", AlbumUrl: "album_url_1", Rating: 8.0, ReleaseDate: "2023-01-01", MovieType: "Action", Country: "USA"},
								},
							},
						},
					}, nil)
			},
			expectedResp: []models.Collection{
				{
					ID:    1,
					Title: "Top Collection",
					Movies: []*models.MovieShortInfo{
						{ID: 1, Title: "Movie 1", CardURL: "card_url_1", AlbumURL: "album_url_1", Rating: 8.0, ReleaseDate: "2023-01-01", MovieType: "Action", Country: "USA"},
					},
				},
			},
			expectedError: nil,
		},
		{
			name:   "Error from service",
			filter: "Unknown",
			mockSetup: func(mock *mockMovie.MockMovieServiceClient) {
				mock.EXPECT().
					GetCollections(gomock.Any(), &movie.GetCollectionsRequest{Filter: "Unknown"}).
					Return(nil, errors.New("gRPC error"))
			},
			expectedResp:  nil,
			expectedError: errors.New("gRPC error"),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockMovieService := mockMovie.NewMockMovieServiceClient(ctrl)
			test.mockSetup(mockMovieService)

			movieClient := client.NewMovieClient(mockMovieService)
			resp, err := movieClient.GetCollection(context.Background(), test.filter)

			assert.Equal(t, test.expectedResp, resp)
			if test.expectedError != nil {
				assert.Error(t, err)
				assert.Contains(t, err.Error(), test.expectedError.Error())
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
