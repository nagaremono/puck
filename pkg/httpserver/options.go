package httpserver

import "net"

type ServerOption func(*Server)

func Port(port string) ServerOption {
	return func(s *Server) {
		s.server.Addr = net.JoinHostPort("", port)
	}
}
