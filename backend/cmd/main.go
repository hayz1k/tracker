package main

import (
	"context"
	"github.com/rs/zerolog"
	_ "net/http"
	"orderTracker/configs"
	"orderTracker/internal/domain/server"
	"orderTracker/internal/observability"
	"orderTracker/internal/store/postgres"
	"os"
)

func main() {
	observability.Init()

	ctx := context.Background()
	_ = ctx
	log := zerolog.New(os.Stdout).With().Timestamp().Logger()

	cfg, err := configs.LoadConfig()
	if err != nil {
		log.Fatal().Msg("failed to load config")
	}

	store, err := postgres.NewPostgresStore(cfg)
	if err != nil {
		log.Fatal().Msg("error connecting database")
	}
	defer store.Close()

	app := server.NewApp(cfg, store)

	srv := server.NewServer(app)
	
	log.Info().Msgf("server is running on port :%s", cfg.Address)
	if err = srv.Run(cfg.Address); err != nil {
		log.Fatal()
	}
}
