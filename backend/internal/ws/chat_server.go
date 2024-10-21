package ws

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/gorilla/websocket"
	"go.uber.org/zap"
)

var upgrader = websocket.Upgrader{
	CheckOrigin:     checkReqOrigin,
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func checkReqOrigin(r *http.Request) bool {
	switch r.Header.Get("Origin") {
	// Only allow frontend server url to connect
	case "http://localhost:8000":
		return true
	default:
		return false
	}
}

type ChatServer struct {
	lg      *zap.SugaredLogger
	clients ClientList
	sync.RWMutex
	handlers map[string]EventHandler
}

func NewChatServer(logger *zap.SugaredLogger) *ChatServer {
	s := &ChatServer{
		lg:       logger,
		clients:  make(ClientList),
		handlers: make(map[string]EventHandler),
	}
	s.setupEventHandlers()
	return s
}

func (s *ChatServer) setupEventHandlers() {
	s.handlers[SendMessage] = func(e Event, c *Client) error {
		fmt.Println(e)
		return nil
	}
}

// Handler that allows ws connections
func (s *ChatServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Println("New connection")

	conn, err := upgrader.Upgrade(w, r, nil) // Upgrade http connection to ws
	if err != nil {
		log.Println(err)
		return
	}

	user := "testuser"
	// r.Header.Get("BearerToken")  // use token info to get username ??

	client := NewClient(s.lg, user, conn, s)
	s.addClient(client)

	go client.readMessages()
	go client.writeMessages()
}

func (s *ChatServer) addClient(client *Client) {
	s.Lock()
	defer s.Unlock()

	s.clients[client] = true
}

func (s *ChatServer) removeClient(client *Client) {
	s.Lock()
	defer s.Unlock()

	if _, ok := s.clients[client]; ok {
		client.conn.Close()
		delete(s.clients, client)
	}
}

func (s *ChatServer) broadcastMessage(message []byte, user string) error {
	s.lg.Debug("new message broadcasted")

	var returnMessage EgressMessageEvent
	if err := json.Unmarshal(message, &returnMessage); err != nil {
		return fmt.Errorf("bad payload in request: %v", err)
	}
	returnMessage.User = user
	returnMessage.Sent = time.Now()

	var outgoingEvent Event
	data, err := json.Marshal(returnMessage)
	if err != nil {
		return fmt.Errorf("failed to marshal broadcast message: %v", err)
	}
	outgoingEvent.Message = string(data)

	for client := range s.clients {
		client.egress <- outgoingEvent // broadcast to egress of all clients
	}
	return nil
}
