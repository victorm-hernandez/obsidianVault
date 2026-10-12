## STAR System Narrative

### dSTS Identity Provider Modernization & Bare-Metal Migration

**Situation:** 
dSTS is Azure's Tier 0 identity provider, serving as the root trust anchor for all Azure infrastructure service-to-service authentication. The team was scaling rapidly (from 5 to 30 engineers) to modernize the platform and introduce protocols like OAuth and OpenID Connect. However, the existing data plane was built on top of retired Azure Access Control Service (ACS) infrastructure. With the original creators gone, there was zero tribal knowledge left. My team was burdened with supporting obsolete legacy infrastructure—from the CI/CD pipelines down to the data store—making any modernization efforts highly risky and operationally expensive.

**Task:** 
I was tasked with defining and leading the technical roadmap to modernize the dSTS data plane. My goal was to migrate the platform to a healthy, modern codebase (Entra ID) capable of running at the Tier 0 bare-metal level. This required overcoming a major technical blocker—Entra ID's dependency on a Tier 1 data store—while simultaneously managing heavy pushback from internal stakeholders who feared losing ownership of the legacy systems.

**Action:** 
* **End-to-End Architectural Ownership:** I researched and formulated a comprehensive proposal to build a new data layer that allowed Entra ID to run on bare metal. To avoid operational DBA overhead for my developer-focused team, I evaluated and selected a SQL Server data store utilizing in-house sharding technology from the Microsoft Account team.
* **Cross-Team Collaboration & Stakeholder Alignment:** I owned the project end-to-end, working directly with multiple technical leaders across different component teams. To overcome organizational pushback, I proactively integrated legacy IdP stakeholders into the new roadmap. I carved out specific leadership roles for them in the new architecture, securing management buy-in and ensuring they retained scope and ownership.
* **Risk Mitigation via Shadow Traffic:** I proposed and led the deployment of a new bare-metal Service Gateway to safely experiment with the architecture. This allowed us to duplicate shadow traffic and compare success rates, latency, and health metrics between the legacy IdP and the new Entra ID stack in real-time.
* **Safe Data Migration Strategy:** I designed a multi-stage migration plan to guarantee zero-downtime consistency. This included an initial static migration, a dual data layer to keep legacy and new stores synced, and a continuous-monitoring daemon to detect any data inconsistencies on the fly.

**Result:** 
I successfully achieved alignment across multiple engineering orgs and delivered a pilot deployment of the modernized bare-metal IdP. The shadow traffic and multi-stage migration pipeline proved that the new SQL-backed data layer could handle Tier 0 requirements safely and consistently. This massive architectural shift successfully decoupled the team from unsupported legacy infrastructure, eliminated the severe operational overhead, and unblocked the roadmap to fully support modern OAuth/OIDC protocols across Azure.


---

## Original Text

As part of the dSTS team, my first task was to bring modern protocols like OAuth and OpenID Connect to the existing identity provider. As historical background, the identity provider data plane was initially part of Azure Access Control Service (ACS). Introduced around 10 years ago in Windows Azure, this was a SaaS offering designed to provision and manage your own identity provider and handle access control.

https://learn.microsoft.com/en-us/archive/blogs/mcsuksoldev/azure-access-control-services-creating-a-custom-identity-provider

This service was eventually retired, and the data plane of this solution was used by the internal Azure infrastructure team to stand up a Secure Token Server (STS). This enabled service-to-service authentication and eventually evolved into a full-fledged identity provider that supports all of the Azure infrastructure running on the lowest Tier 0 of its architecture. Every service within Azure directly or indirectly depends on it.

I joined this team as part of an effort to modernize this service. The effort included growing the team from 5 to 30 engineers and defining the roadmap for the modernization of the platform.

I started my ramp-up by spending some time working closely with engineers to understand the day-to-day pain points. I quickly realized that one of the biggest issues with modernizing this product was the lack of tribal knowledge; the initial team that built the data plane based on this older offering had already left. Furthermore, most of the modern efforts were focused on the control plane of the service. On top of that, the data plane was sitting on a lot of legacy infrastructure that was no longer supported. This extended all the way from the data store that was used to the continuous deployment pipelines. This caused significant overhead for the team, as much of the infrastructure needed to be supported by the team itself since there was no one else left to support it. Another dimension of this was that making changes to the data plane carried considerable risk given the lack of knowledge in the codebase and the aging infrastructure.

At this point, I started working with some of the main stakeholders on Entra ID (formerly Azure Active Directory) to understand what it would take to bring part of the data plane from this Identity Provider to the Tier 0 level. I found that the main blocker was a dependency on the data store, which was another Tier 1 service. 

With this information in hand, I formulated a proposal to create a new data layer for the data plane so that it could run on bare metal. This would enable us to bring over a healthy and modern codebase while retaining the main characteristic required for this system: being bare-metal. This proposal involved multiple pieces, and each piece was owned by different teams. There was the creation of a new data layer that needed to live closely to the data plane of Entra ID. There was the creation of a new service, which would serve as a caching and data store for the information of the identity provider. There were the deployment artifacts of the new data store, which in this case was based on SQL Server and an in-house sharding technology that came from the Microsoft Account team. Finally, there was a proposal to bring another service to bare metal as a way to safely experiment: a Service Gateway. This would enable shadow traffic for every request so we could compare the runtime of both IdPs, including success rate, latency, and other metrics.

Some of the challenges of this project went beyond the technical aspects. Understandably, there was some pushback to this proposal since, fundamentally, it replaced a lot of the work that was being done as part of the initial identity provider. The strategy that I used to mitigate this was to make sure that I involved all the main stakeholders of the original identity provider as part of the migration. I structured it in such a way that they would all have a stake in the project and an important role in the roadmap, establishing themselves as leaders in their own spaces without worries of losing ownership or scope.

When it comes to the technical challenges, I would like to call out finding the proper data store. It had to fulfill the existing requirements of availability and latency, and ideally, it needed to be a data store that could be operated by my organization without requiring new engineering resources or new know-how. I was working with teams of developers, not DB administrators, and database operations was not a task we wanted to take on.

The result of this work was a pilot deployment of this new version, where we migrated the data from the original identity provider using a multi-stage migration with safe deployment practices. The migration had multiple levels to guarantee consistency on the fly. This included an initial static migration, a dual data layer for the legacy IdP to make sure data was current in both data stores, and a daemon that continuously identified any data inconsistencies, among other measures. We monitored the health and performance of the system to guarantee stability.