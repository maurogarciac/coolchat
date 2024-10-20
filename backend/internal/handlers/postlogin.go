package handlers

import (
	"backend/config"
	"backend/internal/domain"
	"backend/internal/middleware"

	"context"
	"encoding/json"
	"io"
	"log"
	"net/http"

	"go.uber.org/zap"
)

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

type LoginHandler struct {
	context   context.Context
	appConfig config.AppConfig
	lg        *zap.SugaredLogger
}

func NewLoginHandler(context context.Context, config config.AppConfig, logger *zap.SugaredLogger) *LoginHandler {
	return &LoginHandler{
		context:   context,
		appConfig: config,
		lg:        logger,
	}
}

func (h *LoginHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}

	var login_creds domain.User

	err = json.Unmarshal(body, &login_creds)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	h.lg.Info("req: %s", login_creds)

	// check if login_creds are valid

	if login_creds.Username == "" {
		http.Error(w, "Username required", http.StatusBadRequest)
		return
	}

	for _, u := range users {
		if u.Username == login_creds.Username {
			if u.Password == login_creds.Password {
				err := middleware.GenerateTokens(u.Username, w) // Generate JWT tokens in a cookie for the user
				if err != nil {
					http.Error(w, "Could not generate tokens for the user", http.StatusInternalServerError)
					return
				}
			} else {
				http.Error(w, "Incorrect password", http.StatusBadRequest)
			}
		} else {
			http.Error(w, "User does not exist", http.StatusBadRequest)
		}
	}

	http.Redirect(w, r, "/home/", http.StatusFound)
}
