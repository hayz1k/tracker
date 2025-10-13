package configs

import (
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
	"github.com/rs/zerolog/log"
)

type WooCommerceConfig struct {
	Url    string `envconfig:"BASE_URL"`
	Key    string `envconfig:"CONSUMER_KEY"`
	Secret string `envconfig:"CONSUMER_SECRET"`
}

type Config struct {
	WooCommerce WooCommerceConfig
}

func LoadConfig() (*Config, error) {
	var cfg Config

	if err := godotenv.Load(); err != nil {
		log.Fatal().Err(err).Msg("Error loading .env file")
		return nil, err
	}

	if err := envconfig.Process("", &cfg); err != nil {
		log.Fatal().Err(err).Msg("Error reading .env file into config")
		return nil, err
	}

	return &cfg, nil
}
