<!--
  Hero — top-of-page composition.

  Anatomy:
    eyebrow            PORTFOLIO · 2026
    name               Cole Eckelberry        (gradient last name)
    tagline            A platform engineer who keeps *making weird things* after 5pm.
    bio                Ten years of building cloud platforms by day. Four side
                       projects below — all open source, all live.
    ctas               [see the work →]   [cv.pdf]

  The terminal slot reserved on the right is filled by Task 1.3.
-->
<script lang="ts">
  import HeroName from './HeroName.svelte';
  import Button from '$lib/components/ui/Button.svelte';
  import type { Snippet } from 'svelte';

  interface Props {
    /** Optional terminal/visual slotted to the right of the hero text on desktop. */
    terminal?: Snippet;
  }

  let { terminal }: Props = $props();
</script>

<section class="hero">
  <div class="text">
    <p class="eyebrow">portfolio · 2026</p>

    <HeroName />

    <p class="tagline">
      A platform engineer who keeps <em class="accent">making weird things</em> after 5pm.
    </p>

    <p class="bio">
      Ten years building cloud platforms — Python, JavaScript, TypeScript, plus the operations
      stack. Four side projects below, all open source, all live.
    </p>

    <div class="ctas">
      <Button href="/#work" variant="primary" size="lg">
        see the work
        <span aria-hidden="true">→</span>
      </Button>
      <Button href="/cv" variant="secondary" size="lg">cv.pdf</Button>
    </div>
  </div>

  {#if terminal}
    <aside class="terminal-slot" aria-label="terminal demo">
      {@render terminal()}
    </aside>
  {/if}
</section>

<style>
  .hero {
    display: grid;
    grid-template-columns: 1fr;
    gap: var(--space-7);
    padding: var(--space-7) 0 var(--space-8);
    align-items: start;
  }

  @media (min-width: 960px) {
    .hero {
      /* Text takes ~56% so the name has room to breathe; terminal sits compactly
         on the right at the design's ~280px target. */
      grid-template-columns: minmax(0, 1.4fr) minmax(280px, 0.8fr);
      gap: var(--space-8);
    }
  }

  .text {
    display: flex;
    flex-direction: column;
    gap: var(--space-5);
    max-width: var(--container-prose);
  }

  .eyebrow {
    font-family: var(--font-mono);
    font-size: var(--type-micro);
    line-height: var(--type-micro-lh);
    text-transform: uppercase;
    letter-spacing: 0.18em;
    color: var(--text-tertiary);
    margin: 0;
  }

  .tagline {
    font-size: var(--type-lead);
    line-height: var(--type-lead-lh);
    color: var(--text-primary);
    margin: 0;
    max-width: var(--container-deck);
  }

  .tagline .accent {
    font-family: var(--font-serif);
    font-style: italic;
    font-weight: 400;
    color: var(--accent-pink-soft);
    /* Soft underline highlight: sits below the baseline, behind the glyphs.
       linear-gradient lets it resize cleanly with text wrapping. */
    background-image: linear-gradient(
      transparent calc(100% - 6px),
      color-mix(in oklab, var(--accent-pink-soft), transparent 65%) calc(100% - 6px),
      color-mix(in oklab, var(--accent-pink-soft), transparent 65%) 96%,
      transparent 96%
    );
    background-repeat: no-repeat;
    padding-inline: 2px;
  }

  .bio {
    font-size: var(--type-body);
    line-height: 1.7;
    color: var(--text-secondary);
    margin: 0;
    max-width: var(--container-deck);
  }

  .ctas {
    display: flex;
    flex-wrap: wrap;
    gap: var(--space-3);
    margin-top: var(--space-2);
  }

  .terminal-slot {
    /* On wide layouts the terminal sits to the right of the text. On narrow
       it falls below — stacked by grid-template-columns: 1fr default. */
    align-self: start;
  }
</style>
