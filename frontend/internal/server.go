package resthttp

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	"frontend/config"
	"frontend/internal/handlers"

	"go.uber.org/zap"
)

const serverShutdownTimeoutDuration = 60 * time.Second

type HTTPServer struct {
	cfg    *config.AppConfig
	server *http.Server
	lg     *zap.SugaredLogger
}

func NewHTTPServer(
	appCfg *config.AppConfig,
	logger *zap.SugaredLogger,
) *HTTPServer {
	server := http.Server{
		Addr:        fmt.Sprintf(":%d", appCfg.ServerPort),
		ReadTimeout: appCfg.ReadTimeout,
	}

	return &HTTPServer{
		cfg:    appCfg,
		server: &server,
		lg:     logger,
	}
}

func (s *HTTPServer) Start() {
	s.lg.Infof("Starting web server on port %d", s.cfg.ServerPort)

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	http.Handle("/home/", handlers.NewHomeHandler(s.lg))
	http.Handle("/chat/", handlers.NewChatHandler(s.lg))
	http.Handle("/", handlers.NewLoginHandler(s.lg))
	http.Handle("/logout/", handlers.NewLogoutHandler(s.lg))

	err := s.server.ListenAndServe()
	if errors.Is(err, http.ErrServerClosed) {
		s.lg.Info("Web server closed")
	} else if err != nil {
		s.lg.Fatal("Failed to start web server", zap.Error(err))
	}
}

func (s *HTTPServer) Shutdown(ctx context.Context) error {
	s.lg.Info("Shutting down web server")
	if s.server != nil {
		shutdownCtx, shutdownRelease := context.WithTimeout(ctx, serverShutdownTimeoutDuration)
		err := s.server.Shutdown(shutdownCtx)
		shutdownRelease()
		return err
	}
	return nil
}
