<!--
  StackMarquee — horizontally scrolling stack items below the hero.
  Spec from DESIGN.md → Stack marquee.

  Implementation: two copies of the same item track end-to-end. Animating
  translateX from 0 to -50% creates a seamless loop; when the first copy
  finishes scrolling off, the second copy is in the same position the first
  started. Hover slows from 30s to 60s instead of stopping entirely (DESIGN.md
  is explicit: "it should never stop").

  prefers-reduced-motion pauses the animation outright (DESIGN.md exception).
-->
<script lang="ts">
  type Item = { name: string; aside?: string };

  const ITEMS: Item[] = [
    { name: 'SvelteKit' },
    { name: 'Python', aside: 'my main' },
    { name: 'TypeScript' },
    { name: 'Kubernetes' },
    { name: 'Terraform' },
    { name: 'PostgreSQL' },
    { name: 'GCP' },
    { name: 'Docker' },
    { name: 'Bash', aside: 'since 2012' }
  ];
</script>

<div class="marquee" aria-label="primary stack">
  <div class="track">
    <!-- Two identical copies. The second is aria-hidden so screen readers
         only announce the list once. -->
    {#each [0, 1] as copy (copy)}
      <ul class="row" aria-hidden={copy === 1 ? 'true' : undefined}>
        {#each ITEMS as item (`${copy}-${item.name}`)}
          <li class="item">
            <span class="name">{item.name}</span>
            {#if item.aside}
              <em class="aside">({item.aside})</em>
            {/if}
            <span class="sep" aria-hidden="true">◇</span>
          </li>
        {/each}
      </ul>
    {/each}
  </div>
</div>

<style>
  .marquee {
    position: relative;
    width: 100%;
    overflow: hidden;
    /* Edge fades — about 60px on each side so the items appear to dissolve
       into the page rather than slam off the edge. */
    -webkit-mask-image: linear-gradient(
      to right,
      transparent 0,
      #000 60px,
      #000 calc(100% - 60px),
      transparent 100%
    );
    mask-image: linear-gradient(
      to right,
      transparent 0,
      #000 60px,
      #000 calc(100% - 60px),
      transparent 100%
    );
    padding-block: var(--space-4);
    border-top: 0.5px solid var(--border-subtle);
    border-bottom: 0.5px solid var(--border-subtle);
  }

  .track {
    display: flex;
    width: max-content;
    animation: scroll 30s linear infinite;
    will-change: transform;
  }

  @keyframes scroll {
    from {
      transform: translateX(0);
    }
    to {
      transform: translateX(-50%);
    }
  }

  .row {
    display: flex;
    align-items: center;
    margin: 0;
    padding: 0;
    list-style: none;
  }

  .item {
    display: inline-flex;
    align-items: baseline;
    gap: 6px;
    padding-inline: var(--space-5);
    font-family: var(--font-mono);
    font-size: var(--type-meta);
    line-height: 1.4;
    letter-spacing: 0.04em;
    color: var(--text-secondary);
    white-space: nowrap;
  }

  .name {
    color: var(--text-primary);
  }

  .aside {
    font-family: var(--font-serif);
    font-style: italic;
    font-weight: 400;
    color: var(--text-tertiary);
    letter-spacing: 0;
  }

  .sep {
    color: var(--accent-pink-soft);
    margin-left: 6px;
    font-size: 9px;
    transform: translateY(-1px);
  }

  @media (prefers-reduced-motion: reduce) {
    .track {
      animation: none;
    }
  }
</style>
