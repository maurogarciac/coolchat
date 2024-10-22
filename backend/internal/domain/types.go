package domain

import (
	"time"
)

type GetAllMessagesResult struct {
	Messages []Message `json:"messages"`
}

type MessageHistory struct {
	MessageList []Message
}

type Message struct {
	ID        string
	Text      string
	From      string
	Timestamp time.Time
}

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type AuthTokens struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}
