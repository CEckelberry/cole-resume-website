<script lang="ts">
  // Phase 4: real markdown blog. v1 ships this placeholder so the writing
  // nav link doesn't 404 — and so visitors who land here have a reason to
  // come back (or follow elsewhere) instead of bouncing.
</script>

<svelte:head>
  <title>writing — cole eckelberry</title>
  <meta
    name="description"
    content="A journal of build notes from the four side projects on this site. Coming soon."
  />
</svelte:head>

<section class="placeholder">
  <p class="eyebrow">section · writing</p>
  <h1>
    A journal of build notes lives <em>here.</em>
  </h1>
  <p class="lead">
    Each side project on this site gets a writing thread — design choices, things that surprised me,
    postmortems when something falls over. The first post lands when the first project ships.
  </p>

  <ul class="links">
    <li>
      <a href="/#work">see the projects this will write about</a>
      <span aria-hidden="true">→</span>
    </li>
    <li>
      <a href="https://github.com/CEckelberry" rel="me noopener">follow on github in the meantime</a
      >
      <span aria-hidden="true">→</span>
    </li>
  </ul>

  <p class="note">
    There's no email list. If something I write is worth coming back for, an RSS feed will be on
    this page too.
  </p>
</section>

<style>
  .placeholder {
    display: flex;
    flex-direction: column;
    gap: var(--space-5);
    max-width: var(--container-prose);
    padding-block: var(--space-7) var(--space-8);
  }

  .eyebrow {
    margin: 0;
    font-family: var(--font-mono);
    font-size: var(--type-micro);
    text-transform: uppercase;
    letter-spacing: 0.18em;
    color: var(--text-tertiary);
  }

  h1 {
    margin: 0;
    font-family: var(--font-sans);
    font-weight: 500;
    font-size: var(--type-h1);
    line-height: var(--type-h1-lh);
    letter-spacing: -0.02em;
    color: var(--text-primary);
  }
  h1 em {
    font-family: var(--font-serif);
    font-style: italic;
    font-weight: 400;
    color: var(--accent-pink-soft);
  }

  .lead {
    margin: 0;
    font-family: var(--font-serif);
    font-style: italic;
    font-size: var(--type-lead);
    line-height: var(--type-lead-lh);
    color: var(--text-secondary);
    max-width: 540px;
  }

  .links {
    list-style: none;
    margin: var(--space-3) 0 0;
    padding: 0;
    display: flex;
    flex-direction: column;
    gap: var(--space-3);
  }

  .links li {
    display: flex;
    align-items: baseline;
    gap: 8px;
  }

  .links a {
    font-family: var(--font-sans);
    font-size: var(--type-body);
    color: var(--text-primary);
    text-decoration: underline;
    text-decoration-color: var(--accent-pink-soft);
    text-decoration-thickness: 1px;
    text-underline-offset: 4px;
    transition: text-decoration-color var(--dur-fast) var(--ease-out);
  }
  .links a:hover {
    text-decoration-color: var(--accent-pink);
  }

  .note {
    margin: var(--space-3) 0 0;
    font-family: var(--font-mono);
    font-size: var(--type-tiny);
    color: var(--text-tertiary);
    letter-spacing: 0.04em;
    max-width: 540px;
  }
</style>
