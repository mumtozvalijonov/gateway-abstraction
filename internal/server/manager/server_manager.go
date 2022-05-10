package service

import "example/microabstraction/internal/server"

type ServerManager struct {
	server server.Server
}

func (s *ServerManager) StartServer() {
	s.server.Run()
}

func NewServerManager(s server.Server) (*ServerManager, error) {
	return &ServerManager{server: s}, nil
}
