### Glossary
- CRL, certificate revocation list
- OCSP, Online Certificate Status Protocol
- TLS, transport layer security.
- Ephemeral Diffie-Hellman (ECDHE). Algorithm to safely share a secret key without an encrypted channel.

### TLS 1.3 (Current)

1. **ClientHello (Client $\rightarrow$ Server):** The client initiates the connection by sending its supported cipher suites, TLS version, and its Diffie-Hellman (DH) key share.
2. **ServerHello (Server $\rightarrow$ Client):** The server locks in the cryptographic algorithm (cipher suite) and responds with its own DH key share. _Because both sides now possess the required key halves, the shared symmetric secret is generated immediately. All subsequent handshake messages from this exact point forward are encrypted._
3. **Server Authentication (Server $\rightarrow$ Client):** The server sends an `EncryptedExtensions` message (to handle minor protocol parameters), followed by its `Certificate` and a `CertificateVerify` message containing a digital signature to prove ownership of the certificate's private key.
4. **Server Finished (Server $\rightarrow$ Client):** The server sends a cryptographic hash of the entire handshake to prove no intermediary tampered with the initial plaintext Hello messages.
5. **Client Finished (Client $\rightarrow$ Server):** The client verifies the server's certificate and Finished hash, then sends its own encrypted `Finished` message to finalize the handshake.
6. **Application Data:** Secure, symmetrically encrypted application traffic begins.
   
### TLS  1.2 (Obsolete)

1. **The client initiates (Correction to Step 1):** The handshake begins with the client sending a **ClientHello**, advertising its supported cipher suites and TLS version. 
2. The server replies (**ServerHello**) to lock in the selected algorithm and its **public certificate**.
3. **Validation:** The client validates expiration, the signature chain up to a trusted root, revocation status (CRL/OCSP), and verifies that the Subject Alternative Name (SAN) matches the requested hostname.
4. Client uses asymmetric encryption, it uses the server public certificate to encrypt a cryptographic key and send it over the wire.
5. **Symmetric encryption:** Once the shared secret is sent to the server and decrypted, symmetric encryption (like AES-GCM) takes over for all application data.
#### Diffie-Hellman Algorithm

Imagine you and a friend want to create a secret paint color to use as a password, but you have to share your ingredients by shouting across a crowded room where everyone can hear you.

Here is how you establish a shared secret without anyone in the room figuring it out.

1. **Agree on a public base:** You and your friend shout across the room: "Let's start with **Yellow**!" Everyone listening now knows your base color is Yellow.
2. **Pick your secret colors:** You secretly pick **Red**. Your friend secretly picks **Teal**. You _never_ say these colors out loud or send them to each other.
3. **Mix them up:**
	- You mix the public Yellow with your secret Red. You get **Orange**.
	- Your friend mixes the public Yellow with their secret Teal. They get **Light Blue**.
4. **Exchange the mixtures:** You pass your Orange mixture to your friend. Your friend passes their Light Blue mixture to you. Everyone in the room sees Orange and Light Blue being exchanged._(Crucial point: It is easy to mix paint, but practically impossible to un-mix it. An eavesdropper cannot look at the Orange paint and easily separate it back into Yellow and your secret Red)._
5. **Add your secrets one last time:**
    - You take your friend's Light Blue mixture and add your secret Red to it.
    - Your friend takes your Orange mixture and adds their secret Teal to it.

**The Result:** You both ultimately mixed Yellow + Red + Teal. You now both possess the exact same muddy **Brown** color.

## Mutual TLS
Almost exactly the same as TLS but with extra steps to validate the identity of the client. 

1. **ClientHello (Client $\rightarrow$ Server):** The client initiates the connection by sending supported cipher suites, TLS version, and its DH key share.
2. **ServerHello (Server $\rightarrow$ Client):** The server locks in the cipher suite and responds with its DH key share. _(Encryption starts here)._
3. **Server Authentication & Certificate Request (Server $\rightarrow$ Client):**    
    - **[mTLS Addition]** The server sends a `CertificateRequest` message, telling the client "I need you to prove who you are."        
    - The server sends its own `Certificate` and `CertificateVerify`.
4. **Server Finished (Server $\rightarrow$ Client):** The server sends its cryptographic hash of the handshake so far.
5. **Client Authentication (Client $\rightarrow$ Server):**
    - **[mTLS Addition]** The client sends its own `Certificate` in response to the server's request.        
    - **[mTLS Addition]** The client sends its own `CertificateVerify` message (a digital signature made with its private key) to prove it owns the certificate it just sent.
6. **Client Finished (Client $\rightarrow$ Server):** The client sends its `Finished` message (the cryptographic hash of the handshake).
7. **Application Data:** Secure, mutually authenticated, and encrypted application traffic begins.