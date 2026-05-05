// SvelteKit /api/contact endpoint — thin proxy to the Go service.
//
// Why proxy through SvelteKit instead of having the browser hit the Go API
// directly:
//   1. Single-origin: avoids CORS preflight on the form submit.
//   2. Hides the API URL: the Go service can move (Cloud Run swap, regional
//      failover) without changing the frontend.
//   3. Lets us forward the visitor's real IP via X-Forwarded-For and
//      User-Agent so the rate-limiter has correct keys.

import { error } from '@sveltejs/kit';
import type { RequestHandler } from './$types';

const API_URL = (typeof process !== 'undefined' && process.env.API_URL) || 'http://localhost:8080';

export const POST: RequestHandler = async ({ request, getClientAddress }) => {
  const body = await request.text();

  let upstream: Response;
  try {
    upstream = await fetch(`${API_URL}/api/contact`, {
      method: 'POST',
      headers: {
        'content-type': 'application/json',
        'x-forwarded-for': getClientAddress(),
        'user-agent': request.headers.get('user-agent') ?? ''
      },
      body
    });
  } catch (err) {
    const message = err instanceof Error ? err.message : 'unknown';
    throw error(502, `contact api unreachable: ${message}`);
  }

  // Pass through the upstream response verbatim (status + body).
  const text = await upstream.text();
  const contentType = upstream.headers.get('content-type') ?? 'application/json';
  return new Response(text, { status: upstream.status, headers: { 'content-type': contentType } });
};

// This endpoint is request-time only; never include in static prerender.
export const prerender = false;
