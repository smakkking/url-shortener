package httpserver

import (
	"context"
	"errors"
	"net/http"
	"sync"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/sirupsen/logrus"
	"github.com/smakkking/url-shortener/internal/app"
	"github.com/smakkking/url-shortener/internal/controllers/httphandlers"
)

type HTTPService struct {
	srv http.Server
	mux *chi.Mux
}

func NewServer(cfg app.Config) *HTTPService {
	service := &HTTPService{
		mux: chi.NewRouter(),
	}

	service.srv = http.Server{
		Addr:         cfg.HTTPAddress,
		Handler:      service.mux,
		ReadTimeout:  cfg.HTTPReadTimeout,
		WriteTimeout: cfg.HTTPWriteTimeout,
		IdleTimeout:  cfg.HTTPIdleTimeout,
	}

	return service
}

func (h *HTTPService) SetupHandlers(urlHandler *httphandlers.Handler) {
	// setup middleware
	h.mux.Use(middleware.RequestID)
	h.mux.Use(middleware.Recoverer)

	h.mux.Post("/create", urlHandler.SaveURL)
	h.mux.Get("/{alias}", urlHandler.GetURL)
}

func (h *HTTPService) Run(wg *sync.WaitGroup) {
	defer wg.Done()

	err := h.srv.ListenAndServe()
	if err != nil {
		if !errors.Is(err, http.ErrServerClosed) {
			logrus.Errorf("cannot start server: %s", err.Error())
		}
	}
}

func (h *HTTPService) Shutdown(ctx context.Context) {
	if err := h.srv.Shutdown(ctx); err != nil {
		logrus.Errorf("failed to stop server %s", err.Error())
		return
	}
}
