package httphandlers

import (
	"encoding/json"
	"net/http"
	"net/url"

	"github.com/go-chi/render"

	"github.com/smakkking/url-shortener/internal/services"
)

type Handler struct {
	urlService *services.Service
}

type R struct {
	URL url.URL `json:"url"`
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

	tmp, _ := url.Parse(outputURL)

	render.JSON(w, r, R{URL: *tmp})
}

func (h *Handler) GetURL(w http.ResponseWriter, r *http.Request) {
	inputURL := new(R)
	err := json.NewDecoder(r.Body).Decode(&inputURL)
	if err != nil {
		render.JSON(w, r, Error("invalid input json"))
		return
	}

	outputURL, err := h.urlService.GetURL(inputURL.URL.String())
	if err != nil {
		render.JSON(w, r, Error("can't save url"))
		return
	}

	render.JSON(w, r, R{URL: outputURL})
}
