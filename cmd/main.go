package main

import (
	"context"
	"github.com/rs/zerolog"
	_ "net/http"
	"orderTracker/configs"
	start "orderTracker/internal"
	"orderTracker/internal/store/postgres"
	"os"
)

func main() {
	ctx := context.Background()
	_ = ctx
	log := zerolog.New(os.Stdout).With().Timestamp().Logger()

	cfg, err := configs.LoadConfig()
	if err != nil {
		log.Fatal().Msg("failed to load config")
	}

	if cfg.WooCommerce.Key == "" || cfg.WooCommerce.Secret == "" {
		log.Fatal().Msg("cfg is empty")
	}

	store, err := postgres.NewPostgresStore(cfg)
	if err != nil {
		log.Info().Msgf("%s", err)
		log.Fatal().Msg("error connecting database")
	}
	defer store.Close()

	app := start.NewApp(cfg, store)

	server := start.NewServer(app)

	if err = server.Run(cfg.Address); err != nil {
		log.Fatal()
	}
}
