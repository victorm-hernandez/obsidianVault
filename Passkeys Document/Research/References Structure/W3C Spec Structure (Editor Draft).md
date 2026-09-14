### Abstract

- **Purpose**: Provides a high-level overview of the entire specification.
- **Summary**: Defines an API that enables the creation and use of strong, attested, scoped, public key-based credentials by web applications for the purpose of strongly authenticating users.

### Status of this document

- **Purpose**: Describes the current state of the document in the W3C standardization process.
- **Summary**: Indicates whether the document is a working draft, candidate recommendation, or full recommendation, and how it can be used or cited.

### Table of Contents

- **Purpose**: Provides navigation for the document.
- **Summary**: A standard list of links to all the normative and non-normative sections within the specification.

### 1. Introduction

- **Purpose**: Offers a non-normative overview of WebAuthn and its architecture.
- **Summary**: Explains the basics of public key credentials, how they are created by a WebAuthn Authenticator at the behest of a WebAuthn Relying Party (a web application), and how they improve security by avoiding shared secrets like passwords.

### 2. Conformance

- **Purpose**: Defines the criteria that implementations must meet to comply with the specification.
- **Summary**: Specifies the different conformance classes (like WebAuthn Relying Party, WebAuthn Client, and WebAuthn Authenticator) so that compliant implementations can interact securely and predictably.

### 3. Dependencies

- **Purpose**: Lists other technical specifications that WebAuthn relies upon.
- **Summary**: Details external documents and normative references (such as RFCs for terminology like "MUST", "SHOULD") that form the foundational rules for implementing WebAuthn.

### 4. Terminology

- **Purpose**: Defines specific terms and concepts used throughout the document.
- **Summary**: Provides definitions for key concepts like "Authenticator," "Relying Party," "Public Key Credential," and other specialized vocabulary to ensure clarity and consistency.

### 5. Web Authentication API

- **Purpose**: Normatively specifies the JavaScript API for creating and using public key credentials.
- **Summary**: Explains the core browser methods (`navigator.credentials.create()` and `navigator.credentials.get()`). It details how scripts request permission to perform authentication operations on the user's behalf without ever exposing the private credential material to the script.

### 6. WebAuthn Authenticator Model

- **Purpose**: Describes the abstract functional model for a WebAuthn Authenticator.
- **Summary**: Explains how client platforms interact with authenticators (like security keys, biometric sensors). While client platforms can implement this model in any way, the exposed behavior must remain consistent with the Web Authentication API requirements.

### 7. WebAuthn Relying Party Operations

- **Purpose**: Outlines the steps and operations a Relying Party (server) must perform.
- **Summary**: Details how a Relying Party creates the configuration options for registration (`PublicKeyCredentialCreationOptions`) or authentication (`PublicKeyCredentialRequestOptions`), and how it should validate the cryptographic responses it receives from the client.

### 8. Defined Attestation Statement Formats

- **Purpose**: Defines how authenticators prove their provenance and characteristics to a Relying Party.
- **Summary**: Specifies an initial set of pluggable attestation statement formats (like Packed, TPM, Android Key Attestation, Apple Anonymous) that verify an authenticator's make and model.

### 9. WebAuthn Extensions

- **Purpose**: Describes the mechanism for extending the WebAuthn API to suit particular use cases.
- **Summary**: Details how client extensions work during credential creation or assertion requests, allowing a Relying Party to request additional processing or data from the client and the authenticator.

### 10. Defined Extensions

- **Purpose**: Specifies a set of official extensions registered in the IANA "WebAuthn Extension Identifiers" registry.
- **Summary**: Lists and explains specific standardized extensions (such as `appid` for FIDO U2F backward compatibility or `credProps` for discovering credential properties) that developers can use.

### 11. User Agent Automation

- **Purpose**: Defines extensions for automated testing of WebAuthn implementations.
- **Summary**: Specifies WebDriver extension commands that allow for the automation of user agents (browsers) and web application testing when dealing with WebAuthn flows (e.g., simulating authenticator interactions).

### 12. IANA Considerations

- **Purpose**: Outlines the registries and namespaces managed by IANA for WebAuthn.
- **Summary**: Details the formal registration of identifiers, formats, and extensions required by the protocol to ensure global uniqueness and interoperability.

### 13. Security Considerations

- **Purpose**: Analyzes the security model, threats, and mitigations related to the specification.
- **Summary**: Discusses how the WebAuthn protocol protects against common attacks like phishing and credential stuffing. It references FIDO Alliance documents for detailed security analyses and authenticator security requirements.

### 14. Privacy Considerations

- **Purpose**: Discusses the privacy implications of using WebAuthn.
- **Summary**: Applies FIDO privacy principles, detailing how the API prevents tracking users across different Relying Parties and providing specific privacy guidance for authenticator, client, and Relying Party implementers.

### 15. Accessibility Considerations

- **Purpose**: Addresses how to make WebAuthn implementations usable for people with disabilities.
- **Summary**: Recommends that authenticators offer multiple user verification methods (e.g., both biometrics and PINs). It also advises Relying Parties to provide features (like naming the authenticator) that help users remember and easily interact with their credentials.

### 16. Test Vectors

- **Purpose**: Provides non-normative examples to help developers validate their implementations.
- **Summary**: Lists example values as pseudocode for both registration and authentication ceremonies using the same credential, heavily annotated to explain the expected binary structures.

### 17. Acknowledgements

- **Purpose**: Recognizes individuals and groups who contributed to the specification.
- **Summary**: Thanks specific members, co-chairs of the Web Authentication Working Group, and W3C Team Contacts for their work in shaping and finalizing the document.

### 18. Revision History

- **Purpose**: Tracks changes made to the specification over time.
- **Summary**: A non-normative summary of the significant modifications, additions, and clarifications that have occurred across different drafts and versions of the document.