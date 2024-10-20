package main

import (
	"backend/config"
	"backend/internal"
	"backend/internal/common"
	"backend/internal/db"

	"context"
	"errors"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/caarlos0/env/v10"
	"go.uber.org/zap"
)

func main() {
	fmt.Println("Backend server started")

	appConfig := readAppConfig()
	appContext := context.Background()

	logger := common.NewLogger(appConfig.LogLevel)
	defer func(logger *zap.SugaredLogger) {
		err := logger.Sync()
		if err != nil && !errors.Is(err, syscall.ENOTTY) {
			logger.Warnf("Failed to flush log buffer: %v", err)
		}
	}(logger)

	dbConfig := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		&appConfig.DbHost, &appConfig.DbPort, &appConfig.DbUser, &appConfig.DbPassword, &appConfig.DbName)

	database := db.NewdbProvider(dbConfig, appContext, logger)

	httpServer := internal.NewHTTPServer(appConfig, logger, database)
	go httpServer.Start(appContext)
	defer func(appContext context.Context) {
		err := httpServer.Shutdown(appContext)
		if err != nil {
			logger.Errorf("Failed to shutdown server: %v", err)
		}
	}(appContext)

	ctx, stop := signal.NotifyContext(appContext, os.Interrupt, os.Kill, syscall.SIGTERM, syscall.SIGINT)
	defer stop()

	<-ctx.Done()
}

func readAppConfig() *config.AppConfig {
	cfg := &config.AppConfig{}
	if err := env.Parse(cfg); err != nil {
		log.Fatalf("Failed to read app configs: %v", err)
	}

	return cfg
}
