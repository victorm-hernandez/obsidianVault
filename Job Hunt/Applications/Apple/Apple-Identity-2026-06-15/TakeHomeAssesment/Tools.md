# Tools reference

## pprof

[pprof](https://github.com/google/pprof/blob/main/doc/README.md) is a tool for visualization and analysis of profiling data. pprof reads a collection of profiling samples in profile.proto format and generates reports.

### How to export enable profiling on your application

Spin a new server with the pproof package cohosted with your application

```go
import _ "net/http/pprof"
import "net/http"

// Run this right before your main server starts
go func() {
    logger.Info("Starting pprof server on :8889")
    if err := http.ListenAndServe("localhost:8889", nil); err != nil {
        logger.Error("pprof server failed", slog.Any("Error", err))
    }
}()
```

### Visualize the profiling information

You may use a full web report, a quick web page with a snapshot of the data, or the CLI to parse the data into graphs.

### Web Report

You may load this page on your browser to get a web report `http://localhost:8889/debug/pprof/`

#### Quick debug on browser

To simply get a snapshot of the go routines on your browser:

- Go routines `http://localhost:8889/debug/pprof/goroutine?debug=1`
- Heap (Allocations) `http://localhost:8889/debug/pprof/goroutine?debug=1`
- Mutext `http://localhost:8889/debug/pprof/goroutine?debug=1`

#### Use the CLI

Fetch the profile with the pprof tool and then use commands to parse it

```bash
# Render a web application instead of the CLI using a .prof file downloaded
go tool pprof -http=:8080 heap.prof 
# Snapshot of the go routines
go tool pprof http://localhost:8889/debug/pprof/goroutine
# Heap (Allocations)
go tool pprof http://localhost:8889/debug/pprof/heap
# Mutex
go tool pprof http://localhost:8889/debug/pprof/mutex
# 10 second profile of the application
go tool pprof http://localhost:8889/debug/pprof/profile?seconds=10
```

Commands once the data is loaded on the CLI

```bash
# Shows the top elemetns that have the most memory (heap), most mutex locks (mutex), most go routines (goroutine)
(pprof) top
# Alternative you can define how many
(pprof) top20
# Focus on a specific process by name/term
(pprof) focus=TakeHomeAssessment
# Generate PDF report
(pprof) pdf
# Generate a web report (requires installing graphviz)
(pprof) web
```

To run the web report, you need to install the pre-requisite graphviz with this command:

```bash
brew install graphviz
```

To learn more about how to read this graph check the [documentation https://git.io/JfYMW](https://github.com/google/pprof/blob/main/doc/README.md#interpreting-the-callgraph).

[Video showing how to use the CLI](https://youtu.be/S45XZMcXgrc)

---

## Load testing

[Hey Load testing tool](https://github.com/rakyll/hey)

```bash
brew install hey
```

Run tests with following command:

```bash
# Options:
#   -n  Number of requests to run. Default is 200.
#   -c  Number of workers to run concurrently. Total number of requests cannot
#       be smaller than the concurrency level. Default is 50.
#   -q  Rate limit, in queries per second (QPS) per worker. Default is no rate limit.
#   -z  Duration of application to send requests. When duration is reached,
#       application stops and exits. If duration is specified, n is ignored.
#       Examples: -z 10s -z 3m.
#   -o  Output type. If none provided, a summary is printed.
#       "csv" is the only supported alternative. Dumps the response
#       metrics in comma-separated values format.

#   -m  HTTP method, one of GET, POST, PUT, DELETE, HEAD, OPTIONS.
#   -H  Custom HTTP header. You can specify as many as needed by repeating the flag.
#       For example, -H "Accept: text/html" -H "Content-Type: application/xml" .
#   -t  Timeout for each request in seconds. Default is 20, use 0 for infinite.
#   -A  HTTP Accept header.
#   -d  HTTP request body.
#   -D  HTTP request body from file. For example, /home/user/file.txt or ./file.txt.
#   -T  Content-type, defaults to "text/html".
#   -a  Basic authentication, username:password.
#   -x  HTTP Proxy address as host:port.
#   -h2 Enable HTTP/2.

#   -host	HTTP Host header.

#   -disable-compression  Disable compression.
#   -disable-keepalive    Disable keep-alive, prevents re-use of TCP
#                         connections between different HTTP requests.
#   -disable-redirects    Disable following of HTTP redirects
hey -n 5000 -c 50 http://localhost:8888/
```

---

## Optimized build

```bash
# Linux
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o output cmd/api/main.go
# MacOS
CGO_ENABLED=0 GOOS=darwin GOARCH=arm64 go build -ldflags="-s -w" -o output cmd/api/main.go
```

```powershell
# Mac OS Target
$env:CGO_ENABLED="0"; $env:GOOS="darwin"; $env:GOARCH="arm64"; go build -ldflags="-s -w" -o output cmd/api/main.go
```

Breakdown of Flags

- CGO_ENABLED=0: Disables cgo. Creates a statically linked binary that runs without external C library dependencies.
- `-ldflags="-s -w"`: Strips debugging information. 
- `-s` removes the symbol table. 
- `-w` removes DWARF debugging symbols. This reduces binary size by roughly 30% to 40%.
- `GOOS=linux GOARCH=amd64`: Target environment overrides. Change these to match your production server architecture (e.g., linux and arm64 for AWS Graviton).

---

## Mac OS adjustments

### File descriptor limit update

Increase the limits in your terminal session.

``` bash
ulimit -n 65535
```

---

## Leak test

```bash
# Terminal 1, Run web server
go run cmd/api/main.go
# Terminal 2, Run load test
hey -n 5000 -c 50 http://localhost:8888/
# Terminal 3, 1 Minute profile CPU report from the CLI
go tool pprof http://localhost:8889/debug/pprof/profile?seconds=60
```

[Web Report](http://localhost:8889/debug/pprof/)


Int64 max value
9,223,372,036,854,775,807
1000 req per sec * 60 secs * 60 min * 24 hrs = 8,640,0000 req per day. 
106,751,991,167.30064591 days = 292,471,208 years