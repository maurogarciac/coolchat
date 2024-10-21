package handlers

import (
	"fmt"
	"net/http"

	"frontend/internal/domain"
	auth "frontend/internal/middleware"
	"frontend/internal/services"
	"frontend/internal/templates"

	"go.uber.org/zap"
)

type LoginHandler struct {
	lg *zap.SugaredLogger
	b  *services.BackendService
}

func NewLoginHandler(logger *zap.SugaredLogger, backend *services.BackendService) *LoginHandler {
	return &LoginHandler{
		lg: logger,
		b:  backend,
	}
}

func (h LoginHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	c := templates.LogIn()

	switch r.Method {

	case http.MethodGet:

		accessToken := ""
		for _, cookie := range r.Cookies() {
			if cookie.Name == auth.AccessTokenCookieName {
				accessToken = cookie.Value
			}
		}

		_, err := auth.VerifyAccessToken(accessToken)
		if err != nil {
			pageRender("login", c, false, h.lg, w, r)
			return
		}
		if err == nil {
			http.Redirect(w, r, "/home/", http.StatusMovedPermanently)
		}

	case http.MethodPost:

		var users = []domain.User{
			{
				Username: "bob",
				Password: "root",
			},
			{
				Username: "alice",
				Password: "root",
			},
		}

		// Should check cookies for jwt?
		err := r.ParseForm()
		if err != nil {
			h.lg.Error(err)
		}

		login_creds := domain.User{
			Username: r.FormValue("username"),
			Password: r.FormValue("password"),
		}

		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}

		h.lg.Info("req: %s", login_creds)

		// check if login_creds are valid

		if login_creds.Username == "" {
			http.Error(w, "Username required", http.StatusBadRequest)
			return
		}
		if login_creds.Password == "" {
			http.Error(w, "Password required", http.StatusBadRequest)
			return
		}

		accepted := false
		for _, u := range users {
			if u.Username == login_creds.Username && u.Password == login_creds.Password {
				err := auth.GenerateTokens(u.Username, w) // Generate JWT tokens in a cookie for the user
				if err != nil {
					h.lg.Error("Token generation error: ", err)
					http.Error(w, "Could not generate tokens for the user", http.StatusInternalServerError)
					return
				}
				accepted = true
			}
		}
		if !accepted {
			// render same template but login = incorrect or something
			http.Redirect(w, r, "/login", http.StatusBadRequest)
		}

		http.Redirect(w, r, "/home", http.StatusMovedPermanently)

	default:
		fmt.Fprintf(w, "only get and post methods are supported")
		return

	}
}
