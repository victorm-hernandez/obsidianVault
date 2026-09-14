# Short Version

As a Principal Software Engineer on the dSTS identity provider (baremetal tier 0), I led the end-to-end design and implementation of an internal WS-FED/SAML token validation library so Azure services in the lower tiers could securely validate dSTS-issued tokens. I owned technical direction, worked closely with the ID4S, dSTS and our first consumer the Instrumentation team, delivered a component that is now used by all tier-0 services across Azure.

# Polished STAR++ Narrative

## Situation

I was a Principal Software Engineer on dSTS, the baremetal tier-0 identity provider for service-to-service authentication in Azure. Microsoft was standardizing identity solutions across services, however there still wasn't a modern library for dSTS and the team that owned the Identity for Services (ID4S) libraries didn't have enough resources to provide a new solution for dSTS. dSTS is different from other identity providers at Microsoft because it uses WS-FED and SAML to issue tokens, has different endpoints for metadata, authentication, and has a different set of claims in the tokens. It also has different certification rotation strategies, caching strategies, and delegation models than OAuth/OIDC.

## Task

I owned the design and implementation of a reusable internal validation library to let ID4S and Azure service teams validate dSTS-issued SAML tokens safely, while handling WS-* protocol differences and meeting security requirements.

## Actions

- I decided to build on top of the existing ID4S libraries rather than creating a identity provider-specific fix, because even if dSTS’s WS-FED/WS-TRUST flow, metadata endpoints, cert rotation, and claim model differed from standard OAuth/OIDC there was still a lot of shared logic that could be reused and maintained in one place.

- I prioritized compatibility-first validation logic and explicit support for dSTS claim sets, there should be zero service disruption and feature parity with the now deprecated dSTS validation libraries. I also prioritized security and performance, since this library would be used by many low-level Azure services to validate tokens.

- Implementation highlights:
    -The library included token parsing, claim handling, metadata discovery, and certificate rotation support.
    - Collaborated with the ID4S and dSTS teams to validate integration with existing libraries and confirm the library worked with live dSTS tokens.
    - Created support documentation, SOPs, and integration guidance so other teams could adopt the library reliably.

## Result

The library was successfully integrated into ID4S’s existing libraries and services and is now used by many low-level Azure services to validate dSTS-issued tokens, reducing token-validation fragmentation and improving secure service access across Azure.

## Reflection & Tradeoffs

I learned that cross-team identity work needs early compatibility validation and clear support docs; if I repeated it, I’d add explicit adoption and error-rate metrics up front and build regression tests for each dSTS token variant. I also considered a more generic middleware approach, but the targeted library was the safer path given the baremetal WS-* differences.

# Quick Pass/Fail Checks

- Situation: yes
- Task: yes
- Actions: yes
- Result: yes
- Reflection: yes

# STAR++ Score

Score: 8/10 — strong ownership and technical decisions, with a bit more metric specificity needed.

# Strengths

- Clear ownership of end-to-end design and implementation.
- Strong technical signal around WS-FED/SAML compatibility and team collaboration.

# Weaknesses & Concrete Improvements

- Weakness 1 → No explicit quantitative metric in the result; fix by adding adoption or security impact numbers if available.
- Weakness 2 → Reflection is a bit general; fix by naming one specific alternative architecture and one concrete lesson.

# One-line Rewrite Suggestion

I led design and delivery of a shared dSTS WS-FED/SAML token validation library that integrated with ID4S, supported live dSTS tokens, and is now used by many Azure services.

# Time-to-Deliver Check

Trimable to ~2 minutes? yes — remove one implementation bullet or shorten the result sentence.

# Two Quick Tips to Increase Technical Signal

- Add a concrete adoption or security metric, even a labeled estimate.
- Mention one specific compatibility test or failure mode you solved.

# Suggested Interviewer Follow-ups

- Q1: How did you validate compatibility with dSTS-issued tokens?
- Q2: What security tradeoffs did you make when supporting WS-FED vs OAuth/OIDC?
- Q3: How did you ensure the library was adopted and supported by other service teams?

# Metadata

- Role: Principal Software Engineer
- Timeframe: not provided
- Metrics provided: no
- Estimates used: no
- 

---
ORIGINAL TEXT

# Identity For Services (ID4S)

I was a Principal Software Engineer on dSTS (datacenter security token service) — Azure's baremetal, tier-0 identity provider.

At the time there was a company wide effort to standardize identity solutions across Microsoft, and the ID4S team was responsible for providing identity solutions for services, however this team didnt have enough resources to provide a new solution for my identity provider which used WS-FED and SAML to issue tokens. I was brought in to help the team design and implement a new internal library that would allow services to validate tokens issued by dSTS. This library was used by many all low-level services to validate tokens and ensure secure access to resources across Azure.

Some of the key differences between dSTS and other identity providers at Microsoft is that dSTS is a baremetal, uses the WS-* protocols (WS-FED, WS-TRUST), uses SAML to issue tokens, uses a different set of endpoints for metadata, authentiation, and has a different set of claims in the tokens. Different certification rotation strategies, different caching strategies, the delegation model for WS-FED was different than that of OAuth/OIDC (on-behalf-of vs delegation).

I was responsible for this project from start to finish, including designing the library, implementing it, and working with the ID4S team to integrate it into their existing libraries and services. I also worked with the dSTS team to ensure that the library was compatible with the tokens issued by dSTS and that it met the security requirements of the service and finally worked with create a baseline for support, SOPs and documentation for the library. The library was successfully integrated into the ID4S team's existing libraries and services, and it is now used by many services across Microsoft to validate tokens issued by dSTS.

Once this project was completed, I was able to contribute to the development of other libraries and services owned by the ID4S team. The ID4S team is responsible for providing identity solutions for services across Microsoft, and they own a number of libraries and services that are used by developers to integrate authentication and authorization into their applications.

Some of the libraries are public facing and open source like the Microsoft Authentication Library (MSAL) and the IdentityModel Extensions for .NET. These libraries are used by developers to integrate authentication and authorization into their applications, enabling secure access to Microsoft services and APIs.

Some resources about the public facing libraries owned by the ID4S team: 
- [What is MSAL?](https://www.youtube.com/watch?v=zufQ0QRUHUk)
- [MSAL GitHub](https://github.com/AzureAD/microsoft-authentication-library-for-dotnet/)
- [Identity Model GitHub](https://github.com/AzureAD/azure-activedirectory-identitymodel-extensions-for-dotnet)