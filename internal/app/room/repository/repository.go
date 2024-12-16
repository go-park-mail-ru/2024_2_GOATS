package repository

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	errVals "github.com/go-park-mail-ru/2024_2_GOATS/internal/app/errors"
	model "github.com/go-park-mail-ru/2024_2_GOATS/internal/app/room/model"
	"github.com/go-redis/redis/v8"
	"github.com/google/uuid"
)

// RoomRepositoryInterface интерфейс репозитория комнаты
type RoomRepositoryInterface interface {
	CreateRoom(ctx context.Context, room *model.RoomState) (*model.RoomState, error)
	UpdateRoomState(ctx context.Context, roomID string, state *model.RoomState) error
	GetRoomState(ctx context.Context, roomID string) (*model.RoomState, error)
	GetFromCookie(ctx context.Context, cookie string) (string, *errVals.RepoError, int)
	//UserById(ctx context.Context, userId string) (*model.User, *errVals.RepoError, int)
}

// Repo структура репоитория
type Repo struct {
	Database *sql.DB
	Redis    *redis.Client
}

// NewRepository конструктор репоитория
func NewRepository(rdb *redis.Client) RoomRepositoryInterface {
	return &Repo{
		Redis: rdb,
	}
}

// CreateRoom создание комнаты
func (r *Repo) CreateRoom(ctx context.Context, room *model.RoomState) (*model.RoomState, error) {
	room.ID = uuid.New().String()
	data, err := json.Marshal(room)
	if err != nil {
		return nil, err
	}
	err = r.Redis.Set(ctx, "room_state:"+room.ID, data, 0).Err()
	if err != nil {
		return nil, err
	}
	return room, nil
}

// UpdateRoomState создание комнаты
func (r *Repo) UpdateRoomState(ctx context.Context, roomID string, state *model.RoomState) error {
	data, err := json.Marshal(state)
	if err != nil {
		return err
	}
	return r.Redis.Set(ctx, "room_state:"+roomID, data, 0).Err()
}

// GetRoomState получение статистики комнаты
func (r *Repo) GetRoomState(ctx context.Context, roomID string) (*model.RoomState, error) {
	data, err := r.Redis.Get(ctx, "room_state:"+roomID).Result()
	if errors.Is(err, redis.Nil) {
		return nil, errors.New("room state not found")
	} else if err != nil {
		return nil, err
	}

	var state model.RoomState
	err = json.Unmarshal([]byte(data), &state)
	if err != nil {
		return nil, err
	}
	return &state, nil
}

// GetFromCookie получение данных из куков
func (r *Repo) GetFromCookie(ctx context.Context, cookie string) (string, *errVals.RepoError, int) {
	var userID string
	err := r.Redis.Get(ctx, cookie).Scan(&userID)
	if err != nil {
		return "", errVals.NewRepoError(
			errVals.ErrCreateUserCode,
			errVals.CustomError{Err: fmt.Sprintf("cannot get cookie from redis: %v", err)},
		), http.StatusForbidden
	}

	return userID, nil, http.StatusOK
}
