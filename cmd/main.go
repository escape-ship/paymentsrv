package main

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	"github.com/escape-ship/paymentsrv/config"
	"github.com/escape-ship/paymentsrv/internal/app"
	"github.com/escape-ship/paymentsrv/pkg/kafka"
	"github.com/escape-ship/paymentsrv/pkg/postgres"
)

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	cfg, err := config.New("config.yaml")
	if err != nil {
		logger.Error("App: config load error", "error", err)
		os.Exit(1)
	}

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer cancel()

	db, err := postgres.New(makeDSN(cfg.Database))
	if err != nil {
		logger.Error("App: database connection error", "error", err)
		os.Exit(1)
	}

	// Initialize Kafka consumer
	kafkaEngine := kafka.NewPublisher(
		[]string{"kafka:9092"}, // TODO: replace with cfg value if available
		"payment-succeeded",    // TODO: replace with cfg value if available
	)

	app := app.New(cfg, db, kafkaEngine)
	if err := app.Run(ctx, logger); err != nil {
		logger.Error("app stopped with error", "err", err)
	}
}

// config.Database 값 사용
func makeDSN(db config.Database) postgres.DBConnString {
	return postgres.DBConnString(
		fmt.Sprintf(
			"postgres://%s:%s@%s:%d/%s?sslmode=%s&search_path=%s",
			db.User, db.Password,
			db.Host, db.Port,
			db.DataBaseName, db.SSLMode, db.SchemaName,
		),
	)
}
