# Final Loop

Discuss the following topics in relation to the take home exercise:

Session 1: Scale + Behavioral Topic
Session 2: Operational Readiness + Behavioral Topic

Operational excellence.
    - Monitoring
    - Observability
    - Unity testing

These last two interviews will be one discussion about live DB migration and the other is a coding session where you will refactor the unfinished take home exercise below.

Session 3: Live Database Migration + Behavioral Topic
Conversation how you would move something from DB 1 to DB 2, how you decommission, do you need to cache something, how do you move this in flight.

Session 4: Refactoring coding exercise + Behavioral Topic
There is two features, understading the start code base, have clear communication, interviewer will give hints and recommendations.
How do you make the code better.

## Questions recruiter

Session 1:
    Is it based on the take home assessment?
    Behavioral also related with it? or project.

Feedback about home assesment:
    - Solution suggested an unfamiliarity with some of the golang concepts. For example:
        -> Did add request timeout projects but not in the idiomatic way using change context
        -> Scaling problems: no cache used, but I did! maybe is not idiomatic
        -> Raw equality tests instead of asserts unfamiliarity with common testing patterns.
        -> 4 applied, only I passed.

    - Behavioral questions:
        -> Collaboration and technical leadership (being vulnerable, work with others). Document light company, all the information is in the people. All the access is via relationships.
        How you foster trust?
        How you fight for quality?
        How you help people?
        How you get to know people at work?
        Time you were able to do this?
        Times you broke trust and how you fixed it?
        During the interview: Be open to mistakes. Are you willing to recognize if you make a mistake. 
        Be authentic: be authentic self. no packaged answers. 

        -> Operated at scale and collaboration and partnership.

        I did super well, compared with experts on go. 

Overview of final interview:

## Plan

- P0 What is the use case?
    We assume one of the following scenarios:
        - Test data generation
        - Dummy information as a filler for a demonstration
        - Entertainment? learn about the weather at any arbitrary location.

- Random location is the slowest unreliable API, the rest is fast.
    - Generate a static list of locations on the US using the unreliable API offline and deploy it with the application as a data source.

- Other optimizations
    - Use the "Expires" header from the response as our own expiration.
    - Add the Retry-After header to our own responses.

- Observer pattern in go
- Decorator pattern in go ( it seems the pipeline pattern is this one) forwarding one channel to another and do something in the middle.

- How to measure how long it takes to fill the buffer?

- What do we need for capacity planning?
    - How can we plan for different throughputs, loads?

- Go architectural patterns and how to apply them to my flow.
    - Bulk Head / Circuit Breaker
    - Hexagonal architecture and how to apply it to current project.

- How to expose information in your health endpoint such that load balancers can consider this?

- Use a service Mesh with my app

- Implement SPIFFE?

- DONE Implement the "Struct-First, Validated Configuration pattern"
- DONE Export the goroutine profile via HTTP in your application.
- DONE Test different configurations to increase throughput by not having so many retries.
- Expose health endpoint such that it integrates with K8.
- Expose metrics endpoint such that it can be consumed by prometheus.
- Check design patterns in go, see what would change if implemented an hexagonal architecture on this project. 

- Capture my own global metrics (atomic counters)
- Prepare the application to be deployed using docker

- Measure how much time each part of the request takes, why the best time we have is 5 secs.

- Detect go routine leaks with test cases.
- Dynamic client and rqs for the generator based on demand.
    - Implement worker and job pattern, that way we can increase the number of workers
    - RPS can be per client not as a whole

- Improve circle log logic, the length part in specific, make sure the logic is clean. There is a potential bug on it when the elements are expired.

- Cache just the random part of the fetcher, this is the only one that shown 429 so far.
- Tests for the handler
- Use generics on the circular log so the type can be changed?

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
