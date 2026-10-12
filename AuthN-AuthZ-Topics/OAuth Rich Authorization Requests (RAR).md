https://www.rfc-editor.org/info/rfc9396/

RFC 9396: OAuth 2.0 Rich Authorization Requests (RAR) defines a standardized mechanism for OAuth 2.0 clients to express complex, fine-grained authorization requirements using structured JSON data.  

## Core Problem & Purpose  

- The Problem with scope: Standard OAuth 2.0 ([RFC 6749](https://datatracker.ietf.org/doc/html/rfc6749)) relies on the space-separated `scope` parameter (e.g., `scope=read write`). While sufficient for coarse-grained permissions (e.g., "read user profile"), `scope` is flat and unstructured. It cannot cleanly represent dynamic, parameter-driven consent requirements like _"transfer €123.50 to Merchant A"_ or _"grant read access to_ _`/folder/A`_ _and write access to_ _`/folder/A/file1.txt`__"_.  
    
- The RAR Solution: Introduces a new request parameter—authorization_details—that carries an array of structured JSON objects. This allows clients to convey granular context, constraints, and parameters directly to the Authorization Server (AS) and Resource Server (RS).

## Key Concepts & Structure  

#### 1. The `authorization_details` Parameter  

The request payload contains a JSON array of authorization objects. Each object must include a `type` field that dictates the expected schema and handling for that specific resource or API.  

```json
{    
"type": "payment_initiation",    
"actions": ["initiate", "status"],    
"locations": ["https://example.com/payments"],    
"instructedAmount": 
{      
	"currency": "EUR",      
	"amount": "123.50"    
},    
"creditorName": "Merchant A",    
"creditorAccount": {      
	"iban": "DE02100100109307118603"    
	}  
}
]
```
  
#### 2. Authorization Details Types  

- type (Required): A string identifying the API or permission type (e.g., `account_information`, `payment_initiation`).  
-
- Namespacing: For open standards or cross-domain APIs, collision-resistant URIs are recommended for the type field (e.g., `"type": "[https://example.org/files](https://example.org/files)"`).  

#### 3. Standardized / Common Data Fields  

While APIs can define custom fields inside an authorization object, RFC 9396 establishes several reusable, standard fields:  

| Field      | Type             | Description                                             |
| ---------- | ---------------- | ------------------------------------------------------- |
| locations  | Array of strings | URIs identifying the target Resource Servers (RSs).     |
| actions    | Array of strings | Operations requested (e.g., ["read", "write"]).         |
| datatypes  | Array of strings | Kinds of data requested (e.g., ["contacts", "photos"]). |
| identifier | String           | Unique ID specifying a targeted resource instance.      |
| privileges | Array of strings | Levels/types of privilege requested.                    |

_Note: When_ _`actions`_ _and_ _`datatypes`_ _are defined together in a single object, permissions act as a Cartesian product (i.e., all listed actions apply to all listed datatypes)._

## OAuth Protocol Extensions & Workflows  

1. Authorization Request:  
    `authorization_details` can be used across various flows: Authorization Code Grant ([RFC 6749](https://datatracker.ietf.org/doc/html/rfc6749)), Device Authorization Grant ([RFC 8628](https://datatracker.ietf.org/doc/html/rfc8628)), or CIBA (Client Initiated Backchannel Authentication).  
    URL Length & Security: Because JSON payloads can exceed URL length limits in browser GET requests, the specification strongly encourages combining RAR with Pushed Authorization Requests (PAR / RFC 9126).  
    
2. Authorization & Consent:  
    The Authorization Server parses `authorization_details` to present human-readable details on the user consent screen (e.g., showing exact transfer amounts and recipient details).  
    
3. Token Response & Issuance:  
    When issuing an access token, the AS can include the granted `authorization_details` in the token response if the granted permissions differ from what was requested.  
    Access tokens (such as JWTs) can embed `authorization_details` so that Resource Servers know precisely what operations were approved by the user.  
    
4. Token Refresh & Introspection:  
    Clients can request a subset of previously granted `authorization_details` during token refresh.  
    Token Introspection ([RFC 7662](https://datatracker.ietf.org/doc/html/rfc7662)) is updated so RSs can query the active `authorization_details` associated with a token.  
## Primary Use Cases  

- Open Banking / PSD2: Authorizing single-payment transactions with strict parameters (amount, payee, IBAN).  
    
- Healthcare & Electronic Health Records: Restricting access to specific medical record types or time ranges.  
    
- Document Management Systems: Specifying exact folder paths, file-level privileges, or action limits.

## How authorization details are passed to the Resource Server

1. In Structured/Self-Contained Access Tokens (e.g., JWTs)  

When the Authorization Server (AS) issues a structured token (like a JSON Web Token), standard RFC 9396 behavior includes the top-level claim authorization_details directly inside the access token payload.  

```json
{  
"iss": "https://as.example.com",  
"sub": "user_123",  
"aud": "https://api.example.com",  
"exp": 1700000000,  
"authorization_details": [    
	{      
	"type": "payment_initiation",      
	"actions": ["initiate"],      
	"instructedAmount": 
		{        
		"currency": "EUR",        
		"amount": "123.50"      
		},      
	"creditorAccount": 
		{        
		"iban": "DE02100100109307118603"      
		}    
	}  
]}
```

The RS decrypts or verifies the digital signature of the JWT, extracts `authorization_details`, and enforces the exact parameters (e.g., rejecting any request attempting to transfer an amount greater than €123.50 or to a different IBAN).

2. Via Token Introspection (RFC 7662) for Opaque Tokens  
If the AS issues opaque (random string) access tokens, the RS cannot read claims directly from the token string. RFC 9396 updates OAuth 2.0 Token Introspection (RFC 7662) by adding the `authorization_details` parameter to the introspection response.

## Key Enforcement Capabilities Defined by the Standard  

- Consent Modification / Enrichment: The AS can modify or enrich the requested permissions during the consent step (e.g., the user selects specific account numbers on screen). The `authorization_details` emitted in the access token or introspection response reflects what the user _actually approved_, not just what the client asked for.  
    
- Granular Downscoping: When refreshing a token (via `grant_type=refresh_token`) or making a token request, the client can request a subset of previously granted `authorization_details`. The resulting access token will reflect only that narrowed scope.  
    
- Flexible Mapping: RFC 9396 allows the AS to transform raw `authorization_details` into internal permission constructs or specialized RS claims if an API implementation requires non-standard token representations.