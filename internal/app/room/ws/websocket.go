package websocket

import (
	models "github.com/go-park-mail-ru/2024_2_GOATS/internal/app/room/model"
	"github.com/gorilla/websocket"
	"net/http"
)

type RoomHub struct {
	Clients    map[*websocket.Conn]bool
	Broadcast  chan models.Action
	Register   chan *websocket.Conn
	Unregister chan *websocket.Conn
}

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

// NewRoomHub создает новый RoomHub
func NewRoomHub() *RoomHub {
	return &RoomHub{
		Clients:    make(map[*websocket.Conn]bool),
		Broadcast:  make(chan models.Action),
		Register:   make(chan *websocket.Conn),
		Unregister: make(chan *websocket.Conn),
	}
}

// HandleConnections обрабатывает новое WebSocket соединение
func (h *RoomHub) HandleConnections(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		http.Error(w, "Failed to upgrade connection", http.StatusInternalServerError)
		return
	}
	defer conn.Close()

	h.Register <- conn //регистрируем новое соединение

	for {
		var msg models.Action
		err := conn.ReadJSON(&msg)
		if err != nil {
			h.Unregister <- conn //при ошибке — удаляем соединение
			break
		}
		h.Broadcast <- msg //отправляем сообщение через Broadcast
	}
}

// Run запускает RoomHub и обрабатывает регистрацию, отключение и рассылку сообщений
func (h *RoomHub) Run() {
	for {
		select {
		case conn := <-h.Register:
			h.Clients[conn] = true
		case conn := <-h.Unregister:
			if _, ok := h.Clients[conn]; ok {
				delete(h.Clients, conn)
				conn.Close()
			}
		case msg := <-h.Broadcast:
			for conn := range h.Clients {
				err := conn.WriteJSON(msg)
				if err != nil {
					h.Unregister <- conn
				}
			}
		}
	}
}
