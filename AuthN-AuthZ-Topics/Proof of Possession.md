A bearer token works for whoever holds it, so a stolen token (logs, XSS, a compromised proxy, a leaked header) is a usable credential. **Sender-constraining** binds the token to a key pair. The token carries a confirmation claim (`cnf`) identifying the key, and every use must prove possession of the matching private key. A stolen token alone is then useless.

### 1. mTLS-bound tokens (RFC 8705)

RFC 8705 covers two related things: mTLS for client authentication, and certificate-bound access tokens.

**Client authentication modes at the token endpoint**

- `tls_client_auth`: a CA-issued certificate. The AS validates the chain and matches the subject DN or SAN registered for the client.
- `self_signed_tls_client_auth`: the client registers its public key (JWKS) and presents a self-signed certificate. The AS checks that the cert's key matches the registered one. There's no PKI, and the trust anchor is the registration.

**Token binding flow**

1. The client opens a TLS connection to the token endpoint with a client certificate (the handshake proves possession of the private key).
2. The AS authenticates the client and computes the SHA-256 thumbprint of the DER-encoded certificate.
3. The AS issues an access token containing `cnf: { "x5t#S256": "<thumbprint>" }`. For opaque tokens, the same `cnf` is returned from introspection.
4. The client calls the resource server over mTLS with the **same** certificate.
5. The RS compares the thumbprint of the certificate on the TLS connection to the token's `cnf`. If they don't match, it rejects.

Notes:

- The token type stays `Bearer` on the wire, with no special header.
- The AS advertises `tls_client_certificate_bound_access_tokens` in metadata, and `mtls_endpoint_aliases` for separate mTLS endpoints.
- The binding is to the certificate's key, so a self-signed cert works for binding even without PKI validation on the RS.

**Strengths:** the private key can live in an HSM or TPM, there's no per-request signing or replay-cache state, and it's mature (it's the FAPI-standard approach for high-security APIs).

**Weaknesses:**

- TLS termination is the biggest operational pain. If a load balancer or CDN terminates TLS, the cert has to be forwarded to the app, usually in a header, and that path must be tightly trusted.
- Certificate provisioning and rotation are hard for browsers and mobile apps.
- It only works where the client controls the TLS layer, so it's a poor fit for SPAs.

### 2. DPoP: Demonstrating Proof-of-Possession (RFC 9449)

DPoP moves the proof to the application layer. There are no certificates and no TLS involvement. The client generates a key pair (in a browser, a non-extractable WebCrypto key; on mobile, the Keystore or Secure Enclave) and signs a small JWT with it for each request.

**The DPoP proof JWT**

- Header: `typ: "dpop+jwt"`, `alg` (asymmetric only), and `jwk` (the **public key embedded in the header**).
- Claims:
    - `jti`: unique ID, for replay detection
    - `htm`: HTTP method
    - `htu`: target URL (without query or fragment)
    - `iat`: issue time
    - `ath`: hash of the access token (only when presenting to an RS)
    - `nonce`: server-provided, if required

**Flow**

1. **Token request:** the client sends `DPoP: <proof>` to the token endpoint, with no `ath` since there's no token yet.
2. The AS validates the proof, computes the JWK thumbprint (RFC 7638), and issues a token with `cnf: { "jkt": "<thumbprint>" }` and `token_type: "DPoP"`.
3. **Resource request:**

```
   Authorization: DPoP <access_token>
   DPoP: <new proof with htm, htu, iat, jti, ath>
```

4. The RS checks all of these:
    - Exactly one DPoP header, and a valid signature using the embedded `jwk`.
    - `typ` and `alg` are acceptable.
    - `htm` and `htu` match the actual request.
    - `iat` is within a small window.
    - `jti` hasn't been seen (replay cache).
    - `ath` equals the hash of the presented token.
    - The thumbprint of the proof's `jwk` equals the token's `cnf.jkt`.

**Refinements**

- **Server nonces:** the AS or RS can respond with a `DPoP-Nonce` header and a `use_dpop_nonce` error, forcing clients to include a fresh nonce. This limits pre-generated proofs with future-dated `iat` values.
- **Auth code binding:** the client can send `dpop_jkt` in the authorization request (ideally via PAR), binding the authorization code to the key and blocking code injection or theft.
- **Refresh tokens:** for public clients, refresh tokens are bound to the key too. Confidential clients are already bound by client authentication.
- **Errors and metadata:** `invalid_dpop_proof`, `dpop_signing_alg_values_supported`, and `WWW-Authenticate: DPoP algs=...`.

**Strengths:** it works for SPAs and mobile, it's independent of TLS termination, and it needs no PKI.

**Weaknesses:**

- Signing happens on every request, and the RS needs `jti` replay tracking (state).
- The proof doesn't cover the request body or other headers.
- If malicious code (XSS) runs in the app, a non-extractable key can't be stolen, but it can still be _used_ to sign proofs while the attacker has execution. That's why nonces and short-lived tokens matter.
- The AS can't verify the key is hardware-backed, because the key is self-asserted (there's no attestation).