
 [WebAuthn](https://en.wikipedia.org/wiki/WebAuthn)

### 1. Background

- **Purpose**: To provide historical context and explain the origins of the WebAuthn standard.
- **Summary**: This section explains that WebAuthn (Web Authentication) is a core component of the FIDO2 Project, developed jointly by the World Wide Web Consortium (W3C) and the FIDO Alliance. It details how the standard evolved from earlier FIDO protocols (like U2F) to provide a standardized web API for passwordless authentication using public-key cryptography.

### 2. Reasons for its design and standardization

- **Purpose**: To explain the fundamental security problems WebAuthn aims to solve and why a standardized approach was necessary.
- **Summary**: Passwords are inherently insecure and vulnerable to phishing, data breaches, and credential stuffing. This section outlines how WebAuthn addresses these vulnerabilities by replacing shared secrets (passwords) with asymmetric public-key cryptography. Standardization across browsers and platforms ensures a unified, secure, and phishing-resistant authentication experience across the entire web ecosystem.

### 3. Passkey branding

- **Purpose**: To clarify the terminology shift and consumer-facing branding of WebAuthn credentials.
- **Summary**: This section explains the term "passkey," which was introduced by an industry coalition (including Apple, Google, and Microsoft) to serve as a more user-friendly, consumer-facing term for discoverable FIDO2/WebAuthn credentials. Passkeys can often be synced across a user's devices via cloud ecosystems (like iCloud Keychain or Google Password Manager), making passwordless login much more convenient.

### 4. Overview

- **Purpose**: To provide a high-level technical explanation of how the WebAuthn protocol operates.
- **Summary**: This section serves as an umbrella for the core mechanics of WebAuthn. It explains the relationship between the Relying Party (the website), the Client (the web browser), and the Authenticator (the hardware key or biometric sensor).
    - **Level 3**: Discusses the latest iteration (Level 3) of the WebAuthn specification, which introduces features like device-bound passkeys and improved UI/UX flows.
    - **Registration**: Details the flow when a user creates a new account or sets up an authenticator. The authenticator generates a new public-private key pair, stores the private key securely, and sends the public key to the server.
    - **Authentication**: Explains the login flow. The server sends a challenge to the client, the authenticator signs the challenge with the private key (after user verification, like a fingerprint), and the server verifies the signature using the stored public key.

### 5. Support

- **Purpose**: To outline the adoption and compatibility of WebAuthn across different browsers, operating systems, and hardware.
- **Summary**: WebAuthn enjoys broad industry support. This section details its implementation across major web browsers (Chrome, Firefox, Safari, Edge) and operating systems (Windows, macOS, Android, iOS). It also notes support from various hardware security key manufacturers, highlighting that WebAuthn is effectively universally supported on modern consumer devices.

### 6. API

- **Purpose**: To provide a technical overview of the JavaScript interfaces used by developers to implement WebAuthn.
- **Summary**: This section briefly touches upon the JavaScript API exposed to web developers. The two primary methods are `navigator.credentials.create()` (used during the registration phase to create a new credential) and `navigator.credentials.get()` (used during the authentication phase to assert a credential).

### 7. Reception

- **Purpose**: To capture the public and industry response to the introduction and rollout of WebAuthn.
- **Summary**: The standard has been widely praised by security researchers and industry professionals for taking a definitive step toward eliminating passwords and neutralizing phishing attacks. However, it also touches on some criticisms or adoption hurdles, such as initial user confusion, the complexity of account recovery if an authenticator is lost, and concerns about vendor lock-in with cloud-synced passkeys.

### 8. In the media

- **Purpose**: To highlight notable mainstream coverage and significant milestones of WebAuthn in the press.
- **Summary**: This section covers major announcements related to WebAuthn, such as the initial W3C standard recommendation, the unified push by Apple, Google, and Microsoft to support passkeys, and major consumer platforms (like PayPal or GitHub) announcing their adoption of the technology.