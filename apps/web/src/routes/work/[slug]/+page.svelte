<script lang="ts">
  import type { PageProps } from './$types';
  import SEO from '$lib/components/shell/SEO.svelte';

  let { data }: PageProps = $props();
  const Content = $derived(data.Content);

  const seoTitle = $derived(
    `${data.meta.title}${data.meta.italicWord ? ' ' + data.meta.italicWord : ''}`
  );
</script>

<SEO title={seoTitle} description={data.meta.summary} path="/work/{data.slug}" type="article" />

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
