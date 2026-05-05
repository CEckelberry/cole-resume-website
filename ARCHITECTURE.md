# Architecture

This document describes how the portfolio site is built, deployed, and operated. It is the source of truth for structural decisions. Visual choices live in `DESIGN.md`; the build sequence lives in `TASKS.md`.

## Goals and non-goals

**Goals.** A fast, accessible, beautiful single-domain portfolio that the owner can iterate on weekly. Cheap to run (target: under $20/month at zero traffic, scales to thousands of visitors per day without intervention). Trivially redeployable from a single `git push`. Each piece of the site (hero, gallery, case study, about) loads in under 1 second on a cold cache from a US East/West coast viewer.

**Non-goals.** A CMS. A user-account system. Real-time multiplayer. Mobile-first only — desktop and mobile both matter, with the desktop experience leading because that is where recruiters and senior engineers tend to evaluate portfolios.

## Service topology

Two services, both running on Google Cloud Run, both behind a single domain:

```
                          ┌──────────────────────────┐
                          │   ce.dev                 │
                          │   (Cloud DNS + LB)       │
                          └────────────┬─────────────┘
                                       │
                     ┌─────────────────┴─────────────────┐
                     │                                   │
                     ▼                                   ▼
        ┌────────────────────────┐         ┌──────────────────────────┐
        │  web  (SvelteKit)      │         │  api  (Go)               │
        │  Cloud Run, public     │         │  Cloud Run, public       │
        │  Renders all pages     │         │  /api/* routes only      │
        │  Reads markdown at     │         │  Talks to Cloud SQL      │
        │  build time            │         │  Talks to GitHub API     │
        └────────────────────────┘         └──────────┬───────────────┘
                                                      │
                                                      ▼
                                          ┌──────────────────────────┐
                                          │  Cloud SQL Postgres      │
                                          │  - contact_submissions   │
                                          │  - page_views (optional) │
                                          └──────────────────────────┘
```

The web service is the thing visitors load. The api service is a small JSON backend the web service calls for the handful of dynamic features (contact form submission, GitHub activity feed, page-view counter).

The split exists because Go is the right language for the small backend (cheap, fast cold starts, easy to reason about) and SvelteKit is the right framework for the frontend. They run as independent Cloud Run services so they can scale and deploy independently. A single Cloud Load Balancer routes `/api/*` to the Go service and everything else to SvelteKit.

## Why Cloud Run for both

- Both services scale to zero — no cost when nobody is on the site
- Both services scale up automatically under load
- Cold starts on Cloud Run gen2 are sub-second for both Go and SvelteKit
- Same deploy mechanism for both (container → Artifact Registry → Cloud Run)
- Owner has deep operational experience with Cloud Run from EPEX SPOT
- Domain mapping and SSL certificate management is built in
- Logs, metrics, and traces flow into Cloud Logging / Monitoring / Trace by default

The alternative (Vercel for web, Fly.io for Go) would be marginally easier for SvelteKit specifically but introduces two vendors, two billing accounts, and two CI flows. Cloud Run wins on operational coherence.

## Frontend architecture (SvelteKit)

### Rendering strategy

The site uses **prerendering** for everything that doesn't need to be dynamic, which is most of the site. Hero, projects gallery, about, footer, and all four case studies are statically prerendered at build time and served from Cloud Run as static HTML. The result is a Lighthouse score in the high 90s and TTFB under 100ms once the container is warm.

The handful of dynamic surfaces are SvelteKit endpoints that call the Go api:
- `/api/contact` — form submission
- `/api/activity` — GitHub recent commits/deploys (cached 5 minutes)

### Svelte 5 with runes

The project uses Svelte 5 (`$state`, `$derived`, `$effect`). Component state should be runes-based throughout. No legacy reactive declarations.

### Project structure

```
apps/web/
├── src/
│   ├── routes/
│   │   ├── +layout.svelte          ← global shell, theme, mesh background
│   │   ├── +layout.server.ts       ← global data (active theme, etc.)
│   │   ├── +page.svelte            ← hero + projects gallery + about + footer
│   │   ├── work/
│   │   │   ├── +page.svelte        ← projects index (also accessible from home)
│   │   │   └── [slug]/
│   │   │       ├── +page.svelte    ← case study renderer
│   │   │       └── +page.ts        ← loads markdown for [slug]
│   │   ├── about/
│   │   │   └── +page.svelte        ← extended about page
│   │   ├── api/
│   │   │   └── activity/+server.ts ← proxies to Go api with caching
│   │   └── cv/
│   │       └── +server.ts          ← serves static PDF
│   ├── lib/
│   │   ├── components/
│   │   │   ├── shell/              ← Nav, Footer, Layout chrome
│   │   │   ├── hero/               ← Hero, Terminal, StackMarquee
│   │   │   ├── projects/           ← ProjectCard, ProjectGallery, FilterChips
│   │   │   ├── case-study/         ← CSHero, CSSection, CodeBlock, ArchDiagram
│   │   │   ├── effects/            ← MeshBackground, ScrollReveal, MagneticButton
│   │   │   └── ui/                 ← Button, Pill, Badge, Tag (primitives)
│   │   ├── content/                ← markdown loader, project metadata
│   │   ├── theme/                  ← theme tokens, dark/light toggle, palette swap
│   │   ├── motion/                 ← reusable animation utilities (svelte/motion wrappers)
│   │   └── utils/                  ← formatters, date helpers, hash, debounce
│   ├── app.html                    ← <html> shell, font preloads, theme bootstrap
│   ├── app.css                     ← Tailwind imports, design tokens (CSS vars)
│   └── hooks.server.ts             ← request logging, no-op for v1
├── static/                         ← favicons, og images, cv.pdf
├── svelte.config.js
├── vite.config.ts
├── tailwind.config.ts              ← extends with design tokens
├── tsconfig.json
└── package.json
```

### Routing rules

- `/` is the marketing page: hero + projects gallery + about + footer all on one scrolling page
- `/work` is an alias to `/#work` that scrolls to the gallery section, useful for direct linking
- `/work/[slug]` is each individual case study (`/work/bakeoff`, `/work/linux`, `/work/diamond`, `/work/terraplane`)
- `/about` is the extended about page (longer than the home-page summary)
- `/cv` returns the static PDF
- `/api/*` routes are SvelteKit server endpoints that proxy to the Go service

### Content strategy

Case studies live in `packages/content/case-studies/*.md`. Each markdown file has frontmatter (title, slug, summary, role, stack, shipped, status) and a body with custom directives for embedded components like live demos, architecture diagrams, and annotated code blocks.

The markdown is parsed at build time using `mdsvex` so case-study pages are statically rendered. The build inlines the project metadata into each page so there is no runtime fetch needed for the gallery either.

### Theming

CSS custom properties at the `:root` level define the palette. Two attributes control which set of variables is active:
- `data-theme="cotton-candy"` (only theme at v1; structure leaves room for `midnight`, `sunset`, etc. in v2)
- `data-mode="light"` or `data-mode="dark"`

Theme + mode are persisted in `localStorage` and read on first paint via a small inline script in `app.html` to avoid flash of wrong theme. The user preference is also reflected in the URL via `?theme=` and `?mode=` for shareability (v2 — at v1 only the toggle is wired up).

### Performance budget

| Page | Target LCP | Target INP | Target CLS | Bundle size (JS, gzipped) |
|---|---|---|---|---|
| Home | < 1.2s | < 200ms | 0 | < 60kb |
| Case study | < 1.5s | < 200ms | 0 | < 80kb |
| About | < 1.0s | < 200ms | 0 | < 40kb |

Strategies: prerender everything possible, avoid runtime fetches on first paint, preload critical fonts, use system font fallbacks, lazy-load anything below the fold including the mesh background animation.

## Backend architecture (Go)

### Why Go for the API

The owner is six months into Go and is using this project as a sustained Go workout. The API surface is small (two real endpoints) so the language choice has more pedagogical value than performance value. Go is also the right idiom for a service that talks to GCP APIs, since the Google client libraries are first-class.

### Service shape

```
apps/api/
├── cmd/
│   └── server/
│       └── main.go               ← entrypoint, wires deps, starts http.Server
├── internal/
│   ├── http/
│   │   ├── router.go             ← chi setup, middleware chain
│   │   ├── middleware.go         ← logging, request ID, panic recovery, CORS
│   │   └── handlers/
│   │       ├── contact.go        ← POST /api/contact
│   │       ├── activity.go       ← GET  /api/activity (GitHub passthrough)
│   │       └── health.go         ← GET  /api/health
│   ├── store/
│   │   ├── postgres.go           ← *pgxpool.Pool, migrations
│   │   └── contacts.go           ← CRUD on contact_submissions
│   ├── github/
│   │   └── client.go             ← thin wrapper around github.com/google/go-github
│   ├── config/
│   │   └── config.go             ← env-var loading (database URL, port, github token)
│   └── logging/
│       └── logger.go             ← structured logging with slog
├── migrations/
│   ├── 001_init.sql
│   └── 002_contact_submissions.sql
├── Dockerfile
├── go.mod
└── go.sum
```

### Endpoint contracts

#### POST /api/contact

Accepts a JSON body with the visitor's name, email, and message. Validates server-side, rate-limits by IP (5 per hour), inserts into Postgres, and returns 201 on success. Sends a Discord webhook (or similar) to notify the owner. No email confirmation to the sender at v1.

```go
type ContactRequest struct {
    Name    string `json:"name"    validate:"required,max=120"`
    Email   string `json:"email"   validate:"required,email"`
    Message string `json:"message" validate:"required,max=4000"`
}
```

#### GET /api/activity

Returns the most recent commits across the owner's public GitHub repos plus the most recent deploy of the portfolio itself. Cached in-memory for 5 minutes to stay well under GitHub's unauthenticated rate limit. Falls back to a static "no recent activity" payload if the GitHub API is unreachable.

```go
type Activity struct {
    Items []ActivityItem `json:"items"`
    UpdatedAt time.Time  `json:"updated_at"`
}

type ActivityItem struct {
    Kind    string    `json:"kind"`        // "commit", "deploy", "release"
    Repo    string    `json:"repo"`        // "ce/portfolio"
    Title   string    `json:"title"`       // "refactored router pool"
    URL     string    `json:"url"`
    AgeText string    `json:"age_text"`    // "2h ago"
    At      time.Time `json:"at"`
}
```

#### GET /api/health

Returns 200 if the database is reachable. Used by Cloud Run's health check.

### Database schema

```sql
CREATE TABLE contact_submissions (
    id           uuid PRIMARY KEY DEFAULT gen_random_uuid(),
    created_at   timestamptz NOT NULL DEFAULT now(),
    name         text        NOT NULL,
    email        text        NOT NULL,
    message      text        NOT NULL,
    source_ip    inet,
    user_agent   text,
    handled_at   timestamptz
);

CREATE INDEX idx_contact_submissions_created_at ON contact_submissions (created_at DESC);
```

Migrations are managed with `golang-migrate` and run in the container entrypoint before the HTTP server starts.

### Configuration

All config comes from environment variables. No config files. The Cloud Run service definition supplies:

- `PORT` — provided by Cloud Run, the server binds to this
- `DATABASE_URL` — full Postgres connection string, sourced from Secret Manager
- `GITHUB_TOKEN` — read-only PAT for the activity endpoint, sourced from Secret Manager
- `DISCORD_WEBHOOK_URL` — optional, for contact-form notifications
- `LOG_LEVEL` — `info` in prod, `debug` in dev
- `ENVIRONMENT` — `production`, `staging`, or `local`

## Deployment

### Container builds

Both services build to small, distroless containers:

- **web**: multi-stage Dockerfile. Stage 1 runs `npm ci && npm run build` to produce the SvelteKit Node adapter output. Stage 2 is `gcr.io/distroless/nodejs20-debian12` running `node build/index.js`.
- **api**: multi-stage Dockerfile. Stage 1 runs `go build -ldflags='-s -w' -o /server ./cmd/server`. Stage 2 is `gcr.io/distroless/static-debian12` running the binary directly.

Both images are pushed to Artifact Registry on every push to main.

### CI/CD pipeline

GitHub Actions workflow on push to main:

1. **Lint and test** — `eslint`, `prettier --check`, `svelte-check`, `vitest run` for web; `go vet`, `golangci-lint`, `go test` for api
2. **Build images** — both Dockerfiles, tagged with the commit SHA
3. **Push to Artifact Registry** — `gcloud auth` via Workload Identity Federation, no static keys
4. **Run migrations** — `migrate up` against Cloud SQL via private connection
5. **Deploy to Cloud Run** — both services, blue-green via Cloud Run revisions; only the new revision receives traffic if its health check passes
6. **Smoke test** — hit `/` and `/api/health` against the new revision; if either fails, revert traffic

Failures at any step abort the deploy and notify via Slack/Discord.

### Infrastructure as code

Terraform manages everything except secrets:

```
infra/terraform/
├── main.tf                    ← provider, backend (gcs)
├── network.tf                 ← VPC, subnets, serverless VPC connector
├── cloud_run.tf               ← both services, IAM, domain mapping
├── cloud_sql.tf               ← Postgres instance, database, user
├── artifact_registry.tf       ← container repo
├── dns.tf                     ← managed zone, A/AAAA records
├── monitoring.tf              ← uptime checks, alert policies
├── iam.tf                     ← service accounts, WIF for GitHub Actions
└── variables.tf
```

Secrets are created manually via `gcloud secrets create` and granted to Cloud Run service accounts via Terraform IAM bindings.

## Observability

- **Logs**: stdout from both services flows into Cloud Logging automatically. Structured JSON via `slog` (Go) and `pino` (Node).
- **Metrics**: Cloud Run emits request counts, latencies, instance counts, and CPU/memory usage out of the box. No additional metrics framework needed at v1.
- **Traces**: skip at v1. Add OpenTelemetry in v2 if the api grows.
- **Alerts**: one uptime check on `https://ce.dev/`, one on `https://ce.dev/api/health`. Alerts fire to email at v1.

## Security

- All secrets in Secret Manager, never in code or env files committed to git
- Cloud SQL is private-IP only; api connects via the serverless VPC connector
- web has no database access; only api does
- Contact form has rate limiting (5/hour/IP) and basic input validation; no captcha at v1
- HTTPS enforced via Cloud Run's managed certificate
- HSTS, CSP, X-Content-Type-Options, Referrer-Policy headers set by SvelteKit middleware

## Cost model

Approximate monthly cost at zero/low traffic:

| Resource | Monthly |
|---|---|
| Cloud Run (web) | $0 — scales to zero |
| Cloud Run (api) | $0 — scales to zero |
| Cloud SQL (db-f1-micro, ZA) | ~$8 |
| Artifact Registry storage | ~$1 |
| Cloud DNS | ~$0.50 |
| Cloud Build | $0 within free tier |
| **Total** | **~$10/month** |

At 10k visitors/day the cost rises to roughly $20–25/month, dominated by Cloud Run CPU-seconds.

## Future considerations (v2 and beyond)

- A `/journal` route with markdown blog posts, same content pipeline as case studies
- The five-theme palette switcher (`midnight`, `sunset`, `mint`, `buttercream`)
- A small analytics endpoint that ingests page views and renders a private dashboard
- Project-level subdomains (`bakeoff.ce.dev`, `linux.ce.dev`) when those projects are built
- An RSS feed for the journal
- A search index for case studies (Pagefind builds at deploy time)
