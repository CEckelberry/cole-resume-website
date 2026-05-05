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
//
// `contactStore` may be nil — happens when the service starts without a
// DATABASE_URL (e.g. local dev with no docker compose). In that case
// /api/contact returns 503 instead of nil-panicking; everything else
// (health, future read endpoints) keeps working.
func NewRouter(cfg *config.Config, log *slog.Logger, contactStore handlers.ContactStore) http.Handler {
	r := chi.NewRouter()

	r.Use(middleware.RealIP)
	r.Use(RequestID)
	r.Use(LogRequests(log))
	r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(15 * time.Second))

	contactRL := ContactSubmissionLimiter()

	r.Route("/api", func(r chi.Router) {
		r.Get("/health", handlers.Health(cfg))

		r.With(contactRL.Middleware).Post("/contact", contactHandler(contactStore))
	})

	return r
}

// contactHandler wraps handlers.Contact so we can return 503 cleanly when
// the store wasn't configured.
func contactHandler(st handlers.ContactStore) http.HandlerFunc {
	if st == nil {
		return func(w http.ResponseWriter, _ *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusServiceUnavailable)
			_, _ = w.Write([]byte(`{"error":"contact form unavailable: no database configured"}`))
		}
	}
	return handlers.Contact(st)
}
