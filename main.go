package main

import (
	"api-go/db"
	"api-go/router"
	"log/slog"
	"net/http"
	"os"
	"time"
)

func main() {
	if err := run(); err != nil {
		slog.Error("failed to execute application", "error", err)
		os.Exit(1)
	}
	slog.Info("application stopped")
}

func run() error {
	appDB, err := db.NewAppDB()
	if err != nil {
		slog.Error("failed to initialize application database", "error", err)
		return err
	}

	appRouter := router.NewRouter(*appDB)
	slog.Info("starting server")

	s := http.Server{
		ReadTimeout:  5 * time.Second,
		IdleTimeout:  120 * time.Second,
		WriteTimeout: 5 * time.Second,
		Addr:         ":8080",
		Handler:      appRouter,
	}

	if err := s.ListenAndServe(); err != nil {
		return err
	}

	return nil
}
