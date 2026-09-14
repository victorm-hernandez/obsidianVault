# ActivityPub Trust and Safety Taskforce Summary

The **ActivityPub Trust and Safety Taskforce** is an initiative operating under the **W3C Social Web Incubator Community Group (SWICG)**. Its primary mission is to identify, document, and establish standards and best practices for trust, safety, user privacy, and moderation within decentralized social networks built on the [ActivityPub](https://www.w3.org/TR/activitypub/) protocol and [Activity Streams 2.0](https://www.w3.org/TR/activitystreams-core/) specifications.

---

## 🎯 Purpose and Objectives

While the foundational ActivityPub and Activity Streams specifications define core mechanisms for federated social networking, they do not comprehensively specify moderation semantics, anti-spam protections, or safety mechanisms required for operating large-scale online communities.

The ActivityPub Trust and Safety Taskforce addresses these gaps by:
- **Identifying Vulnerabilities & Harms**: Analyzing spam, abuse, harassment, and content integrity issues in federated networks (the "Fediverse").
- **Improving Moderation Interoperability**: Ensuring moderation reports, blocks, flags, and account enforcement actions can be cleanly communicated across heterogeneous software platforms (e.g., Mastodon, Pixelfed, PeerTube, Lemmy).
- **Developing Best Practices**: Publishing reports, community notes, and protocol recommendations for developers building ActivityPub-enabled applications.

---

## 🛠️ Workstreams & Key Initiatives

The taskforce's scope of work is divided into three primary workstreams:

1. **Initial Report & Best Practices Overview**
   - Surveying the current state of trust and safety across the Fediverse.
   - Documenting existing ActivityPub features and implementation patterns.
   - Establishing baseline recommendations for developers to maintain safety standards.

2. **Improving Moderation Activities (`Flag`, `Block`, etc.)**
   - Enhancing federated moderation workflows by refining addressing mechanisms for where flags/reports are sent.
   - Standardizing additional metadata fields in moderation activities to prevent dropped reports between differing server software implementations.

3. **Content Labeling, Warnings, and Annotations**
   - Standardizing content warnings (CWs), media sensitive flags, content labeling, and automated safety annotations across federated systems.

---

## 📌 Scope & Governance

- **Governance**: Managed by the W3C Social Web Incubator Community Group (SWICG). Taskforce leadership includes Emelia Smith ([@thisismissem](https://github.com/thisismissem)).
- **Deliverables**: The taskforce produces **Community Group Reports and Notes** (not official W3C standards-track recommendations, but community-supported specifications).
- **Repository Role**: Serves as a centralized issue tracker for trust and safety proposals and as the editing environment for taskforce reports.
- **Software Implementation Boundary**: The taskforce does not directly modify application codebases (e.g., Mastodon, PeerTube). Instead, it establishes protocol-level guidelines and schemas for software maintainers to adopt.

---

## 🔗 Key Links & References

- **GitHub Repository**: [swicg/activitypub-trust-and-safety](https://github.com/swicg/activitypub-trust-and-safety)
- **Editor's Draft Report**: [ActivityPub Trust and Safety Overview](https://swicg.github.io/activitypub-trust-and-safety/)
- **W3C SWICG Home**: [W3C Social Web Incubator Community Group](https://www.w3.org/community/swicg/)
