package handlers

import (
	"encoding/json"
	"net/http"

	"backend/config"
	d "backend/internal/domain"

	"go.uber.org/zap"
)

// Refresh token Handler

type RefreshTokenHandler struct {
	lg  *zap.SugaredLogger
	cfg *config.AppConfig
}

func NewRefreshTokenHandler(logger *zap.SugaredLogger, config *config.AppConfig) *RefreshTokenHandler {
	return &RefreshTokenHandler{
		lg:  logger,
		cfg: config,
	}
}

func (h *RefreshTokenHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	err := r.ParseForm()
	if err != nil {
		h.lg.Error(err)
	}

	token := r.FormValue("token")

	RefreshAccessToken(token, w, r, h.cfg.JwtSecretKey, h.cfg.JwtRefreshSecretKey)
}

// Full login JWT Handler

type JwtHandler struct {
	lg  *zap.SugaredLogger
	cfg *config.AppConfig
}

func NewJwtHandler(logger *zap.SugaredLogger, config *config.AppConfig) *JwtHandler {
	return &JwtHandler{
		lg:  logger,
		cfg: config,
	}
}

func (h *JwtHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	switch r.Method {

	case http.MethodPost:

		var user_candidate d.User
		err := json.NewDecoder(r.Body).Decode(&user_candidate)

		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}

		h.lg.Info("User candidate: ", user_candidate)

		// Check if user credentials are valid

		if user_candidate.Username == "" {
			http.Error(w, "Username required", http.StatusBadRequest)
			return
		}
		if user_candidate.Password == "" {
			http.Error(w, "Password required", http.StatusBadRequest)
			return
		}

		for _, u := range users {
			if u.Username == user_candidate.Username && u.Password == user_candidate.Password {

				h.lg.Infof("%s logged in!", u.Username)
				var tokens d.AuthTokens
				tokens.AccessToken, tokens.RefreshToken, err = GenerateTokens(
					user_candidate.Username, h.cfg.JwtSecretKey, h.cfg.JwtRefreshSecretKey) // Generate JWT tokens in a cookie for the user
				if err != nil {
					http.Error(w, "Could not generate tokens for the user",
						http.StatusInternalServerError)
					return
				}

				tokensJson, err := json.Marshal(tokens)
				if err != nil {
					http.Error(w, "Could not marshal json response",
						http.StatusInternalServerError)
					return
				}

				h.lg.Debug("Tokens: ", tokens)
				h.lg.Debug("Tokens Json: ", tokensJson)

				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusOK)
				w.Write(tokensJson)
			}
		}
	default:
		h.lg.Error("only post method is allowed")
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

var users = []d.User{
	{
		Username: "bob",
		Password: "root",
	},
	{
		Username: "alice",
		Password: "root",
	},
}
