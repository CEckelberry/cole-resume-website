package handlers

import (
	"net/http"

	"github.com/CEckelberry/cole-resume-website/apps/api/internal/github"
	"github.com/CEckelberry/cole-resume-website/apps/api/internal/logging"
)

// Activity returns the GET /api/activity handler.
//
// Always returns 200 with a possibly-empty items list — the about section
// renders a sensible empty state rather than treating "no recent commits"
// as a 500. If the GitHub API is unreachable the response carries a
// `degraded: true` flag so the client can show a "couldn't reach github"
// pill instead of empty rows.
func Activity(client *github.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		items, err := client.Activity(r.Context())

		if err != nil {
			logging.FromContext(r.Context()).Warn("github activity fetch failed", "err", err)
			respondJSON(w, http.StatusOK, map[string]any{
				"items":    []github.ActivityItem{},
				"degraded": true,
				"error":    err.Error(),
			})
			return
		}

		w.Header().Set("Cache-Control", "public, max-age=60")
		respondJSON(w, http.StatusOK, map[string]any{
			"items":    items,
			"degraded": false,
		})
	}
}
