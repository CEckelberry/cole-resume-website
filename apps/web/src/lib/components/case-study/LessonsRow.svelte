<!--
  LessonsRow — two columns of "WORKED" and "DIDN'T" lessons.
  Each entry is a small card with a left-border accent and a short body.
-->
<script lang="ts">
  type Lesson = { title: string; body: string };

  interface Props {
    worked: Lesson[];
    didnt: Lesson[];
  }

  let { worked, didnt }: Props = $props();
</script>

<div class="grid">
  <section class="col" data-flavor="worked">
    <p class="head">worked</p>
    <ul>
      {#each worked as l, i (i)}
        <li>
          <p class="title">{l.title}</p>
          <p class="body">{l.body}</p>
        </li>
      {/each}
    </ul>
  </section>

  <section class="col" data-flavor="didnt">
    <p class="head">didn't</p>
    <ul>
      {#each didnt as l, i (i)}
        <li>
          <p class="title">{l.title}</p>
          <p class="body">{l.body}</p>
        </li>
      {/each}
    </ul>
  </section>
</div>

<style>
  .grid {
    display: grid;
    grid-template-columns: 1fr;
    gap: var(--space-5);
  }
  @media (min-width: 720px) {
    .grid {
      grid-template-columns: 1fr 1fr;
      gap: var(--space-6);
    }
  }

  .head {
    margin: 0 0 var(--space-3);
    font-family: var(--font-mono);
    font-size: var(--type-micro);
    text-transform: uppercase;
    letter-spacing: 0.18em;
    color: var(--text-tertiary);
  }

  ul {
    list-style: none;
    margin: 0;
    padding: 0;
    display: flex;
    flex-direction: column;
    gap: var(--space-3);
  }

  li {
    padding: var(--space-3) var(--space-4);
    border-left: 2px solid;
    border-radius: 0 var(--radius-md) var(--radius-md) 0;
    background: var(--bg-surface-2);
  }

  .col[data-flavor='worked'] li {
    border-left-color: var(--accent-teal);
  }
  .col[data-flavor='didnt'] li {
    border-left-color: var(--accent-pink);
  }

  .title {
    margin: 0 0 4px;
    font-family: var(--font-sans);
    font-size: var(--type-body-sm);
    font-weight: 500;
    color: var(--text-primary);
  }

  .body {
    margin: 0;
    font-size: var(--type-body-sm);
    line-height: 1.65;
    color: var(--text-secondary);
  }
</style>
