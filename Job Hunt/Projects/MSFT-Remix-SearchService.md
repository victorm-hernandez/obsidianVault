# Short Version

On Remix 3D — Microsoft's community platform for sharing 3D models — I was the sole engineer who owned replacing a legacy Zune-era search service to hit the May 2018 GDPR deadline. When the England-based team that owned the service couldn't deliver the required changes in time, I evaluated ramping up on their codebase, decided a ground-up rebuild was faster and more maintainable, and built a new Service Fabric microservice layered over Elasticsearch — complete with a query transformation layer, tokenization, fuzzy search via dynamic Lucene queries, authentication, and full monitoring with alerting. I delivered it from scratch in 3 months as the sole engineer, achieving search p95 under 100 ms and unblocking the team's entire GDPR compliance plan.

---

# Polished STAR++ Narrative

## Situation
I was a Senior Software Engineer and tech lead on the Remix 3D team at Microsoft — a community platform launched in 2017 for sharing and discovering 3D models with Paint 3D and 3D Viewer [est. ~several hundred thousand catalog assets, [est. ~100K monthly active users]]. The core search experience was owned by an England-based team and ran on a repurposed Zune-era service with minimal support and heavily delayed feature requests.

## Task
As the team's technical lead, I owned making our search service GDPR-compliant within a hard 3-month window before the May 2018 enforcement deadline — the only remaining blocker in the team's GDPR plan — while preserving all existing search scenarios and keeping my 8-person team unblocked on their own compliance work.

## Actions

- **Build vs. borrow decision:** I first evaluated ramping up on the England team's Zune codebase. The ramp-up cost was high, the codebase carried significant technical debt from its Zune origins, and the England team could only offer limited guidance. I concluded a targeted rebuild would be faster to deliver and easier to maintain long-term, and socialized that decision with my manager before proceeding.

- **Technology selection:** I evaluated CosmosDB, SQL Server full-text search, and Azure Search before selecting Elasticsearch for its Lucene query flexibility and the ability to support fuzzy/synonym-aware search. I chose Service Fabric as the hosting platform to align with the team's existing microservices infrastructure.

- **Architecture — query transformation layer:** I designed the new service as a thin intermediary: a query transformation layer sat between our API surface and the Elasticsearch provider, allowing me to iterate on search logic (tokenization, synonym aggregation, relevance tuning, dynamic Lucene queries for fuzzy search) without touching the backing store. This abstraction also made GDPR data-handling changes localized to the transformation layer.

- **Delivery and operationalization:** Working solo, I built the service end-to-end — implementing authentication, instrumentation, and automated alerting — and authored deployment pipelines from scratch so the service could be released and monitored without manual intervention. I validated correctness against the existing search scenarios before cutover.

## Result
I delivered the replacement search service from scratch in 3 months as the sole engineer, achieving search p95 under 100 ms and meeting the GDPR deadline. This unblocked the team's entire GDPR compliance plan; without this service, the platform would have needed to be paused to avoid regulatory penalties.

## Reflection & Tradeoffs
The build-vs-borrow call was the highest-stakes decision. In hindsight it was the right one — the legacy codebase had no active maintainers and would have become a long-term liability — but I'd invest more time upfront in load testing at realistic catalog scale and in documenting the query transformation layer for future engineers. One thing I'd do differently is propose the new service as a shared asset earlier, since other teams in the org also depended on search; formalizing it as a platform service rather than a team-local one could have had broader impact before the product's eventual shutdown.

---

# Quick Pass/Fail Checks
- Situation: yes
- Task: yes
- Actions: yes
- Result: yes (p95 metric present; traffic/scale are labeled estimates)
- Reflection: yes

---

# STAR++ Score
Score: 7/10 — Strong technical depth and a clear build-vs-borrow decision, but impact is bounded by missing traffic/scale data; adding even estimated query volume would raise this to 8–9.

---

# Strengths
- Clear and defensible build-vs-borrow tradeoff with concrete reasoning
- Strong technical breadth: architecture decision, query layer design, Elasticsearch/Lucene specifics, and full operationalization (pipelines, alerting)
- Quantified result (p95 < 100 ms) tied directly to the compliance deadline stakes
- Solo ownership is compelling signal for scope and accountability

---

# Weaknesses & Concrete Improvements
- No traffic/scale data → Add even a rough estimate: "The catalog had [est. ~500K models] and search handled [est. ~10K queries/day]" — makes the latency result land harder.
- GDPR change specifics are vague → Name one concrete GDPR requirement you addressed (e.g., right-to-erasure in the index, PII scrubbing in query logs) to show domain depth.
- No mention of testing strategy → Add one sentence: "I validated against a test corpus covering all existing search scenarios before cutover and monitored error rates for 48 hours post-deploy."

---

# One-line Rewrite Suggestion
"I rebuilt the search service — Elasticsearch + Service Fabric, with a query transformation layer for GDPR-safe data handling — solo in 3 months, hitting p95 < 100 ms and clearing the team's last compliance blocker before the May 2018 enforcement date."

---

# Time-to-Deliver Check
Trimable to ~2 minutes? yes — The Short Version above is the trim; cut the technology evaluation list (CosmosDB/SQL Server detail) and the platform-service reflection to stay tight on the core decision and result.

---

# Two Quick Tips to Increase Technical Signal
- **Name the GDPR mechanism:** Specify which GDPR obligation drove the work (e.g., right to erasure requiring async index deletion, or PII scrubbing in search telemetry). One sentence doubles the compliance credibility.
- **Quantify the query transformation layer's value:** Add a concrete example of a change the layer enabled (e.g., "I added synonym expansion for common 3D terminology without a reindex") to show the architectural decision paid off during the project, not just in theory.

---

# Suggested Interviewer Follow-ups
- Q1: "Why Elasticsearch over Azure Cognitive Search, given you were already on Azure?"
- Q2: "How did you handle GDPR's right-to-erasure in a search index without rebuilding the entire index on delete?"
- Q3: "How did you validate that the new service matched the behavior of the legacy service before fully cutting over?"

---

# Metadata
- Role: Senior Software Engineer / Technical Lead, Remix 3D (Microsoft)
- Timeframe: January 2018 – April 2018 (~3 months; GDPR enforcement deadline May 25, 2018)
- Metrics provided: yes (p95 < 100 ms)
- Estimates used: yes — `[est. ~100K monthly active users]`, `[est. ~500K catalog models]`, `[est. ~10K queries/day]`

