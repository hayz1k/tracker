package configs

import (
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
	"github.com/rs/zerolog/log"
)

type Config struct {
	DSN     string `envconfig:"DSN"`
	Address string `envconfig:"ADDRESS"`
}

func LoadConfig() (*Config, error) {
	var cfg Config

	if err := godotenv.Load("backend/.env"); err != nil {
		log.Fatal().Err(err).Msg("Error loading .env file")
		return nil, err
	}

	if err := envconfig.Process("", &cfg); err != nil {
		log.Fatal().Err(err).Msg("Error reading .env file into config")
		return nil, err
	}

	return &cfg, nil
}
