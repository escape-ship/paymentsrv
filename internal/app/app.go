package app

import (
	"context"
	"fmt"
	"log/slog"
	"net"
	"strconv"
	"time"

	"github.com/escape-ship/paymentsrv/config"
	"google.golang.org/grpc"
)

type App struct {
	cfg  *config.Config
	ctx  context.Context
	done chan struct{}
}

func New(cfg *config.Config, ctx context.Context) *App {
	return &App{
		cfg:  cfg,
		ctx:  ctx,
		done: make(chan struct{}),
	}
}

func (a *App) Run() error {
	server := grpc.NewServer()

	// Graceful shutdown handler
	go func() {
		<-a.ctx.Done()
		slog.Info("App: initiating graceful shutdown")
		timer := time.AfterFunc(10*time.Second, func() {
			slog.Warn("Server couldn't stop gracefully in time. Doing force stop.")
			server.Stop()
		})
		defer timer.Stop()
		server.GracefulStop()
		slog.Info("App: completed graceful shutdown")
		close(a.done)
	}()

	host := a.cfg.App.Host
	port := a.cfg.App.Port
	address := net.JoinHostPort(host, strconv.Itoa(port))
	l, err := net.Listen("tcp", address)
	if err != nil {
		return fmt.Errorf("failed to listen on %s: %v", address, err)
	}

	slog.Info("App: listening on", "address", address)
	if err := server.Serve(l); err != nil {
		if err != grpc.ErrServerStopped {
			slog.Error("App: serve error", "error", err)
			return fmt.Errorf("failed to serve: %v", err)
		}
	}

	return nil
}

// Done returns a channel that is closed when the app has completed shutdown
func (a *App) Done() <-chan struct{} {
	return a.done
}
