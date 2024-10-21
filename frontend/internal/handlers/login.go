package handlers

import (
	"context"
	"fmt"
	"net/http"

	"frontend/internal/domain"
	"frontend/internal/services"
	"frontend/internal/templates"

	"go.uber.org/zap"
)

type LoginHandler struct {
	ctx context.Context
	lg  *zap.SugaredLogger
	b   *services.BackendService
}

func NewLoginHandler(context context.Context, logger *zap.SugaredLogger, backend *services.BackendService) *LoginHandler {
	return &LoginHandler{
		ctx: context,
		lg:  logger,
		b:   backend,
	}
}

func (h LoginHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	c := templates.LogIn()

	switch r.Method {

	case http.MethodGet:

		// should check headers to see if user has an access token and redirect to home page

		pageRender("login", c, h.lg, w, r)

	case http.MethodPost:

		// Should check cookies for jwt?
		err := r.ParseForm()
		if err != nil {
			h.lg.Error(err)
		}

		user := domain.User{
			Username: r.FormValue("username"),
			Password: r.FormValue("password"),
		}

		res, err := h.b.PostLogin(h.ctx, user)
		if err != nil {
			h.lg.Error(err)
		}
		if res.Message == "Accepted" {
			http.Redirect(w, r, "/home", http.StatusAccepted)
		}

	default:
		fmt.Fprintf(w, "only get and post methods are supported")
		return

	}
}
