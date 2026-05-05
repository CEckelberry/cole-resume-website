<!--
  Linux Lessons from Hell — preview region.
  Static terminal fragment, pink accent. No animation at v1.
-->
<script lang="ts">
  // No props.
</script>

<div class="preview">
  <pre><span class="prompt">wasteland $</span> <span class="cmd">rm -rf /</span>

<span class="teacher">// the teacher squints at you</span>
<span class="error">no. just no.</span>
<span class="teacher">// "you almost did it, kid"</span>
<span class="prompt">wasteland $</span> <span class="cursor">█</span></pre>
</div>

<style>
  .preview {
    height: 100%;
    width: 100%;
    background:
      radial-gradient(
        circle at 30% 80%,
        color-mix(in oklab, var(--accent-pink), transparent 70%) 0%,
        transparent 60%
      ),
      linear-gradient(180deg, #1a0815 0%, #0e0510 100%);
    overflow: hidden;
    display: flex;
    align-items: stretch;
    padding: 14px;
  }

  pre {
    margin: 0;
    font-family: var(--font-mono);
    font-size: 10px;
    line-height: 1.55;
    color: rgba(255, 255, 255, 0.7);
    width: 100%;
    white-space: pre-wrap;
  }

  .prompt {
    color: var(--accent-pink);
  }
  .cmd {
    color: rgba(255, 255, 255, 0.92);
  }
  .teacher {
    font-family: var(--font-serif);
    font-style: italic;
    font-size: 11px;
    color: rgba(255, 255, 255, 0.42);
  }
  .error {
    color: var(--accent-pink-soft);
    font-weight: 500;
  }
  .cursor {
    color: var(--accent-pink);
    animation: blink 1s step(2) infinite;
  }
  @keyframes blink {
    50% {
      opacity: 0;
    }
  }
  @media (prefers-reduced-motion: reduce) {
    .cursor {
      animation: none;
    }
  }
</style>
