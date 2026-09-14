https://webauthn.guide/

### 1. Introduction

- **Title:** Introduction
- **Purpose:** To highlight the fundamental flaws of traditional password-based authentication and establish the need for a better system.
- **Summary:** This section explains how our online lives are overly reliant on passwords, which cause major headaches for both users and developers. It notes that passwords are a "shared secret" and that 81% of hacking-related breaches leverage stolen or weak passwords. It sets the stage for a new, passwordless alternative.

### 2. About WebAuthn

- **Title:** About WebAuthn
- **Purpose:** To introduce Web Authentication (WebAuthn) and public key cryptography as the modern solution to password vulnerabilities.
- **Summary:** WebAuthn is an API specification by W3C and FIDO that replaces passwords with public key cryptography. Instead of sharing a secret, a private key is stored securely on the user's device while the public key is sent to the server. The section outlines WebAuthn's three major security properties: it is **Strong** (often backed by hardware security modules), **Scoped** (origin-bound to prevent phishing), and **Attested** (servers can verify the public key comes from a trusted authenticator).

### 3. WebAuthn API

- **Title:** WebAuthn API
- **Purpose:** To introduce the technical usage and implementation of the Web Authentication API.
- **Summary:** This section acts as a high-level transition into the technical details, outlining that developers can use the API to manage both the creation of new credentials (registration) and the verification of existing ones (authentication).

### 4. Registering

- **Title:** Registering
- **Purpose:** To detail the client-side and server-side processes for creating and registering a new WebAuthn credential.
- **Summary:** It explains how a server prompts a user to create a new keypair using `navigator.credentials.create()`. It breaks down the required `publicKeyCredentialCreationOptions` (such as the server challenge, relying party info, and user info). It then details how the returned `credential` object—which includes the `clientDataJSON` and `attestationObject`—is passed to the server, parsed, and validated to finalize the registration.

### 5. Authenticating

- **Title:** Authenticating
- **Purpose:** To explain the workflow of how an existing user signs in by proving ownership of their private key.
- **Summary:** During authentication, a user generates an assertion by calling `navigator.credentials.get()`, which creates a cryptographic signature using their device's private key. The section breaks down the `publicKeyCredentialRequestOptions` and the resulting `assertion` object. Finally, it provides pseudo-code to explain how the server verifies this signature using the public key it stored during the registration phase.

### 6. Looking Ahead

- **Title:** Looking Ahead
- **Purpose:** To provide concluding thoughts on the role of WebAuthn within the broader scope of application security.
- **Summary:** It serves as a reminder that while Web Authentication is a massive leap forward (forcing 80% of hacking attacks to adapt or die), security is an overarching mindset that must be incorporated into every step of software design, rather than just relying on a single technology.
