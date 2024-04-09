package httphandlers

import "github.com/smakkking/url-shortener/internal/services"

func NewHandler(service *services.Service) *Handler {
	return &Handler{
		urlService: service,
	}
}
