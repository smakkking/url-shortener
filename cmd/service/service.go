package main

import (
	"os"

	"github.com/sirupsen/logrus"
	"github.com/smakkking/url-shortener/internal/app"
	"github.com/smakkking/url-shortener/internal/controllers/httphandlers"
	"github.com/smakkking/url-shortener/internal/httpserver"
	"github.com/smakkking/url-shortener/internal/infrastructure/inmemory"
	"github.com/smakkking/url-shortener/internal/services"
)

const (
	configPath = "./config/config.yaml"
)

func main() {
	// init логгер
	logrus.SetFormatter(
		&logrus.JSONFormatter{
			PrettyPrint:     true,
			TimestampFormat: "2006-01-02 15:04:05",
		},
	)
	logrus.SetOutput(os.Stdout)

	logrus.Infoln("service started...")
	logrus.Debugln("debug messages are available")

	// загрузка конфига
	config, err := app.NewConfig(configPath)
	if err != nil {
		logrus.Errorf("error reading config: %s", err.Error())
		os.Exit(1)
	}

	// init репозитории
	inmemoryStorage := inmemory.NewStorage()

	// init сервисы
	urlService := services.NewService(inmemoryStorage)

	// init хендлеры
	urlHandler := httphandlers.NewHandler(urlService)

	// запуск сервера
	srv := httpserver.NewServer(config)
	srv.SetupHandlers(urlHandler)
	srv.Run()
	// graceful shutdown
}
