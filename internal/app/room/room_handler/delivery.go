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
	userID := r.URL.Query().Get("user_id")
	if userID == "" {
		http.Error(w, "Missing user_id", http.StatusBadRequest)
		return
	}
	/////////////////
	//log.Println("wwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwww")
	//ck, err := r.Cookie("session_id")
	//log.Println("ckckckck", ck)
	//if errors.Is(err, http.ErrNoCookie) {
	//	log.Println("Session action: No cookie err", err)
	//	return
	//}
	//log.Println("qwerrrewq")
	/////////////////////////////
	cfg, err := config.New(zerolog.Logger{}, false, nil)
	log.Println("qwer222rrewq")
	ctx := config.WrapContext(r.Context(), cfg)
	//log.Println("ck.Valueck.Value", ck.Value)

	sessionSrvResp, errSrvResp := h.roomService.Session(ctx, userID)
	log.Println("asdfdsaasdf1", sessionSrvResp)
	sessionResp, errResp := converter.ToApiSessionResponseForRoom(sessionSrvResp), converter.ToApiErrorResponseForRoom(errSrvResp)
	log.Println("asdfdsaasdf2")

	if errResp != nil {
		log.Println("errResp", errResp)
		api.Response(w, errResp.StatusCode, errResp)
		return
	}

	user := models.User{
		Id:        sessionResp.UserData.Id,
		AvatarUrl: sessionResp.UserData.AvatarUrl,
		Username:  sessionResp.UserData.Email,
		Email:     sessionResp.UserData.Username,
	}
	log.Println("xzxcvvcxzcvx")

	roomID := r.URL.Query().Get("room_id")
	if roomID == "" {
		http.Error(w, "Missing room_id", http.StatusBadRequest)
		return
	}
	log.Println("gfdfghfgd")

	// Обновление соединения до WebSocket
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		http.Error(w, "Failed to upgrade to WebSocket", http.StatusInternalServerError)
		return
	}
	defer conn.Close()
	log.Println("jghfhjghhyjy")

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
	log.Println("6789876")

	for {
		var action models.Action
		if err := conn.ReadJSON(&action); err != nil {
			// Если ошибка — отключаем клиента
			log.Println("Unregister action:", action.TimeCode)
			log.Println("Unregister action:", action.Name)
			h.roomHub.Unregister <- conn
			log.Println("Unregister:", action.TimeCode)
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
	log.Println("broadcastUserList:", userList)

	for conn := range h.roomHub.Clients {
		if err := conn.WriteJSON(userList); err != nil {
			h.roomHub.Unregister <- conn
			delete(h.roomHub.Users, conn)
		}
	}
}
