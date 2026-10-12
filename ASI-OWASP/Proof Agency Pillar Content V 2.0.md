## Document structure

### 5.3 Token exchange and per-request scoping patterns
- Core delegation model: Separating original principal (subject) from executing agent (actor) via RFC 8693 patterns and pre-declared authority constraints (`may_act`).
- Strong agent authentication: Secretless client authentication via RFC 7523 asymmetric assertions (`private_key_jwt`).
- Cross-domain and multi-app orchestration: Identity assertion grants (ID-JAG) for user-present, cross-system authority transfer without session exposure.
- Intent-bound, per-request authorization: Coupling target services via RFC 8707 Resource Indicators and fine-grained operation parameters via RFC 9396 Rich Authorization Requests.
- Assessment evidence.
- Standards gaps.
### 5.4 Sender-constrained tokens

- Threat model: Bearer token compromise in agentic contexts (prompt injection, context window inclusion, trace logs).
- Proof-of-possession mechanisms: Transport-layer certificate binding (mTLS / RFC 8705) versus application-layer request signing (DPoP / RFC 9449).
- Key isolation: Architectural constraints for non-exportable, runtime-isolated key material outside the model’s invocation surface.
- Comparative summary: Technical trade-offs across deployment topologies.
- Assessment evidence.
- Standards gaps.

---

## V 0.1

### 5.3 Token exchange and per-request scoping patterns  

When an agent acts on a user's session token, it inherits the entire reach of that session for its full lifetime. Token exchange replaces this pattern. A trusted Security Token Service (STS) accepts the user’s identity artifact as the subject and the executing agent’s workload identity (Pillar 1) as the actor credential, issuing a short-lived authorization grant explicitly bounded in audience, scope, and operational lifetime. The primary user session credential never enters the agent runtime context. To prevent credential forgery, the presenting actor credential must be the agent's verified workload identity authenticated via asymmetric signed assertions rather than shared static secrets. The resulting grant strictly expresses delegation rather than impersonation: the subject identifies on whose behalf the operation runs, while the actor identifies the presenting entity executing the task. Where downstream delegation paths can be anticipated, the original authorization establishes a binding constraint, such as a pre-declared actor policy (`may_act`), restricting which intermediate workloads are valid actors for that subject before exchange occurs.  
To execute cross-domain tasks across separate application boundaries without exposing ambient user credentials or causing repetitive re-authentication, the agent leverages user-present identity assertions. An initial user consent generates an identity assertion grant that the agent exchanges across target domain endpoints to obtain scoped target grants, maintaining principal provenance while isolating tool execution boundaries.  

Per-request scoping enforces strict authorization boundaries by issuing dedicated, single-use grants for each discrete task or tool call. Resource indicators tightly couple the issued token to a target service endpoint, ensuring a token issued for a calendar interface is immediately rejected if presented to a messaging or payment endpoint. To prevent parameter manipulation, rich authorization structures declare the precise operational parameters, such as target object, action, maximum monetary amount, or counterparty, directly within the grant request. Standard scope strings are insufficient for high-consequence agent actions.  

Within a single execution domain, internal service invocations maintain context through short-lived, signed transaction tokens issued at the ingress boundary. Intermediate workloads pass this contextual token unchanged across internal hops, allowing downstream resources to authorize calls against the original declared user intent rather than the ambient authority of the calling microservice.  

#### Assessment evidence  
Sampled token artifacts verify distinct subject and actor claims, explicit target resource identifiers, and parameter-bound authorization details. Replaying a token against a non-designated resource server or attempting token acquisition using static client secrets fails. No long-lived user tokens, refresh tokens, or static API keys are present in agent configurations, trace logs, or model context windows. Enforcement is evaluated directly at the protected resource boundary rather than at token service configuration. Level criteria are defined in Section 5.1.  

#### ==Gaps  ==
Current token exchange standards do not natively enforce pre-declared actor lists (`may_act`) across non-cooperating authorization servers. Standard identity assertion grants lack unified, cross-framework schemas for representing dynamic tool execution intents, leaving granular authorization parameter semantics dependent on proprietary authorization server implementations.  

## 5.4 Sender-constrained tokens  
Bearer tokens grant authority to any party in possession of the token artifact. In agentic runtimes, authorization tokens frequently transit intermediate processes, tool inputs, stdout logs, context windows, and execution traces, creating high exposure to exfiltration via prompt injection, indirect context leakages, or verbose telemetry logging. Sender-constrained token patterns mitigate bearer risk by cryptographically binding the authorization grant to a private key held exclusively by the legitimate presenting agent runtime. A stolen token presented without proof of access to the bound private key is structurally invalid.  

Proof of possession is enforced at either the transport or application layer:  

- Transport-layer binding: Token presentation is bound to the client certificate presented during mutual TLS negotiation. The resource server validates that the certificate fingerprint embedded within the token's confirmation claim matches the client TLS connection. This pattern is optimal for internal service meshes where agents operate with established workload identities (such as X.509 SVIDs), but fails when HTTP traffic traverses terminating reverse proxies or API gateways.  
    
- Application-layer binding: Each outbound request includes an ephemeral, signed proof assertion covering the HTTP request method, target URI, timestamp, and a hash of the presented token. Application-layer proof travels transparently through intermediate proxies and gateways, but requires protected resource servers to track proof freshness, validate request digests, and reject replayed proof assertions.

| Attribute             | Transport-Layer Binding (mTLS)                                                 | Application-Layer Binding (DPoP)                                             |
| --------------------- | ------------------------------------------------------------------------------ | ---------------------------------------------------------------------------- |
| Enforcement Layer     | Transport layer (TLS handshake)                                                | Application layer (signed HTTP headers)                                      |
| Key Material          | X.509 certificate (e.g., workload identity SVID)                               | Ephemeral or managed asymmetric key pair                                     |
| Proof Validation      | Handshake verifies private key; server matches certificate hash to token claim | Server verifies per-request signature over method, URI, time, and token hash |
| Intermediary Topology | Requires end-to-end TLS or trusted proxy certificate forwarding                | Functions through terminating proxies, gateways, and load balancers          |
| Optimal Context       | Homogeneous internal meshes with native certificate management                 | Heterogeneous networks, cross-domain calls, and proxy-heavy topologies       |

  
An absolute requirement for agentic implementations is that signing key material must remain inaccessible to the model execution context. Keys must reside within isolated runtime processes, cryptographic sidecars, hardware security modules, or secure key management services capable of performing signing operations outside the LLM address space. Key material stored in environment variables or accessible via model-invocable file utilities provides zero effective protection against a compromised or hijacked agent. Ephemeral refresh tokens must maintain identical cryptographic sender constraints to prevent offline token theft. Sender constraints prevent token theft and replay; they do not compensate for compromised runtime environments, which remain governed by Pillar 1 controls.  

#### Assessment evidence  
Sampled token artifacts contain valid cryptographic confirmation claims (`cnf`). Resource servers reject requests presenting missing proof assertions, invalid cryptographic signatures, mismatched target URIs, or replayed proof tokens. Cryptographic signing keys cannot be read, printed, or exported by model-invocable tools, prompt inputs, or system diagnostic functions. Level criteria are defined in Section 5.1.  

#### ==Gaps ==
No unified open standard allows a resource server to remotely verify that a key used for application-layer proof of possession is non-exportable and tied to a hardware-attested agent runtime. Emerging specifications for attestation-based client authentication and workload security address aspects of this requirement, but cross-vendor interoperability remains unstandardized.  

## References  

- [1] M. Jones, A. Nadalin, B. Campbell, J. Bradley, C. Mortimore, "OAuth 2.0 Token Exchange," RFC 8693, IETF, January 2020.  
    
- [2] B. Campbell, J. Bradley, H. Tschofenig, "Resource Indicators for OAuth 2.0," RFC 8707, IETF, February 2020.  
    
- [3] T. Lodderstedt, J. Richer, B. Campbell, "OAuth 2.0 Rich Authorization Requests," RFC 9396, IETF, May 2023.  
    
- [4] A. Tulshibagwale, G. Fletcher, P. Kasselman, "Transaction Tokens," draft-ietf-oauth-transaction-tokens, IETF OAuth Working Group, work in progress.  
    
- [5] B. Campbell, J. Bradley, N. Sakimura, T. Lodderstedt, "OAuth 2.0 Mutual-TLS Client Authentication and Certificate-Bound Access Tokens," RFC 8705, IETF, February 2020.  
    
- [6] D. Fett, B. Campbell, J. Bradley, T. Lodderstedt, M. Jones, D. Waite, "OAuth 2.0 Demonstrating Proof of Possession (DPoP)," RFC 9449, IETF, September 2023.  
    
- [7] M. Jones, J. Bradley, "JWT Profile for OAuth 2.0 Client Authentication and Authorization Grants," RFC 7523, IETF, May 2015.  
    
- [8] OpenID Foundation, "Identity Assertion Authorization Grant Pattern," work in progress.  
    
- [9] T. Lodderstedt, J. Bradley, A. Labunets, D. Fett, "Best Current Practice for OAuth 2.0 Security," RFC 9700, IETF, January 2025.  
    
- [10] B. Campbell, J. Salowey, A. Schwenkschuster, Y. Sheffer, "WIMSE Workload-to-Workload Authentication," draft-ietf-wimse-s2s-protocol, IETF WIMSE Working Group, work in progress.  
    
- [11] T. Looker, P. Bastian, C. Bormann, "OAuth 2.0 Attestation-Based Client Authentication," draft-ietf-oauth-attestation-based-client-auth, IETF OAuth Working Group, work in progress.

## Feedback

Feedback for 5.3 section 

- There is no mention of multi-hop delegation, make sure that there is a proper trail of the delegation as part of the tokens. Read the RFC to learn about this https://datatracker.ietf.org/doc/html/rfc8693
- This is the first time on the document that the concept of a user session token, STS is used. Can we generalize this initially and talk in terms of authorization grants and authorities (Authorization servers) and only start talking about access tokens and authorization servers once we need to?
- Make sure to include the concepts from transactional access tokens, https://datatracker.ietf.org/doc/draft-ietf-oauth-transaction-tokens/, Traditional OAuth Access Tokens: Issued for whole sessions. They convey _what_ a client may do (via `scope`), but carry no context about _why_ or in _which specific transaction_ an action occurs. Once issued, a client holding credentials can obtain new tokens without the Authorization Server (AS) evaluating policy for every single transaction.  Txn-AT Solution: Shifts OAuth from session-level authority to per-transaction authority. The AS evaluates fresh policy decisions for every individual transaction—making it ideal for AI agents performing task-scoped activities.   Using transactional access tokens, allows the authority to be evaluated based on the intent of the action.


---

v0.2

[1] M. Jones, A. Nadalin, B. Campbell, J. Bradley, C. Mortimore, "OAuth 2.0 Token Exchange," RFC 8693, IETF, January 2020.
[2] B. Campbell, J. Bradley, H. Tschofenig, "Resource Indicators for OAuth 2.0," RFC 8707, IETF, February 2020.
[3] T. Lodderstedt, J. Richer, B. Campbell, "OAuth 2.0 Rich Authorization Requests," RFC 9396, IETF, May 2023.
[4] A. Tulshibagwale, G. Fletcher, P. Kasselman, "Transaction Tokens," draft-ietf-oauth-transaction-tokens, IETF OAuth Working Group, work in progress.
[5] B. Campbell, J. Bradley, N. Sakimura, T. Lodderstedt, "OAuth 2.0 Mutual-TLS Client Authentication and Certificate-Bound Access Tokens," RFC 8705, IETF, February 2020.
[6] D. Fett, B. Campbell, J. Bradley, T. Lodderstedt, M. Jones, D. Waite, "OAuth 2.0 Demonstrating Proof of Possession (DPoP)," RFC 9449, IETF, September 2023.
[18] M. Jones, B. Campbell, C. Mortimore, "JSON Web Token (JWT) Profile for OAuth 2.0 Client Authentication and Authorization Grants" RFC 7523, IETF, May 2015.


## 5.3 Token exchange and per-request scoping patterns  
When an agent acts under a generalized, long-lived authorization grant, it inherits the principal’s full reach across all downstream systems for the duration of that session. Bound derivation replaces this ambient privilege. An issuing authority validates an initial authorization grant and evaluates dynamic policy to mint a narrower, task-specific grant for the requesting agent [1].

Primary credentials or session-level grants never enter the executing agent’s runtime context. The requesting agent authenticates to the issuing authority using its own verified workload identity (Pillar 1) via asymmetric signed assertions rather than shared secrets [18].

The resulting authorization grant strictly distinguishes delegation from impersonation . In multi-hop delegation workflows, the grant maintains an explicit, verifiable trail of derivation across intermediate steps. The token payload preserves the original principal as the subject (`sub`), while recording the current requesting workload as the immediate actor (`act`). As tasks pass through chained agents, each successive exchange appends or nests actor contexts, providing an immutable audit trail of every intermediate workload in the execution chain . Where downstream delegation paths are known, the issuing authority enforces pre-declared actor policies (`may_act`) to specify ahead of time which intermediate entities are permitted to act on behalf of the subject, preventing unauthorized workloads from inserting themselves into the delegation chain .  

Per-request scoping shifts authorization from session-level access to per-transaction evaluation [4]. Traditional access tokens indicate what an agent is broadly permitted to do, but carry no context regarding why or within which specific transaction an action takes place. Transactional authorization patterns resolve this by requiring the issuing authority to evaluate policy against the declared intent of each individual transaction. The resulting access tokens are single-purpose and short-lived, carrying immutable transaction context and intent bindings. Resource indicators tightly bind access tokens to specific target resource endpoints, causing a token minted for a designated resource to be rejected if presented to any non-designated resource server [2]. Fine-grained authorization structures declare exact operational bounds, such as permitted parameters, action types, or threshold limits, directly within the request and resulting token, preventing parameter manipulation during execution [3].  

Within a single trust domain, internal system calls maintain context using short-lived transaction tokens issued at the domain boundary. Intermediate workloads pass these contextual assertions unchanged across internal calls. Downstream resource servers authorize requests against the verified transaction intent rather than the ambient privilege of the calling service.  

Assessment evidence  
Decoded authorization tokens demonstrate distinct subject (`sub`) and actor (`act`) claims, a complete multi-hop delegation trail, explicit resource indicators, and parameter-bound authorization details. Attempts to replay a token at a non-designated resource server, alter transaction context parameters, or execute an unauthorized hop in a delegation chain are rejected. No long-lived principal credentials or static secrets exist within agent context windows, trace logs, or runtime memory. Authorization enforcement is verified at the protected resource boundary rather than at token issuance. Level criteria are defined in Section 5.1.  

Gaps  
Standard token exchange specifications treat nested actor claims as informational, leaving multi-hop chain validation and cascading revocation enforcement to custom implementation logic. Cross-domain trust frameworks lack standardized schemas for representing dynamic transaction intent, requiring domain-specific authorization server mappings.  

## 5.4 Sender-constrained tokens  

Bearer tokens grant authority to any party in possession of the token artifact. In agentic runtimes, authorization tokens frequently transit intermediate processes, tool inputs, stdout logs, context windows, and execution traces, creating high exposure to exfiltration via prompt injection, indirect context leakages, or verbose telemetry logging. Sender-constrained token patterns mitigate bearer risk by cryptographically binding the authorization grant to a private key held exclusively by the legitimate presenting agent runtime [T. Lodderstedt, J. Bradley, A. Labunets, D. Fett, "Best Current Practice for OAuth 2.0 Security," RFC 9700, IETF, January 2025.]. A stolen token presented without proof of access to the bound private key is structurally invalid.  

Proof of possession is enforced at either the transport or application layer:  

- Transport-layer binding: Token presentation is bound to the client certificate presented during mutual TLS negotiation [5]. The resource server validates that the certificate fingerprint embedded within the token's confirmation claim matches the client TLS connection. This pattern is optimal for internal service meshes where agents operate with established workload identities, such as X.509 SVIDs [B. Campbell, J. Salowey, A. Schwenkschuster, Y. Sheffer, "WIMSE Workload-to-Workload Authentication," draft-ietf-wimse-s2s-protocol, IETF WIMSE Working Group, work in progress.], but fails when HTTP traffic traverses terminating reverse proxies or API gateways.  
    
- Application-layer binding: Each outbound request includes an ephemeral, signed proof assertion covering the HTTP request method, target URI, timestamp, and a hash of the presented token [6]. Application-layer proof travels transparently through intermediate proxies and gateways, but requires protected resource servers to track proof freshness, validate request digests, and reject replayed proof assertions.
  

| Attribute             | Transport-Layer Binding (mTLS)                                                                                                                                                                                                         | Application-Layer Binding (DPoP)                                             |
| --------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------- |
| Enforcement Layer     | Transport layer (TLS handshake)                                                                                                                                                                                                        | Application layer (signed HTTP headers)                                      |
| Key Material          | X.509 certificate, such as a workload identity SVID [B. Campbell, J. Salowey, A. Schwenkschuster, Y. Sheffer, "WIMSE Workload-to-Workload Authentication," draft-ietf-wimse-s2s-protocol, IETF WIMSE Working Group, work in progress.] | Ephemeral or managed asymmetric key pair                                     |
| Proof Validation      | Handshake verifies private key; server matches certificate hash to token claim                                                                                                                                                         | Server verifies per-request signature over method, URI, time, and token hash |
| Intermediary Topology | Requires end-to-end TLS or trusted proxy certificate forwarding                                                                                                                                                                        | Functions through terminating proxies, gateways, and load balancers          |
| Optimal Context       | Homogeneous internal meshes with native certificate management                                                                                                                                                                         | Heterogeneous networks, cross-domain calls, and proxy-heavy topologies       |

  
An absolute requirement for agentic implementations is that signing key material must remain inaccessible to the model execution context. Keys must reside within isolated runtime processes, cryptographic sidecars, hardware security modules, or secure key management services capable of performing signing operations outside the LLM address space. Key material stored in environment variables or accessible via model-invocable file utilities provides zero effective protection against a compromised or hijacked agent. Ephemeral refresh tokens must maintain identical cryptographic sender constraints to prevent offline token theft [T. Lodderstedt, J. Bradley, A. Labunets, D. Fett, "Best Current Practice for OAuth 2.0 Security," RFC 9700, IETF, January 2025.]. Sender constraints prevent token theft and replay; they do not compensate for compromised runtime environments, which remain governed by Pillar 1 controls.  

Assessment evidence  
Sampled token artifacts contain valid cryptographic confirmation claims (`cnf`). Resource servers reject requests presenting missing proof assertions, invalid cryptographic signatures, mismatched target URIs, or replayed proof tokens. Cryptographic signing keys cannot be read, printed, or exported by model-invocable tools, prompt inputs, or system diagnostic functions. 

Gaps  
No unified open standard allows a resource server to remotely verify that a key used for application-layer proof of possession is non-exportable and tied to a hardware-attested agent runtime. Emerging specifications for attestation-based client authentication and workload security address aspects of this requirement, but cross-vendor interoperability remains unstandardized [T. Looker, P. Bastian, C. Bormann, "OAuth 2.0 Attestation-Based Client Authentication," draft-ietf-oauth-attestation-based-client-auth, IETF OAuth Working Group, work in progress.] [B. Campbell, J. Salowey, A. Schwenkschuster, Y. Sheffer, "WIMSE Workload-to-Workload Authentication," draft-ietf-wimse-s2s-protocol, IETF WIMSE Working Group, work in progress.].

| Attribute             | Transport-Layer Binding (mTLS)                                                 | Application-Layer Binding (DPoP)                                             |
| --------------------- | ------------------------------------------------------------------------------ | ---------------------------------------------------------------------------- |
| Enforcement Layer     | Transport layer (TLS handshake)                                                | Application layer (signed HTTP headers)                                      |
| Key Material          | X.509 certificate, such as a workload identity SVID                            | Ephemeral or managed asymmetric key pair                                     |
| Proof Validation      | Handshake verifies private key; server matches certificate hash to token claim | Server verifies per-request signature over method, URI, time, and token hash |
| Intermediary Topology | Requires end-to-end TLS or trusted proxy certificate forwarding                | Functions through terminating proxies, gateways, and load balancers          |
| Optimal Context       | Homogeneous internal meshes with native certificate management                 | Heterogeneous networks, cross-domain calls, and proxy-heavy topologies       |
