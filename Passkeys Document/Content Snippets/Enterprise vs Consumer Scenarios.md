The WebAuthn Level 3 specification (specifically Section 1.2) explicitly categorizes usage into **Consumer** and **Enterprise (Workforce)** scenarios.

Because consumers prioritize **convenience and privacy**, while enterprises prioritize **strict policy control and un-copyable hardware**, WebAuthn behaves quite differently between these two environments.

---

### 1. Consumer Scenarios (B2C / Public Web)

In consumer scenarios, websites (like e-commerce, social media, or banking) want to eliminate passwords while ensuring users never get locked out of their accounts when they buy a new phone.

#### A. Multi-Device Credentials (Synced Passkeys)

- **The Scenario**: A user registers on a website (e.g., Amazon or eBay) using their iPhone. Later, they visit the website on their iPad or Mac.
- **How it works**: The credential (private key) is generated inside a cloud-synced password manager (such as Apple iCloud Keychain, Google Password Manager, or 1Password).
- **Key Characteristics**:
    - **Zero Friction**: The credential automatically syncs across all of the user's personal devices.
    - **Self-Service Recovery**: If the user loses their phone, they don't get locked out of their account; they just log into their Apple/Google account on a new device, and their passkeys are restored.
    - **Privacy First**: Websites use `attestation: "none"` so they cannot fingerprint or track the user's specific hardware model.

#### B. Cross-Device Authentication (Hybrid / QR Code Login)

- **The Scenario**: A user wants to log into their streaming account on a Smart TV, a friend’s laptop, or a public library computer.
- **How it works**: The computer displays a QR code. The user scans it with their personal smartphone, completes a biometric check (Face ID) on their phone, and is logged in on the computer.
- **Key Characteristics**: The user's credential stays safely on their personal phone; no sensitive data is left behind on the shared/guest computer.

---

### 2. Enterprise / Workforce Scenarios (B2E / B2B)

In enterprise environments (such as corporate SSO, healthcare, finance, or defense), IT administrators need absolute control over identity, strict compliance, and zero trust verification.

#### A. Single-Device Bound Credentials (Hardware Keys & TPMs)

- **The Scenario**: Employees logging into corporate Single Sign-On (SSO) portals (e.g., Microsoft Entra ID, Okta, Ping Identity).
- **How it works**: IT departments require credentials to be bound to a **single physical device** (such as a FIPS-compliant YubiKey or a PC's hardware TPM via Windows Hello).
- **Key Characteristics**:
    - **Non-Exportable**: Credentials **cannot be synced to personal cloud accounts** (like iCloud or Google Password Manager).
    - **Device Loss Protocol**: If an employee loses their YubiKey, IT revokes the credential centrally, and the employee must be issued a new physical key.

#### B. Hardware Attestation & Vendor Whitelisting

- **The Scenario**: A financial institution enforces a policy: _"Employees may only log in using company-issued YubiKey 5 Series keys; personal phones or unapproved USB keys are strictly forbidden."_
- **How it works**:
    - The server requests `attestation: "direct"`.
    - During registration, the server extracts the `AAGUID` and `x5c` certificate chain and cross-references it with the **FIDO Metadata Service (MDS)**.
    - If the certificate isn't signed by Yubico’s Root CA, the registration is rejected.

#### C. Enterprise Attestation (`attestation: "enterprise"`)

- **The Scenario**: Fully managed corporate laptops and devices.
- **How it works**: In standard WebAuthn, attestation certificates are shared across thousands of devices to prevent tracking users. However, in an enterprise setting, privacy between employer and employee is not required. `attestation: "enterprise"` allows authenticators to return uniquely identifiable hardware serial numbers, allowing IT to map specific cryptographic keys directly to individual corporate assets.

#### D. Shared Workstation / Shift Worker Scenarios

- **The Scenario**: Hospital staff or call center employees who don't have assigned PCs and sit at different shared computers every shift.
- **How it works**: Employees carry a physical security key on their badge lanyard. They plug it into any shared terminal, tap their key (or enter a PIN), complete their shift, and unplug the key when leaving.

---

### Summary Comparison

| Dimension                  | Consumer Scenario                  | Enterprise Scenario                       |
| -------------------------- | ---------------------------------- | ----------------------------------------- |
| **Primary Goal**           | Frictionless UX & Account Recovery | Maximum Security, Compliance & Control    |
| **Credential Type**        | Multi-Device (Synced Passkey)      | Single-Device (Bound Hardware Key / TPM)  |
| **Attestation Preference** | `none` (Privacy-preserving)        | `direct` or `enterprise` (Hardware proof) |
| **Device Storage**         | iCloud Keychain, Google, 1Password | YubiKey, Titan Key, Windows Hello TPM     |
| **Recovery Mechanism**     | Cloud Provider Backup              | IT Admin Helpdesk / Backup YubiKey        |