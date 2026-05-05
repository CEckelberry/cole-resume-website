<!--
  ModeToggle — small button that flips dark/light mode.
  Uses inline SVG icons (no emoji per design rule). The current mode is
  inferred from the document attribute on mount and kept in sync via a
  MutationObserver so external changes (devtools, other tabs) are reflected.
-->
<script lang="ts">
  import { onMount } from 'svelte';
  import { getMode, toggleMode, type Mode } from '$lib/theme/mode';

  let mode = $state<Mode>('dark');

  onMount(() => {
    mode = getMode();
    const obs = new MutationObserver(() => (mode = getMode()));
    obs.observe(document.documentElement, { attributes: true, attributeFilter: ['data-mode'] });

    // Cross-tab sync via localStorage events.
    const onStorage = (e: StorageEvent) => {
      if (e.key === 'mode' && (e.newValue === 'light' || e.newValue === 'dark')) {
        document.documentElement.setAttribute('data-mode', e.newValue);
        mode = e.newValue;
      }
    };
    window.addEventListener('storage', onStorage);

    return () => {
      obs.disconnect();
      window.removeEventListener('storage', onStorage);
    };
  });

  function flip() {
    mode = toggleMode();
  }
</script>

<button
  type="button"
  class="toggle"
  onclick={flip}
  aria-label={mode === 'dark' ? 'switch to light mode' : 'switch to dark mode'}
  aria-pressed={mode === 'light'}
  title={mode === 'dark' ? 'light mode' : 'dark mode'}
>
  {#if mode === 'dark'}
    <!-- moon -->
    <svg viewBox="0 0 16 16" fill="none" aria-hidden="true">
      <path
        d="M11.5 9.5 A5 5 0 0 1 6.5 4.5 A5 5 0 0 0 11.5 13.5 A5 5 0 0 0 13.7 13 A5 5 0 0 1 11.5 9.5 Z"
        fill="currentColor"
      />
    </svg>
  {:else}
    <!-- sun -->
    <svg viewBox="0 0 16 16" fill="none" aria-hidden="true">
      <circle cx="8" cy="8" r="3" fill="currentColor" />
      <g stroke="currentColor" stroke-width="1.2" stroke-linecap="round">
        <path d="M8 2 L8 3.5" />
        <path d="M8 12.5 L8 14" />
        <path d="M2 8 L3.5 8" />
        <path d="M12.5 8 L14 8" />
        <path d="M3.8 3.8 L4.9 4.9" />
        <path d="M11.1 11.1 L12.2 12.2" />
        <path d="M3.8 12.2 L4.9 11.1" />
        <path d="M11.1 4.9 L12.2 3.8" />
      </g>
    </svg>
  {/if}
</button>

<style>
  .toggle {
    display: inline-flex;
    align-items: center;
    justify-content: center;
    width: 28px;
    height: 28px;
    margin-left: 4px;
    padding: 0;
    background: transparent;
    border: 0.5px solid transparent;
    border-radius: var(--radius-pill);
    color: var(--text-secondary);
    cursor: pointer;
    transition:
      background-color var(--dur-fast) var(--ease-out),
      color var(--dur-fast) var(--ease-out),
      border-color var(--dur-fast) var(--ease-out);
  }

  .toggle:hover {
    background: var(--bg-surface-3);
    color: var(--text-primary);
    border-color: var(--border-default);
  }

  .toggle svg {
    width: 14px;
    height: 14px;
  }
</style>
