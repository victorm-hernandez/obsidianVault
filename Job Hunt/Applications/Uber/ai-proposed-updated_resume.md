> **Note to Victor (delete before submitting):** This draft repositions your existing, real experience toward security architecture and hyperscale distributed systems — the parts of the JD your history actually supports. Your confirmed Java depth (main language at Salesforce) is now called out explicitly below, since it directly satisfies the JD's production-language requirement. It still does **not** claim bot detection, DDoS mitigation, AWS/GCP, Kubernetes, or Kafka/Flink experience — you confirmed none of that is real, professional experience, so it stays out.

---

# VICTOR HERNANDEZ GUZMAN

victormh@gmail.com | 425-269-8940 | www.linkedin.com/in/victormhdz | victorh.dev

## PRINCIPAL SOFTWARE ENGINEER — SECURITY ARCHITECTURE & DISTRIBUTED SYSTEMS

Technical leader with 20+ years of experience architecting and operating hyperscale, tier-0 distributed systems in both Java and C# — including an identity control plane spanning 500+ datacenters and 1,200+ servers globally, and a 90M-user Java data-synchronization pipeline built on message queues, circuit breakers, and exponential backoff. Deep security-architecture background spanning authentication, cryptographic migration, and threat-detection data pipelines: designed a breach-intelligence ingestion pipeline that automatically enforces remediation at the authentication boundary, and led the industry's first FIDO2/WebAuthn passwordless deployment — a hardware-backed, origin-bound credential built to defeat phishing and credential stuffing. Proven track record leading complex, cross-functional initiatives across up to four teams from inception to production, and mentoring staff/senior engineers as a de facto architecture owner.

## PROFESSIONAL EXPERIENCE

### DOCUSIGN | Principal Software Engineer
*December 2024 – Present*

- **Threat-Intelligence Pipeline & Automated Enforcement:** Architected a breach-data ingestion pipeline that continuously matches externally-sourced compromised credentials against internal password hashes (no plaintext ever stored or compared) and enforces automated remediation — a mandatory reset — at the authentication boundary on the user's next login attempt; built the event logging to double as audit-trail evidence for SOC 2 Type II (CC6.1, CC6.5, CC7.1/CC7.2).
- **Cryptographic Modernization & Controlled Production Rollout:** Owned architecture and hands-on implementation of a company-wide password-hashing migration to meet NIST SP 800-63B and FIPS 140-3, designing a two-stage hashing construction that upgraded both active and inactive credentials without requiring a user login, deployed via a batched, health-gated production rollout (auth success rate, latency, CPU, DB load) with pause/rollback control at every stage.
- **Architectural Strategy & Mentorship:** Serve as Architecture Owner across the Identity organization, defining the authn/authz roadmap and leading the passwordless transition; mentor staff/senior engineers and set cross-functional design standards.

### MICROSOFT | Principal Software Engineer
*January 2022 – December 2024 (3 years)*

- **Hyperscale, Zero-Dependency Distributed Infrastructure:** Technical lead for dSTS, Azure's bare-metal tier-0 identity provider and root of trust for service-to-service authentication — 80 regions, 500+ datacenters, 60+ countries, run as a distributed swarm of 240 clusters (1,200+ servers globally) with zero service dependencies.
- **Data-Layer Migration with Shadow-Traffic Validation:** Designed and led a cross-org migration off an unsupported distributed-hashtable store onto a new SQL-Server-backed sharded data layer, running legacy and new stacks in parallel via a shadow-traffic strategy to validate login-success parity before cutover — de-risking a tier-0 migration with zero tolerance for downtime.
- **SDK Architecture & Open Source:** Led design of an internal WS-FED/SAML token-validation SDK adopted by low-tier Azure services; contributed to the open-source IdentityModel Extensions for .NET.
- **Team Scaling Through Incident Pressure:** As tech lead through 5x team growth (5 → 30+ engineers) amid a chain of livesite incidents, ran bottom-up process diagnostics with the team and landed 8+ process changes (release manager rotation, automated release notes, documentation standards), improving incident turnaround and rebuilding team ownership.

### SALESFORCE | Lead Software Engineer
*February 2019 – February 2022 (3 years)*

- **Resilient, Priority-Tiered Distributed Pipeline (Java):** Designed and built, in Java, a horizontally-scalable change-data-capture pipeline for a new centralized Identity Provider — a Java Message Service (JMS)-based message queue, circuit breakers, exponential backoff, and a three-tier priority model (org-level, multi-user, single-user) — syncing 90M+ users at up to ~2,000 requests/second per instance with zero added latency to live user-facing flows; the architecture became the reference model adopted company-wide.
- **Cross-Team Technical Leadership:** Authored the technical spec parallelizing work across three cross-functional teams; coached a mostly-junior team through spec and PR review, and absorbed a mid-project engineer departure without schedule slip.

### MICROSOFT | Senior Software Engineer
*March 2009 – May 2019 (10 years)*

- **Industry-First Passwordless / Anti-Phishing Authentication:** Owned the first server-side implementation of WebAuthn for Microsoft Account, delivering the industry's first FIDO2 passwordless solution — a hardware-backed, origin-bound credential structurally resistant to phishing and credential stuffing. Coordinated across four teams (MSA, Edge, Windows, W3C working group) against a hard Windows 10 deadline while the spec was still a working draft, feeding implementation requirements back into the W3C standardization process.
- **Production Incident Response at Scale:** Diagnosed a silent regression that collapsed login success to under 20% of baseline on a flow processing tens of millions of logins/month, invisible to every log and pre-production environment; root-caused it via physical device testing, hot-patched via direct CDN cache flush, and restored baseline within ~3 hours — then shipped controller-emulation test automation to permanently close the coverage gap.
- **Cross-Ecosystem Identity Integration:** Integrated Microsoft Account across Windows Core, Xbox One, and HoloLens; engineered QR-code-based and push-notification-based authentication flows.

## INDEPENDENT PROJECTS

- **Agentic Security Standards:** Contributing Member, OWASP GenAI Security Project — Agent Control Standard (ACS), an emerging specification for AI-agent identity, authorization, and runtime observability.
- **Agentic Development Platform (Stanley):** Architected a multi-agent orchestration platform with task-state governance and isolated git worktrees, enabling safe, repeatable AI-assisted development across a distributed team.
- **Applied ML Feasibility Study:** Benchmarked 7 open-source multimodal LLMs against YOLO-based object detection for production inventory counting; delivered a data-backed go/no-go recommendation that prevented investment in an unreliable architecture.
- **RAG-Driven Messaging Assistant:** Built a retrieval-augmented assistant (PostgreSQL/pgvector, MCP tool layer) for Facebook Messenger/Instagram, cutting response time from hours to ~1 second.

## TECHNICAL SKILLS

**Security & Architecture:** Identity & Access Management (IAM), Authentication, Authorization, Threat-Detection & Remediation Pipelines, Zero Trust, Security Frameworks & SDKs

**Distributed Systems & Resilience Engineering:** Hyperscale distributed infrastructure, message queues, circuit breakers, exponential backoff, priority-tiered load shedding, shadow-traffic validation, CDC pipelines

**Protocols & Standards:** OAuth, OpenID Connect, WebAuthn/Passkeys/FIDO2, WS-Trust, WS-FED, SCIM, SAML, JWT, OpenTelemetry

**Languages/Technologies:** Java (production, Salesforce), C# (production, Microsoft/DocuSign), C++, JavaScript, SQL, ASP.NET, React.js, SQL Server, PostgreSQL, CosmosDB, Elasticsearch, Azure, Service Fabric, JMS, Service Mesh

**Currently Expanding Into:** Go (self-directed, personal projects)

## EDUCATION

B.S. Computer Science, Instituto Tecnológico y de Estudios Superiores de Occidente (ITESO)

---

> **Recruiter note (delete before submitting):** This draft still omits AWS/GCP/OCI and Kubernetes/Docker from Technical Skills — you confirmed you have no professional experience with either, only personal-project exposure to K8s, which isn't strong enough to list as a skill without an interviewer immediately testing it. Go is listed separately under "Currently Expanding Into" rather than in Skills proper, since one month of self-study isn't yet a claimable proficiency — but it's honest, low-risk signal that you're actively investing in this pivot.
