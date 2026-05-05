<script lang="ts">
  import type { PageProps } from './$types';

  let { data }: PageProps = $props();
  const Content = $derived(data.Content);
</script>

<svelte:head>
  <title>{data.meta.title} {data.meta.italicWord ?? ''} — cole eckelberry</title>
  <meta name="description" content={data.meta.summary} />
</svelte:head>

<article class="case-study">
  <Content />
</article>

<style>
  .case-study {
    display: flex;
    flex-direction: column;
    gap: var(--space-7);
    padding-bottom: var(--space-8);
  }
</style>
