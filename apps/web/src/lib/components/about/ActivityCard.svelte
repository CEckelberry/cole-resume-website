<!--
  ActivityCard — recent GitHub activity. Calls /api/activity on mount,
  falls back to a sensible empty state if the API is unreachable or
  returns degraded.
-->
<script lang="ts">
  import { onMount } from 'svelte';

  type ActivityKind = 'commit' | 'release' | 'create' | 'pr';

  interface ActivityItem {
    kind: ActivityKind;
    repo: string;
    title: string;
    url: string;
    age_text: string;
    at: string;
  }

  type Status = 'loading' | 'ready' | 'degraded' | 'error';

  let status = $state<Status>('loading');
  let items = $state<ActivityItem[]>([]);

  const REFRESH_MS = 5 * 60 * 1000;

  async function fetchActivity() {
    try {
      const res = await fetch('/api/activity', { headers: { accept: 'application/json' } });
      if (!res.ok) {
        status = 'error';
        return;
      }
      const body = (await res.json()) as { items: ActivityItem[]; degraded: boolean };
      items = body.items;
      status = body.degraded || items.length === 0 ? 'degraded' : 'ready';
    } catch {
      status = 'error';
    }
  }

  onMount(() => {
    void fetchActivity();
    const id = setInterval(() => void fetchActivity(), REFRESH_MS);
    return () => clearInterval(id);
  });

  const KIND_LABEL: Record<ActivityKind, string> = {
    commit: 'commit',
    release: 'release',
    create: 'new repo',
    pr: 'pr'
  };

  const KIND_COLOR: Record<ActivityKind, string> = {
    commit: 'var(--accent-pink-soft)',
    release: 'var(--accent-amber)',
    create: 'var(--accent-teal)',
    pr: 'var(--accent-purple-soft)'
  };
</script>

<aside class="activity" aria-label="recent github activity">
  <header>
    <p class="eyebrow">recent activity</p>
    <p class="dim">github.com/CEckelberry</p>
  </header>

  {#if status === 'loading'}
    <div class="placeholder">
      <span class="dot" aria-hidden="true"></span>
      <span>loading…</span>
    </div>
  {:else if status === 'ready'}
    <ul class="rows">
      {#each items as item, i (i)}
        <li class="row">
          <span class="kind" style:--kind-color={KIND_COLOR[item.kind]}>
            <span class="kind-dot" aria-hidden="true"></span>
            {KIND_LABEL[item.kind]}
          </span>
          <a class="title" href={item.url} rel="noopener" target="_blank">{item.title}</a>
          <span class="detail">{item.repo}</span>
          <time class="when" datetime={item.at}>{item.age_text}</time>
        </li>
      {/each}
    </ul>
  {:else}
    <p class="empty">
      {#if status === 'degraded'}
        no recent public events. a fresh repo or a quiet week.
      {:else}
        couldn't reach github. try refreshing in a minute.
      {/if}
    </p>
  {/if}

  <p class="footnote">
    pulled from the <a href="https://docs.github.com/rest/activity/events" rel="noopener"
      >events api</a
    >, cached 5 min upstream.
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

  /* Loading state */
  .placeholder {
    display: flex;
    align-items: center;
    gap: 8px;
    padding: var(--space-4) 0;
    font-family: var(--font-mono);
    font-size: var(--type-tiny);
    color: var(--text-tertiary);
    letter-spacing: 0.06em;
    text-transform: uppercase;
  }
  .placeholder .dot {
    width: 6px;
    height: 6px;
    border-radius: 50%;
    background: var(--text-muted);
    animation: blink 1.2s ease-in-out infinite;
  }
  @keyframes blink {
    0%,
    100% {
      opacity: 0.3;
    }
    50% {
      opacity: 1;
    }
  }
  @media (prefers-reduced-motion: reduce) {
    .placeholder .dot {
      animation: none;
    }
  }

  /* Empty state */
  .empty {
    margin: 0;
    padding: var(--space-3) 0;
    font-family: var(--font-serif);
    font-style: italic;
    color: var(--text-tertiary);
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
    grid-template-columns: 78px 1fr auto;
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
    text-decoration: none;
    overflow: hidden;
    text-overflow: ellipsis;
    white-space: nowrap;
    transition: color var(--dur-fast) var(--ease-out);
  }
  .title:hover {
    color: var(--accent-pink-soft);
  }

  .detail {
    grid-area: detail;
    font-family: var(--font-mono);
    font-size: var(--type-tiny);
    color: var(--text-tertiary);
    letter-spacing: 0.04em;
    overflow: hidden;
    text-overflow: ellipsis;
    white-space: nowrap;
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
  .footnote a {
    color: var(--text-tertiary);
    text-decoration: underline;
    text-decoration-color: var(--accent-pink-soft);
    text-underline-offset: 2px;
  }
  .footnote a:hover {
    color: var(--text-primary);
  }
</style>
