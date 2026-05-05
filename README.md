# Portfolio site

A personal portfolio for Cole Eckelberry — Senior DevOps / Platform Engineer based in San Jose, California. The site showcases four side projects with deep case studies, alongside a short bio. Designed to feel playful (cotton candy palette, animated hero, gradient meshes) but read as credible to senior engineers.

## Stack at a glance

- **Frontend**: SvelteKit (Svelte 5, runes), TypeScript, Vite, Tailwind v4 with custom design tokens
- **Backend**: Go (stdlib `net/http` + `chi` router), serves the small JSON API endpoints used for live activity, contact form, and analytics ingest
- **Data**: Postgres (Cloud SQL) for contact submissions and project metadata; project content lives in markdown in the repo
- **Hosting**: Google Cloud Run (both services), Cloud SQL Postgres, Artifact Registry for container images, Cloud Build for CI/CD
- **Domain**: TBD — ce.dev or similar via Cloud DNS
- **CDN/edge**: Cloud Run handles its own scaling; static assets served from Cloud CDN in front of a GCS bucket for build artifacts

## Repository layout

```
portfolio/
├── README.md                 ← you are here
├── ARCHITECTURE.md           ← system design, services, data flow, deploy
├── DESIGN.md                 ← design language: tokens, type, motion, components
├── TASKS.md                  ← phased build plan; hand to Claude Code one at a time
├── apps/
│   ├── web/                  ← SvelteKit app (the site)
│   └── api/                  ← Go service (small API surface)
├── packages/
│   └── content/              ← markdown case studies + project metadata (shared)
├── infra/
│   ├── terraform/            ← Cloud Run, Cloud SQL, IAM, DNS
│   └── docker/               ← Dockerfiles for both services
└── .github/workflows/        ← CI: lint, test, build, deploy on push to main
```

## What this site is and isn't

**Is**: a portfolio that treats the site itself as a project. Every interaction should feel intentional. Performance is a feature — Lighthouse 95+ on every page, INP under 200ms, no layout shifts.

**Isn't**: a CV. The /cv route serves a PDF for recruiters who want one. The site is for engineers who want to see how Cole thinks.

## Four projects (each gets its own case study page)

1. **Linux Lessons from Hell** — browser-based bash tutor with a foul-mouthed wasteland teacher. SvelteKit + xterm.js + a Go WASM shell sandbox.
2. **Backend Bake-off** — six runtimes (Go, Rust, Bun, Node, Python, PHP) implementing the same checkout endpoint, hot-swappable via header. GKE-deployed, live latency comparison.
3. **Diamond Departures** — top 100 active MLB players ranked by sabermetrics, styled as a Penn Station split-flap board, updated live during games. Go SSE backend, FanGraphs ingest.
4. **Terraplane** — paste a `terraform plan`, get a walkable architecture graph with cost estimates and blast-radius highlights. HCL parser in Go, D3 force layout in the browser.

The portfolio site links to each project's live deployment but does not host them. Each project lives in its own repo and gets its own subdomain (e.g. `bakeoff.ce.dev`).

## Build phases at a glance

See `TASKS.md` for the full breakdown. High level:

- **Phase 0** — Repo setup, design tokens, Tailwind config, deploy pipeline (1 week)
- **Phase 1** — Marketing pages: hero, projects gallery, about, footer (2 weeks)
- **Phase 2** — Case study template + four case study pages (2 weeks)
- **Phase 3** — Polish: dark/light toggle, animations, perf pass, accessibility audit (1 week)
- **Phase 4 (later)** — GitHub activity feed, additional themes, /journal blog

Total time-to-launch for v1: roughly 6 weeks of evening-and-weekend work.

## Working with Claude Code on this project

Each task in `TASKS.md` is written as a self-contained unit with:
- A goal stated in one sentence
- Inputs (files to read, designs to reference)
- Outputs (files to create or change)
- Acceptance criteria (how you'll know it's done)

Hand tasks one at a time. Do not give Claude Code the whole `TASKS.md` and ask for "everything in Phase 1" — quality drops sharply on multi-task hand-offs. A task should be roughly 30–90 minutes of focused work.

When in doubt about visual choices, `DESIGN.md` is the source of truth. When in doubt about structure, `ARCHITECTURE.md` is. When in doubt about scope, this file is.
