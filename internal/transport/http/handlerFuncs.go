package http

import (
	"net/http"
)

func (h *handler) success(w http.ResponseWriter, r *http.Request) {
	h.logger.Infow("received request", "header", r.Header, "body", r.Body)
	h.service.Success()
	w.WriteHeader(http.StatusOK)
}

func (h *handler) error400(w http.ResponseWriter, r *http.Request) {
	h.logger.Infow("received request", "header", r.Header, "body", r.Body)
	h.service.Error400()
	w.WriteHeader(http.StatusBadRequest)
}

func (h *handler) error500(w http.ResponseWriter, r *http.Request) {
	h.logger.Infow("received request", "header", r.Header, "body", r.Body)
	h.service.Error500()
	w.WriteHeader(http.StatusInternalServerError)
}
