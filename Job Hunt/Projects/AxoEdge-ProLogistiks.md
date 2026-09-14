# AxoEdge ProLogistiks: Evaluating Computer Vision vs. Multimodal LLMs for Pallet Counting

## Technologies and languages used on this project

Languages

- C#
- SQL
- Markdown
- JSON
- PowerShell

Technologies

- .NET 10
- ML.NET
- ONNX Runtime
- YOLO-based object detection
- Ollama
- OllamaSharp
- SkiaSharp
- Npgsql
- PostgreSQL

## Short Version

After my layoff from DocuSign, I co-founded a software company. Our first client, a German/Swiss logistics company, needed to count pallets from warehouse photos for inventory tracking, and computer vision was completely new to me. I benchmarked seven open-source multimodal LLMs against a labeled dataset and built YOLO-based detection baselines. The data was decisive: the top two LLMs (MiniCPM and LLaMA 3.2) had 40% count deviation at 10 seconds per image—too slow and unreliable—while YOLO achieved ~80% accuracy with proper training data. I delivered a go/no-go recommendation in one month, backed by reproducible experiments and a documented handoff repo.

---

## Polished STAR++ Narrative

### Situation

After being laid off from DocuSign, I co-founded a small software consulting company. Our first engagement came from a German/Swiss logistics firm (proLogistik) that needed to automatically count pallets from warehouse photos to track inventory ingress and egress. Computer vision and deep learning were entirely new domains for me, and the hype around multimodal LLMs made the architecture choice genuinely non-obvious.

### Task

I owned the full technical evaluation: determine whether accurate pallet counting was feasible, compare open-source multimodal LLMs against object detection, and produce an evidence-based architecture recommendation and effort estimate—within one month, with no prior baseline to anchor expectations.

### Actions

- **Parallel-track ramp-up**: Rather than evaluating approaches sequentially, I split work into two concurrent tracks—(1) learning object detection workflows and YOLO-family models, (2) surveying the open-source multimodal LLM ecosystem (LLaVA, LLaMA 3.2, Gemma, Granite, Qwen, Moondream, Mistral). Running both in parallel let me make a data-driven comparison instead of arguing from intuition.
- **Benchmark harness design**: I built a structured test bed that ran all 7 LLMs across multiple prompt strategies against a labeled pallet dataset, persisted every result in PostgreSQL, and analyzed count deviation and failure modes per model. This produced reproducible evidence rather than one-off impressions.
- **Object detection baselines**: I implemented Tiny YOLO and YOLO pipelines end-to-end—including data format conversion and preprocessing—to establish a real performance baseline, not just rely on published benchmarks.
- **Structured for handoff**: I documented the full investigation in a GitHub repo (experimental setup, model evaluations, key CV/ML concepts, troubleshooting) so a second engineer could own or extend the work without depending on me.

### Result

The benchmarks were unambiguous: the two best LLMs (MiniCPM and LLaMA 3.2) showed 40% count deviation and required ~10 seconds per image—both disqualifying for production inventory tracking. YOLO-based detection reached ~80% accuracy with appropriate training data and was orders of magnitude faster. I delivered a clear recommendation to invest in training a custom vision model on a representative dataset, giving the client quantitative evidence for a confident go/no-go decision within one month.

### Reflection & Tradeoffs

I ruled out multimodal LLMs based on zero-shot and few-shot performance. A reasonable alternative would have been to explore fine-tuning, but the client had no labeled fine-tuning corpus and needed reliability over flexibility—object detection was the defensible choice. In retrospect, I would have agreed on a minimum-accuracy threshold with the client upfront so the evaluation criteria were externally anchored, not just my judgment. This project cemented my approach to unfamiliar domains: run parallel experiments, let data decide, and document for scale.

---

## Quick Pass/Fail Checks

- Situation: yes
- Task: yes
- Actions: yes
- Result: yes
- Reflection: yes

---

## STAR++ Score

**Score: 8/10** — Strong decision framework, real numbers, reproducible methodology, and clear tradeoffs; docked slightly because the outcome is still pending and client scale/business stakes are underspecified.

---

## Strengths

- Data-driven architecture decision with real benchmark numbers (40% LLM deviation vs. ~80% YOLO accuracy), not intuition
- Parallel-track approach demonstrates efficiency under time constraint
- Knowledge-transfer built into the work (GitHub repo, PostgreSQL-persisted results)
- Explicit tradeoff reasoning: why fine-tuning LLMs was ruled out given client constraints

---

## Weaknesses & Concrete Improvements

- **Weak result framing** → The outcome is "still under evaluation," which reads as inconclusive. Reframe the result around the quality of the recommendation itself: "Eliminated an unreliable architecture before any production commitment" is a concrete outcome even without a signed contract.
- **Client stakes are vague** → Add one line on business impact (e.g., "Miscounting pallets directly affects inventory reconciliation across [est. dozens of warehouse locations]") to give the problem real weight.
- **No rollout or monitoring plan mentioned** → Interviewers may probe: "How would you have validated accuracy in production?" Prepare a short answer about a confidence threshold, a human-review fallback, and a retraining pipeline.

---

## One-line Rewrite Suggestion

"I benchmarked 7 open-source multimodal LLMs against YOLO baselines, found the best LLMs had 40% count deviation at 10s/image vs. ~80% YOLO accuracy, and delivered a data-backed recommendation that prevented committing to an unreliable architecture—in under a month."

---

## Time-to-Deliver Check

**Trimmable to ~2 minutes? Yes** — In live delivery, collapse the Actions to two bullets: (1) benchmark harness + results, (2) YOLO baseline. Drop the handoff detail unless asked. Target 180–200 words spoken.

---

## Two Quick Tips to Increase Technical Signal

- **Quantify the latency constraint**: Say "10 seconds per image was a non-starter for real-time inventory scanning at warehouse intake," which connects the raw number to the business requirement.
- **Name the failure mode, not just the variance**: Instead of "40% deviation," say "LLMs frequently miscounted partially occluded pallets and returned inconsistent answers to identical prompts"—this shows you understood *why* they failed, not just *that* they failed.

---

## Suggested Interviewer Follow-ups

- **Q1**: "How did you decide which 7 models to include in your benchmark, and how did you handle models you couldn't run locally?"
- **Q2**: "YOLO requires labeled training data—how would you advise the client to build that dataset cost-effectively?"
- **Q3**: "What would you change if the client later wanted to deploy this in a low-connectivity warehouse environment?"

---

## Metadata

- **Role**: Co-founder / Lead Engineer, AxoEdge (consulting)
- **Timeframe**: ~1 month
- **Metrics provided**: yes (40% LLM count deviation, ~10s/image latency, ~80% YOLO accuracy)
- **Estimates used**: no
