import type { RequestHandler } from './$types';
import { SITE_URL, SITE_NAME, DEFAULT_DESCRIPTION } from '$lib/config/site';
import { byDateDesc, type JournalMeta } from '$lib/content/journal';

const MODULES = import.meta.glob('$lib/content/journal/*.md', { eager: true }) as Record<
  string,
  { metadata: JournalMeta }
>;

export const prerender = true;

function escapeXml(s: string): string {
  return s
    .replaceAll('&', '&amp;')
    .replaceAll('<', '&lt;')
    .replaceAll('>', '&gt;')
    .replaceAll('"', '&quot;')
    .replaceAll("'", '&apos;');
}

export const GET: RequestHandler = () => {
  const posts: JournalMeta[] = Object.entries(MODULES).map(([path, mod]) => {
    const filenameSlug = (path.split('/').pop() ?? '').replace(/\.md$/, '');
    return { ...mod.metadata, slug: mod.metadata.slug ?? filenameSlug };
  });
  posts.sort(byDateDesc);

  const lastBuild = posts[0]?.date ?? new Date().toISOString().slice(0, 10);

  const feedUrl = `${SITE_URL}/journal/rss.xml`;
  const channelLink = `${SITE_URL}/journal`;

  const items = posts
    .map((p) => {
      const url = `${SITE_URL}/journal/${p.slug}`;
      const pubDate = new Date(p.date + 'T00:00:00Z').toUTCString();
      const title = `${p.title}${p.italicWord ? ' ' + p.italicWord : ''}`;
      return `    <item>
      <title>${escapeXml(title)}</title>
      <link>${url}</link>
      <guid isPermaLink="true">${url}</guid>
      <pubDate>${pubDate}</pubDate>
      <description>${escapeXml(p.summary)}</description>
    </item>`;
    })
    .join('\n');

  const xml = `<?xml version="1.0" encoding="UTF-8"?>
<rss version="2.0" xmlns:atom="http://www.w3.org/2005/Atom">
  <channel>
    <title>${escapeXml(SITE_NAME)} — writing</title>
    <link>${channelLink}</link>
    <atom:link href="${feedUrl}" rel="self" type="application/rss+xml"/>
    <description>${escapeXml(DEFAULT_DESCRIPTION)}</description>
    <language>en-us</language>
    <lastBuildDate>${new Date(lastBuild + 'T00:00:00Z').toUTCString()}</lastBuildDate>
${items}
  </channel>
</rss>
`;

  return new Response(xml, {
    headers: {
      'content-type': 'application/rss+xml; charset=utf-8',
      'cache-control': 'public, max-age=3600'
    }
  });
};
