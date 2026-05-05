<!--
  HeroName — display type "Cole Eckelberry" with the last name rendered as a
  pink → purple → teal gradient at 135°. The full name must fit on one line
  at every viewport ≥320px wide; we use a tighter clamp floor than the
  generic --type-display token to guarantee that at 320px.
-->
<script lang="ts">
  // No props at v1.
</script>

<h1 class="hero-name" aria-label="Cole Eckelberry">
  <span class="first">Cole</span>
  <span class="space" aria-hidden="true"> </span>
  <span class="last">Eckelberry</span>
</h1>

<style>
  .hero-name {
    /* Wider clamp floor than --type-display (40px) so the full 15-char name
       fits on a 320px viewport at the smallest size. The peak (76px) and
       8vw scaling match the design token. */
    font-size: clamp(34px, 8vw, 76px);
    line-height: 0.92;
    font-family: var(--font-sans);
    font-weight: 500;
    letter-spacing: -0.025em;
    margin: 0;
    color: var(--text-primary);
    white-space: nowrap;
  }

  .first,
  .last {
    display: inline;
  }

  .last {
    background: linear-gradient(
      135deg,
      var(--accent-pink) 0%,
      var(--accent-purple) 50%,
      var(--accent-teal) 100%
    );
    -webkit-background-clip: text;
    background-clip: text;
    -webkit-text-fill-color: transparent;
    color: transparent;
  }

  /* Forced-colors / Windows high-contrast: gradient text falls back to a
     solid system color so the name remains readable. */
  @media (forced-colors: active) {
    .last {
      background: none;
      -webkit-text-fill-color: currentColor;
      color: CanvasText;
    }
  }
</style>
