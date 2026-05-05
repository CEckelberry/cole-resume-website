import { error } from '@sveltejs/kit';
import type { PageLoad, EntryGenerator } from './$types';
import type { CaseStudyMeta } from '$lib/content/case-studies';

// Eager glob keeps the import map static so SvelteKit can prerender each
// slug to a self-contained HTML file at build time. With only four case
// studies the bundle cost is negligible.
const MODULES = import.meta.glob('$lib/content/case-studies/*.md', {
  eager: true
}) as Record<
  string,
  {
    default: import('svelte').Component;
    metadata: CaseStudyMeta;
  }
>;

function pathFor(slug: string): string | undefined {
  return Object.keys(MODULES).find((p) => p.endsWith(`/${slug}.md`));
}

export const prerender = true;

export const entries: EntryGenerator = () =>
  Object.keys(MODULES).map((p) => ({
    slug: (p.split('/').pop() ?? '').replace(/\.md$/, '')
  }));

export const load: PageLoad = ({ params }) => {
  const path = pathFor(params.slug);
  if (!path) throw error(404, `case study not found: ${params.slug}`);
  const mod = MODULES[path];
  return {
    slug: params.slug,
    meta: mod.metadata,
    Content: mod.default
  };
};
