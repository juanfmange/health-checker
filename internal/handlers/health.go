package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/juanfmange/health-checker/internal/checker"
	"github.com/juanfmange/health-checker/internal/config"
)

func HealthHandler(cfg config.Config) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		results := checker.CheckServices(r.Context(), cfg.Services)

		response := map[string]interface{}{
			"services":  results,
			"timestamp": time.Now().UTC().Format(time.RFC3339),
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	}
}
