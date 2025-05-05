package main

import (
	"context"
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	"github.com/escape-ship/paymentsrv/config"
	"github.com/escape-ship/paymentsrv/internal/app"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	dir, err := os.Getwd()
	if err != nil {
		slog.Error("App: get current directory error", "error", err)
		os.Exit(1)
	}
	slog.Info("App: current directory", "dir", dir)

	vp := config.NewConfig(dir + "/config.yaml")
	cfg, err := config.Load(vp)
	if err != nil {
		slog.Error("App: config load error", "error", err)
		os.Exit(1)
	}

	// Set up signal handling
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	app := app.New(cfg, ctx)

	// Run application in a separate goroutine
	errChan := make(chan error, 1)
	go func() {
		if err := app.Run(); err != nil {
			slog.Error("App: run error", "error", err)
			errChan <- err
		}
	}()

	// Wait for shutdown signal or error
	select {
	case err := <-errChan:
		slog.Error("App: application error", "error", err)
		cancel()
	case sig := <-quit:
		slog.Info("App: received shutdown signal", "signal", sig)
		cancel()
	}

	// Wait for actual graceful shutdown to complete
	<-app.Done()
	slog.Info("App: shutdown complete")
}
