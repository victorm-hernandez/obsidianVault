## MCP and Identity

### Interesting articles
- https://medium.com/rossoctl-the-agentic-platform/security-in-and-around-mcp-part-1-oauth-in-mcp-3f15fed0dd6e
- https://medium.com/rossoctl-the-agentic-platform/security-in-and-around-mcp-part-2-mcp-in-deployment-65bdd0ba9dc6
- https://medium.com/rossoctl-the-agentic-platform/security-in-and-around-mcp-part-3-mcp-server-identity-10d6768d96c1

### Newsletter Article Ideas Seeded from the MCP Series   (IAM Practitioner)

1. "From Shadow AI to Governed Agents: Designing the Non-Human Identity (NHI) Lifecycle"  

- **The IAM Practitioner Angle:** Enterprises are panicking because business units are deploying AI agents that act like shadow users: storing hardcoded API keys, inheriting bloated human permissions, and bypassing governance reviews.  
    
- **What the article will cover:**  
    ⚬ How to apply traditional **identity lifecycle management (ILM)** principles (onboarding, lifecycle tracking, recertification, and offboarding) to autonomous agents.  
    ⚬ Defining organizational ownership: Who is accountable when an AI agent drifts in its behavior and accesses sensitive data?  
    ⚬ Building a taxonomy for machine identities vs. human delegates.  

2. "The Governance Crisis of Delegated Consent: When AI Acts on Behalf of a Human"  

- **The IAM Practitioner Angle:** Traditional OAuth consent screens assume a human user giving an app permission to read their calendar. When an agent acts as a proxy, making hundreds of autonomous decisions across enterprise SaaS tools, standard user consent models completely collapse.  
    
- **What the Article Covers:**  
    ⚬ The policy challenge of **time-bound, scope-limited delegation**: How do enterprises audit what an agent _is allowed_ to do versus what it _technically can_ do?  
    ⚬ Bridging the gap between CISO risk requirements and business units wanting autonomous productivity tools.  
    ⚬ Structuring access recertification campaigns for automated workflows.  
    
3. "Rethinking Enterprise Access Control: Moving from Static RBAC to Dynamic Context-Aware Policies for Autonomous Systems"  

- **The IAM Practitioner Angle:** Role-Based Access Control (RBAC) and Attribute-Based Access Control (ABAC) were designed for static employees sitting at desks. They break when applied to non-deterministic agents that reason dynamically.  
    
- **What the Article Covers:**  
    ⚬ How enterprise architects must rethink **Policy Decision Points (PDP) and Policy Enforcement Points (PEP)** to handle runtime risk signals.  
    ⚬ Translating compliance frameworks (like NIST or SOC 2) into governance guardrails for AI-driven tool execution.  
    ⚬ How to write enterprise access policies that gracefully handle edge cases where an agent encounters ambiguous data.
    
### Newsletter Article Ideas Seeded from the MCP Series   (Developer)
1. "Beyond the Client ID: Securing the MCP Tool-Chain with OAuth Token Exchange"  

- **The Problem (Inspired by Part 1 & 2):** When an LLM host acts as an MCP client and chains calls to multiple downstream MCP servers, developers are tempted to pass the user's raw access token everywhere or use long-lived service account secrets. Both options destroy the principle of least privilege and expand the blast radius.  
    
- **The Article Focus:** How to implement **RFC 8693 OAuth 2.0 Token Exchange** within agentic pipelines.  
    
- **Key Technical Takeaways for Readers:**  
    ⚬ Using `grant_type=urn:ietf:params:oauth:grant-type:token-exchange` to swap a user's initial bearer token for a tightly-scoped, downstream-specific audience (`aud`) token before an agent invokes an MCP tool.  
    ⚬ Ensuring every component in the loop performs independent token validation (signature, expiration, audience constraints).  
    

2. "The 'Rug Pull' Threat: Establishing Cryptographic Trust and Identity for MCP Servers"  

- **The Problem (Inspired by Part 3):** MCP makes it trivial to plug tools into an agentic host. However, what stops a compromised or malicious MCP server from dynamically altering its behavior or tool definitions ("rug pulling") after gaining widespread adoption? Coarse-grained trust (e.g., namespace-only checks) leaves massive blind spots.  
    
- **The Article Focus:** Designing robust **Workload Identity and Attestation schemas** for non-human entities (NHIs) in agent platforms.  
    
- **Key Technical Takeaways for Readers:**  
    ⚬ Evaluating what makes an identity attribute _attestable, immutable, and granular_ in cloud-native container runtimes.  
    ⚬ Moving away from static API keys toward short-lived, cryptographically signed workload identity documents (akin to SPIFFE/SPIRE patterns) so that an AI agent has strict mathematical guarantees about _which_ server is executing its tool calls.  
    

3. "Architecting the Agentic Perimeter: Least-Privilege Access Control for Non-Deterministic Workflows"  

- **The Problem (Synthesizing Parts 1–3):** Traditional IAM assumes a human user initiating a request with deterministic intent. In an agentic architecture, the _agent_ makes runtime decisions about which tools to call, turning every prompt injection into a potential privilege escalation vector if the ACS is static.  
    
- **The Article Focus:** Designing a **Runtime Access Control Service (ACS)** that evaluates context, delegation chains, and tool-call safety dynamically.  
    
- **Key Technical Takeaways for Readers:**  
    ⚬ How to map user-delegated claims through the agent boundary so that authorization checks happen _at the tool execution boundary_, not just at the initial chatbot login.  
    ⚬ Designing audit logs and observability hooks that trace an agent's reasoning-to-tool-call lifecycle for compliance and incident response.  
    

Recommended Editorial Angle for Your Voice  
Because your background is deeply rooted in systems engineering, concurrency, and identity architecture, you can ground these pieces in **practical reality**. Rather than treating MCP or agentic security as abstract AI policy, you can write from the perspective of _how we actually build the middleware, handle the context propagation, and structure the validation loops in code_ (e.g., handling Go-based micro-gateways that intercept these multi-hop payloads)

## Other ideas

| **Article Title Idea**                                    | **Angle & Focus**                                                                                                                                 | **Target Pain Point**                         |
| --------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------- | --------------------------------------------- |
| **"From Service Accounts to Autonomous Actors"**          | Definitional and architectural shift required to treat AI agents as first-class identities in the enterprise.                                     | Shadow AI and ungoverned API keys.            |
| **"Rewriting OAuth for the Age of Agents"**               | Analyzing how protocols like OAuth 2.0 and OIDC must evolve to handle delegation, prompt context, and token-scoping for non-deterministic actors. | Token abuse and over-privileged delegation.   |
| **"Building a Runtime ACS for Multi-Step LLM Workflows"** | A technical code-and-architecture look at how to intercept agent tool-calls and evaluate authorization policies mid-loop.                         | Lack of real-time control over agent actions. |
Why This Resonates with IDPro Members  
IAM practitioners are currently drowning in "agent-washing" by vendors who claim standard bot management solves AI security. Providing a rigorous, protocol-level breakdown of how to build secure, auditable, and context-aware Access Control Services for agentic workflows positions your newsletter as an indispensable engineering resource.