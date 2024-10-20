package handlers

import (
	"fmt"
	"net/http"

	"frontend/internal/templates"

	"go.uber.org/zap"
)

type ChatHandler struct {
	lg *zap.SugaredLogger
}

func NewChatHandler(logger *zap.SugaredLogger) *ChatHandler {
	return &ChatHandler{
		lg: logger,
	}
}

func (h ChatHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	name := "chat"
	c := templates.ChatBox("joe")

	switch r.Method {

	case http.MethodGet:

		// should check headers to see if user has an access token

		pageRender(name, c, h.lg, w, r)

	default:
		fmt.Fprintf(w, "only get method is supported")
		return
	}

}
