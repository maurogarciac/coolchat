package handlers

import (
	"backend/internal/db"
	"encoding/json"

	"net/http"

	"go.uber.org/zap"
)

type MessageHandler struct {
	lg *zap.SugaredLogger
	db *db.DbProvider
}

func NewMessageHandler(
	logger *zap.SugaredLogger,
	database db.DbProvider) *MessageHandler {

	return &MessageHandler{
		lg: logger,
		db: &database,
	}
}

func (h *MessageHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	switch r.Method {

	case http.MethodGet:
		messages, err := h.db.SelectAllMessages()
		if err != nil {
			h.lg.Errorf("error getting messages: ", err)
			http.Error(w, "Could not get messages", http.StatusInternalServerError)

		}
		h.lg.Debug(messages)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(messages.MessageList)

	default:
		h.lg.Error("only get method allowed")
		http.Error(w, "Only GET method allowed", http.StatusMethodNotAllowed)
	}
}
