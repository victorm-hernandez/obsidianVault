
# Design IAM system

## Original problem statement

1. Clarify authentication, authorization, resource, and policy requirements.
2. Design policy storage, evaluation, propagation, and revocation.
3. Explain online enforcement, caching, availability, and consistency choices.
4. Add ML only where it improves risk detection or adaptive authorization.
5. Define audit logging, evaluation, monitoring, and failure-mode handling.

Do not assume unstated product details. Call out the scale, latency, compliance,
and availability questions you would ask before selecting concrete infrastructure.


-----

## Requirements
- AuthN
- AuthZ: Policy pipeline, storage, evaluation, propagation and revocation.
- Enforcement strategies
- Consistency choices (CAP)
- Where ML improves detection or adaptive AuthZ
- Audit trail should be available (logging, evaluation, monitoring, failure handling)

## Non-Functional requirements
- Call out scale, latency
- Compliance state


## Entities

- Identities
- Access Policies
- Resources

## APIs

Data Plane
- Authenticate(creds) : identity
- IsAllowed(identity) : bool

Control Plane
- CRUD policies
- Revoke policy

## Roles

Identity Provider (AuthN)
- Validate credentials
- Issue access tokens (JWT/SAML/SessionId)
- Expose metadata for token validation (public certs)
- Federate with other IdPs (trust)

Use cases for each
    RBAC / ABAC

    Role: Controls access based on membership, ideal for mover leaver joiner processes. Centralized control of resources shared with who.
        - Source for data: centralized directory or claim mapping (claim source may be federated identity or any other claim)

    Attribute: Controls access based on features, request source IP address range, type of credential used, health of device (intune), ????,
    useful when you want to enforce auth policies (cred type, device health) regardless of the caller (machine auth/agents even)
        - Source for data: Request itself, access tokens. 


