- What is one meaningful contribution you’ve made in a past work experience that is relevant to the role you’re applying for?
- Which of the preferred qualifications have you acquired through your past experiences?
- What unique skills, experiences and contributions will you bring to Robinhood that set you apart from other candidates? [This is your opportunity to tell us about your unique, individual experiences and how they will add value to Robinhood. Consider things like professional, educational, and volunteer experiences.]
- What job-relevant accomplishment are you most proud of? [Tell us about a job-specific achievement that you take immense pride in. Please include specific details.]
- Of the job requirements, in which do you have the least expertise?
- Why are you specifically interested in joining Robinhood?


## Answers

# # Robinhood — Senior Staff Software Security Engineer

## Application Questionnaire — Draft Answers

_Drafted from Victor Hernandez Guzman's resume and project files, mapped against the job posting. Review each answer, tighten as needed, and adjust anything that doesn't match how you'd say it out loud._

---

### 1. What is one meaningful contribution you've made in a past work experience that is relevant to the role you're applying for?

At DocuSign, I owned the architecture and hands-on implementation of two connected security initiatives inside our Identity organization: a company-wide password-hashing migration to meet NIST SP 800-63B and FIPS 140-3 requirements, and a proactive credential-compromise detection system built on breach-data ingestion.

For the hashing migration, I designed a two-stage construction that let me upgrade both active and inactive credentials without requiring a user to log in, then rolled it out as a governed, batched production migration monitored against live authentication success and latency metrics rather than as a one-time database change. That same credential-hashing discipline fed directly into the compromise-detection pipeline: we matched breach data against our password hashes (never storing or comparing plaintext), flagged compromised accounts, and enforced a mandatory reset at the authentication boundary. The audit trail I architected became the evidence our auditors used to close SOC 2 Type II Common Criteria CC6.1 and CC7.1/CC7.2.

I believe this is particularly relevant to this position because it required exactly the cross-domain thinking the role asks for: cryptography, IAM, detection and response, and compliance.

---

### 2. Which of the preferred qualifications have you acquired through your past experiences?

- **Organizational Scale & Influence.** As tech lead (in a non-managerial role) for the bare-metal identity provider for Azure Infrastructure at Microsoft, I drove process and architectural changes across a team that grew from 5 to 30+ engineers in under a year, using bottom-up facilitation rather than mandates. On the Microsoft Account team, I delivered the industry's first WebAuthn implementation, coordinating 3 cross-functional teams (MSA, Edge, and Windows) with no formal authority over any of them.
- **IAM and Cryptography expertise**. 15+ years across OAuth, OIDC, SAML, SCIM, WS-Fed/WS-Trust, and WebAuthN/FIDO2, plus contributions to the core token validation SDKs used by ASP.NET and the .NET ecosystem (including adding support for SAML and WS-Fed), and hands-on cryptographic engineering at DocuSign selecting and implementing a NIST/FIPS-aligned password-hashing scheme.
- **Builder Orientation.** Every role above was hands-on implementation, not just design. I still write the code and ship the systems I architect.
- **Navigating Ambiguity.** While modernizing the baremetal Azure Tier-0 Identity Provider at Microsoft, there was no clean answer for which data store should back Entra ID running at Tier 0 bare metal. I had to weigh my own team's engineering know-how (developers, not DBAs, so a store that demanded new operational expertise was a real cost), the technical feasibility of each option against Tier 0 availability and latency requirements, cost-effectiveness, and how well each option's owning stakeholders would align with and continue supporting my organization long-term. I landed on a SQL Server data store built on in-house sharding technology already proven by the Microsoft Account team, deliberately trading novelty for something my team could actually operate and stakeholders who had a real stake in keeping healthy.
- **Collaborative Leadership.** At Salesforce I coached a mostly-junior team through a 90M-user identity migration and absorbed a mid-project engineer departure without slipping the timeline; at DocuSign I mentor staff and senior engineers as Architecture Owner.
- **AI Security Depth** I'm a contributing member of the OWASP GenAI Security Project, working on the Agent Control Standard (ACS) and the Agentic Security Initiative (ASI) for AI-agent identity, authorization, and runtime observability. I've also built a governance-first orchestration platform for AI coding agents (task-state workflow, centralized policy, isolated execution per task). 

---

### 3. What unique skills, experiences and contributions will you bring to Robinhood that set you apart from other candidates?

Fifteen-plus years of IAM depth built inside three very different hyperscale and regulated environments: Microsoft's tier-0 baremetal identity infrastructure (500+ datacenters, 80 regions), Salesforce's 90-million-user identity consolidation across acquired clouds, and DocuSign's compliance-driven security work (SOC 2 Type II, NIST, FIPS). 

Most directly relevant to where Robinhood is headed: I'm not just aware of agentic AI security as a coming problem, I'm actively contributing to the OWASP standard defining how AI agents get identity, authorization, and observability. That combination,  IAM/crypto foundations plus current, hands-on work on agent identity, is the intersection this role's AI Security Foundations mandate is explicitly asking for.

---

### 4. What job-relevant accomplishment are you most proud of?

Leading the modernization of dSTS, Azure's Tier 0 identity provider and the root trust anchor for all Azure infrastructure service-to-service authentication. The team had scaled from 5 to 30 engineers, but the data plane still sat on retired Azure Access Control Service infrastructure with zero tribal knowledge left, since the original team was gone, and most modernization effort had gone into the control plane, leaving the data plane's legacy footprint, everything from the data store to CI/CD, unaddressed.

I formulated and led the proposal to move the data plane onto Entra ID running at true bare-metal Tier 0, which meant solving a hard technical blocker (Entra ID depended on a Tier 1 data store) and a harder organizational one: the stakeholders who owned the legacy identity provider had every reason to resist a plan that replaced their work. Rather than force the migration through, I built their teams specific leadership roles inside the new architecture, so they retained real ownership and scope instead of watching their system get replaced out from under them. On the technical side, I selected a SQL Server data store using in-house sharding technology already proven at Microsoft, and designed a multi-stage, zero-downtime migration (static migration, a dual data layer, and a continuous-consistency daemon), validated end-to-end with a new bare-metal Service Gateway running shadow traffic against the legacy stack.

We delivered a pilot deployment that proved the new data layer could meet Tier 0 requirements, decoupled the team from unsupported legacy infrastructure, and unblocked the roadmap to modern OAuth/OIDC across Azure. I'm proud of this one because the hard part wasn't the architecture, it was getting a fractured, rapidly growing team and skeptical stakeholders aligned behind a plan that asked several of them to give up ownership of something they'd built, without losing their trust or their expertise along the way.

---

### 5. Of the job requirements, in which do you have the least expertise?

Cloud security in the cloud-native-controls sense, and broad Detection & Response operations. My depth is mostly in IAM, cryptography and distributed systems: identity protocols, token issuance and validation, and the authentication boundary. I've operated at hyperscale infrastructure (Azure's tier-0 baremetal identity layer) and built one detection-and-remediation pipeline (DocuSign's credential-compromise system tied to SOC 2), but I haven't run a security only organization.

I'd expect to ramp fastest on the Cloud Security and Infrastructure Security domains, and I'd rather be direct about that than overstate breadth I don't have yet. What I'd bring instead is a track record of learning a new domain quickly under real deadlines, and enough architectural range to connect what I don't yet own operationally to the identity and cryptography foundations I do.

---

### 6. Why are you specifically interested in joining Robinhood?

Two things line up for me specifically. First, the mission: Robinhood's bet is now extending directly into agentic AI, with the Agentic Trading Platform putting real financial decision-making in the hands of AI agents. That's precisely the identity and authorization problem I've spent the last stretch of my career on: who is this agent, what is it allowed to do, and how do we prove it after the fact. I'm not coming to that problem cold; I'm a contributing member of the OWASP GenAI Security Project's Agent Control Standard, which exists to answer exactly that for AI agents industry-wide.

Second, the shape of the role itself: an autonomous, highly visible IC seat that has to connect Cloud, Infrastructure, Cryptography, IAM, AI Security, and Detection & Response into one coherent strategy, while still writing code and shipping systems, is the natural next step for someone like me whose career has alternated between deep technical ownership (tier-0 identity infrastructure, cryptographic migrations) and cross-team influence without formal authority. Robinhood is one of the few companies shipping agentic financial products at real consumer scale right now, which means the security architecture decisions made in this role won't be theoretical, they'll be protecting real customer money from day one.



https://robinhood.com/us/en/agentic-trading/
