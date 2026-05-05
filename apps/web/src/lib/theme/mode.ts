// Color-mode helpers — keep the localStorage key and DOM attribute in lockstep.
//
// First-paint resolution lives in `src/app.html` so the document loads with
// the correct data-mode before CSS resolves; this module only handles
// runtime toggling once the app is interactive.

export type Mode = 'light' | 'dark';

const STORAGE_KEY = 'mode';
const ATTR = 'data-mode';

function isMode(v: unknown): v is Mode {
  return v === 'light' || v === 'dark';
}

/** Read the current mode. Returns 'dark' during SSR or before hydration. */
export function getMode(): Mode {
  if (typeof document === 'undefined') return 'dark';
  const attr = document.documentElement.getAttribute(ATTR);
  return isMode(attr) ? attr : 'dark';
}

/** Persist a mode to localStorage and reflect it on the document. */
export function setMode(next: Mode): void {
  if (typeof document === 'undefined') return;
  document.documentElement.setAttribute(ATTR, next);
  document.documentElement.style.colorScheme = next;
  try {
    localStorage.setItem(STORAGE_KEY, next);
  } catch {
    // Private mode / disabled storage: the change is still live for this
    // session, just not remembered. That's the correct degraded behavior.
  }
}

/** Flip mode and return the new value. */
export function toggleMode(): Mode {
  const next: Mode = getMode() === 'dark' ? 'light' : 'dark';
  setMode(next);
  return next;
}
