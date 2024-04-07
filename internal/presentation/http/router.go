package http

import (
	"github.com/DrozdovRoman/avito-tech-banner-service/internal/application/common/configuration"
	"github.com/DrozdovRoman/avito-tech-banner-service/internal/presentation/http/api"
	"github.com/go-chi/chi/v5"
)

type Router struct {
	chi.Router
}

func NewRouter(config *configuration.Configuration) *Router {
	r := chi.NewRouter()

	r.Get("/ping", api.PingHandler)

	return &Router{r}
}
