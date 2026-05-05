# Design language

This document defines the visual and interaction language of the portfolio site. If `ARCHITECTURE.md` is the skeleton, this is the skin and the way it moves. When in doubt, this file overrides taste.

## North star

A site that feels like it was built by someone who cares — about the user's time, the user's eyes, and the small details that make a deployment feel hand-finished. **Cotton candy** is the palette but **brutalist confidence** is the layout. Saturated, playful, glassy, fast. Never twee. Never overdesigned. Never apologetic about being colorful.

The acid test: a recruiter should immediately see the four projects. A senior engineer should immediately see that the person who built this knows what they're doing. Both should feel they're looking at one cohesive piece of work, not a portfolio template.

## Palette

The cotton-candy palette is the only palette at v1. The structure leaves room for additional palettes (midnight, sunset, mint, buttercream) in v2 — every color reference should go through a CSS variable so swapping palettes is a one-token change.

### Cotton candy palette — base colors

| Token | Light mode | Dark mode | Use |
|---|---|---|---|
| `--bg-canvas` | `#fefcf8` | `#0e0820` | Page background |
| `--bg-surface` | `#faf7ff` | `#1a0f3a` | Raised surface, cards |
| `--bg-surface-2` | `rgba(255,255,255,0.6)` | `rgba(255,255,255,0.06)` | Glass surface |
| `--bg-surface-3` | `rgba(255,255,255,0.85)` | `rgba(255,255,255,0.1)` | Strong glass |
| `--text-primary` | `#1a1322` | `#ffffff` | Body text |
| `--text-secondary` | `#444441` | `rgba(255,255,255,0.78)` | Subtitles, descriptions |
| `--text-tertiary` | `#888780` | `rgba(255,255,255,0.55)` | Microcopy, metadata |
| `--text-muted` | `#B4B2A9` | `rgba(255,255,255,0.4)` | Placeholders, hints |
| `--border-subtle` | `rgba(0,0,0,0.06)` | `rgba(255,255,255,0.12)` | Hairlines |
| `--border-default` | `rgba(0,0,0,0.12)` | `rgba(255,255,255,0.18)` | Card borders |
| `--border-strong` | `rgba(0,0,0,0.2)` | `rgba(255,255,255,0.32)` | Hover, focus |

### Cotton candy palette — accents

| Token | Hex | Use |
|---|---|---|
| `--accent-pink` | `#ED93B1` | Primary accent — links, "live" indicators, project 01 |
| `--accent-pink-soft` | `#F4C0D1` | Pink fills, underline highlights |
| `--accent-purple` | `#7F77DD` | Secondary accent, project 03 |
| `--accent-purple-soft` | `#CECBF6` | Purple fills |
| `--accent-teal` | `#5DCAA5` | Success, "live" status, project 02 |
| `--accent-teal-soft` | `#9FE1CB` | Teal fills |
| `--accent-amber` | `#FAC775` | Warm accent, project 04 |
| `--accent-amber-soft` | `#FAEEDA` | Amber fills |
| `--accent-coral` | `#D85A30` | Reserved for warnings/destructive only |

Each project has a "home" accent: pink for Linux, teal for Bake-off, purple for Diamond, amber for Terraplane. This is consistent everywhere — card borders on hover, code highlighting in case studies, the gradient direction of mesh blobs near that project's section.

### Mesh blobs (background)

Soft, blurred, drifting circular gradients positioned absolutely behind content. Always at least one per visible viewport-height of scroll, never more than four. Blur radius 60–80px, opacity 0.35–0.55. Animation: 16–22s ease-in-out infinite, translating ±40px and scaling 0.95–1.10. Each section has its own blob configuration so the ambient color shifts as you scroll.

```css
.blob {
  position: absolute;
  border-radius: 50%;
  filter: blur(72px);
  opacity: 0.45;
  will-change: transform;
}
```

### Polygon decorations

Sparse polygonal shapes scattered behind the mesh, each filled with an accent color at low opacity (0.14–0.35). They do not animate at v1 (scroll-tied animation is a v2 enhancement). Their job is to add geometric texture against the soft blobs. 4–6 per viewport, all rendered in a single inline SVG per section.

## Typography

Three families, used with discipline:

| Family | Source | Use |
|---|---|---|
| `Inter` | Google Fonts, weights 400/500 | Default sans — body, UI, navigation |
| `Instrument Serif` | Google Fonts, italic 400 | Editorial accents — pull quotes, project name italics, "deck" subtitles |
| `JetBrains Mono` | Google Fonts, weights 400/500 | Code, microcopy, eyebrow labels, status text |

Two weights in sans: 400 regular, 500 medium. Never 600 or 700 — they read as too heavy against the soft palette. Headlines use 500. Body uses 400. There is no bold-italic.

Sentence case everywhere. Never Title Case. Never ALL CAPS, except for the eyebrow labels and microcopy in JetBrains Mono with `letter-spacing: 0.12em` to 0.18em — that's the one place uppercase is correct.

### Scale

| Token | Size | Line height | Use |
|---|---|---|---|
| `--type-display` | clamp(40px, 8vw, 76px) | 0.92 | Hero name |
| `--type-h1` | clamp(36px, 6vw, 56px) | 0.98 | Case study title |
| `--type-h2` | clamp(28px, 4vw, 38px) | 1.05 | Section titles |
| `--type-h3` | 22px | 1.2 | Card titles |
| `--type-lead` | 20px | 1.4 | Hero deck, section lead-in |
| `--type-body` | 14px | 1.65 | Default prose |
| `--type-body-sm` | 13px | 1.5 | Secondary prose |
| `--type-meta` | 12px | 1.5 | Captions, metadata |
| `--type-micro` | 11px | 1.4 | Eyebrows, status pills (mono, uppercase) |
| `--type-tiny` | 10px | 1.4 | Stack tags, footer microcopy |
| `--type-nano` | 9px | 1.3 | Tertiary microcopy only |

Hero name uses `clamp` for fluid sizing. The full last name "Eckelberry" must fit on one line at every viewport. Test at 320px (smallest realistic mobile) and 1920px (large desktop).

### Italic accents

A subset of words inside hero, section titles, and case-study titles renders in `Instrument Serif italic` and a project accent color. Use sparingly — at most one italic phrase per heading. Examples:
- "Things I built when I should've been *sleeping.*"
- "The Backend *Bake-off.*"
- "Five flavors. Pick the one you can *stand.*"

Implementation:
```html
<h2>Things I built when I should've been <em>sleeping.</em></h2>
```

```css
h1 em, h2 em, h3 em {
  font-family: 'Instrument Serif', serif;
  font-style: italic;
  font-weight: 400;
  color: var(--accent-pink-soft);
}
```

## Spacing and layout

8px base unit. Vertical rhythm uses `rem`-based spacing for relationships between blocks; component-internal gaps use pixels.

| Token | Value | Use |
|---|---|---|
| `--space-1` | 4px | Tight inline gaps |
| `--space-2` | 8px | Default small gap |
| `--space-3` | 12px | Component padding (small) |
| `--space-4` | 16px | Component padding (default) |
| `--space-5` | 22px | Card padding |
| `--space-6` | 32px | Between blocks |
| `--space-7` | 44px | Between sections |
| `--space-8` | 72px | Major section breaks |

Container widths:
- `max-width: 1200px` for the page shell
- `max-width: 720px` for case-study prose
- `max-width: 560px` for hero deck text and section leads

Outer page padding: `24px` on mobile, `48px` on desktop. The mesh background extends to the viewport edge; content respects the padding.

## Radii and borders

| Token | Value | Use |
|---|---|---|
| `--radius-sm` | 4px | Pill tags, small buttons |
| `--radius-md` | 8px | Inputs, default surfaces |
| `--radius-lg` | 12px | Cards |
| `--radius-xl` | 16px | Large cards, hero terminal |
| `--radius-pill` | 9999px | Nav, filter chips, status pills |

Borders are `0.5px` for hairlines, `1px` only when the border carries semantic weight (selected state, focus). No `2px` borders except the "featured" project accent (which we don't use at v1).

## Motion

Motion in this site has a job. It is not decoration. The job is one of: signal hierarchy, signal interaction state, or signal liveness.

### Easing curves

| Token | Curve | Use |
|---|---|---|
| `--ease-out` | `cubic-bezier(0.2, 0.8, 0.2, 1)` | Default for most transitions |
| `--ease-in-out` | `cubic-bezier(0.4, 0, 0.2, 1)` | Drift animations, mesh blobs |
| `--ease-spring` | `cubic-bezier(0.34, 1.56, 0.64, 1)` | Magnetic pulls, playful interactions |

### Durations

| Token | Value | Use |
|---|---|---|
| `--dur-fast` | 150ms | Hover state, focus ring |
| `--dur-default` | 240ms | Card hover lift, color transitions |
| `--dur-slow` | 400ms | Page transitions, reveal animations |
| `--dur-drift` | 18s–22s | Background mesh blobs |

### Specific interactions

**Card hover.** Project cards lift `translateY(-3px)` and brighten the border to `--border-strong` over `--dur-default` `--ease-out`. The accent gradient glyph inside the card shifts position by 6px on cursor parallax (capped, not 1:1 with mouse).

**Magnetic buttons.** Primary buttons pull toward the cursor when within 60px, max displacement 8px, `--ease-spring`. Disable on touch devices.

**Hero terminal.** Boots on first paint with a 1.6s typing animation (`whoami`, then `cat now.txt`, then settle to a blinking cursor). The blinking cursor is a CSS `step(2)` animation at 1s, no JS.

**Stack marquee.** Translates `-50%` over 30s linear, infinite. On hover, slows to 60s instead of stopping outright — it should never stop.

**Mesh blobs.** Each blob has its own keyframe with translate and scale variations. Stagger their timings so they never sync.

**Scroll reveal.** Sections fade in and translate up 12px when their top crosses the 80%-from-top mark, observed via `IntersectionObserver`. Only animate once. Honor `prefers-reduced-motion`.

**Page transitions.** Use SvelteKit view transitions API. From the projects gallery to a case study, the clicked card morphs into the case-study hero (shared element via `view-transition-name`). Other transitions are a 240ms cross-fade.

### Reduced motion

`@media (prefers-reduced-motion: reduce)` disables: mesh blob drift, marquee, scroll reveals, magnetic buttons, page transitions. The terminal still boots but skips to the final state. The blinking cursor stops.

## Components

This section catalogs reusable components by feature. Each entry lists its purpose, anatomy, states, and design rules.

### Nav (top of viewport)

A pill-shaped glass bar that floats at the top of the page with horizontal margin. Contains:
- Left: a live status pill ("deployed 2h ago · main@a3f1c2e")
- Center: nav links (work, about, writing, contact)
- Right: nothing at v1 (theme toggle moves to the about section)

Glass background: `var(--bg-surface-2)` with `backdrop-filter: blur(16px)`. Border: `0.5px solid var(--border-default)`. Padding: `8px 14px`. Radius: `--radius-pill`.

The "deployed N ago" timestamp is real — pulled from `git log -1` at build time and inlined into the page.

States: nav links have a hover state (background `var(--bg-surface-3)`) and an active state (same, plus full opacity text). The active state is set per page.

### Hero terminal

A miniature terminal window that boots on first paint. ~280px wide on desktop, fits below the hero deck on mobile.

Anatomy:
- Title bar: three traffic-light dots (red, yellow, green), title text "~/cole — zsh"
- Body: monospace 11px, line-height 1.7
- Boot sequence: prompt, `whoami`, output, prompt, `cat now.txt`, key-value output, prompt, blinking cursor

The boot sequence is a typed-character animation. Each character appears with a 30ms delay. Total boot time ~1.6s. After boot, the cursor blinks indefinitely. No further interaction.

Color rules:
- Background: `rgba(8,4,18,0.7)` with backdrop blur
- Prompt: `var(--accent-teal)`
- Command text: white
- Output: `rgba(255,255,255,0.7)`
- Keys (e.g. `role`): `var(--accent-pink-soft)`
- Values (e.g. `"DevOps lead"`): `var(--accent-amber-soft)`

### Project card

A glass tile in the projects gallery. ~330px tall, full-width within its grid column. Two parts:
- **Preview** (top, 160px): a custom-rendered visual unique to the project (terminal output for Linux, latency bars for Bake-off, split-flap rows for Diamond, architecture graph for Terraplane)
- **Body** (bottom): project number, title, description, tech tags, footer with status and "case study" link

Hover: `translateY(-3px)`, border brightens, the preview region's contents shift slightly via parallax tied to cursor position relative to the card center.

The preview is the differentiator. A generic colored gradient as a preview would make the cards interchangeable. Each preview is a small piece of *the actual project surfaced as a UI fragment*.

### Filter chips

Pill-shaped buttons above the gallery. Inactive state: glass background, secondary text. Active state: white background, dark text. Filters at v1: All / Games / Tools / Data / Visualizers. Clicking filters the gallery client-side with a 200ms cross-fade between visible/hidden.

### Stack marquee

Below the hero, a horizontal scrolling band of tech stack items separated by diamond dividers (◇). Some items have parenthetical italic asides ("Python *(my main)*"). Edges fade to the page background using horizontal gradient masks.

Items at v1, in order:
- SvelteKit
- Python *(my main)*
- TypeScript
- Kubernetes
- Terraform
- PostgreSQL
- GCP
- Docker
- Bash *(since 2012)*

Note: Go and Ruby do not appear in the marquee. Go appears in individual project tech tags only (since several projects use it); Ruby does not appear at all on the site.

### Case study hero

The top of every `/work/[slug]` page. Anatomy:
- Eyebrow: "PROJECT 02 · A TOOL · MAR 2026"
- Title: large display type, with one italicized word in the project's accent color
- Deck: serif italic, max 540px wide
- Meta row: 4 stats (Role / Stack / Shipped / Status) separated by a top and bottom hairline
- Live demo embed: a glass card with a fake browser bar showing the project URL, then the embedded mini-demo

The mini-demo is mandatory. If a project does not have an embeddable demo, the case study isn't ready to ship.

### Section header

Each numbered section in a case study starts with:
- An eyebrow ("01 · WHY") with a leading hairline
- A section title (h2) with one italicized phrase
- A serif-italic lead sentence (max 560px)
- The section body in standard prose

### Code block

Used for code excerpts in case studies. Anatomy:
- Top bar: traffic lights, file path with language pill, line count on the right
- Body: numbered lines, syntax-highlighted, a small range of lines highlighted with a left border in the project accent color
- Bottom annotation: a soft-tinted strip with a "NOTE" pill explaining the highlighted lines

Syntax highlighting palette:
- Keywords: `var(--accent-pink)` 
- Functions: `var(--accent-amber)`
- Strings: `var(--accent-teal-soft)`
- Comments: `rgba(255,255,255,0.4)` italic
- Types: `var(--accent-purple-soft)`
- Default: `rgba(255,255,255,0.85)`

### Architecture diagram

Inline SVG, styled to match the dark mode (or light mode) page. Boxes use the project's accent color family at the appropriate stop. Connectors are thin (`stroke-width: 0.8`), often dashed, with simple chevron arrowheads.

Each case study has at least one architecture diagram. They are hand-built per case study (not generated) — they're a place to show the system thinking.

### Decision card

Used in case studies to surface tradeoffs. A small glass card with:
- A serif-italic question in the project accent color ("Why GKE and not Cloud Run?")
- A short prose answer (2–3 sentences)
- A pill at the bottom: "picked GKE" (teal) or "skipped JVM" (pink)

Two-column grid on desktop, single column on mobile.

### Lessons row

Two columns, each containing 1–2 cards with a left-border accent (no full border). Two flavors:
- "WORKED" with a teal left border
- "DIDN'T" with a pink left border

The contrast makes scanning easy.

### Footer

Four-column grid on desktop, single column on mobile. Columns:
1. Brand block: name, blurb, uptime status with a pulsing teal dot
2. WORK: links to each case study
3. ELSEWHERE: github, linkedin, writing
4. CONTACT: email, cv.pdf, calendar booking link

Bottom row (full width): build info on the left ("built with sveltekit + go · cloud run · cloud cdn"), copyright on the right ("©2026 — made in Bay Area").

## Voice and tone

Written content on the site has a voice. Reference these patterns when writing or rewriting copy:

**Confident, not boastful.** "Migrated 60+ apps to GCP" not "Successfully orchestrated the migration of over sixty enterprise applications". The numbers do the work.

**Plain English, with occasional dry humor.** "PHP *(reluctantly)*" in the stack marquee. "JVM startup distorts the cold-start tab badly enough to need its own asterisks. Not worth the footnote." The humor only lands when the surrounding copy is straight.

**Specific over vague.** "31 lines of Go" not "a small Go service". "Updates every 30s" not "live updates". "p95 latency 12ms" not "very fast".

**No marketing speak ever.** Words to avoid: leverage, synergy, empower, robust, world-class, cutting-edge, solutions, ecosystem, unlock, enable. If a sentence works without an adverb, drop the adverb.

**Sentences end with periods, including in microcopy.** "Live demo." not "Live demo". The exception is mono-uppercase eyebrows ("SECTION 02 · WORK").

**Personal pronouns.** First person ("I built", "I shipped") in case studies and bio. Site copy in second person ("you can hot-swap", "watch the latency move"). Never "we" — there is no we.

## Accessibility

Mandatory at v1, not v2:

- Color contrast: 4.5:1 for body text, 3:1 for large text and UI components
- Focus rings: visible, 2px outline in `var(--accent-pink)` with 2px offset, on every interactive element
- Keyboard navigation: every interactive element reachable via tab, in a logical order
- Skip-to-content link at the top of every page (visually hidden until focused)
- Semantic HTML: `<nav>`, `<main>`, `<article>`, `<section>`, `<footer>`. Headings in order, no skipped levels.
- ARIA labels on icon-only buttons and on the mesh background SVGs (`aria-hidden="true"`)
- `prefers-reduced-motion` respected (see Motion section)
- `prefers-color-scheme` respected for default mode
- Form fields with proper labels and error states
- The terminal animation has a `aria-live="off"` and an off-screen text alternative

Lighthouse Accessibility score must be 100 on every page.

## Do / don't quick reference

**Do:**
- Use sentence case for every heading
- Use `Instrument Serif italic` for at most one phrase per heading
- Pair the mesh background with the polygon decorations — never one without the other
- Keep card content layered visually (preview → body, never one flat block)
- Use the project accent color consistently across that project's surfaces
- Make all glass surfaces use `backdrop-filter: blur()` (with a fallback solid)
- Put real numbers in copy ("60+ apps", "p95 12ms")

**Don't:**
- Use bold weights heavier than 500
- Title Case anything
- Add drop shadows to surfaces (use blurred mesh + glass instead)
- Animate without a reason
- Use stock photography or stock illustrations
- Add gradients to text outside the hero name (it's overused; once is enough)
- Use emoji in any UI surface (custom SVG icons or shape-based glyphs only)
- Make a card without a unique preview region
- Forget the focus ring
