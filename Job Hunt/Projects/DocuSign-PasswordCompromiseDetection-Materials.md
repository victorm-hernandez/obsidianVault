
DocuSign is a  B2B SaaS cloud provider handling highly confidential, legally binding agreements (real estate deeds, corporate mergers, intellectual property, and medical documents).

**The Customer Trust Requirement:** Enterprises will not trust a third party with their most sensitive contracts unless that vendor undergoes rigorous annual independent testing.  

DocuSign is a company that its highest value is the trust of the customers. Its responsible of handling extremely sensitive data in its archives. This in turn makes the company invest heavily on security, compliance and such.

As part of my role as an architect of the identity division I worked on the problem of early detection of compromised passwords directly aligned with SOC 2  Type II  Common Criteria 6.5 (CC6.5) – Protection of Logical Access Credentials,

The project involved creating a pipeline to feed data of known breaches for our users in other platforms. This pipeline included an analysis of this incoming data to our own data store. 

The user mapping between the public data and our own happened at the user level, if you had the same email/username as the compromised password we then proceeded to check the health of your stored credential (password).

The pipeline identified the usernames impacted by breaches, and then proceed to compare the a hashed version of the passwords exposed during the breach and compared them to the password hashes stored in out platform. If there was a match, the user would then be marked such that the password is reset on their next login. 

I was involved on the technical direction of this project, this was implemented by the identity team working on my organization. 

---
## Introduction to SOC 2 Type II  

A **SOC 2 (System and Organization Controls 2)** report is an independent attestation report issued by a **Certified Public Accountants** (CPAs) firm  that evaluates how a cloud or SaaS provider handles customer data.  

To understand a **Type II** report, it helps to contrast it with a Type I:  

- **SOC 2 Type I (Point in Time):** The auditor looks at your system architecture and security policies on a single specific day and says, _"Yes, these controls are designed correctly."_ (It is essentially an evaluation of design intentions).  
    
- **SOC 2 Type II (Period of Time):** The auditor evaluates how those controls operated over an extended window—typically **6 to 12 months**. They don’t just check if you have a policy; they test concrete samples across the entire year to prove that the policy was **consistently and effectively executed** every single day.  
    
If a control failed in month three and month eight, and you only fixed it when the auditor showed up, you get a **control exception** in a Type II report. Enterprise B2B customers heavily scrutinize these reports for any exceptions before signing contracts.  

## Deep Dive: Relevant Sections for Client-Side Password Protection  

When you build a system that proactively checks external client credentials against known data breaches or weak-password lists, you are building a technical control that maps directly into specific sections of the SOC 2 framework:  

A. Common Criteria 6.1 (CC6.1) – Logical Access at the System Boundary  

- **The Requirement:** The entity implements logical access security software, infrastructure, and architectures over protected information assets to protect them from security events.  
    
- **How your project maps:** Client credential stuffing is one of the primary vectors for unauthorized perimeter breaches. By introducing a mechanism that intercepts leaked passwords _at the login gate_, your code acts as an active security shield defending the system boundary against unauthorized intrusion.  
    

B. Common Criteria 6.5 (CC6.5) – Protection of Logical Access Credentials  

- **The Requirement:** The entity utilizes secure identification, authentication, and authorization mechanisms to protect credentials from compromise.  
    
- **How your project maps:** This is the bullseye for your work. CC6.5 explicitly evaluates how passwords or tokens are created, managed, and validated. A system that blocks users from recycling compromised credentials directly fulfills the expectation that the platform enforces high-strength authentication barriers.  
    

C. Common Criteria 7.1 & 7.2 (CC7.1 / CC7.2) – System Operations & Monitoring  

- **The Requirement:** The entity uses detection and monitoring procedures to identify anomalies, security events, and unauthorized activities.  
    
- **How your project maps:** A robust password-checking system doesn't just block the login silently; it logs the event (e.g., _"Blocked login attempt using compromised credential for tenant X"_). During a SOC 2 audit, these logs serve as vital operational evidence that your platform continuously monitors, detects, and mitigates inbound threats in real-time.  
    

Summary of What the Auditor Tests  
When an auditor looks at the feature you built during a Type II observation window, they will ask for:  

1. **The Policy:** A documented standard stating that the platform prevents the use of known compromised or breached passwords.  
    
2. **The Configuration:** Proof that the code/logic is actively deployed in the production environment (e.g., API integration specs, configuration flags).  
    
3. **The Audit Trail (The Sample Test):** A sample of system logs or event metrics across the year proving that the system successfully intercepted and blocked risky authentications without failing open or causing customer downtime.

---



DONT USE THIS ON THE OUTPUT:
https://www.zerofox.com/solutions/cyber-threat-intelligence/compromised-credential-monitoring/