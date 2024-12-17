package service

import (
	"context"
	"fmt"
	"log"
	"strconv"

	"github.com/go-park-mail-ru/2024_2_GOATS/internal/app/client"
	errVals "github.com/go-park-mail-ru/2024_2_GOATS/internal/app/errors"
	model "github.com/go-park-mail-ru/2024_2_GOATS/internal/app/room/model"
	ws "github.com/go-park-mail-ru/2024_2_GOATS/internal/app/room/ws"
)

//go:generate mockgen -source=service.go -destination=service_mock.go -package=service

// RoomRepositoryInterface defines methods for room repository
type RoomRepositoryInterface interface {
	CreateRoom(ctx context.Context, room *model.RoomState) (*model.RoomState, error)
	UpdateRoomState(ctx context.Context, roomID string, state *model.RoomState) error
	GetRoomState(ctx context.Context, roomID string) (*model.RoomState, error)
	GetFromCookie(ctx context.Context, cookie string) (string, *errVals.RepoError, int)
}

// RoomService service struct
type RoomService struct {
	roomRepository RoomRepositoryInterface
	movieService   client.MovieClientInterface
	userService    client.UserClientInterface
	timerManager   *ws.TimerManager
	hub            *ws.RoomHub
}

// NewService returns an instance of RoomService
func NewService(
	repo RoomRepositoryInterface,
	movieService client.MovieClientInterface,
	userService client.UserClientInterface,
	hub *ws.RoomHub,
	TimerManager *ws.TimerManager,
) *RoomService {
	return &RoomService{
		roomRepository: repo,
		movieService:   movieService,
		userService:    userService,
		hub:            hub,
		timerManager:   TimerManager,
	}
}

// CreateRoom creates room
func (s *RoomService) CreateRoom(ctx context.Context, room *model.RoomState) (*model.RoomState, error) {
	return s.roomRepository.CreateRoom(ctx, room)
}

// HandleAction handles actions from frontend
func (s *RoomService) HandleAction(ctx context.Context, roomID string, action model.Action) error {
	roomState, err := s.roomRepository.GetRoomState(ctx, roomID)
	if err != nil {
		return err
	}

	switch action.Name {
	case "pause":
		roomState.Status = "paused"
		roomState.TimeCode = action.TimeCode
		s.timerManager.Stop(roomID)

	case "play":
		roomState.Status = "playing"
		s.timerManager.Start(roomID, int64(roomState.TimeCode), func(updatedTime int64) {
			roomState.TimeCode = float64(updatedTime)
			_ = s.roomRepository.UpdateRoomState(ctx, roomID, roomState)
		}, int64(roomState.Duration))

	case "rewind":
		roomState.TimeCode = action.TimeCode
		s.timerManager.Stop(roomID)
		if roomState.Status == "playing" {
			s.timerManager.Start(roomID, int64(roomState.TimeCode), func(updatedTime int64) {
				roomState.TimeCode = float64(updatedTime)
				_ = s.roomRepository.UpdateRoomState(ctx, roomID, roomState)
			}, int64(roomState.Duration))
		} else {
			roomState.TimeCode = action.TimeCode
		}

	case "message":
		roomState.Message.Text = action.Message.Text

	case "change":
		movieService, errMovie := s.movieService.GetMovie(ctx, action.MovieID)
		if errMovie != nil {
			log.Println("errMovie", errMovie)
		}
		roomState.Movie.ID = movieService.ID
		roomState.Movie.MovieType = movieService.MovieType

		seasons := []*model.Season{}
		for _, season := range movieService.Seasons {
			sn := season.SeasonNumber
			var eps []*model.Episode
			for _, ep := range season.Episodes {
				cur := &model.Episode{
					ID:            ep.ID,
					Description:   ep.Description,
					EpisodeNumber: ep.EpisodeNumber,
					Title:         ep.Title,
					Rating:        ep.Rating,
					ReleaseDate:   ep.ReleaseDate,
					VideoURL:      ep.VideoURL,
					PreviewURL:    ep.PreviewURL,
				}

				eps = append(eps, cur)
			}

			curSeas := &model.Season{
				SeasonNumber: sn,
				Episodes:     eps,
			}

			seasons = append(seasons, curSeas)
		}
		roomState.Movie.Seasons = seasons

		roomState.Movie.AlbumURL = movieService.AlbumURL
		roomState.Movie.CardURL = movieService.CardURL
		roomState.Movie.TitleURL = movieService.TitleURL
		roomState.Movie.Title = movieService.Title
		roomState.Movie.VideoURL = movieService.VideoURL
		roomState.Movie.Rating = movieService.Rating
		roomState.Movie.ShortDescription = movieService.ShortDescription

		roomState.SeasonNow = 0
		roomState.EpisodeNow = 0

		s.hub.Broadcast <- ws.BroadcastMessage{
			Action: map[string]interface{}{
				"name":  "change_movie",
				"movie": roomState.Movie,
			},
			RoomID: roomID,
		}

	case "change_series":
		roomState.SeasonNow = action.SeasonNow
		roomState.EpisodeNow = action.EpisodeNow
		_ = s.roomRepository.UpdateRoomState(ctx, roomID, roomState)

	case "duration":
		roomState.Duration = action.Duration
		_ = s.roomRepository.UpdateRoomState(ctx, roomID, roomState)
	}

	return s.roomRepository.UpdateRoomState(ctx, roomID, roomState)
}

// GetRoomState gets room state
func (s *RoomService) GetRoomState(ctx context.Context, roomID string) (*model.RoomState, error) {
	roomState, err := s.roomRepository.GetRoomState(ctx, roomID)

	movieService, errMovie := s.movieService.GetMovie(ctx, roomState.Movie.ID)
	if errMovie != nil {
		return nil, fmt.Errorf("errMovie = %+v", errMovie)
	}

	var seasons []*model.Season

	for _, season := range movieService.Seasons {
		sn := season.SeasonNumber
		var eps []*model.Episode
		for _, ep := range season.Episodes {
			cur := &model.Episode{
				ID:            ep.ID,
				Description:   ep.Description,
				EpisodeNumber: ep.EpisodeNumber,
				Title:         ep.Title,
				Rating:        ep.Rating,
				ReleaseDate:   ep.ReleaseDate,
				VideoURL:      ep.VideoURL,
				PreviewURL:    ep.PreviewURL,
			}
			eps = append(eps, cur)
		}

		curSeas := &model.Season{
			SeasonNumber: int(sn),
			Episodes:     eps,
		}

		seasons = append(seasons, curSeas)
	}

	roomState.Movie = model.MovieInfo{
		ID:               movieService.ID,
		Title:            movieService.Title,
		TitleURL:         movieService.TitleURL,
		ShortDescription: movieService.ShortDescription,
		VideoURL:         movieService.VideoURL,
		Seasons:          seasons,
	}
	return roomState, err
}

// Session checks session
func (s *RoomService) Session(ctx context.Context, cookie string) (*model.SessionRespData, *errVals.ServiceError) {
	id, err := strconv.Atoi(cookie)
	log.Println("ididDD = ", id)

	if err != nil {
		return nil, &errVals.ServiceError{
			Code:  "400",
			Error: err,
		}
	}
	user, sesErr := s.userService.FindByID(ctx, uint64(id))
	if sesErr != nil {
		return nil, errVals.NewServiceError(errVals.ErrGetUserCode, fmt.Errorf("failed to login: %w", sesErr))
	}

	return &model.SessionRespData{
		UserData: model.User{
			ID:         user.ID,
			Email:      user.Email,
			Username:   user.Username,
			Password:   user.Password,
			AvatarURL:  user.AvatarURL,
			AvatarName: user.AvatarName,
		},
	}, nil
}
