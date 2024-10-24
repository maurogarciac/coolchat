package handlers

import (
	"context"
	"net/http"
	"time"

	d "frontend/internal/domain"
	auth "frontend/internal/middleware"
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

	redirect_link := "/home/"

	c := templates.LogIn(false)

	switch r.Method {

	case http.MethodGet:

		if auth.GetAccessToken(r) != "" {
			w.Header().Set("HX-Redirect", redirect_link)
			w.WriteHeader(http.StatusOK)
		}

		pageRender("login", c, false, h.lg, w, r)

	case http.MethodPost:

		err := r.ParseForm()
		if err != nil {
			h.lg.Error(err)
		}

		// Maybe hash+salt this later if I have time

		login_creds := d.User{
			Username: r.FormValue("username"),
			Password: r.FormValue("password"),
		}

		h.lg.Info("User candidate: ", login_creds)

		res, err := h.b.PostLogin(h.ctx, login_creds)
		if err != nil {
			h.lg.Error("Could not authenticate user")
			templates.LogIn(true).Render(h.ctx, w)
			// w.Header().Set("HX-Redirect", "/login/")
			return
		}

		auth.SetTokenCookie(
			auth.AccessTokenCookieName,
			res.AccessToken,
			time.Now().Add(15*time.Minute), w)

		auth.SetTokenCookie(
			auth.RefreshTokenCookieName,
			res.RefreshToken,
			time.Now().Add(24*time.Hour), w)

		//http.Redirect(w, r, redirect_link, http.StatusMovedPermanently)
		w.Header().Set("HX-Redirect", redirect_link)
		w.WriteHeader(http.StatusOK)
	default:
		h.lg.Error("Only GET and POST methods are supported")
		http.Error(w, "Only GET and POST method aer supported", http.StatusMethodNotAllowed)
	}
}
