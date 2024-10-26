package service

import (
	"context"
	"fmt"
	errVals "github.com/go-park-mail-ru/2024_2_GOATS/internal/app/errors"
	models "github.com/go-park-mail-ru/2024_2_GOATS/internal/app/room/model"
	//model "github.com/go-park-mail-ru/2024_2_GOATS/internal/app/models"
	movie "github.com/go-park-mail-ru/2024_2_GOATS/internal/app/movie/delivery"
	"github.com/go-park-mail-ru/2024_2_GOATS/internal/app/room/repository"
	"log"
)

//type MovieServiceInterface interface {
//	GetCollection(ctx context.Context) (*model.CollectionsRespData, *models.ErrorRespData)
//	GetMovie(ctx context.Context, mvId int) (*model.MovieInfo, *models.ErrorRespData)
//	GetActor(ctx context.Context, actorId int) (*model.StaffInfo, *models.ErrorRespData)
//}

type RoomService struct {
	roomRepository repository.RoomRepositoryInterface
	movieService   movie.MovieServiceInterface
}

func NewService(repo repository.RoomRepositoryInterface, movieService movie.MovieServiceInterface) *RoomService {
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

	log.Println("roomState === ", roomState.Status)
	log.Println("action === ", action.Name)
	log.Println("action === ", action.TimeCode)
	log.Println("action === ", action.Name)
	log.Println("action === ", action.TimeCode)
	log.Println("action === ", action.Name)

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
		return nil, fmt.Errorf("errMovie", errMovie)
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
	userId, err, code := s.roomRepository.GetFromCookie(ctx, cookie)
	if err != nil || userId == "" {
		return nil, &models.ErrorRespData{
			Errors:     []errVals.ErrorObj{*err},
			StatusCode: code,
		}
	}

	user, sesErr, code := s.roomRepository.UserById(ctx, userId)
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
