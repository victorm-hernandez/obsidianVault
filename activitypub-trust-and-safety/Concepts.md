# ActivityPub Trust & Safety Taskforce: Core Concepts Guide

Welcome to the **ActivityPub Trust and Safety Taskforce** concepts guide. This document provides a comprehensive foundation for understanding the Fediverse, ActivityPub, and the specific trust, safety, and moderation concepts necessary to lead and contribute to this taskforce effectively.

---

## 🌐 1. What is the Fediverse?

### High-Level Overview
The **Fediverse** (a portmanteau of "federation" and "universe") is an interconnected, decentralized web of independently operated social media servers. Unlike centralized platforms (such as X/Twitter, Facebook, or Instagram) controlled by a single corporation, the Fediverse relies on open protocols to allow user accounts on different servers to communicate seamlessly with one another.

```
       +-----------------------+              +-----------------------+
       |   Server A (Mastodon) |              |   Server B (Pixelfed) |
       |   @alice@serverA.org  |  <=========> |   @bob@serverB.com    |
       +-----------------------+  ActivityPub +-----------------------+
                   ^                                      ^
                   |                                      |
                   v                                      v
       +-----------------------+              +-----------------------+
       |    Server C (Lemmy)   |              |  Server D (PeerTube)  |
       |  !forum@serverC.net   |  <=========> |   @video@serverD.tv   |
       +-----------------------+              +-----------------------+
```

### Core Fundamentals & Architecture
* **Decentralized Nodes (Instances)**: Each server in the Fediverse is known as an *instance*. Each instance is owned, hosted, and moderated independently by an individual or organization with its own terms of service, server rules, and moderation policies.
* **Federation**: The process by which independent instances exchange messages, posts, media, follow requests, and moderation signals across the internet using a shared protocol.
* **Global Interoperability**: Users on a microblogging server (e.g., Mastodon) can follow, comment on, or like posts from users on an image-sharing server (e.g., Pixelfed), a video platform (e.g., PeerTube), or a link aggregator (e.g., Lemmy).
* **Decentralized User Identity**: User identities are formatted similarly to email addresses: `@username@domain.com`. User discovery across servers is facilitated via **WebFinger** (RFC 7033).
* **Timelines**:
  * *Home Timeline*: Content from actors the user explicitly follows.
  * *Local Timeline*: Content published by all users on the local server instance.
  * *Federated Timeline*: Public content known to the local instance (received via federation from remote servers).

---

## 📜 2. What is ActivityPub?

### High-Level Overview
**ActivityPub** is an open, W3C Recommendation standard (published in 2018) for decentralized social networking. It defines a protocol for creating, updating, and deleting content, as well as managing social relationships (following, liking, blocking, reporting).

ActivityPub utilizes **Activity Streams 2.0** as its data serialization format, structured in **JSON-LD** (JSON for Linked Data).

### Core Mechanics & Concepts

#### 1. The Actor Model
In ActivityPub, every entity that can perform an action is an **Actor**. Actors can represent individual users, organizations, groups, channels, or automated bots. An Actor object contains:
* Unique URI identifier (`id`).
* Inbox URI (where incoming activities are delivered).
* Outbox URI (where outgoing activities are published).
* Public key material (for cryptographic signatures).
* Collections (e.g., `followers`, `following`, `liked`).

#### 2. Protocol Architecture: C2S vs. S2S
ActivityPub defines two distinct protocols:
* **Server-to-Server (S2S) Federation Protocol**: Used by instances to federate activities with each other. This is the primary domain of the Trust & Safety Taskforce.
* **Client-to-Server (C2S) Protocol**: Used by client applications (mobile apps, web apps) to communicate with a user's home server.

```
+---------------+     C2S      +------------------+     S2S      +------------------+
| Client App    | ===========> | Local Server     | ===========> | Remote Server    |
| (Mobile/Web)  | <=========== | (Outbox / Inbox) | <=========== | (Inbox / Outbox) |
+---------------+              +------------------+              +------------------+
```

#### 3. Inbox and Outbox Pattern
* **Outbox**: A server endpoint where an Actor posts new activities. When a user posts a message, their client posts a `Create` activity to their Outbox. The server then distributes this activity to the inboxes of all remote followers.
* **Inbox**: A server endpoint that receives incoming activities sent from remote instances. The receiving server parses the activity and adds it to the user's stream.
* **Shared Inbox**: A performance optimization endpoint on an instance that receives messages intended for multiple local users simultaneously.

#### 4. Activity Streams 2.0 Objects and Activities
Everything transmitted in ActivityPub is an **Activity** wrapped around an **Object**:
* **Objects**: `Note` (short posts), `Article`, `Image`, `Video`, `Document`, `Event`, `Person`, `Group`.
* **Activities**: `Create`, `Update`, `Delete`, `Follow`, `Accept`, `Reject`, `Add`, `Remove`, `Like`, `Announce` (reblog/retweet), `Undo`, `Flag` (report), `Block`.

*Example: A post creation in JSON-LD / Activity Streams 2.0:*
```json
{
  "@context": "https://www.w3.org/ns/activitystreams",
  "id": "https://example.com/activities/123",
  "type": "Create",
  "actor": "https://example.com/users/alice",
  "to": ["https://www.w3.org/ns/activitystreams#Public"],
  "object": {
    "id": "https://example.com/notes/456",
    "type": "Note",
    "attributedTo": "https://example.com/users/alice",
    "content": "Hello Fediverse!"
  }
}
```

---

## 🛡️ 3. Essential Concepts for Leading the Trust & Safety Taskforce

To effectively guide and lead the SWICG ActivityPub Trust and Safety Taskforce, you must be familiar with the specialized moderation, safety, governance, and protocol extensions unique to decentralized systems.

### A. Moderation Activities & Federated Signal Routing

#### The `Flag` Activity (Abuse Reports)
* **What it is**: The standard ActivityPub activity used when a user or moderator reports abusive, spammy, or illegal content.
* **Key Challenge**: The core ActivityPub spec does not mandate how `Flag` activities should be addressed or processed by remote servers. 
* **Taskforce Focus**:
  * Establishing standard addressing patterns (`to`, `cc`, `bto`, `bcc`) to determine whether a report is routed to the reported user's server admins, the local server admins, or third-party moderation services.
  * Standardizing `Flag` payload metadata: categories (e.g., spam, harassment, CSAM), reporter notes, and attached evidence snippets.

#### `Block` vs. `Undo(Block)` vs. Server Defederation
* **User-Level `Block`**: An Actor sends a `Block` activity targeting another Actor. The receiving server should suppress content delivery between these actors.
* **Server-Level Moderation (Defederation / Domain Block)**: Server administrators block entire remote domains due to systemic abuse or non-responsiveness to reports.
  * *Silence / Limit*: Hides the remote instance's posts from public timelines while allowing individual follows to persist.
  * *Full Suspend / Defederation*: Rejects all incoming and outgoing federation traffic with the remote domain.

---

### B. Federated Safety Harms & Vulnerabilities

| Harm / Vulnerability | Description in ActivityPub Context | Taskforce / Technical Response |
| :--- | :--- | :--- |
| **Cross-Server Harassment** | Malicious users on Instance X targeting users on Instance Y. | Interoperable `Flag` activities, federated blocklists, and shared abuse metadata. |
| **Spam & Botnets** | Automated accounts flooding federated inboxes via `Create` activities. | Challenge-response mechanics, rate-limiting headers, domain reputation systems, HTTP signature checks. |
| **Illegal Content (e.g., CSAM)** | Distribution of illegal imagery via federated media attachments. | Perceptual hashing (PDQ / PhotoDNA) exchange, automated content scanning standards, rapid deletion propagation via `Delete` activities. |
| **Reply Hijacking & Impersonation** | Spoofing actor metadata or altering comment threads hosted on remote servers. | HTTP Signatures, Object Integrity Proofs (LD-Signatures / Data Integrity), Authorized Fetch. |
| **Data Scraping & Consent** | Unconsented bulk harvesting of user data for AI training or surveillance. | Content-use annotations, robots.txt equivalents for ActivityPub objects, access control headers. |

---

### C. Content Warnings (CW) and Safety Labeling

* **`summary` Property**: In Activity Streams 2.0, the `summary` field on a `Note` or `Article` acts as a Content Warning (CW), requiring client applications to collapse the body (`content`) behind a spoiler toggle.
* **`sensitive` Flag**: A boolean flag indicating that attached media contains NSFW or trigger content, requiring media blurring until clicked.
* **Automated Safety Labels & Annotations**: Standardizing schema extensions so third-party moderation tools or automated classifiers can attach safety trust scores, safety tags, or factual labels to activities without mutating original content.

---

### D. Security & Privacy Controls

* **HTTP Signatures**: A cryptographic standard where HTTP requests between ActivityPub servers are signed using the actor's private key (`Signature` HTTP header). Verifies the authenticity and origin of activities.
* **Authorized Fetch (Secure Mode)**: Requiring HTTP Signatures even for `GET` requests fetching public objects, preventing blocked servers or unauthorized scrapers from fetching content.
* **Object Integrity Proofs**: Embedded cryptographic signatures attached directly to JSON-LD objects, ensuring content cannot be tampered with in transit or when relayed through third parties.

---

### E. W3C SWICG Governance & Standards Process

As Taskforce Lead, navigating the W3C community process is vital:

```
+-------------------------------------------------------------------+
|               W3C Social Web Community Group (SWICG)             |
+-------------------------------------------------------------------+
                                  |
         +------------------------+------------------------+
         v                                                 v
+-------------------------------+               +-------------------------------+
|  Trust & Safety Task Force    |               |  Federated Extension          |
|  (Reports, Notes, Guidelines) |               |  Proposals (FEPs)             |
+-------------------------------+               +-------------------------------+
```

* **W3C Social Web Incubator Community Group (SWICG)**: The parent body for ActivityPub maintenance and community proposals.
* **Community Group (CG) Reports & Notes**: Non-normative, community-consensus documents produced by task forces to provide guidelines, best practices, and recommended extensions. (Distinct from W3C Recommendation Track standards).
* **Fediverse Enhancement Proposals (FEPs)**: A merit-based, community-driven process for proposing specific technical extensions to ActivityPub. FEPs are often created or reviewed within SWICG task forces before wide adoption.
* **Consensus Building**: The core role of taskforce leadership is facilitating discussion across diverse software maintainers (microblogs, forums, media sites) to reach rough consensus on safety specifications.

---

## 📚 4. Quick-Reference Glossary

* **Actor**: The primary entity in ActivityPub representing a user, bot, or group with an Inbox and Outbox.
* **Activity Streams 2.0 (AS2)**: The JSON-LD data vocabulary used by ActivityPub to represent social objects and activities.
* **Authorized Fetch**: A server security mode requiring cryptographic HTTP signatures for `GET` requests to prevent blocked instances from viewing public data.
* **C2S / S2S**: Client-to-Server / Server-to-Server protocol distinction in ActivityPub.
* **Defederation (FediBlock)**: Severing federation ties between two instances at the administrator level.
* **FEP (Fediverse Enhancement Proposal)**: Community proposal process for expanding ActivityPub features.
* **`Flag`**: The ActivityPub activity type used for submitting abuse reports.
* **IFTAS (Independent Federated Trust & Safety)**: A non-profit organization supporting moderation safety and standards across the Fediverse.
* **Instance**: An independently hosted server operating in the Fediverse.
* **JSON-LD**: JSON for Linked Data, allowing schema extensibility via `@context` URIs.
* **WebFinger**: An HTTP protocol (RFC 7033) used to discover an actor's ActivityPub profile URI from an address like `@user@domain.com`.
