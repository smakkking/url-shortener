package grpchandlers

import (
	"context"
	"fmt"

	"github.com/smakkking/url-shortener/internal/services"
	"github.com/smakkking/url-shortener/pkg/sdk/go/urlshortener_grpc"

	"google.golang.org/grpc"
)

type serverAPI struct {
	// вот эта хрень нужна, если вы неполностью реализовали service,
	// там просто заглушки вместо методов стоят
	urlshortener_grpc.UnimplementedURLShortenerServer
	urlService *services.Service
}

// вот эту хрень нужно будет вызывать в main, чтобы получить grpc сервер работающий
func Register(gRPC *grpc.Server) {
	urlshortener_grpc.RegisterURLShortenerServer(gRPC, &serverAPI{})
}

func (s *serverAPI) Get(ctx context.Context, req *urlshortener_grpc.GetRequest) (*urlshortener_grpc.GetResponce, error) {
	return &urlshortener_grpc.GetResponce{}
}

func (s *serverAPI) Add(ctx context.Context, req *calcv1.Request) (*calcv1.Responce, error) {
	res := req.A + req.B
	return &calcv1.Responce{Result: res}, nil
}

func (s *serverAPI) Sub(ctx context.Context, req *calcv1.Request) (*calcv1.Responce, error) {
	res := req.A - req.B
	return &calcv1.Responce{Result: res}, nil
}

func (s *serverAPI) Mul(ctx context.Context, req *calcv1.Request) (*calcv1.Responce, error) {
	res := req.A * req.B
	return &calcv1.Responce{Result: res}, nil
}

func (s *serverAPI) Div(ctx context.Context, req *calcv1.Request) (*calcv1.Responce, error) {
	if req.B == 0 {
		return nil, fmt.Errorf("division by 0")
	}

	res := req.A / req.B
	return &calcv1.Responce{Result: res}, nil
}
