<script lang="ts">
  import type { PageProps } from './$types';
  import SEO from '$lib/components/shell/SEO.svelte';
  import { formatPostDate } from '$lib/content/journal';

  let { data }: PageProps = $props();
  const Content = $derived(data.Content);
  const seoTitle = $derived(
    `${data.meta.title}${data.meta.italicWord ? ' ' + data.meta.italicWord : ''}`
  );
</script>

<SEO title={seoTitle} description={data.meta.summary} path="/journal/{data.slug}" type="article" />

<article class="post">
  <header class="head">
    <p class="meta">
      <a href="/journal" class="back">← writing</a>
      <span aria-hidden="true">·</span>
      <time datetime={data.meta.date}>{formatPostDate(data.meta.date)}</time>
      {#if data.meta.tags?.length}
        <span aria-hidden="true">·</span>
        <span class="tags">
          {#each data.meta.tags as tag (tag)}
            <span class="tag">{tag}</span>
          {/each}
        </span>
      {/if}
    </p>
    <h1>
      {data.meta.title}
      {#if data.meta.italicWord}
        <em>{data.meta.italicWord}</em>
      {/if}
    </h1>
    <p class="deck">{data.meta.summary}</p>
  </header>

  <div class="prose">
    <Content />
  </div>
</article>

<style>
  .post {
    display: flex;
    flex-direction: column;
    gap: var(--space-6);
    max-width: var(--container-prose);
    padding-block: var(--space-7) var(--space-8);
  }

  .head {
    display: flex;
    flex-direction: column;
    gap: var(--space-3);
  }

  .meta {
    margin: 0;
    display: flex;
    flex-wrap: wrap;
    align-items: baseline;
    gap: 8px;
    font-family: var(--font-mono);
    font-size: var(--type-tiny);
    text-transform: uppercase;
    letter-spacing: 0.14em;
    color: var(--text-tertiary);
  }
  .back {
    color: var(--text-secondary);
    text-decoration: none;
  }
  .back:hover {
    color: var(--text-primary);
  }
  .tags {
    display: inline-flex;
    gap: 6px;
  }
  .tag {
    padding: 1px 8px;
    border-radius: var(--radius-sm);
    background: var(--bg-surface-2);
    border: 0.5px solid var(--border-subtle);
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
    margin-left: 4px;
  }

  .deck {
    margin: 0;
    font-family: var(--font-serif);
    font-style: italic;
    font-size: var(--type-lead);
    line-height: var(--type-lead-lh);
    color: var(--text-secondary);
    max-width: 540px;
  }

  /* mdsvex-rendered body */
  .prose {
    color: var(--text-secondary);
    font-size: var(--type-body);
    line-height: 1.75;
  }

  .prose :global(> * + *) {
    margin-top: var(--space-4);
  }
  .prose :global(p) {
    margin: 0;
  }
  .prose :global(strong) {
    font-weight: 500;
    color: var(--text-primary);
  }
  .prose :global(em) {
    font-family: var(--font-serif);
    font-style: italic;
    color: var(--text-primary);
  }
  .prose :global(a) {
    color: var(--text-primary);
    text-decoration: underline;
    text-decoration-color: var(--accent-pink-soft);
    text-underline-offset: 3px;
  }
  .prose :global(a:hover) {
    text-decoration-color: var(--accent-pink);
  }
  .prose :global(h2) {
    font-family: var(--font-sans);
    font-weight: 500;
    font-size: var(--type-h2);
    line-height: var(--type-h2-lh);
    letter-spacing: -0.01em;
    color: var(--text-primary);
    margin-top: var(--space-6);
  }
  .prose :global(h3) {
    font-family: var(--font-sans);
    font-weight: 500;
    font-size: var(--type-h3);
    line-height: var(--type-h3-lh);
    color: var(--text-primary);
    margin-top: var(--space-5);
  }
  .prose :global(blockquote) {
    margin: 0;
    padding: var(--space-4) var(--space-5);
    border-left: 2px solid var(--accent-pink-soft);
    background: var(--bg-surface-2);
    border-radius: 0 var(--radius-md) var(--radius-md) 0;
    color: var(--text-primary);
    font-family: var(--font-serif);
    font-style: italic;
    font-size: var(--type-lead);
    line-height: var(--type-lead-lh);
  }
  .prose :global(code) {
    font-family: var(--font-mono);
    font-size: 0.92em;
    color: var(--accent-pink-soft);
    background: var(--bg-surface-2);
    padding: 1px 5px;
    border-radius: var(--radius-sm);
    border: 0.5px solid var(--border-subtle);
  }
  .prose :global(ul),
  .prose :global(ol) {
    margin: 0;
    padding-left: 1.4em;
    display: flex;
    flex-direction: column;
    gap: 6px;
  }
  .prose :global(li::marker) {
    color: var(--accent-pink-soft);
  }
</style>
