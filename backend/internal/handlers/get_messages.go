package handlers

import (
	"backend/config"
	"backend/internal/db"

	"context"
	"net/http"

	"go.uber.org/zap"
)

type MessageHandler struct {
	context   context.Context
	appConfig config.AppConfig
	lg        *zap.SugaredLogger
	db        *db.DbProvider
}

func NewMessageHandler(
	context context.Context,
	config config.AppConfig,
	logger *zap.SugaredLogger,
	database db.DbProvider) *MessageHandler {

	return &MessageHandler{
		context:   context,
		appConfig: config,
		lg:        logger,
		db:        &database,
	}
}

func (h *MessageHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	switch r.Method {

	case http.MethodGet:
		messages, err := h.db.SelectAllMessages()
		if err != nil {
			h.lg.Errorf("error getting messages: ", err)
		}
		h.lg.Info(messages)
	default:
		h.lg.Error("only get method allowed")
	}
}
