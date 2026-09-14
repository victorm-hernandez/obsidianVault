
# Results of evaluation points

## Evaluation Prompt

This project is a take home assessment for a software engineering interview. 

I want you to answer the following questions about the project.

- How will the application handle large spikes in traffic?
- What kind of observability exists, and is it enough to debug production issues and 
aid in capacity planning?

## Traffic spikes handling

#### Limitations / Current Vulnerabilities under Heavy Traffic:

- **No Ingress Rate Limiting**: There is no rate limiter middleware on `GET /`. If traffic completely drains both the channel buffer and the circular log cache, incoming requests will receive `408 Request Timeout` responses.
- **Fixed Worker Pool**: The generator pool size (`ClientCount`) is static at startup and does not auto-scale based on channel drain rate.
- **Single-Instance In-Memory State**: Both the channel and circular log live in local process memory. Horizontal scaling across multiple pods requires an external distributed cache (e.g., Redis).

---
## Observability for debugging and capacity planning

What kind of observability exists, and is it enough to debug production issues and aid in capacity planning?

#### Production Debugging Gaps:

- **No Correlation / Request IDs**: Downstream HTTP requests and incoming requests lack distributed tracing headers or request IDs (`X-Request-ID` / OpenTelemetry context). It is difficult to trace a failed user request to a specific downstream weather API call across logs.
- **Uninstrumented Incoming HTTP Handlers**: `InstrumentedRoundTripper` only measures _outgoing_ client calls. There is no middleware measuring incoming request rates, error distributions (4xx/5xx rates), or handler latency on `GET /`.

#### Capacity Planning Gaps:

- **No Standard Metrics Exporter**: The stats endpoint returns custom JSON rather than Prometheus metrics format (`/metrics`), making it hard to integrate with standard monitoring tools (Prometheus, Datadog, Grafana).
- **Missing Percentile Latencies (P50, P90, P99)**: `InstrumentedRoundTripper` computes standard rolling averages, which hide tail-latency spikes critical for SLA monitoring and capacity sizing.
- **Missing Cache & Buffer Saturation Metrics**: There are no metrics tracking channel buffer depletion frequency, circular log cache hit/miss rates, or cache staleness over time under varying RPS load.

---

## Security 

What steps have been taken to secure the service? Are there mitigations to common 
application security issues?

## Resiliency to external failures

How does the application respond if external service dependencies misbehave?

## Deployment

How will the application be deployed?

## Other production ready considerations

What else does production ready mean to you?


- Expiration of forecast is part of the response, this is ignored.