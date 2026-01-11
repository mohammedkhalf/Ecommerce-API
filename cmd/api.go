package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/jackc/pgx/v5"
	repo "github.com/mohammedkhalf/Ecommerce-API/internal/adapters/postgresql/sqlc"
	"github.com/mohammedkhalf/Ecommerce-API/internal/orders"
	"github.com/mohammedkhalf/Ecommerce-API/internal/products"
	"log"
	"net/http"
	"time"
)

type application struct {
	config config
	// logger
	db *pgx.Conn
}

// mount
func (app *application) mount() http.Handler {
	r := chi.NewRouter()
	// A good base middleware stack
	r.Use(middleware.RequestID) // Important for rate limiting
	r.Use(middleware.RealIP)    // Important for rate limiting , analytics and tracing
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer) // Recover from crash

	// Set a timeout value on the request context (ctx), that will signal
	// through ctx.Done() that the request has timed out and further
	// processing should be stopped.
	r.Use(middleware.Timeout(60 * time.Second))

	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("all good"))
	})

	//Products
	productService := products.NewService(repo.New(app.db))
	productHandler := products.NewHandler(productService)
	r.Get("/products", productHandler.ListProducts)

	//Orders
	orderService := orders.NewService(repo.New(app.db), app.db)
	ordersHandler := orders.NewHandler(orderService)
	r.Post("/orders", ordersHandler.PlaceOrder)

	return r
}

// run
func (app *application) run(h http.Handler) error {
	srv := &http.Server{
		Addr:         app.config.addr,
		Handler:      h,
		WriteTimeout: time.Second * 30,
		ReadTimeout:  time.Second * 10,
		IdleTimeout:  time.Minute,
	}

	log.Printf("starting server on %s", app.config.addr)

	return srv.ListenAndServe()
}

type config struct {
	addr string
	db   dbConfig
}

// Database Config
type dbConfig struct {
	dsn string
}
