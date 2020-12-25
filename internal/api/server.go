package api

import (
	"context"
	"fmt"
	"net/http"

	"github.com/objque/gohan/internal/config"
)

type Server struct {
	server *http.Server
	Addr   string
}

func New(router http.Handler, conf config.HTTPConfig) *Server {
	server := http.Server{
		Addr:         fmt.Sprintf("%s:%d", conf.IP, conf.Port),
		Handler:      router,
		ReadTimeout:  conf.ReadTimeout,
		WriteTimeout: conf.WriteTimeout,
		IdleTimeout:  conf.IdleTimeout,
	}

	return &Server{server: &server, Addr: server.Addr}
}

func (s *Server) ListenAndServe() error {
	return s.server.ListenAndServe()
}

func (s *Server) SetKeepAlivesEnabled(v bool) {
	s.server.SetKeepAlivesEnabled(v)
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.server.Shutdown(ctx)
}
