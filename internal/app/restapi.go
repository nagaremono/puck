package restapi

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/nagaremono/puck/internal/config"
	"github.com/nagaremono/puck/internal/controller"
	"github.com/nagaremono/puck/pkg/httpserver"
	"github.com/nagaremono/puck/pkg/logger"
	"github.com/rs/zerolog"
)

func Run() {
	cfg := config.NewConfig()
	logger := logger.New(zerolog.Level(cfg.Log.Level()))
	handler := controller.NewRouter()
	s := httpserver.New(handler, httpserver.Port(cfg.Server.Port()))

	s.Start()

	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	select {
	case <-quit:
		logger.Info().Msg("Quit signal")
	case err := <-s.Notify():
		logger.Error().Err(err).Msg("listen: %s\n")
	}

	s.Shutdown()
}
