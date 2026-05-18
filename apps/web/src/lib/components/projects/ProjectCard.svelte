<!--
  ProjectCard — generic glass tile used by the gallery. The preview region
  is a slot so each project supplies its own visual fragment.

  DESIGN.md spec:
    - ~330px tall, full-width within its grid column
    - Two parts: 160px preview region + body
    - Hover: translateY(-3px), border brightens, preview parallax shift
-->
<script lang="ts">
  import type { Snippet } from 'svelte';
  import type { Project } from '$lib/content/projects';
  import { ACCENT_COLOR, ACCENT_SOFT, ACCENT_EMPHASIS } from '$lib/content/projects';

  interface Props {
    project: Project;
    /** Visual fragment for the top 160px of the card. */
    preview: Snippet;
  }

  let { project, preview }: Props = $props();

  const statusLabel: Record<Project['status'], string> = {
    live: 'live',
    building: 'building',
    planned: 'planned'
  };
</script>

<a
  class="card"
  href={project.caseStudyHref}
  style:--card-accent={ACCENT_COLOR[project.accent]}
  style:--card-accent-soft={ACCENT_SOFT[project.accent]}
  style:--card-accent-emphasis={ACCENT_EMPHASIS[project.accent]}
  data-accent={project.accent}
  aria-label="case study: {project.title} {project.italicWord ?? ''} — {project.description}"
>
  <div class="preview">
    {@render preview()}
  </div>

  <div class="body">
    <p class="meta-top">
      <span class="number">{project.number}</span>
      <span class="kind">{project.kind}</span>
    </p>

    <h3 class="title">
      {project.title}
      {#if project.italicWord}
        <em>{project.italicWord}.</em>
      {/if}
    </h3>

    <p class="description">{project.description}</p>

    <ul class="tags">
      {#each project.tags as tag (tag)}
        <li>{tag}</li>
      {/each}
    </ul>

    <footer class="card-footer">
      <span class="status status-{project.status}">
        <span class="status-dot" aria-hidden="true"></span>
        {statusLabel[project.status]}
      </span>
      <span class="cta">
        case study <span aria-hidden="true">→</span>
      </span>
    </footer>
  </div>
</a>

<style>
  .card {
    display: flex;
    flex-direction: column;
    border: 0.5px solid var(--border-default);
    border-radius: var(--radius-lg);
    background: var(--bg-surface-2);
    backdrop-filter: blur(14px) saturate(140%);
    -webkit-backdrop-filter: blur(14px) saturate(140%);
    overflow: hidden;
    text-decoration: none;
    color: var(--text-primary);
    min-height: 380px;
    height: 100%;
    transition:
      transform var(--dur-default) var(--ease-out),
      border-color var(--dur-default) var(--ease-out),
      box-shadow var(--dur-default) var(--ease-out);
  }

  .card:hover {
    transform: translateY(-3px);
    border-color: var(--card-accent-soft);
    box-shadow: 0 14px 36px -18px color-mix(in oklab, var(--card-accent), transparent 60%);
  }

  .preview {
    height: 160px;
    flex-shrink: 0;
    overflow: hidden;
    position: relative;
    border-bottom: 0.5px solid var(--border-subtle);
    background: var(--bg-surface);
  }

  .body {
    padding: var(--space-5);
    display: flex;
    flex-direction: column;
    gap: var(--space-3);
    flex: 1;
  }

  .meta-top {
    display: flex;
    gap: 10px;
    align-items: center;
    font-family: var(--font-mono);
    font-size: var(--type-tiny);
    text-transform: uppercase;
    letter-spacing: 0.16em;
    color: var(--text-tertiary);
    margin: 0;
  }

  .number {
    color: var(--card-accent);
  }

  .title {
    margin: 0;
    font-family: var(--font-sans);
    font-weight: 500;
    font-size: var(--type-h3);
    line-height: var(--type-h3-lh);
    letter-spacing: -0.01em;
    color: var(--text-primary);
  }

  .title em {
    font-family: var(--font-serif);
    font-style: italic;
    font-weight: 400;
    color: var(--card-accent-emphasis);
    margin-left: 4px;
  }

  .description {
    margin: 0;
    font-size: var(--type-body-sm);
    line-height: 1.55;
    color: var(--text-secondary);
  }

  .tags {
    list-style: none;
    margin: 0;
    padding: 0;
    display: flex;
    flex-wrap: wrap;
    gap: 6px;
  }
  .tags li {
    font-family: var(--font-mono);
    font-size: var(--type-tiny);
    line-height: 1.2;
    padding: 4px 8px;
    border-radius: var(--radius-sm);
    background: var(--bg-surface-2);
    border: 0.5px solid var(--border-subtle);
    color: var(--text-tertiary);
    letter-spacing: 0.04em;
  }

  .card-footer {
    display: flex;
    align-items: center;
    justify-content: space-between;
    margin-top: auto;
    padding-top: var(--space-3);
    border-top: 0.5px dashed var(--border-subtle);
    font-family: var(--font-mono);
    font-size: var(--type-tiny);
    letter-spacing: 0.06em;
    text-transform: lowercase;
  }

  .status {
    display: inline-flex;
    align-items: center;
    gap: 6px;
    color: var(--text-tertiary);
  }
  .status-dot {
    width: 6px;
    height: 6px;
    border-radius: 50%;
    background: var(--text-muted);
  }
  .status-live .status-dot {
    background: var(--accent-teal);
    animation: pulse 2s var(--ease-in-out) infinite;
  }
  .status-building .status-dot {
    background: var(--card-accent);
  }
  @keyframes pulse {
    0%,
    100% {
      box-shadow: 0 0 0 0 color-mix(in oklab, var(--accent-teal), transparent 60%);
    }
    50% {
      box-shadow: 0 0 0 5px color-mix(in oklab, var(--accent-teal), transparent 100%);
    }
  }

  .cta {
    color: var(--text-secondary);
    transition: color var(--dur-fast) var(--ease-out);
  }
  .card:hover .cta {
    color: var(--card-accent-soft);
  }

  @media (prefers-reduced-motion: reduce) {
    .card,
    .cta,
    .status-dot {
      transition: none;
      animation: none;
    }
    .card:hover {
      transform: none;
    }
  }
</style>
