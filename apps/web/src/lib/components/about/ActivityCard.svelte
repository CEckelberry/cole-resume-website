<!--
  ActivityCard — recent activity card. Placeholder version with static rows
  for v1; structured so the rows can be swapped for live `/api/activity`
  output in Phase 4.
-->
<script lang="ts">
  type ActivityKind = 'deploy' | 'commit' | 'release' | 'reading' | 'outdoors';

  type Item = {
    kind: ActivityKind;
    title: string;
    detail?: string;
    when: string;
  };

  // Placeholder rows. The real data shape mirrors `Activity` in
  // apps/api/internal/http/handlers/activity.go (Phase 4).
  const ITEMS: Item[] = [
    {
      kind: 'deploy',
      title: 'cole-eckelberry.com',
      detail: 'main@build · 47 ms p95',
      when: '2h ago'
    },
    {
      kind: 'commit',
      title: 'fix(api): clamp activity cache TTL when github rate-limited',
      detail: 'cole-eckelberry/portfolio',
      when: '6h ago'
    },
    {
      kind: 'commit',
      title: 'feat(web): split-flap board now reads SSE deltas',
      detail: 'cole-eckelberry/diamond-departures',
      when: '1d ago'
    },
    {
      kind: 'reading',
      title: 'Designing Data-Intensive Applications',
      detail: 'finished · ch. 7 transactions',
      when: '3d ago'
    },
    {
      kind: 'outdoors',
      title: 'morning run · Los Gatos creek',
      detail: '6.2 mi · 9:18/mi',
      when: '4d ago'
    }
  ];

  const KIND_LABEL: Record<ActivityKind, string> = {
    deploy: 'deploy',
    commit: 'commit',
    release: 'release',
    reading: 'reading',
    outdoors: 'outdoors'
  };

  const KIND_COLOR: Record<ActivityKind, string> = {
    deploy: 'var(--accent-teal)',
    commit: 'var(--accent-pink-soft)',
    release: 'var(--accent-amber)',
    reading: 'var(--accent-purple-soft)',
    outdoors: 'var(--accent-amber-soft)'
  };
</script>

<aside class="activity" aria-label="recent activity">
  <header>
    <p class="eyebrow">recent activity</p>
    <p class="dim">last 7 days</p>
  </header>

  <ul class="rows">
    {#each ITEMS as item, i (i)}
      <li class="row">
        <span class="kind" style:--kind-color={KIND_COLOR[item.kind]}>
          <span class="kind-dot" aria-hidden="true"></span>
          {KIND_LABEL[item.kind]}
        </span>
        <span class="title">{item.title}</span>
        {#if item.detail}
          <span class="detail">{item.detail}</span>
        {/if}
        <time class="when">{item.when}</time>
      </li>
    {/each}
  </ul>

  <p class="footnote">
    placeholder data — wired to <code>/api/activity</code> in phase 4.
  </p>
</aside>

<style>
  .activity {
    border: 0.5px solid var(--border-default);
    border-radius: var(--radius-lg);
    background: var(--bg-surface-2);
    backdrop-filter: blur(14px) saturate(140%);
    -webkit-backdrop-filter: blur(14px) saturate(140%);
    padding: var(--space-5);
    display: flex;
    flex-direction: column;
    gap: var(--space-3);
  }

  header {
    display: flex;
    align-items: baseline;
    justify-content: space-between;
    gap: 12px;
  }

  .eyebrow {
    font-family: var(--font-mono);
    font-size: var(--type-micro);
    text-transform: uppercase;
    letter-spacing: 0.18em;
    color: var(--text-tertiary);
    margin: 0;
  }

  .dim {
    font-family: var(--font-mono);
    font-size: var(--type-tiny);
    color: var(--text-muted);
    margin: 0;
    letter-spacing: 0.04em;
  }

  .rows {
    list-style: none;
    margin: 0;
    padding: 0;
    display: flex;
    flex-direction: column;
  }

  .row {
    display: grid;
    grid-template-columns: 88px 1fr auto;
    grid-template-rows: auto auto;
    grid-template-areas:
      'kind title when'
      'kind detail when';
    column-gap: 12px;
    row-gap: 1px;
    padding: 10px 0;
    border-bottom: 0.5px dashed var(--border-subtle);
    align-items: baseline;
  }
  .row:last-child {
    border-bottom: none;
  }

  .kind {
    grid-area: kind;
    display: inline-flex;
    align-items: center;
    gap: 6px;
    font-family: var(--font-mono);
    font-size: var(--type-tiny);
    color: var(--text-tertiary);
    letter-spacing: 0.1em;
    text-transform: uppercase;
  }
  .kind-dot {
    width: 6px;
    height: 6px;
    border-radius: 50%;
    background: var(--kind-color);
  }

  .title {
    grid-area: title;
    font-size: var(--type-body-sm);
    line-height: 1.4;
    color: var(--text-primary);
  }

  .detail {
    grid-area: detail;
    font-family: var(--font-mono);
    font-size: var(--type-tiny);
    color: var(--text-tertiary);
    letter-spacing: 0.04em;
  }

  .when {
    grid-area: when;
    font-family: var(--font-mono);
    font-size: var(--type-tiny);
    color: var(--text-muted);
    letter-spacing: 0.06em;
    align-self: start;
  }

  .footnote {
    margin: 0;
    font-family: var(--font-mono);
    font-size: var(--type-tiny);
    color: var(--text-muted);
    letter-spacing: 0.02em;
  }
  .footnote code {
    color: var(--accent-teal);
    background: transparent;
    padding: 0;
  }
</style>
