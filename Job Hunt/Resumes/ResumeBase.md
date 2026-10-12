# VICTOR HERNANDEZ GUZMAN 

victormh@outlook.com | 425-269-8940 | www.linkedin.com/in/victormhdz 

## PRINCIPAL SOFTWARE ENGINEER 

### _Expertise in Identity & Access Management (IAM) | Hyperscale Cloud Infrastructure_ 

Technical leader with 20+ years of experience, including 15+ years specializing in Distributed Systems and Identity and Access Management (IAM). Demonstrated track record of designing and operating tier-0, hyperscale identity providers at Microsoft, Salesforce, and DocuSign, and of extending that same IAM discipline (governance, policy enforcement, and controlled delegation) into agentic AI systems, including a self-built multi-agent orchestration platform.

Deep expertise across the full identity stack, user authentication, machine-to-machine (M2M) authorization, token exchange, and federation spanning OAuth, OpenID Connect, JWT, SAML, SCIM, and WebAuthN/FIDO2.

## PROFESSIONAL EXPERIENCE 

_DOCUSIGN | Principal Software Engineer_ 

December 2024 – Present 

- **Architectural Strategy:**  Serve as technical leader across the Identity organization, defining the long-term roadmap for core AuthN/AuthZ systems powering DocuSign's global scale (1B+ users across 180+ countries) and leading the transition to passwordless authentication; act as primary Architecture Owner, mentoring staff/senior engineers.

- **Cryptographic Modernization:** Owned architecture and hands-on implementation of an company-wide password-hashing migration to meet NIST SP 800-63B and FIPS 140-3 requirements, designing a two-stage hashing construction that upgraded both active and inactive credentials without requiring user login, and rolling it out as a governed, batched production migration monitored against live authentication success and latency metrics.

-  **Proactive Threat Defense**: Architected a credential-compromise detection system, based on a breach-data ingestion pipeline which concluded in forced remediation at the authentication boundary, directly satisfying SOC 2 Type II Common Criteria (CC6.1, CC6.5, CC7.1/CC7.2) and producing the audit-trail evidence for CPA auditors.

_MICROSOFT | Principal Software Engineer_ 

January 2022 – December 2024  (3 years) 

- **Hyperscale Infrastructure & Machine-to-Machine Identity:** Technical lead for the Secure Token Server (dSTS). Azure's bare-metal, tier-0, zero-service-dependency identity provider that boots immediately after the networking layer in every datacenter and serves as the root-of-trust for service-to-service authentication across 80 regions and 500+ datacenters, run as a distributed swarm of 240 clusters (3 per region, 5+ nodes each) totaling 1,200+ servers globally; while federating with Entra ID for employee authentication.

- **SDK Architecture & Open Source**: Led the end-to-end design of the internal WS-FED/SAML token validation SDK critical to Azure’s tier-0 services, while actively contributing to the foundational open-source IdentityModel Extensions for .NET backing the broader ecosystem.

- **Identity Provider Modernization**: Architected and delivered bare-metal capabilities for the public Microsoft Entra ID Identity Provider, modernizing the Azure Infrastructure legacy identity stack and unblocking the roadmap toward modern OAuth/OIDC protocols. This work included deploying a new sharded data layer backed by SQL Server, executing a cross-organization migration off an unsupported distributed-hashtable store, and implementing a shadow-traffic strategy running both stacks in parallel to validate login-success parity and completely de-risk the cutover.

_SALESFORCE | Staff Software Engineer_ 

February 2019 – February 2022 (3 years) 

- **Identity Consolidation**: Spearheaded the architecture for a centralized Global Identity Provider unifying identity across all of Salesforce's acquired clouds, and owned technical direction for AuthN/AuthZ standards (OAuth, SAML, SCIM, OpenID) supporting 8M+ integrated client applications.

- **Distributed Data Synchronization**: Designed a horizontally-scalable, priority-tiered CDC pipeline (message queues, circuit breakers, exponential backoff, multi-user, and single-user sync tiers). The architecture became the reference model sibling teams adopted company-wide.

_MICROSOFT | Senior Software Engineer_ 

March 2009 – May 2019 (10 years) 

- **Pioneering Passwordless Architecture (WebAuthN/FIDO2)**: Owned the first server-side implementation of WebAuthN for Microsoft Account, delivering the industry's first FIDO2 passwordless solution. Coordinated across four teams (MSA, Edge, Windows, and the W3C Working group) against a hard Windows deadline while the specification was still a working draft, successfully feeding implementation requirements back into the W3C standardization process.

- **Cross-Ecosystem Integration**: Led the project to integrate Microsoft Account with Windows Core, Xbox One, HoloLens, and Windows 10; engineered QR-code-based and push-notification-based authentication.

## INDEPENDENT PROJECTS

**Agentic Security Standards:** Contributing Member, OWASP GenAI Security Project. Agent Control Standard (ACS), an emerging open specification for AI-agent identity, authorization, and runtime observability.

**Agentic Development Framework:** Architected a cohesive, agent‑centric development environment that streamlines task orchestration, enables collaborative coding across distributed nodes, and guarantees high‑quality releases via automated review and testing.

**Multimodal LLMs Benchmark**: Computer Vision vs. Multimodal LLM Feasibility Study Benchmarked 7 open-source multimodal LLMs against YOLO-based object detection for automated inventory counting; delivered a data-backed go/no-go recommendation that prevented investment in an unreliable architecture.

**RAG-Driven Messenger Assistant for Property Leads**: Built an AI messaging assistant for Facebook Messenger and Instagram that let customers get instant answers to property questions. The assistant retrieved answers from a RAG knowledge base (PostgreSQL/pgvector) via an MCP tool connector, cutting response time from hours to ~1second.

## TECHNICAL SKILLS 

**Architectural Domains**: Identity & Access Management (IAM), Authentication, Authorization, Security Frameworks and SDKs, User Authentication, Machine-to-Machine (M2M) Authentication, Zero Trust.

**Protocols & Standards**: OAuth, OpenID Connect, WebAuthN/Passkeys, WS-Trust, WS-FED, SCIM, SAML, JWT, SDLC, XSS, CSRF, OpenTelemetry.

**Languages/Technologies**: C#, Java, C++, JavaScript, SQL, ASP.NET, React.JS, Knockout.JS, Webpack, HTML/CSS, SQL Server, PostgreSQL, CosmosDB, Azure, Service Mesh.

## EDUCATION 

B.S. Computer Science, Instituto Tecnológico y de Estudios Superiores de Occidente (ITESO) 
