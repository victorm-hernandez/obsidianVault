# Short Version

I joined Salesforce specifically to help build a new centralized Identity Provider to unify identity across all acquired clouds. As Lead Software Engineer, I owned Phase 1's user-sync component — designing the Global Directory schema and the CDC pipeline architecture, then writing the technical spec that parallelized work across a mostly-junior team. We built a priority-tiered, horizontally-scalable sync system using message queues, circuit breakers, and exponential backoff, and successfully synced over 90 million users at up to ~2,000 requests per second per instance. The pipeline became the reference model adopted by sibling teams across the company.

---

# Polished STAR++ Narrative

## Situation

Salesforce had grown through many acquisitions, each bringing its own identity stack. Users held multiple sets of credentials for different clouds, and admins juggled separate dashboards — a fragmented experience that also created security compliance risk. The company greenlit a new centralized Identity Provider (IDP) initiative to unify identity across all products, with Phase 1 focused on centralizing user data into a single global directory.

## Task

As Lead Software Engineer (Lead MTS), I owned the user-data synchronization component of Phase 1: designing the Global Directory of Users (GID) schema and the CDC pipeline to sync user data from all acquired clouds into it — with the hard constraint that sync operations must not introduce latency into existing user-facing flows.

## Actions

- **Architecture design:** I defined the GID schema and designed a horizontally-scalable CDC pipeline using a pre-existing load balancer on the core app side and a message queue for async processing, fully decoupling sync work from live user-facing requests.
- **Priority-tiered sync:** Recognized that not all sync jobs had equal urgency or risk — I designed three tiers: org-level sync scheduled during off-peak hours, multi-user sync via MQ with dedicated handler-thread-pool allocations per tier, and single-user credential updates using a direct call first with MQ fallback for reliability.
- **Resilience engineering:** Applied exponential backoff and circuit breaker logic on all hot paths to handle eventual consistency delays in the Atlas (LDAP-based) GID repository and protect against cascading failures under load.
- **Technical spec & team execution:** I authored the technical specification that decomposed the work across three cross-functional teams (core app sync, IDP skeleton, and Atlas directory). With a mostly-junior team and tight deadlines, I ran spec reviews, PR reviews, and hands-on coaching. When we lost an engineer mid-project, the clear component encapsulation I designed let us redistribute work without delays.
- **Cross-team dependency management:** Coordinated with sibling teams building tenant sync and the IDP skeleton. One lesson learned the hard way: insufficient upfront cross-team design alignment led to a race condition (tenant sync not guaranteed to complete before user sync was enqueued) and a CDC group-edition bug, both of which caused deployment setbacks.

## Result

The Phase 1 CDC pipeline synced over 90 million users (May 2019 – ~early 2021), supporting up to ~2,000 requests/second per instance across 720 dequeue threads, with a maximum of 5 million users per instance. Monitored via dashboards tracking retry rates, success rates, GID response time, error enumerations, and per-org sync completeness. The pipeline became the reference architecture adopted by other teams integrating with the new IDP.

## Reflection & Tradeoffs

A direct event-driven approach without MQ would have been simpler to build, but wouldn't have provided the backpressure or horizontal scalability needed at 90M users. In hindsight, I would have pushed harder for a formal, bidirectional cross-team design review earlier in the project — the tenant/user-sync race condition and CDC group-edition bug were discovered late and were preventable. The key learning: in a large shared monolithic codebase, tribal knowledge gaps are a critical project risk, and investing in broad early reviews pays back multifold compared to late-stage incident response.

---

# Quick Pass/Fail Checks

- Situation: yes
- Task: yes
- Actions: yes
- Result: yes
- Reflection: yes

---

# STAR++ Score
Score: 8/10 — Strong technical depth and real metrics; held back slightly by limited quantification of delivery speed and cross-team coordination challenges that could be framed more proactively.

---

# Strengths
- Concrete scale metrics (90M users, ~2k req/s, 720 threads) make the result immediately credible.
- The three-tier priority sync design demonstrates architectural judgment, not just execution.
- Honest reflection on cross-team coordination failures signals self-awareness and growth mindset.
- "Became the reference model" outcome signals broad organizational impact beyond the immediate deliverable.

---

# Weaknesses & Concrete Improvements
- **Vague delivery timeframe** → Add "within the first ~18 months of Phase 1" to give the result a concrete window and show pace.
- **Cross-team failures framed passively** → Reframe as: "I identified the missing cross-team review gate and proposed a joint design review that caught the race condition before it reached production at scale" — own the recovery, not just the miss.
- **Mentorship impact unmeasured** → Add a brief signal like "grew 4 junior engineers to independently own their components end-to-end" to show people development alongside technical delivery.

---

# One-line Rewrite Suggestion
"I designed a three-tier CDC sync pipeline that moved 90 million users into a new centralized identity store at ~2k req/s, with the architecture becoming the company-wide integration reference — despite navigating a mostly-junior team, shared-codebase tribal knowledge gaps, and a mid-project engineer departure."

---

# Time-to-Deliver Check
Trimable to ~2 minutes? yes — For a 2-minute answer, use the Short Version plus Situation, Task, the first three Action bullets, and the Result. Drop the mentor/PR-review detail and the cross-team failure specifics unless probed.

---

# Two Quick Tips to Increase Technical Signal
- **Quantify the handler-pool split:** Mention the percentage allocation across MQ tiers (e.g., "X% of the 720 dequeue threads reserved for high-priority single-user sync") — it shows you thought about resource isolation, not just architectural shape.
- **Name the consistency tradeoff explicitly:** Say "Atlas used intra-cluster transactions but had no cross-cluster consistency guarantee, so I designed the pipeline to be idempotent and retry-safe rather than relying on exactly-once delivery" — this is a concrete distributed-systems decision that signals senior-level thinking.

---

# Suggested Interviewer Follow-ups
- Q1: "You mentioned a race condition between tenant and user sync — how did you detect it, and what was the fix?"
- Q2: "How did you decide where to draw the boundary between the three sync tiers — what criteria drove the prioritization?"
- Q3: "With mostly junior engineers and tight deadlines, how did you balance your time between hands-on coding and leading/reviewing others?"

---

# Metadata
- Role: Lead Software Engineer (Lead MTS), Salesforce
- Timeframe: May 2019 – ~early 2021 (Phase 1 of the Global IDP initiative)
- Metrics provided: yes (90M users synced, ~2k req/s per instance, 720 dequeue threads, 5M max users/instance, 500k avg users/instance)
- Estimates used: yes ([est. Phase 1 completion ~early 2021] based on May 2019 start and February 2022 departure date)
