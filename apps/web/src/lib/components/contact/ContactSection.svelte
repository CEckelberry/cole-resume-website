<!--
  ContactSection — the home page's contact area. Anchored at #contact for
  the nav and footer "contact" links.
-->
<script lang="ts">
  import ContactForm from './ContactForm.svelte';
</script>

<section id="contact" class="section">
  <header class="section-head">
    <p class="eyebrow">section 04 · contact</p>
    <h2>
      Hiring, collaborating, or just curious — <em>say hi.</em>
    </h2>
    <p class="lead">
      Best for staff/principal cloud + devops conversations. Side-project trade-ins also welcome.
    </p>
  </header>

  <div class="card">
    <ContactForm />
  </div>
</section>

<style>
  .section {
    padding: var(--space-8) 0;
    display: flex;
    flex-direction: column;
    gap: var(--space-6);
  }

  .section-head {
    display: flex;
    flex-direction: column;
    gap: var(--space-3);
    max-width: var(--container-prose);
  }

  .eyebrow {
    font-family: var(--font-mono);
    font-size: var(--type-micro);
    text-transform: uppercase;
    letter-spacing: 0.18em;
    color: var(--text-tertiary);
    margin: 0;
  }

  h2 {
    font-family: var(--font-sans);
    font-weight: 500;
    font-size: var(--type-h2);
    line-height: var(--type-h2-lh);
    letter-spacing: -0.01em;
    color: var(--text-primary);
    margin: 0;
  }
  h2 em {
    font-family: var(--font-serif);
    font-style: italic;
    font-weight: 400;
    color: var(--accent-pink-soft);
  }

  .lead {
    font-family: var(--font-serif);
    font-style: italic;
    font-size: var(--type-lead);
    line-height: var(--type-lead-lh);
    color: var(--text-secondary);
    margin: 0;
    max-width: 540px;
  }

  /* Glass-card frame around the form so the section reads as a contained
     surface rather than text floating against the mesh background. Same
     DNA as the activity card and project cards. */
  .card {
    border: 0.5px solid var(--border-default);
    border-radius: var(--radius-lg);
    background: var(--bg-surface-2);
    backdrop-filter: blur(14px) saturate(140%);
    -webkit-backdrop-filter: blur(14px) saturate(140%);
    padding: var(--space-6) var(--space-5);
    max-width: 640px;
  }

  @media (min-width: 720px) {
    .card {
      padding: var(--space-7) var(--space-6);
    }
  }
</style>
