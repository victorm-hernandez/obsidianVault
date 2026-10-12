https://www.rfc-editor.org/info/rfc8705/

RFC 8705: OAuth 2.0 Mutual-TLS Client Authentication and Certificate-Bound Access Tokens defines how to use Mutual Transport Layer Security (mTLS) with X.509 certificates in OAuth 2.0.  
It introduces two distinct mechanisms that significantly strengthen security by replacing shared secret credentials and mitigating token theft/replay attacks:  

## Key Concepts & Components  

### 1. Mutual-TLS for OAuth Client Authentication  
OAuth clients use mTLS to authenticate themselves to authorization server endpoints (such as `/token`, `/introspect`, or `/revoke`). RFC 8705 specifies two ways to validate client credentials:  

- PKI Mutual-TLS Method (tls_client_auth):  
    - Relies on standard X.509 Public Key Infrastructure (PKI) certificate chains signed by a trusted Certificate Authority (CA).  
    - The Authorization Server validates the certificate chain and verifies that a specific identity field matches the client's registered identifier:  
	    - Subject Distinguished Name (`tls_client_auth_subject_dn`)  
	    - Subject Alternative Name (SAN) for DNS (`tls_client_auth_san_dns`), URI (`tls_client_auth_san_uri`), IP address (`tls_client_auth_san_ip`), or Email (`tls_client_auth_san_email`).  
	- **Advantage**: Seamless certificate rotation without updating registration data at the Authorization Server, as long as the CA and Subject name remain identical.  
    
- Self-Signed Certificate Mutual-TLS Method (self_signed_tls_client_auth):  
    - Does not require a public or private PKI/CA infrastructure.  
    - The client registers its specific self-signed X.509 certificate (or public key fingerprint) with the Authorization Server via the `jwks` or `jwks_uri` client metadata parameters.  
    - **Advantage**: Lightweight setup for direct trust relationships without operating a Certificate Authority.  

## 2. Certificate-Bound Access Tokens (Holder-of-Key / Proof-of-Possession)  

Standard OAuth 2.0 access tokens are bearer tokens—anyone who steals or intercepts the token can present it to a resource server and gain access.  

RFC 8705 introduces certificate binding:  

- Issuance: When issuing an access token over an mTLS connection, the Authorization Server computes a cryptographic hash (SHA-256) of the client's X.509 DER certificate. This thumbprint (`x5t#S256`) is embedded into the token (e.g., inside the `cnf` / confirmation claim of a JWT) or linked via token introspection metadata.  
    
- Enforcement: When the client calls a protected API resource server:  
    The client establishes an mTLS connection with the Resource Server.  
    The Resource Server extracts the thumbprint of the client certificate presented during the TLS handshake.  
    The Resource Server verifies that this thumbprint matches the binding inside the presented token.  
    
- Outcome: If an attacker steals a certificate-bound token, it is rendered completely useless unless they also possess the client's private key to complete the mTLS handshake.  

## Metadata & Registration Parameters Summary  

|Parameter Name|Endpoint / Context|Purpose|
|---|---|---|
|tls_client_auth|Token Endpoint Auth|Client uses PKI-validated mTLS authentication.|
|self_signed_tls_client_auth|Token Endpoint Auth|Client uses self-signed mTLS certificate authentication.|
|tls_client_certificate_bound_access_tokens|Dynamic Client Registration|Flag (true/false) indicating the client requires certificate-bound tokens.|
|cnf (Confirmation)|JWT / Introspection|Claims object carrying the x5t#S256 thumbprint of the client certificate.|
  
## Why RFC 8705 Matters  

- Eliminates Token Replay Attacks: Stops stolen tokens from being used by unauthorized third parties.  
    
- High-Security Application Standard: Required by modern financial and open-banking specifications worldwide (such as FAPI / Financial-grade API standards).  
    
- Protects Both Steps of OAuth: Provides strong authentication during token retrieval and strong authorization during resource access.