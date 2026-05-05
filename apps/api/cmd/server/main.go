// Command server starts the portfolio API HTTP service.
package main

import (
	"context"
	"errors"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/CEckelberry/cole-resume-website/apps/api/internal/config"
	apphttp "github.com/CEckelberry/cole-resume-website/apps/api/internal/http"
	"github.com/CEckelberry/cole-resume-website/apps/api/internal/logging"
)

const (
	readHeaderTimeout = 5 * time.Second
	writeTimeout      = 30 * time.Second
	idleTimeout       = 60 * time.Second
	shutdownTimeout   = 10 * time.Second
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		// No logger yet — emit a final-resort message and bail.
		slog.New(slog.NewTextHandler(os.Stderr, nil)).
			Error("config load failed", slog.Any("error", err))
		os.Exit(1)
	}

	log := logging.New(cfg.IsProduction(), cfg.LogLevel)
	slog.SetDefault(log)

	log.Info("server starting",
		slog.String("environment", string(cfg.Environment)),
		slog.String("port", cfg.Port),
		slog.String("version", cfg.Version),
		slog.String("build_sha", cfg.BuildSHA),
	)

	srv := &http.Server{
		Addr:              ":" + cfg.Port,
		Handler:           apphttp.NewRouter(cfg, log),
		ReadHeaderTimeout: readHeaderTimeout,
		WriteTimeout:      writeTimeout,
		IdleTimeout:       idleTimeout,
	}

	// Run the server in a goroutine so we can block on signal handling here.
	errCh := make(chan error, 1)
	go func() {
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			errCh <- err
		}
		close(errCh)
	}()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)

	select {
	case sig := <-stop:
		log.Info("shutdown signal received", slog.String("signal", sig.String()))
	case err := <-errCh:
		if err != nil {
			log.Error("server failed", slog.Any("error", err))
			os.Exit(1)
		}
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), shutdownTimeout)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Error("graceful shutdown failed", slog.Any("error", err))
		os.Exit(1)
	}

	log.Info("server stopped")
}
