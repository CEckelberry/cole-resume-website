---
slug: linux
title: Linux Lessons from
italicWord: Hell
summary: A real bash shell in the browser. A foul-mouthed wasteland teacher heckles you through pipes, grep, and ssh, then hands you broken systems to fix.
role: solo — design, curriculum, full stack
stack:
  - SvelteKit
  - xterm.js
  - Go (Cloud Run)
  - PTY proxy
shipped: in progress
status: planned
liveDemoUrl: https://linux.cole-eckelberry.com
---

<script>
  import CSHero from '$lib/components/case-study/CSHero.svelte';
  import LiveDemo from '$lib/components/case-study/LiveDemo.svelte';
  import CSSection from '$lib/components/case-study/CSSection.svelte';
  import DecisionCard from '$lib/components/case-study/DecisionCard.svelte';
  import AdjacentProjects from '$lib/components/case-study/AdjacentProjects.svelte';
  import LinuxPreview from '$lib/components/projects/previews/LinuxPreview.svelte';
  import { PROJECT_BY_SLUG } from '$lib/content/projects';
  const project = PROJECT_BY_SLUG.linux;
</script>

<CSHero project={project} kind="A GAME" date="CONCEPT · 2026" role="solo — design + dev" shipped="in progress" />

<LiveDemo url={project.liveDemoUrl} title="linux lessons from hell" subtitle="a real shell, a fake teacher" placeholder><LinuxPreview /></LiveDemo>

<CSSection number="01" eyebrow="WHY" title="Most 'learn Linux' tutorials are" italic="sterile." accent="pink" lead="They explain `grep` on a textbook example and never let you feel the satisfaction of finding the one error line in 200,000 lines of nginx logs.">

This is a game where you actually use a real bash environment to solve increasingly broken scenarios while a foul-mouthed post-apocalyptic teacher heckles you through it. He insults your every keystroke when you're slow, congratulates you with backhanded compliments when you finally figure something out, and gives you the kind of tough-love mentoring that actually sticks. The humor lands because the teaching underneath it is real — you walk away genuinely knowing how to navigate a Linux command line.

</CSSection>

<CSSection number="02" eyebrow="CURRICULUM" title="Five tiers," italic="boss-style gates." accent="pink" lead="Each tier is gated by a scenario where the only way out is to type the right command. Banter is template-driven: success, partial, failure, suspicious-fast-success.">

- **Tier 1 — Survival basics.** `cd`, `ls`, `pwd`, `mkdir`, `rm`, `cp`, `mv`, then permissions. Boss: navigate a sprawling vault filesystem to find a hidden cache.
- **Tier 2 — Searching and slicing.** `grep -rABCiv`, `find` with predicates, `awk` and `cut`, `sed`. Boss: extract specific intel from a corrupted log archive.
- **Tier 3 — Pipes and process.** Pipes, redirection, `xargs`, `sort | uniq -c`, background jobs. Boss: pipeline a multi-step intel-gathering operation in one line.
- **Tier 4 — System health.** `ps`, `top`, `df`/`du`, `kill` with the right signals, `systemctl`, reading `/var/log/`. Boss: diagnose why the bunker generator service keeps dying.
- **Tier 5 — Across the wires.** `ssh`, key auth, `~/.ssh/config`, `scp`/`rsync`, `curl`/`wget`. Boss: exfiltrate data from a remote server you don't have a password for.

The endgame is the **sysadmin gauntlet** — a series of broken systems with stated symptoms (the app is returning 500s; users can't log in; disk is full but `du` says we should have space). Diagnose, fix, document. The teacher grades the answer.

</CSSection>

<CSSection number="03" eyebrow="THE TEACHER" title="A character, not" italic="a chatbot." accent="pink" lead="Voice: profane, dismissive of stupid mistakes, oddly invested when you finally get it. Banter is hard-coded for v1; LLM-generated lines are a v2 idea.">

Sample lines, in the spirit of the final voice:

> "Wow. You opened a directory. Should I get the band? Should I throw a parade?"
>
> "That's the third time you've tried `chmod` without numbers. You think the computer can read minds? Numbers. Or letters. Pick a side."
>
> "...okay. That was actually clever. I hate that you made me say that."
>
> "You used `find -exec rm`. Without `-i`. Without `-name`. You absolute weapon. Hope you didn't need that home directory."

Each scenario has 3–5 reaction templates per outcome (success, partial, failure, suspicious-fast-success), so a player solving the same level twice doesn't see identical lines. The reactions are JSON files keyed by scenario id and outcome, which makes adding a new tier a content-only change rather than a code change.

</CSSection>

<CSSection number="04" eyebrow="DECISIONS" title="Things I had to" italic="pick a side on." accent="pink">

<div class="decisions-grid">

<DecisionCard question="Container per session, or in-browser VM?" pickedLabel="picked container" pickedAccent="pink">
A real bash in a Cloud Run container with PTY proxying through a WebSocket beats a v86/webvm in-browser Linux on first-load weight (no 30 MB WASM blob), authenticity (it really is bash, not a re-implementation), and feel (no input lag). The cost is a small Go service that has to manage per-session containers and idle them aggressively.
</DecisionCard>

<DecisionCard question="How to keep the sandbox safe?" pickedLabel="picked harsh limits" pickedAccent="pink">
Per-session containers with no network egress (the `curl` lessons hit a fake DNS that resolves to a fixture server inside the cluster), 90-second idle timeout, 256 MB memory cap, no privilege escalation. If a session goes wild, it gets killed before it costs money. Abuse-mitigation is a bigger deal here than for any other project on the site.
</DecisionCard>

<DecisionCard question="Save and resume, or one-shot sessions?" pickedLabel="picked save" pickedAccent="pink">
Save the player's tier progress in Postgres keyed by a long-lived cookie. The container itself is ephemeral; the *progress* is durable. That way someone can leave mid-tier-3 and come back tomorrow without redoing earlier tiers.
</DecisionCard>

<DecisionCard question="Voice the teacher with TTS?" pickedLabel="skipped voice" pickedAccent="pink">
Adding voice means licensing a TTS, generating dozens of variants per line, and dealing with autoplay restrictions in browsers. Text-only at v1; voice is a stretch goal that probably never gets built unless someone offers to do the voice work.
</DecisionCard>

</div>

</CSSection>

<CSSection number="05" eyebrow="OPEN QUESTIONS" title="What I'm still" italic="figuring out." accent="pink">

- Multiplayer leaderboard for the gauntlet times. Fun, but requires anti-cheat thinking I haven't done.
- Save/resume across devices: tied to an email login, or the same long-lived cookie? Probably cookie at v1, email later.
- Should the curriculum branch by interest (sysadmin vs DevOps vs SRE)? Tempting, but it triples the content load. Sticking to one linear path at v1.
- Localization: the humor is very specifically English-language. Translating the teacher is a content design problem, not a code problem.

</CSSection>

<AdjacentProjects current="linux" />

<style>
  .decisions-grid {
    display: grid;
    grid-template-columns: 1fr;
    gap: var(--space-4);
  }
  @media (min-width: 720px) {
    .decisions-grid {
      grid-template-columns: 1fr 1fr;
    }
  }
</style>
