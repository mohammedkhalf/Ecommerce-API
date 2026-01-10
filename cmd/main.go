package main

import (
	"context"
	"github.com/jackc/pgx/v5"
	"github.com/mohammedkhalf/Ecommerce-API/internal/env"
	"log"
	"log/slog"
	"os"
)

func main() {

	ctx := context.Background()

	cfg := config{
		addr: ":8080",
		db: dbConfig{
			dsn: env.GetString("GOOSE_DBSTRING", "host=localhost user=postgres password=postgres dbname=ecom sslmode=disable"),
		},
	}

	// structure Logging
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	slog.SetDefault(logger)

	// Database
	conn, err := pgx.Connect(ctx, cfg.db.dsn)
	if err != nil {
		panic(err)
	}
	defer conn.Close(ctx)

	logger.Info("connected to postgres", "dsn", cfg.db.dsn)

	api := application{
		config: cfg,
		db:     conn,
	}

	if err := api.run(api.mount()); err != nil {
		log.Printf("Server Failed to start : %v", err)
		os.Exit(1)
	}
}
