<!--
  PolygonField — sparse polygonal shapes at low opacity, behind content but
  in front of the mesh. Procedural placement seeded by a numeric `seed` so
  re-renders are stable (no Math.random flicker between SSR and hydration).

  Scroll-driven rotation: each polygon has its own `rate` (deg/px), some
  positive (clockwise), some negative (counter), so the field never feels
  synced. The base rotation is set on the SVG transform attribute so the
  SSR output already looks intentional; the client takes over on first
  scroll. prefers-reduced-motion skips the listener entirely.
-->
<script lang="ts">
  import { onMount } from 'svelte';

  type AccentClass = 'pink' | 'teal' | 'purple' | 'amber';

  interface Props {
    /** Numeric seed for deterministic placement; vary across pages. */
    seed?: number;
    /** Polygon count, 4–6. Defaults to 5. */
    count?: number;
  }

  let { seed = 1, count = 5 }: Props = $props();

  // Cheap mulberry32 PRNG for stable per-seed layout. Server and client
  // resolve to the same numbers, so the SVG output never reflows on hydrate.
  function makeRng(s: number) {
    let a = s | 0;
    return () => {
      a |= 0;
      a = (a + 0x6d2b79f5) | 0;
      let t = Math.imul(a ^ (a >>> 15), 1 | a);
      t = (t + Math.imul(t ^ (t >>> 7), 61 | t)) ^ t;
      return ((t ^ (t >>> 14)) >>> 0) / 4294967296;
    };
  }

  const accents: AccentClass[] = ['pink', 'teal', 'purple', 'amber'];

  type Poly = {
    points: string;
    accent: AccentClass;
    opacity: number;
    /** Static rotation offset, applied at SSR time. */
    rotate: number;
    /** Scroll rate in deg/px; ±0.04 to ±0.18, signed so directions vary. */
    rate: number;
    cx: number;
    cy: number;
  };

  const polys: Poly[] = (() => {
    const rng = makeRng(seed);
    const out: Poly[] = [];
    for (let i = 0; i < count; i++) {
      const sides = 3 + Math.floor(rng() * 3); // 3, 4, or 5
      const r = 24 + rng() * 56;
      const cx = 6 + rng() * 88; // %
      const cy = 6 + rng() * 88; // %
      const rotate = rng() * 360;
      const accent = accents[Math.floor(rng() * accents.length)];
      const opacity = 0.14 + rng() * 0.21;
      // Rate magnitude 0.04–0.18 deg/px, sign randomized so the field
      // contains a mix of clockwise and counter-clockwise spinners.
      const rateMag = 0.04 + rng() * 0.14;
      const rate = rng() < 0.5 ? rateMag : -rateMag;

      const pts: string[] = [];
      for (let s = 0; s < sides; s++) {
        const a = (s / sides) * Math.PI * 2 - Math.PI / 2;
        const px = Math.cos(a) * r;
        const py = Math.sin(a) * r;
        pts.push(`${px.toFixed(1)},${py.toFixed(1)}`);
      }
      out.push({ points: pts.join(' '), accent, opacity, rotate, rate, cx, cy });
    }
    return out;
  })();

  const accentVar: Record<AccentClass, string> = {
    pink: 'var(--accent-pink)',
    teal: 'var(--accent-teal)',
    purple: 'var(--accent-purple)',
    amber: 'var(--accent-amber)'
  };

  // Refs for the inner rotor groups; the scroll handler mutates their
  // transform attributes directly to skip Svelte's scheduler on a hot loop.
  let rotorEls: SVGGElement[] = $state([]);

  onMount(() => {
    // SSR has already set the base rotation. Reduced-motion stops here.
    if (typeof window === 'undefined') return;
    if (window.matchMedia('(prefers-reduced-motion: reduce)').matches) return;

    let scheduled = false;
    let rafId = 0;

    const apply = () => {
      scheduled = false;
      const y = window.scrollY;
      for (let i = 0; i < polys.length; i++) {
        const el = rotorEls[i];
        if (!el) continue;
        const angle = polys[i].rotate + y * polys[i].rate;
        el.setAttribute('transform', `rotate(${angle.toFixed(2)})`);
      }
    };

    const onScroll = () => {
      if (scheduled) return;
      scheduled = true;
      rafId = requestAnimationFrame(apply);
    };

    apply();
    window.addEventListener('scroll', onScroll, { passive: true });

    return () => {
      window.removeEventListener('scroll', onScroll);
      cancelAnimationFrame(rafId);
    };
  });
</script>

<svg class="polygon-field" viewBox="0 0 100 100" preserveAspectRatio="none" aria-hidden="true">
  {#each polys as p, i (`${p.cx}-${p.cy}-${p.points}`)}
    <g transform="translate({p.cx} {p.cy})">
      <g bind:this={rotorEls[i]} transform="rotate({p.rotate})">
        <polygon
          points={p.points}
          fill={accentVar[p.accent]}
          opacity={p.opacity}
          transform="scale(0.35)"
        />
      </g>
    </g>
  {/each}
</svg>

<style>
  .polygon-field {
    position: fixed;
    inset: 0;
    width: 100vw;
    height: 100vh;
    z-index: 1;
    pointer-events: none;
  }

  :global([data-mode='light']) .polygon-field {
    opacity: 0.65;
  }
</style>
