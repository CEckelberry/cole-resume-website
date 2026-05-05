import adapter from '@sveltejs/adapter-node';
import { mdsvex } from 'mdsvex';

/** @type {import('@sveltejs/kit').Config} */
const config = {
  // mdsvex compiles .md to Svelte components. We keep .svelte for normal
  // components and reserve .md for case-study content.
  extensions: ['.svelte', '.md'],
  preprocess: [
    mdsvex({
      extensions: ['.md'],
      smartypants: { quotes: true, ellipses: true, dashes: 'oldschool' }
    })
  ],
  compilerOptions: {
    // Force runes mode for the project, except for libraries. Can be removed in svelte 6.
    runes: ({ filename }) => (filename.split(/[/\\]/).includes('node_modules') ? undefined : true),
    // mdsvex 0.12 emits `<script context="module">` which Svelte 5 deprecated
    // in favor of `<script module>`. Until mdsvex updates, silence the noise
    // — the code still works.
    warningFilter: (w) => w.code !== 'script_context_deprecated'
  },
  kit: {
    adapter: adapter(),
    prerender: {
      // /cv — Phase 3 PDF (the link works once the PDF is generated/dropped in static/).
      handleHttpError: ({ path, referrer, message }) => {
        const allowedMissing = new Set(['/cv']);
        if (allowedMissing.has(path)) return;
        throw new Error(`prerender error at ${path} (linked from ${referrer}): ${message}`);
      },
      // /journal/[slug] is prerender=true but has no posts at v1, so its
      // entries() returns []. SvelteKit 2.59 requires explicit handling.
      handleUnseenRoutes: 'ignore'
    }
  }
};

export default config;
