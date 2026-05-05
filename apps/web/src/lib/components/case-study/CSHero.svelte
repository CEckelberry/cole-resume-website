<!--
  CSHero — case-study hero header.
  Anatomy: eyebrow / title / serif italic deck / 4-column meta row.
-->
<script lang="ts">
  import type { Project } from '$lib/content/projects';
  import { ACCENT_COLOR, ACCENT_SOFT } from '$lib/content/projects';

  interface Props {
    project: Project;
    /** Eyebrow kind text, e.g. "A TOOL". */
    kind?: string;
    /** Deck (sub-headline). Falls back to project.description. */
    deck?: string;
    /** Eyebrow date string, e.g. "MAR 2026" or "CONCEPT". */
    date?: string;
    /** Optional role override (otherwise renders nothing for that meta cell). */
    role?: string;
    /** Optional shipped date label. */
    shipped?: string;
  }

  let { project, kind, deck, date = 'CONCEPT', role = 'solo', shipped = 'tbd' }: Props = $props();

  const statusLabel: Record<Project['status'], string> = {
    live: 'live',
    building: 'building',
    planned: 'concept'
  };
</script>

<header
  class="cs-hero"
  style:--cs-accent={ACCENT_COLOR[project.accent]}
  style:--cs-accent-soft={ACCENT_SOFT[project.accent]}
>
  <p class="eyebrow">
    project {project.number}
    <span class="sep" aria-hidden="true">·</span>
    {kind ?? project.kind}
    <span class="sep" aria-hidden="true">·</span>
    {date}
  </p>

  <h1 class="title">
    {project.title}
    {#if project.italicWord}
      <em>{project.italicWord}.</em>
    {/if}
  </h1>

  <p class="deck">{deck ?? project.description}</p>

  <dl class="meta-row">
    <div class="meta-cell">
      <dt>role</dt>
      <dd>{role}</dd>
    </div>
    <div class="meta-cell">
      <dt>stack</dt>
      <dd>{project.tags.slice(0, 3).join(', ')}</dd>
    </div>
    <div class="meta-cell">
      <dt>shipped</dt>
      <dd>{shipped}</dd>
    </div>
    <div class="meta-cell">
      <dt>status</dt>
      <dd>
        <span class="status-pill" data-status={project.status}>
          <span class="status-dot" aria-hidden="true"></span>
          {statusLabel[project.status]}
        </span>
      </dd>
    </div>
  </dl>
</header>

<style>
  .cs-hero {
    display: flex;
    flex-direction: column;
    gap: var(--space-4);
    padding-block: var(--space-7) var(--space-6);
  }

  .eyebrow {
    margin: 0;
    font-family: var(--font-mono);
    font-size: var(--type-micro);
    letter-spacing: 0.18em;
    text-transform: uppercase;
    color: var(--text-tertiary);
    display: inline-flex;
    align-items: center;
    gap: 8px;
    flex-wrap: wrap;
  }
  .eyebrow .sep {
    color: var(--text-muted);
  }

  .title {
    margin: 0;
    font-family: var(--font-sans);
    font-weight: 500;
    font-size: var(--type-h1);
    line-height: var(--type-h1-lh);
    letter-spacing: -0.02em;
    color: var(--text-primary);
    max-width: 880px;
  }
  .title em {
    font-family: var(--font-serif);
    font-style: italic;
    font-weight: 400;
    color: var(--cs-accent-soft);
    margin-left: 6px;
  }

  .deck {
    margin: 0;
    font-family: var(--font-serif);
    font-style: italic;
    font-size: var(--type-lead);
    line-height: var(--type-lead-lh);
    color: var(--text-secondary);
    max-width: 540px;
  }

  .meta-row {
    margin: var(--space-3) 0 0;
    padding: var(--space-4) 0;
    display: grid;
    grid-template-columns: repeat(2, minmax(0, 1fr));
    gap: var(--space-4);
    border-top: 0.5px solid var(--border-subtle);
    border-bottom: 0.5px solid var(--border-subtle);
  }
  @media (min-width: 720px) {
    .meta-row {
      grid-template-columns: repeat(4, minmax(0, 1fr));
    }
  }

  .meta-cell {
    display: flex;
    flex-direction: column;
    gap: 4px;
  }

  dt {
    font-family: var(--font-mono);
    font-size: var(--type-tiny);
    text-transform: uppercase;
    letter-spacing: 0.16em;
    color: var(--text-tertiary);
  }

  dd {
    margin: 0;
    font-family: var(--font-sans);
    font-size: var(--type-body-sm);
    color: var(--text-primary);
  }

  .status-pill {
    display: inline-flex;
    align-items: center;
    gap: 6px;
    padding: 3px 10px;
    border-radius: var(--radius-pill);
    border: 0.5px solid var(--border-default);
    background: var(--bg-surface-2);
    font-family: var(--font-mono);
    font-size: var(--type-tiny);
    letter-spacing: 0.06em;
    text-transform: lowercase;
  }
  .status-dot {
    width: 6px;
    height: 6px;
    border-radius: 50%;
    background: var(--text-muted);
  }
  .status-pill[data-status='live'] .status-dot {
    background: var(--accent-teal);
    animation: cs-pulse 2.4s var(--ease-in-out) infinite;
  }
  .status-pill[data-status='building'] .status-dot {
    background: var(--cs-accent);
  }
  @keyframes cs-pulse {
    0%,
    100% {
      box-shadow: 0 0 0 0 color-mix(in oklab, var(--accent-teal), transparent 60%);
    }
    50% {
      box-shadow: 0 0 0 5px color-mix(in oklab, var(--accent-teal), transparent 100%);
    }
  }
  @media (prefers-reduced-motion: reduce) {
    .status-dot {
      animation: none;
    }
  }
</style>
