package main

import (
	"context"
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
	"github.com/smakkking/url-shortener/internal/services"
)

const (
	configPath = "./config/config.yaml"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

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
