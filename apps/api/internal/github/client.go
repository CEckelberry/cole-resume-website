// Package github wraps the public GitHub events API for the user behind
// the portfolio site. The data feeds the about-section activity card.
//
// Rate limits: unauthenticated is 60 req/hour per IP. With a 5-minute
// cache the client hits GitHub 12×/hour, well inside the budget. If a
// GITHUB_TOKEN is supplied (5000 req/hour authenticated), it's used.
package github

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"sync"
	"time"
)

const (
	apiBase  = "https://api.github.com"
	cacheTTL = 5 * time.Minute
	maxItems = 8
)

// ActivityItem mirrors the JSON contract documented in ARCHITECTURE.md.
type ActivityItem struct {
	Kind    string    `json:"kind"`     // "commit" | "release" | "create" | "pr"
	Repo    string    `json:"repo"`     // owner/name
	Title   string    `json:"title"`    // commit message, release name, etc.
	URL     string    `json:"url"`      // link to the commit / release
	AgeText string    `json:"age_text"` // pre-formatted "2h ago" for client convenience
	At      time.Time `json:"at"`
}

// Client fetches and caches the user's recent public GitHub events.
//
// Concurrency: a single mutex guards the cache. Refreshes are serialized;
// the second concurrent caller after expiry waits for the first.
type Client struct {
	httpClient *http.Client
	username   string
	token      string

	mu         sync.Mutex
	cache      []ActivityItem
	cacheUntil time.Time
}

// New returns a Client for `username`. `token` is optional — if empty,
// requests go out unauthenticated.
func New(username, token string) *Client {
	return &Client{
		httpClient: &http.Client{Timeout: 6 * time.Second},
		username:   username,
		token:      token,
	}
}

// Activity returns up to maxItems recent activity items, cached for cacheTTL.
// On fetch error, returns the previous cached value if any (stale-while-error).
func (c *Client) Activity(ctx context.Context) ([]ActivityItem, error) {
	c.mu.Lock()
	defer c.mu.Unlock()

	if time.Now().Before(c.cacheUntil) && len(c.cache) > 0 {
		return c.cache, nil
	}

	fresh, err := c.fetch(ctx)
	if err != nil {
		// Stale-while-error so a transient GitHub blip doesn't blank the UI.
		if len(c.cache) > 0 {
			return c.cache, nil
		}
		return nil, err
	}

	c.cache = fresh
	c.cacheUntil = time.Now().Add(cacheTTL)
	return fresh, nil
}

func (c *Client) fetch(ctx context.Context) ([]ActivityItem, error) {
	url := fmt.Sprintf("%s/users/%s/events/public?per_page=30", apiBase, c.username)
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, fmt.Errorf("build request: %w", err)
	}
	req.Header.Set("Accept", "application/vnd.github+json")
	req.Header.Set("X-GitHub-Api-Version", "2022-11-28")
	req.Header.Set("User-Agent", "cole-eckelberry-portfolio")
	if c.token != "" {
		req.Header.Set("Authorization", "Bearer "+c.token)
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("github request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("github status %d", resp.StatusCode)
	}

	var events []githubEvent
	if err := json.NewDecoder(resp.Body).Decode(&events); err != nil {
		return nil, fmt.Errorf("decode events: %w", err)
	}

	items := make([]ActivityItem, 0, maxItems)
	for _, e := range events {
		if it, ok := mapEvent(e); ok {
			items = append(items, it)
			if len(items) >= maxItems {
				break
			}
		}
	}
	return items, nil
}

type githubEvent struct {
	Type      string          `json:"type"`
	Repo      githubEventRepo `json:"repo"`
	Payload   json.RawMessage `json:"payload"`
	CreatedAt time.Time       `json:"created_at"`
}

type githubEventRepo struct {
	Name string `json:"name"`
}

// mapEvent picks the interesting event types and maps them to the
// ActivityItem shape. Returns ok=false for events the UI shouldn't surface
// (issue comments, watch events, etc.).
func mapEvent(e githubEvent) (ActivityItem, bool) {
	repoBase := "https://github.com/" + e.Repo.Name

	switch e.Type {
	case "PushEvent":
		var p struct {
			Commits []struct {
				Sha     string `json:"sha"`
				Message string `json:"message"`
			} `json:"commits"`
		}
		_ = json.Unmarshal(e.Payload, &p)
		if len(p.Commits) == 0 {
			return ActivityItem{}, false
		}
		last := p.Commits[len(p.Commits)-1]
		title := strings.SplitN(last.Message, "\n", 2)[0]
		return ActivityItem{
			Kind:    "commit",
			Repo:    e.Repo.Name,
			Title:   title,
			URL:     repoBase + "/commit/" + last.Sha,
			AgeText: ageText(e.CreatedAt, time.Now()),
			At:      e.CreatedAt,
		}, true

	case "ReleaseEvent":
		var p struct {
			Action  string `json:"action"`
			Release struct {
				Name    string `json:"name"`
				TagName string `json:"tag_name"`
				HTMLURL string `json:"html_url"`
			} `json:"release"`
		}
		_ = json.Unmarshal(e.Payload, &p)
		if p.Action != "published" {
			return ActivityItem{}, false
		}
		title := p.Release.Name
		if title == "" {
			title = p.Release.TagName
		}
		return ActivityItem{
			Kind:    "release",
			Repo:    e.Repo.Name,
			Title:   "released " + title,
			URL:     p.Release.HTMLURL,
			AgeText: ageText(e.CreatedAt, time.Now()),
			At:      e.CreatedAt,
		}, true

	case "CreateEvent":
		var p struct {
			RefType string `json:"ref_type"` // "repository" | "branch" | "tag"
			Ref     string `json:"ref"`
		}
		_ = json.Unmarshal(e.Payload, &p)
		if p.RefType != "repository" {
			return ActivityItem{}, false
		}
		return ActivityItem{
			Kind:    "create",
			Repo:    e.Repo.Name,
			Title:   "started a new repo",
			URL:     repoBase,
			AgeText: ageText(e.CreatedAt, time.Now()),
			At:      e.CreatedAt,
		}, true

	case "PullRequestEvent":
		var p struct {
			Action      string `json:"action"`
			PullRequest struct {
				Title   string `json:"title"`
				HTMLURL string `json:"html_url"`
			} `json:"pull_request"`
		}
		_ = json.Unmarshal(e.Payload, &p)
		if p.Action != "opened" && p.Action != "closed" {
			return ActivityItem{}, false
		}
		prefix := "opened PR"
		if p.Action == "closed" {
			prefix = "merged PR"
		}
		return ActivityItem{
			Kind:    "pr",
			Repo:    e.Repo.Name,
			Title:   prefix + ": " + p.PullRequest.Title,
			URL:     p.PullRequest.HTMLURL,
			AgeText: ageText(e.CreatedAt, time.Now()),
			At:      e.CreatedAt,
		}, true
	}

	return ActivityItem{}, false
}

func ageText(t time.Time, now time.Time) string {
	d := now.Sub(t)
	switch {
	case d < time.Minute:
		return "just now"
	case d < time.Hour:
		return fmt.Sprintf("%dm ago", int(d/time.Minute))
	case d < 24*time.Hour:
		return fmt.Sprintf("%dh ago", int(d/time.Hour))
	default:
		return fmt.Sprintf("%dd ago", int(d/(24*time.Hour)))
	}
}
