package handlers

import (
	"encoding/json"
	"net/http"
	"strings"

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

	switch r.Method {

	case http.MethodPost:

		var token d.RefreshToken
		err := json.NewDecoder(r.Body).Decode(&token)

		if err != nil {
			h.lg.Error(err)
			http.Error(w, "Missing refresh_token", http.StatusBadRequest)
		}

		if token.RefreshToken == "" {
			http.Error(w, "Value for refresh_token is empty", http.StatusBadRequest)
		} else {
			h.lg.Debugf("Refresh token recieved: %s", token)

			access_token, err := RefreshAccessToken(token.RefreshToken, w, r, h.cfg.JwtSecretKey, h.cfg.JwtRefreshSecretKey)

			if err != nil {
				h.lg.Errorf("Error refreshing token: %s", err)
				http.Error(w, "Could not refresh access_token", http.StatusUnauthorized)
			}

			if access_token != "" {

				tokenJson, err := json.Marshal(d.AccessToken{AccessToken: access_token})
				if err != nil {
					h.lg.Error(err)
				}
				h.lg.Debugf("Access token json returned: %s", tokenJson)

				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusOK)
				w.Write(tokenJson)
			}
		}

	default:
		h.lg.Error("only post method is allowed")
		http.Error(w, "Only POST method allowed", http.StatusMethodNotAllowed)
	}
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

		user_candidate.Username = strings.ToLower(user_candidate.Username)
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

		accepted := false

		for _, u := range users {
			if u.Username == user_candidate.Username && u.Password == user_candidate.Password {

				accepted = true
				h.lg.Infof("%s logged in!", u.Username)

				var res d.LogInResponse
				res.AccessToken, res.RefreshToken, err = GenerateTokens(
					user_candidate.Username, h.cfg.JwtSecretKey, h.cfg.JwtRefreshSecretKey) // Generate JWT tokens in a cookie for the user
				if err != nil {
					http.Error(w, "Could not generate tokens for the user",
						http.StatusInternalServerError)
					return
				}
				res.Username = u.Username

				resJson, err := json.Marshal(res)
				if err != nil {
					http.Error(w, "Could not marshal json response",
						http.StatusInternalServerError)
					return
				}

				h.lg.Debug("Response: ", res)
				h.lg.Debug("Tokens Json: ", resJson)

				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusOK)
				w.Write(resJson)
			}
		}
		if !accepted {
			h.lg.Error("user does not exist")
			http.Error(w, "User does not exist", http.StatusForbidden)
		}

	default:
		h.lg.Error("only post method is allowed")
		http.Error(w, "Only POST method allowed", http.StatusMethodNotAllowed)
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
	{
		Username: "test_bob",
		Password: "test_pw",
	},
	{
		Username: "test_alice",
		Password: "test_pw",
	},
}
