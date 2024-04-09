package grpchandlers

import (
	"context"
	"net/url"

	"github.com/smakkking/url-shortener/internal/services"
	"github.com/smakkking/url-shortener/pkg/sdk/go/urlshortener_grpc"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
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
	outputURL, err := s.urlService.GetURL(req.Alias)
	if err != nil {
		return nil, status.Error(codes.NotFound, "can't find url by this alias")
	}

	return &urlshortener_grpc.GetResponce{
		OriginalUrl: outputURL.String(),
	}, nil
}

func (s *serverAPI) Save(ctx context.Context, req *urlshortener_grpc.SaveRequest) (*urlshortener_grpc.SaveResponce, error) {
	inputURL, err := url.Parse(req.Url)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "give an url")
	}

	alias, err := s.urlService.SaveURL(*inputURL)
	if err != nil {
		return nil, status.Error(codes.Unknown, "internal error")
	}

	return &urlshortener_grpc.SaveResponce{
		Alias: alias,
	}, nil
}
