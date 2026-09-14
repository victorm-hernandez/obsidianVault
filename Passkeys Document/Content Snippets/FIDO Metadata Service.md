The **FIDO Metadata Service (MDS)** is a centralized, authoritative repository maintained by the **FIDO Alliance**. It acts as a "single source of truth" for information about certified FIDO2, WebAuthn, U2F, and UAF authenticators.

When an authenticator (such as a YubiKey, Google Titan Key, or a phone's Secure Enclave) registers with a server, it sends an **AAGUID** (Authenticator Attestation GUID) and a certificate chain. The FIDO MDS allows the server to look up that AAGUID to verify **what the device is**, **who made it**, and **whether it can be trusted**.

---

### What Information Does FIDO MDS Provide?

For every registered authenticator model, the FIDO MDS provides a **Metadata Statement** containing:

1. **Vendor & Model Identity**: The official name of the manufacturer (e.g., _"Yubico"_, _"Feitian"_), device model name, description, and UI icons.
2. **Attestation Root Certificates**: The trusted X.509 Root CA certificates issued by the manufacturer. Your server uses these root certificates to cryptographically verify the device's attestation signature.
3. **Security Capabilities**: Detailed hardware specs, such as:
    - Key storage protection (e.g., hardware-backed TEE, Secure Element, TPM).
    - Biometric capabilities (fingerprint reader, facial recognition) and false-acceptance rates.
    - User verification requirements (PIN, biometrics, simple touch).
4. **FIDO Certification Level**: The security assurance level assigned by the FIDO Alliance evaluation process (e.g., FIDO Certified Level 1, Level 2, Level 3+).
5. **Status Reports & Revocations**: Information on known security vulnerabilities, key compromises, or revocations (e.g., status flags like `USER_VERIFICATION_BYPASS` or `ATTESTATION_KEY_COMPROMISE`).

---

### Why Do WebAuthn Relying Parties Use It?

The FIDO MDS is critical for organizations—especially in **enterprise, banking, and government** environments—for two main reasons:

- **Policy Enforcement**: An enterprise can create access policies such as: _"Only allow users to log in if their authenticator is a FIDO Level 2 Certified hardware key manufactured by Yubico or Kensington."_ Without MDS, a server cannot distinguish a genuine hardware key from a software emulator.
- **Security & Vulnerability Management**: If a specific model of security key is discovered to have a cryptographic flaw, the FIDO Alliance updates its status in the MDS. Relying parties sync with MDS to automatically block or flag compromised key models.

---

### How Developers Interact with FIDO MDS (MDS v3)

Rather than making individual API calls for every user registration, servers fetch and cache the MDS database:

1. **Download the MDS BLOB**: The server periodically (e.g., daily or weekly) downloads the main metadata file from the FIDO Alliance endpoint (`https://mds.fidoalliance.org/`).
2. **Verify the JWT Signature**: The payload is a signed JSON Web Token (JWT). The server verifies the JWT signature using the official FIDO Alliance Root Certificate to ensure the metadata hasn't been tampered with.
3. **Cache the Payload**: The server parses the JSON list, which contains metadata entries for hundreds of authenticator models indexed by their `AAGUID`.
4. **Query During Registration**: When a user registers a new credential, the server extracts the `AAGUID` from the WebAuthn response, looks it up in its cached MDS database, and validates the attestation certificate chain against the Root CA listed in that MDS entry.

3:07 PM