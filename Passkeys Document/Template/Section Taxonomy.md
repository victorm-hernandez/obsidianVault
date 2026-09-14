
## Section Template

```markdown
{SECTION TITLE}
======================

{SECTION CONTENT}
[^{Foot note number}]

{SECTION CONTENT, GRAPHIC/DIAGRAM (IMAGE TAG/MERMAID)}
Figure {Figure #}. {Figure description}

...

---

{Foot note number}. {Author, Title, Date, Source/URL}

...
```

## Example 

```markdown

The Role of Identity Providers and Federation
=============================================

The revelation that passwords are fundamentally flawed is not new -
dating back to at least the '70s, there has been research on how to get
around the need for a human brain in the authentication process.
[^13], [^14] We developed the simple idea that passwords are "something
you know," but also described other options for validating a human's
ownership of a digital account could also include "something you have"
or "something you are". The idea is not that validating the thing you
have can replace the thing you know, but rather that a combination of
things you have, are, and know would require an attacker to compromise
both digital and physical information. Today, the state of the art in
multi-factor authentication is very sophisticated. A growing number of
users protect their phone with a biometric, navigate an SMS message to
confirm a transaction, or use an OTP (one-time password) to improve
security without any need to understand the underlying principles.

We all know that MFA must continue to improve in usability to become
ubiquitous. Specifications like FIDO2 are industry-changing for access
management, not because the problem is solved - but because the problem
is ***decoupled*** - FIDO2 (W3C WebAuthn and FIDO CTAP2) has separated
the problem of negotiating cryptographic keys from the problem of
requiring user gestures. [^15] The
cryptographic key exchange can now stay reliable, while we focus on
innovation - and possibly even revolution - in user interactions.

<img src="Graphics/IdentiBeerLogo-Seattle_upscayl_5x_digital-art-4x.png" alt="Diagram Description" />
Figure 1: The user and their different authentication factors

---


13. Raible, Matt, "What the Heck is OAuth?" DZone Security Zone, 28
    April 2018, <https://dzone.com/articles/what-the-heck-is-oauth> .
   

14.  Wikipedia contributors, \"Federated identity,\" Wikipedia, The Free
    Encyclopedia,
    <https://en.wikipedia.org/w/index.php?title=Federated_identity&oldid=949399706>
    (accessed June 6, 2020). 

15.  Wikipedia contributors, \"Principle of least privilege,\" Wikipedia,
    The Free Encyclopedia,
    <https://en.wikipedia.org/w/index.php?title=Principle_of_least_privilege&oldid=950981064>
    (accessed June 6, 2020).

16.  Rose, Scott, and Oliver Borchert, Stu Mitchell, Sean Connelly, "Zero
    Trust Architecture (2 ^nd^ Draft)," SP 800-207 (Draft), National
    Institute of Standards and Technology, February 2020,
    <https://csrc.nist.gov/publications/detail/sp/800-207/draft> .
   

17.  Paul A. Grassi, James L. Fenton, Elaine M. Newton, Ray A. Perlner,
    Andrew R. Regenscheid, William E. Burr, and Justin P. Richer. 2017.
    Digital identity guidelines - Authentication and Lifecycle
    Management. Technical Report. NIST Special Publication 800-63B.
    

18.  Graham Williamson and Corey Scholefield. Introduction to IAM Project
    Management for IAM Projects. IDPro Body of Knowledge, volume 1,
    issue 1, 31 March 2020. <https://bok.idpro.org/article/id/25/> .
```
