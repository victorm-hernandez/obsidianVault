https://www.rfc-editor.org/info/rfc7592/

RFC 7592 defines the OAuth 2.0 Dynamic Client Registration Management Protocol.  

While RFC 7591 defines how a client initially registers with an Authorization Server (AS), RFC 7592 defines the full CRUD lifecycle (Read, Update, Delete) for managing that registration after it has been created.  

## Core Problem It Solves  

Client applications rarely remain static over their lifetime.  

- Redirect URIs change (e.g., domain updates or new callback paths).      
- Keys and client secrets expire or need rotation.  
- Metadata (like application names, logo URIs, or contact info) evolves.  
- Ephemeral or decommissioned client instances need to be safely deprovisioned.  

Without RFC 7592, dynamically registered clients cannot manage their own configurations programmatically, forcing manual developer intervention or custom, non-standard APIs.  

## Key Concepts & Artifacts  

1. Client Configuration Endpoint (registration_client_uri):  
    A unique management URL returned by the Authorization Server when the client initially registers via RFC 7591 (e.g., `[https://server.example.com/register/s6BhdRkqt3](https://server.example.com/register/s6BhdRkqt3)`).  
    
2. Registration Access Token (registration_access_token):  
    A dedicated OAuth 2.0 Bearer Token issued during registration. It authenticates the client specifically for management operations at its Client Configuration Endpoint.  
    

## Management Operations  
RFC 7592 uses standard HTTP methods on the Client Configuration Endpoint:  

|Action|HTTP Method|Description|
|---|---|---|
|Read|GET|Retrieves the current metadata and status of the registered client.|
|Update|PUT|Replaces the existing client metadata with a new payload (must include all fields).|
|Delete|DELETE|Deprovisions and revokes the client's credentials immediately.|

1. Read Request (`GET`)  
The client sends a `GET` request using its `registration_access_token`:  

```http
GET /register/s6BhdRkqt3 HTTP/1.1
Host: server.example.com
Authorization: Bearer reg-23410913-abewfq.123483
```

- Response (200 OK): Returns the current registration JSON. The server _may_ choose to issue a new client secret or rotation token in this response.      

2. Update Request (`PUT`)  
The client sends a `PUT` request with the complete set of new/existing metadata fields:  

```http
PUT /register/s6BhdRkqt3 HTTP/1.1
Host: server.example.com
Content-Type: application/json
Authorization: Bearer reg-23410913-abewfq.123483

{ 
"client_id": "s6BhdRkqt3",  
"client_secret": "cf13605b380c70e2cd21b1ad",  
"redirect_uris": ["https://client.example.org/callback2"],  
"client_name": "Updated App Name",  
"grant_types": ["authorization_code", "refresh_token"]
}
```

  
- Response (200 OK): Confirms updated fields and optionally returns a rotated `registration_access_token` or `client_secret`.  
    

3. Delete Request (`DELETE`)  
The client revokes its own registration:  

```http
DELETE /register/s6BhdRkqt3 HTTP/1.1
Host: server.example.com
Authorization: Bearer reg-23410913-abewfq.123483
```


- Response (204 No Content or 200 OK): Confirms client deprovisioning. The `client_id`, secrets, and registration access tokens are immediately invalidated.  
    

## Key Features & Security Considerations  

- Credential & Token Rotation: Authorization Servers can automatically cycle secrets and registration access tokens during `GET` or `PUT` calls. When a new token is issued, the previous one is discarded immediately.  
    
- Non-Transferable Identification: The `client_id` cannot be modified via `PUT` requests.  
    
- Strict Revocation: If a read or update attempt fails because the client record was deleted or invalid, the `registration_access_token` is revoked immediately.  
    

## Primary Use Cases  

- Autonomous AI Agents & Ephemeral Instances: An AI agent that dynamically registered via RFC 7591 can cleanly delete (`DELETE`) its client credentials when its task finishes, leaving no lingering security surface.  
    
- Key & Secret Rotation: Systems using `private_key_jwt` or client secrets can periodically update their cryptographic keys via `PUT` requests without downtime.  
    
- Changing OAuth Metadata: Dynamically adjusting allowed redirect URIs or scope entitlements across microservices.