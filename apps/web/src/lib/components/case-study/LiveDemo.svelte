<!--
  LiveDemo — glass-card frame with a fake browser bar.
  The slot is the demo content. For projects that aren't built yet, slot
  the gallery card preview component scaled up.
-->
<script lang="ts">
  import type { Snippet } from 'svelte';

  interface Props {
    /** Public URL where the demo lives — shown in the address bar. */
    url: string;
    /** Title shown above the frame. */
    title?: string;
    /** Optional subtitle in the address bar tooltip-zone. */
    subtitle?: string;
    /** Demo content. */
    children: Snippet;
    /** Whether the demo is real (live) vs placeholder. */
    placeholder?: boolean;
  }

  let { url, title, subtitle, children, placeholder = false }: Props = $props();

  const host = $derived(url.replace(/^https?:\/\//, ''));
</script>

<figure class="live-demo">
  {#if title || subtitle}
    <figcaption class="caption">
      {#if title}<span class="title">{title}</span>{/if}
      {#if subtitle}<span class="subtitle">{subtitle}</span>{/if}
    </figcaption>
  {/if}

  <div class="window">
    <header class="bar">
      <span class="dots" aria-hidden="true">
        <span class="dot d-r"></span>
        <span class="dot d-y"></span>
        <span class="dot d-g"></span>
      </span>
      <div class="address" title={url}>
        <span class="lock" aria-hidden="true">⌁</span>
        <span class="host">{host}</span>
      </div>
      <span class="status-pill" data-placeholder={placeholder}>
        <span class="status-dot" aria-hidden="true"></span>
        {placeholder ? 'preview' : 'live'}
      </span>
    </header>

    <div class="canvas">
      {@render children()}
    </div>
  </div>
</figure>

<style>
  .live-demo {
    margin: 0;
    display: flex;
    flex-direction: column;
    gap: var(--space-3);
  }

  .caption {
    display: flex;
    align-items: baseline;
    gap: 12px;
    font-family: var(--font-mono);
    font-size: var(--type-tiny);
    text-transform: uppercase;
    letter-spacing: 0.16em;
    color: var(--text-tertiary);
  }
  .caption .title {
    color: var(--text-secondary);
  }
  .caption .subtitle {
    color: var(--text-muted);
  }

  .window {
    border: 0.5px solid var(--border-default);
    border-radius: var(--radius-xl);
    background: var(--bg-surface-2);
    backdrop-filter: blur(14px) saturate(140%);
    -webkit-backdrop-filter: blur(14px) saturate(140%);
    overflow: hidden;
    box-shadow: 0 20px 50px -22px color-mix(in oklab, var(--cs-accent, #000), transparent 65%);
  }

  .bar {
    display: grid;
    grid-template-columns: auto 1fr auto;
    align-items: center;
    gap: 12px;
    padding: 8px 12px;
    border-bottom: 0.5px solid var(--border-subtle);
    background: var(--bg-surface-3);
  }

  .dots {
    display: inline-flex;
    gap: 6px;
  }
  .dot {
    width: 9px;
    height: 9px;
    border-radius: 50%;
    box-shadow: 0 0 0 0.5px rgba(0, 0, 0, 0.25) inset;
  }
  .d-r {
    background: #ff5f57;
  }
  .d-y {
    background: #febc2e;
  }
  .d-g {
    background: #28c840;
  }

  .address {
    display: inline-flex;
    align-items: center;
    gap: 8px;
    padding: 4px 12px;
    border-radius: var(--radius-pill);
    background: var(--bg-surface);
    font-family: var(--font-mono);
    font-size: var(--type-tiny);
    color: var(--text-secondary);
    overflow: hidden;
    text-overflow: ellipsis;
    white-space: nowrap;
  }
  .lock {
    color: var(--accent-teal);
  }
  .host {
    overflow: hidden;
    text-overflow: ellipsis;
  }

  .status-pill {
    display: inline-flex;
    align-items: center;
    gap: 6px;
    padding: 3px 8px;
    border-radius: var(--radius-pill);
    border: 0.5px solid var(--border-subtle);
    font-family: var(--font-mono);
    font-size: var(--type-nano);
    text-transform: uppercase;
    letter-spacing: 0.14em;
    color: var(--text-tertiary);
  }
  .status-dot {
    width: 5px;
    height: 5px;
    border-radius: 50%;
    background: var(--accent-teal);
    animation: dpulse 2.4s var(--ease-in-out) infinite;
  }
  .status-pill[data-placeholder='true'] .status-dot {
    background: var(--text-muted);
    animation: none;
  }
  @keyframes dpulse {
    0%,
    100% {
      box-shadow: 0 0 0 0 color-mix(in oklab, var(--accent-teal), transparent 60%);
    }
    50% {
      box-shadow: 0 0 0 4px color-mix(in oklab, var(--accent-teal), transparent 100%);
    }
  }

  @media (prefers-reduced-motion: reduce) {
    .status-dot {
      animation: none;
    }
  }

  .canvas {
    /* The slotted content gets a fixed height for placeholder previews so
       the visual frame holds even when the project isn't built yet. Real
       embeds can override min-height as needed. */
    min-height: 320px;
    background: var(--bg-canvas);
    overflow: hidden;
  }
</style>
