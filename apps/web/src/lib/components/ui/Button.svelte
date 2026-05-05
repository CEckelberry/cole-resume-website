<!--
  Button — primary + secondary glass-pill variants.
  Magnetic cursor pull is a v2 enhancement; v1 uses a CSS hover-lift.
-->
<script lang="ts">
  import type { Snippet } from 'svelte';
  import type { HTMLAnchorAttributes, HTMLButtonAttributes } from 'svelte/elements';

  interface CommonProps {
    variant?: 'primary' | 'secondary';
    size?: 'md' | 'lg';
    children: Snippet;
  }

  type Props =
    | (CommonProps & { href: string } & Omit<HTMLAnchorAttributes, 'href'>)
    | (CommonProps & { href?: undefined } & HTMLButtonAttributes);

  let {
    variant = 'primary',
    size = 'md',
    children,
    href,
    class: className = '',
    ...rest
  }: Props = $props();
</script>

{#if href}
  <a {href} class="btn btn-{variant} btn-{size} {className}" {...rest as HTMLAnchorAttributes}>
    {@render children()}
  </a>
{:else}
  <button class="btn btn-{variant} btn-{size} {className}" {...rest as HTMLButtonAttributes}>
    {@render children()}
  </button>
{/if}

<style>
  .btn {
    display: inline-flex;
    align-items: center;
    gap: 8px;
    border: 0.5px solid var(--border-default);
    border-radius: var(--radius-pill);
    font-family: var(--font-sans);
    font-weight: 500;
    text-decoration: none;
    cursor: pointer;
    user-select: none;
    transition:
      transform var(--dur-default) var(--ease-spring),
      background-color var(--dur-default) var(--ease-out),
      border-color var(--dur-default) var(--ease-out),
      box-shadow var(--dur-default) var(--ease-out),
      color var(--dur-default) var(--ease-out);
  }

  .btn-md {
    padding: 10px 18px;
    font-size: var(--type-meta);
    line-height: 1.2;
  }

  .btn-lg {
    padding: 13px 22px;
    font-size: var(--type-body-sm);
    line-height: 1.2;
  }

  /* Primary: solid white-ish surface, dark text — high-contrast on the dark canvas */
  .btn-primary {
    background: var(--text-primary);
    color: var(--bg-canvas);
    border-color: transparent;
  }
  .btn-primary:hover {
    transform: translateY(-2px);
    box-shadow: 0 8px 20px -8px color-mix(in oklab, var(--accent-pink), transparent 40%);
  }

  /* Secondary: glass surface */
  .btn-secondary {
    background: var(--bg-surface-2);
    color: var(--text-primary);
    backdrop-filter: blur(12px);
    -webkit-backdrop-filter: blur(12px);
  }
  .btn-secondary:hover {
    background: var(--bg-surface-3);
    border-color: var(--border-strong);
    transform: translateY(-2px);
  }

  .btn:active {
    transform: translateY(-1px);
  }

  @media (prefers-reduced-motion: reduce) {
    .btn {
      transition-duration: 0ms;
    }
    .btn:hover,
    .btn:active {
      transform: none;
    }
  }
</style>
