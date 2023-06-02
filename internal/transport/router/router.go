package router

import (
	"chart/internal/config"
	"chart/internal/transport/handler"
	"context"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
)

type Router interface {
	RunServer(ctx context.Context)
}

type router struct {
	serverPort *config.HTTPServerConfiguration
	handler    handler.Handler
}

func New(s *config.HTTPServerConfiguration, h handler.Handler) Router {
	return &router{
		serverPort: s,
		handler:    h,
	}
}

// RunServer starts HTTP server
func (r *router) RunServer(ctx context.Context) {
	engine := gin.Default()
	r.assignRoutes(engine)

	srv := &http.Server{
		Addr:    r.serverPort.Port,
		Handler: engine,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	// Listen for the interrupt signal.
	<-ctx.Done()

	log.Println("shutting down gracefully, press Ctrl+C again to force")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown: ", err)
	}

	log.Println("Server exiting")
}

// AssignRoutes assign the available routes of the HTTP API
func (r *router) assignRoutes(engine *gin.Engine) {
	engine.GET("/create", r.handler.CreateUser)
	engine.GET("/login", r.handler.LoginUser)
}
