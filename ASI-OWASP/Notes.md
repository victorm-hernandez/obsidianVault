## Main document
https://docs.google.com/document/d/1ZHB4o5yn7XJcIN3_Ya35hD-Jta8G4TMKAkyHTHK9Jxw/edit?tab=t.b1hndi14gzya#heading=h.jg6csf9bilyt

Intro document (read before contribute)
https://docs.google.com/document/d/1xRuHe29xMQ1AYFxNd0YGWk1D40HtYKQrM6At7X0Uxn8/edit?tab=t.2fu5dov5ht9

---
## Useful references

Securing Agentic Applications
https://docs.google.com/document/d/116wH4v5d41ZUsFu6kcckJiurHMM81kYw99fgqvYk7kU/edit?tab=t.a4jjur8sa79u

OWASP GenAI LLM Top 10 2026
https://genai.owasp.org/resource/owasp-genai-llm-top-10-2026/

State og Agentic AI Security and Governance
https://genai.owasp.org/resource/state-of-agentic-ai-security-and-governance/

Microsoft Entra Agent ID
https://learn.microsoft.com/en-us/entra/agent-id/what-is-microsoft-entra-agent-id

Transactional Access Tokens
https://www.ietf.org/archive/id/draft-tulshi-oauth-transactional-access-tokens-00.html

---

## Concepts
### RFC 8693: OAuth Token Exchange
https://www.rfc-editor.org/info/rfc8693/

### WIMSE

https://datatracker.ietf.org/wg/wimse/about/

WIMSE stands for Workload Identity in Multi-System Environments, an active IETF Working Group chartered to standardize how software workloads (running instances of applications, microservices, or AI agents) authenticate, obtain credentials, and securely interact across heterogeneous cloud platforms and networks.

Core Focus and Goals
• Machine-to-Machine Trust: Focuses strictly on service-to-service authentication and authorization, distinguishing itself from human user login flows or general enterprise directory IAM.
• Interoperability: Bridges gaps between disparate identity formats—such as JSON Web Tokens (JWT), SPIFFE IDs, mutual TLS (mTLS) certificates, and Remote Attestation Tokens—to eliminate custom "glue" code.
• Standardized Architecture: Developing reference architectures, token formats, and protocols (like HTTP message signatures and workload proof tokens) for secure, least-privilege workload communication


### RATS / EAT
https://www.rfc-editor.org/info/rfc9334/
https://www.rfc-editor.org/info/rfc9711/

RATS (Remote ATtestation ProcedureS) and EAT (Entity Attestation Token) are IETF standards used to cryptographically prove and verify the hardware and software identity, security state, and configuration of a device to a remote service.
While traditional user or server authentication confirms who a user or server is via passwords or certificates, RATS and EAT establish trust in what a device is—proving that its hardware is genuine and its operating software is secure and uncompromised.

Core Concepts

• RATS (Remote Attestation Procedures): An IETF Working Group architecture that defines how a device (the Attester) generates and sends proof of its security posture to a checker (Verifier), which then informs a service (Relying Party) whether the device is safe to trust.
• EAT (Entity Attestation Token): Defined in RFC 9711, an EAT is a signed, machine-readable data structure (built on JSON Web Tokens or CBOR Web Tokens) that packages identity claims about a physical device, chip, or software module.

What EAT Identifies and Measures

An EAT goes far beyond a simple device serial number by carrying verifiable claims such as:
• Unique Device ID (ueid): A globally unique hardware identifier (like a serial number) for a specific phone, IoT sensor, or processor chip.
• Software State: Cryptographic measurements of the bootloader, operating system, and firmware to prove no tampering occurred.
• Security Capabilities: Indicators showing if secure boot is enabled, debugging tools are turned off, or what hardware trust-zone level is active.

How It Works in Identity Systems

1. Evidence Collection: During device bootup, a secure hardware root-of-trust gathers data on the system's components.
2. Token Issuance: The device compiles this information into an Entity Attestation Token (EAT) and signs it.
3. Verification & Trust: A remote application or relying party receives the token, checks its cryptographic signature, and decides whether to grant the device access, allow a high-value financial transaction, or restrict network connectivity.

### W3C VC , Verifiable Credentials

https://www.w3.org/TR/vc-overview-1.0/

A W3C Verifiable Credential (VC) is a tamper-proof, cryptographically signed digital version of physical credentials like a driver’s license, university diploma, or employee badge.
Standardized by the World Wide Web Consortium (W3C), it gives individuals a universal way to store, manage, and share trusted personal data online without relying on a central database.

How It Works

The W3C VC ecosystem relies on three main roles:
• Issuer: An organization or authority (like a university or government) that creates, signs, and issues the digital credential.
• Holder: The individual who stores the credential safely in a digital wallet on their phone or device.
• Verifier: A third party (like an employer or a bank) that requests and checks the digital credential to confirm it is real.

Key Features

• Tamper-Evident: Every credential includes a digital signature from the issuer. Any change to the data invalidates the signature.
• Privacy-Enhancing: VCs support advanced privacy features like selective disclosure, meaning you can share only the specific data required (such as proving you are over 21 without showing your exact birth date).
• Interoperable: Because it follows an open web standard, a credential issued by one organization can be read and verified by software built by an entirely different developer or company


### SPIFFE

https://spiffe.io/

SPIFFE (Secure Production Identity Framework for Everyone) is an open-source standard that defines a universal, cryptographically verifiable identity framework for software workloads across heterogeneous and cloud-native environments. It eliminates the need for long-lived secrets like passwords, API keys, or static credentials by automatically issuing and rotating short-lived certificates. Learn more at the SPIFFE Official Website.

Core Concepts

• SPIFFE ID: A unique URI string (formatted as spiffe://trust-domain/path) that identifies a specific software workload.
• Trust Domain: The root of trust boundary, acting as the authority that manages keys and issues identities for its internal workloads.
• SVID (SPIFFE Verifiable Identity Document): A cryptographically verifiable document—such as an X.509 certificate or JWT—containing the SPIFFE ID used for mutual authentication.
• Workload API: A standardized, environment-agnostic API that local workloads call to securely fetch and refresh their SVIDs without prior credentials.

Implementation & Ecosystem

• SPIRE (SPIFFE Runtime Environment): The primary open-source reference implementation that handles node/workload attestation, SVID issuance, and automatic rotation.
• Use Cases: Enabling mutual TLS (mTLS) between microservices, passwordless database access, and secure zero-trust federation across multi-cloud infrastructure