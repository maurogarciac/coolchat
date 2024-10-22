package handlers

import (
	"fmt"
	"net/http"

	"frontend/internal/templates"

	"go.uber.org/zap"
)

type HomeHandler struct {
	lg *zap.SugaredLogger
}

func NewHomeHandler(logger *zap.SugaredLogger) *HomeHandler {
	return &HomeHandler{
		lg: logger,
	}
}

func (h HomeHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	user, ok := r.Context().Value("User").(string)
	if !ok {
		h.lg.Error("Could not fetch username from context")
	}

	c := templates.Home(user)

	switch r.Method {
	case http.MethodGet:

		pageRender("home", c, true, h.lg, w, r)

	default:

		fmt.Fprintf(w, "only get method is supported")
		return
	}

}
