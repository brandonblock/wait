# wait

Implementing this package in your microservice's main.go (via NewWaiter) with routes for StoreEvent and SetLimit will allow you to limit actions to once per specified duration.

Waiter creates a channel for requests and loads them when StoreEvent is called. The consumer function only handles the request body upon receiving a tick from the <-limiter channel.

If this were to be implemented in a distributed fashion, the limiter channel would be replaced by checking a cache (redis, etc) and the "tick" would be a token expiration. Every second the cache would expire a lease, one of the waiters would "renew" the lease by setting another cache value with an n-second ttl but the pattern would remain the same.
