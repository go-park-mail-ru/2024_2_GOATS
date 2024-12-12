package delivery

import (
	"context"
	"database/sql"
	"errors"
	"testing"

	mockService "github.com/go-park-mail-ru/2024_2_GOATS/movie_service/internal/movie/delivery/mocks"
	"github.com/go-park-mail-ru/2024_2_GOATS/movie_service/internal/movie/models"
	movie "github.com/go-park-mail-ru/2024_2_GOATS/movie_service/pkg/movie_v1"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func TestMovieHandler_GetMovieByGenre(t *testing.T) {
	tests := []struct {
		name          string
		req           *movie.GetMovieByGenreRequest
		mockSetup     func(mockService *mockService.MockMovieServiceInterface)
		expectedResp  *movie.GetMovieByGenreResponse
		expectedError error
	}{
		{
			name: "Success",
			req:  &movie.GetMovieByGenreRequest{Genre: "Action"},
			mockSetup: func(mockService *mockService.MockMovieServiceInterface) {
				mockService.EXPECT().
					GetMovieByGenre(gomock.Any(), "Action").
					Return([]models.MovieShortInfo{
						{
							ID:          1,
							CardURL:     "card_url_1",
							MovieType:   "Action",
							AlbumURL:    "album_url_1",
							Title:       "Movie 1",
							Country:     "USA",
							ReleaseDate: "2023-01-01",
							Rating:      8.5,
						},
						{
							ID:          2,
							CardURL:     "card_url_2",
							MovieType:   "Action",
							AlbumURL:    "album_url_2",
							Title:       "Movie 2",
							Country:     "UK",
							ReleaseDate: "2023-02-01",
							Rating:      7.5,
						},
					}, nil)
			},
			expectedResp: &movie.GetMovieByGenreResponse{
				Movies: []*movie.MovieShortInfo{
					{
						Id:          1,
						CardUrl:     "card_url_1",
						MovieType:   "Action",
						AlbumUrl:    "album_url_1",
						Title:       "Movie 1",
						Country:     "USA",
						ReleaseDate: "2023-01-01",
						Rating:      8.5,
					},
					{
						Id:          2,
						CardUrl:     "card_url_2",
						MovieType:   "Action",
						AlbumUrl:    "album_url_2",
						Title:       "Movie 2",
						Country:     "UK",
						ReleaseDate: "2023-02-01",
						Rating:      7.5,
					},
				},
			},
			expectedError: nil,
		},
		{
			name:          "Missing Genre",
			req:           &movie.GetMovieByGenreRequest{Genre: ""},
			mockSetup:     func(_ *mockService.MockMovieServiceInterface) {},
			expectedResp:  nil,
			expectedError: status.Error(codes.InvalidArgument, "genre is required"),
		},
		{
			name: "Service Error",
			req:  &movie.GetMovieByGenreRequest{Genre: "Drama"},
			mockSetup: func(mockService *mockService.MockMovieServiceInterface) {
				mockService.EXPECT().
					GetMovieByGenre(gomock.Any(), "Drama").
					Return(nil, errors.New("internal error"))
			},
			expectedResp:  nil,
			expectedError: errors.New("internal error"),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockService := mockService.NewMockMovieServiceInterface(ctrl)
			test.mockSetup(mockService)

			handler := NewMovieHandler(mockService)
			resp, err := handler.GetMovieByGenre(context.Background(), test.req)

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

func TestMovieHandler_GetMovie(t *testing.T) {
	tests := []struct {
		name          string
		req           *movie.GetMovieRequest
		mockSetup     func(mockService *mockService.MockMovieServiceInterface)
		expectedResp  *movie.GetMovieResponse
		expectedError error
	}{
		{
			name: "Success",
			req:  &movie.GetMovieRequest{MovieId: 1},
			mockSetup: func(mockService *mockService.MockMovieServiceInterface) {
				mockService.EXPECT().
					GetMovie(gomock.Any(), 1).
					Return(&models.MovieInfo{
						ID:               1,
						CardURL:          "card_url_1",
						AlbumURL:         "album_url_1",
						Rating:           8.5,
						Title:            "Test Movie",
						MovieType:        "Action",
						Country:          "USA",
						ReleaseDate:      "2023-01-01",
						IsFavorite:       true,
						VideoURL:         "video_url",
						Director:         &models.DirectorInfo{Person: models.Person{Name: "John", Surname: "Doe"}},
						FullDescription:  "Full Description",
						ShortDescription: "Short Description",
						TitleURL:         "title_url",
						Actors: []*models.ActorInfo{
							{
								ID: 1,
								Person: models.Person{
									Name:    "Actor Name",
									Surname: "Actor Surname",
								},
								Biography:     "Biography",
								Post:          "Post",
								Birthdate:     sql.NullString{String: "1990-01-01", Valid: true},
								SmallPhotoURL: "small_photo_url",
								BigPhotoURL:   "big_photo_url",
								Country:       "USA",
							},
						},
						Seasons: []*models.Season{
							{
								SeasonNumber: 1,
								Episodes: []*models.Episode{
									{
										ID:            1,
										Description:   "Episode 1 Description",
										EpisodeNumber: 1,
										Title:         "Episode 1",
										Rating:        9.0,
										ReleaseDate:   "2023-01-01",
										VideoURL:      "video_url_1",
										PreviewURL:    "preview_url_1",
									},
								},
							},
						},
					}, nil)
			},
			expectedResp: &movie.GetMovieResponse{
				Movie: &movie.MovieInfo{
					Id:               1,
					CardUrl:          "card_url_1",
					AlbumUrl:         "album_url_1",
					Rating:           8.5,
					Title:            "Test Movie",
					MovieType:        "Action",
					Country:          "USA",
					ReleaseDate:      "2023-01-01",
					IsFavorite:       true,
					VideoUrl:         "video_url",
					DirectorInfo:     &movie.DirectorInfo{Name: "John", Surname: "Doe"},
					FullDescription:  "Full Description",
					ShortDescription: "Short Description",
					TitleUrl:         "title_url",
					ActorsInfo: []*movie.ActorInfo{
						{
							Id:            1,
							Name:          "Actor Name",
							Surname:       "Actor Surname",
							Biography:     "Biography",
							Post:          "Post",
							Birthdate:     "1990-01-01",
							SmallPhotoUrl: "small_photo_url",
							BigPhotoUrl:   "big_photo_url",
							Country:       "USA",
						},
					},
					Seasons: []*movie.Season{
						{
							SeasonNumber: 1,
							Episodes: []*movie.Episode{
								{
									Id:            1,
									Description:   "Episode 1 Description",
									EpisodeNumber: 1,
									Title:         "Episode 1",
									Rating:        9.0,
									ReleaseDate:   "2023-01-01",
									VideoURL:      "video_url_1",
									PreviewURL:    "preview_url_1",
								},
							},
						},
					},
				},
			},
			expectedError: nil,
		},
		{
			name:          "Invalid Movie ID",
			req:           &movie.GetMovieRequest{MovieId: 0},
			mockSetup:     func(_ *mockService.MockMovieServiceInterface) {},
			expectedResp:  nil,
			expectedError: status.Error(codes.InvalidArgument, "invalid movie ID"),
		},
		{
			name: "Service Error",
			req:  &movie.GetMovieRequest{MovieId: 1},
			mockSetup: func(mockService *mockService.MockMovieServiceInterface) {
				mockService.EXPECT().
					GetMovie(gomock.Any(), 1).
					Return(nil, errors.New("internal error"))
			},
			expectedResp:  nil,
			expectedError: errors.New("internal error"),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockService := mockService.NewMockMovieServiceInterface(ctrl)
			test.mockSetup(mockService)

			handler := NewMovieHandler(mockService)
			resp, err := handler.GetMovie(context.Background(), test.req)

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

func TestMovieHandler_GetActor(t *testing.T) {
	tests := []struct {
		name          string
		req           *movie.GetActorRequest
		mockSetup     func(mockService *mockService.MockMovieServiceInterface)
		expectedResp  *movie.GetActorResponse
		expectedError error
	}{
		{
			name: "Success",
			req:  &movie.GetActorRequest{ActorId: 1},
			mockSetup: func(mockService *mockService.MockMovieServiceInterface) {
				mockService.EXPECT().
					GetActor(gomock.Any(), 1).
					Return(&models.ActorInfo{
						ID:            1,
						Person:        models.Person{Name: "John", Surname: "Doe"},
						Biography:     "Actor Biography",
						Post:          "Lead Actor",
						Birthdate:     sql.NullString{String: "1990-01-01", Valid: true},
						SmallPhotoURL: "small_photo_url",
						BigPhotoURL:   "big_photo_url",
						Country:       "USA",
						Movies: []*models.MovieShortInfo{
							{
								ID:          1,
								Title:       "Movie 1",
								CardURL:     "card_url_1",
								AlbumURL:    "album_url_1",
								Rating:      8.0,
								ReleaseDate: "2023-01-01",
								MovieType:   "Action",
								Country:     "USA",
							},
						},
					}, nil)
			},
			expectedResp: &movie.GetActorResponse{
				Actor: &movie.ActorInfo{
					Id:            1,
					Name:          "John",
					Surname:       "Doe",
					Biography:     "Actor Biography",
					Post:          "Lead Actor",
					Birthdate:     "1990-01-01",
					SmallPhotoUrl: "small_photo_url",
					BigPhotoUrl:   "big_photo_url",
					Country:       "USA",
					Movies: []*movie.MovieShortInfo{
						{
							Id:          1,
							Title:       "Movie 1",
							CardUrl:     "card_url_1",
							AlbumUrl:    "album_url_1",
							Rating:      8.0,
							ReleaseDate: "2023-01-01",
							MovieType:   "Action",
							Country:     "USA",
						},
					},
				},
			},
			expectedError: nil,
		},
		{
			name:          "Invalid Actor ID",
			req:           &movie.GetActorRequest{ActorId: 0},
			mockSetup:     func(_ *mockService.MockMovieServiceInterface) {},
			expectedResp:  nil,
			expectedError: status.Error(codes.InvalidArgument, "invalid actor ID"),
		},
		{
			name: "Service Error",
			req:  &movie.GetActorRequest{ActorId: 1},
			mockSetup: func(mockService *mockService.MockMovieServiceInterface) {
				mockService.EXPECT().
					GetActor(gomock.Any(), 1).
					Return(nil, errors.New("internal error"))
			},
			expectedResp:  nil,
			expectedError: errors.New("internal error"),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockService := mockService.NewMockMovieServiceInterface(ctrl)
			test.mockSetup(mockService)

			handler := NewMovieHandler(mockService)
			resp, err := handler.GetActor(context.Background(), test.req)

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

func TestMovieHandler_GetCollections(t *testing.T) {
	tests := []struct {
		name          string
		req           *movie.GetCollectionsRequest
		mockSetup     func(mockService *mockService.MockMovieServiceInterface)
		expectedResp  *movie.GetCollectionsResponse
		expectedError error
	}{
		{
			name: "Success",
			req:  &movie.GetCollectionsRequest{Filter: "Top"},
			mockSetup: func(mockService *mockService.MockMovieServiceInterface) {
				mockService.EXPECT().
					GetCollection(gomock.Any(), "Top").
					Return(&models.CollectionsRespData{
						Collections: []models.Collection{
							{
								ID:    1,
								Title: "Top Movies",
								Movies: []*models.MovieShortInfo{
									{
										ID:          1,
										Title:       "Movie 1",
										CardURL:     "card_url_1",
										AlbumURL:    "album_url_1",
										Rating:      8.0,
										ReleaseDate: "2023-01-01",
										MovieType:   "Action",
										Country:     "USA",
									},
								},
							},
							{
								ID:    2,
								Title: "New Releases",
								Movies: []*models.MovieShortInfo{
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
							},
						},
					}, nil)
			},
			expectedResp: &movie.GetCollectionsResponse{
				Collections: []*movie.Collection{
					{
						Id:    1,
						Title: "Top Movies",
						Movies: []*movie.MovieShortInfo{
							{
								Id:          1,
								Title:       "Movie 1",
								CardUrl:     "card_url_1",
								AlbumUrl:    "album_url_1",
								Rating:      8.0,
								ReleaseDate: "2023-01-01",
								MovieType:   "Action",
								Country:     "USA",
							},
						},
					},
					{
						Id:    2,
						Title: "New Releases",
						Movies: []*movie.MovieShortInfo{
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
					},
				},
			},
			expectedError: nil,
		},
		{
			name: "Service Error",
			req:  &movie.GetCollectionsRequest{Filter: "Popular"},
			mockSetup: func(mockService *mockService.MockMovieServiceInterface) {
				mockService.EXPECT().
					GetCollection(gomock.Any(), "Popular").
					Return(nil, errors.New("internal error"))
			},
			expectedResp:  nil,
			expectedError: errors.New("internal error"),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockService := mockService.NewMockMovieServiceInterface(ctrl)
			test.mockSetup(mockService)

			handler := NewMovieHandler(mockService)
			resp, err := handler.GetCollections(context.Background(), test.req)

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

func TestMovieHandler_GetFavorites(t *testing.T) {
	tests := []struct {
		name          string
		req           *movie.GetFavoritesRequest
		mockSetup     func(mockService *mockService.MockMovieServiceInterface)
		expectedResp  *movie.GetFavoritesResponse
		expectedError error
	}{
		{
			name: "Success",
			req:  &movie.GetFavoritesRequest{MovieIds: []uint64{1, 2}},
			mockSetup: func(mockService *mockService.MockMovieServiceInterface) {
				mockService.EXPECT().
					GetFavorites(gomock.Any(), []uint64{1, 2}).
					Return([]*models.MovieShortInfo{
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
					}, nil)
			},
			expectedResp: &movie.GetFavoritesResponse{
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
			},
			expectedError: nil,
		},
		{
			name: "Empty MovieIds",
			req:  &movie.GetFavoritesRequest{MovieIds: []uint64{}},
			mockSetup: func(mockService *mockService.MockMovieServiceInterface) {
				mockService.EXPECT().
					GetFavorites(gomock.Any(), []uint64{}).
					Return([]*models.MovieShortInfo{}, nil)
			},
			expectedResp: &movie.GetFavoritesResponse{
				Movies: []*movie.MovieShortInfo(nil),
			},
			expectedError: nil,
		},
		{
			name: "Service Error",
			req:  &movie.GetFavoritesRequest{MovieIds: []uint64{1, 2}},
			mockSetup: func(mockService *mockService.MockMovieServiceInterface) {
				mockService.EXPECT().
					GetFavorites(gomock.Any(), []uint64{1, 2}).
					Return(nil, errors.New("internal error"))
			},
			expectedResp:  nil,
			expectedError: errors.New("internal error"),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockService := mockService.NewMockMovieServiceInterface(ctrl)
			test.mockSetup(mockService)

			handler := NewMovieHandler(mockService)
			resp, err := handler.GetFavorites(context.Background(), test.req)

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
