// Package http wires the chi router and middleware chain.
package http

import (
	"log/slog"
	"net/http"
	"time"

	"github.com/CEckelberry/cole-resume-website/apps/api/internal/config"
	"github.com/CEckelberry/cole-resume-website/apps/api/internal/http/handlers"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

// NewRouter returns a fully-configured chi router with the standard
// middleware chain and v1 routes registered.
func NewRouter(cfg *config.Config, log *slog.Logger) http.Handler {
	r := chi.NewRouter()

	// Order matters: real IP and request ID need to run before logging so the
	// log line has the correct attributes; recoverer last so a panic in any
	// upstream middleware still surfaces.
	r.Use(middleware.RealIP)
	r.Use(RequestID)
	r.Use(LogRequests(log))
	r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(15 * time.Second))

	r.Route("/api", func(r chi.Router) {
		r.Get("/health", handlers.Health(cfg))
	})

	return r
}
