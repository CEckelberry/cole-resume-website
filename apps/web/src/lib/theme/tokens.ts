// ============================================================
// Design tokens — type-safe references to the CSS custom
// properties defined in `src/app.css`. Use the helpers below
// when a value needs to be read from JS/TS (e.g. for inline
// style binding, motion calculations, canvas drawing). Prefer
// utility classes / CSS where possible.
// ============================================================

export const ACCENT = {
  pink: 'var(--accent-pink)',
  pinkSoft: 'var(--accent-pink-soft)',
  purple: 'var(--accent-purple)',
  purpleSoft: 'var(--accent-purple-soft)',
  teal: 'var(--accent-teal)',
  tealSoft: 'var(--accent-teal-soft)',
  amber: 'var(--accent-amber)',
  amberSoft: 'var(--accent-amber-soft)',
  coral: 'var(--accent-coral)'
} as const;

export type AccentName = keyof typeof ACCENT;

/** Project → home accent color, per DESIGN.md. */
export const PROJECT_ACCENT: Record<'linux' | 'bakeoff' | 'diamond' | 'terraplane', AccentName> = {
  linux: 'pink',
  bakeoff: 'teal',
  diamond: 'purple',
  terraplane: 'amber'
};

export const BG = {
  canvas: 'var(--bg-canvas)',
  surface: 'var(--bg-surface)',
  surface2: 'var(--bg-surface-2)',
  surface3: 'var(--bg-surface-3)'
} as const;

export const TEXT = {
  primary: 'var(--text-primary)',
  secondary: 'var(--text-secondary)',
  tertiary: 'var(--text-tertiary)',
  muted: 'var(--text-muted)'
} as const;

export const BORDER = {
  subtle: 'var(--border-subtle)',
  default: 'var(--border-default)',
  strong: 'var(--border-strong)'
} as const;

export const TYPE = {
  display: 'var(--type-display)',
  h1: 'var(--type-h1)',
  h2: 'var(--type-h2)',
  h3: 'var(--type-h3)',
  lead: 'var(--type-lead)',
  body: 'var(--type-body)',
  bodySm: 'var(--type-body-sm)',
  meta: 'var(--type-meta)',
  micro: 'var(--type-micro)',
  tiny: 'var(--type-tiny)',
  nano: 'var(--type-nano)'
} as const;

export const SPACE = {
  1: 'var(--space-1)',
  2: 'var(--space-2)',
  3: 'var(--space-3)',
  4: 'var(--space-4)',
  5: 'var(--space-5)',
  6: 'var(--space-6)',
  7: 'var(--space-7)',
  8: 'var(--space-8)'
} as const;

export const RADIUS = {
  sm: 'var(--radius-sm)',
  md: 'var(--radius-md)',
  lg: 'var(--radius-lg)',
  xl: 'var(--radius-xl)',
  pill: 'var(--radius-pill)'
} as const;

export const MOTION = {
  easeOut: 'var(--ease-out)',
  easeInOut: 'var(--ease-in-out)',
  easeSpring: 'var(--ease-spring)',
  durFast: 'var(--dur-fast)',
  durDefault: 'var(--dur-default)',
  durSlow: 'var(--dur-slow)'
} as const;

export const FONT = {
  sans: 'var(--font-sans)',
  serif: 'var(--font-serif)',
  mono: 'var(--font-mono)'
} as const;

/** Read a CSS custom property from the document root, in pixels (or string). */
export function readToken(name: string, fallback = ''): string {
  if (typeof window === 'undefined') return fallback;
  return getComputedStyle(document.documentElement).getPropertyValue(name).trim() || fallback;
}
