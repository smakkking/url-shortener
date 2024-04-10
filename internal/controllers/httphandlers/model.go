package httphandlers

import (
	"encoding/json"
	"fmt"
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
	const op = "httphandler.SaveURL"

	data := new(R)
	err := json.NewDecoder(r.Body).Decode(data)
	if err != nil {
		logrus.Errorf("can't save url! %v", fmt.Errorf("%s: %w", op, err))
		render.JSON(w, r, Error("invalid input json"))
		return
	}

	inputURL, err := url.Parse(data.URL)
	if err != nil {
		logrus.Errorf("can't save url! %v", fmt.Errorf("%s: %w", op, err))
		render.JSON(w, r, Error("invalid input data"))
		return
	}

	outputURL, err := h.urlService.SaveURL(r.Context(), *inputURL)
	if err != nil {
		logrus.Errorf("can't save url! %v", fmt.Errorf("%s: %w", op, err))
		render.JSON(w, r, Error("can't save url"))
		return
	}

	render.JSON(w, r, R{URL: "http://" + r.URL.Host + "/" + outputURL})
}

func (h *Handler) GetURL(w http.ResponseWriter, r *http.Request) {
	const op = "httphandler.GetURL"

	alias := chi.URLParam(r, "alias")

	outputURL, err := h.urlService.GetURL(r.Context(), alias)
	if err != nil {
		logrus.Errorf("can't get url! %v", fmt.Errorf("%s: %w", op, err))
		render.JSON(w, r, Error("can't get url"))
		return
	}

	render.JSON(w, r, R{URL: outputURL.String()})
}
