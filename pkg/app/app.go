package app

import (
	"chart/internal/config"
	"chart/internal/repository"
	"chart/internal/service"
	"chart/internal/transport/http/handler"
	"chart/internal/transport/router"
	"chart/storage"
	"context"
	"log"
)

// Application Config is the top-level configuration object.
type Application struct {
	configuration *config.Configuration
	service       service.Service
}

func Create() (*Application, error) {
	cfg := config.Set()

	db, err := storage.New(cfg.ConfigDB)
	if err != nil {
		log.Fatalf("Couldn't connect to the db: %s\n", err)
	}

	repo := repository.New(db.GetDB())
	serv := service.New(repo)

	return &Application{
		configuration: cfg,
		service:       serv,
	}, nil
}

func (a *Application) Run(ctx context.Context) error {
	router.New(&a.configuration.HTTPServer, handler.New(a.service)).RunServer(ctx)
	return nil
}
