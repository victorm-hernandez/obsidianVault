To validate that a device/authenticator belongs to a specific vendor (e.g., ensuring a key is a genuine Yubico YubiKey, an Apple device, or a specific enterprise hardware key), you must use **WebAuthn Attestation** combined with **Certificate Chain Verification** and the **FIDO Metadata Service (MDS)**.

Here is the step-by-step process of how to perform vendor validation on your server:

---

### Step 1: Request Direct Attestation During Registration

By default, browsers often default to `attestation: "none"` to protect user privacy. To get vendor proof, your server must explicitly request attestation when creating credential options:

```
javascript

const publicKeyCredentialCreationOptions = {

  // ... standard options

  attestation: "direct" // Request vendor/hardware attestation proof

  // Or "enterprise" for enterprise-managed devices

};
```

---

### Step 2: Extract the `AAGUID` and Certificate Chain

When the client returns the `attestationObject`, your server parses the CBOR structure to extract two key pieces of information:

1. **AAGUID (Authenticator Attestation GUID)**: A 128-bit identifier stored in `authData` that indicates the exact make and model of the authenticator (e.g., YubiKey 5 Series vs. YubiKey 5 NFC).
2. **`x5c` Certificate Chain**: Located inside the `attStmt` (Attestation Statement). This contains one or more X.509 certificates embedded in the physical device by the manufacturer during production.

---

### Step 3: Validate Against the FIDO Metadata Service (MDS)

The **FIDO Alliance** maintains the **FIDO Metadata Service (MDS)**—a centralized database containing verified metadata and Root CA certificates for certified authenticators.

Your server performs the following checks against MDS:

1. **Look up by AAGUID**: Query the FIDO MDS using the `AAGUID` extracted from the authenticator data.
2. **Verify Vendor Info**: The MDS entry will return metadata confirming the vendor name (e.g., _"Yubico"_, _"Apple Inc."_), device model, and security certification levels.
3. **Verify Root CA**: The MDS entry provides the official **Attestation Root Certificate(s)** for that specific device model.

---

### Step 4: Perform Cryptographic Chain Verification (PKI)

Once you have the root certificate from FIDO MDS (or directly from the vendor), perform standard Public Key Infrastructure (PKI) validation:

1. **Chain Verification**: Verify that the leaf certificate (`x5c[0]`) from the device was signed by the vendor's intermediate certificate, which in turn anchors to the vendor's **Root CA Certificate**.
2. **Signature Verification**: Verify that the signature over `authData` and `clientDataHash` was created by the public key inside the leaf attestation certificate.
3. **Revocation Check**: Ensure the attestation certificate has not been revoked (via CRLs or OCSP).

---

### Vendor-Specific Examples & Roots of Trust

Different vendors handle root certificates and formats slightly differently:

- **Yubico (YubiKeys)**: Uses the `packed` attestation format. You can verify the certificate chain against Yubico's published Root CA or through FIDO MDS.
- **Apple (iOS / macOS)**: Uses the `apple-anonymous` attestation format. The certificate chain must anchor back to the **Apple WebAuthn Root CA**.
- **Android Devices**: Uses the `android-key` format. The certificate chain must anchor back to Google's **Hardware Attestation Root CA** (`Google_Hardware_Attestation_Root1`).
- **Windows / TPM**: Uses the `tpm` attestation format. The certificate chain anchors back to TPM chip manufacturers (Infineon, Nuvoton, Intel, STMicroelectronics, etc.).

---

### Summary Checklist for Implementation

1. Set `attestation: "direct"` in `create()` options.
2. Parse `authData` to get the `AAGUID`.
3. Parse `attStmt` to get the `x5c` certificate array.
4. Fetch the expected Root CA for that `AAGUID` from **FIDO MDS**.
5. Verify the cryptographic chain from `x5c[0]` →→ Intermediate CAs →→ Vendor Root CA.