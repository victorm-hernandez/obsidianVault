As a side project I recently joined the OWASP organization, in particular the genAI Security project (https://genai.owasp.org/)

In specific the newly adopted Agent Control Standard (ACS) https://genai.owasp.org/resource/agent-control-standard-acs/, the github where the standard is worked out is here https://github.com/GenAI-Security-Project/agent-control-standard


Agent Control Standard (ACS) is an open-source specification for runtime control of AI agents. It defines the interaction between an Observed Agent and a Guardian Agent so enterprises can inspect, trace, and instrument agent behavior, regardless of where the agent runs or which framework built it. 

There is engineering work: porting the reference Guardian to Python, Go, and Rust, hardening it for production, and building adapters for other agent
frameworks. There is specification work on schemas, hook definitions, and the OpenTelemetry and OCSF mappings. And a good deal of what needs doing is not code at all, including documentation, testing, outreach, design, and keeping the issue tracker honest.

I am planning to focus my efforts on the following areas:

Specification: normative spec text and review
Identity: IETF proposals, agent identity, tokenization, authN and authZ for non-human subjects
Building: port the reference Guardian to Python, Go, Rust, or Codex

# Project description

ACS makes AI agents trustworthy by standardizing how they are observed and controlled at runtime. The spec defines a JSON-RPC 2.0 wire format between an "Observed Agent" and a "Guardian Agent," covering native hooks, a small set of dispositions (e.g., allow/deny/modify/ask/defer), and capability negotiation so deployments can adopt only the conformance profiles they need. It also covers observability (OpenTelemetry and OCSF event mapping) and a dynamic Agent Bill of Materials (AgBOM) via CycloneDX, SPDX, and SWID, so agent tools, models, and data access stay inspectable, traceable, and instrumentable. ACS is open-source under the CC BY-SA 4.0 license.

# 🎯 Goals

1. Deliver a reference implementation of the Guardian Agent pattern, benchmarked for interoperability against Microsoft's Agent Governance Toolkit, so the community has a concrete, runnable example of ACS-Core in practice.
2. Establish ACS-Core as the shared, vendor-neutral wire-format standard that any agent framework, gateway, or policy engine can implement to gain runtime observability and control, rather than competing proprietary approaches.
3. Make agents inspectable, traceable, and instrumentable by default, giving enterprises visibility into agent tools, models, and data access, and a way to trace any action back to the reasoning and task that produced it, across in-house, SaaS, cloud, on-prem, and endpoint deployments.
4. Build out Rego and Cedar policy packs for the Guardian Agent's deterministic policy layer, covering common use cases (e.g., data protection, cost control, tool-access allowlisting) and mapped to the OWASP Top 10 for LLM Applications and the OWASP Top 10 for Agentic Applications.
5. Grow an open, vendor-neutral contributor community and governance model, with a Core Team and workstream leads, so ACS evolves through broad industry consensus rather than any single vendor's roadmap.


