---
slug: terraplane
title: Terraplane
summary: Paste a `terraform plan`, get a walkable architecture map with cost estimates and blast-radius highlights. Built so I never have to read raw plan output again.
role: solo — design, parser, frontend
stack:
  - SvelteKit
  - Go (HCL parser, compiled to WASM)
  - D3 force layout
  - Cloud Run
shipped: in progress
status: planned
liveDemoUrl: https://terraplane.cole-eckelberry.com
---

<script>
  import CSHero from '$lib/components/case-study/CSHero.svelte';
  import LiveDemo from '$lib/components/case-study/LiveDemo.svelte';
  import CSSection from '$lib/components/case-study/CSSection.svelte';
  import DecisionCard from '$lib/components/case-study/DecisionCard.svelte';
  import AdjacentProjects from '$lib/components/case-study/AdjacentProjects.svelte';
  import TerraplanePreview from '$lib/components/projects/previews/TerraplanePreview.svelte';
  import { PROJECT_BY_SLUG } from '$lib/content/projects';
  const project = PROJECT_BY_SLUG.terraplane;
</script>

<CSHero project={project} kind="A VISUALIZER" date="CONCEPT · 2026" role="solo — full stack" shipped="in progress" />

<LiveDemo url={project.liveDemoUrl} title="terraplane" subtitle="terraform plan, walkable" placeholder><TerraplanePreview /></LiveDemo>

<CSSection number="01" eyebrow="WHY" title="Most teams don't have an up-to-date" italic="architecture diagram." accent="amber" lead="They have a Terraform repo that *is* their architecture, but reading it requires spelunking through modules. This tool flips that — the graph becomes the document.">

Paste a `terraform plan -json` (or raw HCL with a fallback parser) and get an interactive, walkable graph of what will be created, what depends on what, what each thing will cost, and what breaks if you delete it. Click any node to see its source HCL, attributes, dependencies, blast radius. Toggle a "delete this" mode and the node lights up red along with everything that would die with it.

</CSSection>

<CSSection number="02" eyebrow="CAPABILITIES" title="Parse, layout," italic="cost, highlight." accent="amber" lead="Four reasons to open the tool. Each one already exists in fragments — terraform-graph, infracost, manual diagrams — but never in the same window.">

**Parse and graph.** `terraform plan -json` is the primary input; raw HCL is the fallback. The dependency graph is built from the resource references; the layout is a D3 force-directed simulation with manual override (drag to reposition). Nodes are grouped by provider (gcp, aws, azure) and by module, color-coded by resource family (compute, network, storage, IAM).

**Cost estimation.** Each cost-bearing resource gets a monthly cost from a pricing table. Per-resource and total estimated cost shown above the graph. Cost data is sourced from Infracost's open dataset, refreshed quarterly.

**Blast radius.** Click a node, highlight everything that transitively depends on it. Toggle "what if this fails" mode to show the failure cone in red. Bonus: parse service-account / IAM bindings so the radius can include "what services would lose access."

**Search and filter.** Search by name, type, or tag. Filter by module, by provider, by paid-vs-free-tier. Export the current view as PNG or SVG. Share via URL — paste the plan, get a unique URL that re-renders the same view, signed by a content hash.

</CSSection>

<CSSection number="03" eyebrow="PRIVACY" title="The plan never" italic="leaves your machine." accent="amber" lead="Terraform plans contain real infrastructure details — bucket names, secret references, service-account emails. Engineers will not paste them into a third-party tool unless they trust it.">

The HCL parser is written in Go and compiled to WASM. The browser loads the WASM module once, then parses every plan locally — no upload, no server-side processing, no logs. The "share" feature works by encrypting the plan with a key in the URL hash; the server only ever sees the ciphertext.

This is both a tech decision and a marketing decision. Privacy is the headline on the landing page because it has to be, otherwise the right people will never paste real plans in.

</CSSection>

<CSSection number="04" eyebrow="DECISIONS" title="Things I had to" italic="pick a side on." accent="amber">

<div class="decisions-grid">

<DecisionCard question="HCL parser scope?" pickedLabel="picked plan-json + HCL" pickedAccent="amber">
`terraform plan -json` is the cleanest input — already resolved, no module spelunking, all dependencies pre-computed. Raw HCL is the fallback for users who can't run plan (read-only credentials, plan-time errors). Both paths flow into the same internal graph representation.
</DecisionCard>

<DecisionCard question="WASM in browser, or Go service?" pickedLabel="picked WASM" pickedAccent="amber">
A Go service would be easier to maintain. WASM-in-browser is materially better for trust — and the parser doesn't need anything that WASM can't do. Bundle size is ~3 MB which is fine for a tool people will open occasionally rather than every page load.
</DecisionCard>

<DecisionCard question="Cost data: shipped, live, or both?" pickedLabel="picked shipped quarterly" pickedAccent="amber">
A static JSON file refreshed quarterly is fast (no network on render), works offline, and is good enough for the "rough estimate" that's the actual user need. Live API calls would be more accurate but turn a private tool into a network-dependent one. Quarterly is the right tempo.
</DecisionCard>

<DecisionCard question="Multi-cloud rendering?" pickedLabel="picked unified graph" pickedAccent="amber">
A plan that spans GCP and AWS is rare but real. The graph treats them as one network with provider-colored node clusters; the force layout keeps each provider's resources gravitating together. Cleaner than rendering two separate graphs.
</DecisionCard>

</div>

</CSSection>

<CSSection number="05" eyebrow="OPEN QUESTIONS" title="What I'm still" italic="figuring out." accent="amber">

- **Module aggregation at scale.** A plan with thousands of resources won't render meaningfully. Lazy-render below a threshold? Aggregate by module first, then drill in? Probably both.
- **Saved views.** Beyond the share-URL — a logged-in surface where you can save and re-open plans, diff them across days. That requires accounts, which I don't want at v1.
- **Diffing two plans.** A higher-value feature than viewing one — show me what changed between this morning's plan and now's. Real work, but the right next thing after v1.

</CSSection>

<AdjacentProjects current="terraplane" />

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
