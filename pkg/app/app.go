package app

import (
	"chart/internal/config"
	"chart/internal/repository"
	"chart/internal/service"
	"chart/internal/transport/http/handler"
	"chart/internal/transport/router"
	"chart/internal/transport/ws"
	"chart/storage"
	"context"
	"log"
)

// Application Config is the top-level configuration object.
type Application struct {
	configuration *config.Configuration
	service       service.Service
	wsHub         *ws.Hub
}

func Create() (*Application, error) {
	cfg := config.Set()

	db, err := storage.New(cfg.ConfigDB)
	if err != nil {
		log.Fatalf("Couldn't connect to the db: %s\n", err)
	}

	repo := repository.New(db.GetDB())
	serv := service.New(repo)

	wsHub := ws.NewHub()

	return &Application{
		configuration: cfg,
		service:       serv,
		wsHub:         wsHub,
	}, nil
}

func (a *Application) Run(ctx context.Context) error {
	router.New(&a.configuration.HTTPServer, handler.New(a.service), ws.NewHandler(a.wsHub)).RunServer(ctx)
	a.wsHub.Run()
	return nil
}
