package httphandlers

import (
	"encoding/json"
	"net/http"
	"net/url"

	"github.com/go-chi/render"

	"github.com/smakkking/url-shortener/internal/models"
	"github.com/smakkking/url-shortener/internal/services"
)

type Handler struct {
	urlService *services.Service
}

type R struct {
	URL url.URL `json:"url"`
}

type ResponseWithAlias struct {
	Response
	Alias string `json:"alias,omitempty"`
}

func (h *Handler) SaveURL(w http.ResponseWriter, r *http.Request) {
	inputURL := new(R)
	err := json.NewDecoder(r.Body).Decode(&inputURL)
	if err != nil {
		render.JSON(w, r, Error("invalid input json"))
		return
	}

	outputURL, err := h.urlService.SaveURL(inputURL.URL)
	if err != nil {
		render.JSON(w, r, Error("can't save url"))
		return
	}

	render.JSON(w, r, R{URL: outputURL.Transform()})
}

func (h *Handler) GetURL(w http.ResponseWriter, r *http.Request) {
	inputURL := new(R)
	err := json.NewDecoder(r.Body).Decode(&inputURL)
	if err != nil {
		render.JSON(w, r, Error("invalid input json"))
		return
	}

	outputURL, err := h.urlService.GetURL(models.URLKey(inputURL.URL.String()))
	if err != nil {
		render.JSON(w, r, Error("can't save url"))
		return
	}

	render.JSON(w, r, R{URL: outputURL})
}
