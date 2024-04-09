package http

import (
	"github.com/DrozdovRoman/avito-tech-banner-service/internal/presentation/http/api"
	client "github.com/DrozdovRoman/avito-tech-banner-service/internal/presentation/http/api/client/v1"
	"github.com/go-chi/chi/v5"
)

type Router struct {
	chi.Router
}

func NewRouter(
	bannerHandler *client.BannerHandler,
) *Router {
	r := chi.NewRouter()

	r.Get("/ping", api.PingHandler)

	r.Route("/banner", func(r chi.Router) {
		r.Get("/{id}", bannerHandler.GetBannerByID)
	})

	return &Router{r}
}
