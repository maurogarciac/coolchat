package domain

type User struct {
	Username string
	Password string
}

type LoginResult struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	User         string `json:"user"`
}

type Message struct {
	Text      string `json:"text"`
	User      string `json:"user"`
	Timestamp string `json:"ts"`
}

type MessageHistoryResult struct {
	Messages []Message
}

type RefreshToken struct {
	Token string `json:"refresh_token"`
}

type RefreshResult struct {
	AccessToken string `json:"access_token"`
}
