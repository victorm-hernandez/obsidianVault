

```text
<section #>. <section name>
	- Purpose: <Description of this section purpose>
	- Summary: <A summary of the content that will be included here> 
	  <sub section #>. <sub section name>
		  ... <nested section purpose and summary>
```

1.Introduction
	- **Purpose**: To provide a high-level definition of what WebAuthn is.
	- **Summary**: WebAuthn (Web Authentication) is a W3C API specification that provides a secure way for users to log in to online services without passwords, utilizing biometrics and hardware-based authenticators.
2. Terminology 
	- **Purpose**: Defines specific terms and concepts used throughout the document.
	- **Summary**: Provides precise definitions for foundational terms such as "authenticator", "relying party", "public key credential", "attestation", "user verification", and "ceremony".
3. Reasons for its design and standardization
	- **Purpose**: To explain the fundamental security problems WebAuthn aims to solve and why a standardized approach was necessary.
	- **Summary**: Passwords are inherently insecure and vulnerable to phishing, data breaches, and credential stuffing. This section outlines how WebAuthn addresses these vulnerabilities by replacing shared secrets (passwords) with asymmetric public-key cryptography. Standardization across browsers and platforms ensures a unified, secure, and phishing-resistant authentication experience across the entire web ecosystem.
4. What is WebAuthN?
	1. **Purpose:** To introduce Web Authentication (WebAuthn) and public key cryptography as the modern solution to password vulnerabilities.
		- **Summary:** WebAuthn is an API specification by W3C and FIDO that replaces passwords with public key cryptography. Instead of sharing a secret, a private key is stored securely on the user's device while the public key is sent to the server. The section outlines WebAuthn's three major security properties: it is **Strong** (often backed by hardware security modules), **Scoped** (origin-bound to prevent phishing), and **Attested** (servers can verify the public key comes from a trusted authenticator).
5. Passkey branding
	- **Purpose**: To clarify the terminology shift and consumer-facing branding of WebAuthn credentials.
	- **Summary**: This section explains the term "passkey," which was introduced by an industry coalition (including Apple, Google, and Microsoft) to serve as a more user-friendly, consumer-facing term for **discoverable** FIDO2/WebAuthn credentials. Passkeys can often be synced across a user's devices via cloud ecosystems (like iCloud Keychain or Google Password Manager), making passwordless login much more convenient.
	  
	  Clarify that not all WebAuthN Credentials are passkeys. In the WebAuthn specification, credentials are divided into two main categories based on **where the data is stored**. These two types are **Discoverable Credentials** (often colloquially referred to as "Passkeys") and **Non-Discoverable Credentials**.
6. Overview
	- **Purpose**: To provide a high-level technical explanation of how the WebAuthn protocol operates.
	- **Summary**: This section serves as an umbrella for the core mechanics of WebAuthn. It explains the relationship between the Relying Party (the website), the Client (the web browser), and the Authenticator (the hardware key or biometric sensor).
		1. WebAuthN Roles
			- **Purpose**: To define the primary entities and terms involved in the WebAuthn ecosystem.
			- **Summary**: The section outlines four key roles: the **Relying Party**, the **WebAuthn Client Device**, the **User**, and the **Authenticator** (which can be a built-in platform authenticator like a fingerprint scanner, or an external roaming authenticator like a YubiKey).
	    2. **Registration**: Details the flow when a user creates a new account or sets up an authenticator. The authenticator generates a new public-private key pair, stores the private key securely, and sends the public key to the server. [DIAGRAM]
	    3. **Authentication**: Explains the login flow. The server sends a challenge to the client, the authenticator signs the challenge with the private key (after user verification, like a fingerprint), and the server verifies the signature using the stored public key. [DIAGRAM]
		    1. Discoverable vs Non Discoverable key considerations
			    - Purpose: Describe the types of credentials available and ceremony differences for each. 
7. Consumer and Enterprise scenarios
	- **Purpose**: It details various use cases like multi-device credentials for consumers and single-device credentials for workforce environments
8. Security Considerations
	- **Purpose**: Analyzes the security model, threats, and mitigations related to the specification.
	- **Summary**: Discusses how the WebAuthn protocol protects against common attacks like phishing and credential stuffing. 
		1. Cryptographic Origin Binding
		2. Cross-Device Roaming
9. Support
	- **Purpose**: To outline the adoption and compatibility of WebAuthn across different browsers, operating systems, and hardware.
	- **Summary**: WebAuthn enjoys broad industry support. This section details its implementation across major web browsers (Chrome, Firefox, Safari, Edge) and operating systems (Windows, macOS, Android, iOS). It also notes support from various hardware security key manufacturers, highlighting that WebAuthn is effectively universally supported on modern consumer devices.
10. Deep Dive
	1. WebAuthn API
		- **Purpose:** To introduce the technical usage and implementation of the Web Authentication API.
		- **Summary**: Explains the core browser methods (`navigator.credentials.create()` and `navigator.credentials.get()`). It details how scripts request permission to perform authentication operations on the user's behalf without ever exposing the private credential material to the script.
	2. Credential Types
		1. Discoverable Credentials (formerly "Resident Keys")
		2. Non-Discoverable Credentials (or "Server-Side" Credentials)
	3. Authenticators Attestation
		1. Why do we need attestation?
		2. Attestation Formats
		3. FIDO Metadata Service
		4. Use case: Validate that an authenticator Vendor/Model
	4. WebAuthn Extensions
		- **Purpose**: Describes the mechanism for extending the WebAuthn API to suit particular use cases.
		- **Summary**: Lists and explains specific standardized extensions (such as `appid` for FIDO U2F backward compatibility or `credProps` for discovering credential properties) that developers can use.
