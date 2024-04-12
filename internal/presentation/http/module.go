package http

import (
	"context"
	"github.com/DrozdovRoman/avito-tech-banner-service/internal/presentation/http/api"
	client "github.com/DrozdovRoman/avito-tech-banner-service/internal/presentation/http/api/client/v1"
	"github.com/sirupsen/logrus"
	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(
		NewHttpServer,
		NewRouter,
		api.NewLoginHandler,
		client.NewBannerHandler,
	),

	fx.Invoke(
		// http
		func(lc fx.Lifecycle, httpServer *Server) {
			logrus.Println("Starting HTTP server")
			lc.Append(fx.Hook{
				OnStart: func(ctx context.Context) error {
					return httpServer.Start(ctx)
				},
				OnStop: func(ctx context.Context) error {
					return httpServer.Stop(ctx)
				},
			})
		},
	),
)
