package api

import (
	"frontend/internal/domain"
)

// Login

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func toBackendApiLoginRequest(input domain.User) LoginRequest {
	return LoginRequest{
		Username: input.Username,
		Password: input.Password,
	}
}

type PostLoginResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

func (r *PostLoginResponse) PostLoginResult() domain.LoginResult {

	return domain.LoginResult{
		AccessToken:  r.AccessToken,
		RefreshToken: r.RefreshToken,
	}
}

// Refresh Access Token

type RefreshRequest struct {
	RefreshToken string `json:"refresh_token"`
}

func toBackendApiRefreshRequest(input domain.RefreshToken) RefreshRequest {
	return RefreshRequest{
		RefreshToken: input.Token,
	}
}

type PostRefreshResponse struct {
	AccessToken string `json:"access_token"`
}

func (r *PostRefreshResponse) PostRefreshResult() domain.RefreshResult {

	return domain.RefreshResult{
		AccessToken: r.AccessToken,
	}
}

// Message history

type MessageHistoryResponse struct {
	Messages []domain.Message
}

func (r *MessageHistoryResponse) MessageHistoryResult() domain.MessageHistoryResult {

	return domain.MessageHistoryResult{
		Messages: r.Messages,
	}
}
