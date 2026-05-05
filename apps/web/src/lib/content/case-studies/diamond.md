---
slug: diamond
title: Diamond
italicWord: Departures
summary: Top 100 active MLB players ranked by sabermetrics, styled as a Penn Station split-flap board, updated live during games.
role: solo — design, frontend, ingest pipeline
stack:
  - SvelteKit
  - Go SSE
  - Postgres
  - FanGraphs ingest
shipped: in progress
status: planned
liveDemoUrl: https://diamond.cole-eckelberry.com
---

<script>
  import CSHero from '$lib/components/case-study/CSHero.svelte';
  import LiveDemo from '$lib/components/case-study/LiveDemo.svelte';
  import CSSection from '$lib/components/case-study/CSSection.svelte';
  import DecisionCard from '$lib/components/case-study/DecisionCard.svelte';
  import AdjacentProjects from '$lib/components/case-study/AdjacentProjects.svelte';
  import DiamondPreview from '$lib/components/projects/previews/DiamondPreview.svelte';
  import { PROJECT_BY_SLUG } from '$lib/content/projects';
  const project = PROJECT_BY_SLUG.diamond;
</script>

<CSHero project={project} kind="A DATA TOY" date="CONCEPT · 2026" role="solo — full stack + design" shipped="in progress" />

<LiveDemo url={project.liveDemoUrl} title="diamond departures" subtitle="penn station for sabermetrics" placeholder><DiamondPreview /></LiveDemo>

<CSSection number="01" eyebrow="WHY" title="The split-flap board is one of the great" italic="industrial typefaces." accent="purple" lead="Penn Station, Frankfurt Hauptbahnhof — there's a specific aesthetic where letters and numbers physically rotate when they change. It feels alive in a way an LCD never does.">

A leaderboard of the top 100 active MLB players, ranked by advanced sabermetrics, styled like that, updating live as games progress. Mid-game: a player gets a hit, his wRC+ shifts, his rank changes, the affected rows physically flip to their new values. Players climbing get a green up-arrow; players dropping get a red down-arrow.

It's a sports nerd's dream and a UI animation excuse all at once. Nobody has built this. I want it to exist.

</CSSection>

<CSSection number="02" eyebrow="WHAT IT SHOWS" title="wRC+ on the front," italic="every other stat one tab away." accent="purple" lead="The default landing tab is wRC+ for hitters — the single best 'who is the best hitter in baseball right now' stat. Every other category is a tab.">

**Hitters.** AVG, OBP, SLG, OPS, wRC+, wOBA, BABIP, ISO, fWAR, plus the traditionals (TB, RBI, R) for the people who still ask for them.

**Pitchers.** ERA, FIP, xFIP, SIERA, ERA+, K/9, BB/9, K-BB%, WHIP, fWAR.

**Defensive.** DRS by position, UZR/150 for outfielders and infielders, OAA where available.

**Position views.** A tab per position (C, 1B, 2B, 3B, SS, LF, CF, RF, DH, SP, RP) showing the top players at that position by the most-relevant stat (wRC+ for hitters, FIP for pitchers, framing runs for catchers).

</CSSection>

<CSSection number="03" eyebrow="THE BOARD" title="Every cell is" italic="a flap." accent="purple" lead="Amber monospace on deep purple. When a value changes, the cell rotates upward, exits the top, and the new value rotates in from the bottom. ~600 ms per flip; subtle 'click click click' SFX, off by default.">

When a player changes rank, the entire row slides up or down to its new position with the same flap animation on every column. Multiple simultaneous changes (a player got a hit AND BABIP went up AND wRC+ went up AND rank changed) stagger by 50 ms so the eye can follow which thing happened first.

The board background is `#0d0726` to `#06031a`. Amber is `var(--accent-amber)` with a faint glow via `text-shadow`. The font is JetBrains Mono — close enough to a real split-flap typeface without licensing a custom one.

</CSSection>

<CSSection number="04" eyebrow="DATA PIPELINE" title="A Go ingest" italic="and a wire of SSE." accent="purple" lead="Postgres for the player roster and cached stat history; Server-Sent Events for live deltas to all connected clients.">

The Go backend ingests data from a sports-stats provider every 30 seconds during active games. It diffs the latest snapshot against the previous one, computes which leaderboards moved, and broadcasts only the deltas — not the full board — to subscribed clients. The frontend animates only the cells that changed.

For non-game-time (off-hours, off-season), the board updates from cached data and only changes when daily aggregate stats roll over. There's a "season summary" mode for the off-season so the page isn't dead from November through February.

</CSSection>

<CSSection number="05" eyebrow="DECISIONS" title="Things I had to" italic="pick a side on." accent="purple">

<div class="decisions-grid">

<DecisionCard question="SSE or WebSockets?" pickedLabel="picked SSE" pickedAccent="purple">
SSE is one-direction (server → client), which is exactly what the leaderboard needs. No need for the bidirectional handshake of WebSockets, no proxies misinterpreting frames, easier to debug with `curl`. The latency is identical to WebSockets in practice for this kind of broadcast.
</DecisionCard>

<DecisionCard question="Free stats or paid?" pickedLabel="picked MLB Stats API" pickedAccent="purple">
The free MLB Stats API has the basics; FanGraphs has the rich metrics. v1 ships on the free API plus a small derived-stats layer (compute wRC+ from raw events). If traffic justifies it, swap in a paid Sportradar feed later.
</DecisionCard>

<DecisionCard question="Sound effects?" pickedLabel="picked off-by-default" pickedAccent="purple">
Real split-flap boards make a great mechanical clatter. A subtle 'click' on each flip is faithful to the source. Most people will mute it. Toggle in the UI, off by default, remembered in localStorage.
</DecisionCard>

<DecisionCard question="Mobile?" pickedLabel="picked stacked view" pickedAccent="purple">
The board is naturally horizontal — 6+ columns wide. On a 375 px viewport that's unreadable. Mobile gets a stacked view: each player is a card, columns become rows. The flap animation still works on the rows. The split-flap *feeling* is the goal, not the literal layout.
</DecisionCard>

</div>

</CSSection>

<CSSection number="06" eyebrow="OPEN QUESTIONS" title="What I'm still" italic="figuring out." accent="purple">

- **Historical mode.** "Show me the top 100 by wRC+ from 2019" is a great feature and a v2 trap — every season's data is a separate ingest job. Probably won't ship until someone asks for it three times.
- **Sharing.** A static-image generator so people can post screenshots of "my favorite SS leaderboard right now" — Open Graph image generation that captures the live state at request time.
- **Notifications.** When a chosen player's rank changes by more than N positions, ping me. Useful, but requires a user model. Defer.

</CSSection>

<AdjacentProjects current="diamond" />

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
