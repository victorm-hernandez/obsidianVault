**Important performance improvement:** This requires business logic discussion, do we need to call the underlying APIs as much as possible or is it acceptable to rely on a cache?

 As part of the handler logic we wait for the channel to have elements in it, there is no need for that as the channel main purpose after adding the cache is just client synchronization, we can read directly from the circular cache, which would make our requests immediate and our concurrency CPU bound instead of I/O bound.

**Important improvement**: There is no circuit breaker on the system to avoid bombarding the third party API if there are multiple failures.

**Important Bug:** The fetcher ignores the cache policy HTTP headers which should be used for the forecast expiration. 

**Minor Bug**: Circular cache length protection leaves the reader index at an arbitrary position after all of the elements are expired, it should be left at the write position.

**Very Minor bug**: Write to cache is not synchronized with write to channel, you write to the cache even if the channel is full, there is a slight mismatch on the each data store