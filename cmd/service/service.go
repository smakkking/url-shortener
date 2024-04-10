package main

import (
	"context"
	"errors"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/smakkking/url-shortener/internal/app"
	"github.com/smakkking/url-shortener/internal/controllers/httphandlers"
	"github.com/smakkking/url-shortener/internal/grpcserver"
	"github.com/smakkking/url-shortener/internal/httpserver"
	"github.com/smakkking/url-shortener/internal/infrastructure/inmemory"
	"github.com/smakkking/url-shortener/internal/infrastructure/postgres"
	"github.com/smakkking/url-shortener/internal/services"
	"github.com/smakkking/url-shortener/pkg/keygenerator"
)

const (
	configPath = "./config/config.yaml"
	dbType     = "postgres"
)

var (
	modeDev = os.Getenv("ENV") == ""
)

func main() {
	// init логгер
	setupLogger()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	logrus.Infoln("service started...")
	logrus.Debugln("debug messages are available")

	// загрузка конфига
	config, err := app.MustLoadConfig(configPath)

	// init репозитории
	var store services.Storage

	if dbType == "inmemory" {
		store = inmemory.NewStorage()
	} else if dbType == "postgres" {
		store, err = postgres.NewStorage(config)
		if err != nil {
			panic(err)
		}
	} else {
		panic(errors.New("no such storage type"))
	}

	// init сервисы
	urlService := services.NewService(store, &keygenerator.RandomKeyGenerator{})

	// init хендлеры
	urlHandler := httphandlers.NewHandler(urlService)

	// запуск сервера HTTP
	srvHTTP := httpserver.NewServer(config)
	srvHTTP.SetupHandlers(urlHandler)

	srvGRPC := grpcserver.NewGRPCServer(config)

	// running
	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	wg := &sync.WaitGroup{}
	wg.Add(2)

	go srvHTTP.Run(wg)
	go srvGRPC.Run(wg)

	logrus.Infoln("server started")

	// graceful shutdown
	<-done

	// TODO: move timeout to config
	shutDownCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	srvHTTP.Shutdown(shutDownCtx)
	srvGRPC.Shutdown(shutDownCtx)

	wg.Wait()
	logrus.Infoln("server stopped")
}

func setupLogger() {
	logrus.SetFormatter(
		&logrus.JSONFormatter{
			PrettyPrint:     true,
			TimestampFormat: "2006-01-02 15:04:05",
		},
	)
	logrus.SetOutput(os.Stdout)

	if modeDev {
		logrus.SetLevel(logrus.DebugLevel)
	} else {
		logrus.SetLevel(logrus.InfoLevel)
	}

}
