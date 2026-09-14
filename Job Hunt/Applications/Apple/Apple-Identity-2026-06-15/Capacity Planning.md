Our system today is I/O bound but could be changed to CPU bound with a simple change. 

It is I/O bound since we rely on the throughput of the underlying Weather APIs, even if we keep a cache around, we only use it after attempting to reach from the Weather APIs first. 

---
## Theory

Start with Little's Law: 
		- Throughput = Concurrency / Latency
		- RPS (Request per second) = concurrent thread count / time to process a single request
		- Concurrency = For CPU bound operations is the number of processors for IO bound operations 

Our Latency: 
	Local processing = ~100 ms
	Loc API = ~ 4sec
	Metadata API = ~450 ms
	Forecast API = ~450 ms

Because you have your Latency ($5$ seconds), you are missing one variable to solve the equation. You have to approach this from one of two directions: calculating your **Required Concurrency** based on a business goal, or calculating your **Maximum Concurrency** based on system limits.

### Calculating Required Concurrency (Goal-Driven)

Let's say your goal is a throughput of **200 Requests Per Second (RPS)**.

$$\text{Concurrency} = 200 \text{ req/sec} \times 5 \text{ sec}$$

$$\text{Concurrency} = 1,000$$

To achieve 200 RPS with a 5-second latency, your application architecture must be capable of holding **1,000 concurrent requests** simultaneously.

### Calculating Maximum Concurrency (Resource-Driven)

Memory constraints
- How much memory each request uses, how much memory do you have may determine the ceiling of the number of concurrent requests.

Thread Constraints
- How many threads can your environment spin up? This may determine the ceiling of your max concurrency.

Socket Constraints
- How many I/O requests can you make without exhausting parent OS resources?

Underlying dependency connection limits
- What is the throughput of the underlying APIs? What is the max number of concurrent connections?

### Other interesting data points
atomic operations = 20 ns
channel operations = 100 ns

