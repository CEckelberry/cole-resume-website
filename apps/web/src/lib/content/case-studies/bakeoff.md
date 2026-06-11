---
slug: bakeoff
title: Backend
italicWord: Bake-off
summary: Same checkout endpoint, six runtimes, hot-swappable mid-session. The point isn't to crown a winner — it's to give engineers a number with a unit on it the next time someone says "should we rewrite this in X."
role: solo — design, infra, all six implementations
stack:
  - Go
  - Rust
  - Bun
  - Node.js
  - Python
  - PHP
shipped: 2026
status: live
liveDemoUrl: https://backend-bakeoff.com
---

<script>
  import CSHero from '$lib/components/case-study/CSHero.svelte';
  import LiveDemo from '$lib/components/case-study/LiveDemo.svelte';
  import CSSection from '$lib/components/case-study/CSSection.svelte';
  import DecisionCard from '$lib/components/case-study/DecisionCard.svelte';
  import AdjacentProjects from '$lib/components/case-study/AdjacentProjects.svelte';
  import BakeoffPreview from '$lib/components/projects/previews/BakeoffPreview.svelte';
  import { PROJECT_BY_SLUG } from '$lib/content/projects';
  const project = PROJECT_BY_SLUG.bakeoff;
</script>

<CSHero project={project} kind="A TOOL" date="CONCEPT · 2026" role="solo — full stack" shipped="in progress" />

<LiveDemo url={project.liveDemoUrl} title="bakeoff" subtitle="six runtimes, one checkout endpoint" placeholder><BakeoffPreview /></LiveDemo>

<CSSection number="01" eyebrow="WHY" title="Every backend benchmark you've read was written by" italic="someone with a horse in the race." accent="teal" lead="The numbers aren't wrong. The problem is you can't see them move under workloads you actually care about.">

The site is one frontend — a small e-commerce checkout flow — backed by six independent implementations of the same `POST /checkout` endpoint. Pick which runtime serves your requests via a tab. The page reloads, the runtime changes, and you watch p95, RPS, and cold-start metrics shift in real time.

The point isn't to crown a winner. The point is to let people change variables and see what happens — so the next "should we rewrite this in X" conversation has a number with a unit on it instead of vibes.

</CSSection>

<CSSection number="02" eyebrow="THE LINEUP" title="Six runtimes, one" italic="contract." accent="teal" lead="Same OpenAPI schema, same Postgres, same fault budget. The only thing that varies is the language model under the hood.">

| Runtime | Framework | Why it's here |
|---|---|---|
| Go | `chi` + stdlib | The control implementation |
| Rust | `axum` + `tokio` | Speed-of-light for this workload |
| Bun | `hono` | The new entrant |
| Node | `express` | The baseline most teams actually run |
| Python | `fastapi` | Async Python, what most data orgs ship |
| PHP | `laravel` | The unfair fight that surprises people |

Deliberately not on the list: JVM (cold-start distortion needs too many footnotes), Deno (too similar to Bun for the sample size to matter).

The endpoint itself is small on purpose. Accepts a cart of 3–8 items, validates inventory, computes tax via a fake remote service with simulated latency, runs a fraud score (50 ms of pure CPU), persists the order to Postgres, returns 200 with the order id. Codegen produces server stubs and client types in every language so the implementations can't drift from the schema.

</CSSection>

<CSSection number="03" eyebrow="ARCHITECTURE" title="One cluster, six pods," italic="a header to choose." accent="teal">

The frontend calls a single `POST /api/checkout` on a router service. The router (a small Go service) reads `X-Runtime: rust|go|bun|node|python|php` and forwards to the matching backend pod via Kubernetes' headless service DNS. Switching runtime is a header change in the frontend — no redeploy, no env-var toggle.

```
Browser → Router (Go) → bo-{runtime} pod → Postgres
                                  ↓
                            Prometheus scrape
```

All six pods run continuously, all backed by the same Cloud SQL Postgres. Prometheus scrapes them on the same interval. The frontend pulls live metrics from `/api/metrics` (proxied through the router from Prometheus' HTTP API) and renders bar charts that update every two seconds.

The site has three modes:

- **Casual** — place orders, watch the metrics shift. No scenario, just play.
- **Comparison** — pick two runtimes; the page splits and runs the same request to both side-by-side. Good for "is Rust really 2x Go on this workload?"
- **Stress** — a slider from 1 to 200 RPS. The page generates load (rate-limited per visitor IP) and shows how each runtime degrades. P95, p99, error rate.

</CSSection>

<CSSection number="04" eyebrow="DECISIONS" title="Things I had to" italic="pick a side on." accent="teal">

<div class="decisions-grid">

<DecisionCard question="Why GKE instead of Cloud Run?" pickedLabel="picked GKE" pickedAccent="teal">
Cloud Run isolates services into separate execution environments and scales to zero. Both are wrong for a benchmark that wants the runtimes to share node-level resource pressure and stay warm. GKE Autopilot gives me one cluster with mixed-runtime pods scheduled together — the comparison is honest, and I get to keep the operational simplicity of managed nodes.
</DecisionCard>

<DecisionCard question="Shared Postgres or one per runtime?" pickedLabel="picked shared" pickedAccent="teal">
A shared Postgres is fairer; the variance lives in the runtimes, not in the database. An isolated Postgres per runtime is more realistic for production deployments but introduces noise (different cache states, different connection-pool warmups). The benchmark only works if Postgres is a held-constant — so shared.
</DecisionCard>

<DecisionCard question="Why include PHP?" pickedLabel="picked PHP" pickedAccent="teal">
Because the actual answer is a number, and that's what this whole project is for. I keep meeting senior engineers who assume PHP-FPM is the slowest possible thing, and others who assume it's surprisingly fast, and both are wrong. Frankenphp on persistent workers is the v1 entry; cgi-mode would be a v2 footnote.
</DecisionCard>

<DecisionCard question="Why no JVM?" pickedLabel="skipped JVM" pickedAccent="pink">
JVM startup distorts the cold-start tab badly enough to need its own asterisks. The chart already has six rows; adding a seventh that needs a footnote longer than the row makes the chart worse. If JVM ends up requested often enough I'll add a separate "warmth-sensitive" board.
</DecisionCard>

</div>

</CSSection>

<CSSection number="05" eyebrow="OPEN QUESTIONS" title="What I haven't" italic="decided yet." accent="teal">

The architecture is settled. The product surface still has rough edges:

- **Cost ceiling.** GKE Autopilot is not free. Six pods running 24/7 plus Cloud SQL is roughly $50–80/month at idle. Worth it for the demo to be live? Probably yes; the alternative is a cron job that runs a synthetic workload and posts the results, which is much less interesting.
- **E-commerce surface.** Just `/checkout`, or also `/cart/add` and `/products/list`? More endpoints means more variables. Starting with just `/checkout`.
- **Stress-mode rate limits.** A single visitor cannot be allowed to spend the whole month's Cloud SQL quota in an afternoon. Per-IP token bucket, capped to 50 RPS for 60 seconds, then cooldown.
- **Where the frontend lives.** Same cluster as the backends, or its own Cloud Run service? The latter — keeps the cluster scoped to the experiment.

</CSSection>

<AdjacentProjects current="bakeoff" />

<style>
  .decisions-grid {
    display: grid;
    grid-template-columns: 1fr;
    gap: var(--space-4);
  }
  @media (min-width: 720px) {
    .decisions-grid {
      grid-template-columns: 1fr 1fr;
    }
  }
</style>
