-- Migration tracking. Hand-rolled because golang-migrate is a dependency
-- and a binary we'd otherwise have to ship in the runtime image; for the
-- handful of migrations this service needs, embedded SQL + a tiny applier
-- in store/postgres.go is enough.
CREATE TABLE IF NOT EXISTS schema_migrations (
    version    text        PRIMARY KEY,
    applied_at timestamptz NOT NULL DEFAULT now()
);
