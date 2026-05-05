package http

import (
	"net"
	"net/http"
	"sync"
	"time"

	"golang.org/x/time/rate"
)

// PerIPLimiter throttles requests per source IP using a token bucket.
//
// Defaults at v1: 5 submissions per hour, burst 5. That's enough that a
// human can resubmit after a typo without waiting, but a script trying to
// brute-spam the form gets blocked after 5 in a row.
//
// Stale entries are GC'd after 6 hours of inactivity so the map doesn't
// grow unboundedly under sustained traffic.
type PerIPLimiter struct {
	mu       sync.Mutex
	limiters map[string]*ipState
	r        rate.Limit
	burst    int
	gcAfter  time.Duration
}

type ipState struct {
	limiter  *rate.Limiter
	lastSeen time.Time
}

// NewPerIPLimiter returns a limiter allowing `burst` requests immediately
// and `r` requests per second sustained.
func NewPerIPLimiter(r rate.Limit, burst int) *PerIPLimiter {
	l := &PerIPLimiter{
		limiters: make(map[string]*ipState),
		r:        r,
		burst:    burst,
		gcAfter:  6 * time.Hour,
	}
	go l.gcLoop()
	return l
}

// ContactSubmissionLimiter returns the limiter sized for contact-form
// submissions: 5 per hour with burst 5.
func ContactSubmissionLimiter() *PerIPLimiter {
	return NewPerIPLimiter(rate.Every(time.Hour/5), 5)
}

// Middleware returns http middleware that 429s when an IP exceeds the rate.
func (l *PerIPLimiter) Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ip := clientIP(r)
		if !l.allow(ip) {
			w.Header().Set("Retry-After", "3600")
			http.Error(w, `{"error":"rate limit: too many submissions from this IP"}`, http.StatusTooManyRequests)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func (l *PerIPLimiter) allow(ip string) bool {
	l.mu.Lock()
	defer l.mu.Unlock()
	st, ok := l.limiters[ip]
	if !ok {
		st = &ipState{limiter: rate.NewLimiter(l.r, l.burst)}
		l.limiters[ip] = st
	}
	st.lastSeen = time.Now()
	return st.limiter.Allow()
}

func (l *PerIPLimiter) gcLoop() {
	t := time.NewTicker(15 * time.Minute)
	defer t.Stop()
	for now := range t.C {
		l.mu.Lock()
		for ip, st := range l.limiters {
			if now.Sub(st.lastSeen) > l.gcAfter {
				delete(l.limiters, ip)
			}
		}
		l.mu.Unlock()
	}
}

// clientIP picks the most-trusted source-IP signal available. The
// SvelteKit proxy sets X-Forwarded-For; chi's RealIP middleware also
// rewrites RemoteAddr. The xff parse handles the comma-list shape.
func clientIP(r *http.Request) string {
	if xff := r.Header.Get("X-Forwarded-For"); xff != "" {
		// First entry is the original client.
		for i := 0; i < len(xff); i++ {
			if xff[i] == ',' {
				return trimSpace(xff[:i])
			}
		}
		return trimSpace(xff)
	}
	host, _, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		return r.RemoteAddr
	}
	return host
}

func trimSpace(s string) string {
	for len(s) > 0 && (s[0] == ' ' || s[0] == '\t') {
		s = s[1:]
	}
	for len(s) > 0 && (s[len(s)-1] == ' ' || s[len(s)-1] == '\t') {
		s = s[:len(s)-1]
	}
	return s
}
