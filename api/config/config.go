package config

import "os"

// const (
// 	DB_USER     = "user"
// 	DB_PASSWORD = "password"
// 	DB_NAME     = "fire"
// 	CLIENT_URL  = "http://localhost:3000"
// )

const (
	prod = "production"
)

type Config struct {
	Env      string         `env:"ENV"`
	Postgres PostgresConfig `json:"postgres"`
	Host     string         `env:"APP_HOST"`
	Port     string         `env:"APP_PORT"`
}

func (c Config) IsProd() bool {
	return c.Env == prod
}

func GetConfig() Config {
	return Config{
		Env:      os.Getenv("ENV"),
		Postgres: GetPostgresConfig(),
		Host:     os.Getenv("APP_HOST"),
		Port:     os.Getenv("APP_PORT"),
	}
}
