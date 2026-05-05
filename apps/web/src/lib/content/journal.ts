// Frontmatter contract for journal posts in
// `apps/web/src/lib/content/journal/<slug>.md`.
//
// Posts are mdsvex-rendered with the same component vocabulary as the
// case studies (CSSection, CodeBlock, etc.) plus a JournalHeader for
// the title + date + tags row.

export interface JournalMeta {
  /** URL slug; if omitted, derived from the filename. */
  slug?: string;
  /** Post title — sentence case, optional italic accent appended via italicWord. */
  title: string;
  /** Italic accent word appended to the title in serif italic + accent color. */
  italicWord?: string;
  /** ISO-8601 date string, e.g. "2026-05-05". */
  date: string;
  /** One-paragraph deck shown above the post and used for OG description. */
  summary: string;
  /** Optional tags for filtering/grouping. */
  tags?: string[];
  /** Reading-time minutes (computed at build time if absent). */
  readingMinutes?: number;
}

/** Helper for sorting posts newest-first. */
export function byDateDesc(a: JournalMeta, b: JournalMeta): number {
  return b.date.localeCompare(a.date);
}

export function formatPostDate(iso: string): string {
  const d = new Date(iso + 'T00:00:00Z');
  if (Number.isNaN(d.getTime())) return iso;
  return d.toLocaleDateString('en-US', {
    timeZone: 'UTC',
    year: 'numeric',
    month: 'short',
    day: 'numeric'
  });
}
