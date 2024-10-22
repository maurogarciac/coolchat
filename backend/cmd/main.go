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
		if err := logger.Sync(); err != nil && !errors.Is(err, syscall.ENOTTY) {
			logger.Warnf("Failed to flush log buffer: %v", err)
		}
	}(logger)

	dbConfig := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s?sslmode=disable",
		appConfig.DbUser,
		appConfig.DbPassword,
		appConfig.DbHost,
		appConfig.DbPort,
		appConfig.DbName)

	database, err := db.NewDbProvider(dbConfig, logger)
	if err != nil {
		logger.Error("could not connect to db")
	}
	if err := database.SetupDb(); err != nil {
		logger.Error(err)
	}

	httpServer := internal.NewHTTPServer(appConfig, logger, database)
	go httpServer.Start(appContext)
	defer func(appContext context.Context) {
		if err := httpServer.Shutdown(appContext); err != nil {
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
