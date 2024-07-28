package main

import (
	"Torq_IPGeo_Assignment/config"
	"Torq_IPGeo_Assignment/csv"
	"Torq_IPGeo_Assignment/handlers"
	customMiddleware "Torq_IPGeo_Assignment/middleware"
	"Torq_IPGeo_Assignment/models"
	"Torq_IPGeo_Assignment/ratelimit"
	"context"
	"errors"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}
	log.Printf("Configuration loaded successfully: %v", cfg)

	// Externalize the Store creation to a separated func or even Factory function
	var db models.Store
	switch cfg.DatabaseType {
	case "csv":
		file, err := os.Open(cfg.DatabasePath)
		if err != nil {
			log.Fatalf("Failed to open database path: %v", err)
		}
		defer file.Close()
		db, err = csv.NewStore(file)
		if err != nil {
			log.Fatalf("Failed to load database: %v", err)
		}
	default:
		log.Fatalf("Unsupported database type: %v", cfg.DatabaseType)
	}

	log.Printf("Successfully connected to DB Store of type: %v", cfg.DatabaseType)

	// Externalize to a function to init Http Server [readability]
	rateLimiter := ratelimit.NewRateLimiter(cfg.RateRequests, cfg.RateInterval)
	mwRateLimiter := customMiddleware.NewRateLimiter(rateLimiter)
	handler := handlers.Handler{DB: db}

	router := chi.NewRouter()

	// Middleware
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)
	router.Use(middleware.Timeout(60 * time.Second))
	router.Use(middleware.SetHeader("Content-Type", "application/json"))
	router.Use(mwRateLimiter.Handle)

	// v1 API routes
	router.Route("/v1", func(r chi.Router) {
		r.With(customMiddleware.MethodNotAllowedHandler).Get(
			"/find-country", handler.GetCountry)
	})

	// Default handler for unmatched routes
	router.NotFound(func(w http.ResponseWriter, r *http.Request) {
		http.NotFound(w, r) // Default 404 not found response for non-routed paths
	})

	server := &http.Server{
		Addr:         ":" + cfg.Port,
		Handler:      router,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  15 * time.Second,
	}

	// Graceful shutdown
	go func() {
		if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("Could not listen on %s: %v", cfg.Port, err)
		}
	}()
	log.Printf("Server listening on port %s", cfg.Port)

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c

	//ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	//defer cancel()
	//server.Shutdown(ctx)
	server.Shutdown(context.Background())
	log.Println("Shutting down gracefully...")
}
