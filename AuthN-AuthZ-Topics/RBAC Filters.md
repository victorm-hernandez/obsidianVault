## Definition and Architecture

An **RBAC (Role-Based Access Control) network filter** is an enforcement mechanism embedded directly into the data path of a proxy, API gateway, or service mesh (typically operating at Layer 4 or Layer 7 of the OSI model). Unlike traditional static IP/port firewalls or application-level middleware, an RBAC network filter intercepts raw connections or HTTP streams, extracts caller identity and connection metadata, and evaluates them against defined role bindings and authorization policies _before_ traffic reaches upstream workloads.

### How It Works

1. **Interception:** The filter intercepts incoming packets or requests at the network edge or service boundary.
    
2. **Context Extraction:** It gathers attributes such as:
    
    - **Cryptographic Identity:** SPIFFE IDs or Subject Alternative Names (SAN) extracted from mTLS certificates.
    - **Token Claims:** JWT payloads (scopes, roles, issuer) for L7 filters.
    - **Transport Metadata:** Source IP/CIDR blocks, SNI (Server Name Indication), destination ports, and header values.
        
3. **Policy Evaluation:** The engine checks these attributes against a set of configured `principals`, `permissions`, and `actions` (allow or deny).
4. **Enforcement:** If the evaluation succeeds, the traffic is forwarded; if it fails, the connection is terminated or an HTTP `403 Forbidden` is returned immediately.

## Examples of Use

- **Zero-Trust Service-to-Service Authorization:** In a Kubernetes microservices architecture, a payment service needs to restrict access to its database proxy. An L4/L7 RBAC filter ensures that only workloads presenting a SPIFFE ID matching `spiffe://cluster.local/ns/prod/sa/payment-processor` can establish a TCP connection or execute specific RPC methods.
    
- **Admin Ingress Segmentation:** At an API gateway, an L7 RBAC filter inspects incoming JWT claims. Public endpoints are accessible to any caller, but requests targeting `/admin/*` require the principal to possess the `PlatformAdmin` role claim, blocking external callers even if they can resolve the DNS record.
    
- **Multi-Tenant Cluster Isolation:** Enforcing tenant-level boundaries at the ingress proxy where incoming traffic headers (`X-Tenant-ID`) or mTLS client certificates dictate whether a client role is authorized to communicate with a specific tenant's namespace endpoints.

## Technologies Implementing RBAC Network Filters

- **Envoy Proxy:** Provides native filters including `envoy.filters.network.rbac` (for L4 TCP connections) and `envoy.filters.http.rbac` (for L7 HTTP requests). It evaluates policies based on matchers like `source_ip`, `destination_port`, `authenticated_client_identity`, and header values.
    
- **Istio Service Mesh:** Builds on Envoy's RBAC filtering capabilities via `AuthorizationPolicy` custom resources, allowing operators to define fine-grained namespace-wide, workload-wide, or mesh-wide access control using mTLS identities and JWT claims.
    
- **Cilium:** Implements identity-aware network security at Layer 3 to Layer 7 using eBPF. It maps Kubernetes service accounts to cryptographic security identities, allowing network filters to enforce RBAC rules directly in the Linux kernel data path without userspace proxy overhead.
    
- **Kong Gateway:** Utilizes plugins (such as the Kong ACL or enterprise RBAC plugins) to intercept upstream requests at the proxy layer, validating consumer credentials and assigned groups against route-level permissions.

## Istio / Envoy Integration

Istio uses **Envoy** as its fundamental data plane proxy architecture.

#### How the Integration Works

- **Sidecar Deployment:** When you inject Istio into a Kubernetes pod, it deploys an Envoy proxy container (`istio-proxy`) alongside your application container.
- **Traffic Interception:** Envoy intercepts all inbound and outbound traffic for the pod using `iptables` or eBPF redirection rules
- **Control Plane to Data Plane:** Istio's control plane component (`istiod`) translates high-level configurations—such as `AuthorizationPolicy`, `DestinationRule`, and `VirtualService`—into low-level Envoy configuration via the dynamic **xDS API** (Discovery Service).

When you configure an Istio `AuthorizationPolicy` to enforce RBAC, Istio compiles those rules directly into Envoy's native L4 and L7 RBAC filter configurations (`envoy.filters.network.rbac` and `envoy.filters.http.rbac`), ensuring high-performance enforcement directly in the Envoy data path.