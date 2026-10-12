**

Note: I took the liberty of adding this content as a reference

Provisioning Strategies

Because passkeys rely on asymmetric cryptography, where a private key is generated on the user's local device/authenticator during enrollment, pre-provisioning requires specific enterprise tooling and API support.

1. Hardware Security Keys (FIDO2 Tokens)

If you are deploying physical hardware keys (e.g., YubiKeys or Feitian tokens):

Admin-Led Pre-Registration: Administrators or IT desk support can generate the private key on the physical token and register the public key directly to the user's identity account (such as Microsoft Entra ID or Okta) via administrative APIs or PowerShell scripting (e.g., using Microsoft Graph API for FIDO2 methods). [ [https://agderinthe.cloud/2024/11/04/security-made-easy-pre-provision-fido2-keys-for-your-users-with-powershell/#:~:text=*%20FIDO2%20enabled%20in%20%E2%80%9Cauthentication%20method%E2%80%9D%20for,enabled%20and%20the%20user%20is%20in%20scope](https://agderinthe.cloud/2024/11/04/security-made-easy-pre-provision-fido2-keys-for-your-users-with-powershell/#:~:text=*%20FIDO2%20enabled%20in%20%E2%80%9Cauthentication%20method%E2%80%9D%20for,enabled%20and%20the%20user%20is%20in%20scope). ]

Vendor Pre-Provisioning Services: Vendors (like Yubico's YubiKey Delivery Services) allow you to order keys pre-configured for your enterprise tenant and shipped directly to employees' homes or offices. [[https://www.yubico.com/products/self-service/](https://www.yubico.com/products/self-service/) , CONFIRM]

2. Enterprise Passkey APIs & Identity Platforms

If you are deploying software passkeys across enterprise identity providers:

Directory Provisioning APIs: Platforms like Microsoft Entra ID offer Passkey Provisioning APIs that allow IT administrators to trigger credential creation and auto-bind passkeys to user profiles without waiting for self-service user onboarding.

Enterprise Password Managers: Solutions like Bitwarden, 1Password, or Dashlane allow admins to pre-create vault accounts and push managed enterprise passkeys directly to assigned user accounts via directory sync (SCIM).

3. Managed Device Enrollment (MDM / EMM)

For corporate-managed laptops or mobile devices:

You can push device-bound passkeys or platform authenticators (such as Windows Hello for Business or macOS Touch ID passkeys) during initial device staging using Mobile Device Management (MDM) tools like Intune or Jamf.

When the employee opens their newly staged device and completes zero-touch enrollment (e.g., Windows Autopilot or Apple Automated Device Enrollment), their local passkey binding is generated automatically.

Alternative: Temporary Access Pass (TAP)

If full pre-provisioning is restricted by your identity provider or compliance policies, the standard industry practice is issuing a Temporary Access Pass (TAP) or One-Time Bypass Code.

This allows users to sign in securely on Day 1 without a password so they can register their passkey in a single step upon first login.

Which Identity Provider (e.g., Microsoft Entra ID, Okta, Ping) and passkey types (hardware keys vs. mobile/synced passkeys) are you using for this deployment?

**