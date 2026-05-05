<!--
  FilterChips — pill-shaped tab strip used above the projects gallery.
  Reusable for any filter set; consumer owns the active state.
-->
<script lang="ts" generics="T extends string">
  interface Props {
    filters: ReadonlyArray<{ id: T; label: string }>;
    active: T;
    onchange: (next: T) => void;
    /** Accessible label for the tablist. */
    ariaLabel?: string;
  }

  let { filters, active, onchange, ariaLabel = 'filters' }: Props = $props();
</script>

<div class="chips" role="tablist" aria-label={ariaLabel}>
  {#each filters as f (f.id)}
    <button
      type="button"
      class="chip"
      class:active={active === f.id}
      role="tab"
      aria-selected={active === f.id}
      onclick={() => onchange(f.id)}
    >
      {f.label}
    </button>
  {/each}
</div>

<style>
  .chips {
    display: flex;
    flex-wrap: wrap;
    gap: 8px;
  }

  .chip {
    appearance: none;
    border: 0.5px solid var(--border-default);
    background: var(--bg-surface-2);
    color: var(--text-secondary);
    padding: 6px 14px;
    border-radius: var(--radius-pill);
    font-family: var(--font-mono);
    font-size: var(--type-meta);
    line-height: 1.2;
    cursor: pointer;
    backdrop-filter: blur(10px);
    -webkit-backdrop-filter: blur(10px);
    transition:
      background-color var(--dur-fast) var(--ease-out),
      color var(--dur-fast) var(--ease-out),
      border-color var(--dur-fast) var(--ease-out);
  }

  .chip:hover {
    border-color: var(--border-strong);
    color: var(--text-primary);
  }

  .chip.active {
    background: var(--text-primary);
    color: var(--bg-canvas);
    border-color: transparent;
  }
</style>
