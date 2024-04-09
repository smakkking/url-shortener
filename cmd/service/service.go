package main

import (
	"os"

	"github.com/sirupsen/logrus"
	"github.com/smakkking/url-shortener/internal/app"
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

	// загрузка конфига
	config, err := app.NewConfig(configPath)
	if err != nil {
		logrus.Errorf("error reading config: %s", err.Error())
		os.Exit(1)
	}

	// init репозитории

	// init сервисы

	// init хендлеры

	// запуск сервера

	// graceful shutdown
}
