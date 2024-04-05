package main

import (
	"context"
	"github.com/DrozdovRoman/avito-tech-banner-service/internal/application/common/configuration"
	"github.com/DrozdovRoman/avito-tech-banner-service/internal/presentation/http"
	"github.com/sirupsen/logrus"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var configFile = ""

// @title           Banner Service
// @version         1.0
// @description     How many Scots does it take to change a light bulb?

// @contact.name   Drozdov Roman
// @contact.email  romandrozdov@icloud.com
func main() {
	logrus.SetLevel(logrus.InfoLevel)
	logrus.SetFormatter(&logrus.TextFormatter{})

	config, err := configuration.LoadConfiguration(configFile)
	if err != nil {
		logrus.WithError(err).Fatal("Unable to load configuration")
	}
	logrus.Info("Configuration loaded successfully")

	server := http.NewHttpServer(config)
	server.Start()

	stopChan := make(chan os.Signal, 1)
	signal.Notify(stopChan, os.Interrupt, syscall.SIGTERM)
	<-stopChan

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := server.Stop(ctx); err != nil {
		logrus.WithError(err).Fatal("Failed to gracefully stop the server")
	} else {
		logrus.Info("Server stopped gracefully")
	}

}
