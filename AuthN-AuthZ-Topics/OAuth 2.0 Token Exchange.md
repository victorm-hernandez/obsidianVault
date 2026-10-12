
https://datatracker.ietf.org/doc/html/rfc8693

In OAuth 2.0, Delegation allows an intermediate service or AI agent to act on behalf of a user while maintaining its own distinct identity. This is standardized by RFC 8693 (OAuth 2.0 Token Exchange).  

The core mechanism relies on actor claims (act) to distinguish delegation from impersonation and token chaining to preserve an auditable lineage as requests hop across multiple services.  

## 1. Impersonation vs. Delegation  

Understanding the difference is critical to understanding why the `act` claim exists:  

- Impersonation: Service A presents User B's token to get a new token. The new token's subject (`sub`) is set to User B, and Service A's identity disappears completely. The downstream Resource Server cannot distinguish between the actual user and Service A.  
    
- Delegation: Service A exchanges User B's token for a new token where the subject is still User B, but a composite act claim is added. The downstream API knows: _"This call targets User B's data, but Service A is executing the request."_

## 2. The `act` (Actor) Claim Payload  

When a client performs an OAuth Token Exchange (RFC 8693), it sends both a `subject_token` (representing the user) and an `actor_token` (representing the caller/agent).  

The Authorization Server (Security Token Service) issues a token with an `act` claim containing the actor's identity:

```json
{
  "sub": "user_12345",
  "aud": "https://api.payments.com",
  "iss": "https://auth.enterprise.com",
  "exp": 1775587200,
  "scope": "payments:read",
  "act": {
    "sub": "ai_agent_finance_v2",
    "iss": "https://auth.enterprise.com"
  }
}

```

- sub (Subject): The end-user who owns the resource or data.  
- act (Actor): The party actively making the request under delegated authority.

## 3. Token Chaining Across Multiple Hops  

When a request travels through a chain of microservices or multi-agent pipelines (e.g., `User -> AI Assistant -> Workflow Engine -> Database`), the delegation chain grows.  

RFC 8693 specifies that multi-hop chains are expressed by nesting act claims. The rules dictate:  

1. The outermost `act` claim represents the current/most recent actor.  
2. Nested `act` claims represent prior actors in chronological order (deepest = oldest).  

Multi-Hop Example:  
An AI Assistant calls a Scheduling Service, which in turn calls a Calendar API.

```json
{
  "sub": "user_12345",
  "aud": "https://api.calendar.com",
  "act": {
    "sub": "scheduling_service",
    "act": {
      "sub": "ai_assistant_agent"
    }
  }
}
```

How to read this chain:  

- Resource Owner: `user_12345`  
- Immediate Actor (outermost): `scheduling_service`  
    Prior Actor (nested): `ai_assistant_agent`

## 4. How Resource Servers Enforce Access Policy  
When an API receives a token containing an `act` claim, it evaluates authorization against both identities:  

1. User Permission Check: Does `user_12345` have read access to the target resource?  
2. Actor Scope Minimization: Is `scheduling_service` (or `ai_assistant_agent`) authorized to act on behalf of `user_12345` for this specific action?  

For example, a banking API might allow `user_12345` to perform wire transfers manually, but if the incoming token shows `act.sub = "ai_assistant_agent"`, local policy can reject the wire transfer while permitting read-only balance checks.

## 5. Companion Claim: `may_act`  
RFC 8693 also defines the may_act claim, which works in reverse:  

- `act` records delegation that has occurred.  
- `may_act` asserts that a specific party is authorized to become an actor in the future.  

An initial token might state:  

```json
{  
"sub": "user_12345",  
"may_act": {    
	"sub": "ai_assistant_agent"  
	}
}
```

  
When the `ai_assistant_agent` later presents this token to the Token Endpoint, the Authorization Server inspects `may_act` to confirm the agent is allowed to execute a token exchange on behalf of `user_12345`.


---
## Security Risks in Token Chaining

While OAuth 2.0 Token Exchange (RFC 8693) provides a way to represent nested actors via the `act` claim, applying token exchange across multi-hop pipelines (such as AI agents calling other sub-agents or tool APIs) introduces severe multi-hop security risks.  

## The Main Security Risks in Token Chaining  

### 1. Delegation Chain Splicing (The Primary Risk)  

- How it works: In standard RFC 8693, an intermediary sends a `subject_token` (e.g., User A's token) and an `actor_token` (e.g., Agent B's token) to an Authorization Server (AS). The AS validates both tokens independently, but does not natively cross-validate that they belong to the same execution context.  
    
- The Exploit: A compromised intermediary holding two completely unrelated, genuine tokens—Token 1 (User A authorizing Read access) and Token 2 (High-Privilege Agent X authorizing Write access)—can submit Token 1 as the `subject_token` and Token 2 as the `actor_token`. The AS verifies both individually and issues a valid, signed assertion asserting a delegation relationship (`Agent X acting for User A`) that never actually existed.  
    
- Impact: Fabricated cross-user or cross-agent authority that passes signature verification downstream.  

### 2. Authority Expansion (Scope Bloat across Hops)  

- How it works: When Agent A delegates a task to Agent B, Agent B should only inherit the intersection of Agent A's privileges and User's privileges.  
    
- The Exploit: If the AS mints tokens using the union of permissions (or fails to enforce strict scope reduction at each hop), an agent low in the delegation chain can execute commands with broader authority than the user originally authorized, leading to a "confused deputy" flaw across the chain.  

### 3. Token Passthrough & Replay Across Domains  

- How it works: If an intermediate agent receives a bearer token and forwards it directly to downstream services (instead of performing a fresh token exchange), any captured token can be replayed against any other API or service that trusts the same Authorization Server.  
    
### 4. "Informational-Only" Audit Trails vs. Enforcement  

- How it works: According to RFC 8693, downstream resource servers generally only evaluate the outermost `act` claim and the token subject. The nested `act` array (prior history) is treated as an unverified audit trail rather than cryptographically enforced capability constraints.  
    
- The Exploit: If an agent deep in the chain is compromised, downstream APIs cannot easily verify whether prior agents in the chain had the right to delegate that specific scope offline.  

---
## How to Prevent These Risks  

Modern security profiles, IETF draft proposals (such as _Attenuating Authorization Tokens_ and _Agent Delegation Chains_), and production best practices use several defenses:  

### 1. Cryptographic Byte-Commitments (Fixing Chain Splicing)  

To stop chain splicing, child tokens must be cryptographically tied to their parent token.  

- During token exchange, the Authorization Server or delegating agent embeds a cryptographic hash/commitment of the parent token inside the child token.  
    
- When verifying the chain, the recipient checks that `Hash(Parent Token) == Child.parent_hash`. Splicing two unrelated tokens causes this hash check to fail immediately.  
    
Lets go over an example.

When an agent sends the exchange request to the Authorization Server (AS), it provides both input tokens:  

```http
POST /token HTTP/1.1
Host: auth.enterprise.com
Content-Type: application/x-www-form-urlencoded

grant_type=urn%3Aietf%3Aparams%3Aoauth%3Agrant-type%3Atoken-exchange
&subject_token=eyJhbGciOiJSUzI1NiI... (User A's Token)
&subject_token_type=urn%3Aietf%3Aparams%3Aoauth%3Atoken-type%3Aaccess_token
&actor_token=eyJhbGciOiJSUzI1NiI...   (Agent B's Token)
&actor_token_type=urn%3Aietf%3Aparams%3Aoauth%3Atoken-type%3Aaccess_token
```

  
To create a byte-commitment, the Authorization Server hashes the exact byte representations of the inputs before generating the new token:

1. Calculate Hashes:  
$$H_{sub} = \text{SHA-256}(\text{subject\_token bytes})$$
$$H_{act} = \text{SHA-256}(\text{actor\_token bytes})$$
2. Embed Hashes into Payload: The AS writes these thumbprints directly into the new claims (e.g., inside an `act_hash` or `sth` / _subject token hash_ claim) before signing the JWT:

```json
{  
"sub": "user_12345",  
"aud": "https://api.calendar.com",  
"iss": "https://auth.enterprise.com",  
"sth": "n4bQgYhJi...",   
"act": {    "sub": "ai_agent_b",    "ath": "x9aK2mPz1..."  }}
```

- sth (Subject Token Hash): SHA-256 hash of the exact `subject_token` binary payload.  
- ath (Actor Token Hash): SHA-256 hash of the exact `actor_token` binary payload.

When a downstream Resource Server (or the AS on a subsequent hop) processes the Delegated Access Token:  

1. It reads the `sth` and `ath` claims inside the JWT.  
2. If an attacker attempts to substitute a different `actor_token` or `subject_token` upstream, the cryptographic hashes no longer match the signatures signed by the AS.  
3. Because the AS signed the hash of the exact bytes of both inputs simultaneously, Subject Token A can only ever be combined with Actor Token B. Splicing Subject Token A with a stolen Actor Token C results in an invalid signature or a hash mismatch, breaking the exploit.

To see an example of this attack go to the appendix A "Delegation Chain Splicing Attack"
### 2. Monotonic Attenuation (Strict Downward Scoping)  

Enforce a hard rule: Permissions MUST strictly diminish or stay equal at every hop ($Scope_{N+1} \subseteq Scope_N$).  

- Each delegation step explicitly strips unnecessary scopes using Rich Authorization Requests (RAR - RFC 9396).  
    
- The downstream resource server verifies that no hop gained permissions that its delegating parent lacked.  
    
### 3. Strict Audience Binding & Sender Constraining  

- Audience Binding (aud): Every token issued during a chain exchange must be strictly bound to the specific target resource or downstream agent (`aud: "[https://api.calendar.com](https://api.calendar.com)"`). A token issued for Service B will be rejected if presented to Service C.  
    
- Proof-of-Possession (DPoP / mTLS): Bind tokens to cryptographic key pairs held by the specific agent instance (via DPoP - RFC 9449 or mTLS) so stolen bearer tokens cannot be replayed elsewhere.  

### 4. Explicit Pre-Approval (`may_act` / Consent Binding)  

- Before an agent can issue a delegated token exchange, the original `subject_token` must explicitly state who is allowed to act on its behalf using the may_act claim or user-signed consent rules. If Agent C attempts to perform a token exchange using a token that lacks `may_act: Agent C`, the AS rejects the request.  
    
### 5. Operational Guardrails  

- Enforce Maximum Chain Depth: Limit token chaining to a fixed depth (e.g., maximum 3 hops) at the Authorization Server level.  
    
- Micro-Lifetimes: Issue ultra-short token expiration times (e.g., 2–5 minutes) for multi-hop tokens to minimize the attack window if an agent memory space is exfiltrated.  
    
- Never Pass Through: Prohibit agents from forwarding upstream bearer tokens directly; require every hop to call the AS/STS for an audience-restricted downstream token.


---

## Appendix Delegation Chain Splicing Attack

Here is a step-by-step scenario demonstrating Delegation Chain Splicing (also known as _Context Mixing_ or _Cross-Token Substitution_) without byte-commitment hashes.  

The Setup: Two Legitimate, Unrelated Flows  
Imagine an enterprise environment with two active, legitimate workflows occurring at the same time:  

- Flow 1 (User A & Reading Emails): User A grants Agent 1 (a low-privilege Email Summarizer) permission to read their emails.  
    Token A (subject_token): Belongs to User A, holds scope `email:read`.  
    
- Flow 2 (High-Privilege Service B): Service B is a highly trusted background payment processor that holds a valid identity token proving it is an internal system service.  
    Token B (actor_token): Belongs to Service B, holds high-level execution privileges.  
    

The Exploit: How an Attacker Splices the Tokens  
Imagine Agent 1 is compromised or running untrusted prompt injection code. Agent 1 intercepted Token A (from User A) and managed to snag a copy of Token B (from Service B) from an internal log or memory cache.  
Without byte-commitments, the Authorization Server (AS) treats `subject_token` and `actor_token` as independent, unlinked inputs.  

```
[ Compromised Agent 1 ]       │       ├── Intercepts Token A (User A: email:read)       └── Steals Token B (Service B: payment:admin)       │       ▼   Submits Spliced Request to AS:   - subject_token = Token A   - actor_token   = Token B
```

  
The Malicious Request Over the Wire:  

```http
POST /token HTTP/1.1
Host: auth.enterprise.com
Content-Type: application/x-www-form-urlencoded

grant_type=urn%3Aietf%3Aparams%3Aoauth%3Agrant-type%3Atoken-exchange
&subject_token=eyJhbG... [Token A: User A's Email Token]
&actor_token=eyJhbG...   [Token B: Service B's Payment Identity]
&audience=https%3A%2F%2Fapi.payments.com
&scope=payments:write
```

  
What Happens Without Byte-Commitments  
Because there are no hashes linking the context:  

1. AS Validates subject_token: "Is Token A signed by our IdP? Yes. Belongs to User A."  
    
2. AS Validates actor_token: "Is Token B signed by our IdP? Yes. Belongs to Service B."  
    
3. AS Evaluates Policy (Flawed): "Service B is a trusted payment processor, and User A owns the account. This looks valid!"  
    

The Authorization Server issues a brand new, validly signed Delegated Token:  

```json
{  
"sub": "user_12345",  
"aud": "https://api.payments.com",  
"scope": "payments:write",  
"act": {    "sub": "service_b_payment_processor"  }}
```

  
The Result of the Attack  

- Spliced Identity: The AS just minted a token stating that Service B is acting on behalf of User A to write payments.  
    
- Fabricated Authority: User A _never_ authorized Service B or Agent 1 to execute payments—User A only logged into an email summarizer!  
    
- Bypassed Consent: The attacker bridged User A's identity with Service B's high privileges, completely bypassing user consent and scope limitations.  
    

Why the Byte-Commitment (Hash) Stops This Cold  
When byte-commitments are enforced, Token B must contain a hash of the exact `subject_token` context in which it was issued, or the AS must verify that Token B's `sth` claim matches Token A:  

$$\text{Token B's } \texttt{sth} \neq \text{SHA-256}(\text{Token A})$$
1. The AS calculates
$\text{SHA-256}(\text{Token A})$

2. It compares it to the subject token hash embedded inside Token B.  
    
3. The hashes do not match because Token B was originally issued for a completely different execution context.  
      
    The AS immediately rejects the exchange with an `invalid_grant` error.