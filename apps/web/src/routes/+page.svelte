<script lang="ts">
  import Hero from '$lib/components/hero/Hero.svelte';
  import Terminal from '$lib/components/hero/Terminal.svelte';
  import StackMarquee from '$lib/components/hero/StackMarquee.svelte';
  import ProjectsSection from '$lib/components/projects/ProjectsSection.svelte';
  import AboutSection from '$lib/components/about/AboutSection.svelte';
  import ContactSection from '$lib/components/contact/ContactSection.svelte';
</script>

<Hero>
  {#snippet terminal()}
    <Terminal />
  {/snippet}
</Hero>

<StackMarquee />

<ProjectsSection />

<AboutSection />

<ContactSection />

<style>
  /* Page-level styles will land alongside Tasks 1.4 (marquee) → 1.12 (footer). */
</style>
