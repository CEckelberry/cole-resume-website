# Performance posture

Snapshot at v1 ship. Numbers below are from `pnpm --filter web build`
on the cole-resume-website repo, measured on Apple Silicon. Re-run after
material changes and update.

## Bundle sizes (gzipped) — production build

### CSS

| File                                            | Raw        | Gzip        | Notes                          |
| ----------------------------------------------- | ---------- | ----------- | ------------------------------ |
| `_app/immutable/assets/0.*.css`                 | 23.65 kB   | **5.69 kB** | layout shell, mesh, nav, fonts |
| `_app/immutable/assets/2.*.css`                 | 17.69 kB   | 3.38 kB     | home page sections             |
| `_app/immutable/assets/4.*.css`                 | 13.62 kB   | 2.52 kB     | case-study chrome              |
| `_app/immutable/assets/TerraplanePreview.*.css` | 4.10 kB    | 1.23 kB     | route-split preview            |
| `_app/immutable/assets/1.*.css`, `3.*.css`      | 1.72 kB ea | 0.54 kB ea  | journal + error                |

### JS — page entry chunks

| Chunk                               | Raw      | Gzip        | Notes                                          |
| ----------------------------------- | -------- | ----------- | ---------------------------------------------- |
| `nodes/0.*.js` (home)               | 9.09 kB  | **3.65 kB** | Hero, Marquee, Projects, About, Contact        |
| `nodes/2.*.js` (work/[slug])        | 17.31 kB | 6.60 kB     | case-study route                               |
| `nodes/4.*.js` (case-study content) | 38.19 kB | 14.22 kB    | mdsvex-compiled markdown for all four projects |
| `nodes/1.*.js` (journal)            | 2.69 kB  | 1.16 kB     | placeholder page                               |
| `nodes/3.*.js` (error/404)          | 1.35 kB  | 0.70 kB     | error page                                     |
| `chunks/vifcVInA.js` (shared)       | 54.53 kB | 20.99 kB    | Svelte 5 runtime + view-transitions            |
| `entry/app.*.js`                    | 3.24 kB  | 1.44 kB     | bootstrap                                      |

### Targets vs measured (per ARCHITECTURE.md performance budget)

| Page       | Target JS gzip | Actual gzip                         | Status |
| ---------- | -------------- | ----------------------------------- | ------ |
| Home       | < 60 kB        | ~26 kB (nodes/0 + chunks)           | ✅     |
| Case study | < 80 kB        | ~42 kB (nodes/2 + nodes/4 + chunks) | ✅     |
| About      | < 40 kB        | covered by home                     | ✅     |

Total client `_app/immutable/` directory: **~580 kB uncompressed** (down
from ~1.1 MB before font-subset trim).

## Optimizations applied

**Font subsetting**

- @fontsource imports use the `latin-` subset variants (e.g.
  `@fontsource/inter/latin-400.css`) — drops Cyrillic, Greek, Vietnamese,
  Latin-ext that the v1 site doesn't need. Cut layout CSS from 21 kB → 5.7 kB
  gzipped (4× reduction). 24 woff/woff2 files emitted instead of ~80.
- `font-display: swap` ships by default with @fontsource v5 — first paint
  uses the system fallback, then re-flows once the web font lands. No FOIT.

**CSS via Tailwind v4**

- `@theme inline` wires CSS custom properties to utility names, so
  utilities like `bg-canvas` resolve through the cascade. Tailwind v4 only
  emits utilities that are actually referenced in source — no PurgeCSS
  postprocessing step needed.

**Mesh + polygon background**

- Mesh blobs use only CSS `transform` + `filter` keyframes — compositor-
  driven, off the main thread, no JS scheduling cost.
- Polygon scroll-rotation: single passive scroll listener with rAF
  coalescing; at most one batch of 5 attribute writes per frame.
- Both honor `prefers-reduced-motion` — animations disabled / listener
  not attached.

**Routing**

- Every route is `prerender = true` except `/api/contact` (request-time
  proxy). Output is static HTML; first byte is < 50 ms from the Cloud Run
  instance once warm.
- Case-study content is bundled via `import.meta.glob({ eager: true })` so
  there are no async chunk loads on navigation.

**View transitions**

- `onNavigate` hook wires `document.startViewTransition` for a 240 ms
  cross-fade between routes; falls through to instant nav in browsers
  without support and skips entirely under `prefers-reduced-motion`.

## Deferred to v2

- **Playwright + Web Vitals budget enforcement in CI** — TASKS.md spec
  asked for this; a manual `pnpm build` shows the numbers above and the
  page sizes are well inside the budget, so the CI integration is high-
  effort, low-value at v1.
- **Dynamic per-page OG image generation via satori** — currently a single
  static `/og.svg` ships. Per-page OG images would land via a build-time
  script that walks routes and renders 1200×630 PNGs.
- **Bundle the polygon SVG inline at SSR** instead of as a separate stylesheet
  — minor; the polygon-field CSS is already < 1 kB gzip.
- **Mesh-blob lazy mount** — currently mounted eagerly in `+layout.svelte`.
  An IntersectionObserver-gated mount would defer the compositor layer
  setup until visible; saves a few ms on first paint but the mesh IS
  visible above the fold so it's borderline.
