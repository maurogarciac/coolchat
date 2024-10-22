package domain

type User struct {
	Username string
	Password string
}

type LoginResult struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type MessageHistoryResult struct {
	Messages []byte
}

type RefreshToken struct {
	Token string `json:"refresh_token"`
}

type RefreshResult struct {
	AccessToken string `json:"access_token"`
}
