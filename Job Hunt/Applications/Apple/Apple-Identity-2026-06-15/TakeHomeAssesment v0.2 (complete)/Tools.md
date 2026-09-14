# Tools reference

## Load tsting

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
hey -n 50000 -c 200 http://localhost:8888/
```

## Optimized build

```bash
# Linux
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o server main.go
# MacOS
CGO_ENABLED=0 GOOS=darwin GOARCH=arm64 go build -ldflags="-s -w" -o prod_server main.go
```

Breakdown of Flags

- CGO_ENABLED=0: Disables cgo. Creates a statically linked binary that runs without external C library dependencies.
- `-ldflags="-s -w"`: Strips debugging information. 
- `-s` removes the symbol table. 
- `-w` removes DWARF debugging symbols. This reduces binary size by roughly 30% to 40%.
- `GOOS=linux GOARCH=amd64`: Target environment overrides. Change these to match your production server architecture (e.g., linux and arm64 for AWS Graviton).


## Mac OS adjustments

### File descriptor limit update

Increase the limits in your terminal session.

``` bash
ulimit -n 65535
```


