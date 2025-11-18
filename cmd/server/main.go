package main

import (
	"log"
	"net/http"

	"github.com/juanfmange/health-checker/internal/config"
	handlers "github.com/juanfmange/health-checker/internal/handler"
)

func main() {
	cfg := config.LoadConfig()

	http.HandleFunc("/health", handlers.HealthHandler(cfg))

	log.Printf("Server running on port %s", cfg.Port)
	log.Fatal(http.ListenAndServe(":"+cfg.Port, nil))
}
