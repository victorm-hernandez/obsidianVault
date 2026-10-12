## JWT is really a family

## JWS (signed)

`header.payload.signature`, three dot-separated parts. Readable by anyone, tamper-proof. This is what people usually mean by "JWT."

```json
{
  "header": { "alg": "RS256", "kid": "key-1" },
  "payload": {
    "iss": "https://auth.example.com",
    "sub": "user-42",
    "aud": "https://orders.example.com",
    "scope": "orders.read",
    "exp": 1790000000
  },
  "signature": "SIG_A"
}
```

## JWE (encrypted)

`header.encrypted_key.iv.ciphertext.tag`, five parts. The claims are hidden from everyone except the recipient. JWE is used when the client or an intermediary shouldn't see the claims, such as encrypted ID tokens in OIDC or PII in a token. The catch is that the resource server needs the decryption key, so it's rarer than plain JWS.

```json
{
  "header": { "alg": "RSA-OAEP-256", "enc": "A256GCM", "kid": "orders-key-1" },
  "encrypted_key": "ENC_KEY",
  "iv": "IV",
  "ciphertext": "CIPHERTEXT",
  "tag": "TAG",
  "plaintext_after_decrypt": {
    "iss": "https://auth.example.com",
    "sub": "user-42",
    "email": "alice@example.com"
  }
}
```

## Nested JWT

A combination of both! a signed JWT wrapped inside a JWE (`cty: "JWT"` in the header). This is the standard way to get both **authenticity and confidentiality**.

Why nested JWTs exist:  A signed JWT proves **who issued it** but hides nothing. An encrypted JWE hides the content but, on its own, doesn't prove who issued it. A nested JWT gets both: **authenticity from the inner signature, confidentiality from the outer encryption.**

Why a JWE alone isn't enough: A JWE is encrypted with the _recipient's public key_, and anyone can obtain a public key. So anyone can create a valid JWE addressed to that recipient, and decrypting it proves nothing about who sent it. The AEAD tag only proves the ciphertext wasn't modified, not who produced it. To trust the claims, the recipient needs a signature from a known issuer. The signature goes inside so it isn't exposed.

---
## SPIFFE SVID

SPIFFE defines two forms. Both carry the same identity, a SPIFFE ID such as `spiffe://example.org/ns/prod/sa/orders`.

### X.509-SVID

A normal X.509 certificate with the SPIFFE ID as a **URI SAN**. Decoded (`openssl x509 -text`):

```
Certificate:
    Data:
        Version: 3 (0x2)
        Serial Number: 5f:3a:...:9c
        Signature Algorithm: ecdsa-with-SHA256
        Issuer: C=US, O=example.org
        Validity
            Not Before: Sep 28 14:00:00 2026 GMT
            Not After : Sep 28 15:00:00 2026 GMT
        Subject: C=US, O=SPIRE
        Subject Public Key Info:
            Public Key Algorithm: id-ecPublicKey
                Public-Key: (256 bit)   ASN1 OID: prime256v1
        X509v3 extensions:
            X509v3 Key Usage: critical
                Digital Signature, Key Encipherment
            X509v3 Extended Key Usage:
                TLS Web Server Authentication, TLS Web Client Authentication
            X509v3 Basic Constraints: critical
                CA:FALSE
            X509v3 Subject Alternative Name:
                URI:spiffe://example.org/ns/prod/sa/orders
```

Rules worth knowing: exactly one URI SAN, `CA:FALSE`, and short lifetimes (often an hour or less). Workloads use it for mTLS and the peer checks the URI SAN against its policy.

### JWT-SVID

```json
{
  "header": { "alg": "ES256", "kid": "spire-key-3", "typ": "JWT" },
  "payload": {
    "sub": "spiffe://example.org/ns/prod/sa/orders",
    "aud": ["spiffe://example.org/ns/prod/sa/inventory"],
    "iat": 1789999700,
    "exp": 1790000000
  },
  "signature": "SIG"
}
```

`sub`, `aud`, and `exp` are required. It's a bearer token, so keep the `aud` tight and the lifetime short. Verifiers fetch the trust bundle's signing keys from the SPIFFE bundle endpoint, not from an OIDC discovery URL (though SPIRE can also publish one).

---

## Kubernetes service account token

The modern form is a **projected, bound token**: a JWT tied to a specific pod, audience, and expiry.


```json
{
  "header": { "alg": "RS256", "kid": "k8s-signing-key-1" },
  "payload": {
    "aud": ["https://kubernetes.default.svc.cluster.local"],
    "iss": "https://kubernetes.default.svc.cluster.local",
    "iat": 1789996400,
    "nbf": 1789996400,
    "exp": 1790000000,
    "sub": "system:serviceaccount:prod:orders",
    "kubernetes.io": {
      "namespace": "prod",
      "serviceaccount": { "name": "orders", "uid": "6f1c...a2" },
      "pod": { "name": "orders-7d9f-x4k2p", "uid": "b81e...07" }
    }
  },
  "signature": "SIG"
}
```

- **Where it lives:** `/var/run/secrets/kubernetes.io/serviceaccount/token`, or a projected volume with a custom `audience` and `expirationSeconds`.
- **Bound:** the token dies when the pod is deleted, and the kubelet refreshes it before expiry.
- **`iss`:** the API server's default issuer, or your configured OIDC issuer URL. When the issuer is publicly reachable, external systems (AWS, GCP, Vault) can verify these tokens, which is the workload identity federation pattern from earlier.
- **Legacy tokens:** the old secret-based ones had `iss: "kubernetes/serviceaccount"`, no `exp`, and no pod binding. Avoid them.

---

## X.509 certificate

A TLS server certificate, decoded, then the PEM form:

```
Certificate:
    Data:
        Version: 3 (0x2)
        Serial Number: 0a:1b:...:ef
        Signature Algorithm: sha256WithRSAEncryption
        Issuer: C=US, O=Example CA, CN=Example Issuing CA 1
        Validity
            Not Before: Sep  1 00:00:00 2026 GMT
            Not After : Dec  1 23:59:59 2026 GMT
        Subject: CN=orders.example.com
        Subject Public Key Info:
            Public Key Algorithm: rsaEncryption
                Public-Key: (2048 bit)
        X509v3 extensions:
            X509v3 Subject Alternative Name:
                DNS:orders.example.com, DNS:orders-internal.example.com
            X509v3 Key Usage: critical
                Digital Signature, Key Encipherment
            X509v3 Extended Key Usage:
                TLS Web Server Authentication
            X509v3 Basic Constraints: critical
                CA:FALSE
            X509v3 Authority Key Identifier: keyid:AB:CD:...
            X509v3 CRL Distribution Points:
                URI:http://crl.example-ca.com/issuing1.crl
            Authority Information Access:
                OCSP - URI:http://ocsp.example-ca.com
                CA Issuers - URI:http://certs.example-ca.com/issuing1.crt
```

```
-----BEGIN CERTIFICATE-----
MIIDdzCCAl+gAwIBAgIUCh...(truncated)...
-----END CERTIFICATE-----
```

Compared with the SVID above, it's the same format with a different identity: DNS names in the SAN instead of a SPIFFE URI, and a longer lifetime. Validation is chain building to a trusted root plus checks on validity, key usage, and revocation (CRL or OCSP).

---

## SAML assertion

```xml
<saml:Assertion xmlns:saml="urn:oasis:names:tc:SAML:2.0:assertion"
                ID="_a75adf55-01d7-40cc-929f-dbd8372ebdfc"
                Version="2.0"
                IssueInstant="2026-09-28T14:00:00Z">

  <saml:Issuer>https://idp.example.com/saml</saml:Issuer>

  <ds:Signature xmlns:ds="http://www.w3.org/2000/09/xmldsig#">
    <ds:SignedInfo>
      <ds:CanonicalizationMethod Algorithm="http://www.w3.org/2001/10/xml-exc-c14n#"/>
      <ds:SignatureMethod Algorithm="http://www.w3.org/2001/04/xmldsig-more#rsa-sha256"/>
      <ds:Reference URI="#_a75adf55-01d7-40cc-929f-dbd8372ebdfc">
        <ds:Transforms>
          <ds:Transform Algorithm="http://www.w3.org/2000/09/xmldsig#enveloped-signature"/>
          <ds:Transform Algorithm="http://www.w3.org/2001/10/xml-exc-c14n#"/>
        </ds:Transforms>
        <ds:DigestMethod Algorithm="http://www.w3.org/2001/04/xmlenc#sha256"/>
        <ds:DigestValue>Zm9v...</ds:DigestValue>
      </ds:Reference>
    </ds:SignedInfo>
    <ds:SignatureValue>MIIB...</ds:SignatureValue>
    <ds:KeyInfo><ds:X509Data><ds:X509Certificate>MIID...</ds:X509Certificate></ds:X509Data></ds:KeyInfo>
  </ds:Signature>

  <saml:Subject>
    <saml:NameID Format="urn:oasis:names:tc:SAML:2.0:nameid-format:persistent">user-42</saml:NameID>
    <saml:SubjectConfirmation Method="urn:oasis:names:tc:SAML:2.0:cm:bearer">
      <saml:SubjectConfirmationData
          InResponseTo="_req-8f2a91"
          Recipient="https://sp.example.com/saml/acs"
          NotOnOrAfter="2026-09-28T14:05:00Z"/>
    </saml:SubjectConfirmation>
  </saml:Subject>

  <saml:Conditions NotBefore="2026-09-28T13:59:30Z" NotOnOrAfter="2026-09-28T14:05:00Z">
    <saml:AudienceRestriction>
      <saml:Audience>https://sp.example.com/metadata</saml:Audience>
    </saml:AudienceRestriction>
  </saml:Conditions>

  <saml:AuthnStatement AuthnInstant="2026-09-28T13:59:58Z" SessionIndex="_sess-3d91">
    <saml:AuthnContext>
      <saml:AuthnContextClassRef>
        urn:oasis:names:tc:SAML:2.0:ac:classes:PasswordProtectedTransport
      </saml:AuthnContextClassRef>
    </saml:AuthnContext>
  </saml:AuthnStatement>

  <saml:AttributeStatement>
    <saml:Attribute Name="email">
      <saml:AttributeValue>alice@example.com</saml:AttributeValue>
    </saml:Attribute>
    <saml:Attribute Name="groups">
      <saml:AttributeValue>orders-admins</saml:AttributeValue>
      <saml:AttributeValue>finance</saml:AttributeValue>
    </saml:Attribute>
  </saml:AttributeStatement>
</saml:Assertion>
```

Validation covers the signature (over the assertion or the enclosing Response), `Issuer`, `Audience`, `Recipient`, the `NotBefore`/`NotOnOrAfter` window, `InResponseTo`, and replay (the assertion `ID` should be single-use). Mismatched signature scope, where the signature covers one element and the app reads another, is the classic SAML attack (XML signature wrapping), so read only what the signature covers.