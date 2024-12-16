package delivery

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/go-park-mail-ru/2024_2_GOATS/config"
	errVals "github.com/go-park-mail-ru/2024_2_GOATS/internal/app/errors"
	model "github.com/go-park-mail-ru/2024_2_GOATS/internal/app/room/model"
	ws "github.com/go-park-mail-ru/2024_2_GOATS/internal/app/room/ws"
	websocket "github.com/gorilla/websocket"
	"github.com/rs/zerolog/log"
	//zlog "github.com/rs/zerolog/log"
)

// RoomServiceInterface интерейс сервиса комнаты
type RoomServiceInterface interface {
	CreateRoom(ctx context.Context, room *model.RoomState) (*model.RoomState, error)
	HandleAction(ctx context.Context, roomID string, action model.Action) error
	Session(ctx context.Context, cookie string) (*model.SessionRespData, *errVals.ServiceError)
	GetRoomState(ctx context.Context, roomID string) (*model.RoomState, error)
}

// RoomHandler структура хэндлера комнаты
type RoomHandler struct {
	roomService RoomServiceInterface
	roomHub     *ws.RoomHub
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(_ *http.Request) bool { return true },
}

// NewRoomHandler конструктор хэндлера комнаты
func NewRoomHandler(service RoomServiceInterface, roomHub *ws.RoomHub) *RoomHandler {
	return &RoomHandler{
		roomService: service,
		roomHub:     roomHub,
	}
}

// CreateRoom создание комнаты
func (h *RoomHandler) CreateRoom(w http.ResponseWriter, r *http.Request) {
	var room model.RoomState
	logger := log.Ctx(r.Context())

	if err := json.NewDecoder(r.Body).Decode(&room); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	createdRoom, err := h.roomService.CreateRoom(r.Context(), &room)
	if err != nil {
		http.Error(w, "Failed to create room", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(createdRoom)
	if err != nil {
		logger.Error().Err(err).Msg("Metrics stopped")
	}
}

// JoinRoom функция входа в комнату
func (h *RoomHandler) JoinRoom(w http.ResponseWriter, r *http.Request) {
	logger := log.Ctx(r.Context())

	userID := r.URL.Query().Get("user_id")
	if userID == "" {
		http.Error(w, "Missing user_id", http.StatusBadRequest)
		return
	}

	cfg, err := config.New(false)
	if err != nil {
		http.Error(w, "error initialize app cfg", http.StatusInternalServerError)
	}

	ctx := config.WrapContext(r.Context(), cfg)

	sessionSrvResp, errSrvResp := h.roomService.Session(ctx, userID)
	if errSrvResp != nil {
		http.Error(w, "get session error", http.StatusInternalServerError)
	}

	user := model.User{
		ID:        sessionSrvResp.UserData.ID,
		AvatarURL: sessionSrvResp.UserData.AvatarURL,
		Username:  sessionSrvResp.UserData.Username,
		Email:     sessionSrvResp.UserData.Email,
	}

	roomID := r.URL.Query().Get("room_id")
	if roomID == "" {
		http.Error(w, "Missing room_id", http.StatusBadRequest)
		return
	}

	// Обновление соединения до WebSocket
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		http.Error(w, "Failed to upgrade to WebSocket", http.StatusInternalServerError)
		return
	}

	defer func(conn *websocket.Conn) {
		err := conn.Close()
		if err != nil {
			logger.Error().Err(err).Msg("Failed to close WS connect")
		}
	}(conn)

	// Регистрация клиента в комнате
	h.roomHub.RegisterClient(conn, roomID)
	h.roomHub.Users[conn] = user

	roomState, err := h.roomService.GetRoomState(r.Context(), roomID)
	if err != nil {
		logger.Error().Err(err).Msg("Failed to get room state from Redis")
	} else {
		if err := conn.WriteJSON(roomState); err != nil {
			logger.Error().Err(err).Msg("Failed to send room state")
			return
		}
	}

	h.broadcastUserList(roomID)

	for {
		var action model.Action
		if err := conn.ReadJSON(&action); err != nil {
			h.roomHub.Unregister <- conn
			delete(h.roomHub.Users, conn)
			h.broadcastUserList(roomID)
			break
		}

		h.roomHub.Broadcast <- ws.BroadcastMessage{Action: action, RoomID: roomID, ExcludeConn: conn}
		if err := h.roomService.HandleAction(r.Context(), roomID, action); err != nil {
			logger.Error().Err(err).Msg("Error handling action")
		}
	}
}

func (h *RoomHandler) broadcastUserList(roomID string) {
	userList := make([]model.User, 0)
	for conn := range h.roomHub.GetClients(roomID) {
		if user, ok := h.roomHub.Users[conn]; ok {
			userList = append(userList, user)
		}
	}

	for conn := range h.roomHub.GetClients(roomID) {
		if err := conn.WriteJSON(userList); err != nil {
			h.roomHub.Unregister <- conn
			delete(h.roomHub.Users, conn)
		}
	}
}
