# Short Version

On the Microsoft Account team in 2011, I owned the authentication integration between Windows 8 and MSA. When I needed to adapt the WS-FED web protocol to support OneDrive's roaming features, the protocols architect pushed back hard — and our first meeting ended in a walkout. I recognized I'd been arguing to win rather than listening to understand. I went back, actively paraphrased his concerns, validated the protocol-integrity risks he owned, and then co-designed a compromise that satisfied both sides. The feature shipped on time and unblocked OneDrive's Windows roaming capabilities.

---

# Polished STAR++ Narrative

## Situation
In 2011, as a Senior SDE on the Microsoft Account UX team, I owned the authentication integration for Windows 8, Xbox, and HoloLens. The system used platform-specific protocols (Win8InlineLogin, Win8WAB), and we were under tight release pressure to ship alongside Windows 8.

## Task
I needed to adapt the WS-FED web-authentication protocol — owned by a separate protocols team — to render in a Win8-skinned experience, unblocking OneDrive as the lead partner and enabling Windows roaming features. My constraint was getting sign-off from the protocols architect without compromising protocol integrity.

## Actions
- **Identified the technical gap:** During deep design work I discovered the existing Win8 protocols didn't cover the OneDrive roaming scenario; WS-FED was the only viable path but required a skinned customization that crossed team ownership boundaries.
- **First meeting (failed approach):** I called a meeting with the protocols architect to walk through the design. He had prior friction with the UX team from recent incidents and didn't have full context. I went in prepared to defend my plan — listening only to rebut — and the meeting ended with him walking out.
- **Recognized the mistake:** I reflected on the dynamic and realized I had been arguing to win rather than seeking shared understanding. I revisited the principles from *The 7 Habits of Highly Effective People*: listen actively, paraphrase until the other person confirms you've understood them, validate their perspective before presenting yours, then collaborate toward a better solution.
- **Second meeting (reset approach):** I requested a follow-up and opened by asking him to walk me through his protocol-integrity concerns. I paraphrased each concern back until he confirmed I understood it. With that foundation, I presented the detailed flow and framed the discussion as a joint design problem, inviting him to co-own the solution.
- **Reached a compromise:** Together we identified constraints that protected WS-FED's integrity while enabling the Win8 skin. The resulting design satisfied his requirements and kept my flow intact.

## Result
The feature shipped on time with Windows 8 (~2011), unblocking OneDrive as the primary partner and enabling Windows roaming features for Microsoft Account users. [est. partner unblock: 1 major partner; downstream: millions of Win8 users gaining roaming capability]

## Reflection & Tradeoffs
The alternative would have been to escalate to management — fast but damaging to the cross-team relationship and likely to produce a forced compromise neither team owned. The real lesson was about listening posture: entering a disagreement with the goal of understanding rather than winning produces better technical outcomes and stronger working relationships. If I did this again, I'd seek to understand the other team's concerns *before* the first formal meeting to prevent the initial blowup entirely.

---

# Quick Pass/Fail Checks
- Situation: yes
- Task: yes
- Actions: yes
- Result: yes (with labeled estimate)
- Reflection: yes

# STAR++ Score
Score: 7/10 — Strong narrative arc and interpersonal depth; result metrics are estimated and technical implementation detail is light.

# Strengths
- Clear before/after contrast in listening approach makes the growth arc compelling and memorable.
- Names a concrete framework (*7 Habits*) that grounds the behavioral change in a repeatable method.
- Shows cross-functional ownership: navigating protocol-team dependency without authority.

# Weaknesses & Concrete Improvements
- **Thin technical signal** → Add 1–2 sentences on *what specifically* the WS-FED modification entailed (e.g., "I added a query parameter allowing the Win8 shell to inject a skin token while the protocol endpoint remained unchanged").
- **Unquantified result** → Add any available metric: time saved vs. escalation path, number of roaming scenarios unblocked, or OneDrive DAU/MAU impact at launch.
- **Incident context unexplained** → Briefly note what the prior UX-team incidents were and why they were relevant to his bias, to show you understood the root of the conflict.

# One-line Rewrite Suggestion
"By switching from a defensive listening posture to active paraphrasing, I co-designed a WS-FED skin extension with the protocols architect that shipped on time and unblocked OneDrive roaming for Windows 8."

# Time-to-Deliver Check
Trimable to ~2 minutes? yes — Drop the *7 Habits* enumeration in the Actions and keep only the three verbs (listen, paraphrase, validate); saves ~20 seconds.

# Two Quick Tips to Increase Technical Signal
- **Name the technical constraint explicitly:** Describe in one sentence *why* WS-FED required modification (e.g., the endpoint contract, the skin injection mechanism) — this shows protocol-level fluency.
- **Quantify the compromise:** State what the protocols architect required you to change or preserve (e.g., "we agreed the endpoint URL and token format stayed unchanged; only the UI rendering layer was parameterized") — interviewers will probe exactly here.

# Suggested Interviewer Follow-ups
- Q1: "What specifically did the architect want to protect about WS-FED, and what did you agree to change versus keep unchanged?"
- Q2: "How did you handle the risk that the skinned WS-FED experience could introduce a regression in the web auth path?"
- Q3: "If the architect had continued to refuse, what would your escalation path have been, and what would you have weighed before taking it?"

---

# Metadata
- Role: Senior Software Engineer, Microsoft Account UX Team
- Timeframe: ~2011, Windows 8 release cycle
- Metrics provided: no
- Estimates used: yes — `[est. partner unblock: 1 major partner (OneDrive); downstream: millions of Win8 users gaining roaming capability]`