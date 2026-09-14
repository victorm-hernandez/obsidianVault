# Goals/Evaluation

We ultimately want to deploy this application in production, which implies additional 
concerns beyond just returning correct results to a simple curl call. When we evaluate the 
production readiness of a service, we think about questions like:

## Traffic Spikes

How will the application handle large spikes in traffic?

The application employs an **asynchronous pre-generation and multi-tiered caching strategy** designed to decouple client read requests from slow, rate-limited external weather APIs.

#### Key Architectural Mechanisms:

1. **Asynchronous Background Pre-generation (Buffered Channel)**
    
    - Because the service returns random location forecasts without requiring user input parameters, location forecasts are generated in the background prior to requests. 
    - RandomLocationGenerator runs a configurable worker pool (`ClientCount`) rate-limited by `MaxFetchPerSec`.
    - Forecasts are pushed into a thread-safe, buffered channel (`locDataChan`) of size `BufferSize`.
    - When a user hits `GET /` in GetWeatherHandler, the handler pops a pre-fetched forecast from the channel in $O(1)$ constant time without waiting for external network calls.

2. **Graceful Degradation via Thread-Safe Circular Log Cache**

    - If a traffic spike exhausts the buffered channel or if external weather APIs slow down, the handler waits up to `WriteTimeout` (configured in config.go).
    - Upon timeout, 
        GetWeatherHandler falls back to reading from LocCircularLog.
    - The circular log retains a rolling set of non-expired location records. It cycles sequentially through cached responses so that even under heavy burst load, users receive a valid `200 OK` response with slightly older data rather than a server error or failure.
3. **Connection Pooling & Timeout Protection**
    
    - In cmd/api/main.go, explicit timeouts (`ReadHeaderTimeout`, `ReadTimeout`, `WriteTimeout`, `IdleTimeout`) are set on `http.Server` to protect against file descriptor exhaustion and slow-client DoS attacks.
    - Outgoing HTTP client transports reuse persistent TCP/HTTP2 connections (`MaxIdleConns`, `MaxIdleConnsPerHost`) to avoid opening fresh sockets for every background fetch.
    
---

## Observability for debugging and capacity planning

What kind of observability exists, and is it enough to debug production issues and 
aid in capacity planning?

What Observability Currently Exists:

1. **Structured JSON Logging (`log/slog`)**
    
    - Configured in cmd/api/main.go using standard `log/slog` with a JSON output handler (`os.Stdout`).
    - Logs include structured attributes like `HttpStatusCode`, `FromCache`, `ClientCount`, and `Error`.

2. **Custom Health & Metrics Endpoint (`GET /stats`)**
    
    - GetStatsHandler exposes operational state as JSON via  
        api.HealthResponse:
        - **Generator Metrics**: `ready` status, `bootStartTime`, `bootDurationMs` (time taken to populate the buffer at startup), and `itemCount`.
        - **Outgoing Downstream HTTP Metrics**: Intercepted by InstrumentedRoundTripper, tracking request counts, minimum duration (`DurationFloor`), maximum duration (`DurationCeiling`), and rolling average latency (`DurationAverage`) grouped by downstream domain and HTTP status code.

3. **Go Runtime Profiling (`net/http/pprof`)**
    
    - Enabled via RunProfilingTools. Configures `runtime.SetBlockProfileRate` and `runtime.SetMutexProfileFraction` to diagnose goroutine blocking, lock contention, memory allocations, and goroutine leaks.

## Security

What steps have been taken to secure the service? Are there mitigations to common 
application security issues?

## Resiliency to external failures

How does the application respond if external service dependencies misbehave?

## Deployment
How will the application be deployed?

## Other production ready considerations

What else does production ready mean to you