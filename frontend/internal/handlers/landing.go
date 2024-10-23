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

	switch r.Method {

	case http.MethodGet:

		w.Header().Set("HX-Redirect", "/home?partial=true")

	default:
		h.lg.Error("Only GET method is supported")
		http.Error(w, "Only GET method is supported", http.StatusMethodNotAllowed)
	}
}
