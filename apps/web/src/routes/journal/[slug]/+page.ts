import { error } from '@sveltejs/kit';
import type { PageLoad, EntryGenerator } from './$types';
import type { JournalMeta } from '$lib/content/journal';

const MODULES = import.meta.glob('$lib/content/journal/*.md', { eager: true }) as Record<
  string,
  { default: import('svelte').Component; metadata: JournalMeta }
>;

function pathFor(slug: string): string | undefined {
  // Match either an explicit `slug` in the frontmatter or the filename stem.
  for (const [path, mod] of Object.entries(MODULES)) {
    const explicit = mod.metadata.slug;
    const filenameSlug = (path.split('/').pop() ?? '').replace(/\.md$/, '');
    if (explicit === slug || filenameSlug === slug) return path;
  }
  return undefined;
}

export const prerender = true;

export const entries: EntryGenerator = () =>
  Object.entries(MODULES).map(([path, mod]) => {
    const filenameSlug = (path.split('/').pop() ?? '').replace(/\.md$/, '');
    return { slug: mod.metadata.slug ?? filenameSlug };
  });

export const load: PageLoad = ({ params }) => {
  const path = pathFor(params.slug);
  if (!path) throw error(404, `journal post not found: ${params.slug}`);
  const mod = MODULES[path];
  return {
    slug: params.slug,
    meta: mod.metadata,
    Content: mod.default
  };
};
