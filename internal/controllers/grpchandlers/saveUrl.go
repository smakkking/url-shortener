package grpchandlers

import (
	"context"
	"fmt"
	"net/url"

	"github.com/sirupsen/logrus"
	"github.com/smakkking/url-shortener/pkg/sdk/go/urlshortener_grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *ServerAPI) Save(ctx context.Context, req *urlshortener_grpc.SaveRequest) (*urlshortener_grpc.SaveResponce, error) {
	const op = "grpchandler.Save"
	inputURL, err := url.Parse(req.Url)
	if err != nil {
		logrus.Errorf("invalid params! %v", fmt.Errorf("%s: %w", op, err))
		return nil, status.Error(codes.InvalidArgument, "give an url")
	}

	alias, err := s.urlService.SaveURL(ctx, *inputURL)
	if err != nil {
		logrus.Errorf("can't save url! %v", fmt.Errorf("%s: %w", op, err))
		return nil, status.Error(codes.Unknown, "internal error")
	}

	return &urlshortener_grpc.SaveResponce{
		Alias: "http://localhost:8080/" + alias,
	}, nil
}
