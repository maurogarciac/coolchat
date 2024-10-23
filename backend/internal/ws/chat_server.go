package ws

import (
	"backend/internal/db"
	"backend/internal/domain"
	"encoding/json"
	"fmt"
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
	db       *db.DbProvider
}

func NewChatServer(logger *zap.SugaredLogger, database *db.DbProvider) *ChatServer {
	s := &ChatServer{
		lg:       logger,
		clients:  make(ClientList),
		handlers: make(map[string]EventHandler),
		db:       database,
	}
	s.setupEventHandlers()
	return s
}

func (s *ChatServer) setupEventHandlers() {
	s.handlers[SendMessage] = func(e Event, c *Client) error {
		s.lg.Debug(e)
		return nil
	}
}

// Handler that allows ws connections
func (s *ChatServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.lg.Info("New connection")

	conn, err := upgrader.Upgrade(w, r, nil) // Upgrade http connection to ws
	if err != nil {
		s.lg.Error(err)
		return
	}

	client := NewClient(s.lg, conn, s)
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

func (s *ChatServer) broadcastMessage(message []byte) error {
	s.lg.Debug("new message broadcasted")

	var recievedMessage RecievedEgressMessage
	if err := json.Unmarshal(message, &recievedMessage); err != nil {
		return fmt.Errorf("bad payload in request: %v", err)
	}

	s.lg.Debugf("Recieved msg: %s", recievedMessage)

	var messageContent InnerMessage
	if err := json.Unmarshal([]byte(recievedMessage.Message), &messageContent); err != nil {
		return fmt.Errorf("error in message content in request: %v", err)
	}

	s.lg.Infof("Message content: %s", messageContent)

	returnMessage := ReturnMessage{
		Text:      messageContent.Text,
		User:      messageContent.User,
		Timestamp: time.Now(),
	}

	var outgoingEvent Event
	data, err := json.Marshal(returnMessage)
	if err != nil {
		return fmt.Errorf("failed to marshal broadcast message: %v", err)
	}
	outgoingEvent.Message = string(data)

	msg := domain.InsertMessage{
		Text: returnMessage.Text,
		From: returnMessage.User,
	}

	if s.db == nil {
		errStr := "could not save message. database connection is not initialized"
		return fmt.Errorf("%s", errStr)
	}
	if _, err := s.db.InsertMessage(msg); err != nil {
		return fmt.Errorf("failed to insert message: %v", err)
	}

	for client := range s.clients {
		client.egress <- outgoingEvent // broadcast to egress of all clients
	}
	return nil
}
