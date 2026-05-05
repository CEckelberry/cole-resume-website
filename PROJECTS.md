# Projects

A catalog of the four projects featured on the portfolio site. Each project gets its own repo, its own deployment, and its own subdomain when it's actually built. This file is a holding place for the design and technical ideas we sketched out during portfolio planning, so they're not lost when it's time to build the real thing.

**Status legend.**
- ⚪ Concept — idea only, not started
- 🟡 In progress — actively building
- 🟢 Live — deployed and linked from the portfolio

| # | Project | Slug | Subdomain (planned) | Accent | Status |
|---|---|---|---|---|---|
| 01 | Linux Lessons from Hell | `linux` | `linux.ce.dev` | Pink | ⚪ Concept |
| 02 | Backend Bake-off | `bakeoff` | `bakeoff.ce.dev` | Teal | ⚪ Concept |
| 03 | Diamond Departures | `diamond` | `diamond.ce.dev` | Purple | ⚪ Concept |
| 04 | Terraplane | `terraplane` | `terraplane.ce.dev` | Amber | ⚪ Concept |

Each project's case study on the portfolio site links to its live URL via a "live demo" button. Until each project is actually built, the case-study page shows a static preview of the same fragment used on the gallery card.

The subdomains are illustrative — pick whatever apex you settle on (`ce.dev`, `coleeckelberry.com`, etc.) and substitute consistently.

---

## 01 — Linux Lessons from Hell

> A real bash shell, in the browser. A foul-mouthed wasteland teacher walks you through pipes, grep, and ssh — then quizzes you on broken systems.

### The pitch

Most "learn Linux" tutorials are sterile. They explain `grep` on a textbook example and never let you feel the satisfaction of finding the one error line in 200,000 lines of nginx logs. This is a game where you actually use a real bash environment to solve increasingly broken scenarios, while a foul-mouthed post-apocalyptic teacher (think Fallout/Borderlands raunchy humor) heckles you through it.

The teacher is a character. He insults your every keystroke when you're slow, congratulates you with backhanded compliments when you finally figure something out, and gives you the kind of tough-love mentoring that actually sticks. The humor lands because the teaching underneath it is real — you walk away genuinely knowing how to navigate a Linux command line.

### Curriculum (rough outline)

The game progresses through tiers, each gated by a boss-style challenge:

**Tier 1 — Survival basics**
- `cd`, `ls`, `pwd`, `mkdir`, `rm`, `cp`, `mv`
- Reading file contents: `cat`, `less`, `head`, `tail`
- Permissions and ownership: `chmod`, `chown`, the octal vs symbolic forms
- Boss: navigate a sprawling vault filesystem to find a hidden cache

**Tier 2 — Searching and slicing**
- `grep`, including `-r`, `-i`, `-v`, `-A`, `-B`, `-C`
- `find` with predicates (`-name`, `-mtime`, `-size`, `-exec`)
- `awk` and `cut` for column extraction
- `sed` for substitution
- Boss: extract specific intel from a corrupted log archive

**Tier 3 — Pipes and process**
- Pipes, redirection (`>`, `>>`, `<`, `2>&1`, `/dev/null`)
- `xargs` for command construction
- `sort | uniq -c` patterns
- `tee` for branching output
- Background jobs, `&`, `nohup`, `jobs`, `fg`, `bg`
- Boss: pipeline a multi-step intel-gathering operation in one line

**Tier 4 — System health**
- `ps`, `top`, `htop`-style introspection
- `df`, `du`, finding what's eating disk
- `kill`, `killall`, signals
- `systemctl` basics
- Reading `/var/log/`
- Boss: diagnose why the bunker generator service keeps dying

**Tier 5 — Across the wires**
- `ssh`, key auth, `~/.ssh/config`
- `scp`, `rsync` (with the right flags actually explained)
- `curl` and `wget`
- Boss: exfiltrate data from a remote server you don't have a password for, using only key auth

**Endgame — The sysadmin gauntlet**
A series of broken systems. Each one comes with a stated symptom ("the app is returning 500s", "users can't log in", "disk is full but `du` says we should have space"). The user has to diagnose, fix, and document the fix. The teacher grades the answer.

### The teacher

Voice: profane, dismissive of stupid mistakes, oddly invested in your success when you finally get it. Occasionally philosophical about the wasteland.

Sample lines (style reference, not final):
- *"Wow. You opened a directory. Should I get the band? Should I throw a parade?"*
- *"That's the third time you've tried `chmod` without numbers. You think the computer can read minds? Numbers. Or letters. Pick a side."*
- *"...okay. That was actually clever. I hate that you made me say that."*
- *"You used `find -exec rm`. Without `-i`. Without `-name`. You absolute weapon. Hope you didn't need that home directory."*

Banter is generated from a pool of templates per scenario. Each scenario has 3–5 reaction templates per outcome (success, partial, failure, suspicious-fast-success). Hard-coded for v1; LLM-generated banter is tempting but a v2 idea.

### Tech stack

- **Frontend**: SvelteKit, TypeScript, Tailwind v4
- **Terminal**: `xterm.js` for the terminal UI
- **Shell sandbox**: a real bash running in a Linux container, accessed via WebSocket-attached PTY. Either:
  - **Option A**: a Go backend on Cloud Run that spins up per-session Docker containers and proxies the PTY
  - **Option B**: an in-browser Linux via `v86` or `webvm` — heavier first-load, no backend cost
- **Filesystem fixtures**: each scenario ships a small tarball seeded into the container at session start (a fake `/var/log/`, a fake `/etc/`, broken services for the diagnostic tier)
- **Progress storage**: Postgres if backed by Go; `localStorage` if fully client-side
- **Teacher dialogue**: JSON files keyed by scenario + outcome

### Why this is a good portfolio piece

It demonstrates: real-time backend work (PTY proxying), security thinking (sandboxing untrusted shell input), thoughtful UX (the teacher banter is the hook), and content design (the curriculum is a real curriculum). Engineers will recognize that getting xterm.js + a real bash to feel snappy in the browser is non-trivial. Recruiters will recognize that "I built a game" is more memorable than "I built a CRUD app."

### Open design questions to settle when building

- Container-per-session vs in-browser VM (cost vs first-load weight)
- How to prevent abuse: time limits, rate limits, no network egress from sandboxes
- Save/resume across sessions: required or nice-to-have?
- Should the teacher be voiced (TTS) or text-only? (Text-only at v1, voice as a stretch)
- Multiplayer leaderboard for the gauntlet times?

### Card preview (already built into portfolio)

The portfolio gallery shows a small terminal fragment with sample output. Pink accent. The case study links to this README's spirit — what the project is, why it's built this way, the stack choices, lessons learned.

---

## 02 — Backend Bake-off

> Same e-commerce checkout endpoint, six backend implementations. Hot-swap them mid-session and watch p95 latency lie to your face.

### The pitch

Every backend benchmark you've read was written by someone with a horse in the race. The problem isn't that the numbers are wrong; it's that you can't see them move under workloads you care about. This site is one frontend (a small e-commerce checkout flow) backed by six independent implementations of the same `POST /checkout` endpoint. You pick which runtime serves your requests via a tab. The page reloads, the runtime changes, and you watch p95, RPS, and cold-start metrics shift in real time.

The point isn't to declare a winner. The point is to let people change variables and see what happens.

### The six runtimes

| Runtime | Framework | Notes |
|---|---|---|
| Go | `chi` + stdlib | The "control" implementation |
| Rust | `axum` + `tokio` | The "ceiling" — what speed-of-light looks like |
| Bun | `hono` | The new entrant |
| Node | `express` | The baseline most teams actually run |
| Python | `fastapi` | Async Python, what most data orgs ship |
| PHP | `laravel` | The unfair fight that surprises people |

Deliberately not included: JVM (cold-start distortion needs too many footnotes), Deno (too similar to Bun for the sample size).

### The endpoint

`POST /checkout` accepts a cart (3–8 items, randomized weights), validates inventory, computes tax via a fake remote service (added latency simulated), runs a fraud score (a fixed CPU-bound function: 50ms of pure compute), persists the order to Postgres, and returns 200 with the order ID.

Every runtime implements this exact contract. The OpenAPI spec is the source of truth. Codegen produces server stubs and client types in every language so the implementations can't drift.

### How the hot-swap works

The frontend calls a single `POST /api/checkout` on a router service. The router (a small Go service) reads an `X-Runtime` header and forwards to the matching backend pod via Kubernetes' headless service DNS. Switching runtime is a single header change in the frontend — no redeploy, no env-var toggle.

```
Browser → Router (Go) → bo-{runtime} pod → Postgres
                                  ↓
                            Prometheus scrape
```

All six pods run continuously, all backed by the same Postgres database. Prometheus scrapes them on the same interval. The frontend pulls live metrics from `/api/metrics` (proxied through the router from Prometheus's HTTP API) and renders them as bar charts that update every 2 seconds.

### The interactive surface

The site has three modes:

- **Casual**: place orders, watch the metrics shift. No scenario, just play.
- **Comparison**: pick two runtimes, the page splits and runs the same request to both; results are rendered side-by-side. Useful for "is Rust really 2x Go on this workload?"
- **Stress**: a slider from 1 to 200 RPS. The page generates load (rate-limited per visitor IP) and shows how each runtime degrades. P95, p99, error rate.

### Tech stack

- **Frontend**: SvelteKit, TypeScript, Tailwind v4, `chart.js` for the live metrics
- **Router**: Go (`chi` + `httputil.ReverseProxy`)
- **Backends**: as listed above
- **Infra**: GKE Autopilot cluster, a Helm chart per backend, a single Postgres (Cloud SQL) shared across all
- **Metrics**: Prometheus (deployed in the cluster), Grafana for the public dashboard backend, the frontend pulls from Prometheus's HTTP API
- **Load generation**: a Go service that runs the stress mode, rate-limited per session

### Why GKE and not Cloud Run

Cloud Run scales to zero, which would distort cold-start data. Worse, the cold-start latency for some runtimes (Bun, Python) is itself an interesting metric — but it has to be measured deliberately, not by accident. GKE keeps every pod warm so the latency you see is steady-state.

### Why this is a good portfolio piece

It demonstrates: comfort across six languages (you don't have to be expert in all six, but you have to know enough), Kubernetes knowledge (the routing trick is an indictment of "service mesh" complexity for simple cases), observability thinking (Prometheus + Grafana set up cleanly), and product instinct (live demos beat blog posts).

The "decisions" section of the case study is where you get to say things like "I deliberately did not include JVM because…" — and engineers love reading those.

### Open design questions to settle when building

- Cost ceiling: GKE Autopilot is not free. Six pods running 24/7 plus Postgres might be $50–80/month. Worth it?
- Scope of the e-commerce flow: just `/checkout`? Or `/cart/add` and `/products/list` too?
- Does the frontend itself live on the same cluster, or stay on Cloud Run (probably the latter — simpler)
- Stress mode rate-limits: how to prevent a single visitor from spending your entire monthly Cloud SQL quota
- Should the backends share a Postgres or each get their own with the same schema? Sharing is fairer; isolated is more realistic.

### Card preview (already built into portfolio)

The portfolio gallery shows a six-row stacked bar chart of p95 latencies, with each runtime's row colored. Teal accent.

---

## 03 — Diamond Departures

> Top 100 active MLB players ranked by sabermetrics, styled as a Penn Station split-flap board, updated live during games.

### The pitch

There's a specific aesthetic — the old Penn Station / Frankfurt Hauptbahnhof split-flap departure boards — where the letters and numbers physically rotate when they change. It's iconic. It feels alive in a way that an LCD display doesn't. A leaderboard styled like that, ranking the top 100 active MLB players by advanced sabermetrics and updating live as games progress, is a visual nobody has built and a sports nerd's dream.

The single-page app has a few tabs: All players, by position (top 3rd basemen, top outfielders, top relievers, etc.), and by stat category. As stats update mid-game (a player gets a hit, his wRC+ shifts, his rank changes), the affected rows physically flip to their new values. Players climbing get a green up-arrow that flashes; players dropping get a red down-arrow.

### Stat categories

Players are sortable by:

**Hitters**
- AVG, OBP, SLG, OPS
- wRC+ (weighted runs created plus, the gold-standard rate stat)
- wOBA (weighted on-base average)
- BABIP, ISO
- DRS, UZR (defensive metrics for fielders)
- WAR (Wins Above Replacement, fWAR specifically)
- Total bases, RBI, runs (the traditional ones, for completeness)

**Pitchers**
- ERA, FIP, xFIP, SIERA
- ERA+ (era plus, league-adjusted)
- K/9, BB/9, K/BB, K-BB%
- WHIP, BABIP-against
- WAR (fWAR for pitchers)

**Defensive (split out)**
- DRS by position
- UZR/150 for outfielders and infielders
- OAA (Outs Above Average) when available

The default tab on landing is wRC+ for hitters — it's the single best "who is the best hitter in baseball right now" stat.

### Position views

Tabs for each position: C, 1B, 2B, 3B, SS, LF, CF, RF, DH, SP (starters), RP (relievers). Each tab shows the top players at that position by the most relevant stat for the position (wRC+ for position players, FIP for pitchers, framing runs for catchers if available).

### The split-flap aesthetic

The board sits on a deep purple background with amber monospace text. Every stat cell is a "flap" that animates when the value changes:
- The current value rotates upward (40 frames) and exits the top
- The new value rotates in from the bottom
- A subtle "click click click" sound effect (mutable, off by default)
- Total animation duration: ~600ms

When a player changes rank, the entire row slides up or down to its new position with the same flap-style sub-animations on every column. If multiple things change at once (a player got a hit AND his BABIP went up AND his wRC+ went up AND his rank changed), the animations stagger by 50ms so the eye can follow.

### Live updates

The Go backend ingests data from a sports stats provider (FanGraphs API, MLB Stats API, or a paid provider like Sportradar — costs vary). It re-computes the leaderboards every 30 seconds during active games, broadcasts changes via Server-Sent Events to all connected clients. The frontend animates only the changed cells.

For non-game-time (off-hours, off-season), the board updates from cached data and only changes when daily aggregate stats roll over.

### Tech stack

- **Frontend**: SvelteKit, TypeScript, Tailwind v4
- **Backend**: Go (`chi` + custom SSE implementation)
- **Database**: Postgres for the player roster + cached stat history; Redis for the live leaderboard cache
- **Stats provider**: FanGraphs (scrape-able with care) or MLB Stats API (free, less rich) or a paid provider for production-grade
- **Hosting**: Cloud Run for both the frontend and the Go backend; Cloud SQL Postgres; Memorystore Redis (or skip Redis at v1 and just hit Postgres)

### Why this is a good portfolio piece

It demonstrates: real-time backend work (SSE, change detection, broadcast), beautiful UI animation that has a reason to exist, comfort with sports analytics (which is a niche but a deep one), and full-stack thinking (the data pipeline matters as much as the UI).

It's also genuinely fun. People will share screenshots. That's a portfolio multiplier you can't engineer directly — it has to come from the project being beautiful.

### Open design questions to settle when building

- Stats source: free (MLB Stats API) vs paid (Sportradar). Free is fine for v1 but the stats are less rich.
- How to handle the off-season: leaderboards don't change for 4 months. A "season summary" mode?
- Mobile: the split-flap layout is naturally horizontal; how does it work on a 375px viewport? Stacked columns? Toggle to a list view?
- Sound effects: rights-free, but lots of people will mute them anyway. Worth the build?
- Historical mode: "show me the top 100 by wRC+ from 2019" — interesting but scope creep. v2.

### Card preview (already built into portfolio)

The portfolio gallery shows six rows of player data in the split-flap style with up/down arrows. Purple accent.

---

## 04 — Terraplane

> Paste a `terraform plan`. Get a walkable architecture map with cost estimates and blast-radius highlights.

### The pitch

Most teams don't have an up-to-date architecture diagram. They have a Terraform repo that *is* their architecture, but reading it requires spelunking through modules. This tool flips that: paste a `terraform plan` (or the JSON output from `terraform plan -json`) and get an interactive, walkable graph of what will be created, what depends on what, what each thing will cost, and what breaks if you delete it.

The graph is the document. Click any node to see its source HCL, its attributes, its dependencies, its blast radius (everything that depends on it transitively). Toggle a "delete this" mode and the node lights up red along with everything that would die with it.

### Core capabilities

**Parse and graph**
- Parse `terraform plan -json` (or raw HCL with a fallback parser)
- Build a dependency graph of all resources
- Layout via D3 force-directed simulation, with manual override (drag nodes to reposition)
- Group nodes by provider (gcp, aws, azure) and by module
- Color-code by resource type (compute, network, storage, IAM)

**Cost estimation**
- For each cost-bearing resource, look up its monthly cost from a pricing table
- Show per-resource and total estimated cost
- Support GCP, AWS, and Azure (free tier vs paid)
- Cost data sourced from Infracost's open dataset (or, with permission, from their API)

**Blast radius**
- Click a node, highlight everything that transitively depends on it
- Toggle "what if this node fails" mode: show the failure cone in red
- Bonus: show "what services depend on this" by parsing service-account / IAM bindings

**Search and filter**
- Search by resource name, type, or tag
- Filter by module
- Filter to "only paid resources" or "only free-tier resources"

**Export**
- Save the current view as PNG or SVG
- Share via URL: paste the plan, get a unique URL that re-renders the same view

### Tech stack

- **Frontend**: SvelteKit, TypeScript, Tailwind v4, D3 (force layout)
- **Parser**: a Go service using `github.com/hashicorp/hcl/v2` for raw HCL parsing and Go's `encoding/json` for the plan-JSON path
- **Cost data**: a JSON file shipped with the app (lookup by `resource_type + region + size`), with optional Infracost integration for live pricing
- **State management**: client-side only — the plan is parsed and analyzed in-browser via WASM (the Go HCL parser compiled to WASM), so plans never leave the user's machine
- **Hosting**: Cloud Run for the SvelteKit app, no backend needed in v1 if the WASM approach works

### The privacy story

This is important and worth surfacing in the case study. Terraform plans contain real infrastructure details — bucket names, secret references, service account emails. Users will not paste them into a third-party tool unless they trust it. The WASM-in-browser approach means the plan never touches a server. This is both a tech decision and a marketing decision; lean into it.

### Why this is a good portfolio piece

It demonstrates: deep familiarity with the actual day job (parsing HCL well requires knowing HCL well), front-end ambition (D3 force layouts are a real frontend skill), thoughtful technical decisions (the WASM-in-browser privacy story is the kind of thing senior engineers notice), and a useful product instinct (this is a tool people would actually use, which is a higher bar than "this is a cool demo").

It also reinforces the portfolio's overall thesis: this person is a platform engineer who can also build delightful user-facing things. Terraplane sits at the intersection.

### Open design questions to settle when building

- HCL parser scope: full HCL? Just the `terraform plan -json` output? Both?
- How to handle modules with thousands of resources: lazy-render? Aggregate by module?
- Cost data freshness: ship a static JSON, refresh quarterly, vs. live API calls
- Blast radius: just direct dependencies, or transitively-computed?
- Multi-cloud: how to render a graph that spans GCP and AWS in the same plan
- Saved/shared views: requires a backend; can defer to v2

### Card preview (already built into portfolio)

The portfolio gallery shows a small SVG architecture graph (VPC → GKE/SQL → leaf nodes) with dashed amber connectors and a "$1,247/mo" cost label. Amber accent.

---

## Cross-project notes

### Subdomains and DNS

Each project gets its own subdomain so it can be deployed and scaled independently. The portfolio site at `ce.dev/work/[slug]` shows the case study and links out to the subdomain via the "live demo" button.

### Shared design language

All four projects should feel like siblings. They don't have to use the cotton-candy palette — each can have its own visual identity that suits the content (Diamond's deep-purple split-flap is right; cotton candy on Diamond would be wrong). But all four should feel like they were built by the same person:

- The same typography contrast (sans body + serif italic accents + mono microcopy)
- The same approach to motion (purposeful, never decorative)
- The same writing voice (plain English, dry humor, real numbers)
- The same component DNA (glass surfaces, hairline borders, sentence case)

Think of the portfolio as the "design system showcase" and each project as a "themed application of the system."

### Shared backend patterns

Three of the four projects (Linux Lessons, Bake-off, Diamond) need a Go backend on GCP. Standardize on:

- `chi` for the router
- `slog` for structured logging
- `pgxpool` for Postgres
- `chi/middleware` for the standard middleware chain
- Distroless containers
- Cloud Run for serving (except Bake-off, which is GKE for cold-start reasons)
- Cloud SQL Postgres
- Secret Manager for credentials
- Workload Identity Federation for CI

Reusing this stack across three projects makes each one cheaper to build (you've already debugged the deployment pipeline) and makes the case studies more focused (you spend the words on what's *different* about each project, not on rehashing common infra).

### Build order recommendation

Once the portfolio v1 ships, build the projects in this order:

1. **Backend Bake-off** — most direct showcase of platform skills, the design and infra are well-defined, and it's the project most likely to get traction in DevOps/SRE circles
2. **Terraplane** — natural follow-on, reuses the Go + WASM idea, highest perceived value to platform-engineering hiring managers
3. **Diamond Departures** — the most fun to build, but requires real-time data work and a stats provider integration
4. **Linux Lessons from Hell** — biggest scope, requires content design (the curriculum) which is its own multi-week effort

This order also matches a "low risk to high risk" curve — each project depends less on external content (curriculum design, sports data deals) than the next.

### What to do when a project ships

When a project goes live:
1. Update its row in the table at the top of this file from ⚪ to 🟢
2. Update `apps/web/src/lib/content/projects.ts` in the portfolio with the live URL
3. Replace the case-study page's "demo coming soon" placeholder with the actual live embed (or a high-quality screenshot if iframing isn't right)
4. Add a "Now live →" pill to the gallery card for that project
5. Write a short journal post about the launch (Phase 4 work in `TASKS.md` for the portfolio)
