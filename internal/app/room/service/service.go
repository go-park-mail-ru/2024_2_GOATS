package service

import (
	"context"
	"fmt"
	"github.com/go-park-mail-ru/2024_2_GOATS/internal/app/client"
	errVals "github.com/go-park-mail-ru/2024_2_GOATS/internal/app/errors"
	"github.com/go-park-mail-ru/2024_2_GOATS/internal/app/models"
	model "github.com/go-park-mail-ru/2024_2_GOATS/internal/app/room/model"
	"log"
	"strconv"
)

// TODO раскоментить к 4му РК

//go:generate mockgen -source=service.go -destination=service_mock.go -package=service

//type MovieServiceInterface interface {
//	GetCollection(ctx context.Context) (*model.CollectionsRespData, *errVals.ServiceError)
//	GetMovie(ctx context.Context, mvId int) (*model.MovieInfo, *errVals.ServiceError)
//	GetActor(ctx context.Context, actorId int) (*model.ActorInfo, *errVals.ServiceError)
//	GetMovieByGenre(ctx context.Context, genre string) ([]model.MovieShortInfo, *errVals.ServiceError)
//}

//type MovieServiceInterface interface {
//	GetCollection(ctx context.Context, filter string) (*model.CollectionsRespData, *errVals.ServiceError)
//	GetMovie(ctx context.Context, mvID int) (*models.MovieInfo, *errVals.ServiceError)
//	GetActor(ctx context.Context, actorID int) (*model.ActorInfo, *errVals.ServiceError)
//	GetMovieByGenre(ctx context.Context, genre string) ([]model.MovieShortInfo, *errVals.ServiceError)
//}

type MovieServiceInterface interface {
	GetCollection(ctx context.Context, filter string) (*models.CollectionsRespData, *errVals.ServiceError)
	GetMovie(ctx context.Context, mvID int) (*models.MovieInfo, *errVals.ServiceError)
	GetActor(ctx context.Context, actorID int) (*models.ActorInfo, *errVals.ServiceError)
	GetMovieByGenre(ctx context.Context, genre string) ([]models.MovieShortInfo, *errVals.ServiceError)
}

//type MovieServiceInterface interface {
//	GetCollection(ctx context.Context, filter string) (*models.CollectionsRespData, *errVals.ServiceError)
//	GetMovie(ctx context.Context, mvID int) (*models.MovieInfo, *errVals.ServiceError)
//	GetActor(ctx context.Context, actorID int) (*models.ActorInfo, *errVals.ServiceError)
//	GetMovieByGenre(ctx context.Context, genre string) ([]models.MovieShortInfo, *errVals.ServiceError)
//}

type RoomRepositoryInterface interface {
	CreateRoom(ctx context.Context, room *model.RoomState) (*model.RoomState, error)
	UpdateRoomState(ctx context.Context, roomID string, state *model.RoomState) error
	GetRoomState(ctx context.Context, roomID string) (*model.RoomState, error)
	GetFromCookie(ctx context.Context, cookie string) (string, *errVals.RepoError, int)
	//UserById(ctx context.Context, userId string) (*model.User, *errVals.RepoError, int)
}

type UserServiceInterface interface {
	UpdateProfile(ctx context.Context, profileData *models.User) *errVals.ServiceError
	UpdatePassword(ctx context.Context, passwordData *models.PasswordData) *errVals.ServiceError
	AddFavorite(ctx context.Context, favData *models.Favorite) *errVals.ServiceError
	ResetFavorite(ctx context.Context, favData *models.Favorite) *errVals.ServiceError
	GetFavorites(ctx context.Context, usrID int) ([]models.MovieShortInfo, *errVals.ServiceError)
}

type RoomService struct {
	roomRepository RoomRepositoryInterface
	movieService   client.MovieClientInterface
	userService    client.UserClientInterface
}

func NewService(repo RoomRepositoryInterface, movieService client.MovieClientInterface, userService client.UserClientInterface) *RoomService {
	return &RoomService{
		roomRepository: repo,
		movieService:   movieService,
		userService:    userService,
	}
}

func (s *RoomService) CreateRoom(ctx context.Context, room *model.RoomState) (*model.RoomState, error) {
	return s.roomRepository.CreateRoom(ctx, room)
}

func (s *RoomService) HandleAction(ctx context.Context, roomID string, action model.Action) error {
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
		roomState.Message.Text = action.Message.Text
		//roomState.Message.Avatar = action.Message.Avatar
		//roomState.Message.Sender = action.Message.Sender
	}

	return s.roomRepository.UpdateRoomState(ctx, roomID, roomState)
}

func (s *RoomService) GetRoomState(ctx context.Context, roomID string) (*model.RoomState, error) {
	roomState, err := s.roomRepository.GetRoomState(ctx, roomID)
	log.Println("GetRoomStateGetRoomStateGetRoomStateGetRoomState", roomState)

	movie_service, errMovie := s.movieService.GetMovie(ctx, roomState.Movie.Id)
	if errMovie != nil {
		return nil, fmt.Errorf("errMovie = %+v", errMovie)
	}
	roomState.Movie = model.MovieInfo{
		Id:               movie_service.ID,
		Title:            movie_service.Title,
		TitleUrl:         movie_service.TitleURL,
		ShortDescription: movie_service.ShortDescription, //short_description
		VideoUrl:         movie_service.VideoURL,         //video_url
	}
	return roomState, err
}

//errVals "github.com/go-park-mail-ru/2024_2_GOATS/internal/app/errors"
//*errVals.ServiceError

func (s *RoomService) Session(ctx context.Context, cookie string) (*model.SessionRespData, *errVals.ServiceError) {
	//id := config.CurrentUserID(ctx)
	id, err := strconv.Atoi(cookie)
	log.Println("ididDD = ", id)

	if err != nil {
		return nil, &errVals.ServiceError{
			Code:  "400",
			Error: err,
		}
	}
	log.Println("ididPP = ", id)
	user, sesErr := s.userService.FindByID(ctx, uint64(id))
	if sesErr != nil {
		//errors := make([]errVals.RepoError, 1)
		//errors[0] = *sesErr
		return nil, errVals.NewServiceError(errVals.ErrGetUserCode, fmt.Errorf("failed to login: %w", sesErr))

		//return nil, &model.ErrorRespData{
		//	Errors:     errors,
		//	StatusCode: code,
		//}
	}

	return &model.SessionRespData{
		UserData: *user,
	}, nil
}
