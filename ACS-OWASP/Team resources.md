ACS Google Drive
https://drive.google.com/drive/folders/0AKp51tTQvs4oUk9PVA

ACS Strategic Adoption Plan
Sept 5, 2026
- Definitions
- Project commitments and deadlines
- Competition
- Overall great document, must read for leads.
https://docs.google.com/document/d/1k54-dpbxL32pcxeUUZ2MeZWu7lTb-T4d/edit

ACS Meetings, Meeting materials, transcripts and recordings
	https://drive.google.com/drive/folders/1xk7GXGcn48gCYi_nvxs8PtVYxjkSUGK6

ACS weekly repository status
	https://drive.google.com/drive/folders/1ut6SzVThhS65Ap0Xit_xvTxe09oBw0Wy

Project Overview (canvas slack)
	https://owasp.slack.com/docs/T04T40NHX/F0BPG1USAR3

Specification (Agent Observability Standard) ==Why the name doesnt match the group??==
	https://aos.owasp.org/

Specification GitHub???? (found on the page)
	https://github.com/OWASP/www-project-agent-observability-standard

Public Home Page for the Project
	https://agentcontrolstandard.org/

ACS Slack
	https://agentcontrolstandard.slack.com/join/shared_invite/zt-3upi25akg-iBP68gUN9a1pG71r5oURCw#/shared-invite/email
	
	My invite:
		https://join.slack.com/t/owasp/shared_invite/zt-4ab7i2i9y-QUj7C4Pht5GW7KLaPZhZxQ

New contributor interest form
https://docs.google.com/forms/d/e/1FAIpQLSd0pPoz5aPvgYrw4mCiafd-Do-BngEvl7Vk5f3a5_Psb68RWA/viewform
New
---
## Notes from from strategic adoption plan
https://docs.google.com/document/d/1k54-dpbxL32pcxeUUZ2MeZWu7lTb-T4d/edit

### ACS History
v1 vs v2, donated to OWASP on Sept 2.

### Glossary 
ACS-Core
HMAC-SHA256
ACS-Cryto

### AgBOM
An AgBOM (Agent Bill of Materials) is a structured, dynamic inventory of all components comprising an AI agent system, including tools, models, capabilities, and dependencies.
https://aos.owasp.org/spec/inspect/

What is an AgBOM?
• Definition: A transparent tracking standard for AI agent architectures, behavior, and security postures.
• Purpose: It supports security tracing, version tracking, and regulatory compliance in enterprise and regulated environments.
• Standards Extension: Rather than introducing a brand-new format, it extends existing proven industry standards such as CycloneDX, SPDX, and SWID to accommodate AI agent-specific elements.
• Projexct Context: Developed under the OWASP Agent Observability Standard (AOS)

### CycloneDX
https://cyclonedx.org/

CycloneDX is a lightweight, full-stack Bill of Materials (BOM) standard created by the OWASP Foundation to reduce security risks in supply chains.
It provides a machine-readable format—typically JSON, XML, or Protocol Buffers—to inventory all moving parts in modern applications and systems.
Key Capabilities and Types of BOMs

CycloneDX goes far beyond traditional software lists. It supports multiple advanced BOM formats:
• SBOM (Software Bill of Materials): Inventories first-party and third-party software components, libraries, and transitive dependencies.
• HBOM (Hardware Bill of Materials): Catalogs physical hardware parts and devices.
• CBOM (Cryptography Bill of Materials): Details cryptographic algorithms, certificates, and assets used in code.
• ML-BOM (Machine Learning Bill of Materials): Maps out AI/ML models, training data provenance, and transparency metrics.
• SaaSBOM: Represents cloud services, APIs, and microservices powering cloud-native software.
• VEX (Vulnerability Exploitability Exchange): Acts as an "anti-advisory" to state whether a component with a known vulnerability is actually exploitable in your specific environment.
Core Use Cases
• Vulnerability Management: Helps teams quickly pinpoint outdated dependencies and known security flaws.
• License Compliance: Tracks software licenses and copyrights to avoid legal conflicts.
• DevOps Integration: Easily plugs into build pipelines (like Maven, Gradle, or Python tools) to generate accurate inventories automatically during compilation.

### OCSF
https://ocsf.io/

OCSF stands for the [Open Cybersecurity Schema Framework](https://www.google.com/goto?url=CAESSAHrOzAVHwCZ2rcdm-TOtyU3KBOLny5xkC99ByQ_3SqGMXlP1OB1BKAVNOi2llzMSZf1zWmv4-h7-jvxMYgmdglSTO4LgfFpaw), an open-source, vendor-agnostic standard used to normalize and structure security event data.

What OCSF Does

- **Standardizes Language**: Gives different security tools a common dictionary and taxonomy so an event looks the same regardless of who made the tool.

- **Normalizes Data**: Converts messy, mixed log formats into neat JSON structures.

- **Simplifies Analysis**: Helps security analysts search, correlate, and investigate threats across different products without rewriting custom parsers.

Why It Matters

- **Stops Vendor Lock-In**: Frees organizations from being tied to one proprietary data model.
- **Aids AI and Lakes**: Makes it easier for modern data lakes and AI tools to read and reason across mixed security logs.
- **Industry Backed**: Started by major tech companies and now governed under the [Linux Foundation](https://www.google.com/goto?url=CAESVwHrOzAVwYld6myFg5DmszFW2HBI_qGzfDKXQmNxr7tV7gd-vcrAfiRMViig5_Hhn0w6v2v2HzPi-ODir3PSlJbXrvwjK9Obg4nbFNXJFLEmC0he0LRGww).

### SPDX
https://spdx.dev/

**SPDX** (System Package Data Exchange, formerly Software Package Data Exchange) is an open standard used to share software bill of materials (SBOM) data, including components, licenses, copyrights, and security references

Driven by the Linux Foundation, SPDX provides a common machine-readable format so organizations can track what is inside software. It is also recognized as an international standard (ISO/IEC 5962:2021).

Key Features of SPDX

- **License Compliance:** Maintains a standardized list of short identifiers for open-source licenses to track legal compliance.
- **Supply Chain Security:** Communicates security data, dependencies, and software provenance to vulnerabilities.
- **Expanded Scope (SPDX 3.0):** Covers new domains like AI models, datasets, and build lifecycles beyond traditional software packages

### SWID
https://csrc.nist.gov/Projects/Software-Identification-SWID

SWID most commonly stands for **Software Identification (SWID) Tag**, a standardized file used to name, identify, and track software installed on a computer or device

A SWID tag is an international standard (ISO/IEC 19770-2:2015) machine-readable file—usually written in XML or CBOR format—that provides details about a specific software product. It acts like a digital name tag or fingerprint for an application.

Key Details

- **Metadata included:** Product name, version, edition, vendor/publisher, and unique identifiers.
- **Lifecycle tracking:** The tag is automatically added when software is installed and removed when the software is uninstalled.
- **Main uses:** Helps organizations with [Software Asset Management (SAM)](https://www.google.com/goto?url=CAESfgHrOzAV_cH5plV9CncbNpo9GWKPisvLZCrh4l1L13uMEZDDbYsvKYPugRf8aKV0Yi3Ze9ngR0afVukmVUOXV5ohABQy7DQI2JMIcYuQJOsWGI-U-scU00uBfJwekUKhcqJrI5lp-M4qnBy8rKgoa2cSyPscFdNYSZK0JiM7UA), license compliance, and cybersecurity inventory tracking.

Types of SWID Tags

- **Primary Tag:** Identifies an installed software product.
- **Patch Tag:** Describes an installed update or patch.
- **Corpus Tag:** Represents pre-installation files or installers.
- **Supplemental Tag:** Adds extra metadata or custom tracking info without altering the original tag.

### A2A

**A2A (Agent-to-Agent)** is an open standard protocol that allows different artificial intelligence (AI) agents to discover, communicate, and delegate tasks to one another across different frameworks and platforms.

How A2A Works

- **Agent Discovery:** Agents publish standardized JSON files called **Agent Cards** at a well-known URL to broadcast their specific capabilities, skills, and endpoints.
- **Task Delegation:** A primary coordinator agent assigns a structured goal or task to a secondary peer agent rather than just calling a basic API function.
- **Communication:** The protocol uses standard web technologies like HTTP/HTTPS, JSON-RPC 2.0, and Server-Sent Events (SSE) for real-time streaming of work progress and artifacts. 

A2A vs. MCP (Model Context Protocol)
- **MCP** handles vertical integration by connecting a single AI agent downward to external data sources, databases, and digital tools (like giving an agent a toolbox).
- **A2A** handles horizontal integration by connecting multiple autonomous AI agents together so they can act as a coordinated team
### LangChain

[LangChain](https://www.google.com/goto?url=CAESWwHrOzAVeEiiaXm6nM21UoANDGyDzhzR4Edb9uHMPA-6Kcm4xnxdYnlgnqM0tHTwRrdbjZuPqHS7CwrkGsrMtw1GEukw8J_cGjP3FDZGe-_gTW7_KvaOCMEzE5M) is an open-source software framework designed to help developers build applications powered by large language models (LLMs). 

Core Purpose

LLMs are powerful text processing engines, but they are isolated from private data sources, lack long-term memory, and cannot independently use external tools. LangChain acts as an orchestration layer. It bridges the gap between raw models (like GPT-4, Claude, or Gemini) and external data, APIs, or business logic.

Key Building Blocks

- **Models:** A unified interface to connect with hundreds of different LLM and embedding providers so developers can switch models easily.
- **Prompts:** Tools and templates to manage, reuse, and structure dynamic instructions sent to models.
- **Chains:** Sequences of operations that feed the output of one task directly into the next for multi-step workflows. 
- **Memory:** Utilities that store and recall past conversation history to give stateless models context awareness over time.
- **Retrievers & RAG:** Systems that split private documents, convert them into vector embeddings, and search them to provide relevant context to the LLM (Retrieval-Augmented Generation).
- **Agents:** Logic that lets an LLM act as a reasoning engine to dynamically choose and execute external tools or APIs based on user requests.

Common Use Cases

- **Chatbots:** Interactive assistants with conversation memory and access to customer service data.
- **Question-Answering Systems:** Apps that search through company documentation, internal databases, or PDFs to answer specific queries.
- **Autonomous AI Agents:** Multi-step workflows where independent agents collaborate or use external tools to complete complex tasks.


### JSON-RPC 2.0

**JSON-RPC 2.0 is a stateless, lightweight remote procedure call (RPC) protocol encoded in JSON**. It allows applications running on different systems to communicate over a network by enabling a client to call a function or method on a remote server as if it were a local function.

Because it is **transport-agnostic**, JSON-RPC 2.0 can send messages over any network protocol, including HTTP, WebSockets, TCP/sockets, or standard input/output (STDIO). It is widely used in systems requiring fast, low-overhead communication, such as [Ethereum and blockchain development](https://www.google.com/goto?url=CAESZQHrOzAV_0IeFBguYGuw2s6uEQ17vjS5R12OcgNhWyg_kGX4_r1OJgLf2lbpAR8Y1d6KPEMIJSbBpR5aH1fKTlP6QS21wB9yRlJqX1dRhYL9i8Ru3Cyz0E_-a8q08TQhTU0yENp1), Microsoft's Language Server Protocol (LSP), and the [Model Context Protocol (MCP)](https://www.google.com/goto?url=CAESdgHrOzAV8DeyZWLl1CLUkBLZuQFRgmmrxuD6WF2p-rDyQfPnnsoZcfJW1b_ntm_-PfXGUaLlECC5k2KA0rdGEdRubxfgR6kIY-Ew43-48qdCHbTuDGtFhGms8Nf6DuCeRe1fSUY9U8OckETkXarEEhZlSNURaVE) for AI. 

🧱 Core Message Structures

The protocol is defined by a tiny, simple specification consisting of two primary structures: Requests and Responses.

1. Request Object
A client invokes a remote method by sending a payload containing four specific fields:
- `jsonrpc`: A string that **must** be exactly `"2.0"`.
- `method`: A string containing the name of the function to be invoked.
- `params`: An Object or Array containing the arguments passed to the method.
- `id`: A unique identifier (string, integer, or null) established by the client to map responses back to requests.

**Example Request:**

json

```
{
  "jsonrpc": "2.0",
  "method": "subtract",
  "params": {"minuend": 42, "subtrahend": 23},
  "id": 1
}
```

2. Response Object

When the server finishes executing the method, it returns a response containing the same `id` sent by the client. It must contain either a `result` or an `error` object, but never both:

- `result`: The data returned by the invoked method if it succeeded.
- `error`: An object containing a numeric `code`, a string `message`, and optional `data` if the method failed.

**Example Response:**

```
{
  "jsonrpc": "2.0",
  "result": 19,
  "id": 1
}
```


⚡ Key Features of Version 2.0

- **Notifications:** If a client sends a Request object _without_ an `"id"` member, it is treated as a Notification. This signifies that the client does not want or expect a reply, and the server **must not** return a response. 
- **Batch Requests:** Clients can group multiple Request objects into a single JSON array. The server processes them concurrently or sequentially and returns a single JSON array containing all corresponding responses. 
- **Named Parameters:** While version 1.0 only allowed parameters to be passed sequentially as an array, version 2.0 allows parameters to be structured as a JSON Object, enabling parameters to be looked up by name.
  
🔄 JSON-RPC vs. REST APIs

| Feature          | JSON-RPC 2.0                                               | REST                                                                      |
| ---------------- | ---------------------------------------------------------- | ------------------------------------------------------------------------- |
| **Core Concept** | **Actions / Procedures** (Execute a verb like `getUser()`) | **Resources** (Manipulate a noun like `/users`)                           |
| **Transport**    | Transport-agnostic (HTTP, Sockets, STDIO)                  | Rigidly tied to HTTP methods (`GET`, `POST`, `DELETE`)                    |
| **Efficiency**   | High (Supports batching and lightweight notifications)     | Moderate (Every action requires an individual request/response lifecycle) |
| **Errors**       | Custom internal application error codes                    | Leverages native HTTP Status Codes (`404 Not Found`)                      |


---


Post quantum signatures
44 schema URI of ACS (part of github currently)
MCP
A2A

Terms from the website : https://agentcontrolstandard.org/
OCSF 
SIEM infrastructure

Adapters exists for: Claude Code, Cursor and NVIDIA NAT

Reference implementation lead seat is free according to plan. ==TODO check with Rock==
### Similar projects
- Microsoft AGT
- Galileo Agent Control 
- AgentGateway
- Portkey
- Kong
- AGNTCY


