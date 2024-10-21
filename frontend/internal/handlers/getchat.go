package handlers

import (
	"context"
	"fmt"
	"net/http"

	"frontend/internal/templates"

	"go.uber.org/zap"
)

type ChatHandler struct {
	ctx context.Context
	lg  *zap.SugaredLogger
}

func NewChatHandler(context context.Context, logger *zap.SugaredLogger) *ChatHandler {
	return &ChatHandler{
		ctx: context,
		lg:  logger,
	}
}

func (h ChatHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	user, ok := h.ctx.Value("User").(string)
	if !ok {
		h.lg.Error("Could not fetch username from context")
	}

	c := templates.ChatBox(user)

	switch r.Method {

	case http.MethodGet:

		// should check headers to see if user has an access token

		pageRender("chat", c, h.lg, w, r)

	default:
		fmt.Fprintf(w, "only get method is supported")
		return
	}

}
