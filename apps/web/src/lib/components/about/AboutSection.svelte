<!--
  AboutSection — section 03 of the home page.
  Section header + two-column layout: bio on the left, activity card on the right.
  Single column under 900px.
-->
<script lang="ts">
  import Bio from './Bio.svelte';
  import ActivityCard from './ActivityCard.svelte';
</script>

<section id="about" class="section">
  <header class="section-head">
    <p class="eyebrow">section 03 · about</p>
    <h2>
      Senior DevOps. Long-time scripter. Recent <em>frontend convert.</em>
    </h2>
  </header>

  <div class="grid">
    <Bio />
    <ActivityCard />
  </div>
</section>

<style>
  .section {
    padding: var(--space-8) 0;
    display: flex;
    flex-direction: column;
    gap: var(--space-6);
  }

  .section-head {
    display: flex;
    flex-direction: column;
    gap: var(--space-3);
    max-width: var(--container-prose);
  }

  .eyebrow {
    font-family: var(--font-mono);
    font-size: var(--type-micro);
    line-height: var(--type-micro-lh);
    text-transform: uppercase;
    letter-spacing: 0.18em;
    color: var(--text-tertiary);
    margin: 0;
  }

  h2 {
    font-family: var(--font-sans);
    font-weight: 500;
    font-size: var(--type-h2);
    line-height: var(--type-h2-lh);
    letter-spacing: -0.01em;
    color: var(--text-primary);
    margin: 0;
  }

  .grid {
    display: grid;
    grid-template-columns: 1fr;
    gap: var(--space-6);
    align-items: start;
  }

  @media (min-width: 900px) {
    .grid {
      grid-template-columns: minmax(0, 1.05fr) minmax(0, 0.95fr);
      gap: var(--space-7);
    }
  }
</style>
