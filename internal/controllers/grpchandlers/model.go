package grpchandlers

import (
	"github.com/smakkking/url-shortener/internal/services"
	"github.com/smakkking/url-shortener/pkg/sdk/go/urlshortener_grpc"
)

type ServerAPI struct {
	// вот эта нужна, если вы неполностью реализовали service,
	// там просто заглушки вместо методов стоят
	urlshortener_grpc.UnimplementedURLShortenerServer
	UrlService *services.Service
}
