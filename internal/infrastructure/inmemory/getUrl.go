package inmemory

import (
	"context"
	"fmt"
	"net/url"
)

func (s *Storage) GetURL(ctx context.Context, key string) (url.URL, error) {
	const op = "inmemory.GetURL"

	s.mu.RLock()
	defer s.mu.RUnlock()

	data, ok := s.aliasToURL[key]
	if !ok {
		return url.URL{}, fmt.Errorf("%s: %w", op, ErrNotFound)
	}

	return data, nil
}
