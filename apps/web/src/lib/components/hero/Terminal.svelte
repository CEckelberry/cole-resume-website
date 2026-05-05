<!--
  Terminal — boots on first paint, types out the script, then idles with a
  blinking cursor. Spec from DESIGN.md → Hero terminal.

  Boot sequence (~1.6s total at 30ms/char):

      cole@bay-area $ whoami
      senior platform engineer
      cole@bay-area $ cat now.txt
      role:    "open to staff/principal cloud + devops"
      stack:   ["python", "sveltekit", "go"]
      city:    "san jose, ca"
      open:    true
      cole@bay-area $ █

  SSR renders the final state (full script visible) so visitors without JS
  still see the content. On client mount we reset chars=0 and animate up,
  unless prefers-reduced-motion is set in which case we leave the SSR state
  alone and start the cursor blinking immediately.

  The animation uses requestAnimationFrame (not setInterval) so it pauses
  cleanly when the tab is backgrounded.
-->
<script lang="ts">
  import { onMount } from 'svelte';

  type SegClass = 'prompt' | 'cmd' | 'output' | 'key' | 'val';

  type Segment = { text: string; cls: SegClass };

  // Each segment is a continuous run of text in a single style. Newlines are
  // baked into the segment text so the rendered span flow handles wrapping.
  const SCRIPT: Segment[] = [
    { text: 'cole@bay-area $ ', cls: 'prompt' },
    { text: 'whoami\n', cls: 'cmd' },
    { text: 'senior platform engineer\n', cls: 'output' },
    { text: 'cole@bay-area $ ', cls: 'prompt' },
    { text: 'cat now.txt\n', cls: 'cmd' },
    { text: 'role:    ', cls: 'key' },
    { text: '"open to staff/principal cloud + devops"\n', cls: 'val' },
    { text: 'stack:   ', cls: 'key' },
    { text: '["python", "sveltekit", "go"]\n', cls: 'val' },
    { text: 'city:    ', cls: 'key' },
    { text: '"san jose, ca"\n', cls: 'val' },
    { text: 'open:    ', cls: 'key' },
    { text: 'true\n', cls: 'val' },
    { text: 'cole@bay-area $ ', cls: 'prompt' }
  ];

  const TOTAL = SCRIPT.reduce((n, s) => n + s.text.length, 0);
  const MS_PER_CHAR = 30;

  // SSR: render full script. Client effect resets to 0 then animates up, unless
  // reduced motion is requested.
  let typed = $state(TOTAL);
  let booted = $state(true);

  const visible = $derived.by(() => {
    if (typed >= TOTAL) return SCRIPT;
    const out: Segment[] = [];
    let remaining = typed;
    for (const seg of SCRIPT) {
      if (remaining <= 0) break;
      if (remaining >= seg.text.length) {
        out.push(seg);
        remaining -= seg.text.length;
      } else {
        out.push({ text: seg.text.slice(0, remaining), cls: seg.cls });
        remaining = 0;
      }
    }
    return out;
  });

  onMount(() => {
    const reduced = window.matchMedia('(prefers-reduced-motion: reduce)').matches;
    if (reduced) {
      typed = TOTAL;
      booted = true;
      return;
    }

    // Reset to typing-from-scratch and animate.
    typed = 0;
    booted = false;
    let raf = 0;
    let start = 0;

    const tick = (ts: number) => {
      if (!start) start = ts;
      const elapsed = ts - start;
      typed = Math.min(TOTAL, Math.floor(elapsed / MS_PER_CHAR));
      if (typed < TOTAL) {
        raf = requestAnimationFrame(tick);
      } else {
        booted = true;
      }
    };
    raf = requestAnimationFrame(tick);
    return () => cancelAnimationFrame(raf);
  });
</script>

<div
  class="terminal"
  role="img"
  aria-label="Terminal showing whoami output: senior platform engineer, role: open to staff or principal cloud or devops, stack: python sveltekit go, city: san jose ca, open: true."
>
  <header class="bar" aria-hidden="true">
    <span class="dot d-r"></span>
    <span class="dot d-y"></span>
    <span class="dot d-g"></span>
    <span class="title">~/cole — zsh</span>
  </header>

  <pre class="body" aria-hidden="true">{#each visible as seg, i (i)}<span class={seg.cls}
        >{seg.text}</span
      >{/each}<span class="cursor" class:blinking={booted}>█</span></pre>
</div>

<style>
  .terminal {
    width: 100%;
    max-width: 320px;
    border: 0.5px solid color-mix(in oklab, var(--accent-pink-soft), transparent 70%);
    border-radius: var(--radius-xl);
    background: rgba(8, 4, 18, 0.7);
    backdrop-filter: blur(14px);
    -webkit-backdrop-filter: blur(14px);
    overflow: hidden;
    box-shadow:
      0 12px 40px -16px rgba(0, 0, 0, 0.45),
      0 0 0 0.5px rgba(255, 255, 255, 0.03) inset;
  }

  .bar {
    display: flex;
    align-items: center;
    gap: 6px;
    padding: 8px 12px;
    border-bottom: 0.5px solid rgba(255, 255, 255, 0.08);
    background: rgba(255, 255, 255, 0.02);
  }
  .dot {
    width: 9px;
    height: 9px;
    border-radius: 50%;
    box-shadow: 0 0 0 0.5px rgba(0, 0, 0, 0.25) inset;
  }
  .d-r {
    background: #ff5f57;
  }
  .d-y {
    background: #febc2e;
  }
  .d-g {
    background: #28c840;
  }
  .title {
    margin-left: 8px;
    font-family: var(--font-mono);
    font-size: var(--type-tiny);
    color: rgba(255, 255, 255, 0.55);
    letter-spacing: 0.04em;
  }

  .body {
    margin: 0;
    padding: 14px 14px 16px;
    font-family: var(--font-mono);
    font-size: var(--type-micro);
    line-height: 1.7;
    color: rgba(255, 255, 255, 0.85);
    white-space: pre-wrap;
    word-break: break-word;
    min-height: 220px;
  }

  .prompt {
    color: var(--accent-teal);
  }
  .cmd {
    color: rgba(255, 255, 255, 0.95);
  }
  .output {
    color: rgba(255, 255, 255, 0.7);
  }
  .key {
    color: var(--accent-pink-soft);
  }
  .val {
    color: var(--accent-amber-soft);
  }

  .cursor {
    color: var(--accent-teal);
  }
  .cursor.blinking {
    animation: blink 1s step(2) infinite;
  }

  @keyframes blink {
    50% {
      opacity: 0;
    }
  }

  @media (prefers-reduced-motion: reduce) {
    .cursor.blinking {
      animation: none;
    }
  }

  /* Light mode: keep the terminal's chrome dark — terminals look weird with a
     pale background, and the dark surface gives the colored text its punch. */
</style>
