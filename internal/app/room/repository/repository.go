package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"net/http"

	errVals "github.com/go-park-mail-ru/2024_2_GOATS/internal/app/errors"
	model "github.com/go-park-mail-ru/2024_2_GOATS/internal/app/room/model"
	"github.com/go-park-mail-ru/2024_2_GOATS/internal/app/room/service"
	"github.com/google/uuid"

	"github.com/go-redis/redis/v8"
)

// Repo room struct
type Repo struct {
	Database *sql.DB
	Redis    *redis.Client
}

// NewRepository returns an instance of RoomRepositoryInterface
func NewRepository(rdb *redis.Client) service.RoomRepositoryInterface {
	return &Repo{
		Redis: rdb,
	}
}

// CreateRoom creates room
func (r *Repo) CreateRoom(ctx context.Context, room *model.RoomState) (*model.RoomState, error) {
	room.ID = uuid.New().String()
	data, err := room.MarshalJSON()
	if err != nil {
		return nil, err
	}
	err = r.Redis.Set(ctx, "room_state:"+room.ID, data, 0).Err()
	if err != nil {
		return nil, err
	}
	return room, nil
}

// UpdateRoomState updates room
func (r *Repo) UpdateRoomState(ctx context.Context, roomID string, state *model.RoomState) error {
	data, err := state.MarshalJSON()
	if err != nil {
		return err
	}
	return r.Redis.Set(ctx, "room_state:"+roomID, data, 0).Err()
}

// GetRoomState gets room stats
func (r *Repo) GetRoomState(ctx context.Context, roomID string) (*model.RoomState, error) {
	data, err := r.Redis.Get(ctx, "room_state:"+roomID).Result()
	if errors.Is(err, redis.Nil) {
		return nil, errors.New("room state not found")
	} else if err != nil {
		return nil, err
	}

	var state model.RoomState
	err = state.UnmarshalJSON([]byte(data))
	log.Println("state =", state)
	if err != nil {
		return nil, err
	}
	return &state, nil
}

// GetFromCookie get user data from cookie
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
