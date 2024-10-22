package handlers

// import (
// 	"backend/config"
// 	"backend/internal/domain"
// 	"backend/internal/middleware"

// 	"context"
// 	"encoding/json"
// 	"io"
// 	"log"
// 	"net/http"

// 	"go.uber.org/zap"
// )

// type MessageHandler struct {
// 	context   context.Context
// 	appConfig config.AppConfig
// 	lg        *zap.SugaredLogger
// }

// func NewLoginHandler(context context.Context, config config.AppConfig, logger *zap.SugaredLogger) *LoginHandler {
// 	return &LoginHandler{
// 		context:   context,
// 		appConfig: config,
// 		lg:        logger,
// 	}
// }

// func (h *LoginHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
// 	body, err := io.ReadAll(r.Body)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
