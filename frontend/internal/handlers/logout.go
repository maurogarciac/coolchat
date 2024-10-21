package handlers

import (
	"fmt"
	"net/http"
	"time"

	"frontend/internal/middleware"
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

	c := templates.LogOut()

	switch r.Method {

	case http.MethodGet:

		pageRender("logout", c, true, h.lg, w, r)

	case http.MethodPost:

		deleteTokenCookie(middleware.AccessTokenCookieName, w)
		deleteTokenCookie(middleware.RefreshTokenCookieName, w)

		http.Redirect(w, r, "/login", http.StatusPermanentRedirect)

	default:
		fmt.Fprintf(w, "only get and post methods are supported")
		return

	}
}

func deleteTokenCookie(name string, w http.ResponseWriter) {

	c := &http.Cookie{
		Name:     name,
		Value:    "",
		Path:     "/",
		Expires:  time.Unix(0, 0),
		HttpOnly: true,
		SameSite: 0,
	}

	http.SetCookie(w, c)
}
