<!--
  SEO — drops the right <head> tags for the page it's used in.
  Defaults from $lib/config/site.ts; per-page overrides via props.
-->
<script lang="ts">
  import {
    SITE_URL,
    SITE_NAME,
    DEFAULT_TITLE,
    DEFAULT_DESCRIPTION,
    DEFAULT_OG_IMAGE_PATH
  } from '$lib/config/site';

  interface Props {
    /** Page title without the site name suffix. Omit on the home page to use the full default. */
    title?: string;
    /** Meta description, ~160 chars max. */
    description?: string;
    /** Path-only canonical URL, e.g. '/' or '/work/bakeoff'. */
    path?: string;
    /** Absolute or root-relative og:image. Defaults to the site's static SVG. */
    image?: string;
    /** OpenGraph type. 'article' for case studies, 'website' otherwise. */
    type?: 'website' | 'article';
    /** When true, emit `<meta name="robots" content="noindex">` so crawlers skip the page. */
    noindex?: boolean;
  }

  let {
    title,
    description = DEFAULT_DESCRIPTION,
    path = '/',
    image = DEFAULT_OG_IMAGE_PATH,
    type = 'website',
    noindex = false
  }: Props = $props();

  const fullTitle = $derived(title ? `${title} — ${SITE_NAME}` : DEFAULT_TITLE);
  const canonical = $derived(SITE_URL + path);
  const ogImage = $derived(image.startsWith('http') ? image : SITE_URL + image);
</script>

<svelte:head>
  <title>{fullTitle}</title>
  <meta name="description" content={description} />
  <link rel="canonical" href={canonical} />
  {#if noindex}
    <meta name="robots" content="noindex" />
  {/if}

  <meta property="og:site_name" content={SITE_NAME} />
  <meta property="og:title" content={fullTitle} />
  <meta property="og:description" content={description} />
  <meta property="og:type" content={type} />
  <meta property="og:url" content={canonical} />
  <meta property="og:image" content={ogImage} />
  <meta property="og:image:width" content="1200" />
  <meta property="og:image:height" content="630" />

  <meta name="twitter:card" content="summary_large_image" />
  <meta name="twitter:title" content={fullTitle} />
  <meta name="twitter:description" content={description} />
  <meta name="twitter:image" content={ogImage} />
</svelte:head>
