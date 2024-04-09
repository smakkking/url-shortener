package httphandlers

import "github.com/smakkking/url-shortener/internal/services"

type Handler struct {
	urlService *services.Service
}
