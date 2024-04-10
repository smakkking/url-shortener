package grpchandlers

import (
	"context"
	"fmt"

	"github.com/sirupsen/logrus"
	"github.com/smakkking/url-shortener/pkg/sdk/go/urlshortener_grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *ServerAPI) Get(ctx context.Context, req *urlshortener_grpc.GetRequest) (*urlshortener_grpc.GetResponce, error) {
	const op = "grpchandler.Get"

	outputURL, err := s.UrlService.GetURL(ctx, req.Alias)
	if err != nil {
		logrus.Errorf("can't get url! %v", fmt.Errorf("%s: %w", op, err))
		return nil, status.Error(codes.NotFound, "can't find url by this alias")
	}

	return &urlshortener_grpc.GetResponce{
		OriginalUrl: outputURL.String(),
	}, nil
}
