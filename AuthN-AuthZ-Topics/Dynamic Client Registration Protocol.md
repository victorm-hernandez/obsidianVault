https://www.rfc-editor.org/info/rfc7591/

RFC 7591 defines the OAuth 2.0 Dynamic Client Registration Protocol.  

It standardizes a programmatic mechanism for OAuth 2.0 clients (or developers) to register themselves with an Authorization Server (AS) on the fly, replacing manual client onboarding (such as filling out forms in developer consoles).  

## Core Problem It Solves  

In standard OAuth 2.0 (RFC 6749), every application must obtain a `client_id` and optional `client_secret` before making requests. Manually setting this up works for centralized applications, but breaks down when:  

- Mobile/Native Apps & Installs: Every downloaded instance or local deployment needs its own distinct credentials. 
    
- Open Ecosystems: Applications need to connect automatically to hundreds of independent Authorization Servers (e.g., OpenID Connect IdPs, healthcare systems using SMART-on-FHIR, or federated enterprise networks) without prior manual developer registration on each one. 

## Key Components  

1. Client Registration Endpoint  
An HTTP POST endpoint exposed by the Authorization Server (e.g., `/register`).  

- Open Registration: Anyone can send a request to register.  
- Protected Registration: Requires an Initial Access Token in the `Authorization` header to restrict who can register clients.  

2. Client Metadata  
During registration, the client sends a JSON payload describing its attributes, capabilities, and requirements:  

- redirect_uris (Required): List of allowed redirection endpoints.  
- client_name, client_uri, logo_uri: Human-readable metadata displayed during user consent prompts.  
- grant_types: Requested grant types (e.g., `authorization_code`, `refresh_token`, `urn:ietf:params:oauth:grant-type:jwt-bearer`).  
- response_types: Requested response parameters (e.g., `code`).  
- token_endpoint_auth_method: How the client will authenticate at the token endpoint (e.g., `client_secret_basic`, `private_key_jwt`, `none`).  

Example JSON payload with explicit JWK document

```json
{
  "redirect_uris": ["https://client.example.org/callback"],
  "client_name": "My Secure App",
  "token_endpoint_auth_method": "private_key_jwt",
  "jwks": {
    "keys": [
      {
        "kty": "RSA",
        "e": "AQAB",
        "use": "sig",
        "kid": "key-2026-1",
        "n": "u1W3b...[base64url encoded public key]..."
      }
    ]
  }
}
```

Example JSON with JWKS_uri (URL pointing to keys)

```json
{
  "redirect_uris": ["https://client.example.org/callback"],
  "client_name": "My Secure App",
  "token_endpoint_auth_method": "private_key_jwt",
  "jwks_uri": "https://client.example.org/.well-known/jwks.json"
}

```

3. Software Statements  
To prevent abuse during open registration, clients can include a Software Statement—a signed JSON Web Token (JWT) issued by a trusted third party (such as an app store, developer registry, or software publisher).  

- The JWT asserts metadata values about the client software.  
- The AS validates the JWT signature to ensure the client metadata hasn't been tampered with and comes from an authorized source.  

## Registration Flow  

1. Client Request: The client sends an HTTP POST request containing its metadata JSON to the Client Registration Endpoint:  
    

```http
POST /register HTTP/1.1
Host: server.example.com
Content-Type: application/json
Authorization: Bearer initial_access_token_if_required

{  
"redirect_uris": ["https://client.example.org/callback"],
"client_name": "My App",  
"grant_types": ["authorization_code", "refresh_token"],
"token_endpoint_auth_method": "client_secret_basic"
}
```

    
2. Server Response: The Authorization Server processes the request, assigns credentials, and returns a `201 Created` status with the full registered client configuration:  
    

```http
HTTP/1.1 201 Created
Content-Type: application/json
{  
"client_id": "s6BhdRkqt3",  
"client_secret": "cf13605b380c70e2cd21b1ad",  
"client_id_issued_at": 1577836800,  
"client_secret_expires_at": 0,  
"redirect_uris": ["https://client.example.org/callback"],
"grant_types": ["authorization_code", "refresh_token"]
}
```

1. Subsequent Calls: The client uses the returned `client_id` and `client_secret` to perform standard OAuth 2.0 flows.  

## Primary Use Cases  

- OpenID Connect (OIDC): Formally incorporated as part of the core OpenID Connect Dynamic Client Registration 1.0 specification.  
    
- Healthcare (SMART on FHIR): Allows patient-facing mobile apps to connect to any hospital's EHR system dynamically.  
    
- User-Managed Access (UMA): Enables automated authorization setup between independent resource management apps.  
    
- Federated Networks: Multi-tenant enterprise systems where new software instances register automatically upon installation or deployment.

---

Dynamic Client Registration (RFC 7591) is a critical protocol for autonomous AI agents because it enables them to programmatically discover, establish identity with, and obtain authorized access to arbitrary third-party services and Model Context Protocol (MCP) servers without human developer intervention.  

## Why AI Agents Need Dynamic Client Registration  

1. Zero-Touch Onboarding to Unknown Tools  
When a user asks an AI agent to complete a task (e.g., _"Book a meeting on my calendar using Service X"_), the agent may discover a new API or tool at runtime (e.g., via Open API / OpenAPI specs or Model Context Protocol discovery).  

- Without RFC 7591: A developer must manually log into Service X's developer console, create an application, copy the `client_id` and `client_secret`, and inject them into the agent's environment variables.  
    
- With RFC 7591: The agent calls the authorization server's `/register` endpoint, sends its metadata (including capabilities and redirect URIs), receives a `client_id` instantly, and initiates the authentication flow on the fly.  

2. Short-Lived Ephemeral Agent Instances 
Modern AI agents are increasingly designed as ephemeral micro-agents spun up for a single task and destroyed immediately after.  

- Static client keys stored in long-lived configuration files pose a major security risk.  
    
- Dynamic registration allows an AI agent to create a unique, temporary client_id bound exclusively to that specific task instance, then revoke or delete it (via RFC 7592) when the job is done.  

3. Integration with Model Context Protocol (MCP)  
In ecosystems using protocols like MCP (where AI hosts interact with specialized tool servers):  

- An AI assistant host needs to authenticate against dozens of distributed MCP servers managed by different vendors.  
    
- RFC 7591 provides a standard, interoperable mechanism for the agent to register as an OAuth client with each MCP server's authorization server automatically.  

4. Cryptographic Identity & Software Statements
Allowing agents to register open-ended client configurations poses security risks if malicious scripts abuse registration endpoints.  

- RFC 7591 solves this by using Software Statements—digitally signed JWTs issued by the agent’s platform or vendor.  
    
- An AI agent presents its signed Software Statement to the target API. The target API verifies that the agent was built by a trusted software publisher (e.g., OpenAI, Anthropic, or an enterprise internal registry) before issuing credentials.  

## Combining RFC 7591 with ID-JAG (Draft Spec)  

When combined with the Identity Assertion Authorization Grant (ID-JAG) from your previous spec:  

1. RFC 7591 (Registration): The AI Agent dynamically registers itself with a target third-party Resource Authorization Server to receive a local `client_id`.  
    
2. ID-JAG (Authorization): The AI Agent obtains a signed identity assertion from its central Enterprise IdP and exchanges it with the target Resource AS for a user-scoped access token.  
    
This combination enables fully automated, end-to-end, zero-friction delegation for AI agents across domain boundaries.