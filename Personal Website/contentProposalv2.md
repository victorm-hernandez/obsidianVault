# Website Content & Architecture Proposal v2: victorh.dev

**Target Domain:** [victorh.dev](https://victorh.dev)  
**Target Audience:** Fellow Engineers, Tech Leads, Architects, Engineering Managers, and Technical Peers  
**Core Goal:** Showcase Victor Hernandez Guzman's technical depth, engineering projects, and architecture experience in Identity & Access Management (IAM) and Hyperscale Infrastructure through an authentic, developer-to-developer lens.

---

## Technical & Architectural Principles

### Keep it Lightweight
Performance is your first unspoken technical interview. The site must serve as concrete proof of engineering discipline through ultra-fast load times and lean architecture:
- **Target Metrics:** 100/100 Lighthouse score across Performance, Accessibility, Best Practices, and SEO.
- **Fast Core Web Vitals:** Largest Contentful Paint (LCP) under 1.2 seconds, Interaction to Next Paint (INP) under 50 milliseconds.
- **Zero Heavy Bundles:** Built with a modern Static Site Generator (Astro, Vite, or Next.js static export) hosted on a global CDN edge (Vercel, Netlify, or Cloudflare Pages).
- **Minimal JavaScript Footprint:** Zero heavy framework runtimes shipped to the client where plain HTML/CSS suffices.

### Design with Intent
The UI should feel minimalist, modern, accessible, and clean:
- **Accessibility & Contrast:** High-contrast text adhering strictly to WCAG AAA standards. Clear focus states for keyboard navigation.
- **Clean Typography:** Paired Google Fonts like Inter for body text and JetBrains Mono for technical badges and code snippets.
- **Mobile First & Responsive:** Fluid layouts, flexible grids, and touch targets of at least 44x44 pixels for seamless mobile browsing.
- **Developer Aesthetic:** Dark mode default using clean obsidian slate tones (`#0B0F19`), subtle borders (`#1E293B`), and electric cyan highlight accents (`#00E5FF`).


### Other Design Hints
Inspire your design on the website https://opencode.ai/ 

Add support for both light and dark theme

Dont use complex images and rely on simple graphics, if an image is required to decorate it should be simple, using lines.
    Use the avatar photo I provided for the hero section on the homepage.

**Design and build a modern developer-focused technology website inspired by the visual style of the reference screenshots I provided.**

 The website should feel like a premium, open-source developer tool or AI infrastructure company. The design should be **minimalist, technical, monochromatic, sophisticated, and highly structured**, with a strong emphasis on typography, whitespace, grids, thin borders, and terminal/developer aesthetics.
  
> **Important:** Use the screenshots as visual inspiration for the overall design language and layout, but do not copy their branding, logo, text, illustrations, or exact content. Create an original website using the same general visual principles.

## Overall Visual Style

- Dark, almost-black background, approximately `#0D0B0B` or similar.
- Primary text should be an off-white/light gray rather than pure white.
- Secondary text should be muted gray.
- Use subtle gray borders throughout the page.
- Avoid gradients, excessive shadows, glassmorphism, rounded cards, colorful backgrounds, or generic SaaS styling.
- The visual style should feel **technical, editorial, developer-oriented, and slightly retro-futuristic**.
- Think of a combination of:
    - Unix terminal
    - developer documentation
    - modern open-source project
    - technical research paper
    - premium software company
- The page should feel intentionally sparse rather than crowded.

---

## Content Structure Overview

```
┌────────────────────────────────────────────────────────────────────────┐
│ 1. HERO / INTRO SECTION                                                │
│    Elevator pitch, title, tagline, direct CTAs                         │
├────────────────────────────────────────────────────────────────────────┤
│ 2. ABOUT / CORE EXPERTISE                                              │
│    Developer-to-developer bio, tech stack grid & badge system          │
├────────────────────────────────────────────────────────────────────────┤
│ 3. FEATURED PROJECTS / CASE STUDIES                                   │
│    4 technical case studies: AxoEdge Suite, MSFT dSTS Data Layer       │
│    Migration, MSFT ID4S SAML Support, MSFT MSA WebAuthn              │
├────────────────────────────────────────────────────────────────────────┤
│ 4. PROFESSIONAL EXPERIENCE / CAREER HIGHLIGHTS                         │
│    Career timeline (DocuSign, Microsoft, Salesforce)                   │
│    + PDF resume download link                                          │
├────────────────────────────────────────────────────────────────────────┤
│ 5. WRITING & THOUGHT LEADERSHIP                                        │
│    Curated technical deep-dives and engineering write-ups              │
├────────────────────────────────────────────────────────────────────────┤
│ 6. CONTACT & FOOTER                                                    │
│    1-click email copy widget (contact@victorh.dev), social links       │
└────────────────────────────────────────────────────────────────────────┘
```

---

## 1. Hero / Intro Section

### Purpose
The immediate elevator pitch. When another engineer or team lead lands on `victorh.dev`, they should immediately get who Victor is, what he builds, and how to dive into his work.

### Content Specification

- **Headline / Greeting:**
  > **Hi, I'm Victor Hernandez Guzman**

- **Professional Title:**
  > **Principal Software Engineer & Identity Architect**

- **Core Tagline:**
  > *"Building hyperscale cloud identity providers, zero-trust token systems, and resilient distributed platforms."*

- **Trust Badges / Micro Highlights:**
  - 20+ Years Software Engineering
  - 15+ Years Specializing in Identity & Access Management (IAM)
  - DocuSign, Microsoft, Salesforce

- **Primary Call-to-Action (CTA) Buttons:**
  1. **[ Get in Touch ]** *(Primary Accent Button: Smooth scroll to Contact section or triggers mailto:contact@victorh.dev)*
  2. **[ View Resume ]** *(Secondary Button: Opens downloadable PDF resume)*
  3. **[ Explore Projects ]** *(Ghost Button: Jumps directly to Case Studies section)*

---

## 2. About / Core Expertise

### Purpose
To share Victor's engineering background and technical toolbelt in a straightforward, conversational tone that resonates with fellow software engineers.

### Content Specification

#### Executive Bio (Developer-to-Developer Tone)
> "I've been writing software for over 20 years, focusing mostly on Identity & Access Management (IAM) and hyperscale cloud infrastructure. Over the years at companies like **DocuSign**, **Microsoft**, and **Salesforce**, I've spent my time tackling deep system challenges.
>
> That work has ranged from building bare-metal authentication servers for Microsoft Azure (dSTS) and shipping FIDO2 passwordless WebAuthn to hundreds of millions of Microsoft Accounts, to designing Change Data Capture pipelines for 90M+ identities and building multi-agent AI platforms. I enjoy working at the intersection of complex protocols (OAuth 2.0, OIDC, SAML, WebAuthn), zero-trust security, and distributed systems architecture."

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
Proof of work. Detailed write-ups detailing real problems, architectural decisions, trade-offs, and concrete results.

### Content Specification (4 Featured Projects)

#### Case Study 1: AxoEdge AI Product Suite & Engineering Tools
* **Domain:** AI Agent Orchestration, Computer Vision Benchmarking, RAG & Vector Systems
* **The Context:** Co-founded AxoEdge with fellow Microsoft veterans to build developer platforms and tackle AI engineering challenges for logistics and real estate clients.
* **Architecture & Implementation Highlights:**
  - **Stanley (AI Orchestration Platform):** Built a developer platform that manages multi-agent CLIs (`claude`, `copilot`, `opencode`) with explicit task states (`Working`, `NeedsApproval`, `Done`), automatic Git worktree isolation per task to eliminate merge conflicts, and centralized project policy (`server.settings.json`).
  - **ProLogistiks (Pallet Counting CV vs LLM Study):** Benchmark study evaluating warehouse pallet counting. Ran parallel experiments comparing 7 open-source multimodal LLMs against YOLO object detection on labeled datasets. Discovered top LLMs had 40% count deviation at 10s/image, while YOLO achieved ~80% accuracy in milliseconds. Delivered a data-backed recommendation in under a month.
  - **11 Casitas (Real Estate RAG Bot):** Built a real-time conversational agent for Facebook Messenger and Instagram Direct using a PostgreSQL + pgvector RAG pipeline exposed to LLMs through a Model Context Protocol (MCP) tool layer.
* **The Impact:**
  - Standardized multi-agent development workflows, achieving an estimated 30% reduction in worktree conflicts.
  - Saved logistics client from investing in unreliable LLM vision architecture through benchmark proof.
  - Reduced real estate lead response time from hours to ~1 second, cutting estimated lead drop-off by 30%.

---

#### Case Study 2: MSFT Datacenter STS (dSTS) Data Layer Migration
* **Domain:** Bare-Metal Regional Identity Provider, Shadow Traffic Validation, Storage Migration
* **The Context:** dSTS is Microsoft's bare-metal, tier-0 regional identity provider for Azure, booting immediately after networking in every datacenter with zero service dependencies. The legacy codebase relied on an unsupported distributed-hashtable store, creating technical debt and blocking modern protocol support (OAuth/OIDC).
* **Engineering Approach:**
  - Conducted a codebase assessment and led a strategic pivot to adapt Entra ID's modern login stack over a bare-metal storage model.
  - Evaluated bare-metal database options and selected SQL Server + Webstore (Microsoft's sharding layer) based on operational ownership and dedicated DB management support.
  - Authored the migration design and rollout strategy, using gateway-level shadow traffic to run old and new data layers in parallel without production risk.
* **The Impact:**
  - Delivered a working prototype in Q1 2023 demonstrating equivalent login success rates (~99%+ parity under shadow traffic).
  - De-risked production cutover and unblocked Azure's tier-0 modernization roadmap toward OAuth/OIDC.

---

#### Case Study 3: MSFT ID4S dSTS SAML & WS-FED Token Validation SDK
* **Domain:** Security Libraries, WS-Federation / SAML, Cross-Org Platform Standards
* **The Context:** As Microsoft standardized identity across services, lower-tier Azure services needed a reliable way to validate dSTS-issued SAML/WS-FED tokens. The core Identity for Services (ID4S) team lacked resources for a dSTS-specific fix due to unique endpoints, certificate rotation strategies, and claim sets.
* **Engineering Approach:**
  - Led the design and implementation of a reusable token validation library integrated directly into ID4S shared packages.
  - Built validation logic supporting dSTS-specific WS-FED/WS-TRUST flows, custom metadata endpoints, automated certificate rotation, and delegation models.
  - Partnered with ID4S and dSTS engineering teams to ensure feature parity and zero service disruption.
  - Created integration guides, support SOPs, and documentation for smooth adoption across Azure teams.
* **The Impact:**
  - Integrated into ID4S shared libraries and deployed across Azure services to validate dSTS tokens securely across Microsoft datacenters.
  - Standardized service-to-service token validation and removed custom validation debt.

---

#### Case Study 4: MSFT Account (MSA) Pioneer FIDO2 / WebAuthn Passwordless Launch
* **Domain:** Passwordless Authentication, W3C Standards, Client & UX Integration
* **The Context:** Bringing public key cryptography (WebAuthn/FIDO2) to hundreds of millions of Microsoft Account users while the W3C specification was still an evolving working draft, tied to a strict Windows 10 release timeline.
* **Engineering Approach:**
  - Led browser and UX-layer integration of WebAuthn for Microsoft Account (MSA), making Microsoft the first major identity provider to ship FIDO2 passwordless auth at scale.
  - Coordinated work across 4 engineering groups (MSA, Microsoft Edge, Windows Core, W3C PM group).
  - Advocated for requiring user-verification-only authenticators to achieve true passwordless login instead of simple two-factor auth.
  - Fed implementation telemetry and practical feedback back into W3C working groups.
* **The Impact:**
  - Shipped passwordless authentication on schedule for Windows 10 to hundreds of millions of MSA users.
  - Helped shape international standards for WebAuthn and FIDO2 authentication.

---

## 4. Professional Experience / Career Highlights

### Purpose
Clean timeline of roles and technical scope across top engineering organizations.

### Content Specification

```
┌────────────────────────────────────────────────────────────────────────┐
│ DocuSign                                                  2024 – Present│
│ Principal Software Engineer (Identity Organization)                    │
│ • Technical lead across Identity Org defining core authn/authz roadmap.│
│ • Leading security modernization and passwordless platform evolution.  │
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
│ • Governed OAuth, SAML, SCIM, OIDC standards for 8M+ client integrations.│
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
> 📄 **Formal Resume:** [Download Resume PDF](https://victorh.dev/resume.pdf)

---

## 5. Writing & Thought Leadership

### Purpose
Technical write-ups and architectural deep-dives sharing practical engineering lessons with fellow developers.

### Content Specification (Technical Articles)

1. 📖 **"Benchmarking Multimodal LLMs vs. Dedicated Computer Vision in Production Logistics"**  
   *An empirical look at zero-shot multimodal LLMs vs. YOLO object detection for image-based inventory tracking.*

2. 📖 **"Architecting Task Orchestration and Git Worktree Isolation for Multi-Agent AI Development"**  
   *How Stanley replaces ad-hoc AI coding sessions with structured task states, worktree isolation, and multi-agent policy control.*

3. 📖 **"Bare-Metal Identity: Operating Azure's Tier-0 Datacenter STS at Hyperscale"**  
   *Lessons in fault tolerance, shadow-traffic validation, and zero-dependency identity architectures for root cloud trust anchors.*

4. 📖 **"Passwordless at Scale: Shipping FIDO2 WebAuthn to Hundreds of Millions of Users"**  
   *Navigating evolving W3C specs, browser integration layers, and user verification authenticators in consumer identity.*

5. 📖 **"RAG Infrastructure Design: Leveraging PostgreSQL, pgvector, and MCP for Real-Time Conversational Interfaces"**  
   *Structuring vector chunking pipelines, MCP tool exposure, and sub-second response latency for social messaging integration.*

---

## 6. Contact & Footer

### Purpose
Frictionless way for fellow developers, collaborators, and peers to reach out.

### Content Specification

#### Contact Section
- **Headline:** *"Let's Connect & Talk Engineering"*
- **Sub-headline:** *"Feel free to reach out if you want to chat about IAM, identity protocols, distributed systems, or AI dev tools."*

#### Interactive Elements
- **Direct Email:** `contact@victorh.dev`
- **Actions:**
  - **[ Copy Email Address ]** *(Toast notification: "Copied contact@victorh.dev!")*
  - **[ Send Email ]** *(Triggers mailto:contact@victorh.dev)*
- **Social Links:**
  - **LinkedIn:** [linkedin.com/in/victormhdz](https://www.linkedin.com/in/victormhdz)
  - **GitHub:** [github.com/victormh](https://github.com)
  - **Domain:** [victorh.dev](https://victorh.dev)

#### Footer
> `© 2026 Victor Hernandez Guzman · victorh.dev`  
> *Built with clean semantic HTML, vanilla CSS, and zero tracking scripts.*
