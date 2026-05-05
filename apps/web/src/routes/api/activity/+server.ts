// SvelteKit /api/activity endpoint — proxies to the Go service.
//
// Same proxy pattern as /api/contact: same-origin, hides the API host,
// forwards User-Agent. The Go side caches GitHub responses for 5 min so
// re-hitting this proxy is cheap.

import { error } from '@sveltejs/kit';
import type { RequestHandler } from './$types';

const API_URL = (typeof process !== 'undefined' && process.env.API_URL) || 'http://localhost:8080';

export const GET: RequestHandler = async ({ request, fetch: kitFetch }) => {
  let upstream: Response;
  try {
    upstream = await kitFetch(`${API_URL}/api/activity`, {
      headers: {
        accept: 'application/json',
        'user-agent': request.headers.get('user-agent') ?? 'cole-eckelberry-portfolio'
      }
    });
  } catch (err) {
    const message = err instanceof Error ? err.message : 'unknown';
    throw error(502, `activity api unreachable: ${message}`);
  }

  const text = await upstream.text();
  return new Response(text, {
    status: upstream.status,
    headers: {
      'content-type': upstream.headers.get('content-type') ?? 'application/json',
      // Edge cache for a minute; the upstream's own 5-min cache is the
      // source of truth.
      'cache-control': 'public, max-age=60'
    }
  });
};

export const prerender = false;
