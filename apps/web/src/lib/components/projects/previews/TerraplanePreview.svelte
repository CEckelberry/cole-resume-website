<!--
  Terraplane — preview region.
  Inline SVG architecture graph: VPC at top branching to GKE + Cloud SQL,
  with leaf nodes below. Dashed amber connectors. Cost label "$1,247/mo".
-->
<script lang="ts">
  // No props.
</script>

<div class="preview">
  <div class="cost">
    <span class="cost-label">monthly cost</span>
    <span class="cost-value">$1,247<span class="per">/mo</span></span>
  </div>

  <svg
    class="graph"
    viewBox="0 0 280 130"
    fill="none"
    aria-hidden="true"
    preserveAspectRatio="xMidYMid meet"
  >
    <!-- Connectors -->
    <g class="links" stroke="var(--accent-amber)" stroke-width="0.8" stroke-dasharray="3 3">
      <path d="M140 24 L80 56" />
      <path d="M140 24 L200 56" />
      <path d="M80 76 L40 108" />
      <path d="M80 76 L100 108" />
      <path d="M200 76 L180 108" />
      <path d="M200 76 L240 108" />
    </g>

    <!-- VPC top -->
    <g class="node node-top">
      <rect x="108" y="10" width="64" height="22" rx="4" />
      <text x="140" y="24">vpc</text>
    </g>

    <!-- Mid -->
    <g class="node">
      <rect x="48" y="56" width="64" height="22" rx="4" />
      <text x="80" y="70">gke</text>
    </g>
    <g class="node">
      <rect x="168" y="56" width="64" height="22" rx="4" />
      <text x="200" y="70">cloud sql</text>
    </g>

    <!-- Leaves -->
    <g class="leaf">
      <rect x="14" y="108" width="52" height="18" rx="3" />
      <text x="40" y="120">runner</text>
    </g>
    <g class="leaf">
      <rect x="74" y="108" width="52" height="18" rx="3" />
      <text x="100" y="120">job</text>
    </g>
    <g class="leaf">
      <rect x="154" y="108" width="52" height="18" rx="3" />
      <text x="180" y="120">replica</text>
    </g>
    <g class="leaf">
      <rect x="214" y="108" width="52" height="18" rx="3" />
      <text x="240" y="120">backup</text>
    </g>
  </svg>
</div>

<style>
  .preview {
    height: 100%;
    width: 100%;
    background: linear-gradient(180deg, #412402 0%, #1a0e02 100%);
    color: var(--accent-amber);
    padding: 8px 12px 4px;
    display: flex;
    flex-direction: column;
    gap: 2px;
    overflow: hidden;
  }

  .cost {
    display: flex;
    align-items: baseline;
    justify-content: space-between;
    font-family: var(--font-mono);
    font-size: 9px;
    text-transform: uppercase;
    letter-spacing: 0.14em;
    color: rgba(250, 199, 117, 0.55);
  }

  .cost-value {
    font-size: 13px;
    color: var(--accent-amber);
    letter-spacing: 0.04em;
  }
  .per {
    font-size: 9px;
    color: rgba(250, 199, 117, 0.5);
    margin-left: 2px;
  }

  .graph {
    width: 100%;
    height: 100%;
    flex: 1;
  }

  .graph rect {
    fill: rgba(250, 199, 117, 0.08);
    stroke: var(--accent-amber);
    stroke-width: 0.6;
  }

  .graph text {
    fill: var(--accent-amber);
    font-family: var(--font-mono);
    font-size: 8px;
    text-anchor: middle;
    dominant-baseline: middle;
    letter-spacing: 0.06em;
  }

  .node-top rect {
    fill: rgba(250, 199, 117, 0.16);
    stroke-width: 1;
  }

  .leaf rect {
    fill: rgba(250, 199, 117, 0.05);
  }
  .leaf text {
    fill: rgba(250, 199, 117, 0.78);
  }
</style>
