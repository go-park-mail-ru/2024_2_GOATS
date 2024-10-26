package delivery

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/go-park-mail-ru/2024_2_GOATS/config"
	"github.com/go-park-mail-ru/2024_2_GOATS/internal/app/api"
	"github.com/go-park-mail-ru/2024_2_GOATS/internal/app/api/converter"
	models "github.com/go-park-mail-ru/2024_2_GOATS/internal/app/room/model"
	ws "github.com/go-park-mail-ru/2024_2_GOATS/internal/app/room/ws"
	"github.com/gorilla/websocket"
	"github.com/rs/zerolog"
	"log"
	"net/http"
	"os"
)

type RoomServiceInterface interface {
	CreateRoom(ctx context.Context, room *models.RoomState) (*models.RoomState, error)
	HandleAction(ctx context.Context, roomID string, action models.Action) error
	Session(ctx context.Context, cookie string) (*models.SessionRespData, *models.ErrorRespData)
	GetRoomState(ctx context.Context, roomID string) (*models.RoomState, error)
}

type RoomHandler struct {
	roomService RoomServiceInterface
	roomHub     *ws.RoomHub
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true },
}

func NewRoomHandler(service RoomServiceInterface, roomHub *ws.RoomHub) *RoomHandler {
	return &RoomHandler{
		roomService: service,
		roomHub:     roomHub,
	}
}

func (h *RoomHandler) CreateRoom(w http.ResponseWriter, r *http.Request) {
	var room models.RoomState
	log.Println("room", room)
	if err := json.NewDecoder(r.Body).Decode(&room); err != nil {
		log.Println("err", err)
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	log.Println("room", room)
	createdRoom, err := h.roomService.CreateRoom(r.Context(), &room)
	if err != nil {
		http.Error(w, "Failed to create room", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(createdRoom)
}

func (h *RoomHandler) JoinRoom(w http.ResponseWriter, r *http.Request) {
	ck, err := r.Cookie("session_id")
	if errors.Is(err, http.ErrNoCookie) {
		log.Println("Session action: No cookie err", err)
		return
	}

	logger := zerolog.New(os.Stdout).With().Caller().Timestamp().Logger()

	cfg, err := config.New(logger, true, nil)
	ctx := config.WrapContext(r.Context(), cfg)
	sessionSrvResp, errSrvResp := h.roomService.Session(ctx, ck.Value)

	sessionResp, errResp := converter.ToApiSessionResponseForRoom(sessionSrvResp), converter.ToApiErrorResponseForRoom(errSrvResp)

	if errResp != nil {
		api.Response(w, errResp.StatusCode, errResp)
		return
	}

	user := models.User{
		Id:        sessionResp.UserData.Id,
		AvatarUrl: sessionResp.UserData.AvatarUrl,
		Username:  sessionResp.UserData.Email,
		Email:     sessionResp.UserData.Username,
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
	defer conn.Close()

	// Регистрация клиента в Hub
	h.roomHub.Register <- conn
	h.roomHub.Users[conn] = user

	h.broadcastUserList()

	// Получение состояния комнаты
	roomState, err := h.roomService.GetRoomState(r.Context(), roomID)
	log.Println("roomStateroomStateroomStateroomState:", roomState)

	if err != nil {
		log.Println("Failed to get room state from Redis:", err)
	} else {
		// Отправка текущего состояния новому пользователю
		if err := conn.WriteJSON(roomState); err != nil {
			log.Println("Failed to send room state:", err)
			return
		}
	}

	for {
		var action models.Action
		if err := conn.ReadJSON(&action); err != nil {
			// Если ошибка — отключаем клиента
			log.Println("Unregister action:", action.TimeCode)
			log.Println("Unregister action:", action.Name)
			h.roomHub.Unregister <- conn
			delete(h.roomHub.Users, conn)
			h.broadcastUserList()
			break
		}
		log.Println("Received action:", action.TimeCode)
		log.Println("Received action:", action.Name)
		// Отправляем сообщение в Hub для рассылки всем клиентам
		h.roomHub.Broadcast <- action
		// Обработка действия и обновление состояния комнаты
		if err := h.roomService.HandleAction(r.Context(), roomID, action); err != nil {
			log.Println("Error handling action:", err)
		}
	}

}

func (h *RoomHandler) broadcastUserList() {
	userList := make([]models.User, 0, len(h.roomHub.Users))
	for _, user := range h.roomHub.Users {
		userList = append(userList, user)
	}

	for conn := range h.roomHub.Clients {
		if err := conn.WriteJSON(userList); err != nil {
			h.roomHub.Unregister <- conn
			delete(h.roomHub.Users, conn)
		}
	}
}
