**

# 5.1 Level criteria L1-L5 with measurable evidence per level (Boris)

Assessment principle. Levels are cumulative. Maturity is assessed against controls applicable to the system’s authority model, operational scope, and protected effects. A control that is genuinely not applicable does not by itself cap the maturity level, but the basis for non-applicability shall be documented and reviewable.

L1 - Authority Identified

The system identifies who or what is acting, on whose behalf, which protected resources and consequential operations are involved, where authorization decisions are made, and where those decisions are enforced. Operations designated as critical or high-impact are determined by organizational policy. Agent requests, recommendations, prompts, or model-generated judgments are not treated as authorization by themselves.

Measurable evidence: A sampled protected operation identifies the principal, acting agent or workload identity, requested operation, protected resource, authorization source or decision point, and enforcement boundary, with records that allow these elements to be correlated.

L2 - Authority Bounded

Authority is explicitly bounded by the applicable task or purpose, recipient, permitted operations, target resources or data classes, environment, validity period, transaction or rate limits, and redelegation rights. Delegation is treated as bounded derivation rather than permission inheritance. Derived authority shall not exceed the applicable upstream authority, delegation ceiling, task scope, recipient boundary, contextual constraints, or target-resource policy without new explicit authorization.

Measurable evidence: A sampled delegation traces the parent or upstream authorization to the derived grant and demonstrates that scope and lifetime remain within the applicable inputs. Evidence includes relevant task, resource, operation, expiry, limit, approval, revocation, and redelegation records. Negative tests demonstrate denial of out-of-scope access or operations.

L3 - Authority Runtime-Enforced

Possession of a valid authorization artifact does not by itself establish permission to execute. Authorization is evaluated and enforced at the protected tool or resource boundary using the current grant and applicable runtime conditions. Materially equivalent execution paths shall apply equivalent authorization restrictions. The agent or model cannot independently approve, widen, or bypass the decision.

Measurable evidence: Evidence correlates the authorization or grant, runtime request, policy decision, and protected-resource result. Tests demonstrate rejection of wrong-resource, wrong-operation, expired, replayed, or improperly presented authority where applicable, including alternate execution paths. For denied consequential operations, evidence includes a resource-side deny record and confirmation that no protected side effect occurred.

 L4-Authority Lifecycle- and Consequence-Controlled

Authority remains bounded throughout delegation, use, expiry, revocation, task completion, retry, and recovery. In multi-hop delegation, each hop creates a bounded derived grant that retains or narrows applicable parent constraints; child authority shall not outlive or widen the authority from which it was derived. Critical or irreversible operations use additional controls such as purpose-bound just-in-time authority, independent approval, limited use, and replay prevention where required by organizational policy.

Measurable evidence: Evidence demonstrates applicable lifecycle transitions and parent/child relationships, including scope, audience, expiry, revocation, and use state. Negative tests include widened child authority, child expiry beyond parent expiry, replay after expiry or use, descendant use after parent revocation, and critical operations without required approval.

L5 -Authority Continuously Verifiable Across Applicable Boundaries

Authority, its derivation, and its current validity are machine-verifiable across all applicable trust, organizational, and execution boundaries. The system can establish not only that an authorization artifact is valid, but that the presenting agent is entitled to use it for the actual requested operation under current policy and runtime conditions. Upstream or external authority remains a ceiling, not an automatic permission; the protected resource or receiving domain independently applies its own authorization policy.

Where cross-domain authority is part of the architecture, the receiving organization verifies the issuer, delegation chain, presenting agent, audience, proof binding, current status, and applicable constraints before execution. Where no such boundary exists, cross-domain controls may be documented as not applicable.

Measurable evidence: Evidence supports reconstruction of the authority path from originating authorization through delegation or derivation, presenting agent, runtime decision, protected operation, and resulting state. Where applicable, tests demonstrate rejection of untrusted issuers, broken delegation chains, widened authority, wrong audience or presenter, stale or revoked authority, replay, altered approved operations, and alternate-path execution.

  
  

# 5.2 Permission boundaries, downscoping, and data-sensitivity tiers (Lawrence)

Delegation is a bounded derivation, not permission inheritance. An agent’s permission boundary is the maximum authority it may ever receive: resources, operations, environments, data classes, transaction limits, and delegation rights, not an operational grant.

  

A delegated agent shall receive an explicitly downscoped, task-specific grant rather than the parent’s full authority. The grant shall identify the recipient, permitted operations, target resources or data classes, task or purpose, validity period, transaction limits, and redelegation rights. Each hop shall retain or narrow those constraints; no hop may expand authority, audience, duration, data access, limits, or redelegation rights without new explicit authorization. Equal scope may be retained only where necessary for the same bounded task.

  

Authorization shall be enforced independently at every protected tool or resource boundary. The agent may request an action but shall not approve, widen, or bypass it. Equivalent restrictions must apply to all paths that can produce the same effect, including direct APIs, other tools, retries, downstream agents, and administrative interfaces. Prompts and model-generated judgment are not authorization controls.

  

Data sensitivity must further narrow access. Public or synthetic data may permit bounded read-only access; internal data requires an identified agent, task-specific scope, and short expiry. Personal, confidential, financial, or production data requires narrow object-level scope, separated read and write authority, enhanced evidence, and approval where policy requires. Irreversible financial, legal, identity, or administrative actions require single-purpose authority, explicit transaction and rate limits, rapid revocation, no redelegation by default, and independent approval under Section 5.6. The provisional control-tier table appears in the appendix.

  

A designated policy authority shall map these tiers to enforceable access rules. All paths to the same protected resource must apply equivalent restrictions.

  

Assessment evidence

Assessors shall trace a sample delegation from parent authorization through the derived grant to the resource-side decision. Evidence must show that the grant did not exceed the applicable authority, task, recipient boundary, contextual constraints, or resource policy, and must include relevant approval, expiry, revocation, and redelegation records.

  

Negative tests must demonstrate denial of out-of-scope access, prohibited operations, limit overruns, unauthorized redelegation, alternate-tool use, and indirect-path bypasses. For denied consequential operations, evidence must include a resource-side deny record and confirmation that no protected side effect occurred.

  
  

# 5.3 Token exchange and per-request scoping patterns (Alessandro/Victor)

An agent acting on a user's session token inherits all of the user's reach for the life of the session. Token exchange [1] replaces that. A trusted Security Token Service (STS) takes the user's token as subject_token and the agent's workload identity (Pillar 1) as actor_token, and returns a token narrower in scope, audience and lifetime. The user's token never enters the agent runtime. The actor credential must be the agent's own workload identity, not a shared client secret, or the exchange simply turns a static secret into a legitimate-looking token. The result must express delegation, not impersonation: sub is the user and act is the agent, so every decision and log entry separates who asked from who acted (chaining is covered in 5.5).

Per-request scoping goes further and issues one token per task or tool call. Resource Indicators [2] restrict the token to a single resource server, so a token minted for the calendar API is rejected by the mail API. Rich Authorization Requests [3] state the exact operation and its limits, such as one payment of at most €500 to a named beneficiary, which scopes cannot express. 

Within a domain, Transaction Tokens [4] handle what happens after the first hop. When a request enters the domain, an internal token service issues a short-lived signed token that records the subject, the purpose and the request context. Every internal service passes it along unchanged and cannot alter it, so downstream services authorize against what the user asked rather than the caller's ambient privilege.

Evidence. Decoded tokens show sub, act, aud and exp. A token replayed to another resource server is rejected. No user or refresh tokens are found in agent config, logs or context. Test enforcement at the resource server, not STS configuration. Level criteria are defined in 5.1.

  

# 5.4 Sender-constrained tokens (Alessandro/Victor)

Bearer tokens work for whoever holds them. In agents they pass through context, tool outputs and logs, where prompt injection or a verbose trace can exfiltrate them. Sender constraint binds the token to a key held by the legitimate agent, so a copied token alone is useless [7]. With mTLS [5], the token is bound to the client certificate thumbprint, which fits environments where agents already hold X.509 workload certificates, such as SPIFFE X.509-SVIDs, but breaks when TLS terminates at an intermediary. With DPoP [6], each request carries a signed proof over method, URI, time and the token hash. This works through proxies, but resource servers must check ath and reject replayed proofs.

|   |   |   |
|---|---|---|
|Attribute|mTLS [5]|DPoP [6]|
|Layer|Transport (TLS handshake)|Application (signed JWT in an HTTP header)|
|Key material|X.509 certificate, e.g. an X.509-SVID|Client-generated key pair|
|Proof of possession|TLS handshake proves the key, the resource server matches the certificate to cnf|Fresh signed proof on very request|
|Through gateways|Requires end-to-end TLS or a trusted gateway forwarding the certificate|Works through proxies|
|Best fit|Internal meshes with workload certificates|Traffic crossing gateways, agents without certificates|

The agentic requirement is that the key sits outside the model's reach. The runtime or a sidecar holds it, preferably non-exportable in hardware or a KMS, and signs outside the LLM process. A key in an environment variable that a model-invocable tool can read gives almost no protection. Refresh tokens are bound as well, or an attacker steals those instead. This stops token theft and replay, not runtime compromise, which is Pillar 1's job.

Evidence. cnf is present in sampled tokens. A token without a proof, or with a foreign or replayed proof, is rejected. The key is not readable through model-invocable tools. Level criteria are defined in 5.1.

Gap. No finalized standard lets a resource server verify that the key behind a DPoP-bound token is non-exportable and bound to an attested workload. Attestation-based client authentication [9] and WIMSE [8] address adjacent parts of the problem.

  

# 5.5 Multi-hop delegation: narrowing authority instead of impersonating the principal (Saba)

  

 (Draft) 

 When an agent hands part of a task to another agent, forwarding its own token or the principal's full authority turns delegation into impersonation, producing the delegation dilution and confused deputy failures in 2.2.

Criterion.Each hop is a new token exchange [1]: the delegator's grant is the subject_token and the recipient's workload identity is the actor_token. The derived grant retains or narrows the delegator's current authority (5.2), is audience-bound [2] and sender-constrained (5.4), references its parent grant, carries the remaining delegation depth, and never outlives its parent. Revoking a parent disables all descendants within a defined interval. The principal stays the subject and the presenting agent the actor at every hop. Cross-domain chains follow 5.7; chain reconstruction follows 6.4.

Evidence. A sampled three-hop chain showing subject, actor, parent reference, scope, audience, expiry and remaining depth at each hop, plus the L4 negative tests in 5.1.

Gap. RFC 8693 treats nested act claims as informational only and not for access control [1], and token revocation [12] does not reach grants derived by other parties, so chain verification and cascading revocation are left to the implementation.
 
  
  
  

# 5.6 Purpose-bound just-in-time authority and rule-of-two controls (Savio)

(Draft)

  

Authority is issued only for a declared task and only while that task runs. A critical operation additionally requires approval from two distinct authorized parties. 

  

Definitions

- Declared task: the task or trigger an agent's authority was issued for, set by the issuer, not by the agent. 
    
- Critical operation: an operation the organization's policy designates as high-impact or irreversible, such as moving funds, deleting data, changing privileges, or releasing data externally. The organization maintains this list. 
    

  

Criterion.

1. Task binding. The grant's permitted actions and resources are derived from the declared task, and the protected resource enforces them. The task is recorded in the grant, for example as an authorization_details entry [3]. The issuer sets the task from the request or trigger, not from the agent's own description.
    
2. Just-in-time issuance. No standing grant covers a critical operation. Authority is issued at task start, for example via token exchange [1], and ends at completion or expiry, whichever comes first. Completion triggers revocation without waiting for expiry. The grant has ended only when the protected resource denies it, not when revocation is issued. Short lifetimes bound the interval where revocation has not yet propagated
    
3. Rule-of-two. A critical operation executes only after approval by two authorized parties, neither of which is an agent in the requesting delegation chain. The approver sees the actual action and its consequential parameters, not an agent-written summary. Each approval is bound to that action, expires, and is single-use. This follows dual authorization practice [10]. Transaction and rate limits are separate controls and do not satisfy this requirement.
    

  

Evidence. A register of critical operations. A count of standing grants covering them (expected: zero). Grants showing declared task, derived scope and expiry. Revocation records at task completion showing both the time of the revocation event and the time the protected resource first denied the grant; the interval is within the organization's defined limit. Approval records linking both approvers to the exact action and parameters, with expiry and single-use consumption. Negative tests at the resource server: an out-of-task call is denied; a critical operation without both approvals is denied; an agent cannot approve its own action or approve through a delegate; a replayed, expired or altered approval is denied. For each denied critical operation, the test records the resource-side deny and a state check showing the protected resource is unchanged. Level criteria are defined in 5.1.

  

Gap. RFC 9396 defines no field or shared vocabulary for tasks; type values are specific to each API and authorization server, so task claims are not interoperable across issuers. For approvals, CIBA [11] provides an out-of-band channel, but its core specification defines no structured transaction details, its binding_message is a short interlock code, and the requesting client supplies the hint that selects the approver. Binding an approval to the full action parameters and guaranteeing approver independence are left to the implementation.

  
  

5.7 Cross-domain, machine-verifiable authority (Vladislav)

When authority crosses an organizational boundary, the receiving organization verifies it without relying on what the presenting agent asserts. Without such a boundary, 5.7 is documented as not applicable (5.1); 5.2–5.6 apply at every hop.

Criterion. A named policy owner approves trusted issuers, accepted proof formats and constraint meanings, required runtime conditions, and limits for evidence freshness and revocation propagation. Before execution the resource server verifies signatures against approved trust anchors and confirms the issuer may assert the claimed identities and permissions; discovering an issuer's keys does not establish trust [14]. The principal and acting agent stay separately identifiable within their issuing namespaces and are mapped to local identities; any principal change is rejected [1]. The resource server validates the chain to an accepted root, including each delegator's right to delegate, and rejects constraints it cannot interpret or enforce [1][3]. It checks consent or organizational mandate, approvals (5.6), revocation status, and attestation tied to the presenting agent's key against the approved runtime conditions [15], and denies when any of these is missing or older than the freshness limit. Replay [6] and use limits hold across all receiving instances and retries. The decision is revalidated when request details change or decision validity expires before execution. A valid external grant is only a ceiling; the resource server still applies its own current policy to the actual operation [17].

Evidence. A versioned trust configuration with freshness and propagation limits, identity mappings, delegation proofs, and decision records with trust and policy versions. Tests: successful authorized access; the L3 deny and state-check evidence and the L5 tests in 5.1; and rejection of a principal change, an issuer asserting outside its namespace, an unknown constraint, a stale or unavailable status source, and revoked descendants within the propagation limit. Level criteria are defined in 5.1.

Gap. No standard verifies delegation chains across issuers [1], propagates revocation across issuers [12][16], binds attestation to the agent's key (5.4), or lets DPoP bind request bodies [6].

  

# 5.8  map to owasp threat model (Savio)

  

Pillar 2 controls reduce the identity, delegation and authorization risks in the OWASP Top 10 for Agentic Applications 2026 [13]. The table below shows which risks each control addresses and what remains. Pillar 2 does not stop an agent from being hijacked or poisoned. It limits what a compromised agent can obtain, do and pass on. Risks that originate in the supply chain, code execution or memory (ASI04, ASI05, ASI06) are outside this pillar.

  

Evidence. For each risk in scope, the organization maintains a register entry naming the controls applied, the negative tests that demonstrate them, the controls relied on from other pillars, and the residual risk with an accountable owner.

  

Annex table. M# is the item number in that entry's Prevention and Mitigation Guidelines in [13].

  
  

|   |   |   |   |
|---|---|---|---|
|Risk|Pillar 2 controls (OWASP mitigation|Residual risk|Controls owned elsewhere|
|ASI03 Identity & Privilege Abuse|Downscoped task grants, token exchange with subject/actor, task-bound JIT expiry, sender-constrained tokens, approval for critical operations (M1, M3, M4, M5, M7)|A forged or compromised agent identity still receives the grants issued to it|Agent identity attestation (Pillar 1)|
|ASI02 Tool Misuse|Per-tool scoped grants, ephemeral credentials, approval for destructive actions (M1, M2, M6)|A misled agent misuses tools within its granted scope|Sandboxing, argument validation (outside AIMM)|
|ASI09 Human-Agent Trust|Rule-of-two; approver sees the actual action; approval bound to it (M1, M4, M7)|An approver approves a harmful action despite seeing it, through over-reliance or fatigue|Approval review UI and training (outside AIMM); approval traceability (Pillar 3)|
|ASI08 Cascading Failures|Per-hop narrowing, task expiry, revocation of descendants (M2, M3)|A compromised agent causes fan-out within its valid grants before revocation takes effect|Rate limits, blast-radius caps (Pillar 4)|
|ASI07 Inter-Agent Comms|Replay rejection, audience binding, cross-domain verification (M3, M4, M9|Forged or tampered inter-agent messages that carry no token claims|Message signing, semantic validation (outside AIMM)|
|ASI10 Rogue Agents|Revocation, per-run ephemeral credentials (M4, M6)|An agent drifts while holding valid, unexpired authority|Behavioral attestation, drift detection (Pillar 3)|
|ASI01 Goal Hijack|Impact limits only: least privilege, approval, task binding (M2, M5)|A hijacked agent acts within the authority it legitimately holds|Input validation, goal-lock controls (outside AIMM)|

  
  

# References  
[1] M. Jones, A. Nadalin, B. Campbell, J. Bradley, C. Mortimore, "OAuth 2.0 Token Exchange," RFC 8693, IETF, January 2020.

# [2] B. Campbell, J. Bradley, H. Tschofenig, "Resource Indicators for OAuth 2.0," RFC 8707, IETF, February 2020.

# [3] T. Lodderstedt, J. Richer, B. Campbell, "OAuth 2.0 Rich Authorization Requests," RFC 9396, IETF, May 2023.

# [4] A. Tulshibagwale, G. Fletcher, P. Kasselman, "Transaction Tokens," draft-ietf-oauth-transaction-tokens, IETF OAuth Working Group, work in progress.

[5] B. Campbell, J. Bradley, N. Sakimura, T. Lodderstedt, "OAuth 2.0 Mutual-TLS Client Authentication and Certificate-Bound Access Tokens," RFC 8705, IETF, February 2020.

[6] D. Fett, B. Campbell, J. Bradley, T. Lodderstedt, M. Jones, D. Waite, "OAuth 2.0 Demonstrating Proof of Possession (DPoP)," RFC 9449, IETF, September 2023.

[7] T. Lodderstedt, J. Bradley, A. Labunets, D. Fett, "Best Current Practice for OAuth 2.0 Security," RFC 9700, IETF, January 2025.

[8] B. Campbell, J. Salowey, A. Schwenkschuster, Y. Sheffer, "WIMSE Workload-to-Workload Authentication," draft-ietf-wimse-s2s-protocol, IETF WIMSE Working Group, work in progress.

[9] T. Looker, P. Bastian, C. Bormann, "OAuth 2.0 Attestation-Based Client Authentication," draft-ietf-oauth-attestation-based-client-auth, IETF OAuth Working Group, work in progress.

[10] NIST, "Security and Privacy Controls for Information Systems and Organizations," SP 800-53 Rev. 5, control AC-3(2), "Dual Authorization.”

  

[11] G. Fernandez, F. Walter, A. Nennker, D. Tonge, B. Campbell, "OpenID Connect Client-Initiated Backchannel Authentication Flow – Core 1.0," OpenID Foundation, September 2021.

  

[12] T. Lodderstedt, S. Dronia, M. Scurtescu, "OAuth 2.0 Token Revocation," RFC 7009, IETF, August 2013.

[https://www.rfc-editor.org/rfc/rfc7009](https://www.rfc-editor.org/rfc/rfc7009) 

  

[13] OWASP GenAI Security Project, "OWASP Top 10 for Agentic Applications 2026," December 2025, CC BY-SA 4.0.

  

[14] M. Jones, N. Sakimura, J. Bradley, "OAuth 2.0 Authorization Server Metadata," RFC 8414, IETF, June 2018.

  

[15] H. Birkholz, D. Thaler, M. Richardson, N. Smith, W. Pan, "Remote ATtestation procedureS (RATS) Architecture," RFC 9334, IETF, January 2023.

  

[16] J. Richer, "OAuth 2.0 Token Introspection," RFC 7662, IETF, October 2015.

  

[17] S. Rose, O. Borchert, S. Mitchell, S. Connelly, "Zero Trust Architecture," NIST SP 800-207, August 2020.

**