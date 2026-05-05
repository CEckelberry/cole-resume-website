<!--
  Diamond Departures — preview region.
  Penn Station split-flap aesthetic: amber text on deep purple, mono.
-->
<script lang="ts">
  type Row = {
    rank: string;
    name: string;
    pos: string;
    stat: string;
    arrow: 'up' | 'down' | 'flat';
  };

  const ROWS: Row[] = [
    { rank: '01', name: 'JUDGE', pos: 'RF NYY', stat: '+0.4', arrow: 'up' },
    { rank: '02', name: 'OHTANI', pos: 'DH LAD', stat: '+0.3', arrow: 'up' },
    { rank: '03', name: 'SOTO', pos: 'RF NYM', stat: '-0.1', arrow: 'down' },
    { rank: '04', name: 'WITT JR', pos: 'SS KC', stat: '+0.2', arrow: 'up' },
    { rank: '05', name: 'BETTS', pos: '2B LAD', stat: '+0.0', arrow: 'flat' },
    { rank: '06', name: 'DE LA CRUZ', pos: 'SS CIN', stat: '+0.5', arrow: 'up' }
  ];
</script>

<div class="preview">
  <header class="board-head">
    <span>RANK</span>
    <span>PLAYER</span>
    <span>POS</span>
    <span>±wRC+</span>
  </header>
  <ul class="board">
    {#each ROWS as r (r.rank)}
      <li class="row">
        <span class="cell rank">{r.rank}</span>
        <span class="cell name">{r.name}</span>
        <span class="cell pos">{r.pos}</span>
        <span class="cell stat" data-arrow={r.arrow}>
          {#if r.arrow === 'up'}<span class="arr up" aria-hidden="true">▲</span>{/if}
          {#if r.arrow === 'down'}<span class="arr down" aria-hidden="true">▼</span>{/if}
          {r.stat}
        </span>
      </li>
    {/each}
  </ul>
</div>

<style>
  .preview {
    height: 100%;
    width: 100%;
    padding: 8px 12px;
    background: linear-gradient(180deg, #0d0726 0%, #06031a 100%);
    color: var(--accent-amber);
    font-family: var(--font-mono);
    font-size: 9px;
    line-height: 1.3;
    letter-spacing: 0.06em;
    display: flex;
    flex-direction: column;
    gap: 3px;
  }

  .board-head {
    display: grid;
    grid-template-columns: 28px 1fr 56px 44px;
    gap: 6px;
    color: rgba(250, 199, 117, 0.45);
    text-transform: uppercase;
    padding-bottom: 3px;
    border-bottom: 0.5px dashed rgba(250, 199, 117, 0.25);
  }

  .board {
    list-style: none;
    margin: 0;
    padding: 0;
    display: flex;
    flex-direction: column;
    gap: 1px;
    flex: 1;
  }

  .row {
    display: grid;
    grid-template-columns: 28px 1fr 56px 44px;
    gap: 6px;
    align-items: baseline;
  }

  .cell {
    color: var(--accent-amber);
    text-shadow: 0 0 8px color-mix(in oklab, var(--accent-amber), transparent 80%);
  }

  .name {
    font-weight: 500;
  }

  .pos {
    color: rgba(250, 199, 117, 0.72);
  }

  .stat {
    text-align: right;
    font-variant-numeric: tabular-nums;
    display: inline-flex;
    align-items: center;
    justify-content: flex-end;
    gap: 4px;
  }

  .arr {
    font-size: 7px;
  }
  .arr.up {
    color: var(--accent-teal);
  }
  .arr.down {
    color: var(--accent-pink);
  }
</style>
