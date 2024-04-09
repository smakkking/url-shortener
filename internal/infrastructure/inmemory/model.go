package inmemory

import (
	"errors"
	"net/url"
	"sync"

	"github.com/smakkking/url-shortener/internal/models"
	"github.com/smakkking/url-shortener/pkg/keygenerator"
)

var (
	ErrNotFound = errors.New("no such key")
)

type Storage struct {
	db sync.Map
}

func (s *Storage) SaveURL(urlToSave url.URL) (models.URLKey, error) {
	key := keygenerator.GenRandomString(10)
	s.db.Store(key, urlToSave)
	return key, nil
}

func (s *Storage) GetURL(key models.URLKey) (url.URL, error) {
	data, ok := s.db.Load(key)
	if !ok {
		return url.URL{}, ErrNotFound
	}

	return data.(url.URL), nil
}
