<script lang="ts">
  import '../app.css';
  import { onNavigate } from '$app/navigation';
  import Nav from '$lib/components/shell/Nav.svelte';
  import Footer from '$lib/components/shell/Footer.svelte';
  import MeshBackground from '$lib/components/effects/MeshBackground.svelte';
  import PolygonField from '$lib/components/effects/PolygonField.svelte';

  let { children } = $props();

  // View transitions: 240 ms cross-fade in browsers that support it.
  // Browsers without support fall through to the regular SvelteKit
  // navigation. Honor prefers-reduced-motion by skipping the transition.
  //
  // IMPORTANT: call `document.startViewTransition(...)` as a method, NOT
  // by extracting the function reference and calling it standalone. The
  // method needs its `this` bound to `document` — pulling it into a
  // const drops the binding and the call throws on every navigation,
  // which leaves SvelteKit's nav promise stuck and breaks subsequent
  // links until reload.
  type DocWithVT = Document & {
    startViewTransition?: (cb: () => Promise<void> | void) => { finished: Promise<void> };
  };

  onNavigate((navigation) => {
    if (typeof document === 'undefined') return;
    const doc = document as DocWithVT;
    if (typeof doc.startViewTransition !== 'function') return;
    if (window.matchMedia('(prefers-reduced-motion: reduce)').matches) return;

    return new Promise<void>((resolve) => {
      try {
        doc.startViewTransition!(async () => {
          resolve();
          await navigation.complete;
        });
      } catch {
        // If the API throws for any reason, unblock the navigation so
        // the page still updates. Worst case: no transition this time.
        resolve();
      }
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
