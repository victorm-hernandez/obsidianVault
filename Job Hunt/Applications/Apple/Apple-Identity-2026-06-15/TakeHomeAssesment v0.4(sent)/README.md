# Random Weather Forecast Service

## How to Compile/Test/Run this Service

The followign command will build the service:

```bash
go build -race ./...
```

The following command will execute the underlying tests:

```bash
go test -race ./...
```

The following command will start the service:

```bash
go run cmd/api/main.go
```

To see the random location forecast in action load the following URL on a browser `http://localhost:8888/`.

> IMPORTANT!: The port number used in the final implementation (8888) differs from that of the requirements (5000) due to collisions in MacOS.

---

## Goals/Evaluation

### How will the application handle large spikes in traffic?

My strategy to handle spikes in traffic is to pre-generate a buffer of forecast information required for our responses. The service doesn't require any inputs from the consumer, this makes it a prime candidate to generate forecast info ahead of time.

We will create a **Forecast Generator** that will continously replenish this buffer as it is consumed by our API handlers (using a shared buffered channel).

A **trade off** of our generator is that we will use memory that potentially may or may not be used to improve our overall performance and success rate. The data in the third party weather API seems to be updated 2 to 4 times a day (every 6-12 hours). If we fetch our data just before it changes, we risk serving stale data. To mitigate this risk we will expire the generated records after 10 minutes.

### What kind of observability exists, and is it enough to debug production issues and aid in capacity planning?

**What we have:**

- We have structured logs for the main scenarios, providing useful information to debug production issues.

**What is missing:**

- There are no correlation IDs (Request IDs) propagated to downstream calls which is very useful on livesite investigations.
- Regarding capacity planning, we require metrics, better instrumentation, explicitely capture response times from underlying APIs and from our handlers as well.
- Middleware capturing performance and health for every handler.
- A metrics endpoint.

### What steps have been taken to secure the service? Are there mitigations to common application security issues?

**What we have:** We protect against DOS attacks via file descriptor exhaustion by proactively defining timeouts for TCP connections to our server, outgoing HTTP requests and handlers.

**What is missing:** We need to sanitize the output of our third party APIs before using it in our customers/logs. We should validate the URLs returned against an allow list of domains. We should limit the request size and also throttle incoming requests to avoid abuse, we may benefit from a middleware for this purpose.

### How does the application respond if external service dependencies misbehave?

In all cases we will attempt to use fresh data from the third party service but if their throughput is not enough or if they fail and we risk timing out/failing our request, we will attempt to use a cache (if not expired) to produce a graceful degraded experience.

In practice, the underlying API's throttle limits are very low, which in turn limit our own throughtput. To mitigate this problem I introduced a type a cache tailored for this scenario, a thread safe circular log. This log will hold a copy of the most recent subset of location data produced. This log will be persistent, the data on it will remain available as long it hasn't expired yet, the data will be read sequentially (in cycles) to reduce the chance of repeated responses.

### How will the application be deployed?

There is not much work done towards deployment at this stage. For a rounded project we should:

- Add support for configuration files, all settings are hard coded at this point.
- Add endpoints to assess the health of the service.
- Configure a CI/CD pipeline to push changes from our project to the different environments (test, integration, prod)
- Gracefully shutdown our server.
- Create standard operating guides (SOP) for deployment.
- Troubleshooting guides (TSGs) for livesite operations and incidents.

---

## Additional work to make this project "production ready"

- Add a circuit breaker to avoid bombarding the Weather API service if its down.
- Plenty of testing is missing, increase the test coverage.
