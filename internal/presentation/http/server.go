package http

import (
	"context"
	"fmt"
	"github.com/DrozdovRoman/avito-tech-banner-service/internal/application/common/configuration"
	"github.com/sirupsen/logrus"
	"net"
	"net/http"
)

type Server struct {
	*http.Server
}

func NewHttpServer(config *configuration.Configuration, r *Router) *Server {
	srv := &http.Server{Addr: fmt.Sprintf("%s:%d", config.HTTP.IP, config.HTTP.Port), Handler: r}
	return &Server{srv}
}

func (srv *Server) Stop(ctx context.Context) error {
	err := srv.Shutdown(ctx)
	if err != nil {
		return err
	}
	logrus.Println("HTTP stopped!")
	return nil
}

func (srv *Server) Start(ctx context.Context) error {
	ln, err := net.Listen("tcp", srv.Addr)
	if err != nil {
		return err
	}
	go srv.Serve(ln)
	logrus.Printf("HTTP started! Server listening on %v", srv.Addr)
	return nil
}
