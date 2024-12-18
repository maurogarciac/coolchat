package config

import "time"

type AppConfig struct {
	Ip           string        `env:"SERVER_IP"             envDefault:"localhost"`
	ServerPort   int           `env:"FRONTEND_PORT"         envDefault:"8000"`
	ReadTimeout  time.Duration `env:"FRONTEND_READ_TIMEOUT" envDefault:"5s"`
	LogLevel     string        `env:"FRONTEND_LOG_LEVEL"           envDefault:"debug"`
	JwtSecretKey string        `env:"JWT_SECRET"          envDefault:""`
}

type HTTPConfig struct {
	Timeout       time.Duration `env:"HTTP_CLIENT_TIMEOUT"   envDefault:"1s"`
	RetryMax      int           `env:"HTTP_CLIENT_RETRY_MAX" envDefault:"1"`
	BackendAPIURL string        `env:"BACKEND_API_URL"       envDefault:"http://localhost:1337"`
}

type HTMLConfig struct {
	Css string `env:"CSS" envDefault:"static/styles/index.css"`
}
