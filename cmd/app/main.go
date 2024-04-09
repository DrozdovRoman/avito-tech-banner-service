package main

import (
	"github.com/DrozdovRoman/avito-tech-banner-service/internal/application"
	"github.com/DrozdovRoman/avito-tech-banner-service/internal/application/common/configuration"
	"github.com/DrozdovRoman/avito-tech-banner-service/internal/infrastructure"
	"github.com/DrozdovRoman/avito-tech-banner-service/internal/presentation/http"
	"github.com/sirupsen/logrus"
	fxlogrus "github.com/takt-corp/fx-logrus"
	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"
)

var configFile = ""

// @title           Banner Service
// @version         1.0
// @description     How many Avito developers does it take to change a light bulb? (Interview answer <3)

// @contact.name   Drozdov Roman
// @contact.email  romandrozdov@icloud.com
func main() {
	config, err := configuration.LoadConfiguration(configFile)
	if err != nil {
		logrus.WithError(err).Fatal("unable to load configuration")
	}

	fx.New(
		fx.WithLogger(func() fxevent.Logger {
			return &fxlogrus.LogrusLogger{Logger: logrus.StandardLogger()}
		}),
		fx.Supply(config),
		http.Module,
		infrastructure.Module,
		application.Module,
	).Run()
}
