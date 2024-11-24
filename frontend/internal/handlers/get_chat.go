package handlers

import (
	"net/http"
	"time"

	d "frontend/internal/domain"
	"frontend/internal/services"
	"frontend/internal/templates"

	"github.com/a-h/templ"
	"go.uber.org/zap"
)

type ChatHandler struct {
	lg *zap.SugaredLogger
	b  services.BackendService
	ip string
}

func NewChatHandler(logger *zap.SugaredLogger, backend services.BackendService, serverIp string) *ChatHandler {
	return &ChatHandler{
		lg: logger,
		b:  backend,
		ip: serverIp,
	}
}

func (h ChatHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	switch r.Method {

	case http.MethodGet:

		h.lg.Info("Chat GET")

		user, ok := r.Context().Value("User").(string)
		if !ok {
			h.lg.Error("Could not fetch username from context")
		}

		messageHistory, err := h.b.GetMessageHistory(r.Context())

		var c templ.Component

		if err != nil {
			h.lg.Errorf("could not fetch message history: %s", err)

			m := []d.Message{
				{User: "Error", Text: "Could not fetch message history", Timestamp: time.Now().String()},
			}
			mh := d.MessageHistoryResult{
				Messages: m,
			}

			c = templates.ChatBox(user, mh, h.ip)

		} else {
			h.lg.Debugf("Message history result: %s", messageHistory.Messages)

			c = templates.ChatBox(user, messageHistory, h.ip)

		}
		pageRender("chat", c, true, h.lg, w, r)

	default:
		h.lg.Error("Only GET method is supported for CHAT")
		http.Error(w, "Only GET method is supported", http.StatusMethodNotAllowed)
	}

}
