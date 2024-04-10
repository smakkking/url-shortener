package inmemory

import (
	"context"
	"net/url"
)

func (s *Storage) SaveURL(ctx context.Context, key string, urlToSave url.URL) error {
	s.db.Store(key, urlToSave)
	return nil
}
