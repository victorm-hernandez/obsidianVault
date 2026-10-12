
# OAuth 2.0 & Token Grant Types: Complete Study Guide

## Short reference

1. Roles: Resource, Resource Owner, Client (Confidential/Public), Authorization Server
2. Scenarios: 
	1. Authorization code with PKCE (user facing default public client scenario)
	2. Client credentials (confidential client scenario for non-human identities)
		1. Credential is a client secret (similar to a password).
		2. Credential is a self signed JWT .
		3. Credential is a mTLS certificate.
	3. Device (used for dumb devices, creating a session on the IdP and asynchronously approving on another device)
	4. Token exchange (used in distributed systems to exchange token format or enable delegated authentication)
3. Endpoints
	1. Authorize (entry point to authorize a request over the web)
	2. Token (issuance of tokens and token exchange)
	3. UserInfo (OIDC specific, provides information about the user holding this access token)
	4. Introspection (validates an access token)
	5. Revoke (token revocation) Ex.: sign out.
	6. Device (used to initiate dumb device auth)
	7. ./well-known (discovery metadata, endpoints, supported algorithms)
	8. Back Channel authorize /bc-authorize (similar to device but initiated from the service towards the client)
4. Grant types
	1. Authorization Code
	2. Device Code
	3. Client credentials (client secret/ Self Signed JWT / mTLS)
	4. Refresh Token, produce access tokens without re-auth.
	5. Token Exchange (urn:ietf:params:oauth:grant-type:token-exchange), Impersonation, Delegation, Token format exchange.
	6. Microsoft OBO JWT Bearer (urn:ietf:params:oauth:grant-type:jwt-bearer) , Federation

---

This comprehensive study guide covers OAuth 2.0 authorization flows, grant types, standard endpoints, client classifications, and critical security concepts. Designed specifically for **System Design & Security Architecture Interviews**, this document outlines the fundamental mechanics of authorization protocols.

---

## 1. Core OAuth 2.0 Terminology & Key Definitions

Before diving into flows and endpoints, it is essential to understand the primary roles and client classifications defined in **RFC 6749**.
### Roles

* **Resource Owner**: An entity capable of granting access to a protected resource (typically the end-user).

* **Client**: An application making protected resource requests on behalf of the resource owner and with its authorization.

* **Resource Server**: The server hosting the protected resources, capable of accepting and responding to protected resource requests using access tokens.

* **Authorization Server**: The server issuing access tokens to the client after successfully authenticating the resource owner and obtaining authorization.
### Client Types

Clients are categorized based on their ability to securely maintain confidential credentials (e.g., client secrets):

#### Confidential Clients

* **Definition**: Clients capable of maintaining the confidentiality of their credentials (e.g., applications implemented on a secure server with restricted access to the backend code).

* **Examples**: Traditional web applications (Node.js/Express, Spring Boot, Django) running on a protected server.

* **Security Profile**: Can securely perform client authentication using a `client_secret`, private key JWTs, or mutual TLS (mTLS).
#### Public Clients

* **Definition**: Clients incapable of maintaining the confidentiality of their credentials, or clients where credentials can be extracted via reverse engineering or client-side inspection.

* **Examples**: Single Page Applications (SPAs executing in browsers like React, Angular, Vue), Native Mobile Apps (iOS/Android), Desktop Applications, and IoT/Smart TV apps.

* **Security Profile**: Cannot use `client_secret` securely. **Must use Proof Key for Code Exchange (PKCE)** to prevent authorization code interception attacks.

---
## 2. Comprehensive OAuth 2.0 Flows

OAuth 2.0 defines several grant flows to accommodate different application architectures, trust models, and user interaction capabilities.
### A. Authorization Code Flow (with PKCE)
#### Overview

The gold standard authorization flow for both **Confidential** and **Public** clients. When used with **PKCE (Proof Key for Code Exchange - RFC 7636)**, it protects against authorization code injection and interception attacks.
#### Recommended Use Cases

* Web Applications running on backends (Confidential).
* Mobile Applications (iOS/Android native apps).
* Single Page Applications (SPAs).
#### Sequence Diagram (SPA/Mobile)

```mermaid

sequenceDiagram

autonumber

actor User as Resource Owner (User)

participant Client as Client App (SPA / Mobile)

participant AS as Authorization Server

  

User->>Client: Clicks "Login"

Note over Client: Generate code_verifier<br/>Generate code_challenge (S256)

Client->>AS: Redirect to /authorize<br/>(?response_type=code&client_id=...&code_challenge=...)

User->>AS: Authenticates & Approves Consent

AS-->>Client: Redirect to redirect_uri with Auth Code (?code=...)

Client->>AS: POST /token<br/>(grant_type=authorization_code&code=...&code_verifier=...)

Note over AS: Validates code_verifier against code_challenge

AS-->>Client: Returns Access Token (& Refresh Token)

```

  

---
### B. Device Authorization Grant (Device Flow - RFC 8628)

#### Overview

Designed for browserless or input-constrained devices (e.g., Smart TVs, gaming consoles, CLI tools, IoT devices) that cannot present a standard web browser for user authentication.
 
#### Recommended Use Cases

* Smart TVs (Apple TV, Android TV, Samsung Tizen).
* Command Line Tools (AWS CLI, GitHub CLI, Vercel CLI).
* Hardware devices without screens or keyboards.
#### Sequence Diagram

```mermaid

sequenceDiagram

autonumber

actor User as User

participant Device as Device / CLI

participant Browser as User's Browser / Phone

participant AS as Authorization Server

  

User->>Device: Starts Login

Device->>AS: POST /device/code (client_id=...)

AS-->>Device: Returns device_code, user_code, & verification_uri

Device->>User: Displays user_code & verification_uri

par Polling Loop

loop Every X seconds

Device->>AS: POST /token (grant_type=device_code, device_code=...)

AS-->>Device: authorization_pending

end

and User Authorization

User->>Browser: Opens verification_uri & Enters user_code

Browser->>AS: User Authenticates & Grants Access

end

  

Device->>AS: POST /token (grant_type=device_code, device_code=...)

AS-->>Device: Returns Access Token & Refresh Token

```

  

---

### C. Client Credentials Flow (JWT / Secret / mTLS)
#### Overview

Used strictly for **Machine-to-Machine (M2M)** communications where no direct end-user interaction occurs. The application authenticates as itself using its own credentials.
#### Recommended Use Cases

* Backend microservices communicating with each other.
* Automated cron jobs, daemon processes, and background workers.
* Service-to-service API integrations.

Parameters used to send the credentials based on the type of cred
- Client Secret: parameter is client_secret
- Client JWT Assertion: parameter is client_assertion_type and client_assertion
#### Sequence Diagram

```mermaid

sequenceDiagram

autonumber

participant ServiceA as Service A (Client)

participant AS as Authorization Server

  

ServiceA->>AS: POST /token<br/>(grant_type=client_credentials&client_id=...&client_secret=...)

Note over AS: Validates Client Credentials

AS-->>ServiceA: Returns Access Token

```

Example using JWT

```http
POST /oauth/token HTTP/1.1
Host: authorization-server.com
Content-Type: application/x-www-form-urlencoded

grant_type=client_credentials
&client_assertion_type=urn%3Aietf%3Aparams%3Aoauth%3Aclient-assertion-type%3Ajwt-bearer
&client_assertion=eyJhbGciOiJSUzI1NiIs... (signed JWT client credential)
```

Example using client secret

```http
POST /oauth/token HTTP/1.1
Host: authorization-server.com
Content-Type: application/x-www-form-urlencoded
Authorization: Basic Base64(client_id:client_secret)

grant_type=client_credentials
&scope=read:reports
```

Example using mTLS
In this case authentication happens at the Transport Layer. There is no `Authorization` header and no secret. Because nothing in the HTTP message identifies the client, `client_id` must be sent in the body so the server knows which registered client the certificate should match.

```http
POST /oauth2/token HTTP/1.1
Host: mtls.auth.example.com
Content-Type: application/x-www-form-urlencoded

grant_type=client_credentials&client_id=s6BhdRkqt3&scope=api.read
```


---
### D. Token Exchange (RFC 8693)

#### Overview

Defines a standard mechanism for exchanging one security token for another. Essential in microservices architectures to delegate access, propagate user identity across service boundaries, or impersonate users while maintaining auditing.
#### Recommended Use Cases

* Microservice-to-microservice delegation (Service A calling Service B on behalf of an end-user).
* Impersonation (e.g., admin impersonating a user for troubleshooting).
* Converting an external token (e.g., SAML, legacy OAuth token) into an internal JWT.


#### Sequence Diagram

```mermaid

sequenceDiagram

autonumber

actor User as User

participant ServiceA as Service A

participant ServiceB as Service B

participant AS as Authorization Server

  

User->>ServiceA: Initial Request with Subject Token

ServiceA->>AS: POST /token<br/>(grant_type=urn:ietf:params:oauth:grant-type:token-exchange<br/>&subject_token=USER_TOKEN<br/>&subject_token_type=...access_token)

Note over AS: Validates Subject Token & Delegation Policy

AS-->>ServiceA: Returns Downstream Access Token

ServiceA->>ServiceB: API Request with Downstream Token

```

  

**Scenario:** a user calls `api-gateway` with an access token. The gateway needs to call `orders-service` on the user's behalf, so it exchanges the user's token at the authorization server for one scoped to `orders-service`.

#### Step 1: The client calls the API with the user's token

```http
GET /v1/my-orders HTTP/1.1
Host: gateway.example.com
Authorization: Bearer eyJhbGciOiJSUzI1NiIsImtpZCI6IjEyMyJ9.eyJzdWIiOiJ1c2VyLTQyIiwiYXVkIjoiaHR0cHM6Ly9nYXRld2F5LmV4YW1wbGUuY29tIn0.SIG_A
```

The token's audience is the gateway only, so `orders-service` would reject it.
#### Step 2: The gateway sends the token exchange request

```http
POST /oauth2/token HTTP/1.1
Host: auth.example.com
Authorization: Basic Z2F0ZXdheS1jbGllbnQ6czNjcjN0
Content-Type: application/x-www-form-urlencoded
Accept: application/json

grant_type=urn%3Aietf%3Aparams%3Aoauth%3Agrant-type%3Atoken-exchange
&subject_token=eyJhbGciOiJSUzI1NiIsImtpZCI6IjEyMyJ9.eyJzdWIiOiJ1c2VyLTQyIiwiYXVkIjoiaHR0cHM6Ly9nYXRld2F5LmV4YW1wbGUuY29tIn0.SIG_A
&subject_token_type=urn%3Aietf%3Aparams%3Aoauth%3Atoken-type%3Aaccess_token
&requested_token_type=urn%3Aietf%3Aparams%3Aoauth%3Atoken-type%3Aaccess_token
&audience=https%3A%2F%2Forders.example.com
&scope=orders.read
```

### What changed inside the tokens (decoded claims)

**Subject token (aud = gateway):**
```json
{ "iss": "https://auth.example.com", "sub": "user-42",
  "aud": "https://gateway.example.com", "scope": "orders.read profile",
  "exp": 1790000000 }
```

**Exchanged token (aud = orders):**
```json
{ "iss": "https://auth.example.com", "sub": "user-42",
  "aud": "https://orders.example.com", "scope": "orders.read",
  "client_id": "gateway-client", "exp": 1789999400 }
```

The subject stays the same, while audience and scope are narrowed and lifetime is shorter.

### Variant: delegation with `actor_token`

RFC 8693 distinguishes **impersonation** (the new token looks like the user, as above) from **delegation** (the token records that the gateway is acting for the user). For delegation, add an `actor_token`:

```http
POST /oauth2/token HTTP/1.1
Host: auth.example.com
Content-Type: application/x-www-form-urlencoded

grant_type=urn%3Aietf%3Aparams%3Aoauth%3Agrant-type%3Atoken-exchange
&subject_token=<user token>
&subject_token_type=urn%3Aietf%3Aparams%3Aoauth%3Atoken-type%3Aaccess_token
&actor_token=<gateway's own token>
&actor_token_type=urn%3Aietf%3Aparams%3Aoauth%3Atoken-type%3Aaccess_token
&audience=https%3A%2F%2Forders.example.com
```

The issued token then carries an `act` claim:

```json
{ "sub": "user-42",
  "aud": "https://orders.example.com",
  "act": { "sub": "gateway-client" } }
```

Nested `act` claims form a chain if there are multiple hops. This gives the downstream service and its audit logs a clear "who is acting for whom" record.

This example shows a token that shows that there has been two hops across services first gateway-client and later orders-client.

```json
{ 
"iss": "https://auth.example.com", 
"sub": "user-42", 
"aud": "https://inventory.example.com", 
"scope": "inventory.read", 
"act": { 
	"sub": "orders-client", 
	"act": { "sub": "gateway-client" } 
	}, 
	"exp": 1789999400 }
```


#### Scenarios for Token Exchange

**1. Downscoping (attenuation)**  
The client trades a broadly-scoped token for a narrower one to hand to a less-trusted component, such as a browser, a plugin, or a third-party library. The subject and audience often stay the same and only the permissions shrink.

**2. Workload identity federation**  
A platform credential is exchanged for a cloud or API token, so no long-lived secrets are stored. Examples: a GitHub Actions OIDC token for AWS or GCP credentials, ==a **Kubernetes service account token** for a cloud IAM token, or a **SPIFFE SVID** for an OAuth token==. It's arguably the most widely deployed use of token exchange today.

**3. Cross-domain and cross-tenant exchange**  
A token from partner IdP A is exchanged for a token your own authorization server trusts. This can also be cross-region or cross-cloud. The IETF draft on identity and authorization chaining across domains formalizes this, and it's the basis for enterprise "cross-app access" patterns and AI agent scenarios where an agent calls an app in another trust domain.

**4. Step-up and assurance changes**  
After the user completes MFA, the client exchanges its current token for one with a higher assurance level (`acr`/`amr` claims) to call sensitive APIs. The reverse also exists, where a token is downgraded for low-risk use.

**5. Opaque-to-JWT at the edge (phantom token pattern)**  
Clients hold opaque reference tokens, so nothing readable leaks. The gateway exchanges (or introspects) them into JWTs for internal services. It overlaps with your SAML-to-JWT case, but the motivation is privacy and revocation control rather than protocol compatibility.

**6. Sender-constraining**  
A bearer token is exchanged for one bound to a key, either DPoP or mTLS certificate-bound. This shrinks the damage if the token leaks.

**7. Claim minimization or enrichment**

- _Minimization:_ the new token drops PII or gets a pairwise pseudonymous `sub`, so downstream services can't correlate users.
- _Enrichment:_ the authorization server adds roles, entitlements, or tenant context that the original IdP didn't know about.

**8. Refresh token as the subject**  
A refresh token is exchanged for an access token for a different audience, which avoids re-authenticating for each resource. Microsoft's multi-resource refresh token behavior is similar in spirit.

**9. Legacy bridging**  
A Kerberos ticket, session cookie, or API key is exchanged for an OAuth token. This is common in modernization projects, where a gateway or BFF (backend-for-frontend) translates the old credential into a modern one.

---

### E. Implicit Flow (DEPRECATED & INSECURE)

#### Overview

Originally designed for Single Page Applications when browsers did not support Cross-Origin Resource Sharing (CORS). Access tokens were returned directly in the URI fragment (`#access_token=...`).

#### Security Concerns & Deprecation Status

* **DEPRECATED by OAuth 2.1 Security Best Current Practice (BCP)**.
* Access tokens are exposed in browser history, referrer headers, and access logs.
* No mechanism to authenticate the client or verify token origin.
* **Replacement**: Authorization Code Flow with PKCE.


#### Step 1: The browser is sent to the authorization endpoint

```http
GET /oauth2/authorize?response_type=token
    &client_id=spa-client
    &redirect_uri=https%3A%2F%2Fapp.example.com%2Fcallback
    &scope=orders.read
    &state=af0ifjsldkj HTTP/1.1
Host: auth.example.com
```

`response_type=token` is what selects the implicit flow. The user then logs in and consents, and none of that goes to the app.

#### Step 2: The server redirects back with the token in the fragment

```http
HTTP/1.1 302 Found
Location: https://app.example.com/callback#access_token=eyJhbGciOiJSUzI1NiJ9.eyJzdWIiOiJ1c2VyLTQyIn0.SIG_1&token_type=Bearer&expires_in=300&scope=orders.read&state=af0ifjsldkj
```

The part after `#` is never sent in an HTTP request. The browser keeps it locally, and JavaScript on the callback page reads it from `window.location.hash`. The app then checks that `state` matches what it sent.

---
### F. Resource Owner Password Credentials (ROPC) / Password Flow (DEPRECATED)

#### Overview

Requires the user to provide their raw username and password directly to the client application, which then sends them to the authorization server in exchange for a token.  
#### Security Concerns & Deprecation Status

* **DEPRECATED by OAuth 2.1**.
* Requires complete trust in the client application (defeats the core premise of OAuth).
* Exposes user credentials to third-party applications.
* Prevents modern MFA implementations.
* **Replacement**: Authorization Code Flow with PKCE or OIDC.  

Here the client collects the user's username and password itself and sends them straight to the token endpoint.
#### Request

```http
POST /oauth2/token HTTP/1.1
Host: auth.example.com
Authorization: Basic bGVnYWN5LWFwcDpzM2NyM3Q=
Content-Type: application/x-www-form-urlencoded
Accept: application/json

grant_type=password
&username=alice%40example.com
&password=Correct-Horse-9
&scope=orders.read
```

---
## 3. Grant Types & Token Endpoint Parameters

When communicating directly with the `/token` endpoint (via `POST` requests), the `grant_type` parameter defines how credentials or codes are exchanged for tokens.
### Summary Table of Grant Types

| Grant Type                                        | Flow Name                                          | Key Input Parameters                                                                            | Typical Client Type                                                  | Primary Use Case                                                                                                                                       |
| ------------------------------------------------- | -------------------------------------------------- | ----------------------------------------------------------------------------------------------- | -------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------ |
| `authorization_code`                              | Authorization Code                                 | `code`, `redirect_uri`, `code_verifier` (PKCE), client authentication or `client_id`            | Confidential, or Public with PKCE                                    | Web, mobile, and SPA user sign-in                                                                                                                      |
| `client_credentials`                              | Client Credentials                                 | `scope`, client authentication (`client_secret`, `client_assertion`, or mTLS)                   | Confidential                                                         | Machine-to-machine API calls                                                                                                                           |
| `refresh_token`                                   | Refresh Token                                      | `refresh_token`, optional `scope`, client authentication or `client_id`                         | Confidential or Public                                               | Getting a new access token after expiry                                                                                                                |
| `urn:ietf:params:oauth:grant-type:device_code`    | Device Code                                        | `device_code`, `client_id`                                                                      | Public                                                               | Smart TVs, CLI tools, browserless devices                                                                                                              |
| `urn:ietf:params:oauth:grant-type:token-exchange` | Token Exchange (RFC 8693)                          | `subject_token`, `subject_token_type`, optional `actor_token`, `audience` / `resource`, `scope` | Confidential                                                         | Microservice delegation and impersonation, format swaps                                                                                                |
| `urn:ietf:params:oauth:grant-type:jwt-bearer`     | JWT Bearer Assertion (RFC 7523)                    | `assertion` (signed JWT), optional `scope`, client authentication or `client_id`                | Confidential, or Public where the assertion itself is the credential | Trading a JWT from a trusted issuer for an access token: service accounts, federated identity, Microsoft OBO (with `requested_token_use=on_behalf_of`) |
| `urn:openid:params:grant-type:ciba`               | Client-Initiated Backchannel Authentication (CIBA) | `auth_req_id`, client authentication                                                            | Confidential                                                         | Decoupled sign-in where the user approves on a separate device: call centers, point of sale, banking approvals                                         |
| `password` (deprecated)                           | Resource Owner Password (ROPC)                     | `username`, `password`, `scope`                                                                 | Highly trusted, confidential                                         | Legacy migration only                                                                                                                                  |
| _(none)_ `response_type=token` (deprecated)       | Implicit                                           | Sent to `/authorize`, not `/token`: `client_id`, `redirect_uri`, `scope`, `state`               | Public                                                               | Do not use (replace with Authorization Code + PKCE)                                                                                                    |

---
## 4. Key OAuth 2.0 & OIDC Endpoints

An OAuth 2.0 Authorization Server exposes specific endpoints to facilitate authorization, token issuance, session management, and validation.

```mermaid

graph TD

User[User Browser / App] -->|Interactive Login| Auth[/authorize Endpoint/]

User -->|Browserless Login| Device[/device/code Endpoint/]

Auth -->|Auth Code| ClientApp[Client Application]

Device -->|Device Code| ClientApp

ClientApp -->|Exchange Code / Credentials| Token[/token Endpoint/]

Token -->|Access Token| ResourceServer[Resource Server / API Gateway]

ResourceServer -->|Validate Opaque Token| Introspect[/introspect Endpoint/]

ResourceServer -->|Fetch Public Signing Keys| JWKS[/jwks.json Endpoint/]

ClientApp -->|Get User Profile| UserInfo[/userinfo Endpoint/]

ClientApp -->|Revoke Token| Revoke[/revoke Endpoint/]

```

  

### Endpoints Detail

#### 1. Authorization Endpoint (`/authorize`)

* **Purpose**: Used to interact with the resource owner to obtain an authorization grant via redirection.

* **HTTP Method**: `GET` (or `POST` for large request objects).

* **Key Parameters**:

* `response_type`: `code` (for Authorization Code flow).

* `client_id`: Identifies the client requesting access.

* `redirect_uri`: Location to redirect after authorization.

* `scope`: Space-delimited list of requested permissions.

* `state`: Opaque value used to prevent CSRF attacks.

* `code_challenge` & `code_challenge_method`: Used for PKCE verification.

  

#### 2. Token Endpoint (`/token`)

* **Purpose**: Used by the client to exchange an authorization grant, refresh token, or client credentials for access tokens.

* **HTTP Method**: `POST` (`application/x-www-form-urlencoded`).

* **Key Parameters**: `grant_type`, `code`, `redirect_uri`, `client_id`, `client_secret`, `code_verifier`, `refresh_token`.

  

#### 3. Revocation Endpoint (`/revoke` - RFC 7009)

* **Purpose**: Allows a client to notify the authorization server that a previously obtained access or refresh token is no longer needed (e.g., user logs out).

* **HTTP Method**: `POST`.

* **Key Parameters**: `token`, `token_type_hint` (`access_token` or `refresh_token`).

  

#### 4. Introspection Endpoint (`/introspect` - RFC 7662)

* **Purpose**: Allows Resource Servers to query the Authorization Server to determine the active state and metadata of an opaque token.

* **HTTP Method**: `POST` (Requires Resource Server authentication).

* **Response**: JSON object containing `active` (boolean), `scope`, `client_id`, `exp`, `sub`, etc.

  

#### 5. UserInfo Endpoint (`/userinfo` - OpenID Connect Standard)

* **Purpose**: Returns claims/attributes about the authenticated end-user (requires a valid OIDC Access Token).

* **HTTP Method**: `GET` or `POST` with `Authorization: Bearer <access_token>`.

* **Response**: JSON object containing user identity attributes (e.g., `sub`, `email`, `name`).

  

#### 6. Device Authorization Endpoint (`/device/code` - RFC 8628)

* **Purpose**: Initiates the device flow by returning a user code, device code, and verification URL.

* **HTTP Method**: `POST`.

  

#### 7. Discovery Endpoint (`/.well-known/openid-configuration` / `/.well-known/oauth-authorization-server`)

* **Purpose**: Metadata endpoint returning JSON configuration details listing all supported endpoints, scopes, response types, signing algorithms, and keys.

  

#### 8. JWKS Endpoint (`/jwks.json` or `/keys`)

* **Purpose**: Publishes the public keys (JSON Web Key Set) used by the Authorization Server to sign JWTs. Used by Resource Servers to perform local signature verification.

  

---

  

## 5. Critical Security Concepts for System Design Interviews

  

### 1. PKCE (Proof Key for Code Exchange - RFC 7636)

* **Problem**: Authorization codes sent via browser redirects can be intercepted by malicious apps on mobile devices or via browser extension sniffing.

* **Solution**:

1. Client creates a random string: `code_verifier`.
2. Client hashes it: `code_challenge = SHA256(code_verifier)`.
3. Send `code_challenge` in `/authorize` request.
4. Send `code_verifier` in `/token` request.
5. Authorization Server verifies `SHA256(code_verifier) == code_challenge`.
* **Rule**: Mandatory for public clients; strongly recommended for all clients under OAuth 2.1.

### 2. State Parameter vs. PKCE

* **`state`**: Prevents **Cross-Site Request Forgery (CSRF)** on the authorization endpoint. Ensures the user who initiated login is the same user completing it.
* **`PKCE`**: Prevents **Authorization Code Injection / Interception**. Ensures the app exchanging the authorization code is the app that requested it.

### 3. Opaque Tokens vs. JWTs (JSON Web Tokens)

| Feature          | Opaque Tokens                                      | JWTs (Structured Tokens)                                 |
| :--------------- | :------------------------------------------------- | :------------------------------------------------------- |
| **Structure**    | Random string (e.g., `db3a8f...`)                  | Base64URL encoded header, payload, signature             |
| **Validation**   | Requires network call (`/introspect`) or DB lookup | Self-contained; verified offline using public key (JWKS) |
| **Revocation**   | Instant (server invalidates entry)                 | Harder (requires short TTLs or revocation lists)         |
| **Payload Size** | Minimal (~32-64 bytes)                             | Larger (0.5KB - 4KB+)                                    |
### 4. Refresh Token Rotation

To mitigate the risk of stolen refresh tokens, Authorization Servers implement **Refresh Token Rotation**:

* Every time a refresh token is used, it is invalidated and a **new** refresh token is issued alongside the new access token.
* If a revoked/previously used refresh token is presented, the Authorization Server detects potential theft and invalidates **all tokens** associated with that client/user session.

---

## 6. Summary Cheat Sheet for System Design Interviews

| Use Case / Scenario                    | Recommended Flow & Pattern                                                                                                                                                         |
| :------------------------------------- | :--------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| **Mobile App / SPA Architecture**      | Use **Authorization Code Flow + PKCE**. Never store secrets in client code.                                                                                                        |
| **Microservices Communication**        | Use **Client Credentials Flow** (with Private Key JWT or mTLS) for service-to-service calls. Use **Token Exchange (RFC 8693)** when propagating user context.                      |
| **Smart TV / CLI Tools**               | Use **Device Authorization Grant**.                                                                                                                                                |
| **Stateless API Gateway Verification** | Issue **JWTs**, configure API Gateways to fetch public keys from `/jwks.json`, and cache public keys locally for stateless offline validation.                                     |
| **Instant Revocation Requirement**     | Use **Opaque Tokens** with an introspection endpoint or cache-backed store (Redis), or use **short-lived JWTs (e.g., 5-15 mins)** paired with a revocation blacklisting mechanism. |

---

## Provisioning of mTLS as a credential for the client credentials flow

For MTLS there are actually two distinct provisioning models, and they pre-register different things. 

- In the self-signed variant, the client generates its own keypair and a self-signed certificate, and what gets registered with the server ahead of time isn't the whole certificate but a **thumbprint**, a hash of the public key, or sometimes the JWKS containing that key. The server never validates a certificate chain at all here, **it just checks whether the presented cert's thumbprint matches what's on file**.
- In the PKI based variant, it's different again — the server pre-registers an expected subject distinguished name, and trusts a whole certificate authority rather than one specific certificate. So the client can even rotate its certificate freely, as long as the new one is issued by that trusted CA and **carries the same subject**.