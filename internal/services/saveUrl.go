package services

import (
	"context"
	"fmt"
	"net/url"

	"github.com/smakkking/url-shortener/pkg/keygenerator"
)

func (s *Service) SaveURL(ctx context.Context, urlToSave url.URL) (string, error) {
	const op = "service.SaveURL"

	key := keygenerator.GenRandomString(10)

	err := s.urlStorage.SaveURL(ctx, key, urlToSave)
	if err != nil {
		return "", fmt.Errorf("%s: %w", op, ErrSavingURL)
	}

	return key, nil
}
