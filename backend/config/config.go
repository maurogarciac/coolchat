package config

import (
	"time"
)

type AppConfig struct {
	DbHost              string        `env:"DB_HOST"             envDefault:"localhost"`
	DbPort              string        `env:"DB_PORT"             envDefault:"5432"`
	DbUser              string        `env:"DB_USER"             envDefault:""`
	DbPassword          string        `env:"DB_PASS"             envDefault:""`
	DbName              string        `env:"DB_NAME"             envDefault:"postgres"`
	JwtSecretKey        string        `env:"JWT_SECRET"          envDefault:""`
	JwtRefreshSecretKey string        `env:"RSH_SECRET"          envDefault:""`
	ServerPort          int           `env:"SERVER_PORT"         envDefault:"1337"`
	ReadTimeout         time.Duration `env:"SERVER_READ_TIMEOUT" envDefault:"5s"`
	LogLevel            string        `env:"LOG_LEVEL"           envDefault:"info"`
}
