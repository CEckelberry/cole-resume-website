// Project metadata — single source of truth for the gallery and the
// case-study routes. Keep this typed and aligned with
// `packages/content/case-studies/<slug>.md`.

export type ProjectSlug = 'linux' | 'bakeoff' | 'diamond' | 'terraplane';
export type ProjectAccent = 'pink' | 'teal' | 'purple' | 'amber';
export type ProjectStatus = 'live' | 'building' | 'planned';
export type ProjectFilter = 'games' | 'tools' | 'data' | 'visualizers';

export interface Project {
  slug: ProjectSlug;
  /** Display index, "01"–"04". */
  number: string;
  /** Eyebrow kind label, e.g. "A GAME", "A TOOL". */
  kind: string;
  /** Filter chip this project belongs to. */
  filter: ProjectFilter;
  /** Project title without italic accent. */
  title: string;
  /**
   * Optional italic accent shown after the title in serif italic + project
   * accent color. Renders as: `{title} {italicWord}` with the word italicized.
   */
  italicWord?: string;
  /** One-sentence description for the card body. */
  description: string;
  /** Stack tags rendered below the description. */
  tags: string[];
  accent: ProjectAccent;
  status: ProjectStatus;
  /** Route to the case study page. */
  caseStudyHref: string;
  /** Public URL where the deployed demo lives (used in the case-study browser bar). */
  liveDemoUrl: string;
  /** Optional GitHub URL for the project repo. */
  githubUrl?: string;
}

export const PROJECTS: readonly Project[] = [
  {
    slug: 'linux',
    number: '01',
    kind: 'A GAME',
    filter: 'games',
    title: 'Linux Lessons from',
    italicWord: 'Hell',
    description:
      'A browser-based bash tutor with a foul-mouthed wasteland teacher. Type real shell commands, get real feedback, get yelled at when you forget the dash before -rf.',
    tags: ['SvelteKit', 'xterm.js', 'Go WASM', 'Tailwind'],
    accent: 'pink',
    status: 'planned',
    caseStudyHref: '/work/linux',
    liveDemoUrl: 'https://linux.cole-eckelberry.com',
    githubUrl: 'https://github.com/CEckelberry/linux-lessons-from-hell'
  },
  {
    slug: 'bakeoff',
    number: '02',
    kind: 'A TOOL',
    filter: 'tools',
    title: 'Backend',
    italicWord: 'Bake-off',
    description:
      'Six runtimes (Go, Rust, Bun, Node, Python, PHP) implement the same checkout endpoint. Switch which one serves your request via header. Watch the latency move in real time.',
    tags: ['Go', 'Rust', 'GKE', 'Postgres'],
    accent: 'teal',
    status: 'live',
    caseStudyHref: '/work/bakeoff',
    liveDemoUrl: 'https://backend-bakeoff.com',
    githubUrl: 'https://github.com/CEckelberry/backend-bakeoff'
  },
  {
    slug: 'diamond',
    number: '03',
    kind: 'A DATA TOY',
    filter: 'data',
    title: 'Diamond',
    italicWord: 'Departures',
    description:
      'Top 100 active MLB players ranked by sabermetrics, styled as a Penn Station split-flap board, updated live during games. Watch the rankings flip when someone barrels a fastball.',
    tags: ['Go SSE', 'FanGraphs', 'SvelteKit', 'Postgres'],
    accent: 'purple',
    status: 'planned',
    caseStudyHref: '/work/diamond',
    liveDemoUrl: 'https://diamond.cole-eckelberry.com',
    githubUrl: 'https://github.com/CEckelberry/diamond-departures'
  },
  {
    slug: 'terraplane',
    number: '04',
    kind: 'A VISUALIZER',
    filter: 'visualizers',
    title: 'Terraplane',
    description:
      'Paste a `terraform plan`, get a walkable architecture graph with cost estimates and blast-radius highlights. Built so I never have to read raw plan output again.',
    tags: ['Go', 'HCL parser', 'D3', 'Cloud Run'],
    accent: 'amber',
    status: 'planned',
    caseStudyHref: '/work/terraplane',
    liveDemoUrl: 'https://terraplane.cole-eckelberry.com',
    githubUrl: 'https://github.com/CEckelberry/terraplane'
  }
] as const;

export const PROJECT_BY_SLUG: Readonly<Record<ProjectSlug, Project>> = Object.fromEntries(
  PROJECTS.map((p) => [p.slug, p])
) as Record<ProjectSlug, Project>;

export const FILTERS: ReadonlyArray<{
  id: 'all' | ProjectFilter;
  label: string;
}> = [
  { id: 'all', label: 'all' },
  { id: 'games', label: 'games' },
  { id: 'tools', label: 'tools' },
  { id: 'data', label: 'data' },
  { id: 'visualizers', label: 'visualizers' }
];

export const ACCENT_COLOR: Readonly<Record<ProjectAccent, string>> = {
  pink: 'var(--accent-pink)',
  teal: 'var(--accent-teal)',
  purple: 'var(--accent-purple)',
  amber: 'var(--accent-amber)'
};

export const ACCENT_SOFT: Readonly<Record<ProjectAccent, string>> = {
  pink: 'var(--accent-pink-soft)',
  teal: 'var(--accent-teal-soft)',
  purple: 'var(--accent-purple-soft)',
  amber: 'var(--accent-amber-soft)'
};

export const ACCENT_EMPHASIS: Readonly<Record<ProjectAccent, string>> = {
  pink: 'var(--accent-pink-emphasis)',
  teal: 'var(--accent-teal-emphasis)',
  purple: 'var(--accent-purple-emphasis)',
  amber: 'var(--accent-amber-emphasis)'
};
