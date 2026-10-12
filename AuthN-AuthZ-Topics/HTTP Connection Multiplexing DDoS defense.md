**Connection multiplexing** is a core architectural mechanism that allows multiple independent application-layer requests and responses to share a **single underlying network connection** simultaneously, rather than requiring a dedicated connection for each request.

### 1. The HTTP/1.1 Bottleneck (Why Connection Sprawl Happens)
In older protocols like HTTP/1.1, each TCP connection could only handle **one request at a time**. If a client needed to fetch multiple resources or fire concurrent API requests, it had two bad options:

- **Sequential Execution:** Wait for request A to finish before starting request B (unacceptably slow).
- **Connection Sprawl:** Open dozens of parallel TCP connections to the same server.

Every single one of those parallel connections forces the operating system kernel to allocate a distinct socket, memory buffers, and a **file descriptor (FD)**. Under high traffic or a distributed botnet attack, this connection sprawl rapidly exhausts the server's file descriptor limits, causing legitimate requests to be dropped.
### 2. How HTTP/2 and HTTP/3 Multiplexing Work

HTTP/2 and HTTP/3 eliminate connection sprawl by introducing **streams**:

- **Stream Framing:** Both protocols break down HTTP messages into independent, interleaved binary frames. Each frame carries a unique **Stream ID**.
    
- **Interleaving over One Socket:** Multiple requests (e.g., fetching a JSON payload, submitting a form, loading an asset) are sliced into frames, mixed together over a _single_ TCP connection (HTTP/2) or UDP/QUIC connection (HTTP/3), and reassembled on the receiving end.
    
- **The FD Impact:** Because hundreds of concurrent logical requests share the exact same transport socket, the server only consumes **one file descriptor** for that entire client session.
    
### 3. The Role of Connection Pooling

While HTTP/2 and HTTP/3 multiplexing manage concurrency _per connection_, **connection pooling** manages reuse _across clients and backend services_.

- **Persistent Sockets:** Instead of opening and closing a TCP connection for every database query, microservice call, or client API request, services maintain a pool of warm, pre-established connections.
    
- **Throttled Resource Consumption:** Pooling enforces a strict upper bound on how many total sockets your application can open. Combined with multiplexing, a pool of just a few connections can handle a massive volume of concurrent throughput without spiking file descriptor usage.

### 4. How This Defends Against DoS

- **Resource Preservation:** By reducing the total number of open sockets required to serve traffic, your server's file descriptor table remains well below its exhaustion ceiling (`EMFILE`).
    
- **Mitigating Slow-Rate Attacks:** Attackers cannot easily drain your server pool by opening thousands of idle connections when a single multiplexed connection handles high concurrency.
    
- **Stream-Level Flow Control:** HTTP/2 and HTTP/3 include built-in flow control per stream, allowing servers to cancel or throttle individual malicious or stalled streams without tearing down the entire underlying connection.

### Verification Check

You can verify your server's current multiplexing efficiency and active file descriptor usage under load using standard Linux utilities:

- Check open file descriptors for your service process ID:
    
    `ls -l /proc/<PID>/fd | wc -l`
    
- Inspect active HTTP/2 stream multiplexing ratios in your edge proxy or load balancer metrics (e.g., measuring `connections_active` versus `requests_active` in Envoy or NGINX).