# Tasks

Phased build plan for the portfolio site. Each task is self-contained: hand it to Claude Code one at a time. Each task includes a goal, inputs, outputs, and acceptance criteria.

**How to use this file.** Open the next task in your queue, copy its block (from `## Task X.Y` to the next `## Task` heading), paste it into Claude Code as the prompt, and let it work. When the task is complete, check the acceptance criteria, mark it done in this file (`[x]`), commit, and move on.

**Reference docs.** Every task assumes Claude Code has read `README.md`, `ARCHITECTURE.md`, and `DESIGN.md` first. Pin those three files in the Claude Code context. The current task is the only thing that should change between sessions.

**Bio block (used in multiple tasks).** Cole's bio for the about section, drawn from his actual history:

> I started as a systems engineer at GE in 2012, writing Bash and Perl to keep one of the largest backup environments in the world running. Over a decade I worked my way through cloud and DevOps roles at NIAID, RoundTower, MediaMonks, Hex, and Pinterest, picking up Python and JavaScript along the way. Today I'm a Senior DevOps / Platform Engineer based in the Bay Area.
>
> I came to frontend the long way. SvelteKit clicked about two years ago. I started learning Go six months ago and have been using it ever since. The four projects below are how I learn — by building things I'd actually want to use.
>
> When I'm not at the keyboard you'll find me hiking with my dog Willy, cooking through old family recipes, or watching baseball.

**Concrete numbers to surface (use sparingly):**
- 10+ years building cloud infrastructure
- 60+ apps migrated to GCP/AWS at EPEX SPOT
- 7PB+ of backup data managed at GE across 9,200 servers
- 1,400 clients / 5.5PB managed at NIAID
- $10k+/month in AWS savings delivered to clients

---

# Phase 0 — Foundation (week 1)

Goal: a working SvelteKit + Go monorepo deploying to Cloud Run, with the design system in place but no real content yet. By the end of this phase, pushing to main deploys a "Hello world" page styled with cotton-candy tokens to a real `*.run.app` URL.

- [ ] Task 0.1 — Monorepo scaffold
- [ ] Task 0.2 — SvelteKit app with Tailwind v4 + design tokens
- [ ] Task 0.3 — Go API skeleton with health endpoint
- [ ] Task 0.4 — Dockerfiles for both services
- [ ] Task 0.5 — GitHub Actions CI pipeline
- [ ] Task 0.6 — Terraform for Cloud Run + Cloud SQL + DNS
- [ ] Task 0.7 — First end-to-end deploy

---

## Task 0.1 — Monorepo scaffold

**Goal.** Create the repository structure described in `ARCHITECTURE.md` with empty placeholder files where needed, plus root tooling (pnpm workspace, .gitignore, .editorconfig, basic README).

**Inputs.**
- `README.md` (the repo layout section)
- `ARCHITECTURE.md` (the project structure sections for both apps)

**Outputs.**
- `pnpm-workspace.yaml` declaring `apps/*` and `packages/*` as workspaces
- `package.json` at the root with workspace scripts (`dev`, `build`, `lint`, `test`)
- `.gitignore` covering Node, Go, Terraform, environment files, build outputs
- `.editorconfig` for consistent indentation (2 spaces JS/TS, tabs for Go)
- Empty `apps/web/`, `apps/api/`, `packages/content/`, `infra/terraform/`, `infra/docker/`, `.github/workflows/` directories
- `LICENSE` (MIT)

**Acceptance criteria.**
- `pnpm install` runs at the root without error (it will be a no-op since no packages have dependencies yet)
- The directory tree matches `ARCHITECTURE.md` exactly
- `git status` shows a clean tree after the initial commit
- README.md, ARCHITECTURE.md, DESIGN.md, and TASKS.md sit at the repo root

---

## Task 0.2 — SvelteKit app with Tailwind v4 and design tokens

**Goal.** Initialize a SvelteKit app with TypeScript, Tailwind v4, and the cotton-candy design tokens from `DESIGN.md` wired up as CSS custom properties. The default `+page.svelte` should render a centered "cole.eckelberry" placeholder using the proper fonts and palette.

**Inputs.**
- `DESIGN.md` (the entire Palette and Typography sections)
- `ARCHITECTURE.md` (the Frontend project structure section)

**Outputs.**
- `apps/web/` initialized via `pnpm create svelte@latest` with these answers:
  - Skeleton project, TypeScript, ESLint, Prettier, Vitest, Playwright (skip), no Tailwind from the official prompt (we install manually for v4)
- Tailwind v4 installed via `pnpm add -D tailwindcss@next @tailwindcss/vite@next`
- `vite.config.ts` updated to include the Tailwind plugin
- `apps/web/src/app.css` with:
  - `@import "tailwindcss"`
  - A `:root` block with all CSS variables from `DESIGN.md` (light mode values)
  - A `[data-mode="dark"]` block with dark mode values
  - Font imports for Inter (400, 500), Instrument Serif (italic 400), JetBrains Mono (400, 500) via `@fontsource/*` packages
- `apps/web/src/app.html` updated with `<html lang="en" data-theme="cotton-candy" data-mode="dark">` and a small inline script that reads `localStorage` to set `data-mode` before paint (avoid flash)
- `apps/web/src/lib/theme/tokens.ts` exporting the token names as TypeScript constants (so other modules can reference them type-safely)
- `apps/web/src/routes/+page.svelte` rendering a centered `<h1>cole.eckelberry</h1>` with the display type token, on top of `var(--bg-canvas)`
- `apps/web/tailwind.config.ts` is NOT used (Tailwind v4 uses CSS-based config) — instead, define theme extensions in `app.css` using `@theme`

**Acceptance criteria.**
- `pnpm --filter web dev` starts the dev server on port 5173 with no warnings
- The page renders with `Inter` for the heading, the dark cotton-candy palette by default, and the heading visible in white on the dark canvas
- Toggling `data-mode="light"` on the `<html>` tag in devtools immediately swaps to light mode without errors
- `pnpm --filter web build` produces a successful production build
- All three font families are confirmed loading by checking the Network tab (no fallback flicker)

---

## Task 0.3 — Go API skeleton with health endpoint

**Goal.** Initialize the Go service with a `chi` router, structured logging via `slog`, env-based config, and a single `/api/health` endpoint that returns 200 with a JSON status payload.

**Inputs.**
- `ARCHITECTURE.md` (the Backend section)

**Outputs.**
- `apps/api/go.mod` initialized at module path `github.com/CEckelberry/portfolio/apps/api`
- Dependencies added: `github.com/go-chi/chi/v5`, `github.com/jackc/pgx/v5/pgxpool` (not yet wired up; just add it for later), `github.com/joho/godotenv` for local dev
- `apps/api/cmd/server/main.go` — entrypoint that:
  - Loads config via `internal/config`
  - Sets up structured logging via `internal/logging`
  - Creates the chi router via `internal/http`
  - Starts an `http.Server` on `:${PORT}` (default 8080)
  - Implements graceful shutdown on SIGINT/SIGTERM with a 10s timeout
- `apps/api/internal/config/config.go` — loads `PORT`, `DATABASE_URL`, `LOG_LEVEL`, `ENVIRONMENT` from env, with sensible defaults for local dev
- `apps/api/internal/logging/logger.go` — wraps `slog` with JSON output in production, text in local dev
- `apps/api/internal/http/router.go` — sets up chi with middleware (request ID, logger, recoverer, real-IP, timeout) and registers the health route
- `apps/api/internal/http/middleware.go` — request ID and structured logging middleware
- `apps/api/internal/http/handlers/health.go` — handler returning `{"status":"ok","environment":"local","version":"dev"}`

**Acceptance criteria.**
- `cd apps/api && go run ./cmd/server` starts the server on port 8080
- `curl localhost:8080/api/health` returns `{"status":"ok",...}` with HTTP 200
- Sending SIGINT (Ctrl+C) gracefully shuts down within 1 second
- `go vet ./...` passes with no warnings
- `go test ./...` passes (no tests yet, but the command should succeed)
- Structured logs print to stdout in JSON format when `ENVIRONMENT=production` and human-readable text when not

---

## Task 0.4 — Dockerfiles for both services

**Goal.** Multi-stage, distroless Dockerfiles for `web` and `api`, each producing minimal images suitable for Cloud Run. Plus a top-level `docker-compose.yml` for local development that runs both services and a Postgres container.

**Inputs.**
- `ARCHITECTURE.md` (the Container builds section)

**Outputs.**
- `apps/web/Dockerfile` — multi-stage:
  - Stage 1: `node:20-alpine` builds the SvelteKit app via `pnpm install --frozen-lockfile && pnpm build`
  - Stage 2: `gcr.io/distroless/nodejs20-debian12` runs `node build/index.js` on port 3000
  - Uses `@sveltejs/adapter-node`
- `apps/api/Dockerfile` — multi-stage:
  - Stage 1: `golang:1.23-alpine` builds with `CGO_ENABLED=0 go build -ldflags='-s -w' -o /server ./cmd/server`
  - Stage 2: `gcr.io/distroless/static-debian12` runs `/server` on port 8080
- `docker-compose.yml` at the repo root with three services:
  - `web` (built from `apps/web/Dockerfile`, port 3000)
  - `api` (built from `apps/api/Dockerfile`, port 8080, depends on `db`)
  - `db` (image `postgres:16-alpine`, port 5432, with a named volume)
- `apps/web/.dockerignore` and `apps/api/.dockerignore` excluding node_modules, .git, build outputs

**Acceptance criteria.**
- `docker compose up --build` brings up all three services without errors
- `curl localhost:3000/` returns the SvelteKit hello-world page
- `curl localhost:8080/api/health` returns the JSON status
- Both built images are under 60MB (`docker images`)
- `docker compose down -v` cleans up cleanly

---

## Task 0.5 — GitHub Actions CI pipeline

**Goal.** A single `.github/workflows/ci.yml` that runs on every push and PR: lints both apps, builds both images, and (on push to main only) pushes images to Artifact Registry and deploys to Cloud Run via Workload Identity Federation.

**Inputs.**
- `ARCHITECTURE.md` (the CI/CD section)

**Outputs.**
- `.github/workflows/ci.yml` with two jobs:
  - `test` — runs on every event:
    - Checks out code
    - Sets up pnpm and Node 20
    - Runs `pnpm install --frozen-lockfile`
    - Runs `pnpm --filter web lint && pnpm --filter web check && pnpm --filter web test run`
    - Sets up Go 1.23
    - Runs `cd apps/api && go vet ./... && go test ./...`
  - `deploy` — runs only on push to main, after `test` succeeds:
    - Authenticates to GCP via WIF (assume the secret `WIF_PROVIDER` and service account email `WIF_SERVICE_ACCOUNT` are already configured in repo secrets)
    - Builds both images with `docker buildx`, tags with the commit SHA
    - Pushes to `us-west1-docker.pkg.dev/${PROJECT_ID}/portfolio/web:${SHA}` and `.../api:${SHA}`
    - Deploys both Cloud Run services using `google-github-actions/deploy-cloudrun@v2`
    - Sets traffic to 100% on the new revision after a successful smoke test (a `curl` to `/` and `/api/health`)
- `.github/workflows/README.md` documenting which secrets the workflow expects

**Acceptance criteria.**
- The `test` job runs successfully on a fresh PR with green checkmarks
- The `deploy` job is correctly skipped on PR events
- The workflow file passes `actionlint` (run locally to verify)
- All secret references use repository secrets, never hardcoded values

---

## Task 0.6 — Terraform for Cloud Run + Cloud SQL + DNS

**Goal.** A Terraform module that provisions everything needed for the production environment: VPC, serverless VPC connector, Cloud SQL Postgres, both Cloud Run services, Artifact Registry, IAM, and the GitHub Actions WIF binding.

**Inputs.**
- `ARCHITECTURE.md` (the Infrastructure as code section, including the file layout)

**Outputs.**
- `infra/terraform/main.tf` — provider block, GCS backend
- `infra/terraform/variables.tf` — `project_id`, `region`, `domain_name`, `github_repo`
- `infra/terraform/network.tf` — VPC with a single subnet, serverless VPC connector
- `infra/terraform/cloud_run.tf` — both services (`web`, `api`), with the api service connected to the VPC connector
- `infra/terraform/cloud_sql.tf` — `db-f1-micro` Postgres 16 instance with private IP only, plus a database `portfolio` and a user `portfolio`
- `infra/terraform/artifact_registry.tf` — single repo `portfolio` for both images
- `infra/terraform/dns.tf` — managed zone, A/AAAA records pointing the apex and `www` to a Cloud Run domain mapping
- `infra/terraform/iam.tf` — service accounts for both Cloud Run services, WIF pool and provider for GitHub Actions, role bindings for `roles/run.developer`, `roles/artifactregistry.writer`, `roles/secretmanager.secretAccessor`
- `infra/terraform/monitoring.tf` — uptime checks on `/` and `/api/health`, alert policy notifying email
- `infra/terraform/outputs.tf` — outputs the Cloud Run service URLs, the Artifact Registry repo, the WIF provider name
- `infra/terraform/README.md` — explains the bootstrap (creating the GCS backend bucket, the initial project, enabling APIs)

**Acceptance criteria.**
- `terraform init && terraform validate` succeeds in `infra/terraform/`
- `terraform plan` against an empty state shows the full set of resources without errors
- `terraform fmt -check` passes
- Sensitive values (database password) are generated via `random_password` and stored in Secret Manager, never written to state in plaintext
- The README walks the owner through the manual bootstrap steps clearly

---

## Task 0.7 — First end-to-end deploy

**Goal.** Run the bootstrap, push to main, and confirm both services are reachable at their `*.run.app` URLs (custom domain comes later in Phase 4).

**Inputs.**
- All previous Phase 0 outputs
- A real GCP project (the owner provides this)

**Outputs.**
- A `DEPLOY.md` at the repo root documenting:
  - The exact commands to bootstrap (create the GCS backend bucket, set up Workload Identity)
  - How to set the GitHub Actions secrets (`WIF_PROVIDER`, `WIF_SERVICE_ACCOUNT`, `GCP_PROJECT_ID`)
  - How to manually trigger the first deploy if needed (`gh workflow run ci.yml`)
  - Where to find logs (`gcloud logging read --resource.type=cloud_run_revision`)

**Acceptance criteria.**
- A push to main triggers the workflow
- Both Cloud Run services deploy successfully and serve their respective endpoints
- `https://web-xxxx.run.app/` returns the SvelteKit hello-world page
- `https://api-xxxx.run.app/api/health` returns 200 OK with the JSON status
- `gcloud run services list` shows both services with revisions tagged with the commit SHA
- The owner can run `terraform destroy` and `terraform apply` cleanly to recreate the environment

---

# Phase 1 — Marketing pages (weeks 2–3)

Goal: the home page renders the full hero, projects gallery, about, and footer using real content. The site looks like the mockups; only the case study pages are placeholders.

- [ ] Task 1.1 — Layout shell (nav, footer, mesh background)
- [ ] Task 1.2 — Hero section
- [ ] Task 1.3 — Hero terminal animation
- [ ] Task 1.4 — Stack marquee
- [ ] Task 1.5 — Projects gallery — structure
- [ ] Task 1.6 — Project card — Linux preview
- [ ] Task 1.7 — Project card — Bake-off preview
- [ ] Task 1.8 — Project card — Diamond preview
- [ ] Task 1.9 — Project card — Terraplane preview
- [ ] Task 1.10 — Filter chips + client-side filtering
- [ ] Task 1.11 — About section
- [ ] Task 1.12 — Footer
- [ ] Task 1.13 — Dark/light mode toggle

---

## Task 1.1 — Layout shell (nav, footer placeholder, mesh background)

**Goal.** Build the global layout: the floating glass nav at the top, a placeholder footer at the bottom, and the drifting mesh-blob background that lives behind all content. The mesh should be visible on every page.

**Inputs.**
- `DESIGN.md` (the Components → Nav and the Mesh blobs sections)
- The hero and projects gallery widget references in this conversation (the visuals already shown)

**Outputs.**
- `apps/web/src/routes/+layout.svelte` — wraps `<slot />` in a shell with:
  - The mesh background as the bottom layer (z-index 0, pointer-events none)
  - The polygon decoration SVG as the next layer (z-index 1, pointer-events none, aria-hidden)
  - The nav as a floating element at the top (z-index 3)
  - The main content area (z-index 2) with proper padding
  - A footer placeholder (full implementation in Task 1.12)
- `apps/web/src/lib/components/shell/Nav.svelte` — the floating glass pill nav with:
  - Live status pill on the left (animated pulsing teal dot, "deployed N ago · main@SHA")
  - Center nav links: work, about, writing, contact
  - The `deployed` time and SHA are filled in from a build-time injected constant
- `apps/web/src/lib/components/shell/Footer.svelte` — minimal placeholder (one line: "footer goes here") for now
- `apps/web/src/lib/components/effects/MeshBackground.svelte` — four absolutely-positioned blob divs with the drift animations from `DESIGN.md`. Each blob accepts color, position, size, and animation-duration props
- `apps/web/src/lib/components/effects/PolygonField.svelte` — an inline SVG with 4–6 polygon shapes at low opacity, positioned procedurally based on a seeded random
- `apps/web/vite.config.ts` updated to inject the build-time constants (`__BUILD_SHA__`, `__BUILD_TIME__`) using `define`

**Acceptance criteria.**
- The dev server shows the floating nav at the top of every route
- The mesh blobs visibly drift over a 16–22s loop without stuttering
- The nav's "deployed N ago" updates correctly on each build (verify by changing the build time and checking the rendered string)
- The mesh and polygons have `aria-hidden="true"` and do not interfere with keyboard navigation
- `prefers-reduced-motion: reduce` is respected: the mesh blobs hold still
- The layout shell does not introduce horizontal scroll on any viewport from 320px to 1920px

---

## Task 1.2 — Hero section

**Goal.** Build the hero region for the home page: eyebrow, name (with the gradient-accent last name), serif italic tagline, bio paragraph, and CTA buttons. The terminal goes in a separate task; the stack marquee too.

**Inputs.**
- `DESIGN.md` (the Typography section, the bio block in this file)
- The `concept_b2_hero_above_fold` widget shown earlier as a layout reference

**Outputs.**
- `apps/web/src/lib/components/hero/Hero.svelte` — composes the eyebrow, name, tagline, bio, and CTAs
- `apps/web/src/lib/components/hero/HeroName.svelte` — renders "Cole Eckelberry" in display type. The first name is plain white; the last name "Eckelberry" uses the gradient text-fill from the concept (pink → purple → teal at 135°). The whole name must fit on one line at every viewport ≥320px wide; use `clamp()` and verify
- `apps/web/src/lib/components/ui/Button.svelte` — primary + secondary variants, with the magnetic-pull hover behavior described in `DESIGN.md`. Pure CSS, no JS for the magnetic effect at v1 (transform on `:hover` is enough; full magnetism is v2)
- The hero text content:
  - Eyebrow: `PORTFOLIO · 2026`
  - Name: `Cole Eckelberry`
  - Tagline: `A platform engineer who keeps making weird things after 5pm.` with "making weird things" italicized in serif and underlined with the pink-soft underline highlight
  - Bio: 2 sentences max — "Ten years of building cloud platforms by day. Four side projects below — all open source, all live."
  - CTAs: "see the work" (primary) and "cv.pdf" (secondary)

**Acceptance criteria.**
- The full name "Cole Eckelberry" renders on one line at viewports from 320px through 1920px wide
- The gradient on "Eckelberry" is smooth and uses the three accent colors at 135°
- Tagline italics use Instrument Serif and the pink-soft underline highlight extends across "making weird things" with no gaps
- Both buttons hover correctly with the lift and color shift
- The "cv.pdf" button links to `/cv` (which currently 404s — that's fine for now)
- Lighthouse Accessibility score on the hero region is 100

---

## Task 1.3 — Hero terminal animation

**Goal.** A miniature terminal that boots on first paint, types out a sequence of commands, and ends in a blinking cursor. Sits to the right of the hero text on desktop, below it on mobile.

**Inputs.**
- `DESIGN.md` (the Hero terminal component anatomy and color rules)
- The `concept_b2_hero_above_fold` widget for visual reference

**Outputs.**
- `apps/web/src/lib/components/hero/Terminal.svelte` — a self-contained component that:
  - Renders the title bar (3 traffic-light dots, "~/cole — zsh")
  - Renders the body with a typed-character animation
  - The boot sequence (each character with a 30ms delay):
    ```
    cole@bay-area $ whoami
    senior platform engineer
    cole@bay-area $ cat now.txt
    role:    "open to staff/principal cloud + devops"
    stack:   ["python", "sveltekit", "go"]
    city:    "san jose, ca"
    open:    true
    cole@bay-area $ █
    ```
  - The blinking cursor animation (CSS `step(2)` 1s infinite) starts after the boot completes
  - Uses `requestAnimationFrame` (not `setInterval`) for the typing animation to play nice with the browser
  - Stops gracefully if `prefers-reduced-motion: reduce` — skip to the final state with the cursor blinking
- All terminal coloring uses the tokens defined in `DESIGN.md`'s Hero terminal section

**Acceptance criteria.**
- The terminal boots automatically on page load and finishes within 1.6s
- Each character types in sequence; there's no jank or layout shift during the animation
- The blinking cursor begins immediately after the last command
- With `prefers-reduced-motion: reduce` enabled, the terminal renders the final state instantly
- The terminal does not block the main thread (verify in devtools Performance tab — no long tasks during boot)
- On mobile, the terminal sits below the hero text with appropriate margin

---

## Task 1.4 — Stack marquee

**Goal.** A horizontally scrolling band of tech stack items below the hero, with diamond dividers and edge fades. Loops infinitely; never stops.

**Inputs.**
- `DESIGN.md` (the Stack marquee component section, including the v1 item list)

**Outputs.**
- `apps/web/src/lib/components/hero/StackMarquee.svelte` — full implementation. Use a duplicated content track and a CSS keyframe animation (no JS scroll loop). Edge fades via `mask-image` linear gradient.
- The v1 items list (from `DESIGN.md`):
  1. SvelteKit
  2. Python *(my main)*
  3. TypeScript
  4. Kubernetes
  5. Terraform
  6. PostgreSQL
  7. GCP
  8. Docker
  9. Bash *(since 2012)*
- Hover behavior: when the user hovers the marquee, slow the animation from 30s/loop to 60s/loop. Do not stop.
- Honor `prefers-reduced-motion: reduce` by pausing the animation entirely

**Acceptance criteria.**
- The marquee scrolls smoothly at 30s per loop with no stutter
- Hovering slows it to 60s per loop without snapping
- Edge fades extend ~60px on each side and blend cleanly with the page background
- With `prefers-reduced-motion: reduce`, the marquee is paused and shows the static items
- The marquee does not introduce horizontal scroll on the page

---

## Task 1.5 — Projects gallery — structure

**Goal.** The 2x2 grid of project cards beneath the hero, with section header, filter chips bar (filters wired in Task 1.10), and the four card slots. Each card uses placeholder content for now; real previews come in Tasks 1.6–1.9.

**Inputs.**
- `DESIGN.md` (the Project card and Filter chips component sections)
- The `concept_b2_projects_gallery` widget for visual reference

**Outputs.**
- `apps/web/src/lib/components/projects/ProjectsSection.svelte` — wraps the section header, filter chips, and gallery grid
- `apps/web/src/lib/components/projects/ProjectCard.svelte` — generic card that accepts:
  - `number` (e.g. "01"), `kind` (e.g. "A GAME"), `title` (with optional italic accent), `description`, `tags` (array), `accent` (one of `pink | teal | purple | amber`), `status`, `caseStudyHref`
  - Renders the preview region as a `<slot name="preview">` so each card supplies its own preview
- `apps/web/src/lib/content/projects.ts` — typed metadata for the four projects (title, description, tags, kind, accent, slug). This is the single source of truth for project metadata; both the gallery and the case-study route read from it
- The four project entries with placeholder previews (a colored gradient block per `accent` until 1.6–1.9 land)
- Section header with eyebrow "SECTION 02 · WORK", title "Things I built when I should've been *sleeping.*", and lead-in copy from the mockup
- Filter chips row (chips render but clicking does nothing yet — wired in Task 1.10)

**Acceptance criteria.**
- The 2x2 grid renders correctly on desktop, stacks to single column under 768px
- Cards have the correct hover state (lift, border brighten)
- The four projects appear in the right order: Linux, Bake-off, Diamond, Terraplane
- Each card's accent color matches its assignment (pink/teal/purple/amber)
- Project metadata in `projects.ts` is fully typed (no `any`)

---

## Task 1.6 — Project card preview: Linux Lessons from Hell

**Goal.** Build the preview region for the Linux project card: a small terminal-styled fragment showing wasteland-themed shell output. Pink accent.

**Inputs.**
- `DESIGN.md`
- The Linux preview region in `concept_b2_projects_gallery` for reference

**Outputs.**
- `apps/web/src/lib/components/projects/previews/LinuxPreview.svelte` — renders the static terminal fragment with sample bash output, accented with the pink palette. No animation at v1 (the case study has the live demo)
- Wire the component into the Linux card via the `preview` slot

**Acceptance criteria.**
- The preview fits within the card's 160px preview region without overflow
- All text uses JetBrains Mono
- The pink accent is applied consistently and contrasts adequately with the dark preview background
- The "teacher" lines (italic gray text) read clearly

---

## Task 1.7 — Project card preview: Backend Bake-off

**Goal.** Build the preview region for the Bake-off card: a small "P95 latency" stacked-bar fragment showing the six runtimes with realistic-looking values. Teal accent.

**Inputs.**
- `DESIGN.md`
- The Bake-off preview in `concept_b2_projects_gallery` for reference

**Outputs.**
- `apps/web/src/lib/components/projects/previews/BakeoffPreview.svelte` — renders six rows (go, rust, bun, node, python, php) with bar widths proportional to plausible p95 latencies. Each bar is colored per runtime:
  - go: teal, rust: amber, bun: pink-soft, node: purple-soft, python: pink, php: pink
- Static at v1 (no live data; the case study has the real demo)

**Acceptance criteria.**
- All six rows visible and properly aligned within the 160px preview region
- The bar widths reflect a realistic distribution (Rust shortest, PHP longest)
- Numbers are right-aligned and use mono font

---

## Task 1.8 — Project card preview: Diamond Departures

**Goal.** Build the preview region for the Diamond Departures card: a Penn Station split-flap-style row of MLB players with positions and stats. Purple accent.

**Inputs.**
- `DESIGN.md`
- The Diamond preview in `concept_b2_projects_gallery` for reference

**Outputs.**
- `apps/web/src/lib/components/projects/previews/DiamondPreview.svelte` — renders 6 rows of player data with the split-flap aesthetic (amber-on-deep-purple monospace). Sample data for top 6 players by wRC+ (Judge, Ohtani, Soto, Witt Jr., Betts, De La Cruz)
- Up/down arrows next to stats indicating the direction of recent change

**Acceptance criteria.**
- The split-flap style is unmistakable (amber text, deep purple background, mono font)
- The rows are properly aligned with consistent column widths
- The up/down indicators use teal (up) and pink (down) consistently
- The preview fits within the 160px region

---

## Task 1.9 — Project card preview: Terraplane

**Goal.** Build the preview region for the Terraplane card: a small SVG architecture graph with a VPC node at the top branching to GKE and Cloud SQL, then leaf nodes below. Amber accent.

**Inputs.**
- `DESIGN.md`
- The Terraplane preview in `concept_b2_projects_gallery` for reference

**Outputs.**
- `apps/web/src/lib/components/projects/previews/TerraplanePreview.svelte` — renders an inline SVG with:
  - One central VPC node at the top
  - Two mid-tier nodes (GKE, Cloud SQL)
  - 4 leaf nodes
  - Dashed amber connectors between layers
  - A "$1,247/mo" cost label above the graph
- All amber on a deep brown background (`#412402` → `#1a0e02` gradient)

**Acceptance criteria.**
- The SVG is responsive within its container (uses `viewBox`, no fixed pixel dimensions)
- The architecture is readable at the small preview size
- The dashed connectors render correctly in all browsers

---

## Task 1.10 — Filter chips + client-side filtering

**Goal.** Wire up the filter chips bar so clicking a category filters the visible cards client-side with a smooth transition.

**Inputs.**
- `DESIGN.md` (the Filter chips component section)
- `apps/web/src/lib/content/projects.ts` (each project has a `kind` field that maps to a filter)

**Outputs.**
- `apps/web/src/lib/components/projects/FilterChips.svelte` — accepts a list of filter labels and the active filter, emits `change` events
- Update `ProjectsSection.svelte` to:
  - Track active filter in `$state`
  - Filter the project list in `$derived`
  - Render only matching cards
  - Animate the filter transition with a 200ms cross-fade (svelte's `transition:fade` is fine)
- Filter labels: All / Games / Tools / Data / Visualizers (mapping: Linux → Games, Bake-off → Tools, Diamond → Data, Terraplane → Visualizers)

**Acceptance criteria.**
- Clicking a chip immediately filters the visible cards
- The active chip's appearance changes (white background, dark text)
- Transitions are smooth, no layout shift
- Selecting "All" restores all four cards in original order

---

## Task 1.11 — About section

**Goal.** The about section beneath the gallery: section header, two-column layout with bio on the left and a "recent activity" placeholder card on the right. The activity feed is wired in Phase 4 — for v1, it shows static placeholder content.

**Inputs.**
- `DESIGN.md` (Voice and tone, the about-section reference in `about_footer_and_theme` widget)
- The bio block at the top of this file

**Outputs.**
- `apps/web/src/lib/components/about/AboutSection.svelte` — renders the section header, two-column grid, and both column contents
- `apps/web/src/lib/components/about/Bio.svelte` — renders the bio paragraphs with the serif-italic opening line treatment described in `DESIGN.md`
- `apps/web/src/lib/components/about/ActivityCard.svelte` — placeholder version with static rows; structured so it can later be swapped for a live-data version that calls `/api/activity`

**Acceptance criteria.**
- The bio renders with the first paragraph in serif italic and subsequent paragraphs in sans
- The two-column layout collapses to single column under 900px
- The activity card looks identical to the design at v1 (with static placeholder rows)
- Recent activity placeholder shows a deploy entry, two commit entries, and a non-Last.fm activity entry (e.g. "finished · Designing Data-Intensive Applications" or a hike/run)

---

## Task 1.12 — Footer

**Goal.** The full four-column footer with brand block, work links, elsewhere links, and contact links, plus the bottom row with build info and copyright.

**Inputs.**
- `DESIGN.md` (Footer component section)
- The `about_footer_and_theme` widget reference

**Outputs.**
- `apps/web/src/lib/components/shell/Footer.svelte` — complete implementation, replacing the placeholder from Task 1.1
- Use the build-time constants (`__BUILD_SHA__`, `__BUILD_TIME__`) for the uptime line
- All links functional (work links go to `/work/[slug]`, contact email opens mailto, cv link goes to `/cv`)
- Copyright says "©2026 — made in Bay Area"

**Acceptance criteria.**
- The four columns render correctly on desktop and stack on mobile (under 768px)
- All links have correct hrefs
- The pulsing teal status dot animates correctly
- The "konami code" microcopy is NOT present (per the v1 scoping decisions)

---

## Task 1.13 — Dark/light mode toggle

**Goal.** A small theme toggle accessible from the nav (or a sub-bar) that flips between dark and light mode, persists to `localStorage`, and respects `prefers-color-scheme` on first visit.

**Inputs.**
- `DESIGN.md` (the Palette section, both light and dark values)

**Outputs.**
- `apps/web/src/lib/theme/mode.ts` — exports `getMode()`, `setMode()`, `toggleMode()` that read/write `localStorage["mode"]` and update `document.documentElement.dataset.mode`
- `apps/web/src/lib/components/shell/ModeToggle.svelte` — a small button with sun/moon SVG icons that toggles the mode
- Update `app.html` so the inline pre-paint script reads `localStorage["mode"]` (or falls back to `prefers-color-scheme`) and sets `data-mode` before first paint
- Wire the toggle into the nav (or just below it in a sub-bar — visual designer's call)

**Acceptance criteria.**
- Toggling the theme immediately swaps all colors with no flash
- The choice persists across reloads
- A first visit honors `prefers-color-scheme: dark` (defaults to dark) or light if the OS preference is light
- No flash of unstyled content on initial load — the inline script in `app.html` handles this
- All accent colors (pink, teal, purple, amber) remain visible and accessible in both modes

---

# Phase 2 — Case studies (weeks 4–5)

Goal: every project's case study page is built using a shared template, populated with real markdown content, and includes the live-demo embeds.

- [ ] Task 2.1 — Markdown content pipeline
- [ ] Task 2.2 — Case study page shell + hero
- [ ] Task 2.3 — Case study section components (lead, prose, code block, decision card, lessons)
- [ ] Task 2.4 — Architecture diagram component (per-case-study SVG)
- [ ] Task 2.5 — Live demo embed wrapper
- [ ] Task 2.6 — Adjacent projects component
- [ ] Task 2.7 — Write the four case studies
- [ ] Task 2.8 — Page transitions

---

## Task 2.1 — Markdown content pipeline

**Goal.** Set up `mdsvex` so case-study markdown files in `packages/content/case-studies/` can be rendered as SvelteKit pages with embedded Svelte components for live demos, code blocks, and architecture diagrams.

**Inputs.**
- `ARCHITECTURE.md` (the Content strategy section)

**Outputs.**
- `apps/web/svelte.config.js` updated to include `mdsvex` as a preprocessor
- `packages/content/case-studies/*.md` — four placeholder files (one per project) with frontmatter and a single H1 + paragraph each
- `apps/web/src/routes/work/[slug]/+page.ts` — load function that:
  - Reads the matching markdown file via `import.meta.glob`
  - Returns the parsed module + frontmatter
- `apps/web/src/routes/work/[slug]/+page.svelte` — renders the loaded content as the page body
- `apps/web/src/lib/content/case-studies.ts` — type definitions for case-study frontmatter

**Acceptance criteria.**
- Visiting `/work/bakeoff`, `/work/linux`, `/work/diamond`, `/work/terraplane` each renders the corresponding markdown
- Frontmatter is correctly parsed and passed to the page
- Hot reload works in dev (editing a markdown file refreshes the page)
- Production build prerenders all four case-study pages (verify in `apps/web/build/prerendered/`)

---

## Task 2.2 — Case study page shell + hero

**Goal.** The shared layout for any case-study page: nav, mesh background, the case-study hero (eyebrow, title, deck, meta row), and the live-demo frame placeholder.

**Inputs.**
- `DESIGN.md` (the Case study hero component section)
- The `case_study_backend_bakeoff` widget for visual reference

**Outputs.**
- `apps/web/src/lib/components/case-study/CSHero.svelte` — renders eyebrow, title (with italic phrase support), serif italic deck, and the 4-column meta row
- `apps/web/src/lib/components/case-study/CSDemoFrame.svelte` — the glass-card frame with the fake browser bar; accepts a slot for the actual demo content
- Updates to the case-study page route to use these components

**Acceptance criteria.**
- The hero matches the visual reference exactly (typography, spacing, colors)
- The meta row's hairlines render correctly above and below
- The "live" status pill animates with the pulsing teal dot
- The demo frame's browser bar shows the project's URL correctly per the project metadata

---

## Task 2.3 — Case study section components

**Goal.** The reusable section primitives every case study uses: numbered section header, prose lead, body prose, code block, decision card, lessons row, and adjacent-projects strip.

**Inputs.**
- `DESIGN.md` (the Section header, Code block, Decision card, Lessons row sections)
- The `case_study_backend_bakeoff` widget for reference

**Outputs.**
- `apps/web/src/lib/components/case-study/CSSection.svelte` — wraps an entire section: number, title, lead, slot for the body
- `apps/web/src/lib/components/case-study/CSProse.svelte` — styled prose container (typography, code, strong, em)
- `apps/web/src/lib/components/case-study/CodeBlock.svelte` — renders a code block with a top bar, line numbers, syntax highlighting, line highlighting, and an optional bottom annotation
- `apps/web/src/lib/components/case-study/DecisionCard.svelte` — Q-and-A card with picked/skipped pill
- `apps/web/src/lib/components/case-study/LessonsRow.svelte` — two-column grid; each cell is a Lesson with worked/didn't accent
- `apps/web/src/lib/components/case-study/AdjacentProjects.svelte` — three-card strip showing other projects with shared-stack microcopy

**Acceptance criteria.**
- Each component matches its visual reference
- All components are usable inside markdown via `mdsvex`'s component import support
- Code-block syntax highlighting works for at least Go, TypeScript, Bash, and HCL (via `shiki` or `highlight.js`)
- Line highlighting in the code block uses the project's accent color

---

## Task 2.4 — Architecture diagram component

**Goal.** A reusable wrapper for embedding hand-drawn SVG architecture diagrams inside case-study markdown. Each project will have its own SVG file imported and wrapped.

**Inputs.**
- `DESIGN.md` (the Architecture diagram section)

**Outputs.**
- `apps/web/src/lib/components/case-study/ArchDiagram.svelte` — accepts an SVG `src` and an optional caption; renders the SVG inline (not as an `<img>`) so accent colors can be tokenized
- A directory `apps/web/src/lib/content/diagrams/` containing four SVG files (one per case study). For now, copy the architecture SVG from the relevant widget references as a starting point — the real diagrams get refined when each case study is written

**Acceptance criteria.**
- Each diagram is responsive (`viewBox`, scales to container)
- Diagrams render correctly in both dark and light modes (verify color choices work for both)
- Captions render below the diagram in microcopy style

---

## Task 2.5 — Live demo embed wrapper

**Goal.** A consistent wrapper for the embedded live-demo region of each case study. Each project's actual demo lives in a separate repo/deployment; this wrapper provides the frame and an `<iframe>` (or component swap) for the demo content.

**Inputs.**
- `DESIGN.md` (the Case study hero section, Live demo embed)

**Outputs.**
- `apps/web/src/lib/components/case-study/LiveDemo.svelte` — accepts:
  - `url` (the deployed demo URL, used in the browser bar and as `iframe src`)
  - `title` and `subtitle`
  - A slot for in-page demo content (when the demo is built into the portfolio rather than iframed)
- For Phase 2, all four projects use the in-page slot with a static "demo coming soon — preview shown below" placeholder
- The placeholder for each project is the same visual fragment used in the gallery card preview, scaled up

**Acceptance criteria.**
- The wrapper matches the visual reference (browser bar, glass card, "live" pill)
- The placeholder renders cleanly for each of the four case studies
- Replacing the placeholder with an iframe is a one-prop change

---

## Task 2.6 — Adjacent projects component (refinement)

This is covered by Task 2.3 (`AdjacentProjects.svelte`). Mark this task complete once that component is built, with the additional acceptance criteria that the shared-stack microcopy is correctly computed:

- For each project, list the other three projects and one shared technology
- Linux ↔ Bake-off: "WASM, browser-side"
- Linux ↔ Diamond: "SvelteKit"
- Linux ↔ Terraplane: "Go"
- Bake-off ↔ Diamond: "Go SSE, GKE"
- Bake-off ↔ Terraplane: "Go, GCP infra"
- Diamond ↔ Terraplane: "Go, Postgres"

---

## Task 2.7 — Write the four case studies

**Goal.** Populate the four markdown files with real case-study content following the structure shown in the `case_study_backend_bakeoff` widget.

**Inputs.**
- All previous Phase 2 outputs
- The `case_study_backend_bakeoff` widget as a structural template (every case study should have these sections, in this order: hero / live demo, why, architecture, code excerpt, decisions, lessons, adjacent)
- The voice and tone guidelines in `DESIGN.md`

**Outputs.**
- `packages/content/case-studies/bakeoff.md` — the Bake-off case study, full content
- `packages/content/case-studies/linux.md` — the Linux Lessons case study
- `packages/content/case-studies/diamond.md` — the Diamond Departures case study
- `packages/content/case-studies/terraplane.md` — the Terraplane case study
- Each case study should be 1500–3000 words, include at least one architecture diagram, one annotated code excerpt, 4 decisions, 4 lessons, and the adjacent projects strip

**Acceptance criteria.**
- Each case study reads cleanly from top to bottom and feels like a piece of writing, not a template fill-in
- The voice matches `DESIGN.md`'s Voice and tone section (confident, plain English, occasional dry humor)
- Each case study compiles without errors and renders at its `/work/[slug]` route
- All code excerpts compile (or are explicitly labeled as pseudocode if not)
- Architecture diagrams are tailored to each project, not a generic template

---

## Task 2.8 — Page transitions

**Goal.** Use SvelteKit's view transitions API to animate from the projects gallery to a case study (and back). The clicked card morphs into the case-study hero; other transitions are 240ms cross-fades.

**Inputs.**
- SvelteKit's view transitions documentation
- `DESIGN.md` (the Page transitions section)

**Outputs.**
- Configure `+layout.svelte` to enable view transitions
- Add `view-transition-name` to project cards (computed from project slug) and to case-study heroes (matching slug)
- A small CSS file for the cross-fade keyframes
- Honor `prefers-reduced-motion: reduce` (skip transitions, just navigate)

**Acceptance criteria.**
- Clicking a project card transitions smoothly to its case study with a shared-element morph (in browsers that support view transitions — Chrome, Edge, recent Safari)
- In browsers without support, navigation works correctly with no broken transition
- The back-button transition works the same way
- With `prefers-reduced-motion: reduce`, navigation is instant

---

# Phase 3 — Polish (week 6)

Goal: the site is fast, accessible, and feels finished. Lighthouse scores in the high 90s across the board.

- [ ] Task 3.1 — Performance audit and budget enforcement
- [ ] Task 3.2 — Accessibility audit and fixes
- [ ] Task 3.3 — SEO and Open Graph metadata
- [ ] Task 3.4 — Contact form (frontend + Go backend)
- [ ] Task 3.5 — 404 page
- [ ] Task 3.6 — Custom domain + SSL

---

## Task 3.1 — Performance audit and budget enforcement

**Goal.** The site meets the performance budgets in `ARCHITECTURE.md` (LCP < 1.5s, INP < 200ms, CLS = 0, JS bundle < 80kb gzipped per page).

**Outputs.**
- A `apps/web/playwright.perf.ts` Playwright script that loads each page, captures Web Vitals, and asserts they meet the budgets
- Lazy-load the mesh background animation (it's not above-the-fold critical)
- Audit and remove any unused Tailwind utilities
- Optimize fonts: `font-display: swap`, preload only the weights actually used in the above-the-fold content
- A short report `apps/web/PERF.md` documenting the final numbers per page

**Acceptance criteria.**
- The Playwright perf script runs in CI and fails the build if budgets are exceeded
- Lighthouse Performance score is ≥ 95 on home and case-study pages
- No font flicker (FOIT or FOUT) on any page
- The home page's main JS bundle is under 60kb gzipped

---

## Task 3.2 — Accessibility audit and fixes

**Goal.** Lighthouse Accessibility score of 100 on every page. axe-core passes with zero violations.

**Outputs.**
- A `apps/web/playwright.a11y.ts` script that runs axe-core against every page and asserts no violations
- Fixes for any violations found: missing labels, color contrast, focus management, semantic HTML
- Manual keyboard navigation test: every interactive element reachable and operable via keyboard
- A "skip to content" link visible on focus

**Acceptance criteria.**
- axe-core reports zero violations on every page
- Lighthouse Accessibility is 100 on every page
- The site is fully usable with keyboard only (no mouse needed)
- Screen reader test: VoiceOver or NVDA can read the home page top-to-bottom without confusion

---

## Task 3.3 — SEO and Open Graph metadata

**Goal.** Every page has proper `<title>`, meta description, Open Graph, and Twitter card metadata. A custom-rendered OG image per page (generated at build time).

**Outputs.**
- `apps/web/src/lib/components/shell/SEO.svelte` — accepts title, description, image, and renders all the relevant `<head>` tags
- `apps/web/src/lib/og/generate.ts` — at build time, generates a 1200x630 PNG OG image for each page using `satori` or similar (cotton-candy palette, page title in display type, mesh background)
- `static/robots.txt` allowing all
- `apps/web/src/routes/sitemap.xml/+server.ts` returning a generated sitemap

**Acceptance criteria.**
- Every page has unique `<title>` and `meta name="description"`
- OG image previews correctly on Twitter/X, LinkedIn, Slack (test via opengraph.xyz or similar)
- The sitemap includes all four case studies and the home page
- robots.txt is reachable at `/robots.txt`

---

## Task 3.4 — Contact form (frontend + Go backend)

**Goal.** A working contact form on the home page (or in the footer) that submits to `POST /api/contact` and shows a success/error state.

**Inputs.**
- `ARCHITECTURE.md` (the Contact endpoint contract)

**Outputs.**
- Backend:
  - `apps/api/internal/http/handlers/contact.go` — handler implementing the contract
  - `apps/api/internal/store/contacts.go` — Postgres insert
  - `apps/api/migrations/002_contact_submissions.sql` — schema
  - Rate limiting middleware: 5 submissions per hour per source IP
  - Optional Discord webhook notification on success (configured via `DISCORD_WEBHOOK_URL`)
- Frontend:
  - `apps/web/src/lib/components/shell/ContactForm.svelte` — name, email, message fields with proper validation, submit button, success/error states
  - `apps/web/src/routes/api/contact/+server.ts` — proxies to the Go service

**Acceptance criteria.**
- Submitting the form on the live site results in a row in the `contact_submissions` table
- Validation errors (missing email, invalid email, empty message) show inline error messages
- Rate limiting works: a 6th submission within an hour returns 429
- The Discord webhook fires (if configured)
- The form is fully accessible (proper labels, error association, focus management)

---

## Task 3.5 — 404 page

**Goal.** A custom 404 page styled in the same cotton-candy aesthetic, with a useful message and links back to the major sections.

**Outputs.**
- `apps/web/src/routes/+error.svelte` — the error page

**Acceptance criteria.**
- Visiting `/nonexistent` renders the custom 404 page
- The page includes links to home, work, and contact
- It feels like part of the site, not a default

---

## Task 3.6 — Custom domain + SSL

**Goal.** The site is reachable at the production domain (e.g. `ce.dev`), with SSL and `www` → apex redirect.

**Outputs.**
- Updates to `infra/terraform/dns.tf` and `infra/terraform/cloud_run.tf` to map the custom domain
- A short section in `DEPLOY.md` documenting how to verify domain ownership in GCP

**Acceptance criteria.**
- `https://ce.dev/` resolves to the home page with a valid SSL certificate
- `https://www.ce.dev/` redirects to `https://ce.dev/`
- `http://ce.dev/` redirects to `https://ce.dev/`
- `https://ce.dev/api/health` returns 200

---

# Phase 4 — Optional v2 enhancements (later)

These are explicitly deferred from v1. Reorder freely once v1 ships.

- [ ] Task 4.1 — GitHub activity feed (real `/api/activity`, replacing the placeholder)
- [ ] Task 4.2 — Additional theme palettes (midnight, sunset, mint, buttercream)
- [ ] Task 4.3 — `/journal` markdown blog
- [ ] Task 4.4 — RSS feed for the journal
- [ ] Task 4.5 — Search across case studies (Pagefind)
- [ ] Task 4.6 — Page-view analytics dashboard
- [ ] Task 4.7 — Project subdomains as the actual projects come online (`bakeoff.ce.dev`, etc.)
