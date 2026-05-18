<!--
  Footer — full four-column footer per DESIGN.md → Footer.
  Replaces the Task 1.1 placeholder.
-->
<script lang="ts">
  import { BUILD_SHA, BUILD_TIME } from '$lib/utils/buildInfo';
  import { PROJECTS } from '$lib/content/projects';

  const elsewhereLinks = [
    { label: 'github', href: 'https://github.com/CEckelberry' },
    { label: 'linkedin', href: 'https://www.linkedin.com/in/cole-eckelberry/' }
  ];

  type ContactLink = { label: string; href: string; download?: string };
  const contactLinks: ContactLink[] = [
    { label: 'say hi →', href: '/#contact' },
    { label: 'résumé', href: '/resume' },
    {
      label: 'cv.pdf',
      href: '/Cole_Eckelberry_Senior_DevOps_Platform_Engineer.pdf',
      download: 'cole_eckelberry.pdf'
    },
    { label: 'book a call', href: 'https://cal.com/coleeckelberry' }
  ];

  const buildDate = BUILD_TIME.slice(0, 10);
</script>

<footer class="footer">
  <div class="grid">
    <section class="brand">
      <p class="name">cole eckelberry</p>
      <p class="blurb">
        platform engineer in san jose, ca. ten years of cloud + python, two years of sveltekit, six
        months of go.
      </p>
      <p class="uptime">
        <span class="dot" aria-hidden="true"></span>
        all systems operational
      </p>
    </section>

    <section class="col">
      <p class="col-head">work</p>
      <ul>
        {#each PROJECTS as p (p.slug)}
          <li>
            <a href={p.caseStudyHref}
              >{p.title}{#if p.italicWord}&nbsp;<em>{p.italicWord}</em>{/if}</a
            >
          </li>
        {/each}
      </ul>
    </section>

    <section class="col">
      <p class="col-head">elsewhere</p>
      <ul>
        {#each elsewhereLinks as link (link.href)}
          <li>
            <a href={link.href} rel={link.href.startsWith('http') ? 'me noopener' : undefined}>
              {link.label}
            </a>
          </li>
        {/each}
      </ul>
    </section>

    <section class="col">
      <p class="col-head">contact</p>
      <ul>
        {#each contactLinks as link (link.href)}
          <li>
            <a
              href={link.href}
              rel={link.href.startsWith('http') ? 'noopener' : undefined}
              download={link.download ?? undefined}
            >
              {link.label}
            </a>
          </li>
        {/each}
      </ul>
    </section>
  </div>

  <div class="bottom">
    <p class="meta">
      built with sveltekit + go · cloud run · cloud cdn ·
      <code>main@{BUILD_SHA}</code>
      <span class="dim">· {buildDate}</span>
    </p>
    <p class="copy">©2026 — made in Bay Area</p>
  </div>
</footer>

<style>
  .footer {
    /* Mirror the .page container in +layout.svelte so the brand block isn't
       jammed against the viewport edge. */
    max-width: var(--container-page);
    margin: var(--space-8) auto 0;
    padding: var(--space-8) 24px var(--space-6);
    border-top: 0.5px solid var(--border-subtle);
    color: var(--text-secondary);
    position: relative;
    z-index: 2;
  }

  @media (min-width: 768px) {
    .footer {
      padding: var(--space-8) 48px var(--space-6);
    }
  }

  .grid {
    display: grid;
    grid-template-columns: 1fr;
    gap: var(--space-6);
  }

  @media (min-width: 768px) {
    .grid {
      grid-template-columns: 1.4fr 1fr 1fr 1.2fr;
      gap: var(--space-7);
    }
  }

  .brand .name {
    font-family: var(--font-sans);
    font-weight: 500;
    font-size: var(--type-h3);
    line-height: var(--type-h3-lh);
    color: var(--text-primary);
    margin: 0 0 var(--space-3);
    letter-spacing: -0.01em;
  }

  .brand .blurb {
    font-size: var(--type-body-sm);
    line-height: 1.6;
    color: var(--text-secondary);
    margin: 0 0 var(--space-4);
    max-width: 32ch;
  }

  .uptime {
    display: inline-flex;
    align-items: center;
    gap: 8px;
    font-family: var(--font-mono);
    font-size: var(--type-tiny);
    text-transform: lowercase;
    letter-spacing: 0.06em;
    color: var(--text-tertiary);
    margin: 0;
  }
  .uptime .dot {
    width: 7px;
    height: 7px;
    border-radius: 50%;
    background: var(--accent-teal);
    box-shadow: 0 0 0 0 color-mix(in oklab, var(--accent-teal), transparent 60%);
    animation: footer-pulse 2.4s var(--ease-in-out) infinite;
  }
  @keyframes footer-pulse {
    0%,
    100% {
      box-shadow: 0 0 0 0 color-mix(in oklab, var(--accent-teal), transparent 60%);
    }
    50% {
      box-shadow: 0 0 0 6px color-mix(in oklab, var(--accent-teal), transparent 100%);
    }
  }
  @media (prefers-reduced-motion: reduce) {
    .uptime .dot {
      animation: none;
    }
  }

  .col-head {
    font-family: var(--font-mono);
    font-size: var(--type-micro);
    text-transform: uppercase;
    letter-spacing: 0.18em;
    color: var(--text-tertiary);
    margin: 0 0 var(--space-3);
  }

  .col ul {
    list-style: none;
    margin: 0;
    padding: 0;
    display: flex;
    flex-direction: column;
    gap: 6px;
  }

  .col a {
    color: var(--text-secondary);
    text-decoration: none;
    font-size: var(--type-body-sm);
    line-height: 1.5;
    transition: color var(--dur-fast) var(--ease-out);
  }
  .col a:hover {
    color: var(--text-primary);
  }
  .col a em {
    font-family: var(--font-serif);
    font-style: italic;
    color: var(--accent-pink-soft);
  }

  .bottom {
    margin-top: var(--space-7);
    padding-top: var(--space-4);
    border-top: 0.5px solid var(--border-subtle);
    display: flex;
    flex-wrap: wrap;
    gap: var(--space-3) var(--space-5);
    justify-content: space-between;
    align-items: baseline;
  }

  .meta,
  .copy {
    margin: 0;
    font-family: var(--font-mono);
    font-size: var(--type-tiny);
    color: var(--text-tertiary);
    letter-spacing: 0.04em;
  }

  .meta code {
    color: var(--accent-teal);
    background: transparent;
    padding: 0;
  }
  .meta .dim {
    color: var(--text-muted);
    margin-left: 4px;
  }
</style>
