package main

import (
	"github.com/escape-ship/paymentsrv/internal/app"
)

func main() {
	app := app.NewApp()

	app.Run()
}
