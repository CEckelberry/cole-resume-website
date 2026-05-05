package handlers

import (
	"context"
	"encoding/json"
	"net/http"
	"net/mail"
	"strings"

	"github.com/CEckelberry/cole-resume-website/apps/api/internal/logging"
	"github.com/CEckelberry/cole-resume-website/apps/api/internal/store"
)

// ContactStore is the persistence dependency. Defined as an interface so
// tests can pass a fake without standing up Postgres.
type ContactStore interface {
	InsertContact(ctx context.Context, in store.ContactInsert) (string, error)
}

type contactRequest struct {
	Name    string `json:"name"`
	Email   string `json:"email"`
	Message string `json:"message"`
}

// Contact returns the POST /api/contact handler.
//
// Validation is server-authoritative — the frontend's checks are for UX,
// not enforcement. Rate limiting lives in middleware so this handler stays
// focused on the happy path.
func Contact(st ContactStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Cap the body so a hostile client can't tie up memory by streaming
		// gigabytes; 32 KiB is comfortably more than 4000 chars + headers.
		r.Body = http.MaxBytesReader(w, r.Body, 32<<10)

		var req contactRequest
		dec := json.NewDecoder(r.Body)
		dec.DisallowUnknownFields()
		if err := dec.Decode(&req); err != nil {
			respondErr(w, http.StatusBadRequest, "invalid request body")
			return
		}

		req.Name = strings.TrimSpace(req.Name)
		req.Email = strings.TrimSpace(req.Email)
		req.Message = strings.TrimSpace(req.Message)

		if err := validateContact(req); err != "" {
			respondErr(w, http.StatusBadRequest, err)
			return
		}

		id, err := st.InsertContact(r.Context(), store.ContactInsert{
			Name:      req.Name,
			Email:     req.Email,
			Message:   req.Message,
			SourceIP:  r.Header.Get("X-Forwarded-For"),
			UserAgent: r.Header.Get("User-Agent"),
		})
		if err != nil {
			logging.FromContext(r.Context()).Error("contact insert failed", "err", err)
			respondErr(w, http.StatusInternalServerError, "could not save submission")
			return
		}

		logging.FromContext(r.Context()).Info("contact submission received", "id", id, "email", req.Email)
		respondJSON(w, http.StatusCreated, map[string]any{"status": "ok", "id": id})
	}
}

func validateContact(req contactRequest) string {
	if req.Name == "" {
		return "name is required"
	}
	if len(req.Name) > 120 {
		return "name is too long (max 120)"
	}
	if req.Email == "" {
		return "email is required"
	}
	if _, err := mail.ParseAddress(req.Email); err != nil {
		return "email is invalid"
	}
	if len(req.Email) > 254 {
		return "email is too long (max 254)"
	}
	if req.Message == "" {
		return "message is required"
	}
	if len(req.Message) > 4000 {
		return "message is too long (max 4000)"
	}
	return ""
}

func respondJSON(w http.ResponseWriter, status int, body any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(body)
}

func respondErr(w http.ResponseWriter, status int, msg string) {
	respondJSON(w, status, map[string]string{"error": msg})
}
