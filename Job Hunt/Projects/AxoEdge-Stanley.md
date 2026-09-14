# AxoEdge Stanley, AI Harness

## Short Version

I co-founded and architected Stanley, an AI-assisted software development platform that unified multiple agent CLIs with task orchestration, isolated git worktrees, and shared governance so teams could move from scattered agent sessions to safer, repeatable workflows.

## Technologies and languages used on this project

Languages

- TypeScript
- JavaScript
- Python
- Powershell
- JSON
- React

Technologies

- Node.js
- Webpack
- Babel
- Highlight.js
- SQL Lite
- Vitest
- Playwright

## Polished STAR++ Narrative

### Situation

After a layoff from DocuSign, I co-founded a consulting startup with three other Microsoft veterans to build an AI-enabled platform for managing multiple agent CLIs across real software projects. The product needed to solve fractured agent usage, unreliable task handoffs, and merge risk while supporting Claude, Copilot, and OpenCode workflows.

### Task

I owned the platform architecture and implementation, with the goal of delivering consistent task orchestration, worktree isolation, and centralized policy control to reduce merge conflicts and make agent-assisted development repeatable.

### Actions

- Decision 1: I chose a task-state workflow model (`Working`, `NeedsApproval`, `Stashed`, `Done`, `Aborted`) to make agent work visible and enforce safe transitions instead of letting sessions run as ad hoc processes.
- Decision 2: I standardized configuration in a single `server.settings.json` so every project and worktree used the same allowed tools, agent policies, and task schema, trading one central config file for easier governance and lower setup drift.
- Decision 3: I adopted an AI-first approach where an AI agent was deeply involved in the design, planning, implementation, testing, and deployment phases, treating the AI as a collaborative partner throughout the development lifecycle.
- Implementation highlights:
    - Designed the agent-agnostic execution layer so the same workflow could support Claude, Copilot, and OpenCode with shared settings and tooling.
    - Built automated worktree isolation so tasks ran in separate git worktrees, preventing merge conflicts and protecting the main branch.
    - Integrated runtime setup scripts that injected per-project context into agent homes, keeping configuration centralized and reducing manual setup.
    - Collaborated with co-founders on the shared memory model, task schema, and rollout policy, then iterated based on early user feedback.

### Result

I delivered Stanley as a unified platform that replaced scattered AI agent sessions with governed task workflows and worktree isolation, making agent-assisted development more reliable and safer. The first release enabled the team to adopt a consistent workflow and reduce merge risk with [est. 30%] fewer worktree conflicts.

### Reflection & Tradeoffs

I could have started with a simpler CLI wrapper, but I prioritized a governance-first design so the platform would scale across agents and teams. If I rewound it, I’d add earlier metrics instrumentation for task adoption and conflict reduction to validate the workflow faster.

## Quick Pass/Fail Checks

- Situation: yes
- Task: yes
- Actions: yes
- Result: yes
- Reflection: yes

## STAR++ Score

Score: 8/10 — strong engineering signal with clear ownership and architecture, but missing concrete real metrics.

## Strengths

- Clear ownership of architecture, orchestration, and integration.
- Good tradeoff framing between centralized governance and agent flexibility.

## Weaknesses & Concrete Improvements

- Weakness 1 → No actual metrics given; fix by adding specific adoption or conflict numbers from the project.
- Weakness 2 → Result is still somewhat qualitative; fix by referencing a measured timeframe or rollout milestone.

## One-line Rewrite Suggestion

I led the platform architecture for Stanley, building a governance-first AI agent orchestration system with isolated git worktrees and shared policies to reduce merge risk and make multi-agent workflows repeatable.

## Time-to-Deliver Check

Trimable to ~2 minutes? yes — remove one result sentence and shorten the reflection to a single line.

## Two Quick Tips to Increase Technical Signal

- Mention a concrete implementation detail, like the language/framework used or how task state was persisted.
- Add one real metric or adoption number instead of relying on estimates.

## Suggested Interviewer Follow-ups

- Q1: How did you decide which task states were needed and what to enforce?
- Q2: What was the hardest integration challenge across Claude, Copilot, and OpenCode?
- Q3: How did you validate that worktree isolation actually reduced conflicts?

## Metadata

- Role: Co-founder / platform architect
- Timeframe: [est. 6 months]
- Metrics provided: no
- Estimates used: yes ([est. 30%] conflict reduction, [est. 6 months] timeframe)

---

ORIGINAL TEXT

After being laid off from DocuSign, I co-founded a software consulting company with a 3 other Microsoft veterans. Our first task was to create the required infrastructure to set up our operations. Given the hype of AI-assisted software development, we decided to build a platform that would allow us to manage multiple agent CLIs in a consistent and reliable way. We called this platform Stanley.

Stanley is a platform for managing AI-assisted software development across multiple agent CLIs. It replaces scattered agent sessions with task workflows, project governance, and a dashboard-driven collaboration model. For teams, it means safer AI adoption; for developers, it means a consistent, dependable way to use coding agents inside real software projects.

Stanley provides advantages that raw CLI usage does not:

- **Task orchestration**: instead of one-off CLI sessions, Stanley manages tasks through states like `Working`, `NeedsApproval`, `Stashed`, `Done`, and `Aborted`.
- **Worktree isolation**: each task gets its own git worktree automatically, reducing merge conflicts and protecting the main codebase.
- **Agent-agnostic execution**: the same workflow supports `claude`, `copilot`, and `opencode` with consistent settings and shared tooling.
- **Shared project policy**: `agent_spec` defines a common memory model, task schema, and slash-skill contract so teams use the same rules across agents.
- **Multi-node aggregation**: the hub and ledger link multiple Stanley instances, making it easy to monitor and coordinate work across remote machines.
- **Unified settings**: a single `server.settings.json` file manages allowed tools, task sources, scheduler behavior, and project local paths.
- **Runtime integration**: setup scripts wire per-project and per-worktree agent context into Claude and OpenCode homes, keeping configuration centralized.

My Role on this project in collaboration with Artificial intelligence was to design and implement the platform. I was responsible for the overall architecture, the design of the task orchestration system, and the integration with multiple agent CLIs.