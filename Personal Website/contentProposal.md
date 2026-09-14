# Website Content & Architecture Proposal: victorh.dev

**Target Domain:** [victorh.dev](https://victorh.dev)  
**Target Audience:** Engineering Executives (VPs, CTOs), Hiring Managers, Staff/Principal Engineers, Technical Recruiters  
**Core Goal:** Position Victor Hernandez Guzman as a premier **Principal Software Engineer & Identity Architect** with 20+ years of experience (15+ in IAM & Hyperscale Cloud Infrastructure), driving high-conversion hiring leads and establishing technical thought leadership.

---

## Technical & Architectural Principles

### Keep it Lightweight
Performance is your first unspoken technical interview. The site must serve as concrete proof of engineering discipline through ultra-fast load times and lean architecture:
- **Target Metrics:** 100/100 Lighthouse score across Performance, Accessibility, Best Practices, and SEO.
- **Fast Core Web Vitals:** Largest Contentful Paint (LCP) under 1.2 seconds, Interaction to Next Paint (INP) under 50 milliseconds.
- **Zero Heavy Bundles:** Built with a modern Static Site Generator (Astro, Vite, or Next.js static export) hosted on a global CDN edge (Vercel, Netlify, or Cloudflare Pages).
- **Minimal JavaScript Footprint:** Zero heavy framework runtimes shipped to the client where plain HTML/CSS suffices.

### Design with Intent
The UI must feel minimalist, premium, fully responsive, and accessible:
- **Accessibility & Contrast:** High-contrast text adhering strictly to WCAG AAA standards. Clear focus states for keyboard navigation.
- **Clean Typography:** Paired Google Fonts like Inter for body text and JetBrains Mono for technical badges and code snippets.
- **Mobile First & Responsive:** Fluid layouts, flexible grids, and touch targets of at least 44x44 pixels for seamless mobile browsing.
- **Minimalist Aesthetic:** Dark mode default using clean obsidian slate tones (`#0B0F19`), subtle borders (`#1E293B`), and electric cyan highlight accents (`#00E5FF`).

---

## Content Structure Overview

```
┌────────────────────────────────────────────────────────────────────────┐
│ 1. HERO / INTRO SECTION                                                │
│    Elevator pitch, principal title, value proposition, CTAs            │
├────────────────────────────────────────────────────────────────────────┤
│ 2. ABOUT / CORE EXPERTISE                                              │
│    Executive bio, interactive tech stack grid & badge system           │
├────────────────────────────────────────────────────────────────────────┤
│ 3. FEATURED PROJECTS / CASE STUDIES                                   │
│    4 deep-dive case studies: AxoEdge Suite, MSFT dSTS Data Layer       │
│    Migration, MSFT ID4S SAML Support, MSFT MSA WebAuthn              │
├────────────────────────────────────────────────────────────────────────┤
│ 4. PROFESSIONAL EXPERIENCE / CAREER HIGHLIGHTS                         │
│    High-impact career timeline (DocuSign, Microsoft, Salesforce)       │
│    + Downloadable PDF resume link                                      │
├────────────────────────────────────────────────────────────────────────┤
│ 5. WRITING & THOUGHT LEADERSHIP                                        │
│    Curated technical deep-dives and architectural articles             │
├────────────────────────────────────────────────────────────────────────┤
│ 6. CONTACT & FOOTER                                                    │
│    1-click email copy widget, social links, availability badge         │
└────────────────────────────────────────────────────────────────────────┘
```

---

## 1. Hero / Intro Section

### Purpose
The immediate 3-second elevator pitch. When a VP of Engineering or Recruiter lands on `victorh.dev`, they must instantly recognize Victor as an enterprise-grade Principal Software Engineer and Identity Architect who builds resilient, global-scale identity platforms.

### Content Specification

- **Headline / Greeting:**
  > **Hi, I'm Victor Hernandez Guzman**

- **Compelling Professional Title:**
  > **Principal Software Engineer & Identity Architect**

- **Core Tagline (1-Sentence Value Proposition):**
  > *"Architecting hyperscale cloud identity providers, zero-trust token infrastructure, and resilient distributed systems serving hundreds of millions of users worldwide."*

- **Trust Badges / Micro Highlights:**
  - 20+ Years Software Engineering Experience
  - 15+ Years Specializing in Identity & Access Management (IAM)
  - Ex-DocuSign (Principal), Ex-Microsoft (Principal & Senior), Ex-Salesforce (Lead MTS)

- **Primary Call-to-Action (CTA) Buttons:**
  1. **[ Get in Touch ]** *(Primary Accent Button: Smooth scroll to Contact section or mailto trigger)*
  2. **[ View Resume (PDF) ]** *(Secondary Outline Button: Direct link to downloadable PDF resume)*
  3. **[ Explore Case Studies ]** *(Ghost Button: Jump directly to Section 3)*

---

## 2. About / Core Expertise

### Purpose
To communicate Victor's engineering philosophy, technical identity, and deep domain mastery without forcing visitors to read dense paragraphs.

### Content Specification

#### Executive Bio
> "I am a Principal Software Engineer with over two decades of engineering experience, specializing in large-scale **Identity & Access Management (IAM)** and **Hyperscale Cloud Infrastructure**.
>
> My career spans foundational technical leadership at **DocuSign**, **Microsoft**, and **Salesforce**. I have engineered bare-metal authentication servers powering Microsoft Azure (dSTS), led pioneer FIDO2 passwordless WebAuthn rollouts for hundreds of millions of accounts, and architected multi-agent AI execution systems and high-throughput data sync pipelines. I combine protocol mastery (OAuth 2.0, OIDC, SAML, WebAuthn) with zero-trust security, distributed systems design, and engineering governance."

#### Tech Stack & Competency Grid

| Category | Core Technologies & Protocols |
| :--- | :--- |
| **Identity Protocols & Standards** | OAuth 2.0 · OpenID Connect (OIDC) · SAML 2.0 · SCIM · WebAuthn / FIDO2 · WS-Federation · DPoP · CIMD · DCR |
| **System Architecture & Security** | Identity Providers (IdP) · Zero Trust Architecture · Datacenter STS · Change Data Capture (CDC) · Microservices · High Availability |
| **Languages & Frameworks** | C# / .NET · ASP.NET Core · Java · C++ · TypeScript / Node.js · Python · SQL |
| **AI Systems & Vector DBs** | Retrieval-Augmented Generation (RAG) · Model Context Protocol (MCP) · PostgreSQL pgvector · LLM Benchmarking · YOLO CV |
| **Cloud & Distributed Infrastructure**| Microsoft Azure · Entra ID / Azure AD · Service Fabric · SQL Server · Message Queues · ONNX Runtime |

---

## 3. Featured Projects / Case Studies

### Purpose
Proof of work. This section provides detailed technical case studies highlighting problem framing, system architecture, engineering trade-offs, and quantitative results.

### Content Specification (4 Featured Projects)

#### Case Study 1: AxoEdge AI Product Suite & Applied Engineering
* **Domain:** AI Agent Orchestration, Computer Vision Benchmarking, RAG & Vector Systems
* **The Problem:** Following a layoff from DocuSign, Victor co-founded AxoEdge, solving multi-agent software development fragmentation and engineering challenges for logistics and real estate clients.
* **Solutions & Architectural Highlights:**
  - **Stanley (AI Orchestration Platform):** Architected a platform unifying multi-agent CLIs (`claude`, `copilot`, `opencode`) with structured task workflows (`Working`, `NeedsApproval`, `Done`), automated Git worktree isolation per task to eliminate merge conflicts, and centralized policy control (`server.settings.json`).
  - **ProLogistiks (Pallet Counting CV vs LLM Study):** Evaluated automated warehouse pallet counting by running parallel benchmarks across 7 open-source multimodal LLMs vs YOLO object detection. Discovered top LLMs had 40% count deviation at 10 seconds per image, while YOLO achieved ~80% accuracy in milliseconds. Delivered a data-backed recommendation in under a month.
  - **11 Casitas (Real Estate RAG Bot):** Built a real-time conversational agent on Facebook Messenger and Instagram Direct using a PostgreSQL + pgvector RAG pipeline exposed to LLMs via a Model Context Protocol (MCP) tool layer.
* **The Impact:**
  - Replaced ad-hoc AI coding sessions with governed task execution, achieving an estimated 30% reduction in worktree conflicts.
  - Saved logistics client from investing in unreliable LLM vision architecture through quantitative benchmark proof.
  - Reduced real estate lead response time from hours to ~1 second, cutting estimated lead drop-off by 30%.

---

#### Case Study 2: MSFT Datacenter STS (dSTS) Data Layer Migration
* **Domain:** Bare-Metal Regional Identity Provider, Shadow Traffic Validation, Storage Migration
* **The Problem:** dSTS is Microsoft's bare-metal, tier-0 regional identity provider for Azure, booting immediately after networking in every datacenter with zero service dependencies. The legacy codebase relied on an unsupported distributed-hashtable store, causing recurring livesite incidents and blocking the roadmap to modern auth protocols (OAuth/OIDC).
* **Your Solution (As Principal Software Engineer):**
  - Conducted a comprehensive landscape assessment and led the strategic pivot to adapt Entra ID's modern login stack over a new bare-metal storage model.
  - Evaluated bare-metal database options and selected SQL Server + Webstore (Microsoft's sharding layer) based on operational ownership and dedicated DB management support.
  - Authored the end-to-end migration design and rollout strategy, utilizing gateway-level shadow traffic to run old and new data layers in parallel without production risk.
* **The Impact:**
  - Delivered a validated prototype in Q1 2023 demonstrating equivalent login success rates (~99%+ parity under shadow traffic).
  - De-risked full production cutover and unblocked Azure's tier-0 modernization roadmap toward OAuth/OIDC.

---

#### Case Study 3: MSFT ID4S dSTS SAML & WS-FED Token Validation SDK
* **Domain:** Security Libraries, WS-Federation / SAML, Cross-Org Platform Standards
* **The Problem:** As Microsoft standardized identity across services, lower-tier Azure services needed a reliable, secure way to validate dSTS-issued SAML/WS-FED tokens. The core Identity for Services (ID4S) team lacked resources to build custom support for dSTS's unique endpoints, certificate rotation strategies, and claim sets.
* **Your Solution (As Principal Software Engineer):**
  - Led the end-to-end design and implementation of a reusable token validation library integrated directly into ID4S shared packages.
  - Architected validation logic supporting dSTS-specific WS-FED/WS-TRUST flows, custom metadata endpoints, automated certificate rotation, and delegation models.
  - Partnered with ID4S and dSTS engineering teams to ensure zero service disruption and full parity with deprecated legacy validators.
  - Produced integration guides, support SOPs, and developer documentation for cross-team adoption.
* **The Impact:**
  - Adopted across Azure services to validate dSTS tokens securely across Microsoft datacenters.
  - Eliminated token validation fragmentation and established a unified standard for service-to-service authentication.

---

#### Case Study 4: MSFT Account (MSA) Pioneer FIDO2 / WebAuthn Passwordless Launch
* **Domain:** Passwordless Authentication, W3C Standards, Client & UX Integration
* **The Problem:** Transitioning hundreds of millions of Microsoft Account users to public key cryptography (WebAuthn/FIDO2) while the underlying W3C specification was still an evolving working draft, tied to a strict Windows 10 OS ship date.
* **Your Solution (As Senior Software Engineer):**
  - Spearheaded the browser and UX-layer integration of WebAuthn for Microsoft Account (MSA), making Microsoft the first major identity provider to ship FIDO2 passwordless auth at scale.
  - Coordinated across 4 engineering groups (MSA, Microsoft Edge, Windows Core, W3C PM group).
  - Successfully advocated for requiring user-verification-only authenticators to deliver true passwordless login rather than simple 2-factor authentication.
  - Contributed implementation telemetry and feedback directly back to W3C working groups.
* **The Impact:**
  - Shipped passwordless authentication on schedule for Windows 10 to hundreds of millions of MSA users.
  - Influenced global web standards for WebAuthn and FIDO2 authentication.

---

## 4. Professional Experience / Career Highlights

### Purpose
High-level career timeline highlighting technical leadership scope, organizational impact, and career progression across industry leaders.

### Content Specification

```
┌────────────────────────────────────────────────────────────────────────┐
│ DocuSign                                                  2024 – Present│
│ Principal Software Engineer (Identity Organization)                    │
│ • Defining long-term architectural roadmap for core authn/authz.       │
│ • Orchestrating platform security modernization and passwordless auth. │
└────────────────────────────────────────────────────────────────────────┘
                                    │
                                    ▼
┌────────────────────────────────────────────────────────────────────────┐
│ Microsoft                                                  2022 – 2024 │
│ Principal Software Engineer (Azure Datacenter STS & ID4S)              │
│ • Led bare-metal data layer migration & shadow traffic validation.     │
│ • Built cross-Azure WS-FED/SAML token validation SDKs for ID4S.        │
└────────────────────────────────────────────────────────────────────────┘
                                    │
                                    ▼
┌────────────────────────────────────────────────────────────────────────┐
│ Salesforce                                                 2019 – 2022 │
│ Lead Software Engineer / Lead MTS (Identity Platform)                  │
│ • Designed CDC user sync pipeline for Global IDP (90M+ users synced).  │
│ • Governed OAuth, SAML, SCIM, OIDC standards across 8M+ app client integrations.│
└────────────────────────────────────────────────────────────────────────┘
                                    │
                                    ▼
┌────────────────────────────────────────────────────────────────────────┐
│ Microsoft                                                  2009 – 2019 │
│ Senior Software Engineer (Microsoft Account & Windows Identity)        │
│ • First IdP engineer to ship FIDO2 WebAuthn passwordless at scale.     │
│ • Integrated MSA across Windows 8/10, Xbox One, and HoloLens.          │
└────────────────────────────────────────────────────────────────────────┘
```

#### Resume Download Link
> 📄 **Formal Resume Available:** [Download Full Executive Resume (PDF)](https://victorh.dev/resume.pdf)

---

## 5. Writing & Thought Leadership

### Purpose
Demonstrate technical communication skills and domain mastery through curated technical deep-dives and architectural analysis.

### Content Specification (Proposed Technical Articles)

1. 📖 **"Benchmarking Multimodal LLMs vs. Dedicated Computer Vision in Production Logistics"**  
   *A real-world empirical study comparing zero-shot multimodal LLMs against YOLO object detection for image-based inventory tracking.*

2. 📖 **"Architecting Task Orchestration and Git Worktree Isolation for Multi-Agent AI Development"**  
   *How Stanley replaces ad-hoc AI coding CLI sessions with structured task states, worktree isolation, and multi-agent policy enforcement.*

3. 📖 **"Bare-Metal Identity: Operating Azure's Tier-0 Datacenter STS at Hyperscale"**  
   *Lessons in fault tolerance, shadow-traffic validation, and zero-dependency identity architectures for root cloud trust anchors.*

4. 📖 **"Passwordless at Scale: Shipping FIDO2 WebAuthn to Hundreds of Millions of Users"**  
   *Navigating evolving W3C specifications, browser integration layers, and user verification authenticators in consumer identity.*

5. 📖 **"RAG Infrastructure Design: Leveraging PostgreSQL, pgvector, and MCP for Real-Time Conversational Interfaces"**  
   *Structuring vector chunking pipelines, MCP tool exposure, and sub-second response latency for social messaging integration.*

---

## 6. Contact & Footer

### Purpose
Low-friction conversion. Enable hiring managers and executives to reach out with zero barrier to entry.

### Content Specification

#### Contact Card
- **Headline:** *"Let's Connect & Build Resilient Systems"*
- **Sub-headline:** *"Open to Principal/Staff Software Engineering, Identity Architecture, and technical advisory roles."*

#### Interactive Elements
- **Direct Email:** `victormh@outlook.com`
- **Actions:**
  - **[ Copy Email Address ]** *(Instant toast message: "Email address copied!")*
  - **[ Send Email ]** *(Triggers mailto:victormh@outlook.com)*
- **Social Links:**
  - **LinkedIn:** [linkedin.com/in/victormhdz](https://www.linkedin.com/in/victormhdz)
  - **GitHub:** [github.com/victorm-hernandez](https://github.com)
  - **Domain:** [victorh.dev](https://victorh.dev)

#### Availability Status Badge
> 🟢 **Status:** Available for select Principal / Staff IAM & Cloud Architecture opportunities.

#### Footer
> `© 2026 Victor Hernandez Guzman · victorh.dev`  
> *Built with clean semantic HTML, vanilla CSS, and zero tracking scripts.*


## Create Socials

Medium: https://medium.com/@victorhdev
dev.to: https://dev.to/victormh