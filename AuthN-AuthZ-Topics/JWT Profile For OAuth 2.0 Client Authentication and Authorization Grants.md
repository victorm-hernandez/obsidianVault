
https://www.rfc-editor.org/info/rfc7523/

RFC 7523 (_JSON Web Token (JWT) Profile for OAuth 2.0 Client Authentication and Authorization Grants_) defines how to use a signed JSON Web Token (JWT) as a cryptographic credential when interacting with an OAuth 2.0 Authorization Server.  

It profiles the generic OAuth 2.0 Assertion Framework (RFC 7521) specifically for JWTs, enabling passwordless and secret-free authorization flows.  

## Key Concepts  

RFC 7523 defines two independent (and combinable) capabilities:  

1. JWT as an Authorization Grant  
    Purpose: Exchanging a trusted JWT (assertion) directly for an OAuth 2.0 Access Token without requiring interactive user approval.  
    Use Case: Server-to-server integrations, service accounts, or cross-domain token exchanges where a trust relationship already exists.  
    
2. JWT for Client Authentication (e.g., private_key_jwt)  
    Purpose: Authenticating the OAuth client application itself using a digitally signed JWT instead of a static shared `client_secret`.  
    Use Case: High-security applications where asymmetric key pairs (RSA/ECDSA) are preferred over storing symmetric shared secrets.  
    

## HTTP Request Specification  

1. Authorization Grant  
To present a JWT as a grant, the client sends a `POST` request to the token endpoint with the specified URN parameter values:  

```http
POST /token HTTP/1.1
Host: server.example.com
Content-Type: application/x-www-form-urlencoded

grant_type=urn%3Aietf%3Aparams%3Aoauth%3Agrant-type%3Ajwt-bearer
&assertion=eyJhbGciOiJFUzI1NiIs...
```

  
2. Client Authentication  
To authenticate using a JWT while exchanging another grant type (such as an `authorization_code`), the client adds `client_assertion_type` and `client_assertion`:  

```http
POST /token HTTP/1.1
Host: server.example.com
Content-Type: application/x-www-form-urlencoded

grant_type=authorization_code
&code=SplxlOBeZQQYbYS6WxSbIA
&client_assertion_type=urn%3Aietf%3Aparams%3Aoauth%3Aclient-assertion-type%3Ajwt-bearer&client_assertion=eyJhbGciOiJSUzI1NiIs...
```

  
Core Requirements & JWT Claims  
To be accepted by the Authorization Server, the JWT must be digitally signed (or MACed) and contain specific valid claims:  

|Claim|Name|Description & Requirement|
|---|---|---|
|iss|Issuer|Identifies the entity that issued the JWT (e.g., the client ID or an external identity provider).|
|sub|Subject|Identifies the principal.  <br>• For Authorization Grants: The resource owner, user ID, or authorized entity.  <br>• For Client Authentication: Must be the OAuth client_id.|
|aud|Audience|Must identify the intended Authorization Server (typically its Token Endpoint URL). Assertions with mismatched audiences are rejected.|
|exp|Expiration Time|Unix timestamp limiting the window of validity. Expired tokens are rejected.|
|nbf / iat|Not Before / Issued At|(Optional) Specifies when the token becomes active or when it was generated.|
|jti|JWT ID|(Optional) Unique token identifier used by the Authorization Server to prevent replay attacks.|

Security Considerations  

- Asymmetric Security: Reduces the vulnerability of static `client_secret` leaks by utilizing public/private key pairs.      
- Replay Protection: Authorization Servers are advised to track the `jti` (JWT ID) claim until expiration to ensure a token cannot be submitted twice.  
- Strict Audience Matching: The `aud` check ensures a JWT generated for one server cannot be intercepted and replayed against a different Authorization Server.

---

## What is an "Authorization Grant" in OAuth?

In OAuth 2.0, an Authorization Grant is an intermediate credential that represents the Resource Owner's (user's) permission given to the client. It is used by the client application to obtain an Access Token from the Authorization Server.


Think of it like a claim ticket:  

- The user approves access
- The Authorization Server gives the client a claim ticket (Authorization Grant).  
- The client takes the claim ticket to the token endpoint
- The server exchanges it for an Access Token (the actual key to the API).  

Examples of Authorization Grants:  

- Authorization Code: A temporary, short-lived code returned after a user logs in via a web browser.  
    
- JWT Bearer (RFC 7523): A signed JWT asserting that a user/service authorized this client.  
    
- Refresh Token: A long-lived credential used to request new access tokens without prompting the user.


## Client Authentication vs. Authorization Grant

Let's distinguish between Who is making the request versus What permission is being presented.  

```
           +--------------------------------------------------------+           |                OAuth Token Request                     |           +--------------------------------------------------------+           |                                                        |           |  1. WHO ARE YOU? (Client Authentication)               |           |     "I am Client Application 'App_123'."               |           |     (e.g., client_secret OR client_assertion JWT)      |           |                                                        |           |  2. WHAT PERMISSION ARE YOU SHOWING? (Authorization Grant) |           |     "User 'Alice' gave me permission to access data." |           |     (e.g., code=xyz OR grant_type=jwt-bearer)          |           |                                                        |           +--------------------------------------------------------+
```

How RFC 7523 Uses Both  

RFC 7523 allows a JSON Web Token (JWT) to fulfill either (or both) of these roles, which is why it often confuses people:  

Scenario A: JWT as Client Authentication (`private_key_jwt`)  

- Goal: Prove the client's identity without sending a static password (`client_secret`).  
- Request Payload:  
    `grant_type=authorization_code` _(I have a standard user approval code)_  
    `code=splx_123`  
    `client_assertion_type=...:jwt-bearer` _(I am proving my client identity via a signed JWT)_  
    `client_assertion=eyJhbG...` _(JWT signed by the client's private key)_  
    

Scenario B: JWT as an Authorization Grant (`jwt-bearer`)  
- Goal: Use a signed JWT directly as permission to get an access token (no interactive browser login required).  
- Request Payload:  
    `grant_type=...:jwt-bearer` _(My authorization proof IS this JWT)_  
    `assertion=eyJhbG...` _(JWT representing user delegation or service-to-service trust)_