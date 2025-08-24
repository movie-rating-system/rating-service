package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/kirillApanasiuk/kit/pkg/discovery"
	"github.com/kirillApanasiuk/kit/pkg/discovery/counsul"
	"github.com/kirillApanasiuk/movie-rating/infrastructure/persistence/reporitory"
	http2 "github.com/kirillApanasiuk/movie-rating/internal/controller/http"
	"github.com/kirillApanasiuk/movie-rating/usecase/rating"
)

const (
	portNumber   = 8082
	serviceName  = "rating"
	discoveryUrl = "localhost:8500"
)

type app struct {
	server     *http.Server
	port       int
	registry   *counsul.Registry
	instanceId string
	errChan    chan error
}

func New() *app {
	repo := reporitory.NewRepository()
	ctrl := rating.New(repo)
	h := http2.New(ctrl)

	mux := http.NewServeMux()
	mux.Handle("/rating", http.HandlerFunc(h.Handle))

	var port int
	flag.IntVar(&port, "port", portNumber, "Port to listen on")
	flag.Parse()

	registry, err := counsul.NewRegistry(discoveryUrl)
	if err != nil {
		log.Panic(err)
	}

	return &app{
		server: &http.Server{
			Addr:    fmt.Sprintf(":%d", port),
			Handler: mux,
		},
		port:       port,
		registry:   registry,
		instanceId: discovery.GenerateInstanceID(serviceName),
	}
}

func (a *app) Run(ctx context.Context) error {
	if err := a.registry.Register(ctx, a.instanceId, serviceName, discoveryUrl); err != nil {
		panic(err)
	}

	go a.reportHealthy(ctx)

	serverErrors := make(chan error, 1)
	go func() {
		log.Printf("Starting the rating service on port %d\n", a.port)
		serverErrors <- a.server.ListenAndServe()
	}()

	select {
	case err := <-serverErrors:
		if err != nil && !errors.Is(err, http.ErrServerClosed) {
			return fmt.Errorf("server error: %w", err)
		}
		return err
	case <-ctx.Done():
		log.Printf("Received shutdown signal for service on port %d", a.port)
		return ctx.Err()
	}
}

func (a *app) reportHealthy(ctx context.Context) {
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()
	for {
		select {
		case <-ticker.C:
			if err := a.registry.ReportHealthyState(a.instanceId, serviceName); err != nil {
				log.Println("Failed to log healthy state " + err.Error())
				return
			}
		case <-ctx.Done():
			log.Printf("Stopping the rating service on port %d", a.port)
			return
		}
	}
}

func (a *app) Stop() error {
	log.Println("Shutting down server...")
	shutdownCtx, shutdownCancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer shutdownCancel()

	if err := a.server.Shutdown(shutdownCtx); err != nil {
		log.Printf("Graceful shutdown failed: %v", err)
		log.Println("Forcing server close...")
		if err = a.server.Close(); err != nil {
			log.Printf("Forse close failed: %v", err)
			return err
		}
	}

	deregisterCtx, deregisterCancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer deregisterCancel()

	if err := a.registry.Deregister(deregisterCtx, a.instanceId, serviceName); err != nil {
		log.Printf("Failed to deregister service: %v", err)
		return err
	}

	log.Println("Server successfully stopped")
	log.Println("Server deregistered successfully")
	return nil
}
