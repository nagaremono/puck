package restapi

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/nagaremono/puck/internal/controller"
	"github.com/nagaremono/puck/pkg/httpserver"
)

func Run() {
	handler := controller.NewRouter()
	s := httpserver.New(handler)

	s.Start()

	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	select {
	case <-quit:
		log.Println("Shutting down server")
	case err := <-s.Notify():
		log.Fatalf("listen: %s\n", err)
	}

	s.Shutdown()
}
