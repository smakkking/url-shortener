package inmemory

import (
	"context"
	"errors"
	"fmt"
	"net/url"
	"sync"
)

var (
	ErrNotFound = errors.New("no such key, can't find value")
)

type Storage struct {
	db sync.Map
}

func (s *Storage) SaveURL(ctx context.Context, key string, urlToSave url.URL) error {
	s.db.Store(key, urlToSave)
	return nil
}

func (s *Storage) GetURL(ctx context.Context, key string) (url.URL, error) {
	const op = "inmemory.GetURL"

	data, ok := s.db.Load(key)
	if !ok {
		return url.URL{}, fmt.Errorf("%s: %w", op, ErrNotFound)
	}

	return data.(url.URL), nil
}
