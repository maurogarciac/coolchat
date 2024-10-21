package domain

type User struct {
	Username string
	Password string
}

type LoginResult struct {
	Message string
}

type MessageHistoryResult struct {
	Messages []byte
}
