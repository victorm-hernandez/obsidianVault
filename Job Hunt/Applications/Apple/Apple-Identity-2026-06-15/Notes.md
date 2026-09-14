# Apple Identity Team

1. Manager, Daniel Boromisa, 26 years working, Java Dev, SRE, PostgreSQL.

2. Dev, Patrick East, 15 years, different team (RBAC?), storage knowledge, open policy, principal at oracle infra.
3. Dev, Chris Rice, 9 years, Docker, Linux, MicroServices.
4. Dev, David Johnston, 7 years, different team (RBAC?), build infra, deployment (ansible/podman)
5. Dev, Matt Hawkins, 6 years, security (VISA)

## Contacts

SRE background
- Engineering Manager: https://www.linkedin.com/in/danielboromisa/details/experience/


Derek McQuay
https://www.linkedin.com/in/derek-mcquay-057a4b66/

RIGHT Matthew Hawkins
    https://www.linkedin.com/in/matthew-hawkins-1b8695189/
    - Skills: Go, Java, Python, Shell, Security (worked for VISA)
    - Experience: 1 year on Apple, 6 years total.


WRONG Matt Hawkins
    https://www.linkedin.com/in/reillyhawk/
    - Experience: 10 years apple, 20 years total.
    - Skills: Web Dev experience, Ruby, Ruby on Rails, Test Driven Development, Mobile Developer (iOS/Android), Agile Methodologies.

Chris Rice
- Skills: Go, Docker, Linux, Microservices, Java, Ruby, Clustered Systems
- Experience: 7 Years at apple, total 9 years as engineer.
Graduated on 2016, Colorado School of Mines BS in CS
https://www.linkedin.com/in/chris-rice-388577156/

---

David Johnston (IC from different Manager)

- Area: Familiar with DevOps
- Skills: Go, DevOps, IAM, RUST, PKI, Developer Infrastructure (Compiler, Build Systems, Test harness
- Dev Ops tools: Podman/Ansible
- Experience: 5 at apple, 7 total.

https://www.linkedin.com/in/dwtj/

---

Patrick East (IC from different Manager)

- Area: Familiar with storage solutions and RBAC/ABAC SaaS solutions.
- Interview assigned to him: migration due to storage affinity
- Experience: 5 years at apple, total 15 years, principal at Oracle Infrastructure before.
- Skills: Go, Python, Java, C++
- Notes:
    Developer at Pure storage for 5 years, EverPure Storage Solutions.
    Open Policy Agent, an open-source, general-purpose policy engine that decouples policy decision-making from software code.
- LinkedIn: https://www.linkedin.com/in/patrick-east-9789ba84/

## Tech stack

- Cockroach DB
- Golang
- SPIFFE/SPIRE

## Tech Topics

### mTLS

Why Use mTLS Over Passwords?
Because mTLS happens at the network layer (Layer 4) rather than the application layer (Layer 7), it is incredibly secure for machine-to-machine communication.

If an attacker tries to connect to an mTLS-protected API without a valid, CA-signed client certificate, the server's network stack will drop the connection instantly during the handshake. The attacker never even gets the chance to guess a password, exploit an application vulnerability, or send a malicious HTTP payload, because the connection is severed before the application layer is even reached.
