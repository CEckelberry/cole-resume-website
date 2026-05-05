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
    I started as a systems engineer at GE in 2012, writing Bash and Perl to keep one of the largest
    backup environments in the world running.
  </p>

  <p>
    Over a decade I worked my way through cloud and DevOps roles at NIAID, RoundTower, MediaMonks,
    Hex, and Pinterest, picking up Python and JavaScript along the way. Today I'm a Senior DevOps /
    Platform Engineer based in the Bay Area.
  </p>

  <p>
    I came to frontend the long way. SvelteKit clicked about two years ago. I started learning Go
    six months ago and have been using it ever since. The four projects above are how I learn — by
    building things I'd actually want to use.
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
    font-family: var(--font-serif);
    font-style: italic;
    font-size: var(--type-lead);
    line-height: var(--type-lead-lh);
    color: var(--text-primary);
  }

  .closer {
    color: var(--text-tertiary);
  }
</style>
