// Centralized site-wide constants used by SEO + sitemap.
// When the production domain changes, update SITE_URL and the OG image
// host references will follow.

export const SITE_URL = 'https://cole-eckelberry.com';
export const SITE_NAME = 'cole eckelberry';
export const SITE_TAGLINE = 'platform engineer';

export const DEFAULT_TITLE = `${SITE_NAME} — ${SITE_TAGLINE}`;
export const DEFAULT_DESCRIPTION =
  'Senior DevOps / platform engineer in the Bay Area. Ten years of cloud and Python; SvelteKit and Go on the side. Four open-source side projects below.';

/** Used as the default og:image — a static SVG; dynamic per-page OG comes in v2 via satori. */
export const DEFAULT_OG_IMAGE_PATH = '/og.svg';
