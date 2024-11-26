package websocket

// TODO раскоментить к 4му РК

//
//// BroadcastMessage представляет сообщение, которое будет отправлено в комнату
//type BroadcastMessage struct {
//	Action      interface{}     // действие, которое должно быть выполнено
//	RoomID      string          // идентификатор комнаты
//	ExcludeConn *websocket.Conn // соединение, которое не должно получать сообщение
//}
//
//type RoomHub struct {
//	Rooms      map[string]map[*websocket.Conn]bool // клиенты по комнатам
//	Users      map[*websocket.Conn]models.User     // соответствие клиента и пользователя
//	Register   chan *Client                        // канал регистрации клиентов
//	Unregister chan *websocket.Conn                // канал для удаления клиентов
//	Broadcast  chan BroadcastMessage               // канал для широковещательной рассылки сообщений
//	mu         sync.RWMutex                        // защита для работы с Rooms и Users
//}
//
//// Client представляет подключенного пользователя
//type Client struct {
//	Conn   *websocket.Conn
//	RoomID string
//}
//
//func NewRoomHub() *RoomHub {
//	return &RoomHub{
//		Rooms:      make(map[string]map[*websocket.Conn]bool),
//		Users:      make(map[*websocket.Conn]models.User),
//		Register:   make(chan *Client),
//		Unregister: make(chan *websocket.Conn),
//		Broadcast:  make(chan BroadcastMessage),
//	}
//}
//
//func (hub *RoomHub) Run() {
//	for {
//		select {
//		case client := <-hub.Register:
//			hub.addClientToRoom(client)
//		case conn := <-hub.Unregister:
//			hub.removeClient(conn)
//		case message := <-hub.Broadcast:
//			hub.broadcastToRoom(message)
//		}
//	}
//}
//
//// Регистрация клиента в комнате
//func (hub *RoomHub) RegisterClient(conn *websocket.Conn, roomID string) {
//	client := &Client{Conn: conn, RoomID: roomID}
//	hub.Register <- client
//}
//
//// Получение всех клиентов из указанной комнаты
//func (hub *RoomHub) GetClients(roomID string) map[*websocket.Conn]bool {
//	hub.mu.RLock()
//	defer hub.mu.RUnlock()
//	return hub.Rooms[roomID]
//}
//
//func (hub *RoomHub) addClientToRoom(client *Client) {
//	hub.mu.Lock()
//	defer hub.mu.Unlock()
//
//	// Создаем комнату, если её нет
//	if hub.Rooms[client.RoomID] == nil {
//		hub.Rooms[client.RoomID] = make(map[*websocket.Conn]bool)
//	}
//	hub.Rooms[client.RoomID][client.Conn] = true
//}
//
//func (hub *RoomHub) removeClient(conn *websocket.Conn) {
//	hub.mu.Lock()
//	defer hub.mu.Unlock()
//
//	for roomID, clients := range hub.Rooms {
//		if clients[conn] {
//			delete(clients, conn)
//			if len(clients) == 0 {
//				delete(hub.Rooms, roomID)
//			}
//			break
//		}
//	}
//	delete(hub.Users, conn)
//	conn.Close()
//}
//
//func (hub *RoomHub) broadcastToRoom(message BroadcastMessage) {
//	hub.mu.RLock()
//	defer hub.mu.RUnlock()
//
//	clients := hub.Rooms[message.RoomID]
//	for conn := range clients {
//		if conn != message.ExcludeConn {
//			if err := conn.WriteJSON(message.Action); err != nil {
//				hub.Unregister <- conn
//			}
//		}
//	}
//}
