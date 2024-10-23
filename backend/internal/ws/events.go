package ws

import (
	"time"
)

const (
	SendMessage = "send_message"
	NewMessage  = "new_message"
)

type EventHandler func(event Event, c *Client) error

type InnerMessage struct {
	Text string `json:"text"`
	User string `json:"user"`
}

type RecievedEgressMessage struct {
	Message string `json:"message"`
}

type ReturnMessage struct {
	Text      string    `json:"text"`
	User      string    `json:"user"`
	Timestamp time.Time `json:"ts"`
}

type Event struct {
	Message string `json:"message"`
}
