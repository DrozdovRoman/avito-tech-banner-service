package application

import (
	"github.com/DrozdovRoman/avito-tech-banner-service/internal/application/service"
	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(
		service.NewBannerService,
	),
)
