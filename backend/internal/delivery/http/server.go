package http

import (
	"context"
	"net/http"
	"time"
)

// Структура сервера
type Server struct {
	httpServer *http.Server
}

// Создание нового сервера
func NewServer(addr string, handler http.Handler) *Server {
	return &Server{
		httpServer: &http.Server{
			Addr:         addr,
			Handler:      handler,
			ReadTimeout:  10 * time.Second,
			WriteTimeout: 10 * time.Second,
		},
	}
}

// Старт сервера
func (s *Server) Start() error {
	return s.httpServer.ListenAndServe()
}

// Отключение сервера
func (s *Server) Shutdown(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}
