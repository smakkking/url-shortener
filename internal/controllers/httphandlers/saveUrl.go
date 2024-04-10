package httphandlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"github.com/go-chi/render"
	"github.com/sirupsen/logrus"
)

func (h *Handler) SaveURL(w http.ResponseWriter, r *http.Request) {
	const op = "httphandler.SaveURL"

	data := new(Request)
	err := json.NewDecoder(r.Body).Decode(data)
	if err != nil {
		logrus.Errorf("can't save url! %v", fmt.Errorf("%s: %w", op, err))
		w.WriteHeader(http.StatusBadRequest)
		render.JSON(w, r, Error("invalid json"))
		return
	}

	inputURL, err := url.Parse(data.URL)
	if err != nil {
		logrus.Errorf("can't save url! %v", fmt.Errorf("%s: %w", op, err))
		w.WriteHeader(http.StatusBadRequest)
		render.JSON(w, r, Error("invalid input data"))
		return
	}

	outputURL, err := h.urlService.SaveURL(r.Context(), *inputURL)
	if err != nil {
		logrus.Errorf("can't save url! %v", fmt.Errorf("%s: %w", op, err))
		w.WriteHeader(http.StatusInternalServerError)
		render.JSON(w, r, Error("can't save url"))
		return
	}

	render.JSON(w, r, OK("http://"+r.URL.Host+"/"+outputURL))
}
