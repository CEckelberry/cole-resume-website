<!--
  MeshBackground — soft, blurred, drifting circular gradients positioned
  absolutely behind all content. Sits at z-0 in the layout stack.

  DESIGN.md spec:
    - 4 blobs across the visible viewport, never more
    - blur 60–80px, opacity 0.35–0.55
    - 16–22s ease-in-out infinite drift, ±40px translate, 0.95–1.10 scale
    - stagger so they never sync
    - prefers-reduced-motion: hold still
-->
<script lang="ts">
  // No props at v1. A future variant could accept per-section overrides.
</script>

<div class="mesh" aria-hidden="true">
  <div class="blob blob-1"></div>
  <div class="blob blob-2"></div>
  <div class="blob blob-3"></div>
  <div class="blob blob-4"></div>
</div>

<style>
  .mesh {
    position: fixed;
    inset: 0;
    z-index: 0;
    pointer-events: none;
    overflow: hidden;
  }

  .blob {
    position: absolute;
    border-radius: 50%;
    filter: blur(72px);
    will-change: transform;
  }

  /* Top-left: pink */
  .blob-1 {
    top: -8%;
    left: -6%;
    width: 520px;
    height: 520px;
    background: var(--accent-pink);
    opacity: 0.4;
    animation: drift-1 18s var(--ease-in-out) infinite;
  }

  /* Top-right: teal */
  .blob-2 {
    top: 6%;
    right: -8%;
    width: 580px;
    height: 580px;
    background: var(--accent-teal);
    opacity: 0.38;
    animation: drift-2 21s var(--ease-in-out) infinite;
    animation-delay: -7s;
  }

  /* Middle-left: purple */
  .blob-3 {
    top: 48%;
    left: -10%;
    width: 620px;
    height: 620px;
    background: var(--accent-purple);
    opacity: 0.42;
    animation: drift-3 19s var(--ease-in-out) infinite;
    animation-delay: -3s;
  }

  /* Bottom-right: amber */
  .blob-4 {
    bottom: -6%;
    right: -4%;
    width: 540px;
    height: 540px;
    background: var(--accent-amber);
    opacity: 0.45;
    animation: drift-4 22s var(--ease-in-out) infinite;
    animation-delay: -11s;
  }

  @keyframes drift-1 {
    0%,
    100% {
      transform: translate(0, 0) scale(1);
    }
    33% {
      transform: translate(28px, -22px) scale(1.06);
    }
    66% {
      transform: translate(-18px, 32px) scale(0.97);
    }
  }
  @keyframes drift-2 {
    0%,
    100% {
      transform: translate(0, 0) scale(1.02);
    }
    50% {
      transform: translate(-34px, 18px) scale(0.96);
    }
  }
  @keyframes drift-3 {
    0%,
    100% {
      transform: translate(0, 0) scale(0.98);
    }
    40% {
      transform: translate(36px, 24px) scale(1.08);
    }
    80% {
      transform: translate(-14px, -28px) scale(1.02);
    }
  }
  @keyframes drift-4 {
    0%,
    100% {
      transform: translate(0, 0) scale(1);
    }
    25% {
      transform: translate(-26px, -16px) scale(1.04);
    }
    75% {
      transform: translate(22px, 28px) scale(0.95);
    }
  }

  /* Light mode dials the saturation back so the soft canvas isn't overrun. */
  :global([data-mode='light']) .mesh .blob {
    opacity: 0.28;
    filter: blur(80px);
  }

  @media (prefers-reduced-motion: reduce) {
    .blob {
      animation: none !important;
    }
  }
</style>
