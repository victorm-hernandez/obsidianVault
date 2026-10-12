
In the context of JSON Web Tokens (JWT) and the broader JSON Object Signing and Encryption (JOSE) framework, the **JOSE header** is the first component of a token. It provides metadata about how the token was securely processed—specifically, which cryptographic algorithms were used to sign or encrypt it, and how to verify it.

### Key Components of a JOSE Header

A JWT is divided into three parts separated by periods (`.`): **Header**, **Payload**, and **Signature**. The header is a JSON object that is Base64Url-encoded.

Common parameters found inside a JOSE header include:

- **`alg` (Algorithm):** The cryptographic algorithm used to secure the JWT (e.g., `HS256` for HMAC using SHA-256, or `RS256` for RSA signature with SHA-256). This is a **required** field for security.
    
- **`typ` (Type):** Declares the media type of the token. For standard JWTs, this is typically set to `JWT`.
    
- **`cty` (ContentType):** Used when payloads are nested (for instance, if a JWT is encrypted inside another JWT).
    
- **`kid` (Key ID):** A hint indicating which specific key (out of a set of keys) a client should use to verify the signature. This is especially useful in systems that rotate keys frequently.
    
- **`jku` (JWK Set URL):** A URI pointing to a secure set of JSON Web Keys (JWKs) that can be used to verify the signature.
    

### Example of a JOSE Header

Here is an example of an unencoded JSON header:

JSON

```
{
  "alg": "HS256",
  "typ": "JWT",
  "kid": "key-id-12345"
}
```

When Base64Url-encoded, this JSON object becomes the first segment of your JWT (e.g., `eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCIsImtpZCI6ImtleS1pZC0xMjM0NSJ9`).

### Security Considerations

- **Algorithm Verification (`alg` header attack):** Applications should always explicitly validate the `alg` parameter against an allowed list of algorithms. A well-known vulnerability involves attackers modifying the `alg` header to `none` (skipping verification) or swapping an asymmetric algorithm (like RS256) for a symmetric one (like HS256) using the public key as a secret.
    
- **Header Tampering:** Because the header is only Base64Url-encoded (not encrypted, in standard signed JWTs), anyone can decode and read it. However, because the header is included in the cryptographic signature, **any tampering with the header will invalidate the token's signature**.