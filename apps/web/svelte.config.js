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
    runes: ({ filename }) => (filename.split(/[/\\]/).includes('node_modules') ? undefined : true)
  },
  kit: {
    adapter: adapter(),
    prerender: {
      // Routes that aren't built yet but are linked from prerendered pages.
      // /journal — Phase 4 markdown blog.
      // /cv — Phase 3 PDF (the link works once the PDF is generated/dropped in static/).
      handleHttpError: ({ path, referrer, message }) => {
        const allowedMissing = new Set(['/journal', '/cv']);
        if (allowedMissing.has(path)) return;
        throw new Error(`prerender error at ${path} (linked from ${referrer}): ${message}`);
      }
    }
  }
};

export default config;
