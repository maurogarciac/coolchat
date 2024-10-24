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
	ID        string    `json:"id"`
	Text      string    `json:"text"`
	From      string    `json:"user"`
	Timestamp time.Time `json:"ts"`
}

type InsertMessage struct {
	Text string
	From string
}

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type AuthTokens struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type RefreshToken struct {
	RefreshToken string `json:"refresh_token"`
}

type AccessToken struct {
	AccessToken string `json:"access_token"`
}
