package main

import (
	"log"
	"net/http"

	"github.com/juanfmange/health-checker/internal/config"
	"github.com/juanfmange/health-checker/internal/handlers"
	"github.com/juanfmange/health-checker/internal/middleware"
)

func main() {
	cfg := config.LoadConfig()

	mux := http.NewServeMux()
	mux.HandleFunc("/health", handlers.HealthHandler(cfg))

	// Wrap your mux with CORS middleware
	handlerWithMiddleware := middleware.CORS(mux)

	log.Printf("Server running on port %s", cfg.Port)
	log.Fatal(http.ListenAndServe(":"+cfg.Port, handlerWithMiddleware))
}
