package handlers

import (
	"frontend/internal/templates"
	"net/http"

	"github.com/a-h/templ"
	"go.uber.org/zap"
)

// Checks if the GET request requires a full page or just main content
func pageRender(templateName string, c templ.Component, hasToken bool, lg *zap.SugaredLogger, w http.ResponseWriter, r *http.Request) {

	partial := r.URL.Query()["partial"]

	if partial != nil {
		err := c.Render(r.Context(), w) // Render only main content

		if err != nil {
			http.Error(w, "Error rendering partial template", http.StatusInternalServerError)
			return
		} else {
			lg.Debugf("partial %s template rendered", templateName)
		}
	} else {
		err := templates.Layout(c, hasToken).Render(r.Context(), w) // Render full page

		if err != nil {
			http.Error(w, "Error rendering full template", http.StatusInternalServerError)
			return
		} else {
			lg.Debugf("%s template rendered", templateName)
		}
	}

}
