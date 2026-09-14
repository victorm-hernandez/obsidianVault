# Observability for debugging and capacity planning

## Requirements

We need to cover the fundamentals:

- Logs, timed events with context about the event.
- Metrics, health and performance indicators
- Traces, visibility across the different systems of any given request. Traces (entire flow) and Spans (segments of the flow)

The following information should be recorded as part of the lifecycle of our service.

1 Generation
    1.1 Third party requests
        1.1.1 Performance stats. How long it takes for each succesful request (P99, P95, P80, Average)? Group by request type (fetch random, get URL, get forecast), group by domain.
        1.2.1 Reliability stats. What is the status code distribution across requests? How many times do we need to retry (P99, P95, P80)? When we retry, what is the success rate after retries (P99, P95, P80)? Group by request type (fetch random, get URL, get forecast), group by domain.
    1.2 Buffer status
        1.2.0 Client usage, is the number of clients optimal? what is the usage of the clients at any given times, do we have all 50 running? how else to assess this?
        1.2.1 Boot insights, how long it takes to boot? how many errors happen during boot?, time for each request?
        1.2.2 Depletion/Replenish stats, how fast is the buffer consumed, how often is it empty? how long it takes to replenish again?
        1.2.3 Aging rate (expiration stats), how often a record in the buffer is not used since it was expired.
    1.3 Cache status
        1.3.1 Usage. How many times is the cache used over time?
        1.3.2 Eviction. How many times the records on the cache are expired?
        1.3.3 Reuse. How often a record is seen more than once by a request (cache full loop)
        1.3.4 Reuse with same client. How many times the same client receives the same data due to a cache full loop?

2. Serving data
    2.1 Usage stats. What is the usage of our API (RPS P99, P80, P50, P10)? At what time do we have traffic spikes, how does the stats look on that time range? 
    2.2 Performance stats. What is the response time for succesful requests? How does it change over time based on RPS.
    2.3 Reliability stats. What are the non Success HTTP status code returned and how many. What is the reason for the failure? What are the top k reasons for failure?

Correlation IDs
- Incoming Client ID (thumbprint)
- Incoming Request ID
- Fecth ID (client-id-timeFirstRequest)
- Forecast ID (generatorId-latt-long-time)
- Error ID (per originating location, action performed, Ex.: parsing JSON response from third party request) [package ID-ErrorID]

Debug level logs:
- Configure SLog to include call stacks.
- Additional logging upon failure: Detailed request, response data for both incoming and outgoing requests.

## Observability Plan

### Share correlation IDs

We will use a context object to share some correlation IDs:

- During incoming request we will include client ID and request ID. Inserted by the middleware into the context.
- During outgoing request we will include fetch ID (clientId-time of first request). Inserted by the generator into the context.

### Instrumented Generator

Instrumented Generator, decorator on top of the Generator, responsible for:

- Consumption, how fast is data consumed, reads per second. 
- Generation, how fast data is generated, writes per second, write duration. What is the fill rate of the buffer over time? How often is the buffer empty?
- Correlation IDs used: forecast ID is part of the generated forecast [clientId-lat-long-time].

### Fingerprint struct 

Algorithm per scenario:
- If behind a load balancer, read the headers injected by the load balancer which performs the fingerprint on our behalf.
- Read the TLS connection information (capture the hello request) and create a hash (J3).
- If non HTTP, read the HTTP headers, create a hash with them to generate a fingerprint.

- Additional checks to identify bots: look for inconsistencies on the headers, order, information sent by each browser.

### Middleware 

Logging Middleware responsible for:

- Correlation Ids generated and used: Request ID [req-timeInNano] and calling the fingerprint package to generate client ID correlation IDs [from cookie, or if not found, save cookie as hash of user agent, IP]
- Logging request start, duration. Response is not here since it is not accesible.
- [Debug enabled] Record the entire HTTP request.

### Handlers

Location handler responsible for:
- Logging any error processing the data
- Log the information source (cache, buffer) and the specific information served (forecast ID)
- Log the HTTP status response
- [Debug enabled] Record the entire HTTP response.
- Correlation IDs used: client ID (fingerprint) and request IDs, forecast ID.

## Open Telemetry Strategy

Metrics
    Counters
        HttpStatusCodes
            Code
            Count
        RequestCount
        ExpiredCacheItemsCount

    Histograms
        IncomingRequestPerSecond
        OutgoingRequestsPerSecond
        DependenciesResponseTime
        ServerResponseTime
        BufferLength
        ActiveFetchClients

    Gauges
        BufferLength (0-Max)
        CacheLength (0-Max)
        ActiveFetchClients (0-ClientMax)
        
        
        MAYBE IRRELEVANT AT THIS LEVEL
        CacheOldest (timestamp) Does it have expired records already?
        CacheNewest (timestamp)


Traces
    Trace - Location Generation
        Span - Fetch Random Location (our server)
            Attribute: ClientID, WithRetries, With500, With409 (too many requests)
            Events: requests for every third party, include retries and error data.

            Span - Random Location (generated on the remote server)
            Span - URL retrieval (generated on the remote server)
            Span - Forecast retrieval (generated on the remote server)

QUESTION: When an event is attached to a span, how to determine level, info, warn, debug?
QUESTION: What about duplicate instruments, gauges and histograms, is this a bad practice?
QUESTION: What should I use to slice my data based on HTTP status code, an attribute (With429/WithRetries) or an event or error? 
