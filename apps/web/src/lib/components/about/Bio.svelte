<!--
  Bio — three-paragraph autobiography.
  First paragraph rendered in serif italic per DESIGN.md; rest in sans.
  Voice notes (DESIGN.md voice section): real numbers stay in copy, no
  marketing words.
-->
<script lang="ts">
  // No props.
</script>

<div class="bio">
  <p class="opener">
    I started as a systems engineer at GE in 2012, writing Bash and Perl to keep 7+ PB of backup
    data flowing across 9,200 servers — one of the largest backup environments in the world at the
    time. From there I moved through infrastructure and cloud roles at NIAID (1,400 federal-research
    clients, 5.5 PB of data), RoundTower, MediaMonks, Hex, and Pinterest before landing at EpexSpot
    as DevOps Technical Lead. Each stop taught a different flavor of "ops is half the product."
  </p>

  <p>
    My main languages have been <strong>Python</strong> for about eight years and
    <strong>JavaScript / TypeScript</strong> alongside it, with <strong>React</strong> and
    <strong>Svelte</strong> in professional settings. The infrastructure stack I lean on is the
    boring-on-purpose one: <strong>Kubernetes</strong>, <strong>Terraform</strong>,
    <strong>GCP</strong> and <strong>AWS</strong>, <strong>Postgres</strong>, <strong>Docker</strong>,
    plus enough Bash since 2012 that I will fight you about <code>set -euo pipefail</code>. Recently
    I've added <strong>Go</strong> (six months) and <strong>SvelteKit</strong> (two years) — this
    site is what I made with them.
  </p>

  <p>
    The four projects above are how I learn — by building things I'd actually want to use. They
    cover the surface area I'm interested in: real-time backends, browser-side WASM, observability,
    and the unglamorous infrastructure that holds all of it up.
  </p>

  <p class="closer">
    When I'm not at the keyboard you'll find me hiking with my dog Willy, cooking through old family
    recipes, or watching baseball.
  </p>
</div>

<style>
  .bio {
    display: flex;
    flex-direction: column;
    gap: var(--space-4);
    max-width: var(--container-prose);
  }

  .bio p {
    margin: 0;
    font-size: var(--type-body);
    line-height: 1.7;
    color: var(--text-secondary);
  }

  .opener {
    font-size: var(--type-lead);
    line-height: var(--type-lead-lh);
    color: var(--text-primary);
  }

  .closer {
    color: var(--text-tertiary);
  }
</style>
