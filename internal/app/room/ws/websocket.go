package websocket

import (
	models "github.com/go-park-mail-ru/2024_2_GOATS/internal/app/room/model"
	"github.com/gorilla/websocket"
	"sync"
	"time"
)

type BroadcastMessage struct {
	Action      interface{}
	RoomID      string
	ExcludeConn *websocket.Conn
}

type RoomHub struct {
	Rooms        map[string]map[*websocket.Conn]bool
	Users        map[*websocket.Conn]models.User
	Register     chan *Client
	Unregister   chan *websocket.Conn
	Broadcast    chan BroadcastMessage
	mu           sync.RWMutex
	timerManager *TimerManager
}

type Client struct {
	Conn   *websocket.Conn
	RoomID string
}

func NewRoomHub() *RoomHub {
	return &RoomHub{
		Rooms:      make(map[string]map[*websocket.Conn]bool),
		Users:      make(map[*websocket.Conn]models.User),
		Register:   make(chan *Client),
		Unregister: make(chan *websocket.Conn),
		Broadcast:  make(chan BroadcastMessage),
	}
}

func (hub *RoomHub) Run() {
	for {
		select {
		case clients := <-hub.Register:
			hub.addClientToRoom(clients)
		case conn := <-hub.Unregister:
			hub.removeClient(conn)
		case message := <-hub.Broadcast:
			hub.broadcastToRoom(message)
		}
	}
}

func (hub *RoomHub) RegisterClient(conn *websocket.Conn, roomID string) {
	clients := &Client{Conn: conn, RoomID: roomID}
	hub.Register <- clients
}

func (hub *RoomHub) GetClients(roomID string) map[*websocket.Conn]bool {
	hub.mu.RLock()
	defer hub.mu.RUnlock()
	return hub.Rooms[roomID]
}

func (hub *RoomHub) addClientToRoom(clients *Client) {
	hub.mu.Lock()
	defer hub.mu.Unlock()

	// Создаем комнату, если её нет
	if hub.Rooms[clients.RoomID] == nil {
		hub.Rooms[clients.RoomID] = make(map[*websocket.Conn]bool)
	}
	hub.Rooms[clients.RoomID][clients.Conn] = true
}

func (hub *RoomHub) removeClient(conn *websocket.Conn) {
	hub.mu.Lock()
	defer hub.mu.Unlock()

	for roomID, clients := range hub.Rooms {
		if clients[conn] {
			delete(clients, conn)
			if len(clients) == 0 {
				delete(hub.Rooms, roomID)

				if hub.timerManager != nil {
					hub.timerManager.Stop(roomID)
				}
			}
			break
		}
	}
	delete(hub.Users, conn)
	conn.Close()
}

func (hub *RoomHub) broadcastToRoom(message BroadcastMessage) {
	hub.mu.RLock()
	defer hub.mu.RUnlock()

	clients := hub.Rooms[message.RoomID]
	for conn := range clients {
		if conn != message.ExcludeConn {
			if err := conn.WriteJSON(message.Action); err != nil {
				hub.Unregister <- conn
			}
		}
	}
}

func (hub *RoomHub) SetTimerManager(manager *TimerManager) {
	hub.timerManager = manager
}

type TimerManager struct {
	mu     sync.Mutex
	timers map[string]chan struct{}
	hub    *RoomHub
}

func NewTimerManager(hub *RoomHub) *TimerManager {
	return &TimerManager{
		timers: make(map[string]chan struct{}),
		hub:    hub,
	}
}
func (tm *TimerManager) Start(roomID string, startTime int64, updateFunc func(int64), duration int64) {
	tm.mu.Lock()
	if _, exists := tm.timers[roomID]; exists {
		tm.mu.Unlock()
		return
	}
	quit := make(chan struct{})
	tm.timers[roomID] = quit
	tm.mu.Unlock()

	go func() {
		timeCode := startTime
		ticker := time.NewTicker(3 * time.Second)
		defer ticker.Stop()

		for {
			select {
			case <-ticker.C:
				timeCode += 3

				if duration > 0 && timeCode >= duration {
					tm.Stop(roomID)
					return
				}

				tm.hub.Broadcast <- BroadcastMessage{
					Action: map[string]interface{}{
						"type":     "timer",
						"timeCode": timeCode,
					},
					RoomID: roomID,
				}
				updateFunc(timeCode)
			case <-quit:
				return
			}
		}
	}()
}

func (tm *TimerManager) Stop(roomID string) {
	tm.mu.Lock()
	if quit, exists := tm.timers[roomID]; exists {
		close(quit)
		delete(tm.timers, roomID)
	}
	tm.mu.Unlock()
}
