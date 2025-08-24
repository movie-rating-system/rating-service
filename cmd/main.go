package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	app := New()
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer cancel()

	if err := app.Run(ctx); err != nil {
		log.Printf("App.Run error %v", err)
	}

	if err := app.Stop(); err != nil {
		log.Println(err)
		os.Exit(1)
	}
}
