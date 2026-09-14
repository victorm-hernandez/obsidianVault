IMPORTANT WITHOUT SOLUTION:
Data freshness: 
- Data should be evicted from the channel continously to keep it fresh.


Problems to solve:
- Rate limit requests to third party
- Prepare for spikes in traffic
- Resiliency if third party is down

Scalability considerations:
Given that there is no input parameters to the API we can pre calculate a buffer of responses.
- TODO: No Circuit Breaker / Bulkhead Pattern for Weather APIs.

Security considerations TODOs:
- Sanitize name string from the random location API before including it on the response.
- Validate forecast URL  returned by the points API is in a whitelist before proceeding.
- Sanitize the forecast string from the weather API before including it on the response.

Testing TODOs:
- Poor test coverage
- Missing E2E test cases
- Missing negative cases

Observability considerations:
- Better Instrument requests to third party APIs to keep track of performance and health.

No clean shutdown:
No handling of os signals to shutdown the server.