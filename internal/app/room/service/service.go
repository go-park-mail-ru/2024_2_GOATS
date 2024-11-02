package service

import (
	"context"
	"fmt"
	errVals "github.com/go-park-mail-ru/2024_2_GOATS/internal/app/errors"
	model "github.com/go-park-mail-ru/2024_2_GOATS/internal/app/models"
	models "github.com/go-park-mail-ru/2024_2_GOATS/internal/app/room/model"

	"log"
)

//go:generate mockgen -source=service.go -destination=service_mock.go -package=service

type MovieServiceInterface interface {
	GetCollection(ctx context.Context) (*model.CollectionsRespData, *model.ErrorRespData)
	GetMovie(ctx context.Context, mvId int) (*model.MovieInfo, *model.ErrorRespData)
	GetActor(ctx context.Context, actorId int) (*model.StaffInfo, *model.ErrorRespData)
}

type RoomRepositoryInterface interface {
	CreateRoom(ctx context.Context, room *models.RoomState) (*models.RoomState, error)
	UpdateRoomState(ctx context.Context, roomID string, state *models.RoomState) error
	GetRoomState(ctx context.Context, roomID string) (*models.RoomState, error)
	GetFromCookie(ctx context.Context, cookie string) (string, *errVals.ErrorObj, int)
	UserById(ctx context.Context, userId string) (*models.User, *errVals.ErrorObj, int)
}

type RoomService struct {
	roomRepository RoomRepositoryInterface
	movieService   MovieServiceInterface
}

func NewService(repo RoomRepositoryInterface, movieService MovieServiceInterface) *RoomService {
	return &RoomService{
		roomRepository: repo,
		movieService:   movieService,
	}
}

func (s *RoomService) CreateRoom(ctx context.Context, room *models.RoomState) (*models.RoomState, error) {
	return s.roomRepository.CreateRoom(ctx, room)
}

func (s *RoomService) HandleAction(ctx context.Context, roomID string, action models.Action) error {
	roomState, err := s.roomRepository.GetRoomState(ctx, roomID)
	if err != nil {
		return err
	}

	log.Println("roomState.Status === ", roomState.Status)
	log.Println("action.Name === ", action.Name)
	log.Println("action.TimeCode === ", action.TimeCode)

	switch action.Name {
	case "pause":
		roomState.Status = "paused"
		roomState.TimeCode = action.TimeCode
	case "play":
		roomState.Status = "playing"
	case "rewind":
		roomState.TimeCode = action.TimeCode
	case "timer":
		roomState.TimeCode = action.TimeCode
	case "message":
		roomState.Message = action.Message
	}

	return s.roomRepository.UpdateRoomState(ctx, roomID, roomState)
}

func (s *RoomService) GetRoomState(ctx context.Context, roomID string) (*models.RoomState, error) {
	roomState, err := s.roomRepository.GetRoomState(ctx, roomID)
	log.Println("GetRoomStateGetRoomStateGetRoomStateGetRoomState", roomState)

	movie, errMovie := s.movieService.GetMovie(ctx, roomState.Movie.Id)
	if errMovie != nil {
		return nil, fmt.Errorf("errMovie = %+v", errMovie)
	}
	roomState.Movie = models.Movie{
		Id:         movie.Id,
		Title:      movie.Title,
		TitleImage: movie.TitleUrl,
		Video:      movie.VideoUrl,
	}
	return roomState, err
}

func (s *RoomService) Session(ctx context.Context, cookie string) (*models.SessionRespData, *models.ErrorRespData) {

	user, sesErr, code := s.roomRepository.UserById(ctx, cookie)
	if sesErr != nil {
		errors := make([]errVals.ErrorObj, 1)
		errors[0] = *sesErr

		return nil, &models.ErrorRespData{
			Errors:     errors,
			StatusCode: code,
		}
	}

	return &models.SessionRespData{
		StatusCode: code,
		UserData:   *user,
	}, nil
}
