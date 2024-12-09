package repository

// TODO раскоментить к 4му РК

//import (
//	"context"
//	"database/sql"
//
//	errVals "github.com/go-park-mail-ru/2024_2_GOATS/internal/app/errors"
//	model "github.com/go-park-mail-ru/2024_2_GOATS/internal/app/room/model"
//
//	// user "github.com/go-park-mail-ru/2024_2_GOATS/internal/app/user/repository/userdb"
//	"github.com/go-redis/redis/v8"
//)
//
//type RoomRepositoryInterface interface {
//	CreateRoom(ctx context.Context, room *model.RoomState) (*model.RoomState, error)
//	UpdateRoomState(ctx context.Context, roomID string, state *model.RoomState) error
//	GetRoomState(ctx context.Context, roomID string) (*model.RoomState, error)
//	GetFromCookie(ctx context.Context, cookie string) (string, *errVals.RepoError, int)
//	UserById(ctx context.Context, userId string) (*model.User, *errVals.RepoError, int)
//}
//
//type Repo struct {
//	Database *sql.DB
//	Redis    *redis.Client
//}
//
//func NewRepository(db *sql.DB, rdb *redis.Client) RoomRepositoryInterface {
//	return &Repo{
//		Database: db,
//		Redis:    rdb,
//	}
//}
//
//func (r *Repo) CreateRoom(ctx context.Context, room *model.RoomState) (*model.RoomState, error) {
//	room.Id = uuid.New().String() // Генерация уникального ID для комнаты
//	data, err := json.Marshal(room)
//	if err != nil {
//		return nil, err
//	}
//	log.Println("data", string(data))
//	err = r.Redis.Set(ctx, "room_state:"+room.Id, data, 0).Err() // Сохранение комнаты в Redis
//	if err != nil {
//		return nil, err
//	}
//	return room, nil
//}
//
//func (r *Repo) UpdateRoomState(ctx context.Context, roomID string, state *model.RoomState) error {
//	data, err := json.Marshal(state)
//	if err != nil {
//		return err
//	}
//	return r.Redis.Set(ctx, "room_state:"+roomID, data, 0).Err()
//}
//
//func (r *Repo) GetRoomState(ctx context.Context, roomID string) (*model.RoomState, error) {
//	data, err := r.Redis.Get(ctx, "room_state:"+roomID).Result()
//	if err == redis.Nil {
//		return nil, errors.New("room state not found")
//	} else if err != nil {
//		return nil, err
//	}
//
//	var state model.RoomState
//	err = json.Unmarshal([]byte(data), &state)
//	log.Println("state =", state)
//	if err != nil {
//		return nil, err
//	}
//	return &state, nil
//}
//
//func (r *Repo) GetFromCookie(ctx context.Context, cookie string) (string, *errVals.RepoError, int) {
//	var userID string
//	err := r.Redis.Get(ctx, cookie).Scan(&userID)
//	log.Println("err =", err)
//	if err != nil {
//		return "", errVals.NewRepoError(
//			errVals.ErrCreateUserCode,
//			errVals.CustomError{Err: fmt.Errorf("cannot get cookie from redis: %w", err)},
//		), http.StatusForbidden
//	}
//
//	return userID, nil, http.StatusOK
//}
//
//func (r *Repo) UserById(ctx context.Context, userId string) (*model.User, *errVals.RepoError, int) {
//	userIdInt, err := strconv.Atoi(userId)
//	if err != nil {
//		log.Println("Ошибка перевода str в int", err)
//	}
//	usr, err := user.FindByID(ctx, userIdInt, r.Database)
//	if err != nil {
//		if errors.Is(err, sql.ErrNoRows) {
//			return nil, errVals.NewRepoError(errVals.ErrUserNotFoundCode, errVals.ErrUserNotFound), http.StatusNotFound
//		}
//
//		return nil, errVals.NewRepoError(errVals.ErrServerCode, errVals.CustomError{Err: err}), http.StatusUnprocessableEntity
//	}
//
//	return &model.User{
//		ID:        usr.ID,
//		Email:     usr.Email,
//		Username:  usr.Username,
//		Password:  usr.Password,
//		AvatarURL: usr.AvatarURL,
//	}, nil, http.StatusOK
//	//return *model.User(usr), nil, http.StatusOK
//}
