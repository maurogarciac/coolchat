package handlers

import (
	"net/http"

	"frontend/internal/services"
	"frontend/internal/templates"

	"go.uber.org/zap"
)

type ChatHandler struct {
	lg *zap.SugaredLogger
	b  services.BackendService
}

func NewChatHandler(logger *zap.SugaredLogger, backend services.BackendService) *ChatHandler {
	return &ChatHandler{
		lg: logger,
		b:  backend,
	}
}

func (h ChatHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	user, ok := r.Context().Value("User").(string)
	if !ok {
		h.lg.Error("Could not fetch username from context")
	}

	messageHistory, err := h.b.GetMessageHistory(r.Context())
	if err != nil {
		h.lg.Error(err)
		http.Error(w, "Could not fetch message history", http.StatusInternalServerError)
	}

	h.lg.Info(messageHistory.Messages)

	c := templates.ChatBox(user)

	switch r.Method {

	case http.MethodGet:

		pageRender("chat", c, true, h.lg, w, r)

	default:
		h.lg.Error("Only GET method is supported for CHAT")
		http.Error(w, "Only GET method is supported", http.StatusMethodNotAllowed)
	}

}
