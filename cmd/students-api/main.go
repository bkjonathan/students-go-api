package main

import (
	"context"
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/bkjonathan/students-go-api/internal/config"
	"github.com/bkjonathan/students-go-api/internal/http/handlers/student"
)

func main() {
	cfg := config.LoadConfig()

	router := http.NewServeMux()

	router.HandleFunc("POST /api/students", student.Create())

	server := &http.Server{
		Addr:    fmt.Sprintf("%s:%d", cfg.Server.Host, cfg.Server.Port),
		Handler: router,
	}

	slog.Info("Starting server...", "host", cfg.Server.Host, "port", cfg.Server.Port)
	fmt.Printf("Starting server on %s:%d...\n", cfg.Server.Host, cfg.Server.Port)

	done := make(chan os.Signal, 1)

	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Error starting server: %v", err)
		}
	}()
	<-done

	slog.Info("Shutting down server...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("Error shutting down server: %v", err)
	}

	slog.Info("Server stopped gracefully")
}
