package services

import (
	"errors"
	"net/url"

	"github.com/smakkking/url-shortener/internal/models"
)

var (
	ErrSavingURL  = errors.New("error while saving url")
	ErrGettingURL = errors.New("error while saving url")
)

type Service struct {
	urlStorage Storage
}

type Storage interface {
	SaveURL(url.URL) (models.URLKey, error)
	GetURL(models.URLKey) (url.URL, error)
}

func (s *Service) SaveURL(urlToSave url.URL) (models.URLKey, error) {
	return s.urlStorage.SaveURL(urlToSave)

}

func (s *Service) GetURL(key models.URLKey) (url.URL, error) {
	return s.urlStorage.GetURL(key)
}
