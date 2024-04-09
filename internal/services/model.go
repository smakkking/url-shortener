package services

import (
	"context"
	"errors"
	"net/url"

	"github.com/smakkking/url-shortener/pkg/keygenerator"
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

func (s *Service) SaveURL(ctx context.Context, urlToSave url.URL) (string, error) {
	key := keygenerator.GenRandomString(10)

	err := s.urlStorage.SaveURL(ctx, key, urlToSave)
	if err != nil {
		return "", ErrSavingURL
	}

	return "http://localhost:8080/" + key, nil

}

func (s *Service) GetURL(ctx context.Context, key string) (url.URL, error) {
	return s.urlStorage.GetURL(ctx, key)
}
