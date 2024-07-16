package main

import (
	"log/slog"
	"s3ai-testtask/internal/application"
	"s3ai-testtask/internal/infrastructure/config"
	"s3ai-testtask/internal/infrastructure/logger"
	"s3ai-testtask/internal/presentation/rest"
)

func main() {
	if err := run(); err != nil {
		panic(err)
	}
}

// run is the main entry point for the application
func run() error {
	// config
	cfg, err := config.Get()
	if err != nil {
		return err
	}

	// logger
	err = logger.Set(cfg.Env)
	if err != nil {
		return err
	}

	// service
	atmService := application.NewATMService()

	// rest
	srv, err := rest.NewServer(atmService, cfg)
	if err != nil {
		return err
	}

	slog.Info("Starting server...")

	return srv.ListenAndServe()
}
