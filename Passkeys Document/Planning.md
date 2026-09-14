Collaboration link
https://docs.google.com/document/d/1w_ajZEYL8AgkVtb8WBwSXdfm0vmy7jxi2JhU5hK7p7M/edit?tab=t.0

BoK index
- there are elements missing in our BoK index for some reason
- https://bok.idpro.org/article/id/113/
- https://bok.idpro.org/article/id/102/

 ## BoK resources
 1. Ask about process to publish, and access to github repo. Talk to Elizabeth.
2. Find a reviewer who is an IAM practitioner.
3. Submit first draft of content.
4. Check out https://pandoc.org/ for content transformation
5. Increase the context for the Agent to be able to cross reference to other BoK documents.
	1. Create summaries of every doc in BoK so the Agent knows if it should go a check it out. 
	2. Create an action meant to link with other documents post creation.
	
## BoK resources
ID Pro BoK website https://idpro.org/body-of-knowledge/
ID Pro BoK github https://github.com/IDPros/bok
## Working group
Message on slack about this
- https://idproconnect.slack.com/archives/C0AQT0K45EY/p1787581508911719

People interested in contributing that I pinged already:
- [David Treece, (VP at Yubico),](https://www.linkedin.com/in/treeced/) [david.treece@yubico.com](mailto:david.treece@yubico.com)
- [Libby Brown (PM at Okta)](https://www.linkedin.com/in/libbro/), [libby.brown@outlook.com](mailto:libby.brown@outlook.com)

Other people interested in contributing:
- [Anant Wairagade (Security Engineer at AMEX)](https://www.linkedin.com/in/anant-w-17866720a/), vanant@gmail.com
	- [Passwordless Authentication: Risk, Reward, and Readiness](https://www.isaca.org/resources/news-and-trends/industry-news/2026/passwordless-authentication-risk-reward-and-readiness)
- [Moumita Saha (Security Specialist at Amazon)](https://www.linkedin.com/in/moumita-saha/), moumis@amazon.com

Consider inviting to ID Pro and to contributing to this document:
- [Ricky Mondello](https://www.linkedin.com/in/rmondello/), Principal Eng at Apple
- [Tim Cappalli](https://www.linkedin.com/in/timcappalli/), Architect at Okta
- [Anthony Nadalin](https://www.linkedin.com/in/anthony-n-8a4395/), Partner Architect at Microsoft , Chair of the working group

## W3C publications

Working group: https://www.w3.org/groups/wg/webauthn/
WebAuthN github: https://github.com/w3c/webauthn
Specification: https://w3c.github.io/webauthn/
Mailing list: https://lists.w3.org/Archives/Public/public-webauthn/2026Aug/
Related Spec: [CTAP](https://fidoalliance.org/specs/fido-v2.2-ps-20250714/fido-client-to-authenticator-protocol-v2.2-ps-20250714.html)

## References

- WebAuthN playground: https://webauthn.io/, by [Nick Steele (Identity at OpenAI)](https://www.linkedin.com/in/nickelsteele/)
- WebAuthN guide: https://webauthn.guide/ , by [Suby Raman (Security Infra at Stripe)](https://www.linkedin.com/in/suby-raman-6a051950/)
- Wikipedia: https://en.wikipedia.org/wiki/WebAuthn
- Okta:
	- [WebAuthn - A Short Introduction ](https://auth0.com/blog/webauthn-a-short-introduction/) \, by [Pradheepa Pullanieswaran](https://www.linkedin.com/in/pradheepa/) (Adobe Experience Manager)
- Passkeys playground: 
	- https://passkeys.dev/
	- Advanced: https://passkeys.dev/docs/advanced/
	- Tools: https://passkeys.dev/docs/tools-libraries/test-sites/
	- Specs: https://passkeys.dev/docs/reference/specs/
	- Glossary: https://passkeys.dev/docs/reference/terms/
- FIDO Website
	- https://fidoalliance.org/passkeys/
	- Fido metadata service https://fidoalliance.org/metadata/

## Other References
- Descope: 
	- [WebAuthn: How it Works & Example Flows](https://www.descope.com/learn/post/webauthn)
	- [What Is FIDO2 & How Does FIDO Authentication Work?](https://www.descope.com/learn/post/fido2)
	- [Phishing-Resistant MFA: How It Works and Why You Need It](https://www.descope.com/learn/post/phishing-resistant-mfa)
	- [Strong Authentication: What It Is and Why You Need It](https://www.descope.com/learn/post/strong-authentication)
	- [A Guide to Authentication Protocols](https://www.descope.com/learn/post/authentication-protocols)
	- [How Does Facial Recognition Work, Benefits & Use Cases](https://www.descope.com/learn/post/facial-recognition)
	- [What Is a Passkey & How Does It Work?](https://www.descope.com/learn/post/passkeys)
	- [Biometric Authentication: A Comprehensive Guide](https://www.descope.com/learn/post/biometric-authentication)


---
## Open questions
### 1. Best Place to Add the Article

Depending on the specific focus of your article, there are a few excellent candidate directories in the repository:
- **`Architecture/`**: This is a great fit, especially since it already contains articles like `designing-mfa-for-humans-final.md`. If your article focuses on the implementation, protocol details, or system design of WebAuthN/Passkeys, this is the best place.
- **`Workforce IAM/`**: This directory contains `authentication-methods.md`. If your article focuses on how Passkeys apply to employees and workforce authentication, this is a suitable home.
- **`CIAM/`** (Customer Identity and Access Management): If your focus is on consumer adoption, user experience, and retail authentication with Passkeys, this would be the most appropriate section.
- **`Digital Identity/`**: If the article takes a broader view of Passkeys as a mechanism for establishing and proving digital identity.

### 2. Top 10 Most Recent Articles Added

All of the 10 most recent articles were committed by **Heather Flanagan**. Here they are in chronological order of their addition:

1. **`idpro-bok-incident-response-framework.md`** (Digital Identity) - _Dec 23, 2025_
2. **`non-human-identity-management_-designing-and-governingmachine-actors.md`** (Non-Human Entities) - _Dec 23, 2025_
3. **`hipaa-security-rule-updates-iam-compliance-recommendations.md`** (Laws Regulations Standards) - _Apr 24, 2025_
4. **`pkce-bok.md`** (Laws Regulations Standards) - _Apr 24, 2025_
5. **`authentication-methods.md`** (Workforce IAM) - _Apr 24, 2025_
6. **`optimizing-access-recertification-final.md`** (Workforce IAM) - _Apr 24, 2025_
7. **`intro-to-PAM-v2.md`** (Access Control) - _Nov 28, 2024_
8. **`tokens-in-oauth2.md`** (Access Control) - _Nov 28, 2024_
9. **`An-Introduction-to-OIDC.md`** (Laws Regulations Standards) - _Nov 28, 2024_
10. **`ethics-and-digital-identity-kiser.md`** (Introduction) - _Jul 31, 2024_

### 3. Required Structure for an Article

By reviewing the formatting of existing articles in the repository, a standard IDPro article follows this structure:

1. **Title Header:** e.g., `# Authentication Methods`
2. **Author (Optional but common):** e.g., `By [Your Name]`
3. **Copyright Statement:** e.g., `© 2025 IDPro, [Your Name]`
4. **Standardized Comment Prompt:** An italicized paragraph directing readers to GitHub to leave feedback. You should copy this exactly:
    
    > _*To comment on this article, please visit our [GitHub repository](https://github.com/IDPros/bok) and [submit an issue](https://docs.github.com/en/github/managing-your-work-on-github/opening-an-issue-from-code).*_
    
5. **Introduction Section:** `## Introduction` or `Introduction \n ============`
6. **Terminology Section:** A dedicated section (`## Terminology`) containing a standard Markdown table with three columns: `Term`, `Source`, and `Definition`.
7. **Main Content Sections:** The body of your article.

### 4. Article Word Count Statistics

Excluding structural files like `README.md` and `terminology.md`, there are currently 45 articles in the repository.

- **Average Word Count:** ~4,113 words per article.
- **Smallest Article:** `idpro-article-review-iso_iec-24760-final.md` (in _Laws Regulations Standards_) at **540 words**.
- **Biggest Article:** `intro-to-ciam.md` (in _CIAM_) at **7,692 words**.

