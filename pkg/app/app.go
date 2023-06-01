package app

import (
	"chart/internal/config"
	"chart/internal/repository"
	"chart/storage"
	"context"
	"log"
)

// Application Config is the top-level configuration object.
type Application struct {
	configuration *config.Configuration
}

func Create() (*Application, error) {
	cfg := config.Set()

	db, err := storage.New(cfg.ConfigDB)
	if err != nil {
		log.Fatalf("Couldn't connect to the db: %s\n", err)
	}

	repo := repository.New(db.GetDB())

	return &Application{
		configuration: cfg,
	}, nil
}

func (a *Application) Run(ctx context.Context) error {
	return nil
}
