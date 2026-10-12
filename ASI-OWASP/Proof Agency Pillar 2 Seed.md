
Identity establishes who is acting. Authorization governs what that identity may do, and narrows it at every delegation hop. A correctly identified agent with inherited permissions is still an Excessive Agency incident waiting to happen, which is why this pillar matters even when Identity is strong.


- Full-stack attestation: At Level 1, agents operate under an authenticated identity with no anonymous access, and access incidents are tracked. 
    
- Explicit downscoping:At Level 2, the model shifts to default-inherit with explicit downscoping: per-agent permission boundaries, writes that require approval, permissions calibrated to data-sensitivity tiers, and separated sandbox, staging, and production environments. Token exchange swaps a human session for a scoped, short-lived agent token, the pattern defined in RFC 8693 [9].
    
- Policy-engine enforcement: At Level 3, a policy engine applies dynamic, risk-adjusted scoping, and the agent itself enforces policy across all output paths rather than trusting upstream checks. Per-request scopes are expressed through Resource Indicators [10] and Rich Authorization Requests [11], and tokens are sender-constrained through mutual-TLS binding or DPoP so a stolen token alone cannot act [12][13].
    
- Purpose-bound JIT delegation: Levels 4 and 5 carry the weight of multi-agent delegation, the hardest authorization problem agents introduce. Agents rarely act alone: a user asks an agent to do something, that agent calls a tool, exchanges a token, and delegates to another agent, and each hop can dilute context and create a confused-deputy risk. At Level 4, permissions are scoped to a declared purpose per request, provisioned just in time with automatic expiry, and gated by a rule-of-two control for critical operations. Each hop receives narrowed, recipient-bound authority, so delegation narrows scope instead of impersonating the original principal. 
    
- Cross-domain verifiable authority: At Level 5, authority boundaries become machine-verifiable across trust domains, with real-time consent verification and agent-to-agent chains whose scope continuously narrows. Authority is granted only because the current request, runtime attestation, delegation path, and policy together justify it.