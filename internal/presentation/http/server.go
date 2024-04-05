package http

import (
	"context"
	"errors"
	"fmt"
	"github.com/DrozdovRoman/avito-tech-banner-service/internal/application/common/configuration"
	"github.com/sirupsen/logrus"
	"net/http"
)

type Server struct {
	server *http.Server
}

func NewHttpServer(config *configuration.Configuration) *Server {
	router := NewRouter()

	server := &http.Server{
		Addr:    fmt.Sprintf("%s:%d", config.HTTP.IP, config.HTTP.Port),
		Handler: router,
	}

	return &Server{server: server}
}

func (s *Server) Stop(ctx context.Context) error {
	logrus.Infof("Shutting down server...")
	return s.server.Shutdown(ctx)
}

func (s *Server) Start() {
	go func() {
		logrus.Infof("Starting server on %s", s.server.Addr)
		if err := s.server.ListenAndServe(); !errors.Is(err, http.ErrServerClosed) {
			logrus.Fatalf("Failed to listen and serve: %+v", err)
		}
	}()
}
