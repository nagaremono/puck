package httpserver

import (
	"context"
	"log"
	"net/http"
	"time"
)

const (
	_defaultAddress = ":8000"
)

type Server struct {
	server *http.Server
	notify chan error
}

func New(handler http.Handler) *Server {
	s := http.Server{
		Handler: handler,
		Addr:    _defaultAddress,
	}

	server := Server{
		server: &s,
		notify: make(chan error, 1),
	}

	return &server
}

func (s *Server) Start() {
	go func() {
		s.notify <- s.server.ListenAndServe()
		close(s.notify)
	}()
}

func (s *Server) Notify() <-chan error {
	return s.notify
}

func (s *Server) Shutdown() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := s.server.Shutdown(ctx); err != nil {
		log.Fatal("Server forced shutdown: ", err)
	}
}
