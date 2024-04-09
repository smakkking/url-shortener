package services

import (
	"net/url"

	"github.com/smakkking/url-shortener/internal/models"
)

type Service struct {
	urlStorage Storage
}

type Storage interface {
	SaveURL(url.URL) (models.URLKey, error)
	GetURL(models.URLKey) (url.URL, error)
}
