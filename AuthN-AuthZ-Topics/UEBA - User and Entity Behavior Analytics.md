User and Entity Behavior Analytics (UEBA) moves threat detection away from static, rule-based alerts and toward machine learning-driven behavioral profiling. Instead of just asking, "Did this user type the correct password?" or "Does this file contain a virus?", UEBA asks, "Is this action normal for this specific user, their peers, and the devices they use?".

## Core Concepts of UEBA

To understand how UEBA fits into a modern security architecture, it relies on four primary mechanisms:

1. **Broad Data Collection:** UEBA ingests logs from across the IT environment—identity providers, VPNs, SaaS applications, endpoints, and network traffic.
2. **Machine Learning and Baselining:** Rather than relying on human-written rules, UEBA algorithms analyze this data to establish a "normal" behavioral baseline for every user and entity (like a server, service account, or IoT device).
3. **Anomaly Detection:** Once a baseline is established, the system constantly compares real-time activity against it. Deviations—such as logging in at 3 AM when the user usually works 9-to-5, or downloading 50GB of data when their daily average is 50MB—are flagged.
4. **Contextual Risk Scoring:** Instead of firing an alert for every single anomaly (which causes alert fatigue), UEBA stitches discrete events together over time and assigns a cumulative risk score. Only when a score breaches a specific threshold does it trigger a response.


## Main Threat Scenarios

UEBA is specifically designed to catch threats that successfully bypass perimeter defenses and static rules:

- **Compromised Accounts:** External attackers who steal valid credentials via phishing or credential stuffing. UEBA catches them through indicators like "impossible travel" (logging in from two distant countries in an hour) or accessing resources the original user never touches.
- **Malicious Insider Threats:** Employees or contractors intentionally stealing intellectual property or committing fraud. Because they have legitimate access, rule-based systems ignore them, but UEBA catches the sudden spike in data exfiltration or off-hours access.
- **Negligent Users:** Well-meaning employees who accidentally violate policy, such as repeatedly failing authentication or moving sensitive data to unauthorized cloud storage.

## Industry Standards and Frameworks

UEBA maps directly to major cybersecurity and compliance frameworks, translating behavioral anomalies into standardized threat language:

- **MITRE ATT&CK:** UEBA detections map heavily to specific MITRE techniques, such as T1078 (Valid Accounts), T1098 (Account Manipulation), and T1048 (Exfiltration Over Alternative Protocol).
- **NSA Zero Trust Architecture:** Behavioral analytics is explicitly recommended as a core capability under the "Visibility and Analytics" pillar of the NSA's Zero Trust framework.
- **NIST CSF & Compliance:** UEBA satisfies the "Continuous Monitoring" and "Anomalies and Events" requirements of the NIST Cybersecurity Framework, while also providing the audit trails necessary for GDPR, HIPAA, and SOC 2.

## Leading Tools and Market Landscape

Historically a standalone category, UEBA is now most commonly deployed as an integrated layer within Security Information and Event Management (SIEM) or Data Loss Prevention (DLP) platforms to give those tools behavioral context.

Leading tools in the space include:

- **Exabeam:** A heavily ML-focused SIEM that pioneered session stitching and risk scoring.
- **Splunk User Behavior Analytics:** Adds UEBA capabilities on top of Splunk's massive data aggregation engine.
- **IBM Security QRadar SIEM:** Integrates behavioral analytics directly into its core event correlation.
- **Varonis & Microsoft Defender for Identity:** Focus heavily on data-centric security and identity-driven behavioral monitoring.

---

## The bridge between UEBA and ABAC for AuthZ

UEBA is almost never placed in the synchronous "hot path" of authorization.

When you are architecting a Zero Trust system using Attribute-Based Access Control (ABAC), latency is the primary enemy. The policy decision point (PDP) needs to make allow/deny decisions in single-digit milliseconds so it doesn't bottleneck the microservices or user requests relying on it.

UEBA, by its very nature, is computationally heavy. Evaluating a request against a machine learning model, retrieving historical baselines over a 30-day window, and calculating a behavioral risk score takes seconds (or minutes) to process. Putting that in the hot path would bring a distributed architecture to a halt.

Here is how these two systems are architecturally decoupled to maintain performance while still enforcing Zero Trust:

## The Hot Path: Synchronous ABAC

The hot path is strictly for evaluating _available_ state and context instantly.

- **The Actor:** A fast Policy Decision Point (PDP)—such as an **Open Policy Agent (OPA)** engine executing Rego policies.
- **The Process:** When a request hits the API Gateway or Service Mesh, the PDP intercepts it. It evaluates the identity claims (e.g., from a JWT), the resource attributes, and any environment attributes it can pull from a low-latency, in-memory cache (like Redis).
- **Latency Budget:** < 10 milliseconds.

## The Cold Path: Asynchronous UEBA

The cold path happens out-of-band. It ingests logs continuously but does not block real-time traffic.
- **The Actor:** The UEBA engine running inside a SIEM or a dedicated behavioral analytics cluster.
- **The Process:** It asynchronously consumes event logs (from the API gateway, identity provider, network flows). It runs its ML models to detect anomalies and calculates an updated **Risk Score** for the entity.
- **Latency Budget:** Seconds to minutes.


## The Bridge: Eventual Consistency for Risk

To combine the intelligence of UEBA with the speed of ABAC, you bridge the hot and cold paths using eventual consistency:

1. The UEBA engine calculates that an identity's risk score has suddenly spiked from `10` (Low) to `95` (Critical) due to anomalous data exfiltration.
2. The UEBA engine asynchronously pushes this updated `risk_score=95` attribute to a fast-read data store (like a Redis cache) used by the PDP.
3. On the _very next_ API request made by that identity, the PDP (OPA) evaluates the ABAC policy synchronously in the hot path.
4. The Rego policy logic dictates: `deny if user.risk_score > 80`. Because Redis now reflects the updated score, the request is instantly denied.

## Entities that can get a risk score

### 1. Network and Infrastructure Entities (The "Where")

Risk can be assigned to routing or network boundaries rather than individual people or machines.

- **IP Addresses and Subnets:** A specific IP address—or an entire subnet block—can accrue risk if it suddenly initiates port scanning or repeatedly attempts to authenticate against closed services.
- **Geographic Regions (ASN):** An entire country or Autonomous System Number (ASN) can be scored. If a company does no business in a specific region, but begins seeing traffic from there, the region's entity risk score spikes.
- **How ABAC uses it:** In the hot path, an OPA policy might evaluate the environment attribute `request.ip.risk_score`. If that subnet is currently flagged by UEBA as compromised, OPA denies access, even if the user identity is perfectly valid.

## 2. Device and Endpoint Entities (The "What")

A single laptop, mobile device, or IoT sensor has its own behavioral baseline entirely independent of who logs into it.

- **Hardware and MAC Addresses:** If a corporate laptop typically runs developer tools and connects to GitHub, but suddenly begins running PowerShell scripts to scrape memory and beaconing to an unknown external domain, the _device_ risk score spikes.
- **Unmanaged Devices:** A printer or a smart thermostat (which has no traditional user identity) can be assigned a risk score based on its network flow data. If a printer suddenly tries to access the Active Directory server, the UEBA engine flags the device entity.

## 3. Workload and Service Entities (The "Machine")

In a microservices architecture, non-human identities often outnumber human users by orders of magnitude.

- **Service Accounts and API Keys:** A service account used by a CI/CD pipeline might pull code from a repo and push containers to a registry every day at 2 PM. If that same service account tries to delete an S3 bucket at midnight, its risk score triggers an alert.    
- **Kubernetes Pods:** Modern Zero Trust systems use SPIFFE/SPIRE to give cryptographic identities to workloads. UEBA tracks the behavior of these workloads (e.g., "Service A only ever talks to Service B").