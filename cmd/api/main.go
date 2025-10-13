package main

import (
	"context"
	"github.com/rs/zerolog"
	_ "net/http"
	"orderTracker/configs"
	"orderTracker/internal/repositories"
	"os"
)

func main() {
	ctx := context.Background()
	log := zerolog.New(os.Stdout).With().Timestamp().Logger()

	cfg, err := configs.LoadConfig()
	if err != nil {
		log.Fatal().Msg("failed to load config")
	}

	if cfg.WooCommerce.Key == "" || cfg.WooCommerce.Secret == "" {
		log.Fatal().Msg("cfg is empty")
	}

	store.ApiRequest(ctx, cfg)

}
