package http

import (
	"github.com/DrozdovRoman/avito-tech-banner-service/internal/presentation/http/api"
	"github.com/go-chi/chi/v5"
)

type Router struct {
	chi.Router
}

func NewRouter() *chi.Mux {
	router := chi.NewRouter()

	router.Get("/ping", api.PingHandler)

	return router
}
