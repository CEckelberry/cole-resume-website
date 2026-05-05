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
    Python has been my main language for about eight years; JavaScript and TypeScript came right
    alongside it, with React in professional settings. Over a decade I worked through cloud and
    DevOps roles at NIAID, RoundTower, MediaMonks, Hex, and Pinterest. Today I'm a Senior DevOps /
    Platform Engineer based in the Bay Area.
  </p>

  <p>
    Frontend frameworks came later. SvelteKit clicked about two years ago — this site is built on
    it. Go I picked up six months ago, mostly to have a sharper tool for the small services in the
    projects above. The four projects there are how I learn — by building things I'd actually want
    to use.
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
