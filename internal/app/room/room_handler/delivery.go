package delivery

import (
	"context"
	"fmt"

	"github.com/go-park-mail-ru/2024_2_GOATS/config"
	"github.com/go-park-mail-ru/2024_2_GOATS/internal/app/api"
	errVals "github.com/go-park-mail-ru/2024_2_GOATS/internal/app/errors"
	model "github.com/go-park-mail-ru/2024_2_GOATS/internal/app/room/model"
	ws "github.com/go-park-mail-ru/2024_2_GOATS/internal/app/room/ws"
	websocket "github.com/gorilla/websocket"

	//zlog "github.com/rs/zerolog/log"
	"log"
	"net/http"
)

type RoomServiceInterface interface {
	CreateRoom(ctx context.Context, room *model.RoomState) (*model.RoomState, error)
	HandleAction(ctx context.Context, roomID string, action model.Action) error
	Session(ctx context.Context, cookie string) (*model.SessionRespData, *errVals.ServiceError)
	GetRoomState(ctx context.Context, roomID string) (*model.RoomState, error)
}

type RoomHandler struct {
	roomService RoomServiceInterface
	roomHub     *ws.RoomHub
	cfg         *config.Config
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true },
}

func NewRoomHandler(service RoomServiceInterface, roomHub *ws.RoomHub, cfg *config.Config) *RoomHandler {
	return &RoomHandler{
		roomService: service,
		roomHub:     roomHub,
		cfg:         cfg,
	}
}

func (h *RoomHandler) CreateRoom(w http.ResponseWriter, r *http.Request) {
	room := &model.RoomState{}
	if !api.DecodeBody(w, r, room) {
		return
	}

	createdRoom, err := h.roomService.CreateRoom(r.Context(), room)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to create room: %v", err), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	api.Response(r.Context(), w, http.StatusOK, createdRoom)
}

func (h *RoomHandler) JoinRoom(w http.ResponseWriter, r *http.Request) {
	userID := r.URL.Query().Get("user_id")
	if userID == "" {
		http.Error(w, "Missing user_id", http.StatusBadRequest)
		return
	}

	ctx := config.WrapContext(r.Context(), h.cfg)

	sessionSrvResp, errSrvResp := h.roomService.Session(ctx, userID)

	if errSrvResp != nil {
		return
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
		log.Println(err)
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
		var action model.Action
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
