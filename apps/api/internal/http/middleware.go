package http

import (
	"crypto/rand"
	"encoding/hex"
	"net/http"
	"time"

	"github.com/CEckelberry/cole-resume-website/apps/api/internal/logging"
	"log/slog"
)

const requestIDHeader = "X-Request-ID"

// RequestID assigns each incoming request a unique ID. If the client supplies
// X-Request-ID we trust and reuse it; otherwise we generate a 16-hex-char
// token. The ID is set on the response header and stored on the context.
func RequestID(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id := r.Header.Get(requestIDHeader)
		if id == "" {
			id = newRequestID()
		}
		w.Header().Set(requestIDHeader, id)
		ctx := contextWithRequestID(r.Context(), id)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// LogRequests attaches a request-scoped logger to the context (with method,
// path, request_id pre-filled) and emits a structured access log on
// completion, including status code and duration.
func LogRequests(base *slog.Logger) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			start := time.Now()
			id := requestIDFromContext(r.Context())
			scoped := base.With(
				slog.String("request_id", id),
				slog.String("method", r.Method),
				slog.String("path", r.URL.Path),
			)
			ctx := logging.IntoContext(r.Context(), scoped)

			// Wrap the writer so we can capture the status code without
			// pulling in chi's middleware.WrapResponseWriter (keep deps thin).
			ww := &statusRecorder{ResponseWriter: w, status: http.StatusOK}
			next.ServeHTTP(ww, r.WithContext(ctx))

			scoped.LogAttrs(ctx, slog.LevelInfo, "http.request",
				slog.Int("status", ww.status),
				slog.Duration("duration", time.Since(start)),
				slog.String("remote_addr", r.RemoteAddr),
			)
		})
	}
}

type statusRecorder struct {
	http.ResponseWriter
	status int
}

func (s *statusRecorder) WriteHeader(code int) {
	s.status = code
	s.ResponseWriter.WriteHeader(code)
}

func newRequestID() string {
	var buf [8]byte
	if _, err := rand.Read(buf[:]); err != nil {
		return "0000000000000000"
	}
	return hex.EncodeToString(buf[:])
}

// --- internal context plumbing for request id ---

type ridKey int

const ridCtxKey ridKey = 0

func contextWithRequestID(ctx requestCtx, id string) requestCtx {
	return contextWithValue(ctx, ridCtxKey, id)
}

func requestIDFromContext(ctx requestCtx) string {
	if v, ok := ctx.Value(ridCtxKey).(string); ok {
		return v
	}
	return ""
}
