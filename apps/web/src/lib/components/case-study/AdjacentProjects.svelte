<!--
  AdjacentProjects — three-card strip at the foot of a case study showing the
  other three projects. Includes shared-stack microcopy ("WASM, browser-side"
  etc.) per Task 2.6 spec.
-->
<script lang="ts">
  import { PROJECTS, type ProjectSlug } from '$lib/content/projects';

  interface Props {
    current: ProjectSlug;
  }

  let { current }: Props = $props();

  // Pair-wise shared-stack labels, per Task 2.6.
  const SHARED: Record<string, string> = {
    'linux-bakeoff': 'WASM, browser-side',
    'linux-diamond': 'SvelteKit',
    'linux-terraplane': 'Go',
    'bakeoff-diamond': 'Go SSE, GKE',
    'bakeoff-terraplane': 'Go, GCP infra',
    'diamond-terraplane': 'Go, Postgres'
  };

  function sharedLabel(a: ProjectSlug, b: ProjectSlug): string {
    const key = [a, b].sort().join('-');
    return SHARED[key] ?? '';
  }

  const others = $derived(PROJECTS.filter((p) => p.slug !== current));
</script>

<section class="adjacent">
  <p class="head">other projects</p>
  <ul class="row">
    {#each others as p (p.slug)}
      <li>
        <a href={p.caseStudyHref} data-accent={p.accent}>
          <span class="num">{p.number}</span>
          <span class="title">
            {p.title}{#if p.italicWord}&nbsp;<em>{p.italicWord}</em>{/if}
          </span>
          <span class="shared">shares {sharedLabel(current, p.slug)}</span>
          <span class="cta" aria-hidden="true">→</span>
        </a>
      </li>
    {/each}
  </ul>
</section>

<style>
  .adjacent {
    display: flex;
    flex-direction: column;
    gap: var(--space-4);
    padding-top: var(--space-7);
    border-top: 0.5px solid var(--border-subtle);
  }

  .head {
    font-family: var(--font-mono);
    font-size: var(--type-micro);
    text-transform: uppercase;
    letter-spacing: 0.18em;
    color: var(--text-tertiary);
    margin: 0;
  }

  .row {
    list-style: none;
    margin: 0;
    padding: 0;
    display: grid;
    grid-template-columns: 1fr;
    gap: var(--space-3);
  }
  @media (min-width: 720px) {
    .row {
      grid-template-columns: repeat(3, minmax(0, 1fr));
    }
  }

  a {
    display: grid;
    grid-template-columns: auto 1fr auto;
    grid-template-rows: auto auto;
    grid-template-areas:
      'num title cta'
      'num shared cta';
    align-items: center;
    column-gap: 12px;
    row-gap: 2px;
    padding: var(--space-4);
    text-decoration: none;
    background: var(--bg-surface-2);
    border: 0.5px solid var(--border-default);
    border-radius: var(--radius-md);
    transition:
      border-color var(--dur-default) var(--ease-out),
      transform var(--dur-default) var(--ease-out);
  }

  a:hover {
    transform: translateY(-2px);
  }
  a[data-accent='pink']:hover {
    border-color: var(--accent-pink-soft);
  }
  a[data-accent='teal']:hover {
    border-color: var(--accent-teal-soft);
  }
  a[data-accent='purple']:hover {
    border-color: var(--accent-purple-soft);
  }
  a[data-accent='amber']:hover {
    border-color: var(--accent-amber-soft);
  }

  .num {
    grid-area: num;
    font-family: var(--font-mono);
    font-size: var(--type-tiny);
    letter-spacing: 0.16em;
    color: var(--text-tertiary);
  }
  a[data-accent='pink'] .num {
    color: var(--accent-pink);
  }
  a[data-accent='teal'] .num {
    color: var(--accent-teal);
  }
  a[data-accent='purple'] .num {
    color: var(--accent-purple);
  }
  a[data-accent='amber'] .num {
    color: var(--accent-amber);
  }

  .title {
    grid-area: title;
    font-family: var(--font-sans);
    font-size: var(--type-body-sm);
    font-weight: 500;
    color: var(--text-primary);
  }
  .title em {
    font-family: var(--font-serif);
    font-style: italic;
    font-weight: 400;
  }

  .shared {
    grid-area: shared;
    font-family: var(--font-mono);
    font-size: var(--type-tiny);
    color: var(--text-muted);
    letter-spacing: 0.04em;
  }

  .cta {
    grid-area: cta;
    color: var(--text-tertiary);
    font-size: 16px;
    transition:
      transform var(--dur-fast) var(--ease-out),
      color var(--dur-fast) var(--ease-out);
  }
  a:hover .cta {
    transform: translateX(3px);
    color: var(--text-primary);
  }

  @media (prefers-reduced-motion: reduce) {
    a,
    a:hover,
    .cta {
      transition: none;
      transform: none;
    }
  }
</style>
