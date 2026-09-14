# Question: Tell me about yourself / walk me through your background

## Answer

**Victor Hernandez Guzman | Principal Software Engineer**
*Prepared for: Technical Lead Software Engineer Interviews*

---

## 🎤 Short Version (2-Minute Interview Answer)

I have over 20 years of experience in software engineering, with the last 15 focused specifically on Identity and Access Management. Throughout my career, I've had the opportunity to work on some of the most security-critical and large-scale identity systems in the industry.

At Microsoft, I spent nearly a decade on the Microsoft Account and Entra ID teams, where I led some genuinely pioneering work — including architecting the industry's first implementation of passwordless authentication using WebAuthn and FIDO. That was a milestone I'm particularly proud of.

After that, I moved to Salesforce as the technical lead of the Core Authentication team, where my biggest contribution was designing the foundational architecture of a new, unified Identity Provider — built from scratch to consolidate identity systems across multiple acquisitions.

I then returned to Microsoft as a Principal Engineer, this time leading the Datacenter Secure Token Server — a bare-metal identity provider that underpins all of Azure, including Entra ID itself. My work there spanned high-availability design, protocol standardization, and driving a major architectural migration from a stateful to a stateless service model.

After that, I joined DocuSign as a Principal Engineer on the Identity team, where I defined the long-term architectural roadmap and led the transition to passwordless authentication. Following a restructuring in 2025, I was laid off — and rather than immediately jumping to the next role, I used that window as an opportunity to go deep on something I'd been evangelizing internally but never had hands-on time with: applied AI.

I co-founded a small software company called AxoEdge, focused on building AI-powered solutions for real business problems. Our first client was a German-Swiss logistics company that needed to automatically count pallets from photos to track inventory. Computer vision was completely new territory for me, so I ran two tracks in parallel: I ramped up quickly on object detection while simultaneously benchmarking open-source multimodal LLMs against labeled data to determine the right approach. The data made the decision clear — LLMs were too inconsistent for reliable counting, while YOLO-based object detection was predictable and controllable. I delivered a concrete feasibility recommendation that let the client make a confident go/no-go decision, and structured the entire codebase and documentation so another engineer could own it without me.

That experience reinforced something I believe strongly: the best technical leaders can enter an unfamiliar domain, make fast evidence-based decisions, and structure their work so others can build on it.

At this stage of my career, I'm focused on expanding my impact at the organizational level — shaping architecture, setting engineering standards, and elevating the teams around me. That's exactly the kind of role I'm looking for.

---

## 📖 Long Version (Full Narrative)

I have over 20 years of experience in software engineering, with the last 15 years deeply focused on Identity and Access Management. Over that time, I've had the privilege of working on some of the most security-critical, high-scale identity systems in the industry — from consumer-facing products used by hundreds of millions of people to bare-metal infrastructure that supports global cloud platforms.

### Early Career & Foundations

I started my career as a full-stack developer, which gave me a solid grounding across the entire stack — from database design and business logic to front-end interfaces. Early on, I also co-founded a software company, which taught me how to lead teams, manage client relationships, and deliver software end-to-end. I later worked in a technical lead role at HP, where I introduced test-driven development and continuous integration to a geo-distributed team spanning five countries, cutting bug counts and development time significantly.

### Microsoft Account & Entra ID (2009–2019)

In 2009, I joined Microsoft and spent nearly a decade on the identity team. This is where my deep specialization in IAM began. I worked across both Microsoft Account — the consumer-facing identity provider — and what is now known as Entra ID, the enterprise identity platform.

During this period, I led several industry-defining initiatives. Most notably, I architected and delivered the **industry's first implementation of passwordless authentication** using public key credentials — the foundation of what is now WebAuthn and FIDO2. I also owned the integration of Microsoft Account across Windows Core, extending seamless authentication to platforms like Xbox One, Windows 8 and 10, Windows Phone, and HoloLens. Beyond that, I built systems for compromised account detection and remediation, QR code-based session authentication, and push notification login — features used by millions of users globally.

### Salesforce (2019–2022)

I then joined Salesforce as the technical lead of the Core Authentication team. The team owned the authentication and authorization standards — OAuth, SAML, SCIM, and OpenID Connect — that powered the integration of 8 million client applications and over 65 million package installations worldwide.

My biggest contribution there was leading the design and early architecture of a **new, unified Identity Provider** built from the ground up. The goal was to consolidate the fragmented identity silos created by years of corporate acquisitions. I led the critical first phase: architecting the synchronization of user datastores across multiple acquired platforms to establish a single source of truth for identity data.

### Microsoft — Datacenter STS (2022–2024)

I returned to Microsoft as a Principal Engineer to lead the **Datacenter Secure Token Server (STS)** — a bare-metal identity provider that is foundational to the entire Azure ecosystem, including Entra ID itself. This is infrastructure that must work before Azure itself can boot.

My work spanned several dimensions: ensuring high availability and fault tolerance at hyperscale, driving improvements in identity protocols and service reliability, and leading a major architectural evolution — migrating the service from a stateful to a stateless model. That migration involved selecting the right database platform, designing schemas based on IMOS definitions, configuring replication, setting up the deployment pipeline, and coordinating with multiple platform and security teams. I also established "secure-by-design" engineering practices and mentored engineers across the team.

### DocuSign (2024–2025)

I joined DocuSign as a Principal Engineer within the Identity organization, where I defined the long-term architectural roadmap for core authentication and authorization systems, led the company's transition to passwordless authentication, and acted as the primary Architecture Owner across the Identity org. In 2025, following an organizational restructuring after a leadership change, I was laid off.

### AxoEdge — Co-Founder (2025–Present)

Rather than immediately searching for the next role, I chose to treat that transition as a deliberate opportunity. I had been actively evangelizing AI internally at DocuSign — Copilot adoption, MCPs, automation for developer productivity — but all of my hands-on experience had been with closed, paid LLMs. I wanted to go much deeper into open-source AI.

I co-founded **AxoEdge**, a software company focused on applying AI to real business problems, together with another engineer who had also been affected by layoffs.

Our first client was a German-Swiss logistics company that needed to automatically count pallets from warehouse photos to track inventory ingress and egress. Computer vision and deep learning were completely new territory for me, and the hype around multimodal LLMs made the right technical approach genuinely non-obvious.

My goal was to determine viability quickly and avoid committing to the wrong architecture. I ran two tracks in parallel: I ramped up rapidly on object detection workflows while simultaneously evaluating open-source multimodal LLMs — models like LLaVA, Gemma, LLaMA, and Granite. Rather than debating approaches theoretically, I built a benchmarking harness, tested multiple models and prompt strategies against a labeled dataset, persisted results in PostgreSQL, and analyzed accuracy variance and failure modes. I also implemented YOLO-based baselines to compare reliability firsthand.

The data made the decision clear: even the strongest multimodal LLMs were too inconsistent for reliable counting, while object detection was predictable and controllable. I delivered a concrete feasibility assessment and recommended training a vision model on a representative dataset — giving the client the evidence they needed to make a confident go/no-go decision. I also structured the entire codebase and documentation so another engineer could own or extend the work independently.

This project sharpened my ability to rapidly enter an unfamiliar technical domain, make fast evidence-based architectural decisions, and organize work so that knowledge doesn't stay siloed with a single contributor.

### What I'm Looking For

At this stage of my career, I'm seeking to expand my impact at the organizational level. I want to drive architectural strategy, shape engineering culture, and mentor the next generation of technical leaders. I'm drawn to roles where I can combine deep technical expertise with organizational influence — and that's exactly what I bring to a Technical Lead position.

---

*Languages: C#, Java, C++, JavaScript, and web technologies (HTML, CSS, Angular, etc.)*
*Protocols: OAuth 2.0, SAML, SCIM, OpenID Connect, WebAuthn, FIDO2*
