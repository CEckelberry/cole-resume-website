// Build metadata helpers — read the constants injected by vite.config.ts at
// build time and format them for display in the nav, footer, and tooling.

export const BUILD_SHA: string = typeof __BUILD_SHA__ !== 'undefined' ? __BUILD_SHA__ : 'dev';

export const BUILD_TIME: string =
  typeof __BUILD_TIME__ !== 'undefined' ? __BUILD_TIME__ : new Date().toISOString();

const MINUTE = 60_000;
const HOUR = 60 * MINUTE;
const DAY = 24 * HOUR;

/**
 * Format a timestamp as "N{unit} ago", coarse-grained so the value doesn't
 * thrash for visitors who linger. We deliberately stop at days; older builds
 * collapse to "Nd ago" rather than weeks/months — the deploy timestamp is
 * meant to feel fresh, and a stale build is a problem to surface, not hide.
 */
export function timeAgo(iso: string, now: number = Date.now()): string {
  const t = Date.parse(iso);
  if (Number.isNaN(t)) return 'just now';
  const delta = Math.max(0, now - t);

  if (delta < MINUTE) return 'just now';
  if (delta < HOUR) return `${Math.floor(delta / MINUTE)}m ago`;
  if (delta < DAY) return `${Math.floor(delta / HOUR)}h ago`;
  return `${Math.floor(delta / DAY)}d ago`;
}
