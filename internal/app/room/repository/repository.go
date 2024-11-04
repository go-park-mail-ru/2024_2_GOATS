package repository

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strconv"

	errVals "github.com/go-park-mail-ru/2024_2_GOATS/internal/app/errors"
	models "github.com/go-park-mail-ru/2024_2_GOATS/internal/app/room/model"
	"github.com/go-park-mail-ru/2024_2_GOATS/internal/app/user/repository/user"
	"github.com/go-redis/redis/v8"
	"github.com/google/uuid"
)

type RoomRepositoryInterface interface {
	CreateRoom(ctx context.Context, room *models.RoomState) (*models.RoomState, error)
	UpdateRoomState(ctx context.Context, roomID string, state *models.RoomState) error
	GetRoomState(ctx context.Context, roomID string) (*models.RoomState, error)
	GetFromCookie(ctx context.Context, cookie string) (string, *errVals.ErrorObj, int)
	UserById(ctx context.Context, userId string) (*models.User, *errVals.ErrorObj, int)
}

type Repo struct {
	Database *sql.DB
	Redis    *redis.Client
}

func NewRepository(db *sql.DB, rdb *redis.Client) RoomRepositoryInterface {
	return &Repo{
		Database: db,
		Redis:    rdb,
	}
}

func (r *Repo) CreateRoom(ctx context.Context, room *models.RoomState) (*models.RoomState, error) {
	room.Id = uuid.New().String() // Генерация уникального ID для комнаты
	data, err := json.Marshal(room)
	if err != nil {
		return nil, err
	}
	log.Println("data", string(data))
	err = r.Redis.Set(ctx, "room_state:"+room.Id, data, 0).Err() // Сохранение комнаты в Redis
	if err != nil {
		return nil, err
	}
	return room, nil
}

func (r *Repo) UpdateRoomState(ctx context.Context, roomID string, state *models.RoomState) error {
	data, err := json.Marshal(state)
	if err != nil {
		return err
	}
	return r.Redis.Set(ctx, "room_state:"+roomID, data, 0).Err()
}

func (r *Repo) GetRoomState(ctx context.Context, roomID string) (*models.RoomState, error) {
	data, err := r.Redis.Get(ctx, "room_state:"+roomID).Result()
	if err == redis.Nil {
		return nil, errors.New("room state not found")
	} else if err != nil {
		return nil, err
	}

	var state models.RoomState
	err = json.Unmarshal([]byte(data), &state)
	log.Println("state =", state)
	if err != nil {
		return nil, err
	}
	return &state, nil
}

func (r *Repo) GetFromCookie(ctx context.Context, cookie string) (string, *errVals.ErrorObj, int) {
	var userID string
	err := r.Redis.Get(ctx, cookie).Scan(&userID)
	log.Println("err =", err)
	if err != nil {
		return "", errVals.NewErrorObj(
			errVals.ErrCreateUserCode,
			errVals.CustomError{Err: fmt.Errorf("cannot get cookie from redis: %w", err)},
		), http.StatusForbidden
	}

	return userID, nil, http.StatusOK
}

func (r *Repo) UserById(ctx context.Context, userId string) (*models.User, *errVals.ErrorObj, int) {
	userIdInt, err := strconv.Atoi(userId)
	if err != nil {
		log.Println("Ошибка перевода str в int", err)
	}
	usr, err := user.FindById(ctx, userIdInt, r.Database)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errVals.NewErrorObj(errVals.ErrUserNotFoundCode, errVals.ErrUserNotFoundText), http.StatusNotFound
		}

		return nil, errVals.NewErrorObj(errVals.ErrServerCode, errVals.CustomError{Err: err}), http.StatusUnprocessableEntity
	}

	return (*models.User)(usr), nil, http.StatusOK
}
