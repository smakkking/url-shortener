package httphandlers

import (
	"encoding/json"
	"net/http"
	"net/url"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"github.com/sirupsen/logrus"

	"github.com/smakkking/url-shortener/internal/services"
)

type Handler struct {
	urlService *services.Service
}

type R struct {
	URL string `json:"url"`
}

func (h *Handler) SaveURL(w http.ResponseWriter, r *http.Request) {
	data := new(R)
	err := json.NewDecoder(r.Body).Decode(data)
	if err != nil {
		logrus.Errorf("%s", err.Error())
		render.JSON(w, r, Error("invalid input json"))
		return
	}

	inputURL, err := url.Parse(data.URL)
	if err != nil {
		logrus.Errorf("%s", err.Error())
		render.JSON(w, r, Error("invalid input data"))
		return
	}

	outputURL, err := h.urlService.SaveURL(r.Context(), *inputURL)
	if err != nil {
		render.JSON(w, r, Error("can't save url"))
		return
	}

	render.JSON(w, r, R{URL: outputURL})
}

func (h *Handler) GetURL(w http.ResponseWriter, r *http.Request) {
	alias := chi.URLParam(r, "alias")

	outputURL, err := h.urlService.GetURL(r.Context(), alias)
	if err != nil {
		render.JSON(w, r, Error("can't get url"))
		return
	}

	render.JSON(w, r, R{URL: outputURL.String()})
}
