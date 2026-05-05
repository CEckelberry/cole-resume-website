<!--
  Footer — minimal placeholder for Task 1.1. Full four-column implementation
  lands in Task 1.12.
-->
<script lang="ts">
  import { BUILD_SHA, BUILD_TIME } from '$lib/utils/buildInfo';
</script>

<footer class="footer">
  <p class="line">
    <span>©2026 cole eckelberry</span>
    <span aria-hidden="true">·</span>
    <span class="dim">made in Bay Area</span>
    <span aria-hidden="true">·</span>
    <span class="dim">build</span>
    <code>{BUILD_SHA}</code>
    <time datetime={BUILD_TIME} class="dim">{BUILD_TIME.slice(0, 10)}</time>
  </p>
</footer>

<style>
  .footer {
    border-top: 0.5px solid var(--border-subtle);
    padding: var(--space-7) var(--space-6) var(--space-6);
    margin-top: var(--space-8);
  }

  .line {
    display: flex;
    flex-wrap: wrap;
    gap: 8px;
    align-items: baseline;
    font-family: var(--font-mono);
    font-size: var(--type-tiny);
    color: var(--text-secondary);
    letter-spacing: 0.04em;
    margin: 0;
  }

  .dim {
    color: var(--text-tertiary);
  }

  code {
    font-family: var(--font-mono);
    color: var(--accent-teal);
  }
</style>
