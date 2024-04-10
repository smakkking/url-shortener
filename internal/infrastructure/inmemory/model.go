package inmemory

import (
	"errors"
	"net/url"
	"sync"
)

var (
	ErrNotFound = errors.New("no such key, can't find value")
)

type Storage struct {
	mu         sync.RWMutex
	aliasToURL map[string]url.URL
	urlToAlias map[url.URL]string
}

func NewStorage() *Storage {
	return &Storage{
		aliasToURL: make(map[string]url.URL),
		urlToAlias: make(map[url.URL]string),
	}
}
