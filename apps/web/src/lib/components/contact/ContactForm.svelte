<!--
  ContactForm — name + email + message + submit.
  Posts to /api/contact (SvelteKit endpoint, proxies to the Go service).
  Inline validation, success/error states, honeypot for low-effort spam.
-->
<script lang="ts">
  import Button from '$lib/components/ui/Button.svelte';

  type Status = 'idle' | 'submitting' | 'success' | 'error';

  let name = $state('');
  let email = $state('');
  let message = $state('');
  let honeypot = $state(''); // bots fill this; humans don't see it
  let status = $state<Status>('idle');
  let errors = $state<{ name?: string; email?: string; message?: string; form?: string }>({});

  // Server-side validation is the source of truth; this exists for UX.
  function validate(): boolean {
    const next: typeof errors = {};
    if (!name.trim()) next.name = 'name is required.';
    else if (name.length > 120) next.name = 'name is too long (max 120).';
    if (!email.trim()) next.email = 'email is required.';
    else if (!/^[^\s@]+@[^\s@]+\.[^\s@]+$/.test(email)) next.email = 'that email looks off.';
    if (!message.trim()) next.message = 'message is required.';
    else if (message.length > 4000) next.message = 'message is too long (max 4000).';
    errors = next;
    return Object.keys(next).length === 0;
  }

  const FORM_ACTION =
    'https://docs.google.com/forms/d/e/1FAIpQLSfGn3v8lSYvStCxbJ1mCnkAUtmP0AGpMAEkhNzFOdVGJ5d3WA/formResponse';

  async function submit(event: SubmitEvent) {
    event.preventDefault();
    if (!validate()) return;
    if (honeypot.trim()) {
      status = 'success';
      return;
    }

    status = 'submitting';
    errors = {};
    try {
      const body = new URLSearchParams({
        'entry.1334858943': name,
        'entry.701278000': email,
        'entry.51477920': message,
        fvv: '1',
        pageHistory: '0',
        fbzx: String(Math.floor(Math.random() * 9e15))
      });

      // Google Forms doesn't allow CORS reads — no-cors means we can't inspect
      // the response, so we optimistically treat every non-thrown fetch as success.
      await fetch(FORM_ACTION, { method: 'POST', mode: 'no-cors', body });

      status = 'success';
      name = '';
      email = '';
      message = '';
    } catch {
      errors = { form: 'network error. check your connection and try again.' };
      status = 'error';
    }
  }
</script>

<form class="contact" onsubmit={submit} novalidate>
  {#if status === 'success'}
    <div class="success" role="status" aria-live="polite">
      <p class="success-title">message sent.</p>
      <p class="success-body">i'll reply within a couple days. promise.</p>
    </div>
  {:else}
    <div class="row">
      <label for="contact-name">name</label>
      <input
        id="contact-name"
        type="text"
        autocomplete="name"
        maxlength="120"
        bind:value={name}
        aria-invalid={!!errors.name}
        aria-describedby={errors.name ? 'contact-name-err' : undefined}
        disabled={status === 'submitting'}
      />
      {#if errors.name}
        <p id="contact-name-err" class="error">{errors.name}</p>
      {/if}
    </div>

    <div class="row">
      <label for="contact-email">email</label>
      <input
        id="contact-email"
        type="email"
        autocomplete="email"
        maxlength="254"
        bind:value={email}
        aria-invalid={!!errors.email}
        aria-describedby={errors.email ? 'contact-email-err' : undefined}
        disabled={status === 'submitting'}
      />
      {#if errors.email}
        <p id="contact-email-err" class="error">{errors.email}</p>
      {/if}
    </div>

    <div class="row">
      <label for="contact-message">message</label>
      <textarea
        id="contact-message"
        rows="5"
        maxlength="4000"
        bind:value={message}
        aria-invalid={!!errors.message}
        aria-describedby={errors.message ? 'contact-message-err' : undefined}
        disabled={status === 'submitting'}
      ></textarea>
      {#if errors.message}
        <p id="contact-message-err" class="error">{errors.message}</p>
      {/if}
    </div>

    <!-- Honeypot: visually hidden + aria-hidden + tabindex=-1 so humans skip it. -->
    <div class="honeypot" aria-hidden="true">
      <label for="contact-company">company (leave blank)</label>
      <input
        id="contact-company"
        type="text"
        tabindex="-1"
        autocomplete="off"
        bind:value={honeypot}
      />
    </div>

    {#if errors.form}
      <p class="error form-error" role="alert">{errors.form}</p>
    {/if}

    <div class="actions">
      <Button type="submit" variant="primary" size="lg" disabled={status === 'submitting'}>
        {status === 'submitting' ? 'sending…' : 'send'}
      </Button>
      <p class="hint">replies usually within a couple days.</p>
    </div>
  {/if}
</form>

<style>
  .contact {
    display: flex;
    flex-direction: column;
    gap: var(--space-4);
    max-width: 560px;
  }

  .row {
    display: flex;
    flex-direction: column;
    gap: 6px;
  }

  label {
    font-family: var(--font-mono);
    font-size: var(--type-tiny);
    text-transform: uppercase;
    letter-spacing: 0.16em;
    color: var(--text-tertiary);
  }

  input,
  textarea {
    font-family: var(--font-sans);
    font-size: var(--type-body);
    line-height: 1.5;
    color: var(--text-primary);
    background: var(--bg-surface-2);
    border: 0.5px solid var(--border-default);
    border-radius: var(--radius-md);
    padding: 10px 12px;
    backdrop-filter: blur(10px);
    -webkit-backdrop-filter: blur(10px);
    transition:
      border-color var(--dur-fast) var(--ease-out),
      background-color var(--dur-fast) var(--ease-out);
  }
  input:focus,
  textarea:focus {
    outline: none;
    border-color: var(--accent-pink);
    background: var(--bg-surface-3);
  }
  input:disabled,
  textarea:disabled {
    opacity: 0.6;
    cursor: not-allowed;
  }
  input[aria-invalid='true'],
  textarea[aria-invalid='true'] {
    border-color: var(--accent-coral);
  }
  textarea {
    resize: vertical;
    min-height: 120px;
  }

  .error {
    margin: 0;
    font-family: var(--font-mono);
    font-size: var(--type-tiny);
    color: var(--accent-coral);
    letter-spacing: 0.02em;
  }
  .form-error {
    padding: 8px 12px;
    background: color-mix(in oklab, var(--accent-coral), transparent 88%);
    border: 0.5px solid color-mix(in oklab, var(--accent-coral), transparent 60%);
    border-radius: var(--radius-md);
  }

  .actions {
    display: flex;
    align-items: center;
    flex-wrap: wrap;
    gap: var(--space-3) var(--space-4);
    margin-top: var(--space-2);
  }
  .hint {
    margin: 0;
    font-family: var(--font-mono);
    font-size: var(--type-tiny);
    color: var(--text-tertiary);
    letter-spacing: 0.04em;
  }

  /* Off-screen but focusable would let a screen reader hit the honeypot.
     This stays fully removed from a11y tree and tab order. */
  .honeypot {
    position: absolute;
    left: -10000px;
    width: 1px;
    height: 1px;
    overflow: hidden;
  }

  .success {
    padding: var(--space-5);
    border: 0.5px solid color-mix(in oklab, var(--accent-teal), transparent 50%);
    border-radius: var(--radius-md);
    background: color-mix(in oklab, var(--accent-teal), transparent 88%);
  }
  .success-title {
    margin: 0 0 4px;
    font-family: var(--font-sans);
    font-weight: 500;
    color: var(--text-primary);
  }
  .success-body {
    margin: 0;
    font-family: var(--font-serif);
    font-style: italic;
    color: var(--text-secondary);
  }
</style>
