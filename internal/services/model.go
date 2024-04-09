package services

import (
	"errors"
	"net/url"

	"github.com/smakkking/url-shortener/pkg/keygenerator"
)

var (
	ErrSavingURL  = errors.New("error while saving url")
	ErrGettingURL = errors.New("error while saving url")
)

type Service struct {
	urlStorage Storage
}

type Storage interface {
	SaveURL(string, url.URL) error
	GetURL(string) (url.URL, error)
}

func (s *Service) SaveURL(urlToSave url.URL) (string, error) {
	key := keygenerator.GenRandomString(10)

	err := s.urlStorage.SaveURL(key, urlToSave)
	if err != nil {
		return "", ErrSavingURL
	}

	return "http://localhost:8080/" + key, nil

}

func (s *Service) GetURL(key string) (url.URL, error) {
	return s.urlStorage.GetURL(key)
}
