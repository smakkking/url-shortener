package httphandlers

import (
	"github.com/smakkking/url-shortener/internal/services"
)

type Handler struct {
	urlService *services.Service
}

func NewHandler(service *services.Service) *Handler {
	return &Handler{
		urlService: service,
	}
}
