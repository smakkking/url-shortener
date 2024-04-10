package httphandlers

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"github.com/sirupsen/logrus"
)

func (h *Handler) GetURL(w http.ResponseWriter, r *http.Request) {
	const op = "httphandler.GetURL"

	alias := chi.URLParam(r, "alias")

	outputURL, err := h.urlService.GetURL(r.Context(), alias)
	if err != nil {
		logrus.Errorf("can't get url! %v", fmt.Errorf("%s: %w", op, err))

		w.WriteHeader(http.StatusNotFound)
		render.JSON(w, r, Error("can't get url"))
		return
	}

	render.JSON(w, r, OK(outputURL.String()))
}
