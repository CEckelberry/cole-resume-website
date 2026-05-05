<!--
  DecisionCard — surfaces a tradeoff inside a case study.
  Anatomy: serif-italic question + 2–3 sentence answer + picked/skipped pill.
-->
<script lang="ts">
  import type { Snippet } from 'svelte';

  type Accent = 'pink' | 'teal' | 'purple' | 'amber';

  interface Props {
    question: string;
    pickedLabel: string;
    /** Accent token for the pill — usually the project accent for "picked"
        and pink for "skipped". */
    pickedAccent: Accent;
    children: Snippet;
  }

  let { question, pickedLabel, pickedAccent, children }: Props = $props();
</script>

<article class="decision" data-accent={pickedAccent}>
  <p class="question">{question}</p>
  <div class="answer">
    {@render children()}
  </div>
  <span class="pill">{pickedLabel}</span>
</article>

<style>
  .decision {
    display: flex;
    flex-direction: column;
    gap: var(--space-3);
    padding: var(--space-5);
    border: 0.5px solid var(--border-default);
    border-radius: var(--radius-lg);
    background: var(--bg-surface-2);
    backdrop-filter: blur(12px);
    -webkit-backdrop-filter: blur(12px);
  }

  .question {
    margin: 0;
    font-family: var(--font-serif);
    font-style: italic;
    font-size: var(--type-lead);
    line-height: var(--type-lead-lh);
    color: var(--cs-accent-soft, var(--accent-pink-soft));
  }

  .answer {
    font-family: var(--font-sans);
    font-size: var(--type-body-sm);
    line-height: 1.65;
    color: var(--text-secondary);
  }

  .pill {
    align-self: flex-start;
    padding: 3px 10px;
    border-radius: var(--radius-pill);
    font-family: var(--font-mono);
    font-size: var(--type-tiny);
    text-transform: lowercase;
    letter-spacing: 0.06em;
    border: 0.5px solid;
    background: transparent;
    margin-top: 4px;
  }

  .decision[data-accent='pink'] .pill {
    color: var(--accent-pink);
    border-color: color-mix(in oklab, var(--accent-pink), transparent 50%);
    background: color-mix(in oklab, var(--accent-pink), transparent 88%);
  }
  .decision[data-accent='teal'] .pill {
    color: var(--accent-teal);
    border-color: color-mix(in oklab, var(--accent-teal), transparent 50%);
    background: color-mix(in oklab, var(--accent-teal), transparent 88%);
  }
  .decision[data-accent='purple'] .pill {
    color: var(--accent-purple);
    border-color: color-mix(in oklab, var(--accent-purple), transparent 50%);
    background: color-mix(in oklab, var(--accent-purple), transparent 88%);
  }
  .decision[data-accent='amber'] .pill {
    color: var(--accent-amber);
    border-color: color-mix(in oklab, var(--accent-amber), transparent 50%);
    background: color-mix(in oklab, var(--accent-amber), transparent 88%);
  }
</style>
