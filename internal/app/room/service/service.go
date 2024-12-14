package service

import (
	"context"
	"fmt"
	"github.com/go-park-mail-ru/2024_2_GOATS/internal/app/client"
	errVals "github.com/go-park-mail-ru/2024_2_GOATS/internal/app/errors"
	"github.com/go-park-mail-ru/2024_2_GOATS/internal/app/models"
	model "github.com/go-park-mail-ru/2024_2_GOATS/internal/app/room/model"
	ws "github.com/go-park-mail-ru/2024_2_GOATS/internal/app/room/ws"
	"log"
	"strconv"
)

// TODO раскоментить к 4му РК

//go:generate mockgen -source=service.go -destination=service_mock.go -package=service

type MovieServiceInterface interface {
	GetCollection(ctx context.Context, filter string) (*models.CollectionsRespData, *errVals.ServiceError)
	GetMovie(ctx context.Context, mvID int) (*models.MovieInfo, *errVals.ServiceError)
	GetActor(ctx context.Context, actorID int) (*models.ActorInfo, *errVals.ServiceError)
	GetMovieByGenre(ctx context.Context, genre string) ([]models.MovieShortInfo, *errVals.ServiceError)
}

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
	timerManager   *ws.TimerManager
	hub            *ws.RoomHub
}

func NewService(repo RoomRepositoryInterface, movieService client.MovieClientInterface, userService client.UserClientInterface, hub *ws.RoomHub, TimerManager *ws.TimerManager) *RoomService {
	return &RoomService{
		roomRepository: repo,
		movieService:   movieService,
		userService:    userService,
		hub:            hub,
		timerManager:   TimerManager,
	}
}

func (s *RoomService) CreateRoom(ctx context.Context, room *model.RoomState) (*model.RoomState, error) {
	return s.roomRepository.CreateRoom(ctx, room)
}

//func (s *RoomService) HandleAction(ctx context.Context, roomID string, action model.Action) error {
//	roomState, err := s.roomRepository.GetRoomState(ctx, roomID)
//	if err != nil {
//		return err
//	}
//
//	log.Println("roomState.Status === ", roomState.Status)
//	log.Println("action.Name === ", action.Name)
//	log.Println("action.TimeCode === ", action.TimeCode)
//
//	switch action.Name {
//	case "pause":
//		roomState.Status = "paused"
//		roomState.TimeCode = action.TimeCode
//	case "play":
//		roomState.Status = "playing"
//	case "rewind":
//		roomState.TimeCode = action.TimeCode
//	case "timer":
//		roomState.TimeCode = action.TimeCode
//		//roomStateRepo, _ := s.roomRepository.GetRoomState(ctx, roomID)
//		//roomState = roomStateRepo
//	case "message":
//		roomState.Message.Text = action.Message.Text
//		//roomState.Message.Avatar = action.Message.Avatar
//	}
//	return s.roomRepository.UpdateRoomState(ctx, roomID, roomState)
//}

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
		movie_service, errMovie := s.movieService.GetMovie(ctx, action.MovieId)
		if errMovie != nil {
			log.Println("errMovie", errMovie)
		}
		log.Println("movie_service==", movie_service.VideoURL)
		//roomState.Message.Text = action.Message.Text
		roomState.Movie.ID = movie_service.ID
		roomState.Movie.MovieType = movie_service.MovieType
		log.Println("MovieType==", roomState.Movie.MovieType)

		//if movie_service.MovieType == "serial" {
		seasons := []*model.Season{}
		for _, season := range movie_service.Seasons {
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
		log.Println("SeasonNumber==")

		roomState.Movie.AlbumURL = movie_service.AlbumURL
		log.Println("roomState.Movie.AlbumURL==", roomState.Movie.AlbumURL)
		roomState.Movie.CardURL = movie_service.CardURL
		log.Println("roomState.Movie.CardURL==", roomState.Movie.CardURL)
		roomState.Movie.TitleURL = movie_service.TitleURL
		log.Println("roomState.Movie.TitleURL==", roomState.Movie.TitleURL)
		roomState.Movie.Title = movie_service.Title
		log.Println("roomState.Movie.Title==", roomState.Movie.Title)
		roomState.Movie.VideoURL = movie_service.VideoURL
		roomState.Movie.Rating = movie_service.Rating
		roomState.Movie.ShortDescription = movie_service.ShortDescription

		s.hub.Broadcast <- ws.BroadcastMessage{
			Action: map[string]interface{}{
				"name":   "change_movie",
				"movie ": roomState.Movie,
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

	//case "change":
	//	movie_service, errMovie := s.movieService.GetMovie(ctx, roomState.Movie.ID)
	//	roomState.Message.Text = action.Message.Text
	//
	//	roomState, err := h.roomService.GetRoomState(r.Context(), roomID)
	//	if err != nil {
	//	log.Println("Failed to get room state from Redis:", err)
	//	} else {
	//	if err := conn.WriteJSON(roomState); err != nil {
	//	log.Println("Failed to send room state:", err)
	//	return
	//	}
	//	}

	//if errMovie != nil {
	//return nil, fmt.Errorf("errMovie = %+v", errMovie)
	//}
	//}

	//movie_service, errMovie := s.movieService.GetMovie(ctx, roomState.Movie.ID)
	//	if errMovie != nil {
	//		return nil, fmt.Errorf("errMovie = %+v", errMovie)
	//	}x
	//	roomState.Movie = model.MovieInfo{
	//		ID:               movie_service.ID,
	//		Title:            movie_service.Title,
	//		TitleURL:         movie_service.TitleURL,
	//		ShortDescription: movie_service.ShortDescription,
	//		VideoURL:         movie_service.VideoURL,
	//	}
	//	return roomState, err

	return s.roomRepository.UpdateRoomState(ctx, roomID, roomState)
}

func (s *RoomService) GetRoomState(ctx context.Context, roomID string) (*model.RoomState, error) {
	roomState, err := s.roomRepository.GetRoomState(ctx, roomID)
	log.Println("GetRoomStateGetRoomStateGetRoomStateGetRoomState", roomState)

	movie_service, errMovie := s.movieService.GetMovie(ctx, roomState.Movie.ID)
	if errMovie != nil {
		return nil, fmt.Errorf("errMovie = %+v", errMovie)
	}
	roomState.Movie = model.MovieInfo{
		ID:               movie_service.ID,
		Title:            movie_service.Title,
		TitleURL:         movie_service.TitleURL,
		ShortDescription: movie_service.ShortDescription,
		VideoURL:         movie_service.VideoURL,
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

//func (s *RoomService) startTimer(ctx context.Context, roomID string, initialTimeCode int) {
//	if _, ok := s.timers[roomID]; ok {
//		s.stopTimer(roomID)
//	}
//
//	log.Println("startTimerqwedwedwewd")
//	log.Println("initialTimeCode ==", initialTimeCode)
//	log.Println("roomID ==", roomID)
//
//	timer := &Timer{
//		Ticker:   time.NewTicker(3 * time.Second),
//		Quit:     make(chan struct{}),
//		TimeCode: initialTimeCode,
//	}
//	log.Println("startTimer11", roomID)
//	log.Println("startTim", timer)
//	log.Println("startTimqwd", timer.TimeCode)
//
//	s.timers[roomID] = timer
//
//	log.Println("startTimer22", roomID)
//	go func() {
//		for {
//			log.Println("TTTTTTTTTT")
//			select {
//			case <-timer.Ticker.C:
//				timer.TimeCode += 3
//				log.Println("timer", timer.TimeCode)
//
//				// Отправляем обновление в RoomHub через канал Broadcast
//				s.hub.Broadcast <- ws.BroadcastMessage{
//					Action: map[string]interface{}{
//						"type":     "timer",
//						"timeCode": timer.TimeCode,
//					},
//					RoomID: roomID,
//				}
//
//				log.Println("timer", timer.TimeCode)
//				// Обновляем состояние комнаты
//				roomState, _ := s.roomRepository.GetRoomState(ctx, roomID)
//				roomState.TimeCode = float64(timer.TimeCode)
//				s.roomRepository.UpdateRoomState(ctx, roomID, roomState)
//
//			case <-timer.Quit:
//				return
//			}
//		}
//	}()
//	log.Println("TTTTTTTTTT22222")
//
//}
