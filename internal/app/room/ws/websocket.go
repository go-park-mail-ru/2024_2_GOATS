package websocket

import (
	models "github.com/go-park-mail-ru/2024_2_GOATS/internal/app/room/model"
	"github.com/gorilla/websocket"
)

type RoomHub struct {
	Clients    map[*websocket.Conn]bool
	Users      map[*websocket.Conn]models.User
	Broadcast  chan BroadcastMessage
	Register   chan *websocket.Conn
	Unregister chan *websocket.Conn
}

type BroadcastMessage struct {
	Action      models.Action   `json:"action"`
	ExcludeConn *websocket.Conn `json:"exclude"`
}

func NewRoomHub() *RoomHub {
	return &RoomHub{
		Clients:    make(map[*websocket.Conn]bool),
		Users:      make(map[*websocket.Conn]models.User),
		Broadcast:  make(chan BroadcastMessage),
		Register:   make(chan *websocket.Conn),
		Unregister: make(chan *websocket.Conn),
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
				if msg.ExcludeConn == conn {
					continue
				}
				err := conn.WriteJSON(msg)
				if err != nil {
					h.Unregister <- conn
				}
			}
		}
	}
}
