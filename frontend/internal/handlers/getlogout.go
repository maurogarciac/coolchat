package handlers

import (
	"fmt"
	"net/http"

	"frontend/internal/templates"

	"go.uber.org/zap"
)

type LogoutHandler struct {
	lg *zap.SugaredLogger
}

func NewLogoutHandler(logger *zap.SugaredLogger) *LogoutHandler {
	return &LogoutHandler{
		lg: logger,
	}
}

func (h LogoutHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	name := "logout"
	c := templates.Home()

	switch r.Method {

	case http.MethodGet:

		// should check headers to see if user has an access token and redirect to either home or login

		pageRender(name, c, h.lg, w, r)

	case http.MethodPost:
		// Should check cookies for jwt

	default:
		fmt.Fprintf(w, "only get and post methods are supported")
		return

	}
}
