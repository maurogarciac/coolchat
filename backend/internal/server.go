package internal

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	"backend/config"
	"backend/internal/db"
	"backend/internal/ws"

	"go.uber.org/zap"
)

const serverShutdownTimeoutDuration = 60 * time.Second

type HTTPServer struct {
	cfg      *config.AppConfig
	server   *http.Server
	logger   *zap.SugaredLogger
	database *db.DbProvider
}

func NewHTTPServer(
	appCfg *config.AppConfig,
	lg *zap.SugaredLogger,
	db *db.DbProvider,
) *HTTPServer {
	server := http.Server{
		Addr:        fmt.Sprintf(":%d", appCfg.ServerPort),
		ReadTimeout: appCfg.ReadTimeout,
	}

	return &HTTPServer{
		cfg:    appCfg,
		server: &server,
		logger: lg,
	}
}

func (s *HTTPServer) Start(ctx context.Context) {
	s.logger.Infof("Starting server on port %d", s.cfg.ServerPort)

	// http.Handle("/login", handlers.NewLoginHandler(ctx, *s.cfg, s.logger))
	http.Handle("/ws", ws.NewChatServer(s.logger))

	err := s.server.ListenAndServe()
	for {
		if errors.Is(err, http.ErrServerClosed) {
			s.logger.Info("Server closed")
		} else if err != nil {
			s.logger.Fatal("Failed to start server", zap.Error(err))
		}
	}
}

func (s *HTTPServer) Shutdown(ctx context.Context) error {
	s.logger.Info("Shutting down server")
	if s.server != nil {
		shutdownCtx, shutdownRelease := context.WithTimeout(ctx, serverShutdownTimeoutDuration)
		err := s.server.Shutdown(shutdownCtx)
		shutdownRelease()
		return err
	}
	return nil
}
