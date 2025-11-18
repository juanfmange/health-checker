package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/juanfmange/health-checker/internal/checker"
	"github.com/juanfmange/health-checker/internal/config"
)

func HealthHandler(cfg config.Config) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		results := checker.CheckServices(r.Context(), cfg.Services, cfg.TimeoutSeconds)

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]interface{}{
			"status":   "ok",
			"services": results,
		})
	}
}
