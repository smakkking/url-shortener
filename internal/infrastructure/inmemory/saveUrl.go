package inmemory

import (
	"context"
	"net/url"
)

func (s *Storage) SaveURL(ctx context.Context, key string, urlToSave url.URL) (string, error) {
	// если на эту ссылку уже есть запись, то мы
	s.mu.Lock()
	defer s.mu.Unlock()

	if oldKey, ok := s.urlToAlias[urlToSave]; ok {
		return oldKey, nil
	}

	s.aliasToURL[key] = urlToSave
	return key, nil
}
