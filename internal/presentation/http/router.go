package http

import (
	"github.com/DrozdovRoman/avito-tech-banner-service/internal/application/common/configuration"
	"github.com/DrozdovRoman/avito-tech-banner-service/internal/presentation/http/api"
	admin "github.com/DrozdovRoman/avito-tech-banner-service/internal/presentation/http/api/admin/v1"
	client "github.com/DrozdovRoman/avito-tech-banner-service/internal/presentation/http/api/client/v1"
	"github.com/DrozdovRoman/avito-tech-banner-service/internal/presentation/http/api/common/middlewares"
	"github.com/go-chi/chi/v5"
	httpSwagger "github.com/swaggo/http-swagger"
)

type Router struct {
	chi.Router
}

func NewRouter(
	config *configuration.Configuration,
	userBannerHandler *client.UserBannerHandler,
	adminBannerHandler *admin.AdminBannerHandler,
	loginHandler *api.LoginHandler,
	authMiddleware *middlewares.AuthMiddleware,
) *Router {
	r := chi.NewRouter()

	r.Get("/ping", api.PingHandler)
	r.Post("/login", loginHandler.Login)

	r.With(authMiddleware.Handler).Get("/user_banner", userBannerHandler.FetchActiveUserBannerContent)

	r.Route("/banner", func(r chi.Router) {
		r.Post("/", adminBannerHandler.CreateBanner)
	})

	r.Get("/swagger/*", httpSwagger.Handler(
		httpSwagger.URL(config.ApiExternalURL+"/swagger/doc.json"),
	))

	return &Router{r}
}
