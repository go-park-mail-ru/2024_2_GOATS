package service

import (
	"context"
	errVals "github.com/go-park-mail-ru/2024_2_GOATS/internal/app/errors"
	models "github.com/go-park-mail-ru/2024_2_GOATS/internal/app/room/model"
	"github.com/go-park-mail-ru/2024_2_GOATS/internal/app/room/repository"
	"log"
)

type RoomService struct {
	roomRepository repository.RoomRepositoryInterface
}

func NewService(repo repository.RoomRepositoryInterface) *RoomService {
	return &RoomService{
		roomRepository: repo,
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
	}

	return s.roomRepository.UpdateRoomState(ctx, roomID, roomState)
}

//const movie = {
//id: 1,
//title: 'Сопрано',
//titleImage:
//'https://i.pinimg.com/originals/93/c7/54/93c754126bcdecb6e540e02631f5eda1.png',
//shortDescription:
//'Мафиозный босс Нью-Джерси обращается за помощью к психологу. Культовый сериал, ставший образцом гангстерскогокино',
//longDescription:
//'Сюжет и Основная Идея: "Клан Сопрано" предлагает зрителю уникальную возможность погрузиться в жизнь Тони Сопрано, главы мафиозной семьи в Нью-Джерси. Великолепно сыгранный Джеймсом Гандольфини, Тони Сопрано становится центром вселенной, где переплетаются его внутренние конфликты, семейные проблемы и криминальные дела.',
//image:
//'https://avatars.mds.yandex.net/i?id=47e4f9e1d2423964f4db57f4db284b2e_l-5350095-images-thumbs&n=13',
//rating: 10,
//releaseDate: '13.09.1999',
//country: 'США',
//director: 'Тимоти Ван Паттен',
//isSerial: false,
//video:
//'http://commondatastorage.googleapis.com/gtv-videos-bucket/sample/BigBuckBunny.mp4',
//};

func (s *RoomService) GetRoomState(ctx context.Context, roomID string) (*models.RoomState, error) {
	qwer, err := s.roomRepository.GetRoomState(ctx, roomID)
	log.Println("GetRoomStateGetRoomStateGetRoomStateGetRoomState", qwer)
	qwer.Movie = models.Movie{
		Id:         1,
		Title:      "Сопрано",
		TitleImage: "https://i.pinimg.com/originals/93/c7/54/93c754126bcdecb6e540e02631f5eda1.png",
		Video:      "http://commondatastorage.googleapis.com/gtv-videos-bucket/sample/BigBuckBunny.mp4",
	}
	return qwer, err
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
