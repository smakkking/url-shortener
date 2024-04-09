package inmemory

import "sync"

func NewStorage() *Storage {
	return &Storage{
		db: sync.Map{},
	}
}
