import { PROJECTS } from '$lib/content/projects';
import { SITE_URL } from '$lib/config/site';
import type { RequestHandler } from './$types';

// Generated at build time; the link list is small enough that hand-rolling
// the XML is cleaner than pulling in a sitemap library.
export const prerender = true;

interface Url {
  loc: string;
  changefreq: 'always' | 'hourly' | 'daily' | 'weekly' | 'monthly' | 'yearly' | 'never';
  priority: number;
}

const STATIC_URLS: Url[] = [
  { loc: '/', changefreq: 'weekly', priority: 1.0 },
  { loc: '/resume', changefreq: 'monthly', priority: 0.7 }
];

export const GET: RequestHandler = () => {
  const today = new Date().toISOString().slice(0, 10);

  const urls: Url[] = [
    ...STATIC_URLS,
    ...PROJECTS.map(
      (p): Url => ({
        loc: p.caseStudyHref,
        changefreq: 'monthly',
        priority: 0.8
      })
    )
  ];

  const xml = `<?xml version="1.0" encoding="UTF-8"?>
<urlset xmlns="http://www.sitemaps.org/schemas/sitemap/0.9">
${urls
  .map(
    (u) => `  <url>
    <loc>${SITE_URL}${u.loc}</loc>
    <lastmod>${today}</lastmod>
    <changefreq>${u.changefreq}</changefreq>
    <priority>${u.priority.toFixed(1)}</priority>
  </url>`
  )
  .join('\n')}
</urlset>
`;

  return new Response(xml, {
    headers: {
      'content-type': 'application/xml; charset=utf-8',
      'cache-control': 'public, max-age=3600'
    }
  });
};
