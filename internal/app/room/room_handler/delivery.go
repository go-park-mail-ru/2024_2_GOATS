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
	"github.com/rs/zerolog/log"
	ll "log"
	"net/http"
	"strconv"
)

// RoomServiceInterface defines methods for room service layer
type RoomServiceInterface interface {
	CreateRoom(ctx context.Context, room *model.RoomState) (*model.RoomState, error)
	HandleAction(ctx context.Context, roomID string, action model.Action) error
	Session(ctx context.Context, cookie string) (*model.SessionRespData, *errVals.ServiceError)
	GetRoomState(ctx context.Context, roomID string) (*model.RoomState, error)
}

// RoomHandler handler struct
type RoomHandler struct {
	roomService RoomServiceInterface
	roomHub     *ws.RoomHub
	cfg         *config.Config
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(_ *http.Request) bool { return true },
}

// NewRoomHandler returns an instance of RoomHandler
func NewRoomHandler(service RoomServiceInterface, roomHub *ws.RoomHub, cfg *config.Config) *RoomHandler {
	return &RoomHandler{
		roomService: service,
		roomHub:     roomHub,
		cfg:         cfg,
	}
}

// CreateRoom creates room
func (h *RoomHandler) CreateRoom(w http.ResponseWriter, r *http.Request) {
	room := &model.RoomState{}
	logger := log.Ctx(r.Context())

	if !api.DecodeBody(w, r, room) {
		return
	}

	//userID := r.URL.Query().Get("user_id")
	//if userID == "" {
	//	http.Error(w, "Missing user_id", http.StatusBadRequest)
	//	return
	//}

	//ctx := config.WrapContext(r.Context(), h.cfg)
	//
	//sessionSrvResp, errSrvResp := h.roomService.Session(ctx, userID)
	//if errSrvResp != nil {
	//	http.Error(w, "get session error", http.StatusInternalServerError)
	//}

	//if !sessionSrvResp.UserData.SubscriptionStatus {
	createdRoom, err := h.roomService.CreateRoom(r.Context(), room)
	if err != nil {
		logger.Error().Err(err).Msg("cannot_create_room")
		http.Error(w, fmt.Sprintf("Failed to create room: %v", err), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	api.Response(r.Context(), w, http.StatusOK, createdRoom)
	//} else {
	//	w.Header().Set("Content-Type", "application/json")
	//	api.Response(r.Context(), w, http.StatusBadRequest, fmt.Errorf("no subscription"))
	//}
}

func (h *RoomHandler) JoinRoom(w http.ResponseWriter, r *http.Request) {
	logger := log.Ctx(r.Context())

	userID := r.URL.Query().Get("user_id")
	if userID == "" {
		http.Error(w, "Missing user_id", http.StatusBadRequest)
		return
	}

	ctx := config.WrapContext(r.Context(), h.cfg)
	sessionSrvResp, errSrvResp := h.roomService.Session(ctx, userID)
	if errSrvResp != nil {
		http.Error(w, "get session error", http.StatusInternalServerError)
		return
		//ll.Println("getsessionerror")
	}

	user := model.User{
		ID:        sessionSrvResp.UserData.ID,
		AvatarURL: sessionSrvResp.UserData.AvatarURL,
		Username:  sessionSrvResp.UserData.Username,
		Email:     sessionSrvResp.UserData.Email,
	}

	roomID := r.URL.Query().Get("room_id")
	if roomID == "" {
		logger.Error().Msg("empty room id")
		http.Error(w, "Missing room_id", http.StatusBadRequest)
		return
	}

	ErrorAlreadyConnection := false
	ErrorManyConnections := false

	clients := h.roomHub.Rooms[roomID]

	for conn := range clients {
		usrId, err := strconv.Atoi(userID)
		if err != nil {
			return
		}

		if h.roomHub.Users[conn].ID == usrId {
			ll.Println("already_connected11111")
			ErrorAlreadyConnection = true
		}

	}

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		logger.Error().Err(err).Msg("Failed to upgrade to WebSocket")
		http.Error(w, "Failed to upgrade to WebSocket", http.StatusInternalServerError)
		return
	}

	if len(clients)+1 > 2 {
		ErrorManyConnections = true
	}

	if ErrorAlreadyConnection == true {
		Action := map[string]interface{}{
			"name": "already_connected",
		}
		if err := conn.WriteJSON(Action); err != nil {
			ll.Println("already_connected", conn.WriteJSON(Action))
		}
		return
	}
	if ErrorManyConnections == true {
		Action := map[string]interface{}{
			"name": "many_connections",
		}
		if err := conn.WriteJSON(Action); err != nil {
			ll.Println("many_connections", conn.WriteJSON(Action))
		}
		return
	}

	defer func(conn *websocket.Conn) {
		err := conn.Close()
		if err != nil {
			logger.Error().Err(err).Msg("Failed to close WS connect")
		}
	}(conn)

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

		validActions := []string{"pause", "play", "rewind", "message", "change", "change_series", "duration", "already_connected"}
		isValid := false

		for _, validAction := range validActions {
			if action.Name == validAction {
				isValid = true
				break
			}
		}

		if !isValid {
			logger.Error().Msg("Invalid action")
		} else {
			h.roomHub.Broadcast <- ws.BroadcastMessage{Action: action, RoomID: roomID, ExcludeConn: conn}
			if err := h.roomService.HandleAction(r.Context(), roomID, action); err != nil {
				logger.Error().Err(err).Msg("Error handling action")
			}
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
