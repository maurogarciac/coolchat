package handlers

import (
	"net/http"

	"go.uber.org/zap"
)

type LandingHandler struct {
	lg *zap.SugaredLogger
}

func NewLandingHandler(logger *zap.SugaredLogger) *LandingHandler {
	return &LandingHandler{
		lg: logger,
	}
}

func (h LandingHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	redirect_link := "/home/"

	switch r.Method {

	case http.MethodGet:

		h.lg.Info("Landing GET")

		http.Redirect(w, r, redirect_link, http.StatusMovedPermanently)

	default:
		h.lg.Error("Only GET method is supported")
		http.Error(w, "Only GET method is supported", http.StatusMethodNotAllowed)
	}
}
