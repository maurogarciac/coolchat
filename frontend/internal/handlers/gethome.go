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

	name := "home"
	c := templates.Home()

	switch r.Method {
	case http.MethodGet:

		// should check headers to see if user has an access token

		pageRender(name, c, h.lg, w, r)

	default:
		fmt.Fprintf(w, "only get method is supported")
		return
	}

}
