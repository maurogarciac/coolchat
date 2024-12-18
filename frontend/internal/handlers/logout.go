package handlers

import (
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

		h.lg.Info("logout GET")

		pageRender("logout", c, true, h.lg, w, r)

	case http.MethodPost:

		h.lg.Infof("logout POST: %s", r.Body)

		deleteTokenCookie(middleware.AccessTokenCookieName, w)
		deleteTokenCookie(middleware.RefreshTokenCookieName, w)

		// Needs to use hx-redirect header for partial page rendering
		w.Header().Set("HX-Redirect", "/login/")
		w.WriteHeader(http.StatusOK)

	default:
		h.lg.Error("Only GET and POST methods are supported")
		http.Error(w, "Only GET and POST methods are supported", http.StatusMethodNotAllowed)
	}
}

func deleteTokenCookie(name string, w http.ResponseWriter) {

	c := &http.Cookie{
		Name:     name,
		Value:    "",
		Path:     "/",
		MaxAge:   -1,
		Expires:  time.Unix(0, 0),
		HttpOnly: true,
		SameSite: 0,
	}

	http.SetCookie(w, c)
}
