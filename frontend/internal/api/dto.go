package api

import (
	"frontend/internal/domain"
)

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
	Message string `json:"message"`
}

func (r *PostLoginResponse) PostLoginResult() domain.LoginResult {

	return domain.LoginResult{
		Message: r.Message,
	}
}

type MessageHistoryResponse struct {
	Messages []byte `json:"messages"`
}

func (r *MessageHistoryResponse) MessageHistoryResult() domain.MessageHistoryResult {
	return domain.MessageHistoryResult{
		Messages: r.Messages,
	}
}
