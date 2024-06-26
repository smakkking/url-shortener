package grpcserver

import (
	"context"
	"net"
	"sync"

	"github.com/sirupsen/logrus"
	"github.com/smakkking/url-shortener/internal/app"
	"github.com/smakkking/url-shortener/internal/controllers/grpchandlers"
	"github.com/smakkking/url-shortener/internal/services"
	"github.com/smakkking/url-shortener/pkg/sdk/go/urlshortener_grpc"
	"google.golang.org/grpc"
)

type GRPCService struct {
	grpcServer *grpc.Server
	Port       string
}

// вот эту нужно будет вызывать в main, чтобы получить grpc сервер работающий
func NewGRPCServer(config app.Config) *GRPCService {
	return &GRPCService{
		grpcServer: grpc.NewServer(),
		Port:       config.GrpcPort,
	}
}

func (s *GRPCService) RegisterHandlers(service *services.Service) {
	// здесь регистрирутся вообще все сервисы, которые участвуют в GRPC

	urlshortener_grpc.RegisterURLShortenerServer(
		s.grpcServer,
		&grpchandlers.ServerAPI{
			UrlService: service,
		})
}

func (s *GRPCService) Run(wg *sync.WaitGroup) {
	defer wg.Done()

	l, err := net.Listen("tcp", ":"+s.Port)
	if err != nil {
		panic("cant listen socket")
	}
	defer l.Close()

	if err = s.grpcServer.Serve(l); err != nil {
		logrus.Errorf("cannot start server: %s", err.Error())
	}
}

func (s *GRPCService) Shutdown(ctx context.Context) {
	s.grpcServer.GracefulStop()
}
