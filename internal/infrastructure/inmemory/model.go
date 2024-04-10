package inmemory

import (
	"errors"
	"sync"
)

var (
	ErrNotFound = errors.New("no such key, can't find value")
)

type Storage struct {
	db sync.Map
}

func NewStorage() *Storage {
	return &Storage{
		db: sync.Map{},
	}
}
