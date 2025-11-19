package main

import (
	"log"
	"net/http"

	"github.com/juanfmange/health-checker/internal/config"
	"github.com/juanfmange/health-checker/internal/handlers"
)

func main() {
	cfg := config.LoadConfig()

	http.HandleFunc("/health-check", handlers.HealthHandler(cfg))

	log.Printf("Server running on port %s", cfg.Port)
	log.Fatal(http.ListenAndServe(":"+cfg.Port, nil))
}
