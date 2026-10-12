Here are the 7 levels of the OSI (Open Systems Interconnection) model

==**A**ll **P**eople **S**eem **T**o **N**eed **D**ata **P**rocessing==

1.  **Physical Layer** – Transmits raw bit streams (electrical voltages, light pulses, or radio waves) over a physical medium (e.g., cables, hubs, repeaters).
2. **Data Link Layer** – Provides node-to-node data transfer, physical addressing (MAC addresses), and error detection/correction over a single physical link (e.g., Ethernet, Wi-Fi/802.11).
3. **Network Layer** – Handles packet routing, forwarding, and logical addressing across multiple interconnected networks (e.g., IP, ICMP, IPsec).
4. **Transport Layer** – Ensures reliable, sequential, and error-checked delivery of data packets between end systems (e.g., TCP, UDP).
5. **Session Layer** – Manages and maintains communication sessions (establishing, maintaining, and terminating connections) between applications across a network.
6. **Presentation Layer** – Translates, encrypts, and compresses data so that application layers across different systems can understand each other (e.g., SSL/TLS, JPEG, ASCII).
7. **Application Layer** – Directly serves the end user and application processes, handling communication protocols and data exchange (e.g., HTTP, FTP, SMTP, DNS).
    
_(A common mnemonic to remember the order from Layer 7 down to Layer 1 is: **A**ll **P**eople **S**eem **T**o **N**eed **D**ata **P**rocessing)._

7 - Application (HTTP/HTTPS)
6 - Presentation (TLS, )
5 - Session
4 - Transport (UDP/TCP)
3 - Network (IP)
2 - Data (Ethernet)
1 -Physical (cables, repeaters)

---
## Load balancers on the OSI model

Load balancers operate at different levels of the OSI model depending on their architecture and capability. They are typically categorized into two main types:

- **Layer 4 (Transport Layer) Load Balancers** – These operate on network transmission data without inspecting the actual application content. They route traffic based on IP addresses and TCP/UDP ports (e.g., forwarding all traffic hitting port 443 to a backend pool). They are extremely fast and efficient because they don't terminate or parse application payloads.
    
- **Layer 7 (Application Layer) Load Balancers** – These operate at the highest level of the OSI model, acting as a reverse proxy. They can inspect the actual application traffic—such as HTTP/HTTPS headers, cookies, URLs, and SSL/TLS certificates—allowing for content-based routing (e.g., routing `/api/*` to one set of servers and `/images/*` to another).

---
## Level 5 - Session, concepts

### What is a session at the end of the day?

At its core, **a connection or session is fundamentally an allocated block of state shared between two endpoints.**

While the OSI model treats the Session Layer (Layer 5) as a neat abstraction for managing dialogues (establishing, maintaining, and synchronizing interactions), physical and software engineering reality dictates that a "connection" is nothing more than data residing in memory, data structures, or backing stores.

### What State Actually Lives on Either Side?

Depending on whether you are looking at the Transport/TLS layer or the Application/Session layer, that persisted state typically includes:

- **OS and Network Buffers:** File descriptors, socket structures, and TCP Control Blocks (TCBs) tracking sequence numbers, acknowledgement numbers, and congestion windows.
    
- **Cryptographic Context:** Negotiated cipher suites, symmetric encryption keys, or TLS session tickets (which allow session resumption without a full handshake).
    
- **Application and Identity Context:** Associated user IDs, roles, permissions, transaction states, or temporary tokens cached in memory (e.g., in a Redis cluster or local process memory).
    
- **Dialogue Trackers:** State machines indicating whether the conversation is waiting for a request, processing, streaming, or shutting down.
    
### The Lifecycle: Allocation, Persistence, and Teardown

1. **Allocation (The Handshake):** When a connection is opened, both ends allocate resources. The server allocates a memory footprint (e.g., a slot in a connection pool or an entry in a session table) to map incoming packets to a specific context.
    
2. **Persistence (The Lifetime):** Throughout the connection's life, that state is maintained. If it's a stateful protocol or a sticky session, the server relies on that stored context to route subsequent requests without re-authenticating or re-negotiating.
    
3. **Discarding (The Teardown):** Once the connection is closed, timed out, or explicitly terminated (via a FIN/RST packet or an explicit logout/session destroy call), that state is **purged and reclaimed**.
    
    - Memory is freed, file descriptors are released, cryptographic keys are wiped from active memory, and session caches or database rows are deleted (or marked for garbage collection).
        

### The Modern Architectural Twist

In modern distributed systems and microservices architectures, the classic OSI Session Layer is often heavily decoupled.

Because tying a user's session strictly to a single server's in-memory state breaks horizontal scaling and load balancing, modern systems separate the **transport connection** (ephemeral TCP/TLS state that dies when the socket drops) from the **session state** (externalized into distributed caches like Redis, or encoded into cryptographically signed stateless tokens like JWTs).

### File Descriptors (session state persisted at the OS level)

#### What Are File Descriptors?

At a systems level, a **file descriptor (FD)** is an abstract, non-negative integer handle (an index in the kernel's per-process file descriptor table) that an operating system kernel uses to track and manage open I/O resources.

In Unix-like operating systems, **everything is a file**—including regular files, directories, pipes, anonymous memory, and crucially, **network sockets**. Every time a distributed service accepts an incoming TCP connection or opens an outbound client connection, the kernel allocates a new file descriptor to represent that socket.

#### How File Descriptors Enable Denial of Service (DoS)

Because every open connection consumes an FD, systems have a hard ceiling on concurrent connections determined by OS limits (`ulimit -n`) and physical memory. This creates a primary vector for **Resource Exhaustion DoS attacks**:

- **Socket Flooding:** An attacker or botnet opens thousands of TCP connections simultaneously without closing them. Once the process or kernel hits its maximum FD limit, legitimate incoming requests are rejected with errors like `EMFILE` (Too many open files) or `ENFILE` (File table overflow).
    
- **Slow-Rate Attacks (e.g., Slowloris):** Attackers open connections and send HTTP headers extremely slowly, byte by byte. The server holds the file descriptors and memory buffers open waiting for the request to complete, tying up resources until the pool is completely drained.

### Protecting Distributed Systems Against FD Exhaustion

Mitigating FD-based DoS requires a defense-in-depth strategy spanning the operating system, edge infrastructure, and application code:

#### 1. Operating System and Kernel Tuning

- **Increase Limits (`ulimit`/`sysctl`):** Raise hard and soft limits for open files (`nofile`) for service processes, and tune kernel parameters like `fs.file-max` to allow a higher ceiling of total system-wide open files.
    
- **SYN Cookies and Backlog Queues:** Enable TCP SYN cookies to protect the kernel's half-open connection queue from SYN-flood attacks before an FD is even allocated.

#### 2. Edge and Load Balancer Offloading

- **Edge Rate Limiting:** Terminate and filter malicious traffic at the API gateway or load balancer (e.g., Envoy, NGINX, Cloudflare) before it hits your core backend services.
    
- **Connection Multiplexing:** Use connection pooling and HTTP/2 or HTTP/3 multiplexing so multiple logical requests share a single underlying TCP connection (and thus a single file descriptor).

#### 3. Aggressive Timeouts and Lifecycle Controls

- **Short Read/Write Timeouts:** Implement strict idle, read, and write timeouts at the socket layer. If a client stops communicating or transmits too slowly, forcibly drop the connection and reclaim the FD.
    
- **Connection Shedding:** When a service detects high resource utilization (e.g., FD usage crossing 80%), proactively drop incoming connections or return `533 Service Unavailable` before exhausting the table.

#### 4. Application-Level Backpressure and Bounded Queues

- **Bounded Concurrency:** Use non-blocking I/O event loops (such as Linux `epoll`, Go's netpoller, or Java NIO) combined with bounded worker pools to reject or queue excess load gracefully rather than spawning unbounded threads/sockets.

---
### What is a Socket?

A **socket** is an abstract software endpoint that enables two-way communication between two programs running on a network (or on the same machine).

Think of it as a virtual telephone jack or a designated mailing slot managed by the operating system kernel. When an application wants to send or receive data over a network, it doesn't talk directly to the network interface card or raw network packets; instead, it opens a socket, binds it to an IP address and port, and reads/writes data to it just like it would a standard file.

### Is a Socket a File Descriptor?

**Yes, in Unix-like operating systems (Linux, macOS, BSD), a network socket is represented as a file descriptor.**

This is a direct implementation of the classic Unix philosophy that **"everything is a file."** When your application creates a TCP or UDP socket, the kernel allocates a file descriptor (an integer handle, such as `3` or `12`) to track that socket.

Because a socket is a file descriptor, your program can use standard system calls like `read()`, `write()`, `poll()`, `select()`, and `close()` on a network connection in the exact same way it would interact with a text file or hardware device on disk. (Windows uses a similar abstraction called a _socket handle_).

### The Anatomy of a Socket
A network socket is formally defined by a combination of software constructs and network parameters:

1. **The Transport Protocol:** The underlying protocol governing how data moves (typically TCP for reliable, connection-oriented streams, or UDP for fast, connectionless datagrams).
    
2. **The Socket Address Tuple:** A unique combination of **Source IP, Source Port, Destination IP, and Destination Port** (known as a 4-tuple for TCP) that identifies a specific active connection across the network.
    
3. **Kernel I/O Buffers:** Memory spaces allocated by the operating system kernel to hold incoming and outgoing data packets while the application processes them.