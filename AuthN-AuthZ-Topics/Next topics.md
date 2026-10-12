How to 
- Node attestation specifics, does the public key registers with the IdP? so later the TPM can be used to validate the identity of the workload? how does this work?
	- TPM APIs, are the vaults created on TPM linked to the DLL which created it?
		- If not how to protect, use something like DPAPI? to protect a secret that unlocks the vault?
- Handle revocation of an OAuth token?
- Handle revocation in SPIFFE/SPIRE workloads?
- How to handle PKI revocation?
- Enumerate the RFCs for OAuth, there is so many and each have relevance in technical conversations. 