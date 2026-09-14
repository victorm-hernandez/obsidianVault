Passkeys (WebAuthn) prevent Man-in-the-Middle (MitM) and phishing attacks through two distinct mechanisms, addressing both **domain spoofing** (phishing sites) and **cross-device authentication** (using a phone to log into a computer).

---

### 1. The Core MitM Defense: Cryptographic Origin Binding

Even if a user is tricked into visiting a fake website (e.g., `http://evil-bank.com` instead of `https://realbank.com`), **passkeys cannot be phished.** Here is why:

1. **The Browser Enforces the Origin**: When a webpage calls `navigator.credentials.get()`, the browser independently looks at the domain in the address bar (`evil-bank.com`). The webpage cannot fake this.
2. **Origin Hashing**: The browser puts this real origin (`https://evil-bank.com`) into a data structure called `clientDataJSON`.
3. **Signed Challenge**: The passkey creates a digital signature that covers both the server's challenge and the `clientDataJSON` containing the browser-verified origin.
4. **Validation Failure**: The attacker on `evil-bank.com` intercepts the signature and tries to relay it to the real `realbank.com` server. However, when `realbank.com` inspects the signature, it sees that it was signed for `evil-bank.com` rather than `realbank.com`, and immediately rejects the login.

Because the private key never leaves the device and the signature is mathematically bound to the domain in the browser bar, relaying the login attempt through a proxy or fake site is impossible.

---

### 2. How Cross-Device Roaming Works Without MitM Risks

Passkeys support roaming in two main ways: **Cloud Synchronization** and **Cross-Device / Hybrid Authentication** (e.g., scanning a QR code on a PC with your phone). Both have strict security controls:

#### A. Cross-Device Authentication (QR Code + Bluetooth)

If you are logging into a browser on a desktop/Smart TV using a passkey stored on your mobile phone:

- **E2E Encrypted Tunnel**: Scanning the QR code establishes a secure, end-to-end encrypted direct tunnel between the phone and the desktop browser using ephemeral cryptographic keys.
- **Proximity Check via Bluetooth (BLE)**: To prevent a remote attacker in another country from displaying a fake QR code on a phishing website and tricking you into scanning it, **FIDO Hybrid transport requires Bluetooth Low Energy (BLE) proximity.**
    - The phone and the computer must detect each other's BLE signals.
    - If you are sitting in front of your computer, BLE succeeds.
    - If a hacker halfway around the world presents a QR code on a fake site, your phone will scan it, but BLE proximity verification will fail because your phone is not physically near the hacker's computer. The authentication is aborted.

#### B. Cloud Synchronization (Multi-Device Sync)

When passkeys sync across your own devices (via Apple iCloud Keychain, Google Password Manager, 1Password, etc.):

- **End-to-End Encryption (E2EE)**: The private keys are encrypted on your device using keys derived from your device passcode/biometrics before being synced to the cloud.
- **No Provider Access**: Neither Apple, Google, nor any cloud provider can read or use your passkeys. They can only be decrypted on your trusted devices that share your account chain.

---

### Summary

|Threat Scenario|How Passkeys Prevent It|
|---|---|
|**Fake Web Site / Phishing Proxy**|The signature is locked to the domain in the browser's address bar. Relay fails on the real server.|
|**Remote QR Code Interception**|Bluetooth (BLE) proximity checks guarantee the phone and browser are physically near each other.|
|**Intercepted Cloud Passkeys**|End-to-End Encryption ensures passkeys cannot be intercepted in transit or read in the cloud.|