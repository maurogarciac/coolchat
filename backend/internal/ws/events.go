package ws

import (
	"time"
)

const (
	SendMessage = "send_message"
	NewMessage  = "new_message"
)

type EventHandler func(event Event, c *Client) error

type Event struct {
	Message string `json:"message"`
}

type EgressMessageEvent struct {
	Text string    `json:"message"`
	User string    `json:"user"`
	Sent time.Time `json:"sent"`
}
