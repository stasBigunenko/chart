package router

import (
	"chart/internal/config"
	"chart/internal/transport/http/handler"
	"chart/internal/transport/ws"
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
	wsHandler  ws.Handler
}

func New(s *config.HTTPServerConfiguration, h handler.Handler, ws ws.Handler) Router {
	return &router{
		serverPort: s,
		handler:    h,
		wsHandler:  ws,
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
	engine.Use(r.handler.VerifyUser())

	engine.POST("/signup", r.handler.CreateUser)
	engine.POST("/signin", r.handler.LoginUser)
	engine.GET("/welcome", r.handler.Welcome)
	engine.GET("/logout", r.handler.Logout)

	engine.POST("/ws/createRoom", r.wsHandler.CreateRoom)
	engine.GET("/ws/joinRoom/:roomId", r.wsHandler.JoinRoom)
	engine.GET("/ws/getRooms", r.wsHandler.GetRooms)
	engine.GET("/ws/getClients/:roomId", r.wsHandler.GetClients)
}
