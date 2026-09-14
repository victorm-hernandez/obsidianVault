# dSTS Data Layer Migration

---

## Short Version

As a Principal Software Engineer on dSTS — Azure's baremetal, tier-0 identity provider — I led the design of a cross-org data layer migration to address mounting technical debt and repeated livesite incidents. I authored a proposal to replace the legacy distributed-hashtable store with SQL Server + Webstore, and to substitute the legacy codebase with Entra ID's modern login stack, validated via shadow traffic through an existing gateway. In Q1 2023, the team delivered a prototype with equivalent login success rates [est. ~99%+ parity], proving the migration path and unblocking the roadmap toward modern auth protocols (OAuth/OIDC).

---

## Polished STAR++ Narrative

### Situation
dSTS is Microsoft's baremetal, tier-0 regional identity provider for all of Azure — it boots immediately after the networking layer in every new datacenter, has zero service dependencies, and serves as the root trust anchor for [est. hundreds of datacenters globally]. The service had been inherited from a legacy codebase, was running on obsolete storage technology, had near-zero institutional knowledge remaining (1 of 5 original engineers), and was generating repeated livesite incidents driving high on-call attrition.

### Task

As the Principal Software Engineer brought in during a targeted hiring wave, I owned end-to-end design of the data layer migration strategy — including technology selection, architecture, rollout plan, and team task distribution — with the goal of stabilizing the service and enabling roadmap features (OAuth/OIDC migration).

### Actions

- **Landscape assessment:** Mapped the fragmented codebase (legacy + modern control-plane wrapper), cataloged obsolete storage and source control, and assessed the knowledge gap. Concluded a full replacement was lower risk than continued patching.

- **Strategic pivot to Entra ID:** Recognized that Entra ID already had the feature parity dSTS needed (token caching, OAuth/OIDC, healthy livesite) and that its login stack was sufficiently decoupled from its data store to allow a drop-in replacement — the only blocker was its dependency on a non-tier-0 directory service.
- **Data store selection:** Evaluated baremetal-compatible options (open-source alternatives, in-house low-tier stores) and selected SQL Server + Webstore — an in-house sharding layer actively used by Microsoft's consumer identity provider (MSA). Key decision factors: long-term org ownership, pre-existing operational expertise, and ability to secure dedicated DB management resources.
- **Design proposal & rollout strategy:** Authored a detailed proposal covering component breakdown, team task distribution (Webstore deployment, Entra ID data layer, dSTS dual-source data layer, data migration tooling, gateway baremetal updates), and a shadow traffic strategy using an existing gateway to run both stacks in parallel and compare login outcomes before cutover.

### Result
In Q1 2023, the team delivered a working prototype: Entra ID running against the new Webstore-backed data layer loaded with dSTS data, with forked shadow traffic achieving equivalent login success rates [est. ~99%+ parity]. This validated the migration path, de-risked the full cutover, and unblocked the team's modernization roadmap toward OAuth/OIDC protocol support.

### Reflection & Tradeoffs
The biggest tradeoff was choosing Webstore — an in-house legacy sharding technology — over a modern open-source distributed store. Org ownership and pre-existing operational expertise outweighed technical modernity, especially given dSTS had no DB management skillset of its own. In hindsight, I would have established quantitative livesite baselines (incident rate, auth error rate) earlier to measure impact more rigorously post-migration. The key learning: in high-stakes baremetal infrastructure, organizational support and operational continuity often matter more than technical elegance.

---

## Quick Pass/Fail Checks
- Situation: yes
- Task: yes
- Actions: yes
- Result: yes
- Reflection: yes

---

## STAR++ Score
Score: 7/10 — Strong technical depth and explicit ownership; held back by absence of hard metrics and a prototype (not production) result.

---

## Strengths
- Clear strategic thinking: identified the Entra ID pivot rather than patching legacy code
- Explicit tradeoff reasoning for technology selection (Webstore vs. open-source alternatives)
- Owned the full design end-to-end: architecture, rollout, cross-team coordination
- Risk mitigation via shadow traffic validation before committing to cutover

---

## Weaknesses & Concrete Improvements
- No hard metrics → Add: livesite incident frequency before/after, auth error rate baseline, shadow traffic parity percentage, and delivery date vs. plan
- Result is a prototype, not production → Frame explicitly: "this prototype de-risked the migration and enabled the next phase" to show clear business value at the milestone
- Scale context is vague → Add approximate request volume (e.g., [est. ~X million auth requests/day across Azure]) for stronger impact framing

---

## One-line Rewrite Suggestion
"I designed a cross-org data layer migration for Azure's tier-0 identity provider — replacing an unsupported distributed-hashtable store with SQL Server + Webstore and validating equivalence via shadow traffic — delivering a prototype with [est. ~99%+ login parity] in Q1 2023 and unblocking the team's OAuth/OIDC roadmap."

---

## Time-to-Deliver Check
Trimable to ~2 minutes? yes — Remove Webstore historical background; compress the landscape assessment to one sentence; skip the knowledge-gap detail and lead directly with the strategic Entra ID pivot.

---

## Two Quick Tips to Increase Technical Signal
- **Quantify the shadow traffic experiment:** mention the percentage of traffic shadowed, duration, and how you defined "equivalent" (e.g., HTTP status codes, token signatures, error rates) — this shows production-grade validation thinking.
- **Emphasize the dual-source data layer:** calling it out explicitly demonstrates you designed for safe, incremental migration rather than a big-bang cutover — a strong signal for staff+ level interviews.

---

## Suggested Interviewer Follow-ups
- Q1: "Why SQL Server + Webstore over a modern distributed store like Cassandra or CockroachDB?"
- Q2: "How did you validate that Entra ID's data layer was decoupled enough for a drop-in replacement?"
- Q3: "What was the biggest technical risk in the shadow traffic strategy, and how did you mitigate it?"

---

## Metadata
- Role: Principal Software Engineer
- Timeframe: Q1 2023
- Metrics provided: no
- Estimates used: yes ([est. hundreds of datacenters globally], [est. ~99%+ login parity], [est. ~X million auth requests/day across Azure])