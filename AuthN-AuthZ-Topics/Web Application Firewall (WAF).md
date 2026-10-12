A **Web Application Firewall (WAF)** is a specialized security system that protects web applications by filtering, monitoring, and blocking HTTP/HTTPS traffic between a client application (such as a browser or mobile app) and a web service.

Unlike traditional network firewalls (Layers 3 and 4) that inspect IP addresses, ports, and transport protocols, a WAF operates at **Layer 7 (Application Layer)**. It understands application-level protocols and semantics, allowing it to inspect payloads, cookies, headers, and query strings to detect and block sophisticated attacks such as:

- SQL Injection (SQLi)
- Cross-Site Scripting (XSS)
- Local/Remote File Inclusion (LFI/RFI)
- Cross-Site Request Forgery (CSRF)
- Application-layer Denial of Service (DoS) and bot-driven credential stuffing

## Examples of WAFs

WAF solutions are typically deployed as cloud-managed services, reverse proxies, or software plugins:

- **Cloud-Native / Edge WAFs:** Cloudflare WAF, AWS WAF, Akamai App & API Protector, and Google Cloud Armor.
- **Open-Source & Software WAFs:** ModSecurity (an open-source engine frequently paired with NGINX or Apache), Coraza (a modern Go-based rewrite of ModSecurity), and NAXSI (an NGINX module focused on heuristic rule sets).
- **Enterprise Next-Gen WAFs (NGWAF):** F5 Advanced WAF (formerly BIG-IP ASM), Imperva Cloud WAF, and Palo Alto Prisma Cloud.

## Deployment Topology: Does it Live on the Same Node as a Load Balancer / API Gateway?

In modern architectures, a WAF is **rarely deployed as a separate physical "node"** or middlebox in the traditional sense. Instead, its placement depends on the deployment model:

1. **Integrated with API Gateways & Reverse Proxies (Same Node / Process):** In cloud-native or Kubernetes environments, WAF capabilities are frequently embedded directly into the edge proxy, API gateway, or ingress controller layer (e.g., Kong with a ModSecurity/Coraza plugin, Envoy proxy equipped with WASM-based security filters, or NGINX Plus). In this setup, the load balancing, API routing, and L7 security inspection execute on the **same compute node and within the same network proxy layer** to minimize network hops and latency.
    
2. **Cloud/Edge Proxy (SaaS / CDN Layer):** Services like Cloudflare or AWS WAF operate at edge PoPs (Points of Presence) distributed globally, far upstream from your application servers and load balancers. They intercept traffic at the network edge before it ever reaches your infrastructure nodes.
    
3. **Dedicated Virtual/Hardware Appliance:** In strict enterprise or on-premises data centers, a dedicated WAF appliance (such as an F5 BIG-IP) may sit inline right behind the external load balancer—or act as the load balancer itself—processing traffic before it is routed to internal microservices.

## Architecture of a WAF System

A modern L7 WAF architecture consists of several core components designed to process high-throughput HTTP traffic with minimal latency:

```
[ Client / Browser ] 
       │
       ▼
[ Edge / CDN / DNS ] (Geo-routing & DDoS mitigation)
       │
       ▼
[ WAF Engine / Reverse Proxy ] 
  ├── 1. TLS Termination & Parsing (HTTP/1.1, HTTP/2, HTTP/3 / QUIC)
  ├── 2. Inspection Pipeline
  │     ├── Signature / Rule Engine (OWASP Top 10, CVE patterns)
  │     ├── Protocol Compliance Checks (RFC validation)
  │     ├── Behavioral & Anomaly Detection (Machine Learning)
  │     └── Rate Limiting & Bot Management (Fingerprinting/CAPTCHA)
  └── 3. Decision Matrix (Allow, Block, Challenge, Log)
       │
       ├──► [ Block / Challenge ] ──► Return 403 Forbidden / CAPTCHA
       │
       ▼ (Pass)
[ API Gateway / Load Balancer ] 
       │
       ▼
[ Microservices / Backend Nodes ]
```

### Core Architectural Components

- **TLS Termination & Protocol Parser:** The WAF must first decrypt TLS traffic (supporting modern standards up to TLS 1.3 and QUIC) and parse raw byte streams into structured HTTP request objects (methods, headers, cookies, query parameters, and body payloads).
    
- **Inspection Pipeline & Rule Engine:** Requests flow through a modular rule evaluation pipeline. Engines like ModSecurity utilize regular expressions and declarative rule languages (such as the OWASP Core Rule Set) to scan strings for attack signatures. Modern engines also incorporate **machine learning models** to detect anomalous data transfer patterns and zero-day exploits.
    
- **State Management & Rate Limiting Engine:** WAFs track client state across requests (e.g., tracking request frequency per IP, token bucket algorithms for rate limiting, or session behavior analysis) to mitigate brute-force and credential-stuffing attacks.
    
- **Decision and Action Module:** Based on the inspection score, the WAF executes a policy action: **Allow** (forward to upstream backend), **Block** (return a `403 Forbidden`), **Log** (audit mode), or **Challenge** (issue a JavaScript challenge, proof-of-work, or CAPTCHA).