
### **Talk Proposal: From Post-Mortems to Pre-Emptive Control: Runtime Governance for AI Agents with the Agent Control Standard (ACS)**

#### **Suggested Title**

- **Primary:** _Governing the Autonomous Loop: Runtime Interception and Policy Enforcement for AI Agents_
    
- **Alternative:** _Stop Guessing What Your Agents Are Doing: Building Runtime Control Planes with ACS_
    

#### **Target Audience**

Software engineers, platform architects, security engineers, and tech leads building or scaling LLM-driven applications and agentic workflows.

#### **Objective**

To provide software engineers with a practical mental model and architectural blueprint for moving away from brittle system prompts and after-the-fact observability, demonstrating how the **Agent Control Standard (ACS)** introduces standardized middleware hooks, deterministic guardrails (Guardian Agents), and dynamic inventorial tracking (AgBOM) into modern agent runtimes.

#### **Overall Description**

AI agents are moving from experimental proof-of-concepts to autonomous production systems—calling tools, executing code, modifying databases, and chaining decisions across networks in milliseconds. Yet, engineering teams are still trying to govern this autonomy using brittle system prompts, custom ad-hoc validation logic, or post-execution observability tools that only tell you _how_ your system broke after the damage is done.

Just as OAuth standardized authentication and OpenTelemetry standardized metrics, the **Agent Control Standard (ACS)** provides an open, framework-agnostic runtime control plane for AI agents. In this talk, we will explore how ACS decouples safety and governance from agent application logic. We'll examine the three-layer architecture—exposing standardized middleware hooks at decision points (Tier 1), enforcing deterministic policies via open-source Guardian Agents (Tier 2), and plugging in domain-specific classifiers (Tier 3)—while showing how dynamic Agent Bills of Materials (AgBOM) solve supply chain visibility for software that rewrites itself.

#### **Suggested Talking Points & Outline**

1. **The Architectural Blind Spot in Agentic Systems**
    
    - Why system prompts, model fine-tuning, and static guardrails fail to act as true runtime boundaries.
        
    - The limitation of post-execution observability: _Logs tell you why it failed; control planes stop it from happening._
        
    - Parallels to web security: How we moved from unchecked callbacks to API gateways and middleware filters.
        
2. **Anatomy of the Agent Control Standard (ACS)**
    
    - What is ACS? Transitioning from observability (AOS) to active runtime control.
        
    - The core primitive: Standardizing decision-point middleware hooks across input, output, tool selection, code execution, and memory ops.
        
    - The Three-Layer Model:
        
        - _Tier 1 (Platform Layer):_ Exposing hooks without requiring agent platforms to implement policy logic.
            
        - _Tier 2 (Enforcement Layer):_ Framework-agnostic SDKs and Guardian Agents returning `allow`, `deny`, or `modify` verdicts.
            
        - _Tier 3 (Enterprise Customization):_ Plugging in custom classifiers, PHI filters, and organizational policy-as-code.
            
3. **Under the Hood: How Interception Works**
    
    - Walkthrough of a tool-call lifecycle intercept.
        
    - Integrating structured telemetry with OpenTelemetry semantic conventions and OCSF mappings.
        
    - Supply chain security for dynamic runtimes: Tracking capabilities via dynamic Agent Bills of Materials (AgBOM) extending CycloneDX and SPDX.
        
4. **Adopting ACS in Your Stack**
    
    - How to instrument an existing agent framework without rewriting business logic.
        
    - Contributing to the open-source reference implementations, upstream OTel mappings, and shaping the future of autonomous governance.