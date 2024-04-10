package services

import (
	"context"
	"fmt"
	"net/url"
)

func (s *Service) SaveURL(ctx context.Context, urlToSave url.URL) (string, error) {
	const op = "service.SaveURL"

	key := s.keyGenerator.GenRandomString(10)

	key, err := s.urlStorage.SaveURL(ctx, key, urlToSave)
	if err != nil {
		return "", fmt.Errorf("%s: %w", op, err)
	}

	return key, nil
}
