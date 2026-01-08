package main

import (
	"log"
	"log/slog"
	"os"
)

func main() {
	cfg := config{
		addr: ":8080",
		db:   dbConfig{},
	}
	api := application{
		config: cfg,
	}

	// structure Logging
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	slog.SetDefault(logger)

	if err := api.run(api.mount()); err != nil {
		log.Printf("Server Failed to start : %v", err)
		os.Exit(1)
	}
}
