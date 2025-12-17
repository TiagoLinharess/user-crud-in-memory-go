package main

import (
	"log/slog"
	"net/http"
	"os"
	"time"
	"userCrud/internal/api"
	"userCrud/internal/models"
)

func main() {
	if err := run(); err != nil {
		slog.Error("Failed to execute code", "error", err)
		os.Exit(1)
	}
	slog.Info("All systems offline")
}

func run() error {
	db := models.NewApplication()

	handler := api.NewHandler(db)

	s := http.Server{
		ReadTimeout:  10 * time.Second,
		IdleTimeout:  time.Minute,
		WriteTimeout: 10 * time.Second,
		Addr:         ":8080",
		Handler:      handler,
	}

	if err := s.ListenAndServe(); err != nil {
		return err
	}

	return nil
}
