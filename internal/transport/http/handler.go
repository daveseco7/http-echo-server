package http

import (
	"net/http"

	"github.com/go-chi/chi"
	"go.uber.org/zap"
)

type Service interface {
	Success()
	Error400()
	Error500()
}

type handler struct {
	handler http.Handler
	service Service
	logger  *zap.SugaredLogger
}

// New instantiates a new http handler.
func New(s Service, l *zap.SugaredLogger) *handler {
	h := &handler{
		service: s,
		logger:  l,
	}

	r := chi.NewRouter()
	r.Get("/success", h.success)
	r.Get("/error400", h.error400)
	r.Get("/error500", h.error500)

	r.Post("/success", h.success)
	r.Post("/error400", h.error400)
	r.Post("/error500", h.error500)

	r.Put("/success", h.success)
	r.Put("/error400", h.error400)
	r.Put("/error500", h.error500)

	h.handler = r

	return h
}

// ServeHTTP implements the http handler interface.
func (h *handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.handler.ServeHTTP(w, r)
}
