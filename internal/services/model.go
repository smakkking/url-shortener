package services

import (
	"context"
	"errors"
	"fmt"
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
	const op = "service.SaveURL"

	key := keygenerator.GenRandomString(10)

	err := s.urlStorage.SaveURL(ctx, key, urlToSave)
	if err != nil {
		return "", fmt.Errorf("%s: %w", op, ErrSavingURL)
	}

	return key, nil

}

func (s *Service) GetURL(ctx context.Context, key string) (url.URL, error) {
	const op = "service.GetURL"

	originalURL, err := s.urlStorage.GetURL(ctx, key)
	if err != nil {
		return originalURL, fmt.Errorf("%s: %w", op, ErrGettingURL)
	}

	return originalURL, nil
}
