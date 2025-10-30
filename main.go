package main

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"time"
)

func main() {
	httpServer, err := NewServer()
	if err != nil {
		slog.Error("Failed to construct HTTP server", "error", err)
		os.Exit(1)
	}

	fmt.Println("Starting HTTP server")
	if err := httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		slog.Error("Failed to start HTTP server", "error", err)
		os.Exit(1)
	}
}

func NewServer() (*http.Server, error) {
	thingService := NewThingService()
	strictHandler := NewStrictHandler(thingService, nil)
	h := Handler(strictHandler)

	srv := &http.Server{
		Addr:              ":8080",
		Handler:           h,
		ReadHeaderTimeout: 5 * time.Second,
	}

	return srv, nil
}

type thingService struct{}

func NewThingService() *thingService {
	return &thingService{}
}

// Create Thing
// (POST /things)
func (s *thingService) CreateThing(ctx context.Context, request CreateThingRequestObject) (CreateThingResponseObject, error) {
	slog.Info("CreateThing", "request", request)
	return CreateThing201Response{}, nil
}
