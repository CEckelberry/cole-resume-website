<!--
  Custom error page. SvelteKit renders this for any `error` thrown during
  load() or navigation, including 404s.
  Cotton-candy aesthetic; gives the visitor a way back instead of a wall.
-->
<script lang="ts">
  import { page } from '$app/state';

  const status = $derived(page.status);
  const message = $derived(page.error?.message ?? '');

  const isNotFound = $derived(status === 404);
</script>

<svelte:head>
  <title>{status} — cole eckelberry</title>
  <meta name="robots" content="noindex" />
</svelte:head>

<section class="error">
  <p class="eyebrow">error · {status}</p>

  {#if isNotFound}
    <h1>
      That page lives <em>somewhere else.</em>
    </h1>
    <p class="lead">
      Either the link rotted or it's a route that hasn't shipped yet. The home page is a good place
      to land while we figure it out.
    </p>
  {:else}
    <h1>
      Something <em>broke.</em>
    </h1>
    <p class="lead">
      The server returned a {status}. {message ? `(${message})` : ''} Refresh in a minute, or use one
      of the links below to keep moving.
    </p>
  {/if}

  <ul class="links">
    <li><a href="/">home</a><span aria-hidden="true">→</span></li>
    <li><a href="/#work">see the projects</a><span aria-hidden="true">→</span></li>
    <li><a href="/#about">about</a><span aria-hidden="true">→</span></li>
    <li><a href="/#contact">say hi</a><span aria-hidden="true">→</span></li>
  </ul>

  <p class="footnote">
    if you got here from a link on this site, that's a bug — please mention it on the contact form
    so I can fix it.
  </p>
</section>

<style>
  .error {
    display: flex;
    flex-direction: column;
    gap: var(--space-5);
    max-width: var(--container-prose);
    padding-block: var(--space-8) var(--space-7);
    min-height: calc(100dvh - 280px);
  }

  .eyebrow {
    margin: 0;
    font-family: var(--font-mono);
    font-size: var(--type-micro);
    text-transform: uppercase;
    letter-spacing: 0.18em;
    color: var(--accent-coral);
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
    gap: var(--space-2);
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
    text-underline-offset: 4px;
    transition: text-decoration-color var(--dur-fast) var(--ease-out);
  }
  .links a:hover {
    text-decoration-color: var(--accent-pink);
  }

  .footnote {
    margin: var(--space-3) 0 0;
    font-family: var(--font-mono);
    font-size: var(--type-tiny);
    color: var(--text-tertiary);
    letter-spacing: 0.04em;
    max-width: 540px;
  }
</style>
