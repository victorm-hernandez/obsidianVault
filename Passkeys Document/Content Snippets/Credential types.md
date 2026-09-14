In the WebAuthn specification, credentials are divided into two main categories based on **where the data is stored**.

These two types are **Discoverable Credentials** (often colloquially referred to as "Passkeys") and **Non-Discoverable Credentials**.

---

### 1. Discoverable Credentials (formerly "Resident Keys")

**What it is:** A Discoverable Credential means that the private key, along with the user's metadata (like their username, display name, and the user ID), is **stored locally on the physical authenticator itself** (e.g., taking up memory on your YubiKey or inside your iPhone's Secure Enclave).

**Why it's called "Discoverable":** Because the credential lives on the device, the server doesn't need to know who you are before asking for authentication. The server can simply ask the authenticator: _"Do you have any accounts for `example.com`?"_ The authenticator looks in its local memory, discovers the accounts, and shows the user a UI prompt saying: _"Which account do you want to use? Alice or Bob?"_

**The Use Case: "Usernameless" Login (Passkeys)** This enables the most frictionless login experience possible. The user goes to a website, clicks "Sign In", touches their fingerprint sensor, and they are immediately logged in. They never have to type their username or password.

---

### 2. Non-Discoverable Credentials (or "Server-Side" Credentials)

**What it is:** Authenticators (especially hardware security keys) traditionally have very little internal memory. To avoid running out of space, the authenticator does something clever: it generates the private key, but instead of saving it, it encrypts (wraps) the private key using a master hardware secret. It then sends this encrypted blob—called the **Credential ID**—to the server. **The server stores the credential, not the authenticator.**

**How it works:** Because the authenticator didn't save anything, it has no memory of the user. If the server asks, _"Do you have any accounts for `example.com`?"_, the authenticator will say _"I don't know."_

Instead, the authentication ceremony must work like this:

1. The user types their username into the website.
2. The server looks up the user's encrypted **Credential ID** in its database.
3. The server sends the Credential ID to the authenticator.
4. The authenticator uses its master hardware secret to decrypt the Credential ID, reconstructing the private key on the fly, and uses it to sign the login challenge.

**The Use Case: Traditional Two-Factor Authentication (2FA)** Because this requires the user to identify themselves first, it necessitates a **"username-first" flow**. This is how traditional U2F security keys have worked for years. It is highly scalable because a single $20 security key can be registered on an infinite number of websites without ever running out of storage space.

---

### Summary of Differences

|Feature|Discoverable Credential (Resident Key)|Non-Discoverable Credential|
|---|---|---|
|**Storage Location**|Takes up memory on the authenticator|Stored in the server's database|
|**Login Flow**|"Usernameless" (User just clicks Sign In)|"Username-first" (User types username, then uses key)|
|**Storage Limit**|Hardware limited (e.g., a hardware key might only hold ~25 to 100 credentials)|Practically infinite|
|**Modern Terminology**|Often branded as **Passkeys**|Standard 2FA Security Key|

_(Note: During registration, a developer can dictate which type they want by setting the `residentKey` property to `"required"`, `"preferred"`, or `"discouraged"` in the `authenticatorSelection` options)._
