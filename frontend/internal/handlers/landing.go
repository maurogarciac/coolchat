package handlers

import (
	"fmt"
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
		fmt.Fprintf(w, "only get and post methods are supported")
		return

	}
}
