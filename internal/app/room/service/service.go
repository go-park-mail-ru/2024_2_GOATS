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
	"sync"
	"time"
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
	timerManager   *TimerManager
	hub            *ws.RoomHub
}

type TimerManager struct {
	mu     sync.Mutex
	timers map[string]chan struct{}
	hub    *ws.RoomHub
}

func NewTimerManager(hub *ws.RoomHub) *TimerManager {
	return &TimerManager{
		timers: make(map[string]chan struct{}),
		hub:    hub,
	}
}

func NewService(repo RoomRepositoryInterface, movieService client.MovieClientInterface, userService client.UserClientInterface, hub *ws.RoomHub, TimerManager *TimerManager) *RoomService {
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
		})

	case "rewind":
		roomState.TimeCode = action.TimeCode
		s.timerManager.Stop(roomID)
		s.timerManager.Start(roomID, int64(roomState.TimeCode), func(updatedTime int64) {
			roomState.TimeCode = float64(updatedTime)
			_ = s.roomRepository.UpdateRoomState(ctx, roomID, roomState)
		})

	case "timer":
		roomState.TimeCode = action.TimeCode

	case "message":
		roomState.Message.Text = action.Message.Text
	}

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

func (tm *TimerManager) Start(roomID string, startTime int64, updateFunc func(int64)) {
	tm.mu.Lock()
	if _, exists := tm.timers[roomID]; exists {
		tm.mu.Unlock()
		return
	}
	quit := make(chan struct{})
	tm.timers[roomID] = quit
	tm.mu.Unlock()

	go func() {
		timeCode := startTime
		ticker := time.NewTicker(3 * time.Second)
		defer ticker.Stop()

		for {
			select {
			case <-ticker.C:
				timeCode += 3
				tm.hub.Broadcast <- ws.BroadcastMessage{
					Action: map[string]interface{}{
						"type":     "timer",
						"timeCode": timeCode,
					},
					RoomID: roomID,
				}
				updateFunc(timeCode)
			case <-quit:
				return
			}
		}
	}()
}

func (tm *TimerManager) Stop(roomID string) {
	tm.mu.Lock()
	if quit, exists := tm.timers[roomID]; exists {
		close(quit)
		delete(tm.timers, roomID)
	}
	tm.mu.Unlock()
}

//func (s *RoomService) startTimer(ctx context.Context, roomID string, roomState *model.RoomState) {
//	timeCode := int(roomState.TimeCode)
//	ticker := time.NewTicker(1 * time.Second)
//	defer ticker.Stop()
//
//	for {
//		select {
//		case <-ticker.C:
//			timeCode++
//			roomState.TimeCode = float64(int64(timeCode))
//
//			// Сохраняем обновлённое состояние комнаты
//			err := s.roomRepository.UpdateRoomState(ctx, roomID, roomState)
//			if err != nil {
//				log.Println("Error updating room state:", err)
//				return
//			}
//
//		case <-roomState.TimerQuit:
//			log.Println("Timer stopped for room:", roomID)
//			return
//		}
//	}
//}

//func (s *RoomService) stopTimer(roomID string) {
//	if timer, ok := s.timers[roomID]; ok {
//		timer.Ticker.Stop()
//		close(timer.Quit)
//		delete(s.timers, roomID)
//	}
//}

//
//func (s *RoomService) HandleAction(ctx context.Context, roomID string, action model.Action) error {
//	// Получаем текущее состояние комнаты
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
//		// Ставим статус в "paused" и сохраняем текущий timeCode
//		roomState.Status = "paused"
//		roomState.TimeCode = action.TimeCode
//
//		// Останавливаем таймер
//		//if timer, ok := s.timers[roomID]; ok {
//		//	timer.Quit <- struct{}{}
//		//	delete(s.timers, roomID)
//		//}
//
//		if roomState.TimerQuit != nil {
//			close(roomState.TimerQuit)
//			roomState.TimerQuit = nil
//		}
//
//	case "play":
//		// Ставим статус в "playing"
//		roomState.Status = "playing"
//
//		// Запускаем таймер, если он ещё не запущен
//		//if _, ok := s.timers[roomID]; !ok {
//		//	s.startTimer(ctx, roomID, int(roomState.TimeCode))
//		//}
//		if roomState.TimerQuit == nil {
//			roomState.TimerQuit = make(chan struct{})
//			go s.startTimer(ctx, roomID, roomState)
//		}
//
//	case "rewind":
//		// Обновляем timeCode
//		roomState.TimeCode = action.TimeCode
//
//		// Останавливаем текущий таймер и запускаем новый
//		//if timer, ok := s.timers[roomID]; ok {
//		//	timer.Quit <- struct{}{}
//		//	delete(s.timers, roomID)
//		//}
//		//s.startTimer(ctx, roomID, int(roomState.TimeCode))
//
//		if roomState.TimerQuit == nil {
//			roomState.TimerQuit = make(chan struct{})
//			go s.startTimer(ctx, roomID, roomState)
//		}
//
//	case "timer":
//		// Просто обновляем текущее время
//		roomState.TimeCode = action.TimeCode
//
//	case "message":
//		// Обновляем сообщение
//		roomState.Message.Text = action.Message.Text
//	}
//
//	// Обновляем состояние комнаты
//	return s.roomRepository.UpdateRoomState(ctx, roomID, roomState)
//}
