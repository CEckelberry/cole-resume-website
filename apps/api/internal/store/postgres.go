// Package store wraps the Postgres connection pool and runs embedded
// migrations on startup. Exists so cmd/server doesn't have to know about
// pgx directly.
package store

import (
	"context"
	"embed"
	"fmt"
	"io/fs"
	"sort"
	"strings"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

//go:embed migrations/*.sql
var migrationsFS embed.FS

const migrationsDir = "migrations"

// Pool is the typed wrapper. Currently a thin alias; a future version may
// add metrics, query timeouts, or RLS context plumbing here.
type Pool struct {
	*pgxpool.Pool
}

// Connect creates the pool, pings it, and runs migrations. Returns a Pool
// ready to use. Caller owns Close().
func Connect(ctx context.Context, dsn string) (*Pool, error) {
	cfg, err := pgxpool.ParseConfig(dsn)
	if err != nil {
		return nil, fmt.Errorf("parse dsn: %w", err)
	}

	// Conservative defaults — Cloud SQL has a low ceiling on max conns and
	// the API workload is request-scoped, not bulk.
	cfg.MaxConns = 10
	cfg.MinConns = 0
	cfg.MaxConnIdleTime = 30 * time.Second

	pool, err := pgxpool.NewWithConfig(ctx, cfg)
	if err != nil {
		return nil, fmt.Errorf("create pool: %w", err)
	}

	pingCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	if err := pool.Ping(pingCtx); err != nil {
		pool.Close()
		return nil, fmt.Errorf("ping: %w", err)
	}

	p := &Pool{Pool: pool}
	if err := p.Migrate(ctx); err != nil {
		pool.Close()
		return nil, fmt.Errorf("migrate: %w", err)
	}
	return p, nil
}

// Migrate applies any embedded SQL files in migrations/ that haven't yet
// been recorded in schema_migrations. Versions are derived from the file
// name (e.g. "002_contact_submissions.sql" → "002").
func (p *Pool) Migrate(ctx context.Context) error {
	entries, err := fs.ReadDir(migrationsFS, migrationsDir)
	if err != nil {
		return fmt.Errorf("read embedded migrations: %w", err)
	}

	// Always run the schema_migrations table creation (001) first so we can
	// query it. ReadDir is alphabetical, but be explicit.
	files := make([]string, 0, len(entries))
	for _, e := range entries {
		if !e.IsDir() && strings.HasSuffix(e.Name(), ".sql") {
			files = append(files, e.Name())
		}
	}
	sort.Strings(files)

	for _, name := range files {
		version := strings.SplitN(name, "_", 2)[0]
		applied, err := p.alreadyApplied(ctx, version)
		if err != nil {
			return err
		}
		if applied {
			continue
		}

		body, err := fs.ReadFile(migrationsFS, migrationsDir+"/"+name)
		if err != nil {
			return fmt.Errorf("read %s: %w", name, err)
		}

		tx, err := p.Begin(ctx)
		if err != nil {
			return fmt.Errorf("begin %s: %w", version, err)
		}
		if _, err := tx.Exec(ctx, string(body)); err != nil {
			_ = tx.Rollback(ctx)
			return fmt.Errorf("exec %s: %w", name, err)
		}
		if _, err := tx.Exec(ctx,
			`INSERT INTO schema_migrations (version) VALUES ($1) ON CONFLICT DO NOTHING`,
			version,
		); err != nil {
			_ = tx.Rollback(ctx)
			return fmt.Errorf("record %s: %w", version, err)
		}
		if err := tx.Commit(ctx); err != nil {
			return fmt.Errorf("commit %s: %w", version, err)
		}
	}
	return nil
}

func (p *Pool) alreadyApplied(ctx context.Context, version string) (bool, error) {
	// schema_migrations may not exist yet on a fresh database; the first
	// migration creates it. Treat "missing relation" as "not applied".
	var exists bool
	err := p.QueryRow(ctx,
		`SELECT EXISTS (SELECT 1 FROM schema_migrations WHERE version = $1)`,
		version,
	).Scan(&exists)
	if err != nil {
		// pgx returns the SQLSTATE in the error string; 42P01 is undefined_table.
		if strings.Contains(err.Error(), "42P01") {
			return false, nil
		}
		return false, fmt.Errorf("check version %s: %w", version, err)
	}
	return exists, nil
}
