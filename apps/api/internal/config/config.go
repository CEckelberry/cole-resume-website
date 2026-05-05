// Package config loads runtime settings from environment variables.
//
// Cloud Run injects PORT; everything else has sensible defaults so a bare
// `go run ./cmd/server` works locally without a .env file. For local dev
// with a Postgres container or secrets, drop a .env in apps/api and it will
// be loaded automatically.
package config

import (
	"fmt"
	"log/slog"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

type Environment string

const (
	EnvLocal      Environment = "local"
	EnvStaging    Environment = "staging"
	EnvProduction Environment = "production"
)

type Config struct {
	Port               string
	DatabaseURL        string
	LogLevel           slog.Level
	Environment        Environment
	GitHubToken        string
	DiscordWebhookURL  string
	BuildSHA           string
	Version            string
}

// Load reads env vars (and an optional .env file in CWD) and returns a Config.
// Returns an error if required values are missing for non-local environments.
func Load() (*Config, error) {
	// Load .env if present; ignore "file not found" — it's optional in prod
	// (Cloud Run / Docker / k8s supply env directly).
	_ = godotenv.Load()

	env := Environment(getenv("ENVIRONMENT", string(EnvLocal)))

	cfg := &Config{
		Port:              getenv("PORT", "8080"),
		DatabaseURL:       os.Getenv("DATABASE_URL"),
		LogLevel:          parseLogLevel(getenv("LOG_LEVEL", defaultLogLevel(env))),
		Environment:       env,
		GitHubToken:       os.Getenv("GITHUB_TOKEN"),
		DiscordWebhookURL: os.Getenv("DISCORD_WEBHOOK_URL"),
		BuildSHA:          getenv("BUILD_SHA", "dev"),
		Version:           getenv("VERSION", "dev"),
	}

	if env == EnvProduction && cfg.DatabaseURL == "" {
		return nil, fmt.Errorf("DATABASE_URL is required when ENVIRONMENT=production")
	}

	return cfg, nil
}

// IsProduction reports whether the service is running in production mode.
// Used to switch logger format and other prod-only behavior.
func (c *Config) IsProduction() bool {
	return c.Environment == EnvProduction
}

func getenv(key, fallback string) string {
	if v, ok := os.LookupEnv(key); ok && v != "" {
		return v
	}
	return fallback
}

func parseLogLevel(s string) slog.Level {
	switch strings.ToLower(s) {
	case "debug":
		return slog.LevelDebug
	case "info":
		return slog.LevelInfo
	case "warn", "warning":
		return slog.LevelWarn
	case "error":
		return slog.LevelError
	default:
		return slog.LevelInfo
	}
}

func defaultLogLevel(env Environment) string {
	if env == EnvLocal {
		return "debug"
	}
	return "info"
}
