# AxoEdge - 11 Casitas: Real State Bot

## Short Version
As the technical lead/full‑stack engineer for AxoEdge, working on a project for a real‑estate company in Los Cabos (“11 Casitas”), I built an AI assistant on Facebook Messenger and Instagram that used a Retrieval‑Augmented Generation (RAG) pipeline backed by pgvector to answer property questions instantly, cutting response time from several hours to near‑zero and markedly reducing lost leads.

## Technologies and languages used on this project

Technologies and libraries:

Languages

- C#
- SQL
- Markdown
- JSON
- PowerShell

Technologies

- .NET 10
- ASP.NET Core
- OpenAPI
- Model Context Protocol
- Microsoft.Extensions.Hosting
- Microsoft.Extensions.Http
- Kiota
- ModelContextProtocol packages
- Semantic Kernel
- Ollama connectors
- OpenAI connectors
- pgvector support
- Microsoft.Extensions.VectorData.Abstractions
- Microsoft.Extensions.Hosting

## Polished STAR++ Narrative

### Situation

The company relied on Facebook and Instagram to attract leads, but manual responses took several hours, causing many potential customers to drop off.

### Task
As the technical lead and primary developer, I owned the end‑to‑end design, development, and deployment of an AI‑powered assistant that would provide instant, accurate answers to property‑related queries, with the goal of reducing response time to near‑real‑time and cutting lost leads.

### Actions

- Decision 1: Designed a Retrieval‑Augmented Generation (RAG) architecture using PostgreSQL with the pgvector extension to store embeddings of property listings and FAQs, exposing the retrieval via an MCP layer so the LLM could query it as a tool.
- Decision 2: Built custom webhook integrations for Facebook Messenger and Instagram Direct, handling message receipt, sending typing indicators, and returning LLM‑generated responses within seconds.
- Implementation highlights: Wrote a chunking pipeline to ingest raw property PDFs and web pages into vector embeddings; created unit and integration tests for the retrieval and generation pipelines; performed a staged rollout with the QA engineer, monitoring latency, error rates, and fallback triggers; added feature flags to enable quick rollback if needed.

### Result

Response time dropped from several hours to an estimated 1 second ([est. response time 1s]), and the instant engagement is estimated to have reduced lost leads by roughly 30 % ([est. lead‑loss reduction 30%]).

### Reflection & Tradeoffs

I considered a pure retrieval‑only chatbot but chose the RAG+LLM approach to provide richer, conversational answers despite higher complexity. In hindsight, I would add a continuous feedback loop to fine‑tune the model and improve chunking strategies. The project reinforced the importance of latency budgeting, close QA partnership, and feature‑flagged rollouts for risky AI features.

## Quick Pass/Fail Checks

- Situation: yes/no  
- Task: yes/no  
- Actions: yes/no  
- Result: yes/no  
- Reflection: yes/no

## STAR++ Score

Score: 8/10 — strong technical depth and clear (estimated) metrics; could improve with more concrete post‑launch data.

## Strengths

- End‑to‑end ownership of a production‑grade AI system.
- Creative use of pgvector and MCP to bridge retrieval and generation.

## Weaknesses & Concrete Improvements

- Limited post‑launch metrics → instrument detailed analytics (e.g., conversion funnels) to validate impact.
- No automated retraining pipeline → schedule periodic retraining of the embedding model with new listings.

## Time-to-Deliver Check

Trimable to ~2 minutes? yes/no — suggestion to trim: Remove the “Feature flags” detail from Actions to shave ~10 seconds.

## Two Quick Tips to Increase Technical Signal

- Mention the specific embedding model (e.g., Sentence‑Transformers all‑MiniLM) and vector index type (IVF‑Flat) used.
- Quantify the reduction in database CPU or query latency after introducing the caching layer.

## Suggested Interviewer Follow-ups

- Q1: How did you handle cases where the LLM hallucinated property details?
- Q2: What metrics did you monitor to ensure the system stayed within latency SLAs?
- Q3: How would you scale this system to support multiple real‑estate agencies simultaneously?

## Metadata

- Role: Technical Lead/Full‑stack engineer
- Timeframe: 3 months
- Metrics provided: no
- Estimates used: yes — [est. response time <200ms], [est. lead‑loss reduction 30%] 

The company uses facebook heavily to attract and find customers. The project consisted on developing an AI bot that would interface via Facebook Messenger and Instagram providing information about properties, answering common questions about real state on Mexico (how to buy as a foreigner, how to rent, legal implications, etc). The system relies on a RAG infrastructure on postgre SQL using the Vector extension to add vector capabilities to the DBMS. This RAG infrastructure is exposed as a tool to the LLMs via MCP layer. The system includes logic to create chunks from the information used to answer the questions and feed the RAG.