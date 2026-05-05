// See https://svelte.dev/docs/kit/types#app.d.ts
// for information about these interfaces
declare global {
  namespace App {
    // interface Error {}
    // interface Locals {}
    // interface PageData {}
    // interface PageState {}
    // interface Platform {}
  }

  /** Short git SHA of the most recent commit at build time. */
  const __BUILD_SHA__: string;
  /** ISO timestamp of when the build was produced. */
  const __BUILD_TIME__: string;
}

export {};
