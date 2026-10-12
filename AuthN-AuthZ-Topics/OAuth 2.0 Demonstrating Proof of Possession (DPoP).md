[https://www.rfc-editor.org/info/rfc8693/](https://www.rfc-editor.org/info/rfc9449/)

RFC 9449 specifies OAuth 2.0 Demonstrating Proof-of-Possession at the Application Layer (DPoP). It is an IETF Standards Track document designed to strengthen OAuth 2.0 security by preventing token theft and replay attacks.  
## Key Takeaways  

### 1. The Core Problem: Bearer Tokens are Vulnerable  
Standard OAuth 2.0 access tokens and refresh tokens are traditionally bearer tokens: anyone who holds or captures the token can use it. If an attacker steals a token via cross-site scripting (XSS), server logging leaks, or network interception, they gain unauthorized access to protected resources.  
### 2. The Solution: Sender-Constrained Tokens  
DPoP solves this by sender-constraining tokens at the application layer:  
- When requesting a token, the client presents a freshly signed DPoP proof JWT created with its asymmetric private key.  
- The Authorization Server (AS) binds the issued access token (and refresh token, if applicable) to the corresponding public key.  
- When making requests to a Resource Server (RS), the client must present the access token AND a new signed DPoP proof for that specific HTTP request. 
- Even if an attacker steals the access token, it is useless without access to the client’s private key.  


How DPoP Works  
```mermaid
sequenceDiagram
    autonumber
    participant Client
    participant AS as Authorization Server
    participant RS as Resource Server

    Note over Client: 1. Generate asymmetric key pair<br/>(Public/Private Key)

    rect rgb(240, 244, 248)
        Note over Client, AS: Token Request (Issuance)
        Client->>AS: POST /token<br/>Header: DPoP: <DPoP Proof JWT signed with Private Key>
        Note over AS: Verify DPoP signature & claims<br/>Bind token to Public Key (cnf claim)
        AS-->>Client: 200 OK<br/>{ "access_token": "...", "token_type": "DPoP" }
    end

    rect rgb(240, 248, 240)
        Note over Client, RS: Resource Request (Usage)
        Client->>RS: GET /api/user<br/>Authorization: DPoP <Access Token><br/>Header: DPoP: <New DPoP Proof JWT with ath claim>
        Note over RS: 1. Extract Public Key bound to Token<br/>2. Validate DPoP signature & HTTP claims (htm, htu, ath)
        RS-->>Client: 200 OK (Protected Resource Response)
    end

```

1. Token Issuance:  
    The client generates an ephemeral or stored asymmetric key pair (e.g., EC or RSA). During the token request, it sends a signed `DPoP` HTTP header. The AS attaches the public key thumbprint to the issued token (usually as a `cnf` / confirmation claim).  
    
2. Resource Request:  
    When accessing a protected API, the client passes the token in the `Authorization: DPoP <token>` header alongside a new `DPoP: <DPoP JWT Proof>` header.  
    
3. Verification:  
    The Resource Server verifies that:  
    The signature on the DPoP JWT matches the public key bound to the access token.  
    The DPoP JWT contains claims matching the HTTP method, request URI, timestamp, unique ID (`jti`), and a hash of the access token (`ath`).  

## What's Inside a DPoP Proof JWT?  
A DPoP proof is a standard Signed JWT with specific header parameters and claims:  

- Header:  
    `typ`: MUST be `dpop+jwt`.  
    `alg`: Asymmetric signature algorithm (e.g., `ES256`, `RS256`).  
    `jwk`: The public key corresponding to the private key used to sign the proof.  
    
- Payload Claims:  
    `jti`: Unique identifier for the JWT (prevents replay attacks).  
    `htm`: The HTTP method of the request (e.g., `GET`, `POST`).  
    `htu`: The HTTP target URI without query or fragment parameters.  
    `iat`: Issue timestamp.  
    `ath`: Hash of the access token (included on resource requests to bind the proof directly to the token).  
    `nonce`: _(Optional)_ Server-provided challenge string to combat pre-computed proof replay attacks.  

## Why DPoP vs. Alternatives (mTLS / Token Binding)?  

- Application-Layer Operation: Unlike mTLS (RFC 8705), DPoP operates entirely in HTTP application headers. It does not require special TLS infrastructure or client certificate management.  
    
- Single-Page Applications (SPAs): SPAs running in browser sandboxes cannot manage TLS client certificates or use mTLS, but they can generate non-extractable keys using the browser's Web Crypto API to create DPoP proofs.  
    
- Mobile & Native Apps: Works cleanly across native apps where hardware-backed key storage (e.g., Secure Enclave, Android Keystore) is available.


### DPoP vs. JWT Client Assertion (`private_key_jwt`)  
DPoP and the JWT Profile for OAuth Client Authentication (RFC 7523 / private_key_jwt) serve distinct security roles.

| Feature                 | private_key_jwt (Client Authentication)                           | DPoP (Proof of Possession)                                           |
| ----------------------- | ----------------------------------------------------------------- | -------------------------------------------------------------------- |
| Purpose                 | Authenticate who the client is during token requests.             | Prove ownership of the token on every API call.                      |
| Where it is used        | Only at the Authorization Server (/token endpoint).               | At both the Token Endpoint AND every API call (Resource Server).     |
| Protects                | The client registration identity.                                 | The access token and refresh token against theft/replay.             |
| Key Registration        | Client public key must be pre-registered/provisioned with the AS. | Key is declared on-the-fly inside the JWT header (jwk claim).        |
| Public vs. Confidential | Only for Confidential Clients (backends).                         | Works for both Public Clients (SPAs, Mobile) & Confidential Clients. |

## HTTP Request Flow & Wire Protocol Example  

### Step 1: Token Request to Authorization Server (AS)  

HTTP Request:  

```http
POST /oauth/v2/token HTTP/1.1
Host: auth.example.com
Content-Type: application/x-www-form-urlencoded
DPoP:eyJhbGciOiJFUzI1NiIsInR5cCI6ImRwb3Arand0Iiwid2siOnsia3R5IjoiRUMiLCJjcnYiOiJQLTI1NiIsIngiOiJmODNKLWFsTFQiLCJ5IjoiWTBiM3ptS3AifX0.eyJqdGkiOiItQjR0S2Z2R211RiIsImh0bSI6IlBPU1QiLCJodHUiOiJodHRwczovL2F1dGguZXhhbXBsZS5jb20vb2F1dGgvdjIvdG9rZW4iLCJpYXQiOjE3MTAwMDAwMDB9.SignatureValue

grant_type=authorization_code
&code=SplxlOBeZQQYbYS6WxSbIA
&client_id=my-spa-client
```


Decoded DPoP Proof (Header & Payload sent to AS):  

```json
// JOSE Header
{  
"typ": "dpop+jwt",  
"alg": "ES256",  
"jwk": {    "kty": "EC",    "crv": "P-256",    "x": "f83J-alLT...",    "y": "Y0b3zmKp..."  }
}

// Payload
{  
"jti": "-B4tKfvGmuF",  
"htm": "POST",  
"htu": "https://auth.example.com/oauth/v2/token",  
"iat": 1710000000
}
```
  
AS Response (DPoP-Bound Access Token Returned):  

```http
HTTP/1.1 200 OK
Content-Type: application/json

{  
"access_token": "eyJhbGciOiJSUzI1NiI...",  
"token_type": "DPoP",  
"expires_in": 3600,  
"refresh_token": "r-8x239n10..."
}
```

Decoded Access Token Payload (Issued by AS):  
Notice the cnf (Confirmation) claim containing the JWK SHA-256 thumbprint (`jkt`):  

```json
{  
"iss": "https://auth.example.com",  
"sub": "user_12345",  
"aud": "https://api.example.com",  
"exp": 1710003600,  
"cnf": {    "jkt": "0Z3we0A01FFRByR-2M2_q3XS1S33D10-9S234_ABC"  }
}
```

  
### Step 2: Requesting Protected Resource from Resource Server (RS)  
To call the API, the client passes both the Access Token and a brand-new DPoP Proof that includes the hash of the access token (`ath`).  
HTTP Request:  

```http
GET /api/user/profile HTTP/1.1
Host: api.example.com
Authorization: DPoP eyJhbGciOiJSUzI1NiI...
DPoP: eyJhbGciOiJFUzI1NiIsInR5cCI6ImRwb3Arand0Iiwid2siOn...
```

  
Decoded DPoP Proof Payload (Sent to RS):  

```json
{ 
 "jti": "8f870234-a123",  
 "htm": "GET",  
 "htu": "https://api.example.com/api/user/profile",  
 "iat": 1710000005,  
 "ath": "fU851BJ3K5QuM-384910385910358102384102934" 
 }
```

  
_(Note:_ _`ath`_ _is the base64url-encoded SHA-256 hash of the Access Token string)._

## Key Lifecycle & Pre-Provisioning Scenarios  

Is the key generated for every token request, or is it once in a lifetime?  
Neither extreme is required; the RFC allows flexibility:  

- Per-Session Key (Most Common for SPAs/Browsers): A client generates an in-memory key pair via `window.crypto.subtle` when the app loads. It uses this key pair for its token request and all subsequent API requests during that browser session.  
    
- Persistent Key (Common for Mobile Apps/Native Clients): A key pair is generated upon initial installation, stored in hardware-backed secure storage (iOS Secure Enclave / Android Keystore), and reused for months or years across token refreshes.  
    
- Per-Request Key (Rare): A client could theoretically generate a new key pair per token request, but it must reuse that same key for all resource calls using the access token issued under it.
## What is the TTL of a DPoP Proof JWT?  

- Short-lived window (typically 60 to 300 seconds): DPoP proofs usually omit an `exp` claim and rely solely on `iat` (issued-at time). The receiving server verifies that `iat` is close to the current time to limit replay windows.  
    
- Nonce-Based TTL (Server Challenge): If an RS wants strict control, it returns a `DPoP-Nonce` header in a `401 Unauthorized` response. The client must then regenerate the DPoP proof with that exact `nonce` claim and retry. This turns the DPoP proof into a single-use token with instantaneous expiry controlled by the server.