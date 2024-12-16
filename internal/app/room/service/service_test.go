package service

// import (
// 	"context"
// 	"errors"
// 	model "github.com/go-park-mail-ru/2024_2_GOATS/internal/app/models"
// 	models "github.com/go-park-mail-ru/2024_2_GOATS/internal/app/room/model"
// 	"github.com/stretchr/testify/assert"
// 	"go.uber.org/mock/gomock"
// 	"testing"
// )

// type serviceMock struct {
// 	movieService   *MockMovieServiceInterface
// 	roomRepository *MockRoomRepositoryInterface

// 	s *RoomService
// }

// func newServiceMock(t *testing.T) *serviceMock {
// 	ctrl := gomock.NewController(t)

// 	movieService := NewMockMovieServiceInterface(ctrl)
// 	roomRepository := NewMockRoomRepositoryInterface(ctrl)

// 	return &serviceMock{
// 		movieService:   movieService,
// 		roomRepository: roomRepository,

// 		s: &RoomService{
// 			movieService:   movieService,
// 			roomRepository: roomRepository,
// 		},
// 	}
// }

// func TestNewExampleService(t *testing.T) {
// 	t.Parallel()

// 	t.Run("успех", func(t *testing.T) {
// 		t.Parallel()

// 		m := newServiceMock(t)

// 		s := NewService(
// 			m.roomRepository,
// 			m.movieService,
// 		)

// 		assert.Equal(t, m.s, s)
// 	})
// }

// func TestRoomService_GetRoomState(t *testing.T) {

// 	roomIdTest := 1

// 	type in struct {
// 		roomID string
// 	}

// 	type out struct {
// 		roomState *models.RoomState
// 		err       error
// 	}

// 	tests := []struct {
// 		name   string
// 		in     in
// 		setup  func(m *serviceMock)
// 		assert func(t *testing.T, o out)
// 	}{
// 		{
// 			name: "успех",
// 			in: in{
// 				roomID: "1",
// 			},
// 			setup: func(m *serviceMock) {

// 				m.roomRepository.EXPECT().GetRoomState(
// 					gomock.Any(), "1").
// 					Return(&models.RoomState{
// 						ID: "1",
// 					}, nil)

// 				m.movieService.EXPECT().GetMovie(
// 					gomock.Any(), 0).
// 					Return(&model.MovieInfo{
// 						ID: 1,
// 					}, nil)
// 			},
// 			assert: func(t *testing.T, o out) {
// 				assert.NoError(t, o.err)
// 				assert.Equal(t, roomIdTest, o.roomState.Movie.ID)
// 			},
// 		},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			t.Parallel()

// 			m := newServiceMock(t)

// 			tt.setup(m)

// 			got, err := m.s.GetRoomState(context.Background(), tt.in.roomID)

// 			tt.assert(t, out{
// 				roomState: got,
// 				err:       err,
// 			})
// 		})
// 	}
// }

// func TestRoomService_HandleAction(t *testing.T) {

// 	terr := errors.New("test error")

// 	type in struct {
// 		roomID string
// 		action models.Action
// 	}

// 	type out struct {
// 		err error
// 	}

// 	tests := []struct {
// 		name   string
// 		in     in
// 		setup  func(m *serviceMock)
// 		assert func(t *testing.T, o out)
// 	}{
// 		{
// 			name: "ошибка",
// 			in: in{
// 				roomID: "1",
// 			},
// 			setup: func(m *serviceMock) {
// 				m.roomRepository.EXPECT().GetRoomState(
// 					gomock.Any(), "1").
// 					Return(&models.RoomState{
// 						ID: "1",
// 					}, terr)
// 			},
// 			assert: func(t *testing.T, o out) {
// 				assert.Error(t, o.err)
// 				assert.ErrorIs(t, o.err, terr)
// 			},
// 		},

// 		{
// 			name: "ошибка2",
// 			in: in{
// 				roomID: "1",
// 			},
// 			setup: func(m *serviceMock) {
// 				m.roomRepository.EXPECT().GetRoomState(
// 					gomock.Any(), "1").
// 					Return(&models.RoomState{
// 						ID: "1",
// 					}, nil)

// 				m.roomRepository.EXPECT().UpdateRoomState(
// 					gomock.Any(), "1", &models.RoomState{
// 						ID: "1",
// 					}).
// 					Return(nil)
// 			},
// 			assert: func(t *testing.T, o out) {
// 				assert.NoError(t, o.err)
// 			},
// 		},

// 		{
// 			name: "case_pause",
// 			in: in{
// 				roomID: "1",
// 				action: models.Action{Name: "pause"},
// 			},
// 			setup: func(m *serviceMock) {
// 				m.roomRepository.EXPECT().GetRoomState(
// 					gomock.Any(), "1").
// 					Return(&models.RoomState{
// 						ID: "1",
// 					}, nil)

// 				m.roomRepository.EXPECT().UpdateRoomState(
// 					gomock.Any(), "1", &models.RoomState{
// 						ID:     "1",
// 						Status: "paused",
// 					}).
// 					Return(nil)
// 			},
// 			assert: func(t *testing.T, o out) {
// 				assert.NoError(t, o.err)
// 			},
// 		},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			t.Parallel()

// 			m := newServiceMock(t)

// 			tt.setup(m)

// 			err := m.s.HandleAction(context.Background(), tt.in.roomID, tt.in.action)

// 			tt.assert(t, out{
// 				err: err,
// 			})
// 		})
// 	}
// }

// ////

// //func (s *RoomService) Session(ctx context.Context, cookie string) (*models.SessionRespData, *models.ErrorRespData) {
