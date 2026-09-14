# Final Loop

Discuss the following topics in relation to the take home exercise:

Session 1: Scale + Behavioral Topic
Session 2: Operational Readiness + Behavioral Topic

These last two interviews will be one discussion about live DB migration and the other is a coding session where you will refactor the unfinished take home exercise below.

Session 3: Live Database Migration + Behavioral Topic
Session 4: Refactoring coding exercise + Behavioral Topic

## Pending tasks

### Deployment

- Add support for configuration files, all settings are hard coded at this point.
- Add endpoints to assess the health of the service.
- Configure a CI/CD pipeline to push changes from our project to the different environments (test, integration, prod)
- Gracefully shutdown our server.
- Create standard operating guides (SOP) for deployment.
- Troubleshooting guides (TSGs) for livesite operations and incidents.
- Create a docker image for our service.

### Security

- Sanitize the output of our third party APIs before using it in our customers/logs.
- Validate the URLs returned against an allow list of domains.
- Limit the request size and also throttle incoming requests to avoid abuse, we may benefit from a middleware for this purpose.

### Observability

- There are no correlation IDs (Request IDs) propagated to downstream calls which is very useful on livesite investigations.
- Regarding capacity planning, we require metrics, better instrumentation, explicitely capture response times from underlying APIs and from our handlers as well.
- Middleware capturing performance and health for every handler.
- A metrics endpoint.
- Use prometheus in our metrics.

### Testing

- Run Go routine leak profiler in our API
- Create tests for every component.
- Create tests that validate the concurrent behavior of our componetns

### Functional

- Make sure that if the loc info channel is iddle for a while its refreshed. After 10 minutes of inactivity, start refreshing the entire buffer.
- Add a circuit breaker to avoid calling the underlying APIs if they are down.
