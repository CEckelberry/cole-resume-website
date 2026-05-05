// Package handlers contains the per-endpoint HTTP handlers.
package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/CEckelberry/cole-resume-website/apps/api/internal/config"
)

// HealthResponse is the body returned by GET /api/health.
type HealthResponse struct {
	Status      string `json:"status"`
	Environment string `json:"environment"`
	Version     string `json:"version"`
}

// Health returns an http.HandlerFunc serving the readiness probe.
//
// Cloud Run wires this to its own health check; we keep the handler
// dependency-free at v1 so a failed Postgres connection doesn't take the
// service down. Database-aware health checks come in Phase 3 with the
// contact-form work.
func Health(cfg *config.Config) http.HandlerFunc {
	body := HealthResponse{
		Status:      "ok",
		Environment: string(cfg.Environment),
		Version:     cfg.Version,
	}
	encoded, _ := json.Marshal(body)

	return func(w http.ResponseWriter, _ *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Cache-Control", "no-store")
		_, _ = w.Write(encoded)
	}
}
