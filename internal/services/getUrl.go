package services

import (
	"context"
	"fmt"
	"net/url"
)

func (s *Service) GetURL(ctx context.Context, key string) (url.URL, error) {
	const op = "service.GetURL"

	originalURL, err := s.urlStorage.GetURL(ctx, key)
	if err != nil {
		return originalURL, fmt.Errorf("%s: %w", op, ErrGettingURL)
	}

	return originalURL, nil
}
