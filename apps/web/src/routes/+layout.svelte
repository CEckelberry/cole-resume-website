<script lang="ts">
  import '../app.css';
  import Nav from '$lib/components/shell/Nav.svelte';
  import Footer from '$lib/components/shell/Footer.svelte';
  import MeshBackground from '$lib/components/effects/MeshBackground.svelte';
  import PolygonField from '$lib/components/effects/PolygonField.svelte';

  let { children } = $props();
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
