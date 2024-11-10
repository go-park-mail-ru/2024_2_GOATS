package delivery

import (
	"context"
	"encoding/json"
	"github.com/go-park-mail-ru/2024_2_GOATS/config"
	"github.com/go-park-mail-ru/2024_2_GOATS/internal/app/api"
	"github.com/go-park-mail-ru/2024_2_GOATS/internal/app/api/converter"
	models "github.com/go-park-mail-ru/2024_2_GOATS/internal/app/room/model"
	ws "github.com/go-park-mail-ru/2024_2_GOATS/internal/app/room/ws"
	"github.com/gorilla/websocket"
	"github.com/rs/zerolog"
	"log"
	"net/http"
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
	json.NewEncoder(w).Encode(createdRoom)
}

func (h *RoomHandler) JoinRoom(w http.ResponseWriter, r *http.Request) {
	userID := r.URL.Query().Get("user_id")
	if userID == "" {
		http.Error(w, "Missing user_id", http.StatusBadRequest)
		return
	}

	cfg, err := config.New(zerolog.Logger{}, false, nil)
	ctx := config.WrapContext(r.Context(), cfg)

	sessionSrvResp, errSrvResp := h.roomService.Session(ctx, userID)
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

	// Регистрация клиента в комнате
	h.roomHub.RegisterClient(conn, roomID)
	h.roomHub.Users[conn] = user

	roomState, err := h.roomService.GetRoomState(r.Context(), roomID)
	if err != nil {
		log.Println("Failed to get room state from Redis:", err)
	} else {
		if err := conn.WriteJSON(roomState); err != nil {
			log.Println("Failed to send room state:", err)
			return
		}
	}

	h.broadcastUserList(conn, roomID)

	for {
		var action models.Action
		if err := conn.ReadJSON(&action); err != nil {
			h.roomHub.Unregister <- conn
			delete(h.roomHub.Users, conn)
			h.broadcastUserList(conn, roomID)
			break
		}

		h.roomHub.Broadcast <- ws.BroadcastMessage{Action: action, RoomID: roomID, ExcludeConn: conn}
		if err := h.roomService.HandleAction(r.Context(), roomID, action); err != nil {
			log.Println("Error handling action:", err)
		}
	}
}

//func (h *RoomHandler) broadcastUserList(excludeConn *websocket.Conn, roomID string) {
//	userList := make([]models.User, 0, len(h.roomHub.Users))
//	for _, user := range h.roomHub.Users {
//		userList = append(userList, user)
//	}
//
//	for conn := range h.roomHub.GetClients(roomID) {
//		if err := conn.WriteJSON(userList); err != nil {
//			h.roomHub.Unregister <- conn
//			delete(h.roomHub.Users, conn)
//		}
//	}
//}

//func (h *RoomHandler) broadcastUserList(excludeConn *websocket.Conn, roomID string) {
//	// Получаем пользователей, которые находятся только в указанной комнате
//	userList := make([]models.User, 0)
//	for conn := range h.roomHub.GetClients(roomID) {
//		// Извлекаем пользователя, связанного с каждым соединением, если он существует
//		if user, ok := h.roomHub.Users[conn]; ok {
//			userList = append(userList, user)
//		}
//	}
//
//	// Рассылаем обновленный список пользователей всем клиентам в указанной комнате
//	for conn := range h.roomHub.GetClients(roomID) {
//		if conn != excludeConn { // исключаем отправителя
//			if err := conn.WriteJSON(userList); err != nil {
//				h.roomHub.Unregister <- conn
//				delete(h.roomHub.Users, conn)
//			}
//		}
//	}
//}

func (h *RoomHandler) broadcastUserList(excludeConn *websocket.Conn, roomID string) {
	// Получаем список пользователей только для комнаты roomID
	userList := make([]models.User, 0)
	for conn := range h.roomHub.GetClients(roomID) {
		if user, ok := h.roomHub.Users[conn]; ok {
			userList = append(userList, user)
		}
	}

	// Отправляем список пользователей всем подключенным клиентам в комнате
	for conn := range h.roomHub.GetClients(roomID) {
		if conn != excludeConn { // исключаем соединение, не нуждающееся в обновлении
			if err := conn.WriteJSON(userList); err != nil {
				h.roomHub.Unregister <- conn
				delete(h.roomHub.Users, conn)
			}
		}
	}
}
