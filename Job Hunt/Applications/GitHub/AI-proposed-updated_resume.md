> **Note to Victor — read before sending, then delete this block:**
> 1. **Kubernetes framing below is intentionally modest.** You said DocuSign's deployments ran on Kubernetes but you weren't hands-on with deployment ops — so the resume describes the *environment*, not personal K8s ownership. Don't let an interview question drift into overclaiming here.
> 2. **The OWASP GenAI ACS line now reads as brand-new**, since you joined about a week ago. Don't build interview answers around it as an established track record — it's a genuine, current signal of interest in the space, not a project yet.
> 3. **A few metrics below are still your own estimates, not measured baselines** (dSTS's "~30% incident-turnaround improvement," the "~99%+ login parity," Stanley's "~30% fewer worktree conflicts," MessageBot's "~30% fewer lost leads" and "~1 second" response time). Fine to keep if you can defend them, but firm up with real numbers if you have them — a Principal-level panel will probe.
> 4. Everything else (Salesforce's 90M users/2,000 req/s, Xbox's incident metrics, Remix 3D's p95 < 100ms, ProLogistiks' benchmark numbers, and the dSTS footprint numbers you gave me) is either a measured result from your source materials or a number you stated directly — not an estimate.

---

# Victor Hernandez Guzman
victormh@outlook.com | 425-269-8940 | www.linkedin.com/in/victormhdz

## Principal Software Engineer — Identity & Access Management | Agentic Systems | Hyperscale Cloud Infrastructure

Technical leader with 20+ years of experience, including 15+ years specializing in Identity and Access Management (IAM). Deep expertise across the full identity stack — user authentication, machine-to-machine (M2M) authorization, token exchange, and federation — spanning OAuth, OpenID Connect, JWT, SAML, SCIM, and WebAuthn/FIDO2. Track record of designing and operating tier-0, hyperscale identity providers at Microsoft, Salesforce, and DocuSign, and of extending that same IAM discipline — governance, policy enforcement, and controlled delegation — into agentic AI systems, including a contribution to the OWASP GenAI Security Project's Agent Control Standard (ACS) and a self-built multi-agent orchestration platform.

## Professional Experience

### DocuSign | Principal Software Engineer
*December 2024 – Present*

- **Cryptographic Modernization:** Owned architecture and hands-on implementation of a company-wide password-hashing migration to meet NIST SP 800-63B and FIPS 140-3, designing a two-stage hashing construction that upgraded both active and inactive credentials without requiring a user login, rolled out as a governed, batched production migration across a Kubernetes-based deployment environment, monitored against live authentication success and latency signals — with defined pause/rollback control points as a formal risk-mitigation framework.
- **Proactive Threat Defense:** Architected a credential-compromise detection system — a breach-data ingestion pipeline driving forced remediation at the authentication boundary — directly satisfying SOC 2 Type II Common Criteria (CC6.1, CC6.5, CC7.1/CC7.2) and producing the audit-trail evidence CPA auditors required.
- **Architectural Strategy & Governance:** Serve as technical leader across the Identity organization, defining the long-term roadmap for core authn/authz systems and leading the transition to passwordless authentication; act as primary Architecture Owner, mentoring staff/senior engineers and establishing cross-functional design standards.

### Microsoft | Principal Software Engineer
*January 2022 – December 2024 (3 years)*

- **Tier-0 Identity Leadership at Global Scale:** Owned the Datacenter Secure Token Server (dSTS) — Azure's bare-metal, tier-0, zero-service-dependency identity provider that boots immediately after the networking layer in every datacenter and serves as the root-of-trust for service-to-service authentication across 80 regions and 500+ datacenters in 60+ countries, run as a distributed swarm of 240 clusters (3 per region, 5+ nodes each) totaling 1,200+ servers globally — while federating with Azure AD for employee authentication.
- **Data-Layer Modernization:** Designed a cross-org migration off an unsupported distributed-hashtable store, pivoting dSTS onto Entra ID's modern login stack backed by SQL Server + Webstore; de-risked the cutover with a shadow-traffic strategy running both stacks in parallel, validating ~99%+ login-success parity and unblocking the roadmap toward modern OAuth/OIDC protocols.
- **Open-Source SDK Architecture:** Led end-to-end design of the internal WS-FED/SAML token-validation library adopted by all tier-0 Azure services, while contributing to the open-source IdentityModel Extensions for .NET underpinning it.
- **Team & Process Scale-Up:** As tech lead, grew the team from 5 to 30+ engineers in under a year; ran bottom-up sessions to surface pain points and facilitated the team in landing 8+ process changes (documentation standards, release-manager rotation, automated release notes), improving incident turnaround by an estimated 30%.

### Salesforce | Lead Software Engineer
*February 2019 – February 2022 (3 years)*

- **Identity Consolidation:** Spearheaded the architecture for a centralized Global Identity Provider unifying identity across all of Salesforce's acquired clouds, and owned technical direction for authn/authz standards (OAuth, SAML, SCIM, OpenID) supporting 8M+ integrated client applications.
- **Large-Scale Data Synchronization:** Designed a horizontally-scalable, priority-tiered CDC pipeline (message queues, circuit breakers, exponential backoff across org-level, multi-user, and single-user sync tiers) that synced 90M+ users at up to ~2,000 requests/second per instance across 720 dequeue threads; the architecture became the reference model sibling teams adopted company-wide.
- **Technical & People Leadership:** Authored the spec that parallelized delivery across 3 cross-functional teams; coached a mostly-junior team through spec and PR review, absorbing a mid-project engineer departure without schedule slip.

### Microsoft | Senior Software Engineer
*March 2009 – May 2019 (10 years)*

- **Pioneering Passwordless Authentication (WebAuthn/FIDO2):** Owned the browser/UX-layer integration of WebAuthn for Microsoft Account — the industry's first FIDO2 passwordless implementation shipped at consumer scale (hundreds of millions of users) — coordinating across 4 teams (MSA, Edge, Windows, W3C PM group) against a hard Windows 10 deadline while the spec was still a working draft, and feeding implementation requirements back into the W3C standardization process.
- **Production Incident Response:** Diagnosed a silent Xbox One/Microsoft Account login regression (>80% success-rate drop across a flow processing tens of millions of logins/month, invisible to every log and pre-production environment) by reproducing it on physical console hardware; hot-patched and CDN-deployed the fix within ~3 hours of detection, then shipped controller-emulation automation to permanently close the testing gap.
- **Solo Compliance-Driven Rebuild:** As sole engineer on Remix 3D, rebuilt a legacy search service (Service Fabric + Elasticsearch, custom query-transformation layer) from scratch in 3 months to meet a hard GDPR enforcement deadline, achieving p95 < 100ms and clearing the team's last compliance blocker.
- **Cross-Ecosystem Integration & Security Innovation:** Led identity integration across Windows Core, Xbox One, HoloLens, and Windows 10; engineered QR-code-based and push-notification-based authentication; architected a scalable "flighting" framework for A/B testing authentication flows.

## Independent Projects

- **Agentic Security Standards — OWASP GenAI Security Project (New Contributing Member):** Recently joined the working group developing the Agent Control Standard (ACS), an emerging open specification for AI-agent identity, authorization, and runtime observability.
- **Stanley — Agentic Development Governance Platform (Co-Founder / Architect, AxoEdge):** Architected a governance-first platform unifying multiple coding-agent CLIs (Claude, Copilot, OpenCode) under a single task-state authorization model (`Working → NeedsApproval → Stashed → Done/Aborted`), an agent-agnostic policy layer, and isolated git-worktree execution — applying IAM governance principles (least-privilege tool allow-lists, centralized policy, auditable state transitions) directly to agentic software development; reduced worktree merge conflicts by an estimated 30%.
- **Computer Vision vs. Multimodal LLM Feasibility Study:** Benchmarked 7 open-source multimodal LLMs against YOLO-based object detection for automated inventory counting; the top LLMs showed ~40% count deviation at ~10 seconds/image versus ~80% accuracy for YOLO, and delivered a data-backed go/no-go recommendation that prevented investment in an unreliable architecture.
- **AI-Powered Lead Engagement Flow (RAG):** Built a RAG-based messaging assistant (PostgreSQL/pgvector, exposed to the LLM via an MCP tool connector) for Facebook Messenger and Instagram, cutting response time from hours to an estimated ~1 second and reducing an estimated 30% of lost leads.

## Technical Skills

**Identity & Security:** IAM, Authentication, Authorization, Federation, Zero Trust, M2M Authentication, Agentic Identity & Policy Enforcement, Security Frameworks & SDKs, NIST SP 800-63B, FIPS 140-3, SOC 2 Type II, GDPR

**Protocols & Standards:** OAuth, OpenID Connect, WebAuthn/Passkeys, WS-FED, WS-TRUST, SAML, SCIM, JWT, OpenTelemetry, SDLC, XSS, CSRF

**Distributed Systems & Architecture:** System Design, Distributed Systems, Cloud-Native Architecture, High Availability, Disaster Recovery (break-glass), Service Mesh, Service Fabric, Kubernetes (deployment environments), Elasticsearch/Lucene, Message Queues, Circuit Breakers, CDC Pipelines, Observability & Telemetry (Logging, Tracing, Profiling)

**Agentic AI / LLM:** Retrieval-Augmented Generation (RAG), Model Context Protocol (MCP), Semantic Kernel, Multi-agent orchestration (Claude, Copilot, OpenCode), LLM evaluation & benchmarking, ML.NET, ONNX Runtime, YOLO-based object detection, Ollama

**Languages / Technologies:** C#, Java, C++, JavaScript, TypeScript, Python, SQL, ASP.NET, React.js, Node.js, Knockout.js, HTML/CSS, SQL Server, PostgreSQL/pgvector, CosmosDB, Azure, Entra ID

## Education
B.S. Computer Science, Instituto Tecnológico y de Estudios Superiores de Occidente (ITESO)
