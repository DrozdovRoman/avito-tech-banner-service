package http

import (
	"github.com/DrozdovRoman/avito-tech-banner-service/internal/application/common/configuration"
	"github.com/DrozdovRoman/avito-tech-banner-service/internal/presentation/http/api"
	client "github.com/DrozdovRoman/avito-tech-banner-service/internal/presentation/http/api/client/v1"
	"github.com/go-chi/chi/v5"
	httpSwagger "github.com/swaggo/http-swagger"
)

type Router struct {
	chi.Router
}

func NewRouter(
	config *configuration.Configuration,
	bannerHandler *client.BannerHandler,
) *Router {
	r := chi.NewRouter()

	r.Get("/ping", api.PingHandler)

	r.Route("/banner", func(r chi.Router) {
		r.Get("/{id}", bannerHandler.GetBannerByID)
	})

	r.Get("/swagger/*", httpSwagger.Handler(
		httpSwagger.URL(config.ApiExternalURL+"/swagger/doc.json"),
	))

	return &Router{r}
}
