package internal

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	"backend/config"
	"backend/internal/db"
	"backend/internal/handlers"
	"backend/internal/ws"

	"go.uber.org/zap"
)

const serverShutdownTimeoutDuration = 60 * time.Second

type HTTPServer struct {
	cfg *config.AppConfig
	sv  *http.Server
	lg  *zap.SugaredLogger
	db  *db.DbProvider
}

func NewHTTPServer(
	appConfig *config.AppConfig,
	logger *zap.SugaredLogger,
	database *db.DbProvider,
) *HTTPServer {
	server := http.Server{
		Addr:        fmt.Sprintf(":%d", appConfig.ServerPort),
		ReadTimeout: appConfig.ReadTimeout,
	}

	return &HTTPServer{
		cfg: appConfig,
		sv:  &server,
		lg:  logger,
	}
}

func (s *HTTPServer) Start(ctx context.Context) {
	s.lg.Infof("Starting server on port %d", s.cfg.ServerPort)

	http.Handle("/ws", ws.NewChatServer(s.lg, s.db))
	http.Handle("/auth", handlers.NewJwtHandler(s.lg, s.cfg))
	http.Handle("/refresh", handlers.NewRefreshTokenHandler(s.lg, s.cfg))
	http.HandleFunc("/health",
		func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte("OK"))
		})

	err := s.sv.ListenAndServe()
	for {
		if errors.Is(err, http.ErrServerClosed) {
			s.lg.Info("Server closed")
		} else if err != nil {
			s.lg.Fatal("Failed to start server", zap.Error(err))
		}
	}
}

func (s *HTTPServer) Shutdown(ctx context.Context) error {
	s.lg.Info("Shutting down server")
	if s.sv != nil {
		shutdownCtx, shutdownRelease := context.WithTimeout(ctx, serverShutdownTimeoutDuration)
		err := s.sv.Shutdown(shutdownCtx)
		shutdownRelease()
		return err
	}
	return nil
}
