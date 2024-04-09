package inmemory

import (
	"errors"
	"net/url"
	"sync"
)

var (
	ErrNotFound = errors.New("no such key, can't find value in memory")
)

type Storage struct {
	db sync.Map
}

func (s *Storage) SaveURL(key string, urlToSave url.URL) error {
	s.db.Store(key, urlToSave)
	return nil
}

func (s *Storage) GetURL(key string) (url.URL, error) {
	data, ok := s.db.Load(key)
	if !ok {
		return url.URL{}, ErrNotFound
	}

	return data.(url.URL), nil
}
