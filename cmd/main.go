package main

import (
	"context"
	"log"

	"github.com/your-org/my-app/internal/config"
	"github.com/your-org/my-app/internal/handler"
	"github.com/your-org/my-app/internal/logger"
	"github.com/your-org/my-app/internal/metrics"
	"github.com/your-org/my-app/internal/middleware"
	"github.com/your-org/my-app/internal/server"
)

func main() {
	// Load configuration
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	// Initialize logger
	logger, err := logger.New(cfg.LogLevel)
	if err != nil {
		log.Fatalf("Failed to initialize logger: %v", err)
	}

	// Initialize metrics
	metrics := metrics.New()

	// Create HTTP handler
	httpHandler := handler.New(logger, metrics)

	// Apply middleware
	httpHandler = middleware.AccessLog(logger)(httpHandler)
	httpHandler = middleware.InputSanitization()(httpHandler)
	httpHandler = metrics.Middleware()(httpHandler)

	// Create HTTP server
	srv := server.New(cfg.ServerAddress, httpHandler)

	// Start server
	ctx := context.Background()
	err = srv.Start(ctx)
	if err != nil {
		logger.Error("Failed to start server: %v", err)
	}
}
