import type { ProjectSlug } from './projects';

/**
 * Frontmatter contract for case-study markdown files in
 * `packages/content/case-studies/<slug>.md`. Mirrored at runtime via the
 * +page.ts loader.
 */
export interface CaseStudyMeta {
  slug: ProjectSlug;
  title: string;
  /** Optional italic word inside the title, project accent color. */
  italicWord?: string;
  /** One-paragraph summary; renders as the case-study deck. */
  summary: string;
  /** Role on the project. */
  role: string;
  /** Stack array shown in the meta row. */
  stack: string[];
  /** Ship date (or "in progress" if building). */
  shipped: string;
  /** Status pill text. */
  status: 'live' | 'building' | 'planned';
  /** Live demo URL (mirrors PROJECTS, used in the demo browser bar). */
  liveDemoUrl: string;
}
