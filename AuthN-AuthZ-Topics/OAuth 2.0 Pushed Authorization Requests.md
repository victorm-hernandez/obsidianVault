
https://www.rfc-editor.org/info/rfc9126/

RFC 9126: OAuth 2.0 Pushed Authorization Requests (PAR) defines a mechanism for clients to send authorization request parameters directly to the Authorization Server (AS) via a back-channel HTTP POST request, rather than passing them through the user's browser via front-channel query parameters.  
Key Takeaways & Workflow  

1. The Core Problem Solved  
In standard OAuth 2.0 (RFC 6749), the authorization request payload is appended directly to the URL string during browser redirection (`GET /authorize?client_id=...&scope=...`). This introduces several risks:  

- Integrity/Authenticity Risk: Parameters sent through the front channel can be manipulated or intercepted by a malicious user agent or attacker.  
    
- Confidentiality/Privacy Leaks: Query parameters easily leak into web server logs, browser history, and HTTP `Referer` headers (particularly problematic in Open Banking, healthcare, or when sharing PII).  
    
- URL Length Limits: Complex, fine-grained access requests or long JSON Web Tokens (JWTs) can exceed maximum URL lengths allowed by browsers and proxies.  
    

2. The PAR Flow  
Instead of building a long URL and redirecting the browser immediately, the client uses a two-step process:  

```
+--------+                                 +---------------+|        |--(1) POST authorization request->| Pushed Auth   ||        |      payload directly (Back-Ch)  | Req Endpoint  ||        |<-(2) Returns "request_uri" ------|               || Client |      & "expires_in"             +---------------+| App    |                                 | Authorization ||        |--(3) Redirect user agent with--->| Endpoint      ||        |      "request_uri" (Front-Ch)   | (Browser Flow)|+--------+                                 +---------------+
```

  

1. Push Payload (Back-Channel): The client makes an HTTP `POST` request to the authorization server's dedicated PAR endpoint (`/as/par`) containing all request parameters (`response_type`, `scope`, `redirect_uri`, `PKCE parameters`, etc.).  
    
2. Receive Request URI: The server validates the payload and client credentials, stores the parameters securely, and returns a JSON payload containing:  
    `request_uri`: A short, opaque reference string (e.g., `urn:ietf:params:oauth:request_uri:...`).  
    `expires_in`: Time-to-live for the reference (typically short, e.g., 60–90 seconds).  
    
3. Redirect User (Front-Channel): The client redirects the user’s browser to the standard Authorization Endpoint (`/authorize`) passing only two parameters: `client_id` and the `request_uri`.  
    

Core Technical Benefits  

- Early Client Authentication: Because the request goes directly to the AS over the back-channel, the server can authenticate the client _before_ initiating any user interaction. This prevents client spoofing or unauthorized authorization attempts early in the process.  
    
- Prevents Parameter Tampering: Attackers cannot modify parameters like `scope`, `redirect_uri`, or transaction values while in transit through the browser.  
    
- Enforces Short-Lived References: The `request_uri` is bound to a single client and expires quickly, reducing window of opportunity for replay attacks.  
    
- Works Seamlessly with Native/Mobile Apps: Unlike cross-site HTML form POST workarounds (which fail or cause UX friction on mobile OS native browser tabs), PAR works cleanly with standard browser redirects (`GET`).  
    
- Complements JWT-Secured Authorization Requests (JAR / RFC 9101): Clients requiring non-repudiation can still wrap request parameters inside a signed JWT (Request Object) and push that object directly to the PAR endpoint.  
    

Server Discovery Metadata  
Authorization Servers supporting PAR publish the endpoint in their OpenID Connect or OAuth 2.0 Discovery metadata (`/.well-known/openid-configuration` or `/.well-known/oauth-authorization-server`):  

```json
{  "pushed_authorization_request_endpoint": "https://as.example.com/as/par",  "require_pushed_authorization_requests": true}
```