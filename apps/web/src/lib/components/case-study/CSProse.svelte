<!--
  CSProse — typography wrapper for case-study body text.
  Provides defaults for p, ul, ol, h3, strong, em, code, blockquote.
-->
<script lang="ts">
  import type { Snippet } from 'svelte';

  let { children }: { children: Snippet } = $props();
</script>

<div class="prose">
  {@render children()}
</div>

<style>
  .prose {
    max-width: var(--container-prose);
    color: var(--text-secondary);
    font-size: var(--type-body);
    line-height: 1.75;
  }

  .prose :global(> * + *) {
    margin-top: var(--space-4);
  }

  .prose :global(p) {
    margin: 0;
  }

  .prose :global(strong) {
    font-weight: 500;
    color: var(--text-primary);
  }

  .prose :global(em) {
    font-family: var(--font-serif);
    font-style: italic;
    color: var(--text-primary);
  }

  .prose :global(a) {
    color: var(--text-primary);
    text-decoration: underline;
    text-decoration-color: var(--cs-accent-soft, var(--accent-pink-soft));
    text-decoration-thickness: 1px;
    text-underline-offset: 3px;
    transition: text-decoration-color var(--dur-fast) var(--ease-out);
  }
  .prose :global(a:hover) {
    text-decoration-color: var(--cs-accent, var(--accent-pink));
  }

  .prose :global(ul),
  .prose :global(ol) {
    margin: 0;
    padding-left: 1.4em;
    display: flex;
    flex-direction: column;
    gap: 6px;
  }

  .prose :global(li) {
    color: var(--text-secondary);
  }

  .prose :global(li::marker) {
    color: var(--cs-accent, var(--accent-pink-soft));
  }

  .prose :global(h3) {
    font-family: var(--font-sans);
    font-weight: 500;
    font-size: var(--type-h3);
    line-height: var(--type-h3-lh);
    letter-spacing: -0.005em;
    color: var(--text-primary);
    margin-top: var(--space-5);
  }

  .prose :global(code) {
    font-family: var(--font-mono);
    font-size: 0.92em;
    color: var(--cs-accent-soft, var(--accent-pink-soft));
    background: var(--bg-surface-2);
    padding: 1px 5px;
    border-radius: var(--radius-sm);
    border: 0.5px solid var(--border-subtle);
  }

  .prose :global(blockquote) {
    margin: 0;
    padding: var(--space-4) var(--space-5);
    border-left: 2px solid var(--cs-accent-soft, var(--accent-pink-soft));
    background: var(--bg-surface-2);
    border-radius: 0 var(--radius-md) var(--radius-md) 0;
    color: var(--text-primary);
    font-family: var(--font-serif);
    font-style: italic;
    font-size: var(--type-lead);
    line-height: var(--type-lead-lh);
  }

  .prose :global(table) {
    width: 100%;
    border-collapse: collapse;
    font-size: var(--type-body-sm);
  }
  .prose :global(th),
  .prose :global(td) {
    text-align: left;
    padding: 8px 12px 8px 0;
    border-bottom: 0.5px solid var(--border-subtle);
  }
  .prose :global(th) {
    font-family: var(--font-mono);
    font-weight: 500;
    font-size: var(--type-tiny);
    text-transform: uppercase;
    letter-spacing: 0.12em;
    color: var(--text-tertiary);
  }
</style>
