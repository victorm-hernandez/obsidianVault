# Projects Summary

Quick-reference document for behavioral interviews. Use it to pick the right story for any question.

---

## Theme → Project Map

| Theme | What interviewers look for | Projects |
|---|---|---|
| **Failure** | Accountability + learning velocity | Win8 Protocol Disagreement · Xbox Live Site Incident · Salesforce Global IDP |
| **Conflict** | Communication + pragmatism | Win8 Protocol Disagreement · dSTS Rapid Team Growth |
| **Leadership** | Influence without authority | MSA WebAuthn · dSTS Rapid Team Growth · Salesforce Global IDP · Remix Search Service |
| **Ambiguity** | Decision-making with incomplete data | AxoEdge / ProLogistiks · MSA WebAuthn · Remix Search Service · Salesforce Global IDP |
| **Impact** | Business + technical results | MSA WebAuthn · Xbox Live Site Incident · Salesforce Global IDP · Remix Search Service |
| **Execution** | Delivering under constraints | Remix Search Service · MSA WebAuthn · Xbox Live Site Incident · Salesforce Global IDP |

---

## Project Summaries

### AxoEdge / ProLogistiks — Pallet Detection Feasibility Study
**Role:** Co-founder / sole engineer — software consulting startup post-DocuSign layoff.  
**What happened:** A logistics client needed to count pallets from warehouse photos. Computer vision was a new domain. Built parallel evaluation tracks: benchmarked 7 open-source multimodal LLMs and YOLO-based detection baselines against a labeled dataset. The best LLMs showed 40% count deviation at ~10 s/image; YOLO reached ~80% accuracy orders of magnitude faster. Delivered a clear go/no-go recommendation in one month backed by reproducible experiments.  
**Key signals:** Rapid learning in an unknown domain · data-driven architecture decision · parallel experimentation under time constraint.

---

### MSFT — dSTS Rapid Team Growth: Low Morale & Unstable Service
**Role:** Tech lead, dSTS (Microsoft).  
**What happened:** Team scaled from 5 to 30+ engineers in under a year. A chain of incidents left the team burnt out and knowledge dangerously siloed. Ran structured sessions to surface pain points bottom-up, seeded starter proposals to anchor discussion, then facilitated the team in owning and extending them — landing 8+ process changes (documentation standards, release manager rotation, automated release notes, PR review standards, branch cleanup policy). Incident turnaround improved [est. ~30%] within ~3 months.  
**Key signals:** Leadership without authority · bottom-up change management · team health under scale pressure.

---

### MSFT — MSA WebAuthn (FIDO2 Passwordless Authentication)
**Role:** Senior Software Engineer, Microsoft Account UX team.  
**What happened:** Led the browser/UX-layer integration of WebAuthn — the first major identity provider to ship FIDO2 passwordless authentication — for hundreds of millions of MSA users. Coordinated across four teams (MSA, Edge, Windows, W3C PM group) against a hard Windows 10 ship deadline while the W3C spec was still a working draft. Advocated for requiring user-verification-only authenticators to deliver full passwordless (not just second-factor) at launch. Delivered on time and fed requirements back into the W3C spec process.  
**Key signals:** Cross-functional coordination · delivering on a hard deadline against a moving spec · technical influence on industry standards.

---

### MSFT — MSA Win8 Protocol Disagreement
**Role:** Senior Software Engineer, Microsoft Account UX team (~2011).  
**What happened:** Needed the protocols architect to sign off on a WS-FED customization to enable OneDrive roaming on Windows 8. The first meeting ended in a walkout — I was defending, not listening. Recognized the mistake, reset using active paraphrasing, and co-designed a compromise that protected protocol integrity while shipping the feature. The feature shipped on time for Windows 8 and unblocked OneDrive roaming.  
**Key signals:** Conflict resolution · listening posture shift · cross-team negotiation without authority.

---

### MSFT — Xbox Live Site Incident (Login Regression)
**Role:** Engineer, Xbox / Microsoft Account Integration (~late 2013).  
**What happened:** An hour after a routine deployment, an alert fired: login success collapsed to <20% of baseline on a flow processing tens of millions of logins per month. No errors in any log or environment. Exhausted all remote diagnostics, then reproduced the issue using a physical Xbox One controller — a junior engineer's custom WinJS bundle had silently stripped the controller-input mapping module. Rebuilt the bundle, validated in PROD via HTTP proxy injection, and hot-deployed by flushing the CDN cache. Fully resolved ~3 hours after detection. Added controller-emulation automation post-incident.  
**Key signals:** Debugging under pressure · incident ownership · prevent-recurrence follow-through.

---

### MSFT — Remix 3D Search Service Rebuild
**Role:** Senior Software Engineer / Tech Lead, Remix 3D (Microsoft).  
**What happened:** The team's core search was owned by an England-based team running a legacy Zune-era service. A hard GDPR deadline (May 2018) required changes the external team could not deliver in time. Evaluated ramping up on their codebase versus a rebuild; chose rebuild given high ramp-up cost and unsupportable technical debt. Built a new Service Fabric microservice over Elasticsearch with a query transformation layer (tokenization, fuzzy/synonym search, dynamic Lucene queries), authentication, monitoring, and deployment pipelines — solo, in 3 months. Achieved p95 < 100 ms and cleared the team's last GDPR compliance blocker.  
**Key signals:** Build vs. borrow decision · solo delivery under a hard compliance deadline · full-stack technical ownership.

---

### Salesforce — Auth Global IDP (Phase 1 User Sync)
**Role:** Lead Software Engineer (Lead MTS), Salesforce.  
**What happened:** Salesforce greenlit a new centralized Identity Provider to unify identity across all acquired clouds. As Lead SWE, owned the user-data synchronization component: designed the Global Directory schema and a horizontally-scalable CDC pipeline using message queues, circuit breakers, exponential backoff, and a three-tier priority model (org-level, multi-user, single-user). Authored the technical spec that parallelized work across three cross-functional teams, coached a mostly-junior team, and managed a mid-project engineer departure without delays. Successfully synced 90M+ users at ~2,000 req/s per instance. The pipeline became the reference architecture adopted by sibling teams across the company.  
**Key signals:** Large-scale distributed systems design · technical leadership with a junior team · cross-team dependency management at Salesforce scale.

---

## Common Questions → Project Fit

1. **Tell me about yourself / walk me through your background.**  
   → Draw from all projects to build a narrative arc: Microsoft (MSA identity → dSTS leadership → Remix ownership) → Salesforce (large-scale distributed systems + team lead) → AxoEdge (founding, applied ML). No single project owns this answer.

2. **Tell me about a time you faced conflict on a team.**  
   → **Primary:** Win8 Protocol Disagreement (direct interpersonal conflict, walkout, reset)  
   → **Secondary:** dSTS Rapid Team Growth (friction from burned-out team after rapid scaling)

3. **Tell me about a time you made a mistake / failed.**  
   → **Primary:** Win8 Protocol Disagreement (wrong listening posture caused the meeting to fail)  
   → **Secondary:** Xbox Live Site Incident (testing gap that enabled the regression), Salesforce Global IDP (insufficient cross-team design review led to race condition)

4. **Describe a challenging technical problem you solved.**  
   → **Primary:** Xbox Live Site Incident (silent regression invisible to every log and environment)  
   → **Secondary:** MSA WebAuthn (building against a living W3C draft across 4 teams), Remix Search Service (GDPR-compliant rebuild solo in 3 months), Salesforce Global IDP (90M-user CDC pipeline design)

5. **Tell me about a time you had to learn something quickly.**  
   → **Primary:** AxoEdge / ProLogistiks (computer vision and LLM benchmarking from zero in 1 month)  
   → **Secondary:** MSA WebAuthn (learning FIDO2/WebAuthn spec from a working draft)

6. **Describe a time you took ownership of a project.**  
   → **Primary:** Remix Search Service (sole engineer, solo rebuild under GDPR deadline)  
   → **Secondary:** Xbox Live Site Incident (owned root cause, fix, validation, and permanent prevention), Salesforce Global IDP (owned design + delivery of entire user-sync component)

7. **Tell me about a time you influenced without authority.**  
   → **Primary:** MSA WebAuthn (coordinated 4 teams across MSA/Edge/Windows/W3C; shaped spec requirements)  
   → **Secondary:** dSTS Rapid Team Growth (facilitated bottom-up process change across a 30+ person team as tech lead, not manager)

8. **Tell me about prioritizing conflicting tasks or requests.**  
   → **Primary:** Salesforce Global IDP (three-tier sync priority model balancing latency, scale, and reliability)  
   → **Secondary:** Remix Search Service (kept 8-person team unblocked on GDPR work while solo-building the new service), dSTS Rapid Team Growth (diagnosed and ranked pain points across documentation, reliability, and onboarding)  
   > *Consider adding a story with a clearer personal prioritization tradeoff under explicit competing stakeholder demands.*

9. **Tell me about working with ambiguous requirements.**  
   → **Primary:** AxoEdge / ProLogistiks (undefined architecture, unknown domain, no prior baseline)  
   → **Secondary:** MSA WebAuthn (spec was a living draft; requirements changed mid-implementation), Salesforce Global IDP (greenfield system design with no existing data model)

10. **Tell me about disagreeing with a peer or manager.**  
    → **Primary:** Win8 Protocol Disagreement (direct technical disagreement with the protocols architect)  
    → **Secondary:** MSA WebAuthn (advocated for full passwordless vs. second-factor-only over internal skepticism)

11. **Describe improving system performance or reliability.**  
    → **Primary:** Xbox Live Site Incident (diagnosed >80% login drop, restored baseline in ~3 hrs, shipped prevention automation)  
    → **Secondary:** Remix Search Service (replaced slow legacy service; achieved p95 < 100 ms), Salesforce Global IDP (circuit breakers + backoff design for 90M-user pipeline reliability)

12. **Tell me about a cross-functional project you worked on.**  
    → **Primary:** MSA WebAuthn (4 teams: MSA, Edge, Windows, W3C PM)  
    → **Secondary:** Salesforce Global IDP (3 cross-functional teams; tribal knowledge gaps as a key risk)

13. **Describe delivering under tight deadlines.**  
    → **Primary:** Remix Search Service (GDPR hard deadline, solo rebuild in 3 months)  
    → **Secondary:** MSA WebAuthn (hard Windows 10 ship date against a moving spec), Xbox Live Site Incident (live incident, ~3 hr resolution window)

14. **Tell me about mentoring or coaching someone.**  
    → **Partial:** Salesforce Global IDP (coached mostly-junior team through spec reviews, PR reviews, and hands-on guidance; redistributed work when engineer departed)  
    → **Partial:** dSTS Rapid Team Growth (facilitated team ownership of processes, pairing, and PR standards)  
    > *These examples are partial. Consider developing a dedicated story with a specific mentee, explicit goals, and measurable growth outcome.*

15. **Tell me about handling a major setback at work.**  
    → **Primary:** Xbox Live Site Incident (production login failure affecting tens of millions of logins/month)  
    → **Secondary:** Salesforce Global IDP (race condition and CDC group-edition bug caused deployment setbacks mid-project; lost an engineer mid-delivery)  
    > *Both are strong on resilience but lean technical. Consider preparing a version that emphasizes personal/team resilience over debugging mechanics.*

---
*Last updated: April 2026*
