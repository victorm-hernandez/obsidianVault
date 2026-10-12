
### Current References
[1] M. Jones, A. Nadalin, B. Campbell, J. Bradley, C. Mortimore, "OAuth 2.0 Token Exchange," RFC 8693, IETF, January 2020.
[2] B. Campbell, J. Bradley, H. Tschofenig, "Resource Indicators for OAuth 2.0," RFC 8707, IETF, February 2020.
[3] T. Lodderstedt, J. Richer, B. Campbell, "OAuth 2.0 Rich Authorization Requests," RFC 9396, IETF, May 2023.

### New references
- CIMD
- Dynamic Client Registration (DCR)
	- https://www.linkedin.com/feed/update/urn:li:activity:7441842233231949824/
- ID-JAG stands for [Identity Assertion JWT Authorization Grant](https://datatracker.ietf.org/doc/draft-ietf-oauth-identity-assertion-authz-grant/), an emerging security standard that lets a central identity provider give a temporary, restricted token to an application or AI agent so it can safely access another app's API 
	- Identity Assertion (JAG) https://datatracker.ietf.org/doc/draft-ietf-oauth-identity-assertion-authz-grant/
- [Cross App Access (XAA)](https://www.okta.com/solutions/cross-app-access/) es un ==protocolo de identidad basado en OAuth diseñado para asegurar y gestionar la comunicación entre aplicaciones y agentes de inteligencia artificial en entornos empresariales==
- https://www.ietf.org/archive/id/draft-tulshi-oauth-transactional-access-tokens-00.html

## Current Text

### 5.3 Token exchange and per-request scoping patterns (Alessandro/Victor)

An agent acting on a user's session token inherits all of the user's reach for the life of the session. Token exchange [1] replaces that. A trusted Security Token Service (STS) takes the user's token as subject_token and the agent's workload identity (Pillar 1) as actor_token, and returns a token narrower in scope, audience and lifetime. The user's token never enters the agent runtime. The actor credential must be the agent's own workload identity, not a shared client secret, or the exchange simply turns a static secret into a legitimate-looking token. The result must express delegation, not impersonation: sub is the user and act is the agent, so every decision and log entry separates who asked from who acted (chaining is covered in 5.5).

Per-request scoping goes further and issues one token per task or tool call. Resource Indicators [2] restrict the token to a single resource server, so a token minted for the calendar API is rejected by the mail API. Rich Authorization Requests [3] state the exact operation and its limits, such as one payment of at most €500 to a named beneficiary, which scopes cannot express. 

Within a domain, Transaction Tokens [4] handle what happens after the first hop. When a request enters the domain, an internal token service issues a short-lived signed token that records the subject, the purpose and the request context. Every internal service passes it along unchanged and cannot alter it, so downstream services authorize against what the user asked rather than the caller's ambient privilege.

Evidence. Decoded tokens show sub, act, aud and exp. A token replayed to another resource server is rejected. No user or refresh tokens are found in agent config, logs or context. Test enforcement at the resource server, not STS configuration. Level criteria are defined in 5.1.


## Proposed Text

