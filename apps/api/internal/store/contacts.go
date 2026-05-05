package store

import (
	"context"
	"fmt"
)

// ContactInsert is the input shape for InsertContact. Pointer-free so the
// handler can pass in a value with no aliasing concerns.
type ContactInsert struct {
	Name      string
	Email     string
	Message   string
	SourceIP  string // optional; empty string skips the inet column
	UserAgent string // optional
}

// InsertContact persists a contact-form submission and returns its UUID.
func (p *Pool) InsertContact(ctx context.Context, in ContactInsert) (string, error) {
	var id string
	// Cast empty string source_ip to NULL via the conditional cast — pgx
	// rejects an empty string against the inet column.
	err := p.QueryRow(ctx, `
		INSERT INTO contact_submissions (name, email, message, source_ip, user_agent)
		VALUES ($1, $2, $3, NULLIF($4, '')::inet, NULLIF($5, ''))
		RETURNING id
	`, in.Name, in.Email, in.Message, in.SourceIP, in.UserAgent).Scan(&id)
	if err != nil {
		return "", fmt.Errorf("insert contact: %w", err)
	}
	return id, nil
}
