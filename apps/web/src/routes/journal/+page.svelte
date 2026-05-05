<script lang="ts">
  import type { PageProps } from './$types';
  import SEO from '$lib/components/shell/SEO.svelte';
  import { formatPostDate } from '$lib/content/journal';

  let { data }: PageProps = $props();
  const hasPosts = $derived(data.posts.length > 0);
</script>

<SEO
  title="writing"
  description="Build notes from the four side projects on this site — design choices, things that surprised me, postmortems when something falls over."
  path="/journal"
/>

{#if hasPosts}
  <section class="journal">
    <header class="head">
      <p class="eyebrow">section · writing</p>
      <h1>
        Build notes from the <em>side projects.</em>
      </h1>
      <p class="lead">
        Design choices, decisions I had to pick a side on, things that surprised me. New posts as
        each project ships.
      </p>
    </header>

    <ul class="posts">
      {#each data.posts as post (post.slug)}
        <li class="post">
          <a href="/journal/{post.slug}">
            <time datetime={post.date} class="date">{formatPostDate(post.date)}</time>
            <h2 class="title">
              {post.title}
              {#if post.italicWord}
                <em>{post.italicWord}</em>
              {/if}
            </h2>
            <p class="summary">{post.summary}</p>
            {#if post.tags?.length}
              <ul class="tags">
                {#each post.tags as tag (tag)}
                  <li>{tag}</li>
                {/each}
              </ul>
            {/if}
          </a>
        </li>
      {/each}
    </ul>

    <p class="rss">
      <a href="/journal/rss.xml" rel="alternate" type="application/rss+xml">rss</a>
    </p>
  </section>
{:else}
  <section class="empty">
    <p class="eyebrow">section · writing</p>
    <h1>
      Build notes will live <em>here.</em>
    </h1>
    <p class="lead">
      Each side project on this site gets a writing thread — design choices, things that surprised
      me, postmortems when something falls over. The first post lands when the first project ships.
    </p>

    <ul class="links">
      <li>
        <a href="/#work">see the projects this will write about</a>
        <span aria-hidden="true">→</span>
      </li>
      <li>
        <a href="https://github.com/CEckelberry" rel="me noopener"
          >follow on github in the meantime</a
        >
        <span aria-hidden="true">→</span>
      </li>
      <li>
        <a href="/journal/rss.xml" rel="alternate" type="application/rss+xml">subscribe via rss</a>
        <span aria-hidden="true">→</span>
      </li>
    </ul>
  </section>
{/if}

<style>
  .journal,
  .empty {
    display: flex;
    flex-direction: column;
    gap: var(--space-6);
    max-width: var(--container-prose);
    padding-block: var(--space-7) var(--space-8);
  }

  .head,
  .empty {
    display: flex;
    flex-direction: column;
    gap: var(--space-4);
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

  /* --- post list --- */

  .posts {
    list-style: none;
    margin: 0;
    padding: 0;
    display: flex;
    flex-direction: column;
    border-top: 0.5px solid var(--border-subtle);
  }

  .post a {
    display: flex;
    flex-direction: column;
    gap: 8px;
    padding: var(--space-5) 0;
    border-bottom: 0.5px solid var(--border-subtle);
    text-decoration: none;
    color: inherit;
    transition:
      background-color var(--dur-fast) var(--ease-out),
      padding-left var(--dur-default) var(--ease-out);
  }
  .post a:hover {
    background: var(--bg-surface-2);
    padding-left: var(--space-3);
  }

  .date {
    font-family: var(--font-mono);
    font-size: var(--type-tiny);
    text-transform: uppercase;
    letter-spacing: 0.16em;
    color: var(--text-tertiary);
  }

  .title {
    margin: 0;
    font-family: var(--font-sans);
    font-weight: 500;
    font-size: var(--type-h3);
    line-height: var(--type-h3-lh);
    letter-spacing: -0.005em;
    color: var(--text-primary);
  }
  .title em {
    font-family: var(--font-serif);
    font-style: italic;
    font-weight: 400;
    color: var(--accent-pink-soft);
    margin-left: 4px;
  }

  .summary {
    margin: 0;
    font-size: var(--type-body-sm);
    line-height: 1.6;
    color: var(--text-secondary);
    max-width: 560px;
  }

  .tags {
    list-style: none;
    margin: 4px 0 0;
    padding: 0;
    display: flex;
    flex-wrap: wrap;
    gap: 6px;
  }
  .tags li {
    font-family: var(--font-mono);
    font-size: var(--type-tiny);
    color: var(--text-tertiary);
    padding: 2px 8px;
    border-radius: var(--radius-sm);
    background: var(--bg-surface-2);
    border: 0.5px solid var(--border-subtle);
    letter-spacing: 0.04em;
  }

  .rss {
    margin: var(--space-3) 0 0;
    font-family: var(--font-mono);
    font-size: var(--type-tiny);
    text-transform: uppercase;
    letter-spacing: 0.16em;
  }
  .rss a {
    color: var(--text-tertiary);
    text-decoration: underline;
    text-decoration-color: var(--accent-pink-soft);
    text-underline-offset: 4px;
  }
  .rss a:hover {
    color: var(--accent-pink);
  }

  /* --- empty state --- */

  .empty .links {
    list-style: none;
    margin: var(--space-3) 0 0;
    padding: 0;
    display: flex;
    flex-direction: column;
    gap: var(--space-3);
  }
  .empty .links li {
    display: flex;
    align-items: baseline;
    gap: 8px;
  }
  .empty .links a {
    font-family: var(--font-sans);
    font-size: var(--type-body);
    color: var(--text-primary);
    text-decoration: underline;
    text-decoration-color: var(--accent-pink-soft);
    text-underline-offset: 4px;
    transition: text-decoration-color var(--dur-fast) var(--ease-out);
  }
  .empty .links a:hover {
    text-decoration-color: var(--accent-pink);
  }
</style>
