package app

import (
	"context"
	"fmt"
	"log/slog"
	"net"
	"strconv"
	"time"

	"github.com/escape-ship/paymentsrv/config"
	"github.com/escape-ship/paymentsrv/internal/service"
	"github.com/escape-ship/paymentsrv/pkg/kafka"
	"github.com/escape-ship/paymentsrv/pkg/postgres"
	pb "github.com/escape-ship/paymentsrv/proto/gen"
	"google.golang.org/grpc"
)

type App struct {
	cfg   *config.Config
	pg    postgres.DBEngine
	kafka kafka.Engine
}

func New(cfg *config.Config, pg postgres.DBEngine, kafkaEngine kafka.Engine) *App {
	return &App{
		cfg:   cfg,
		pg:    pg,
		kafka: kafkaEngine,
	}
}

func (a *App) Run(ctx context.Context, logger *slog.Logger) error {
	server := grpc.NewServer()

	// Register Services
	pb.RegisterPaymentServiceServer(server, service.NewPaymentServer(a.cfg, a.pg, a.kafka))

	host := a.cfg.App.Host
	port := a.cfg.App.Port
	address := net.JoinHostPort(host, strconv.Itoa(port))
	l, err := net.Listen("tcp", address)
	if err != nil {
		return fmt.Errorf("failed to listen on %s: %v", address, err)
	}

	logger.Info("App: listening on", "address", address)

	// Run server in a separate goroutine
	go func() {
		if err := server.Serve(l); err != nil {
			if err != grpc.ErrServerStopped {
				logger.Error("App: serve error", "error", err)
			}
		}
	}()

	// Wait for context cancellation
	<-ctx.Done()
	logger.Info("App: initiating graceful shutdown")

	timer := time.AfterFunc(10*time.Second, func() {
		logger.Warn("Server couldn't stop gracefully in time. Doing force stop.")
		server.Stop()
	})
	defer timer.Stop()

	server.GracefulStop()
	logger.Info("App: completed graceful shutdown")

	return nil
}
