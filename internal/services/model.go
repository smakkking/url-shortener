package services

import (
	"context"
	"net/url"
)

type Service struct {
	urlStorage   Storage
	keyGenerator KeyGen
}

type KeyGen interface {
	GenRandomString(size int) string
}

//go:generate mockgen -source=model.go -destination=mocks/mock.go
type Storage interface {
	SaveURL(context.Context, string, url.URL) (string, error)
	GetURL(context.Context, string) (url.URL, error)
}

func NewService(storage Storage, kg KeyGen) *Service {
	return &Service{
		urlStorage:   storage,
		keyGenerator: kg,
	}
}
