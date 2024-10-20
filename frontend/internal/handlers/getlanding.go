package handlers

import (
	"fmt"
	"net/http"

	"frontend/internal/templates"

	"go.uber.org/zap"
)

type LoginHandler struct {
	lg *zap.SugaredLogger
}

func NewLoginHandler(logger *zap.SugaredLogger) *LoginHandler {
	return &LoginHandler{
		lg: logger,
	}
}

func (h LoginHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	name := "login"
	c := templates.LogIn()

	switch r.Method {

	case http.MethodGet:

		// should check headers to see if user has an access token and redirect

		pageRender(name, c, h.lg, w, r)

	case http.MethodPost:
		// Should check cookies for jwt

	default:
		fmt.Fprintf(w, "only get and post methods are supported")
		return

	}
}
