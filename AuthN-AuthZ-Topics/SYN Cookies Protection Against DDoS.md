
To understand how **SYN cookies** protect against Denial of Service (DoS) attacks, it helps to look at the exact mechanics of how the Linux kernel handles incoming TCP connections before any file descriptor (FD) is ever created.

### What is SYN?

**SYN** stands for **Synchronize**. It is a 1-bit control flag located within the TCP header. Its primary purpose is to initiate a connection synchronization sequence between two networking endpoints, signaling that the sender wants to establish a communication channel and agree on starting parameters.

### What is a SYN Packet?

A **SYN packet** is simply a TCP segment where the SYN flag is set to `1` (and the ACK flag is `0`).

It represents **Step 1** of the classic TCP **3-way handshake**:

1. **Client $\to$ Server:** The client sends a **SYN packet**, proposing an Initial Sequence Number (ISN) so both sides can track packet order and detect missing data.
    
2. **Server $\to$ Client:** The server replies with a **SYN-ACK packet** (acknowledging the client's sequence number and providing its own server-side ISN).
    
3. **Client $\to$ Server:** The client sends a final **ACK packet** confirming the server's sequence number.

### 1. The Normal TCP Handshake and Backlog Queues

When a client initiates a TCP connection, the operating system kernel manages two distinct queues in memory:

- **The Half-Open Queue (SYN Queue):** When the server receives a client's initial `SYN` packet, it replies with a `SYN-ACK` and creates a temporary entry in the SYN queue to track the handshake. **At this stage, no application socket or file descriptor exists yet.**
    
- **The Fully Established Queue (Accept Queue):** Once the client responds with the final `ACK`, the connection is fully established, moved to the accept queue, and handed off to the application via `accept()`—which is when the kernel finally allocates a file descriptor.

### 2. Why is the Half-Open Queue Needed?

The **Half-Open Queue (SYN Queue)** is required because TCP is a **stateful, reliable connection-oriented protocol** that operates over unreliable networks. The queue exists for three core architectural reasons:

- **Sequence Number Tracking:** Before data can flow securely and in order, both hosts must record each other's Initial Sequence Numbers. The server needs temporary storage to hold the client's proposed sequence number while it waits to receive the final handshake confirmation.
    
- **Network Latency and Dropped Packets:** Because the internet is asynchronous, the time between the server sending a `SYN-ACK` and receiving the client's final `ACK` is variable. The server must maintain a holding area for these "in-flight" handshakes so it knows a connection is actively being formed rather than assuming it failed immediately.
    
- **Resource Isolation (Shielding the Application):** By keeping half-open connections in a dedicated kernel queue, the operating system prevents unverified connection attempts from polluting the application's actual connection pool or spawning process threads/file descriptors before the client has proven it actually exists and is reachable.
- 
### 3. The Vulnerability: SYN-Flooding

An attacker exploits the Half-Open Queue by flooding the server with `SYN` packets using **spoofed source IP addresses**.

- The server responds with `SYN-ACK` packets to non-existent or unreachable IPs, filling up its SYN queue with half-open states.
- Because the SYN queue has a fixed size (governed by `somaxconn` and backlog limits), a high-rate flood quickly exhausts it.
- Once the queue is full, **legitimate incoming connection requests are dropped**, locking out real users even though the application server and its file descriptor limits are completely untouched.
    
### 4. How SYN Cookies Defeat the Attack

When SYN cookies are enabled (e.g., via `net.ipv4.tcp_syncookies = 1` in Linux), the kernel changes how it handles a full SYN queue:

- **Stateless Operation:** When the SYN queue fills up, the kernel **stops allocating memory for incoming SYN packets**.
    
- **Cryptographic Encoding:** Instead of storing state, the kernel encodes the connection details—such as source/destination IP, ports, a secret server key, and a timestamp—into a cryptographic hash. This hash is sent back to the client as the initial sequence number in the `SYN-ACK` packet.
    
- **Verification on ACK:** The server completely forgets the connection attempt and frees up any pending overhead. If the client is legitimate, it replies with an `ACK` containing that sequence number.
    
- **Late Allocation:** The kernel recalculates the cryptographic hash using the incoming packet's details. If it matches, the kernel verifies the client is real, creates the socket structure, allocates a file descriptor, and hands the connection over to the application.
    

### Summary

SYN cookies act as a stateless cryptographic shield. They allow the server to handshake with legitimate clients under heavy attack without dedicating memory or queue space to unverified connections, preventing resource exhaustion before an FD is ever in the picture.