<script lang="ts">
  import '../app.css';
  import { onNavigate } from '$app/navigation';
  import Nav from '$lib/components/shell/Nav.svelte';
  import Footer from '$lib/components/shell/Footer.svelte';
  import MeshBackground from '$lib/components/effects/MeshBackground.svelte';
  import PolygonField from '$lib/components/effects/PolygonField.svelte';

  let { children } = $props();

  // View transitions: 240ms cross-fade in browsers that support it. Browsers
  // without support fall through to the regular SvelteKit navigation. Honor
  // prefers-reduced-motion by skipping the transition entirely.
  onNavigate((navigation) => {
    if (typeof document === 'undefined') return;
    const startViewTransition = (
      document as Document & {
        startViewTransition?: (cb: () => Promise<void>) => { finished: Promise<void> };
      }
    ).startViewTransition;
    if (!startViewTransition) return;
    if (window.matchMedia('(prefers-reduced-motion: reduce)').matches) return;

    return new Promise<void>((resolve) => {
      startViewTransition(async () => {
        resolve();
        await navigation.complete;
      });
    });
  });
</script>

<svelte:head>
  <link rel="icon" type="image/svg+xml" href="/favicon.svg" />
</svelte:head>

<a href="#main" class="skip-link">skip to content</a>

<MeshBackground />
<PolygonField seed={20260505} count={5} />

<Nav />

<main id="main" class="page">
  {@render children()}
</main>

<Footer />

<style>
  .page {
    position: relative;
    z-index: 2;
    max-width: var(--container-page);
    margin: 0 auto;
    padding: 88px 24px 0; /* top padding leaves room for the floating nav */
  }

  @media (min-width: 768px) {
    .page {
      padding: 112px 48px 0;
    }
  }
</style>
