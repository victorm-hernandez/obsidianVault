**Common practice:** many systems combine them, using ==RBAC for coarse roles, ReBAC for resource sharing, and ABAC for contextual conditions==.

RBAC for baseline entitlements, ReBAC for sharing and hierarchies, and ABAC conditions for guardrails like device posture or data residency.

### RBAC: Role-Based Access Control

Access is granted to **roles**, and users get access by being assigned roles. Permissions never attach directly to individual users.

**Model:** `User → Role → Permission`

- Roles: `Admin`, `Editor`, `Viewer`, `Billing Manager`
- Permissions: `document:read`, `document:write`, `user:delete`, `invoice:approve`
- Alice has role `Editor`, and `Editor` includes `document:read` and `document:write`, so Alice can read and write documents.

**Common extensions:**
- **Hierarchical roles:** `Admin` inherits everything `Editor` has, which inherits from `Viewer`.
- **Constraints:** separation of duty (nobody can hold both `Requester` and `Approver`).
- **Scoped roles:** a role assigned within a tenant, project, or org (`Editor` on Project X only).

**Strengths**
- Simple to understand, implement, and audit ("who has Admin?")
- Maps well to job functions and org charts
- Easy to certify in access reviews, which is why compliance frameworks like SOC 2 lean on it

**Weaknesses**
- **Role explosion:** as requirements get finer ("Editor of Project X but only in EU, only for drafts"), you mint more and more roles.
- **No context:** it can't natively say "only during business hours" or "only from managed devices."
- **Coarse per-resource control:** "this one document" requires a role per resource or a bolt-on.

---
### ABAC: Attribute-Based Access Control

Access is decided by evaluating a **policy** against **attributes** of the request. No fixed role is required.

The four attribute categories (from NIST SP 800-162):

- **Subject:** department, clearance level, job title, employment type
- **Resource:** classification, owner, data type, region
- **Action:** read, write, delete, approve
- **Environment:** time of day, IP or network, device posture, risk score

**Example policies:**

- "Allow `read` if `user.department == resource.department`"
- "Allow `view` on documents where `resource.classification <= user.clearance`"
- "Deny `export` if `environment.device_managed == false`"
- "Allow `approve` if `resource.amount < 10000` and `user.title == 'Manager'`"

**Typical architecture:**

- **PEP** (Policy Enforcement Point) intercepts the request
- **PDP** (Policy Decision Point) evaluates the policy and returns permit/deny
- **PIP** (Policy Information Point) fetches the needed attributes
- **PAP** (Policy Administration Point) is where policies are authored

Policy languages include XACML (the original standard), AWS IAM conditions, Cedar (AWS), and OPA/Rego.

**Strengths**

- Very expressive: fine-grained, contextual, dynamic decisions
- No role explosion. Adding a new resource with the right attributes just works.
- Policies change centrally without reassigning users

**Weaknesses**

- **Hard to audit:** "who can access X?" requires evaluating policies against all possible attribute combinations, not reading a list.
- **Policy sprawl and debugging:** interacting allow/deny rules get hard to reason about.
- **Attribute dependency:** decisions are only as good as attribute quality and freshness, and fetching them adds latency.
- **Heavier to build and operate** than RBAC.

### Side by side

|                    | RBAC                                  | ABAC                                              |
| ------------------ | ------------------------------------- | ------------------------------------------------- |
| **Decision input** | User's assigned roles                 | Attributes of user, resource, action, environment |
| **Granularity**    | Coarse                                | Fine                                              |
| **Context-aware**  | No (without extensions)               | Yes                                               |
| **Auditability**   | Easy ("who has role X?")              | Harder ("what does this policy allow?")           |
| **Complexity**     | Low                                   | High                                              |
| **Best for**       | Job-function permissions, admin tiers | Data classification, regulatory, contextual rules |

---

## ReBAC

REBAC is **relationship-based access control**. Instead of deciding access from a user's role (RBAC) or from attributes of the user, resource, and environment (ABAC), it decides access from the **relationships between entities** in a graph.
### The core idea
You store facts as relationship tuples, typically `object # relation @ subject`:

- `doc:roadmap # owner @ user:alice`
- `folder:planning # viewer @ team:eng#member`
- `doc:roadmap # parent @ folder:planning`
- `team:eng # member @ user:bob`

A permission check like "can Bob view `doc:roadmap`?" becomes a **graph traversal**: Bob is a member of `team:eng`, which is a viewer on `folder:planning`, which is the parent of `doc:roadmap`, and the schema says viewers of a parent are viewers of the child. So yes.

Permission rules are defined in a schema that says how relations compose, for example "a document's viewers = its direct viewers + owners + viewers of its parent folder."

### How it differs from RBAC and ABAC

|                | RBAC                                         | ABAC                                         | ReBAC                                            |
| -------------- | -------------------------------------------- | -------------------------------------------- | ------------------------------------------------ |
| **Decides by** | Role assigned to user                        | Attributes + policy rules                    | Path through a relationship graph                |
| **Good at**    | Coarse, org-wide permissions                 | Contextual rules (time, location, clearance) | Per-resource sharing, hierarchies, nested groups |
| **Pain point** | Role explosion; can't express "this one doc" | Policies get hard to reason about and audit  | Needs a graph store and careful schema design    |
RBAC says "Bob is an editor." ReBAC says "Bob is an editor _of this document_ because he's on a team that owns its parent folder." That per-resource granularity is what RBAC handles poorly without inventing a role for every resource.


---

## Architectural Elements of Access Controls

- **PEP (Enforcement):** intercepts the request and enforces the decision. It exists in _every_ system that checks permissions: an API gateway, middleware, or a check in your service code.
- **PDP (Decision):** evaluates "is this allowed?" and returns permit/deny.
- **PIP (Information):** supplies the data the PDP needs to decide.
- **PAP (Administration):** where policies or permission data are authored and managed.

They're commonly associated with ABAC because XACML popularized the vocabulary, but they describe **authorization architecture**, not a specific access model.

### How they map to each model

|         | RBAC                                                     | ABAC                                                                         | ReBAC                                                                |
| ------- | -------------------------------------------------------- | ---------------------------------------------------------------------------- | -------------------------------------------------------------------- |
| **PEP** | Middleware checking the user's role or permission        | Same                                                                         | Same (calls the check API)                                           |
| **PDP** | Role-to-permission lookup (often trivial, or in-process) | Policy engine (OPA, Cedar, XACML engine)                                     | Graph evaluator (SpiceDB, OpenFGA, Zanzibar `Check`)                 |
| **PIP** | User-to-role assignments (IdP, directory, JWT claims)    | Attribute sources: HR system, device posture, resource metadata, risk engine | The relationship tuple store itself                                  |
| **PAP** | Admin UI for roles and assignments                       | Policy authoring tools, policy repo                                          | Schema definition plus tuple writes (share a doc, add a team member) |
### Where the fit differs

- **RBAC:** the components exist but are usually collapsed. Roles often ride in the token (a JWT `roles` claim), so the ==PEP makes the decision locally. The PDP and PIP become nearly invisible, which is why nobody talks about them for RBAC.==
- **ABAC:** the separation matters most here. The PIP is a real problem: ==attributes come from many systems, with latency and freshness concerns==. That's why the four-part decomposition got emphasized.
- **ReBAC:** the PDP and PIP blur together. In Zanzibar-style systems, the relationship store _is_ the information source, and the decision is a graph traversal over it. Your app is the PEP, calling `Check(user, relation, object)`. The schema plus tuple writes play the PAP role, and the tuple-write path is often driven by your application events (a user shares a doc, so a tuple is written).


### The ABAC PIP problem


The PIP is a problem because the policy is simple to write, but the attributes it needs are scattered across systems that were never designed to answer authorization queries in real time. Here are some examples where each failure mode shows up.

### Example 1: Hospital record access

> Allow `read` on a patient record if `doctor.department == patient.assigned_department` AND `doctor.on_shift == true` AND `patient.consent_restrictions` doesn't exclude the doctor's role.

One policy, but the attributes come from four systems:

|Attribute|Source|
|---|---|
|`doctor.department`|HR system|
|`doctor.on_shift`|Scheduling system|
|`patient.assigned_department`|EHR|
|`patient.consent_restrictions`|Consent registry|

**PIP problems this exposes:**

- **Latency stacking:** one authorization check fans out to four network calls, and the slowest one sets your p99.
- **Partial failure:** if the scheduling system is down, do you deny (a doctor can't treat a patient) or allow (you may violate policy)? You have to decide fail-open vs. fail-closed per attribute, and the PDP usually can't tell you which is safe.
- **Ownership:** each source has its own team, SLA, and schema, and none of them think of themselves as part of your security path.

### Example 2: Device posture

> Deny `export` unless `device.managed == true` AND `device.disk_encrypted == true` AND `device.risk_score < 70`.

**PIP problems:**

- **Freshness vs. cost:** the MDM or EDR state changes minute to minute. Query live and you add latency and load. Cache for 15 minutes and a compromised device keeps access for up to 15 minutes.
- **Revocation lag:** this is the ABAC version of the ReBAC "new enemy problem." The attribute changed, but the decision was made on the old value.
- **Attribute in the token:** if you embed device posture in a JWT at login, it's stale for the token's whole lifetime.

### Example 3: Expense approval with delegation

> Allow `approve` if `resource.amount <= user.approval_limit` AND `user` is in the management chain of `resource.submitter`.

**PIP problems:**

- **Derived attributes:** "management chain" isn't stored anywhere. It's computed by walking the org hierarchy, which makes the PIP a query engine, not a lookup.
- **Org changes lag:** a reorg or a manager departure takes days to propagate through HR, so the policy is evaluated against a stale org chart. Someone can approve for a former report.
- **Temporary state:** delegation ("Sam approves while I'm out") is easy to model in policy but has to live in some system as an attribute, with expiry.

### Example 4: The "list what I can see" problem

> Allow `read` if `resource.classification <= user.clearance` AND `resource.region == user.region`.

Checking one document is fine. But the UI asks "show me all documents I can read," and now the PDP needs attributes for **every candidate resource**.

**PIP problems:**

- **N+1 attribute fetches:** evaluating 10,000 documents means 10,000 attribute lookups, or you filter after fetching, which breaks pagination and leaks counts.
- **Can't push down to the database:** a relational store can filter on `classification` and `region` natively, but the PDP doesn't know that unless you do **partial evaluation** (compile the policy into a query filter). That's hard, and it's one place ReBAC systems with `LookupResources` have a more natural answer.

### Example 5: Self-asserted attributes

> Allow `view` if `user.department == 'Finance'`.

Where did `department` come from? If it's a profile field users or a team admin can edit, or a claim from an external IdP you don't control, you've built a privilege escalation path.

**PIP problems:**

- **Provenance and trust:** an attribute is only as trustworthy as its source and who can write to it. Every attribute used in a policy inherits the security requirements of whatever system stores it.
- **Attribute injection:** if the PEP passes attributes from the request (headers, body, unsigned claims), the caller can lie.

### Example 6: Missing or inconsistent attributes

> Allow `access` if `user.employment_type == 'FTE'` AND `resource.data_class != 'restricted'`.

**PIP problems:**

- **Null semantics:** a contractor record has no `employment_type`, or a new resource has no `data_class` yet. Does the policy treat it as false, error, or match? Different engines behave differently, and this is a classic source of accidental allows.
- **Naming drift:** HR says `Full-Time`, the directory says `FTE`, the IdP says `employee`. The policy silently fails to match.

### The pattern

|PIP problem|Root cause|Typical mitigation|
|---|---|---|
|Latency|Attributes live in remote systems|Cache, co-locate, or push data into the PDP (OPA bundles)|
|Staleness|Cache TTL vs. real-time changes|Short TTLs for high-risk attributes, event-driven invalidation|
|Availability|PDP depends on N upstream systems|Explicit fail-closed defaults, degrade gracefully per attribute|
|Consistency|Many sources of truth|Designate an authoritative source per attribute|
|Trust|Attribute writers ≠ policy authors|Only use attributes from sources with security-grade write controls|
|Enumeration|Need attributes for all candidates|Partial evaluation, denormalize into the resource store|

### Why RBAC and ReBAC dodge most of this

RBAC needs one lookup (the user's roles) from one source. ReBAC keeps the relationship data inside the authorization system itself, so the PDP owns its inputs. ABAC's flexibility comes from letting the policy reach into arbitrary external data, and the PIP is where you pay for that.

If it helps, I can walk through how a specific engine handles this. For instance, OPA pushes data to the PDP, while Cedar requires the caller to pass attributes in with the request, which are very different answers to the same problem.