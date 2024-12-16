package websocket

import (
	models "github.com/go-park-mail-ru/2024_2_GOATS/internal/app/room/model"
	"github.com/gorilla/websocket"

	"log"
	"sync"
	"time"
)

// BroadcastMessage структура сообщения рассылаемое бродкастом
type BroadcastMessage struct {
	Action      interface{}
	RoomID      string
	ExcludeConn *websocket.Conn
}

// RoomHub структура хаба
type RoomHub struct {
	Rooms        map[string]map[*websocket.Conn]bool
	Users        map[*websocket.Conn]models.User
	Register     chan *Client
	Unregister   chan *websocket.Conn
	Broadcast    chan BroadcastMessage
	mu           sync.RWMutex
	timerManager *TimerManager
}

// Client структура клиента
type Client struct {
	Conn   *websocket.Conn
	RoomID string
}

// NewRoomHub конструктор хаба
func NewRoomHub() *RoomHub {
	return &RoomHub{
		Rooms:      make(map[string]map[*websocket.Conn]bool),
		Users:      make(map[*websocket.Conn]models.User),
		Register:   make(chan *Client),
		Unregister: make(chan *websocket.Conn),
		Broadcast:  make(chan BroadcastMessage),
	}
}

// Run запуск
func (hub *RoomHub) Run() {
	for {
		select {
		case client := <-hub.Register:
			hub.addClientToRoom(client)
		case conn := <-hub.Unregister:
			hub.removeClient(conn)
		case message := <-hub.Broadcast:
			hub.broadcastToRoom(message)
		}
	}
}

// RegisterClient функция регистрации клиента
func (hub *RoomHub) RegisterClient(conn *websocket.Conn, roomID string) {
	clients := &Client{Conn: conn, RoomID: roomID}
	hub.Register <- clients
}

// GetClients функция получения клиента
func (hub *RoomHub) GetClients(roomID string) map[*websocket.Conn]bool {
	hub.mu.RLock()
	defer hub.mu.RUnlock()
	return hub.Rooms[roomID]
}

func (hub *RoomHub) addClientToRoom(client *Client) {
	hub.mu.Lock()
	defer hub.mu.Unlock()

	// Создаем комнату, если её нет
	if hub.Rooms[client.RoomID] == nil {
		hub.Rooms[client.RoomID] = make(map[*websocket.Conn]bool)
	}
	hub.Rooms[client.RoomID][client.Conn] = true
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
	err := conn.Close()
	if err != nil {
		log.Println("close websocket err:", err)
	}
}

// broadcastToRoom бродкаст для комнаты
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

// SetTimerManager установка таймера
func (hub *RoomHub) SetTimerManager(manager *TimerManager) {
	hub.timerManager = manager
}

// TimerManager структура таймера
type TimerManager struct {
	mu     sync.Mutex
	timers map[string]chan struct{}
	hub    *RoomHub
}

// NewTimerManager конструктор таймера
func NewTimerManager(hub *RoomHub) *TimerManager {
	return &TimerManager{
		timers: make(map[string]chan struct{}),
		hub:    hub,
	}
}

// Start старт таймерв
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

// Stop остановка таймера
func (tm *TimerManager) Stop(roomID string) {
	tm.mu.Lock()
	if quit, exists := tm.timers[roomID]; exists {
		close(quit)
		delete(tm.timers, roomID)
	}
	tm.mu.Unlock()
}
