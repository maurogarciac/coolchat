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

	c := templates.Home() // make logout template

	switch r.Method {

	case http.MethodGet:

		// are you sure you want to logout??? with 2 buttons
		// if yes, redirect to login
		// if no just uhh redirect to home page

		pageRender("logout", c, h.lg, w, r)

	default:
		fmt.Fprintf(w, "only get and post methods are supported")
		return

	}
}
