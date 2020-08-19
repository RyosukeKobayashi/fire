package config

import "os"

type PostgresConfig struct {
	URI      string `env:"POSTGRES_URI"`
	USER     string `env:"POSTGRES_USER"`
	PASSWORD string `env:"POSTGRES_PASSWORD"`
	NAME     string `env:"POSTGRES_DB"`
}

func GetPostgresConfig() PostgresConfig {
	return PostgresConfig{
		URI:      os.Getenv("POSTGRES_URI"),
		USER:     os.Getenv("POSTGRES_USER"),
		PASSWORD: os.Getenv("POSTGRES_PASSWORD"),
		NAME:     os.Getenv("POSTGRES_DB"),
	}
}
