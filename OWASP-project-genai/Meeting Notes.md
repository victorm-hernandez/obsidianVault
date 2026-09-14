Rock Blog
https://www.rockcybermusings.com/

Interest link
https://docs.google.com/forms/d/e/1FAIpQLSd0pPoz5aPvgYrw4mCiafd-Do-BngEvl7Vk5f3a5_Psb68RWA/viewform?usp=dialog

Github
https://github.com/GenAI-Security-Project/agent-control-standard

Start here: https://github.com/GenAI-Security-Project/agent-control-standard/tree/integration/reference-implementations/agt

I am planning to focus my efforts on the following areas:

Specification: normative spec text and review
Identity: IETF proposals, agent identity, tokenization, authN and authZ for non-human subjects
Building: port the reference Guardian to Python, Go, Rust, or Codex


## All Open areas for contribution

Specification: JSON Schema and hook definitions
Specification: normative spec text and review
Specification: Trace mappings for OpenTelemetry and OCSF
Specification: AgBOM and the CycloneDX, SPDX, and SWID derivations
Building: port the reference Guardian to Python, Go, Rust, or Codex
Building: production hardening, batching, and OpenTelemetry collection
Building: adapters for other agent frameworks and harnesses
Building: SDK work
Coding Agents: the IDE coding-agent spec and hook coverage across harnesses
Identity: IETF proposals, agent identity, tokenization, authN and authZ for non-human subjects
Testing and Validation: conformance suite, negative vectors, behavioral tests, requirement ledger
Using it: run ACS against your own harness and file conformance reports
Using it: deploy it and tell us what living with it is like
Documentation: guides, concepts pages, tutorials, editing
Outreach: talks, writing, social, community, analyst and vendor conversations
Program: issue triage, project board, release coordination, meeting notes
Design: site, diagrams, presentation material

## Contacts
Lead
Rock.lambros@owasp.org

He coded the demo reference implementation 
https://www.linkedin.com/in/arielfogel/

### Main contributors
https://www.linkedin.com/in/russelltait/
https://www.linkedin.com/in/bar-kaduri/ (https://hooks.security/)
https://www.linkedin.com/in/evgeniykokuykin/

### Other attendees
https://www.linkedin.com/in/paulgrabow/
https://www.linkedin.com/in/charisa-orwig/

https://www.linkedin.com/in/ivanmelia/
https://www.linkedin.com/in/aruneeshsalhotra/
https://www.linkedin.com/in/janmast/
https://www.linkedin.com/in/ghigliottyc/
https://www.linkedin.com/in/taha-ansari/
https://www.linkedin.com/in/information-security-cyber-ai/
https://www.linkedin.com/in/vadim-bulavintsev-10957b167/
https://www.linkedin.com/in/danielastocker/
https://www.linkedin.com/in/nilasisb/
https://www.linkedin.com/in/vanessagerardo/
https://www.linkedin.com/in/altazvalani/
https://www.linkedin.com/in/anna-tischenko-a26b73335/
https://www.linkedin.com/in/karmvir-jadeja-a21374201/
## Competition 
This project by MSFT doesnt follow ACS to the T.
https://github.com/microsoft/agent-governance-toolkit

## Related Projects

https://genai.owasp.org/resource/owasp-aibom-generator/

## Resources
https://hooks.security/
Framework of the reference implementation https://bun.com/
Reference implementation: https://github.com/GenAI-Security-Project/agent-control-standard/tree/integration/reference-implementations/agt
## Concepts

**JSON RPC 2.0**
JSON-RPC 2.0 is a stateless, lightweight remote procedure call (RPC) protocol that uses JavaScript Object Notation (JSON) as its data format. It allows a client application to execute functions or methods on a remote server as if they were local functions, passing parameters and receiving results back over any transport medium like HTTP, WebSockets, or TCP sockets

**Rego**
**Rego** is an open-source, declarative policy language created by the Open Policy Agent (OPA) project. It is specifically designed to allow developers and security teams to write policy as code, separating authorization logic from application code across modern cloud-native environments.

**SIEM**
SIEM stands for Security Information and Event Management. It is a cybersecurity solution that acts as a centralized "brain" for an organization's digital security, helping teams detect, investigate, and respond to cyber threats before they can cause damage

### Hash-Chained SessionContext

A **hash-chained SessionContext** is a security and architectural pattern used in distributed systems, zero-trust frameworks, and advanced identity architectures to maintain the integrity, provenance, and unforgeability of a user's session state.

Instead of treating a session context as a static token or a mutable database record, a hash-chained SessionContext links every state transition or request cryptographically to the one that preceded it—forming a tamper-evident audit trail similar to a lightweight blockchain or an append-only log.

### Key Architectural Components

- **Cryptographic Hashing:** Each state update, operation, or token issuance contains a cryptographic hash (e.g., SHA-256) of the _previous_ session state payload, combined with the current operation's metadata.
    
- **Append-Only Event Stream:** The session doesn't just hold "current values" (like user ID and permissions); it maintains an ordered chain of state-change events. If any intermediate event is modified, deleted, or injected, the hash chain breaks instantly.
    
- **Token Binding / Cryptographic Proof:** Often combined with token-binding mechanisms (like DPoP—Demonstrating Proof-of-Possession or mTLS), the session context ties the cryptographic keys used by the client directly to the hash chain.
    

### How It Works (Conceptual Flow)

1. **Session Initialization ($E_0$):**
    
    When a user authenticates, the identity provider or gateway creates the initial session context ($E_0$), generating a genesis hash:
    
    $$\text{Hash}_0 = \text{HMAC}(\text{Secret}, \text{Payload}_0)$$
    
2. **State Transition / Request ($E_1$):**
    
    When the user performs a sensitive action or refreshes their token, a new event payload ($\text{Payload}_1$) is generated. This payload includes a reference to $\text{Hash}_0$:
    
    $$\text{Hash}_1 = \text{HMAC}(\text{Secret}, \text{Payload}_1 \parallel \text{Hash}_0)$$
    
3. **Verification:**
    
    Downstream microservices or authorization gateways can verify the entire session trajectory by walking backward or forward through the chain, ensuring no man-in-the-middle tampering, replay attacks, or unauthorized privilege escalations occurred mid-session.
    

### Security Benefits

- **Tamper Evidence:** If an attacker manages to compromise a Redis cache or database where session attributes are stored and attempts to elevate privileges (e.g., changing a `role: user` to `role: admin`), the subsequent hashes will fail validation.
    
- **Replay Protection:** Because each step incorporates the prior state's hash and often a nonce or timestamp, captured requests cannot be replayed out of sequence.
    
- **Enhanced Auditability & Non-Repudiation:** It creates a strict, mathematically verifiable forensic trail of what the session did and when it changed state, which is critical in high-security zero-trust architectures and SPIFFE/SPIRE-adjacent environments.