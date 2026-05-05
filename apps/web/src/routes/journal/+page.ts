import type { PageLoad } from './$types';
import type { JournalMeta } from '$lib/content/journal';
import { byDateDesc } from '$lib/content/journal';

// Eager glob all journal markdown — bundle is small and lets us list every
// post without an async load. Slug falls back to the filename stem.
const MODULES = import.meta.glob('$lib/content/journal/*.md', { eager: true }) as Record<
  string,
  { metadata: JournalMeta }
>;

export const prerender = true;

export const load: PageLoad = () => {
  const posts: JournalMeta[] = Object.entries(MODULES).map(([path, mod]) => {
    const filenameSlug = (path.split('/').pop() ?? '').replace(/\.md$/, '');
    return { ...mod.metadata, slug: mod.metadata.slug ?? filenameSlug };
  });
  posts.sort(byDateDesc);
  return { posts };
};
