<!--
  ProjectsSection — section header + filter chips bar + 2x2 gallery grid.
  Filter chip wiring lands in Task 1.10; for Task 1.5 the chips render but
  do nothing.
-->
<script lang="ts">
  import { fade } from 'svelte/transition';
  import { PROJECTS, FILTERS, type ProjectFilter } from '$lib/content/projects';
  import ProjectCard from './ProjectCard.svelte';
  import FilterChips from './FilterChips.svelte';
  import LinuxPreview from './previews/LinuxPreview.svelte';
  import BakeoffPreview from './previews/BakeoffPreview.svelte';
  import DiamondPreview from './previews/DiamondPreview.svelte';
  import TerraplanePreview from './previews/TerraplanePreview.svelte';

  type ActiveFilter = 'all' | ProjectFilter;
  let active = $state<ActiveFilter>('all');

  const filtered = $derived(
    active === 'all' ? PROJECTS : PROJECTS.filter((p) => p.filter === active)
  );

  const previewBySlug = {
    linux: LinuxPreview,
    bakeoff: BakeoffPreview,
    diamond: DiamondPreview,
    terraplane: TerraplanePreview
  };
</script>

<section id="work" class="section">
  <header class="section-head">
    <p class="eyebrow">section 02 · work</p>
    <h2>
      Things I built when I should've been <em>sleeping.</em>
    </h2>
    <p class="lead">
      Four side projects, all open source, all live. Each one is a place I went to learn something I
      couldn't learn on the job — Go, WASM, sabermetrics, terraform internals.
    </p>
  </header>

  <FilterChips
    filters={FILTERS}
    {active}
    ariaLabel="filter projects"
    onchange={(next) => (active = next)}
  />

  <div class="grid">
    {#each filtered as project (project.slug)}
      <div class="cell" in:fade={{ duration: 200 }} out:fade={{ duration: 120 }}>
        <ProjectCard {project}>
          {#snippet preview()}
            {@const PreviewComponent = previewBySlug[project.slug]}
            <PreviewComponent />
          {/snippet}
        </ProjectCard>
      </div>
    {/each}
  </div>
</section>

<style>
  .section {
    padding: var(--space-8) 0 var(--space-7);
    display: flex;
    flex-direction: column;
    gap: var(--space-6);
  }

  .section-head {
    display: flex;
    flex-direction: column;
    gap: var(--space-3);
    max-width: var(--container-prose);
  }

  .eyebrow {
    font-family: var(--font-mono);
    font-size: var(--type-micro);
    line-height: var(--type-micro-lh);
    text-transform: uppercase;
    letter-spacing: 0.18em;
    color: var(--text-tertiary);
    margin: 0;
  }

  h2 {
    font-family: var(--font-sans);
    font-weight: 500;
    font-size: var(--type-h2);
    line-height: var(--type-h2-lh);
    letter-spacing: -0.01em;
    color: var(--text-primary);
    margin: 0;
  }

  .lead {
    font-size: var(--type-body);
    line-height: 1.6;
    color: var(--text-secondary);
    margin: 0;
    max-width: var(--container-deck);
  }

  .grid {
    display: grid;
    grid-template-columns: 1fr;
    gap: var(--space-5);
  }

  @media (min-width: 768px) {
    .grid {
      grid-template-columns: 1fr 1fr;
      gap: var(--space-6);
    }
  }
</style>
