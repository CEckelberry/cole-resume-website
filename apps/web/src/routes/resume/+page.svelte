<script lang="ts">
  import SEO from '$lib/components/shell/SEO.svelte';
</script>

<SEO
  title="Résumé"
  description="Cole Eckelberry — Senior DevOps / Platform Engineer. 10+ years across GCP, AWS, and Azure."
  path="/resume"
/>

<article class="resume">
  <!-- ── Header ── -->
  <header class="r-header">
    <p class="eyebrow">résumé · 2026</p>
    <h1 class="r-name">Cole Eckelberry</h1>
    <p class="r-title">Senior DevOps / Platform Engineer</p>
    <div class="r-meta">
      <span>San Jose, United States</span>
      <span class="sep" aria-hidden="true">·</span>
      <a href="mailto:cole.eckelberry@gmail.com">cole.eckelberry@gmail.com</a>
      <span class="sep" aria-hidden="true">·</span>
      <a href="https://github.com/CEckelberry" target="_blank" rel="noopener">GitHub</a>
      <span class="sep" aria-hidden="true">·</span>
      <a href="https://linkedin.com/in/cole-eckelberry" target="_blank" rel="noopener">LinkedIn</a>
    </div>
    <p class="r-pitch">
      Cloud Architect and DevOps Technical Lead with 14 years of experience across GCP, AWS, Azure,
      and on-prem/hybrid systems. I've guided organizations through full cloud transformations,
      modernized legacy infrastructure, containerized monolithic stacks, and built the CI/CD
      pipelines that eliminate manual bottlenecks. I'm drawn to the hard problems — the migrations
      that seem impossible, the systems that haven't been touched in years, and the architectures
      that need to be reinvented to support the next order of magnitude.
    </p>
    <div class="r-actions">
      <a
        href="/Cole_Eckelberry_Senior_DevOps_Platform_Engineer.pdf"
        download="cole_eckelberry.pdf"
        class="btn-download"
      >
        <svg width="14" height="14" viewBox="0 0 14 14" fill="none" aria-hidden="true">
          <path
            d="M7 1v8M4 6l3 3 3-3M2 11h10"
            stroke="currentColor"
            stroke-width="1.5"
            stroke-linecap="round"
            stroke-linejoin="round"
          />
        </svg>
        download pdf
      </a>
    </div>
  </header>

  <div class="r-divider" aria-hidden="true"></div>

  <!-- ── Profile ── -->
  <section class="r-section">
    <p class="eyebrow section-label">profile</p>
    <div class="r-prose">
      <p>
        Cloud Architect and DevOps Technical Lead with a mission to help organizations thrive
        through the power of technology. Expertise in designing scalable, efficient solutions and
        translating complex infrastructure challenges into clean, maintainable systems. Over 10
        years of experience across AWS, GCP, and Azure — leading large-scale cloud migrations,
        modernizing legacy application stacks, and building the DevOps foundations that engineering
        teams rely on daily.
      </p>
      <p>
        Collaborative and forward-thinking: equally effective working with developers and
        stakeholders as going deep in infrastructure. Leads by example, mentors teams on
        cloud-native best practices, and consistently drives initiatives from concept through to
        production.
      </p>
    </div>
  </section>

  <div class="r-divider" aria-hidden="true"></div>

  <!-- ── Experience ── -->
  <section class="r-section">
    <p class="eyebrow section-label">experience</p>

    <div class="timeline">
      <div class="entry">
        <div class="entry-head">
          <div>
            <h2 class="entry-role">DevOps Technical Lead</h2>
            <p class="entry-org">EpexSpot <span class="dim">via Skyworkz</span> · Amsterdam</p>
          </div>
          <time class="entry-date">Jan 2024 — present</time>
        </div>
        <p class="entry-body">
          Driving cloud transformation and platform modernization across the full engineering
          organization. Led the lift-and-shift of over 60 applications from on-premises datacenters
          to GCP and AWS, then continued pushing the platform forward — containerizing applications,
          standardizing deployment practices, and building the developer tooling and workflows that
          teams depend on daily.
        </p>
        <ul class="entry-achievements">
          <li>
            Orchestrated the migration of 60+ applications to GCP and AWS, establishing foundational
            Terraform IaC templates for GCS, Pub/Sub, Cloud SQL, GCE, and Secret Manager as the
            organization-wide standard
          </li>
          <li>
            Containerized and deployed applications to Google Cloud Run and GKE, modernizing the
            deployment architecture away from legacy GCE VM-based workflows
          </li>
          <li>
            Built and standardized a company-wide CI/CD release process in GitLab, incorporating
            code scanning (SonarQube), PHP linters, bug fixers, and vulnerability scanners
          </li>
          <li>
            Architected Lisaria — a unified batch and queue management platform consolidating a
            fragmented series of legacy applications into a single scalable solution
          </li>
          <li>
            Overhauled developer onboarding and enabled local development workflows for
            containerized applications, significantly reducing friction across the engineering org
          </li>
          <li>
            Established comprehensive GCP-native logging, monitoring, and alerting, replacing legacy
            Centreon infrastructure with real-time observability across all services
          </li>
        </ul>
        <p class="entry-tools">
          <span class="tools-label">Toolset:</span>
          GCP, GKE, Cloud Run, GCS, Cloud SQL, GitLab, Terraform, Ansible, Docker, Kubernetes, SonarQube,
          PHP, Jira, AWS
        </p>
      </div>

      <div class="entry">
        <div class="entry-head">
          <div>
            <h2 class="entry-role">Senior Cloud Engineer</h2>
            <p class="entry-org">Pinterest · San Francisco</p>
          </div>
          <time class="entry-date">Mar 2023 — Oct 2023</time>
        </div>
        <p class="entry-body">
          Embedded within a large-scale infrastructure security team responsible for securing cloud
          environments serving hundreds of millions of users. Contributed to the design and
          enforcement of IAM security posture across both AWS and GCP, working closely with internal
          application teams to ensure all service-to-service communication followed least-privilege
          principles.
        </p>
        <ul class="entry-achievements">
          <li>
            Provided Cloud Architecture guidance on AWS IAM security stacks — role design, policy
            scoping, and secure resource access patterns across a massive multi-account AWS
            environment
          </li>
          <li>
            Authored and maintained Terraform modules for AWS IAM roles, subnets, security groups,
            and firewall rules, contributing to a standardized IaC approach used across the
            infrastructure org
          </li>
          <li>
            Managed GKE clusters: service account bindings, Workload Identity configurations, and
            Kubernetes RBAC ensuring secure pod-level access to GCP secrets and downstream services
          </li>
          <li>
            Contributed to ArgoCD-based GitOps workflows for deploying and managing GKE workloads,
            supporting continuous delivery pipelines across internal platform services
          </li>
        </ul>
        <p class="entry-tools">
          <span class="tools-label">Toolset:</span>
          AWS, GCP, GKE, Lambda, Terraform, ArgoCD, Kubernetes, IAM, Secret Manager, Workload Identity,
          GitHub
        </p>
      </div>

      <div class="entry">
        <div class="entry-head">
          <div>
            <h2 class="entry-role">Solutions Architect</h2>
            <p class="entry-org">Hex Technologies · San Francisco</p>
          </div>
          <time class="entry-date">May 2022 — Jan 2023</time>
        </div>
        <p class="entry-body">
          Collaborated with data scientists worldwide to develop machine learning solutions within
          Hex's hosted Jupyter Notebook environment. Provided feedback to engineering teams on
          product performance, customer sentiment, and new feature ideas, while leveraging Python
          and cloud technologies to create custom solutions for clients.
        </p>
        <ul class="entry-achievements">
          <li>
            Consistently worked with Python packages and data science models for customer machine
            learning implementations
          </li>
          <li>
            Helped create custom solutions for customers using Google Cloud and AWS integrations
          </li>
          <li>Created architecture diagrams, demo projects, dbt models, and data visualizations</li>
        </ul>
        <p class="entry-tools">
          <span class="tools-label">Toolset:</span>
          AWS, Kubernetes, GitHub, CircleCI, Terraform, Python, JavaScript, Datadog, Sentry
        </p>
      </div>

      <div class="entry">
        <div class="entry-head">
          <div>
            <h2 class="entry-role">Cloud Engineer</h2>
            <p class="entry-org">MediaMonks · San Francisco</p>
          </div>
          <time class="entry-date">May 2021 — May 2022</time>
        </div>
        <p class="entry-body">
          Built serverless architectures and automated workflows for internal tools and projects.
          Played a key role in implementing ETL processes using Python and Node.js across Google
          Cloud. Developed API Gateway solutions with custom serverless architecture. Provided
          technical documentation, training, and engineering services, collaborating with diverse
          teams on data integration, platform migrations, and cloud architecture consulting.
        </p>
        <p class="entry-tools">
          <span class="tools-label">Toolset:</span>
          GCP, Kubernetes, Cloud Functions, Cloud Run, GitHub, BigQuery, Terraform, Fivetran
        </p>
      </div>

      <div class="entry">
        <div class="entry-head">
          <div>
            <h2 class="entry-role">Sr. Systems Engineer</h2>
            <p class="entry-org">RoundTower Technologies · Cincinnati</p>
          </div>
          <time class="entry-date">Sept 2016 — Oct 2018</time>
        </div>
        <p class="entry-body">
          Implemented advanced snapshot technology reducing backup times by over 70% for VMware and
          NAS systems. Deep expertise in cloud networking — VPCs, firewall rules, disaster recovery,
          and high availability — delivering robust infrastructure solutions. Saved a client over
          $6,000 monthly in AWS S3 storage costs.
        </p>
        <ul class="entry-achievements">
          <li>
            Implemented team-first array-based snapshot technology, decreasing backup times by over
            70% for VMware and NAS backups while reducing performance load
          </li>
          <li>
            Saved a customer over $6,000 a month in AWS S3 storage costs due to a backup
            misconfiguration discovery
          </li>
        </ul>
        <p class="entry-tools">
          <span class="tools-label">Toolset:</span>
          AWS, GCP, Azure, S3, Ansible, GitHub, Perl, PowerShell, Bash, VMware, ServiceNow, Python, TypeScript/JS,
          Node
        </p>
      </div>

      <div class="entry">
        <div class="entry-head">
          <div>
            <h2 class="entry-role">Cloud / Systems Engineer</h2>
            <p class="entry-org">
              National Institute of Allergic and Infectious Diseases (NIAID) · Washington, D.C.
            </p>
          </div>
          <time class="entry-date">Apr 2014 — May 2016</time>
        </div>
        <p class="entry-body">
          Administered Commvault backups for over 1,400 clients, managing a total data size of 5.5
          petabytes. Managed NetApp storage arrays — capacity adjustments, new shares, and regular
          maintenance — and supported cloud adoption initiatives across the institute.
        </p>
        <ul class="entry-achievements">
          <li>
            Developed automated backup monitoring scripts to streamline identification of potential
            failures and ensure timely data protection
          </li>
          <li>
            Assisted in scaling storage solutions to accommodate growing research data, ensuring
            high availability
          </li>
          <li>
            Commvault backup administration for 1,400 clients including VMs, NAS, and physical
            clients with a total data size of 5.5 PB
          </li>
        </ul>
        <p class="entry-tools">
          <span class="tools-label">Toolset:</span>
          NetApp, VMware, AWS, Commvault, Bash, Puppet, ServiceNow, Linux, HPC
        </p>
      </div>

      <div class="entry">
        <div class="entry-head">
          <div>
            <h2 class="entry-role">Sr. Systems Administrator</h2>
            <p class="entry-org">General Electric · Cincinnati</p>
          </div>
          <time class="entry-date">Jun 2012 — Jun 2014</time>
        </div>
        <p class="entry-body">
          Oversaw an extensive environment of over 9,200 servers and more than 7 petabytes of data.
          Responsible for daily operational backup status reports to executive-level management,
          ensuring the protection of one of the world's largest data storage environments.
        </p>
        <ul class="entry-achievements">
          <li>
            Monitored and managed environments with a combined monthly volume exceeding 7 PB and
            9,200 servers
          </li>
          <li>
            Ran daily operations backup call reporting environment status to GE executive management
          </li>
          <li>
            Worked on a team of 26 engineers to protect one of the largest data storage environments
            in the world
          </li>
          <li>
            Helped migrate over 600 servers from Hamilton, Ohio to the Cincinnati datacenter in
            under 3 months
          </li>
        </ul>
        <p class="entry-tools">
          <span class="tools-label">Toolset:</span>
          NetBackup, Avamar (EMC), VMware, Nagios, ServiceNow, Linux/Unix, Bash
        </p>
      </div>
    </div>
  </section>

  <div class="r-divider" aria-hidden="true"></div>

  <!-- ── Education ── -->
  <section class="r-section">
    <p class="eyebrow section-label">education</p>

    <div class="timeline">
      <div class="entry">
        <div class="entry-head">
          <div>
            <h2 class="entry-role">B.S. Information Technology</h2>
            <p class="entry-org">Western Governors University</p>
          </div>
          <time class="entry-date">May 2014 — May 2019</time>
        </div>
        <ul class="entry-achievements">
          <li>Cloud technologies · Network and Security</li>
          <li>Oracle SQL &amp; MySQL Database Administration</li>
          <li>Data Science, Python programming, Java programming</li>
          <li>Project Management, Encryption policies, PII</li>
        </ul>
      </div>

      <div class="entry">
        <div class="entry-head">
          <div>
            <h2 class="entry-role">Full-Stack Web Development</h2>
            <p class="entry-org">Springboard</p>
          </div>
          <time class="entry-date">Jun 2019 — Sept 2020</time>
        </div>
        <p class="entry-body">
          700 hours of hands-on course material with 1:1 industry expert mentor oversight and
          completion of 4 in-depth portfolio projects. Front-end, back-end, databases, and data
          structures &amp; algorithms.
        </p>
      </div>
    </div>
  </section>

  <div class="r-divider" aria-hidden="true"></div>

  <!-- ── Skills ── -->
  <section class="r-section">
    <p class="eyebrow section-label">skills</p>
    <div class="skills-grid">
      {#each ['Google Cloud', 'AWS', 'Azure', 'Terraform', 'Kubernetes', 'Docker', 'Python', 'Go', 'TypeScript / JS', 'Postgres', 'Kafka', 'REST / GraphQL', 'GitLab CI', 'ArgoCD', 'Ansible', 'GKE', 'Cloud Run', 'Pub/Sub'] as skill (skill)}
        <span class="skill-tag">{skill}</span>
      {/each}
    </div>
  </section>
</article>

<style>
  .resume {
    max-width: 800px;
    margin: 0 auto;
    padding-bottom: var(--space-8);
    display: flex;
    flex-direction: column;
    gap: var(--space-7);
  }

  /* ── Header ── */
  .r-header {
    display: flex;
    flex-direction: column;
    gap: var(--space-3);
  }

  .r-name {
    font-family: var(--font-sans);
    font-weight: 500;
    font-size: var(--type-h1);
    line-height: var(--type-h1-lh);
    letter-spacing: -0.02em;
    color: var(--text-primary);
    margin: 0;
  }

  .r-title {
    font-family: var(--font-mono);
    font-size: var(--type-meta);
    letter-spacing: 0.14em;
    text-transform: uppercase;
    color: var(--accent-teal);
    margin: 0;
  }

  .r-meta {
    display: flex;
    flex-wrap: wrap;
    align-items: center;
    gap: 8px;
    font-size: var(--type-body-sm);
    color: var(--text-secondary);
  }

  .r-meta a {
    color: var(--text-secondary);
    text-decoration: none;
    transition: color var(--dur-fast) var(--ease-out);
  }

  .r-meta a:hover {
    color: var(--text-primary);
  }

  .sep {
    color: var(--text-muted);
  }

  .r-pitch {
    font-size: var(--type-body);
    line-height: 1.75;
    color: var(--text-secondary);
    margin: 0;
    max-width: 640px;
  }

  .r-actions {
    margin-top: var(--space-2);
  }

  .btn-download {
    display: inline-flex;
    align-items: center;
    gap: 6px;
    padding: 9px 16px;
    border-radius: var(--radius-pill);
    border: 0.5px solid var(--border-default);
    background: var(--bg-surface-2);
    color: var(--text-secondary);
    font-family: var(--font-mono);
    font-size: var(--type-meta);
    letter-spacing: 0.04em;
    text-decoration: none;
    backdrop-filter: blur(12px);
    transition:
      background var(--dur-fast) var(--ease-out),
      color var(--dur-fast) var(--ease-out),
      border-color var(--dur-fast) var(--ease-out),
      transform var(--dur-default) var(--ease-spring);
  }

  .btn-download:hover {
    background: var(--bg-surface-3);
    border-color: var(--border-strong);
    color: var(--text-primary);
    transform: translateY(-2px);
  }

  /* ── Divider ── */
  .r-divider {
    height: 0.5px;
    background: var(--border-subtle);
  }

  /* ── Section ── */
  .r-section {
    display: grid;
    grid-template-columns: 120px 1fr;
    gap: var(--space-5);
    align-items: start;
  }

  @media (max-width: 600px) {
    .r-section {
      grid-template-columns: 1fr;
      gap: var(--space-4);
    }
  }

  .section-label {
    padding-top: 3px;
  }

  /* ── Prose ── */
  .r-prose {
    display: flex;
    flex-direction: column;
    gap: var(--space-3);
  }

  .r-prose p {
    font-size: var(--type-body);
    line-height: 1.75;
    color: var(--text-secondary);
    margin: 0;
  }

  /* ── Timeline ── */
  .timeline {
    display: flex;
    flex-direction: column;
    gap: var(--space-7);
  }

  .entry {
    display: flex;
    flex-direction: column;
    gap: var(--space-3);
    padding-left: var(--space-5);
    border-left: 1.5px solid var(--border-subtle);
    position: relative;
  }

  .entry::before {
    content: '';
    position: absolute;
    left: -5px;
    top: 6px;
    width: 8px;
    height: 8px;
    border-radius: 50%;
    background: var(--accent-teal);
    box-shadow: 0 0 0 3px var(--bg-canvas);
  }

  .entry-head {
    display: flex;
    justify-content: space-between;
    align-items: flex-start;
    gap: var(--space-4);
    flex-wrap: wrap;
  }

  .entry-role {
    font-family: var(--font-sans);
    font-weight: 500;
    font-size: var(--type-h3);
    line-height: var(--type-h3-lh);
    color: var(--text-primary);
    margin: 0;
  }

  .entry-org {
    font-size: var(--type-body-sm);
    color: var(--text-tertiary);
    margin: 2px 0 0;
  }

  .dim {
    color: var(--text-muted);
  }

  .entry-date {
    font-family: var(--font-mono);
    font-size: var(--type-meta);
    color: var(--text-muted);
    white-space: nowrap;
    flex-shrink: 0;
  }

  .entry-body {
    font-size: var(--type-body);
    line-height: 1.75;
    color: var(--text-secondary);
    margin: 0;
  }

  .entry-achievements {
    display: flex;
    flex-direction: column;
    gap: 6px;
    padding-left: var(--space-4);
    margin: 0;
  }

  .entry-achievements li {
    font-size: var(--type-body);
    line-height: 1.7;
    color: var(--text-secondary);
  }

  .entry-tools {
    font-family: var(--font-mono);
    font-size: var(--type-meta);
    color: var(--text-tertiary);
    line-height: 1.6;
    margin: 0;
  }

  .tools-label {
    color: var(--text-muted);
    margin-right: 4px;
  }

  /* ── Skills ── */
  .skills-grid {
    display: flex;
    flex-wrap: wrap;
    gap: var(--space-2);
  }

  .skill-tag {
    display: inline-block;
    padding: 5px 12px;
    border-radius: var(--radius-pill);
    border: 0.5px solid var(--border-default);
    background: var(--bg-surface-2);
    font-family: var(--font-mono);
    font-size: var(--type-meta);
    color: var(--text-secondary);
    letter-spacing: 0.02em;
  }

  /* ── Print ── */
  @media print {
    .r-actions {
      display: none;
    }

    .resume {
      gap: 24px;
    }

    .entry::before {
      box-shadow: none;
    }
  }
</style>
