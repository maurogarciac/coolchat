package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"frontend/config"
	i "frontend/internal"
	"frontend/internal/api"
	"frontend/internal/common"
	"frontend/internal/services"

	"github.com/caarlos0/env/v10"
	"github.com/hashicorp/go-retryablehttp"
	"go.uber.org/zap"
)

func main() {
	fmt.Println("Web server starting")

	appConfig := readAppConfig()
	appContext := context.Background()
	httpConfig := readHTTPConfig()
	httpClient := initHTTPClient(httpConfig)
	logger := common.NewLogger(appConfig.LogLevel)
	defer func(logger *zap.SugaredLogger) {
		err := logger.Sync()
		if err != nil && !errors.Is(err, syscall.ENOTTY) {
			logger.Warnf("Failed to flush log buffer: %v", err)
		}
	}(logger)

	backendApi := api.NewBackendApi(httpClient, httpConfig, logger)
	backendService := services.NewBackendService(appConfig, backendApi)

	httpServer := i.NewHTTPServer(appConfig, logger, backendService)
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

func readHTTPConfig() *config.HTTPConfig {
	cfg := &config.HTTPConfig{}
	if err := env.Parse(cfg); err != nil {
		log.Fatalf("Failed to read http configs: %v", err)
	}

	return cfg
}

func initHTTPClient(cfg *config.HTTPConfig) *http.Client {
	retryClient := retryablehttp.NewClient()
	retryClient.RetryMax = cfg.RetryMax

	return retryClient.StandardClient() // timeout for retryable http client is 30s
}
