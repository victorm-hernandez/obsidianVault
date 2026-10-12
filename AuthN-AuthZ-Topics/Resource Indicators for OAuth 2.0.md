[https://datatracker.ietf.org/doc/html/rfc8707](https://datatracker.ietf.org/doc/html/rfc8707)

RFC 8707: Resource Indicators for OAuth 2.0 introduces a standardized parameter—`resource`—that allows an OAuth 2.0 client to explicitly indicate the target resource server (API/application) where it intends to use an access token.  

## The Problem RFC 8707 Solves  

In core OAuth 2.0 ([RFC 6749](https://datatracker.ietf.org/doc/html/rfc6749)), scopes define _what permissions_ are requested (e.g., `read:profile`, `write:orders`), but there is no standardized way to specify _where_ those permissions will be used.  

Historically, this created two main problems:  

1. Token Misuse / Replay Attacks: Bearer tokens used across multiple resource servers could be stolen by a malicious resource server and replayed against another.     
2. Overloaded Scopes: Implementations often overloaded the `scope` parameter to encode resource target locations (e.g., `[https://api.example.com/read](https://api.example.com/read)`), mixing permission types with audience targets.  

## Key Concepts & Mechanisms  

### 1. The `resource` Parameter  
Clients append one or more `resource` parameters to their authorization or token requests.  
- Format: MUST be an absolute URI without a fragment component (e.g., `[https://api.example.com/](https://api.example.com/)` or `[https://api.example.com/v1/users](https://api.example.com/v1/users)`).  
    
- Multi-Resource Requests: A request can contain multiple `resource` parameters to indicate access across multiple services.  

### 2. Protocol Integration  
The `resource` parameter applies across standard OAuth 2.0 flows:  

- Authorization Endpoint (/authorize): In the authorization code flow, the resource parameter applies to the issued authorization grant. In the implicit flow, it applies directly to the access token.  
    
- Token Endpoint (/token): When redeeming an authorization code, refresh token, or client credentials grant, the client can specify or narrow down the target `resource`.  
    
### 3. Audience Restriction (`aud`)  
When the Authorization Server (AS) processes a request containing a `resource` parameter:  

- It issues access tokens restricted strictly to that resource.      
- For JSON Web Tokens (JWTs), the AS sets the `aud` (audience) claim to the requested URI(s).  
- Resource Servers validate the token's audience and reject any token not minted for them.  

4. Token Downscoping with Refresh Tokens  
When using a `refresh_token`, a client can request an access token for a subset of the original resources granted during user authorization. This enables clients to fetch short-lived tokens restricted to individual microservices on demand rather than reusing a broad token across all APIs.  
5. New Error Code: `invalid_target`  
If a requested `resource` is missing, malformed, unknown, or incompatible with the requested scopes, the Authorization Server returns:  

```json
{  "error": "invalid_target",  "error_description": "The requested resource indicator is unknown or invalid."}
```

  
## Example Flow  

### Step 1: Authorization Request  
The client requests access to both a calendar and contacts API.  

```http
GET /authorize?  response_type=code  &client_id=client-123  &redirect_uri=https%3A%2F%2Fclient.example.com%2Fcb  &scope=read  &resource=https%3A%2F%2Fcalendar.example.com%2F  &resource=https%3A%2F%2Fcontacts.example.com%2F HTTP/1.1Host: auth.example.com
```

  
### Step 2: Access Token Request (Narrowed)  
When redeeming the authorization code, the client asks for an access token specific to the calendar API only.  

```http
POST /token HTTP/1.1Host: auth.example.comContent-Type: application/x-www-form-urlencodedgrant_type=authorization_code&code=SplxlOBeZQQYbYS6WxSbIA&redirect_uri=https%3A%2F%2Fclient.example.com%2Fcb&client_id=client-123&resource=https%3A%2F%2Fcalendar.example.com%2F
```

  
The resulting access token will have an audience (`aud`) matching `[https://calendar.example.com/](https://calendar.example.com/)` and will be rejected if sent to `[https://contacts.example.com/](https://contacts.example.com/)`.  
Summary Checklist  

|Topic|RFC 8707 Requirement|
|---|---|
|Parameter Name|resource|
|Value Requirement|Absolute URI (no fragment #)|
|Cardinality|Multiple parameters allowed per request|
|Token Representation|Reflected in token audience (aud claim or introspection)|
|Error Handling|invalid_target for invalid, missing, or mismatched targets|
Example JWT 

```json 
{
  "sub": "1234567890",
  "iss": "https://auth.example.com",
  "aud": [
    "https://cal.example.com/",
    "https://contacts.example.com/"
  ],
  "exp": 1812345678,
  "client_id": "example-client"
}

```