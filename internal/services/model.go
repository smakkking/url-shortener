package services

import (
	"context"
	"errors"
	"net/url"
)

var (
	ErrSavingURL  = errors.New("error while saving url")
	ErrGettingURL = errors.New("error while saving url")
)

type Service struct {
	urlStorage Storage
}

type Storage interface {
	SaveURL(context.Context, string, url.URL) error
	GetURL(context.Context, string) (url.URL, error)
}
