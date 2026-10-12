https://datatracker.ietf.org/doc/draft-ietf-oauth-identity-assertion-authz-grant/

The Internet-Draft draft-ietf-oauth-identity-assertion-authz-grant defines the Identity Assertion JWT Authorization Grant (often referred to as ID-JAG or Cross-App Access / XAA).  

It specifies a standard way for a client application to take an identity assertion (like an ID Token from a central Enterprise Identity Provider) and exchange it for a scoped access token to call a third-party API in a different authorization domain—without prompting the user to re-authenticate or re-grant consent.  

## Core Problem It Solves  

In enterprise environments, users authenticate with a central Enterprise Identity Provider (IdP) via Single Sign-On (SSO). However, modern applications (like AI agents, email clients, or SaaS integrations) often need to access third-party resource servers that manage their own OAuth authorization servers (AS).  

Without ID-JAG:  
- Users face constant OAuth consent prompts across connected services.  
- Service-to-service calls lack user-scoped delegated authority.  
    
With ID-JAG:  
- The common Enterprise IdP acts as the central trust broker to issue an identity-based assertion.  
- The application exchanges that assertion with the target Resource AS to obtain a user-scoped access token seamlessly.  

## Key Components & Building Blocks  

ID-JAG bridges two existing RFC standards:  

1. RFC 8693 (OAuth 2.0 Token Exchange): Used at the Enterprise IdP to convert the initial user session or ID token into a specially crafted Identity Assertion JWT.  
2. RFC 7523 (JWT Profile for OAuth 2.0 Authorization Grants): Used at the Target Resource AS to exchange the Identity Assertion JWT for a local access token.  

## How the Flow Works  

1. User Authentication: The user logs into the Client App via the Enterprise IdP (IdP AS).  
2. Token Exchange at IdP:  
    The Client sends a Token Exchange (`RFC 8693`) request to the IdP's `/token` endpoint.  
    `requested_token_type`: `urn:ietf:params:oauth:token-type:id-jag`  
    `audience`: The target Resource AS's Issuer URL.  
    The IdP issues an Identity Assertion JWT Authorization Grant specifically targeted to that Resource AS.  
3. Authorization Grant at Resource AS:  
    The Client presents the ID-JAG assertion to the Resource AS's /token endpoint.  
    `grant_type`: `urn:ietf:params:oauth:grant-type:jwt-bearer`  
    `assertion`: The ID-JAG token obtained from Step 2.  
    The Resource AS verifies the assertion against the Enterprise IdP (via trust metadata/JWKS) and returns a local Access Token.  
4. API Call: The Client calls the target Resource Server using the local Access Token.  
## Primary Use Cases  
- Model Context Protocol (MCP) & AI Agents: AI assistants operating on a user's behalf to access enterprise data across distinct microservices or third-party tools while maintaining exact user identity constraints.  
- Cross-App Automation / Cross-Domain Single Sign-On (XAA): Allowing an email/calendar app to seamlessly gain access to an external videoconferencing API without re-prompting the user.  
- Delegated Authorization Chaining: Propagating original end-user identities safely across security domain boundaries.

---

## Security Trade Offs

Eliminating interactive OAuth consent prompts streamlines the user experience—especially for autonomous AI agents—but it raises a critical risk: How do you keep an agent from asking for data it doesn't need without alerting the user?  
The ID-JAG spec (and the underlying enterprise trust model) addresses this problem by shifting consent from interactive, real-time user prompts to pre-configured policy, strict scoping, and centralized governance.  

### How the Risk Is Mitigated  

1. Scope Containment & Downward Scoping  
An agent cannot simply request any scope it wants. The ID-JAG draft specifies a strict containment chain:  

- Initial SSO Scope: When the user logs in via the Enterprise IdP, a baseline set of user privileges is established.  
    
- Request Minimization: When requesting an ID-JAG assertion, the client passes a `scope` parameter. The IdP will only grant scopes that the user explicitly holds and that the target application policy allows.  
    
- Resource AS Enforcers: When the agent presents the ID-JAG token to the target Resource Authorization Server (AS), that server applies its own local policy. If the requested scope exceeds what the IdP assertion signed off on, the request is rejected with an `invalid_scope` error.  
    

2. Dual Authorization Policies (IdP + Resource AS)  
Instead of relying on the end user to read consent screens, authorization decisions are enforced at two administrative layers:  

- Enterprise Policy (The IdP): Administrators configure pre-authorized trust policies (e.g., _"AI Agent X is allowed to read Calendar events for Department Y, but cannot access HR or Payroll data"_). If an agent attempts to forge or escalate its requested scope, the IdP blocks issuance of the ID-JAG JWT.  
    
- Resource Policy (Target Service): The Resource AS evaluates the `aud` (audience) and `scope` claims inside the assertion. It retains final authority on whether to issue the final access token.  
    

3. Client Identity Bound to the Token (Actor Claims)  
In classic OAuth, an access token often just says "User A". With ID-JAG, the assertion explicitly names both the user and the agent (via nested actor claims `act`).  

- The target API knows the call is coming from _Agent X acting on behalf of User A_.  
    
- This allows Resource Servers to apply finer-grained access policies (e.g., _"User A can edit records manually, but when acting via Agent X, the access is strictly read-only"_).  
    

## What if an Agent Requests Data It Doesn't Need?  
If an agent behaves maliciously or is poorly designed (requesting overly broad data), several guardrails are built into the architecture:  

Real-Time Step-Up Authentication / Interactive Fallback  
The ID-JAG spec explicitly limits when non-interactive assertions are allowed:  

- Confidential Clients Only: ID-JAG is restricted to authenticated, confidential clients.  
    
- Fallback to Interactive Authorization Code Grant: If an agent requests a sensitive scope or an action that falls outside pre-approved policy, the Resource AS or IdP can reject the ID-JAG flow (`interaction_required`). The agent is forced to fallback to a traditional OAuth Authorization Code flow, prompting the user with an explicit consent screen.  
    

Auditability & Admin Dashboards  
Because token exchanges funnel through the enterprise IdP and target Resource AS:  

- Security Operations (SecOps) teams get centralized audit logs showing every service-to-service token request, including the requested scopes and target audiences.  
    
- Organizations can instantly revoke an agent's grant globally at the IdP layer if anomalous scope requests are detected.  
    

Summary: The Shifts in the Security Model  

|Traditional OAuth Consent|ID-JAG Enterprise Model|
|---|---|
|User-Driven: The end-user evaluates every prompt.|Policy-Driven: Enterprise admins define acceptable scopes beforehand.|
|Interactive: Interrupts workflow with pop-ups.|Non-Interactive: Seamless execution within pre-approved boundaries.|
|Coarse Scopes: Users often blindly click "Allow All".|Least Privilege & Delegation: IdP and Resource AS enforce scope minimization and actor-based controls.|

---

## Deep dives

### Understanding of the Core Flow

In traditional OAuth 2.0, when Application A wants to access a user's data in Application B, the user is redirected through an interactive browser flow to consent to the scopes.  
With ID-JAG (Cross-App Access):  

1. SSO Context: The user is already authenticated with an Identity Provider (IdP) via single sign-on (SSO).  
    
2. Token Exchange: Application A presents an Identity Assertion (like an OpenID Connect ID Token or SAML assertion) to the IdP and receives back an ID-JAG (a specially scoped JWT authorization grant).  
    
3. Bypassing Interactive Consent: Application A takes this ID-JAG to Application B's Resource Authorization Server and exchanges it directly for an Access Token.  
    
4. Policy-Driven: Instead of prompting the end user with interactive consent screens, authorization is governed by pre-established trust relationships and access policies managed by the administrator (or organization platform) across both domains.


### Enterprise vs. Consumer Scenarios

While enterprise workforce environments (e.g., Okta, Microsoft Entra, Ping Identity) are the primary driver for this spec, it is not exclusively limited to enterprise.  

The specification explicitly outlines multi-domain and platform trust patterns beyond enterprise workforce management:  

- Customer Identity & Developer Platforms (CIAM): In developer platforms or SaaS ecosystems, a consumer/customer signs into a main platform account (e.g., a developer console or creative hub). The platform IdP can broker Cross-App Access to first-party or third-party partner tools without repeatedly asking the user to consent to interconnected services.  
    
- Email, Calendaring, & Integrated Apps: Connected consumer productivity suites where the identity provider brokers data access between companion applications on behalf of the user under account-level policies.  
    
- Why it works in consumer contexts: In consumer platforms, the "administrator policy" translates to platform-level trust policies or account-level pre-authorizations agreed to during account setup/terms, removing friction from app-to-app integrations.

### Confidential Client restrictions of ID-JAG

Confidential Client Requirement  
The ID-JAG spec heavily emphasizes Confidential Clients for both stages of the exchange:  

1. Token Exchange with the IdP: The client must authenticate with the IdP.  
2. Access Token Request with the Resource AS: The client must authenticate itself when presenting the ID-JAG.  

This requirement exists because the spec relies on client authentication (e.g., mTLS, private key JWTs) to ensure that only authorized, trusted applications can request cross-app tokens on behalf of users without explicit consent. Bare public clients (like standard browser SPA JS apps) cannot safely hold a secret to identify themselves at both trust boundaries.  
#### How This Applies to AI Agents on a User's Desktop 

When an AI agent (or local desktop app) runs on a user's machine, it is traditionally treated as a Public Client. To participate in ID-JAG securely, architecture patterns bridge local agents to confidential client standards:  

1. **Backend / Gateway Proxy (Recommended Architecture)**:  
    The local AI agent running on the desktop communicates with a trusted cloud/backend service (which acts as the Confidential Client).  
    The local agent authenticates the user, but the backend performs the ID-JAG exchange with the IdP and third-party Resource Authorization Servers.  
    
2. **Sender-Constrained Tokens & Proof-of-Possession (Local Agent)**:  
    The specification details Proof-of-Possession (PoP) and sender-constraining mechanisms (such as DPoP — _Demonstrating Proof-of-Possession_ via RFC 9449).  
    
    If a local AI agent uses a public client model, it generates a asymmetric key pair locally. The ID-JAG token is cryptographic bound to the agent’s public key. Even if the local agent holds the token, it must sign requests using its private key to prove ownership, preventing token theft/hijacking on the endpoint.  
    
3. AI Agent Tool Use (Appendix A.4 of the Spec):  
    Section A.4 specifically addresses AI Agents executing external tools.  
    When an AI Agent needs to perform actions on third-party APIs (e.g., booking a meeting, querying a database), it obtains an ID-JAG targeting that specific tool/resource server. The IdP can enforce fine-grained constraints (like narrowing scopes or adding step-up authentication triggers) before issuing the grant for the AI tool call.

#### How Proof-of-Possession Solves the "Public Client" Problem

In traditional OAuth, a public client cannot safely store a client secret (`client_secret`), making standard confidential client authentication impossible.  
With Proof-of-Possession (most notably via DPoP - _Demonstrating Proof-of-Possession_, RFC 9449):  

1. Local Key Pair: The AI agent running on the desktop generates an ephemeral or persistent asymmetric key pair (Public/Private Key) stored securely in the local operating system's keyring (e.g., TPM, Secure Enclave, or Windows Credential Manager).  
    
2. Client Assertion Signing: When requesting the ID-JAG from the IdP or presenting the ID-JAG to the Resource Authorization Server, the AI agent signs a client assertion (`private_key_jwt`) or DPoP proof using its private key.  
    
3. Cryptographic Binding: The server verifies the signature against the agent's public key. The issued token is strictly bound to that specific key pair.  
    
Because the private key never leaves the local machine and cannot be extracted, the agent proves its identity through cryptography rather than a shared static secret.

When a desktop AI agent uses DPoP with ID-JAG, proof-of-possession applies at two distinct levels:  

- Grant Level (ID-JAG Binding): When the IdP issues the ID-JAG, it binds the grant assertion to the AI agent's public key (using the `cnf` / confirmation claim). When the agent passes the ID-JAG to the Resource Server's Authorization Server, it must prove possession of the private key to redeem it.  
    
- Access Token Level (API Access): The final Access Token issued by the Resource Authorization Server is also sender-constrained via DPoP. If a malicious actor intercepts the access token, it is completely useless on another device because they do not hold the desktop agent's private key to sign API requests.