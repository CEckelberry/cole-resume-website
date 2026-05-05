<!--
  Backend Bake-off — preview region.
  Six rows of P95 latency bars. Static at v1; case study has the live demo.
-->
<script lang="ts">
  type Row = { runtime: string; p95: number; color: string };

  // Plausible synthetic numbers; ordered fastest → slowest within "fast group"
  // and "slow group" so Rust shortest, PHP longest as required by spec.
  const ROWS: Row[] = [
    { runtime: 'rust', p95: 9, color: 'var(--accent-amber)' },
    { runtime: 'go', p95: 12, color: 'var(--accent-teal)' },
    { runtime: 'bun', p95: 18, color: 'var(--accent-pink-soft)' },
    { runtime: 'node', p95: 24, color: 'var(--accent-purple-soft)' },
    { runtime: 'python', p95: 41, color: 'var(--accent-pink)' },
    {
      runtime: 'php',
      p95: 58,
      color: 'color-mix(in oklab, var(--accent-pink), var(--accent-coral))'
    }
  ];

  const MAX = Math.max(...ROWS.map((r) => r.p95));
</script>

<div class="preview">
  <p class="header">
    <span class="label">p95 latency</span>
    <span class="unit">ms</span>
  </p>
  <ul class="rows">
    {#each ROWS as row (row.runtime)}
      <li class="row">
        <span class="name">{row.runtime}</span>
        <span class="bar-wrap">
          <span class="bar" style:width="{(row.p95 / MAX) * 100}%" style:background={row.color}
          ></span>
        </span>
        <span class="value">{row.p95}</span>
      </li>
    {/each}
  </ul>
</div>

<style>
  .preview {
    height: 100%;
    width: 100%;
    padding: 12px 14px 10px;
    background: linear-gradient(180deg, #0a1814 0%, #060f0c 100%);
    color: rgba(255, 255, 255, 0.85);
    display: flex;
    flex-direction: column;
    gap: 6px;
  }

  .header {
    display: flex;
    justify-content: space-between;
    margin: 0 0 2px;
    font-family: var(--font-mono);
    font-size: 9px;
    text-transform: uppercase;
    letter-spacing: 0.16em;
    color: rgba(255, 255, 255, 0.45);
  }

  .label {
    color: var(--accent-teal);
  }

  .rows {
    margin: 0;
    padding: 0;
    list-style: none;
    display: flex;
    flex-direction: column;
    gap: 4px;
    flex: 1;
  }

  .row {
    display: grid;
    grid-template-columns: 44px 1fr 28px;
    align-items: center;
    gap: 8px;
    font-family: var(--font-mono);
    font-size: 10px;
    line-height: 1.2;
  }

  .name {
    color: rgba(255, 255, 255, 0.78);
    text-transform: lowercase;
  }

  .bar-wrap {
    height: 8px;
    background: rgba(255, 255, 255, 0.05);
    border-radius: 1px;
    overflow: hidden;
  }
  .bar {
    display: block;
    height: 100%;
    border-radius: 1px;
  }

  .value {
    text-align: right;
    color: rgba(255, 255, 255, 0.62);
    font-variant-numeric: tabular-nums;
  }
</style>
