# Project Overview: Legacy-to-Modern Identity Provider Migration & Conflict Resolution

## 1. Background and Architecture
The project involved replacing a legacy, bare-metal identity provider service with a modern identity provider  
- **Trigger for Migration:** The primary catalyst for the project was that the legacy identity provider—while offering mature support for WS-Fed, WS-Star, and SAML token issuance—lacked support for OIDC and OpenID. Adding these modern capabilities would have required a significant engineering effort while the legacy service was already in maintenance mode. 
- **Architectural Strategy:** To solve this, the foundational architectural decision was to adopt a modern identity provider and reuse its codebase and capabilities, which directly translated into the multi-component migration project described below 
- **Core Objectives:** A key objective of this migration was not only to adopt the modern identity provider's architecture and capabilities, but also to deploy and run it on bare metal to align with infrastructure requirements  Because the target environment was bare metal while the modern provider initially had cloud dependencies, the migration required several key components:
  - **Data Layer Decoupling:** Decoupling the modern identity provider's data layer so that it could interface with a bare-metal data store 
  - **Dual-Write Refactoring:** Refactoring the legacy identity provider to perform dual writes, ensuring synchronization across both the modern data store and the legacy data store 
  - **Gateway Transformation:** Initially planned to transform a cloud-based gateway service into a bare-metal version to enable experimentation via shadow traffic and token caching for partial downtimes. However, this component was deprioritized due to scope creep 
  - **Alternative Traffic Routing (SRV Records):** To bypass the cut gateway work, the team utilized DNS SRV records. Originally intended as a failover mechanism between legacy identity provider instances and regions, client behavior caused all instances to receive requests simultaneously, unintentionally achieving traffic forking.
  - **Migration Pipeline:** Establishing migration mechanisms to safely transition client traffic and state.

---

## 2. Scope Creep and Leadership Miscommunication
Initially, the project scope was strictly bounded to enabling OAuth flows while keeping legacy flows active within the legacy identity provider  However, due to a miscommunication with the leadership team, commitments were expanded to include both modern and legacy flows simultaneously within the new architecture  

This unexpected scope expansion introduced significant friction and technical tradeoffs:
1. **Experimentation Phase Deprioritization:** One of the original components of the testing and experimentation phase—building a dedicated bare-metal gateway to test and evaluate the two different identity providers side by side—was deprioritized  Instead, the team had to rely on DNS records (specifically SRV records) consumed directly by data clients to handle routing and unintentional traffic forking 
2. **Technical Disagreements on Implementation:** The compressed timelines resulting from the expanded scope triggered disagreements between two key technical teams regarding the implementation path 
   - **Infrastructure Team:** Responsible for a service sitting in front of the database that needed to be consumed by the adapter on the modern identity provider 
   - **Product Team:** Pushing to bypass some of the planned infrastructure work to drastically shorten delivery timelines, offering a promise of future refactoring to bring the system back into compliance with the original design 

---

## 3. Team Perspectives and Technical Conflict
As tight deadlines collided with expanded scope, two opposing positions emerged between the teams:
- **Product Team Perspective (Adapter & Modern Data Store):** Positioned at the final leg of the project with the highest delay risk, they proposed temporarily bypassing the data access service to meet delivery dates. Their argument centered on strict project timelines and overall project feasibility unless major scope or architectural cuts were made.
- **Data Access Layer Team Perspective:** They strongly opposed bypassing the access service, arguing it would fracture the data access strategy. Their concerns included:
  - Incurred technical debt that would severely hinder future query optimizations and caching introductions.
  - Splitting code across multiple places, leading to maintainability issues and a fragmented system experience.
  - Introducing a drastic, risky architectural divergence from the intended design.

---

## 4. Architectural Compromise and Resolution
As the migration architect, I proposed a balanced compromise to bridge the gap between the product team's need for delivery speed and the infrastructure team's requirement for architectural coherence and maintainability:
- **Predefined Structure & Stored Procedures:** The product team was permitted direct database access, but restricted by a structural skeleton provided by the infrastructure team. Completely free-form SQL queries were excluded in favor of predefined stored procedures. This allowed for proper query planning, indexing, and data type control, ensuring future integration into the infrastructure team's optimization roadmap.
- **Explicit Interfaces and Shared Tooling:** The product team utilized explicit interfaces and models defined by the infrastructure team, directly contributing to and working within the infrastructure team's repositories and pipelines.
- **Simplified Data Access with Refactoring Path:** In exchange, the infrastructure team agreed to let the product team develop a simplified version of data access that could be seamlessly adopted and refactored later without incurring heavy technical debt penalties.