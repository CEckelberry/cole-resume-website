<!--
  ArchDiagram — wraps a hand-drawn inline SVG architecture diagram with an
  optional caption. Each case study supplies its own SVG via the `svg` snippet
  (kept as a snippet rather than `src` so token colors resolve from the
  parent's accent context).
-->
<script lang="ts">
  import type { Snippet } from 'svelte';

  interface Props {
    /** SVG markup as a snippet so styles cascade from the parent. */
    children: Snippet;
    caption?: string;
  }

  let { children, caption }: Props = $props();
</script>

<figure class="arch">
  <div class="canvas">
    {@render children()}
  </div>
  {#if caption}
    <figcaption>{caption}</figcaption>
  {/if}
</figure>

<style>
  .arch {
    margin: 0;
    display: flex;
    flex-direction: column;
    gap: var(--space-3);
  }

  .canvas {
    border: 0.5px solid var(--border-default);
    border-radius: var(--radius-lg);
    background: var(--bg-surface-2);
    backdrop-filter: blur(12px);
    -webkit-backdrop-filter: blur(12px);
    padding: var(--space-5);
    overflow: hidden;
  }

  /* Stretch the slotted SVG to fill its container responsively. */
  .canvas :global(> svg) {
    width: 100%;
    height: auto;
    display: block;
  }

  figcaption {
    font-family: var(--font-mono);
    font-size: var(--type-tiny);
    color: var(--text-tertiary);
    letter-spacing: 0.04em;
  }
</style>
