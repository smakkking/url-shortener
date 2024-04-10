package inmemory

import (
	"context"
	"fmt"
	"net/url"
)

func (s *Storage) GetURL(ctx context.Context, key string) (url.URL, error) {
	const op = "inmemory.GetURL"

	data, ok := s.db.Load(key)
	if !ok {
		return url.URL{}, fmt.Errorf("%s: %w", op, ErrNotFound)
	}

	return data.(url.URL), nil
}
