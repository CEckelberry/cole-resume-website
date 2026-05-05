<!--
  PolygonField — sparse polygonal shapes at low opacity, behind content but
  in front of the mesh. Procedural placement seeded by a numeric `seed` so
  re-renders are stable (no Math.random flicker between SSR and hydration).

  At v1: 4–6 polygons per viewport-height of layout. No scroll-tied animation.
-->
<script lang="ts">
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
    rotate: number;
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

      const pts: string[] = [];
      for (let s = 0; s < sides; s++) {
        const a = (s / sides) * Math.PI * 2 - Math.PI / 2;
        const px = Math.cos(a) * r;
        const py = Math.sin(a) * r;
        pts.push(`${px.toFixed(1)},${py.toFixed(1)}`);
      }
      out.push({ points: pts.join(' '), accent, opacity, rotate, cx, cy });
    }
    return out;
  })();

  const accentVar: Record<AccentClass, string> = {
    pink: 'var(--accent-pink)',
    teal: 'var(--accent-teal)',
    purple: 'var(--accent-purple)',
    amber: 'var(--accent-amber)'
  };
</script>

<svg class="polygon-field" viewBox="0 0 100 100" preserveAspectRatio="none" aria-hidden="true">
  {#each polys as p (`${p.cx}-${p.cy}-${p.points}`)}
    <g transform="translate({p.cx} {p.cy}) rotate({p.rotate})">
      <polygon
        points={p.points}
        fill={accentVar[p.accent]}
        opacity={p.opacity}
        transform="scale(0.35)"
      />
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
