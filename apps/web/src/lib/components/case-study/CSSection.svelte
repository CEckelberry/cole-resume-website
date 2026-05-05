<!--
  CSSection — wraps a numbered case-study section.
  Anatomy: eyebrow (number + label) + title (with optional italic) + serif
  italic lead + body slot in CSProse.
-->
<script lang="ts">
  import type { Snippet } from 'svelte';
  import CSProse from './CSProse.svelte';

  interface Props {
    /** Section index, "01"–"NN". */
    number: string;
    /** Eyebrow label, e.g. "WHY", "ARCHITECTURE", "DECISIONS". */
    eyebrow: string;
    /** Section title, sentence case. */
    title: string;
    /** Italicized phrase appended after the title in the section accent. */
    italic?: string;
    /** Optional serif-italic lead sentence (max ~560 px). */
    lead?: string;
    /** Accent for italic phrase: defaults to inheriting --cs-accent-soft. */
    accent?: 'pink' | 'teal' | 'purple' | 'amber';
    /** Section body content. */
    children: Snippet;
  }

  let { number, eyebrow, title, italic, lead, accent, children }: Props = $props();

  const accentVar = $derived(
    accent ? `var(--accent-${accent}-soft)` : 'var(--cs-accent-soft, var(--accent-pink-soft))'
  );
</script>

<section class="cs-section" style:--section-accent={accentVar}>
  <header class="head">
    <p class="eyebrow">
      <span class="hairline" aria-hidden="true"></span>
      <span class="num">{number}</span>
      <span class="sep" aria-hidden="true">·</span>
      <span class="label">{eyebrow}</span>
    </p>
    <h2>
      {title}
      {#if italic}
        <em>{italic}</em>
      {/if}
    </h2>
    {#if lead}
      <p class="lead">{lead}</p>
    {/if}
  </header>

  <CSProse>
    {@render children()}
  </CSProse>
</section>

<style>
  .cs-section {
    display: flex;
    flex-direction: column;
    gap: var(--space-5);
    padding-block: var(--space-7) 0;
  }

  .head {
    display: flex;
    flex-direction: column;
    gap: var(--space-3);
    max-width: var(--container-prose);
  }

  .eyebrow {
    margin: 0;
    display: inline-flex;
    align-items: center;
    gap: 10px;
    font-family: var(--font-mono);
    font-size: var(--type-micro);
    text-transform: uppercase;
    letter-spacing: 0.18em;
    color: var(--text-tertiary);
  }
  .hairline {
    width: 36px;
    height: 0.5px;
    background: var(--border-default);
  }
  .num {
    color: var(--text-primary);
  }
  .sep {
    color: var(--text-muted);
  }

  h2 {
    margin: 0;
    font-family: var(--font-sans);
    font-weight: 500;
    font-size: var(--type-h2);
    line-height: var(--type-h2-lh);
    letter-spacing: -0.01em;
    color: var(--text-primary);
  }
  h2 em {
    font-family: var(--font-serif);
    font-style: italic;
    font-weight: 400;
    color: var(--section-accent);
    margin-left: 4px;
  }

  .lead {
    margin: 0;
    font-family: var(--font-serif);
    font-style: italic;
    font-size: var(--type-lead);
    line-height: var(--type-lead-lh);
    color: var(--text-secondary);
    max-width: 560px;
  }
</style>
