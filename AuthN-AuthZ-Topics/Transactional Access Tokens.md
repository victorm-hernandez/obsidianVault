
==Very new, proposal issued on September 29 2026==
https://datatracker.ietf.org/doc/draft-ietf-oauth-transaction-tokens/
https://www.ietf.org/archive/id/draft-tulshi-oauth-transactional-access-tokens-00.html

`draft-tulshi-oauth-transactional-access-tokens-00` (published September 2026 by A. Tulshibagwale) defines Transactional Access Tokens (Txn-ATs)—a profile of JSON Web Token (JWT) access tokens designed to extend OAuth 2.0 with per-transaction authorization and contextual metadata.  
Key Concepts & Mechanics  

## 1. Problem Addressed  

Traditional OAuth Access Tokens: Issued for whole sessions. They convey _what_ a client may do (via `scope`), but carry no context about _why_ or in _which specific transaction_ an action occurs. Once issued, a client holding credentials can obtain new tokens without the Authorization Server (AS) evaluating policy for every single transaction.  

Txn-AT Solution: Shifts OAuth from session-level authority to per-transaction authority. The AS evaluates fresh policy decisions for every individual transaction—making it ideal for AI agents performing task-scoped activities.  
    
## 2. Core Properties of a Txn-AT  

JWT Profile with Distinct Type: Extends standard RFC 9068 JWT access tokens, but uses `typ: "txnat+jwt"` in the JOSE header to prevent legacy resource servers (RS) from accepting them as ordinary, long-lived access tokens.  

Very Short Lifetime: The difference between `exp` (expiration) and `iat` (issued at) should not exceed 300 seconds (5 minutes).  

Strict Audience Restriction: Every Txn-AT is bound to a single, specific resource server string (`aud`).  

Transaction Identifier (txn): Contains a transaction ID. To prevent cross-resource 
tracking/correlation, the AS issues distinct, unlinkable `txn` values to different resource servers for the same transaction.  

Transaction Context (tctx & rctx): Carries AS-evaluated context parameters (e.g., risk signals, subject attributes, task boundaries) explaining _why_ the transaction was approved.  
    
## 3. Transaction Flow & Handles  
Initiating a Transaction: The client requests a Txn-AT for a resource server (RS1) using standard token endpoints (e.g., RFC 8693 Token Exchange, Refresh Token Grant, or Client Credentials Grant).  

Transaction Handle (txn_handle): Upon approving a transaction, the AS issues the initial Txn-AT alongside an opaque `txn_handle`.  

Multi-Resource Workflows: If the client needs to interact with another resource server (RS2) as part of the _same_ overall transaction, it presents the `txn_handle` to the AS. The AS evaluates policy again and issues a new Txn-AT scoped specifically to RS2.  
    
## 4. Integration with OAuth Transaction Tokens (Txn-Tokens)  

Txn-ATs align with the semantics of OAuth Transaction Tokens (used _within_ trust domains).  

When an external client presents a Txn-AT to a Resource Server, that RS can translate or carry the AS-asserted transaction context and `txn` ID into its internal microservice network using Txn-Tokens.  
    

## Key Benefits  

- Task-Scoped AI Agent Control: Allows organizations to safely delegate tasks to autonomous AI agents by requiring fresh, per-transaction authorization.  
    
- Context Preservation: Downstream services receive verifiable context about the transaction rather than just static broad permissions.  
    
- Privacy & Anti-Correlation: Different resource servers involved in the same transaction cannot track or correlate requests using the `txn` identifier.


---
## 1. HTTP Token Request Comparison  

### Standard OAuth 2.0 Token Exchange Request  
A client exchanges an existing token for a standard access token for a target resource (`RS1`).  

```http
POST /token HTTP/1.1
Host: as.example.com
Content-Type: application/x-www-form-urlencoded

grant_type=urn%3Aietf%3Aparams%3Aoauth%3Agrant-type%3Atoken-exchange
&client_id=agent_app_123&subject_token=eyJhbGciOiJSUzI1Ni...
&subject_token_type=urn%3Aietf%3Aparams%3Aoauth%3Atoken-type%3Aaccess_token&resource=https%3A%2F%2Frs1.example.com%2F
&scope=read%20write
```

  
### Transactional Access Token (Txn-AT) Token Request  
When requesting a Txn-AT, the client can pass explicit transactional parameters (like `authorization_details`) or reuse a Transaction Handle (`txn_handle`) to chain requests across multiple services under one transaction.  

```http
POST /token HTTP/1.1
Host: as.example.com
Content-Type: application/x-www-form-urlencoded

grant_type=urn%3Aietf%3Aparams%3Aoauth%3Agrant-type%3Atoken-exchange
&client_id=agent_app_123&subject_token=eyJhbGciOiJSUzI1Ni...
&subject_token_type=urn%3Aietf%3Aparams%3Aoauth%3Atoken-type%3Aaccess_token&resource=https%3A%2F%2Frs2.example.com%2F
&txn_handle=th_9876543210_opaque_handle
&authorization_details=%5B%7B%22type%22%3A%22payment%22%2C%22amount%22%3A100%7D%5D
```

  
## 2. Resulting Token Comparison  

A. Standard OAuth JWT Access Token ([RFC 9068](https://datatracker.ietf.org/doc/html/rfc9068))  
Encoded String over the wire:  
`eyJhbGciOiJSUzI1NiIsInR5cCI6ImF0K2p3dCIsImtpZCI6ImsxIn0.eyJpc3MiOiJodHRwczovL2FzLmV4YW1wbGUuY29tIiwic3ViIjoidXNlcl80NSIsImF1ZCI6Imh0dHBzOi8vcnMxLmV4YW1wbGUuY29tLyIsImNsaWVudF9pZCI6ImFnZW50X2FwcF8xMjMiLCJpYXQiOjE3MDAwMDAwMDAsImV4cCI6MTcwMDAwMzYwMCwianRpIjoiYWJjMTIzIiwic2NvcGUiOiJyZWFkIHdyaXRlIn0.signature_bytes...`  

Decoded JOSE Header:  

```json
{  "alg": "RS256",  "typ": "at+jwt",  "kid": "k1"}
```

  
Decoded Payload:  

```json
{  
"iss": "https://as.example.com",  
"sub": "user_45",  
"aud": "https://rs1.example.com/",  
"client_id": "agent_app_123",  
"iat": 1700000000,  
"exp": 1700003600,  
"jti": "abc123",  
"scope": "read write"
}
```

  
B. Transactional Access Token (Txn-AT)  

Encoded String over the wire:  
`eyJhbGciOiJSUzI1NiIsInR5cCI6InR4bmF0K2p3dCIsImtpZCI6ImsxIn0.eyJpc3MiOiJodHRwczovL2FzLmV4YW1wbGUuY29tIiwic3ViIjoidXNlcl80NSIsImF1ZCI6Imh0dHBzOi8vcnMyLmV4YW1wbGUuY29tLyIsImNsaWVudF9pZCI6ImFnZW50X2FwcF8xMjMiLCJpYXQiOjE3MDAwMDAwMDAsImV4cCI6MTcwMDAwMDMwMCwianRpIjoieHl6Nzg5IiwidHhuIjoiOTBhMWI3YzM4ZjkyIiwidGN0eCI6eyJyaXNrX3Njb3JlIjo1LCJhdXRoX2xldmVsIjoiaGlnaCJ9fQ.signature_bytes...`  
Decoded JOSE Header:  

```json
{  "alg": "RS256",  "typ": "txnat+jwt",  "kid": "k1"}
```

  
Decoded Payload:  

```json
{  
"iss": "https://as.example.com",  
"sub": "user_45",  
"aud": "https://rs2.example.com/",  
"client_id": "agent_app_123",  
"iat": 1700000000,  
"exp": 1700000300,  
"jti": "xyz789",  
"txn": "90a1b7c38f92",  
"tctx": {    "risk_score": 5,    "auth_level": "high"  }}
```

  
Summary of Wire Differences  

|Dimension|Standard JWT Access Token (at+jwt)|Transactional Access Token (txnat+jwt)|
|---|---|---|
|JOSE Header typ|"at+jwt"|"txnat+jwt" (Forces non-supporting RSs to reject it safely)|
|Lifetime (exp - iat)|Long-lived (typically 1 hour / 3600s)|Extremely short-lived (\le 300 seconds)|
|Transaction ID (txn)|None|Unique per-audience transaction identifier|
|Context Assertions (tctx)|None|Contains AS-evaluated transaction context|
|Response Payload|Returns token only|Returns token and a txn_handle for subsequent calls|