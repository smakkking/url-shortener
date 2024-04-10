package services

import (
	"context"
	"net/url"
)

type Service struct {
	urlStorage Storage
}

type Storage interface {
	SaveURL(context.Context, string, url.URL) error
	GetURL(context.Context, string) (url.URL, error)
}
