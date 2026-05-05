<!--
  Nav — floating glass pill at the top of every page.
  Spec from DESIGN.md → Components → Nav.

  Anatomy:
    [pulsing teal dot] deployed Nh ago · main@SHA  |  work · about · writing · contact
-->
<script lang="ts">
  import { BUILD_SHA, BUILD_TIME, timeAgo } from '$lib/utils/buildInfo';
  import { onMount } from 'svelte';
  import ModeToggle from './ModeToggle.svelte';

  // The "deployed N ago" string is computed on the client so it stays fresh
  // without invalidating the prerendered HTML. SSR shows a static fallback
  // (computed once at server-render time); the client reconciles on mount.
  let ago = $state(timeAgo(BUILD_TIME));

  onMount(() => {
    ago = timeAgo(BUILD_TIME);
    const id = setInterval(() => (ago = timeAgo(BUILD_TIME)), 30_000);
    return () => clearInterval(id);
  });

  type LinkDef = { href: string; label: string };
  const links: LinkDef[] = [
    { href: '/#work', label: 'work' },
    { href: '/#about', label: 'about' },
    { href: '/journal', label: 'writing' },
    { href: '/#contact', label: 'contact' }
  ];
</script>

<header class="nav-wrap">
  <nav class="nav" aria-label="primary">
    <div class="status">
      <span class="dot" aria-hidden="true"></span>
      <span class="status-text">
        <span class="dim">deployed</span>
        <span class="value">{ago}</span>
        <span class="sep" aria-hidden="true">·</span>
        <span class="dim">main@</span><span class="value">{BUILD_SHA}</span>
      </span>
    </div>

    <ul class="links">
      {#each links as link (link.href)}
        <li>
          <a href={link.href}>{link.label}</a>
        </li>
      {/each}
    </ul>

    <ModeToggle />
  </nav>
</header>

<style>
  .nav-wrap {
    position: fixed;
    top: 16px;
    left: 0;
    right: 0;
    display: flex;
    justify-content: center;
    z-index: 30;
    padding: 0 16px;
    pointer-events: none;
  }

  .nav {
    display: flex;
    align-items: center;
    gap: 18px;
    padding: 8px 14px;
    background: var(--bg-surface-2);
    border: 0.5px solid var(--border-default);
    border-radius: var(--radius-pill);
    backdrop-filter: blur(16px) saturate(160%);
    -webkit-backdrop-filter: blur(16px) saturate(160%);
    pointer-events: auto;
    max-width: min(100%, 720px);
    color: var(--text-primary);
  }

  .status {
    display: flex;
    align-items: center;
    gap: 8px;
    padding-right: 10px;
    border-right: 0.5px solid var(--border-subtle);
    white-space: nowrap;
  }

  .dot {
    width: 6px;
    height: 6px;
    border-radius: 50%;
    background: var(--accent-teal);
    box-shadow: 0 0 0 0 color-mix(in oklab, var(--accent-teal), transparent 60%);
    animation: pulse 2s var(--ease-in-out) infinite;
  }

  @keyframes pulse {
    0%,
    100% {
      box-shadow: 0 0 0 0 color-mix(in oklab, var(--accent-teal), transparent 60%);
    }
    50% {
      box-shadow: 0 0 0 6px color-mix(in oklab, var(--accent-teal), transparent 100%);
    }
  }

  @media (prefers-reduced-motion: reduce) {
    .dot {
      animation: none;
    }
  }

  .status-text {
    font-family: var(--font-mono);
    font-size: var(--type-tiny);
    line-height: 1.4;
    letter-spacing: 0.06em;
    text-transform: lowercase;
  }
  .status-text .dim {
    color: var(--text-tertiary);
  }
  .status-text .value {
    color: var(--text-primary);
  }
  .status-text .sep {
    color: var(--text-muted);
    margin: 0 4px;
  }

  .links {
    list-style: none;
    margin: 0;
    padding: 0;
    display: flex;
    align-items: center;
    gap: 4px;
  }

  .links a {
    display: inline-block;
    padding: 6px 10px;
    border-radius: var(--radius-pill);
    color: var(--text-secondary);
    text-decoration: none;
    font-size: var(--type-meta);
    line-height: 1.2;
    letter-spacing: 0.01em;
    transition:
      background-color var(--dur-fast) var(--ease-out),
      color var(--dur-fast) var(--ease-out);
  }

  .links a:hover {
    background: var(--bg-surface-3);
    color: var(--text-primary);
  }

  /* Hide the build-info text on very small screens but keep the dot as a
     subtle "we're alive" indicator. */
  @media (max-width: 540px) {
    .status-text {
      display: none;
    }
    .status {
      padding-right: 6px;
      border-right: none;
    }
    .nav {
      gap: 8px;
    }
  }
</style>
